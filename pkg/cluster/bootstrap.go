package cluster

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/rancher/k3s/pkg/bootstrap"
	"github.com/rancher/k3s/pkg/clientaccess"
	"github.com/rancher/k3s/pkg/version"
	"github.com/sirupsen/logrus"
)

// Bootstrap attempts to load a managed database driver, if one has been initialized or should be created/joined.
// It then checks to see if the cluster needs to load boostrap data, and if so, loads data into the
// ControlRuntimeBoostrap struct, either via HTTP or from the datastore.
func (c *Cluster) Bootstrap(ctx context.Context) error {
	if err := c.assignManagedDriver(ctx); err != nil {
		return err
	}

	runBootstrap, err := c.shouldBootstrapLoad(ctx)
	if err != nil {
		return err
	}
	c.shouldBootstrap = runBootstrap

	if runBootstrap {
		if err := c.bootstrap(ctx); err != nil {
			return err
		}
	}

	return nil
}

// shouldBootstrapLoad returns true if we need to load ControlRuntimeBootstrap data again.
// This is controlled by a stamp file on disk that records successful bootstrap using a hash of the join token.
func (c *Cluster) shouldBootstrapLoad(ctx context.Context) (bool, error) {
	// Non-nil managedDB indicates that the database is either initialized, initializing, or joining
	if c.managedDB != nil {
		c.runtime.HTTPBootstrap = true

		isInitialized, err := c.managedDB.IsInitialized(ctx, c.config)
		if err != nil {
			return false, err
		}

		if isInitialized {
			// If the database is initialized we skip bootstrapping; if the user wants to rejoin a
			// cluster they need to delete the database.
			logrus.Infof("Managed %s cluster bootstrap already complete and initialized", c.managedDB.EndpointName())
			return false, nil
		} else if c.config.JoinURL == "" {
			// Not initialized, not joining - must be initializing (cluster-init)
			logrus.Infof("Managed %s cluster initializing", c.managedDB.EndpointName())
			return false, nil
		} else {
			// Not initialized, but have a Join URL - fail if there's no token; if there is then validate it.
			if c.config.Token == "" {
				return false, errors.New(version.ProgramUpper + "_TOKEN is required to join a cluster")
			}

			token, err := clientaccess.NormalizeAndValidateTokenForUser(c.config.JoinURL, c.config.Token, "server")
			if err != nil {
				return false, err
			}

			info, err := clientaccess.ParseAndValidateToken(c.config.JoinURL, token)
			if err != nil {
				return false, err
			}

			logrus.Infof("Managed %s cluster not yet initialized", c.managedDB.EndpointName())
			c.clientAccessInfo = info
		}
	}

	// Check the stamp file to see if we have successfully bootstrapped using this token.
	// NOTE: The fact that we use a hash of the token to generate the stamp
	//       means that it is unsafe to use the same token for multiple clusters.
	stamp := c.bootstrapStamp()
	if _, err := os.Stat(stamp); err == nil {
		logrus.Info("Cluster bootstrap already complete")
		return false, nil
	}

	// No errors and no bootstrap stamp, need to bootstrap.
	return true, nil
}

// bootstrapped touches a file to indicate that bootstrap has been completed.
func (c *Cluster) bootstrapped() error {
	if err := os.MkdirAll(filepath.Dir(c.bootstrapStamp()), 0700); err != nil {
		return err
	}

	if _, err := os.Stat(c.bootstrapStamp()); err == nil {
		return nil
	}

	f, err := os.Create(c.bootstrapStamp())
	if err != nil {
		return err
	}

	return f.Close()
}

// httpBootstrap retrieves bootstrap data (certs and keys, etc) from the remote server via HTTP
// and loads it into the ControlRuntimeBootstrap struct. Unlike the storage bootstrap path,
// this data does not need to be decrypted since it is generated on-demand by an existing server.
func (c *Cluster) httpBootstrap() error {
	content, err := clientaccess.Get("/v1-"+version.Program+"/server-bootstrap", c.clientAccessInfo)
	if err != nil {
		return err
	}

	return bootstrap.Read(bytes.NewBuffer(content), &c.runtime.ControlRuntimeBootstrap)
}

// bootstrap performs cluster bootstrapping, either via HTTP (for managed databases) or direct load from datastore.
func (c *Cluster) bootstrap(ctx context.Context) error {
	c.joining = true

	// bootstrap managed database via HTTP
	if c.runtime.HTTPBootstrap {
		return c.httpBootstrap()
	}

	// Bootstrap directly from datastore
	return c.storageBootstrap(ctx)
}

// bootstrapStamp returns the path to a file in datadir/db that is used to record
// that a cluster has been joined. The filename is based on a portion of the sha256 hash of the token.
// We hash the token value exactly as it is provided by the user, NOT the normalized version.
func (c *Cluster) bootstrapStamp() string {
	return filepath.Join(c.config.DataDir, "db/joined-"+keyHash(c.config.Token))
}
