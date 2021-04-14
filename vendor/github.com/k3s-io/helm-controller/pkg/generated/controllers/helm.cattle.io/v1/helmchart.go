/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type HelmChartHandler func(string, *v1.HelmChart) (*v1.HelmChart, error)

type HelmChartController interface {
	generic.ControllerMeta
	HelmChartClient

	OnChange(ctx context.Context, name string, sync HelmChartHandler)
	OnRemove(ctx context.Context, name string, sync HelmChartHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() HelmChartCache
}

type HelmChartClient interface {
	Create(*v1.HelmChart) (*v1.HelmChart, error)
	Update(*v1.HelmChart) (*v1.HelmChart, error)
	UpdateStatus(*v1.HelmChart) (*v1.HelmChart, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.HelmChart, error)
	List(namespace string, opts metav1.ListOptions) (*v1.HelmChartList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HelmChart, err error)
}

type HelmChartCache interface {
	Get(namespace, name string) (*v1.HelmChart, error)
	List(namespace string, selector labels.Selector) ([]*v1.HelmChart, error)

	AddIndexer(indexName string, indexer HelmChartIndexer)
	GetByIndex(indexName, key string) ([]*v1.HelmChart, error)
}

type HelmChartIndexer func(obj *v1.HelmChart) ([]string, error)

type helmChartController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewHelmChartController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) HelmChartController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &helmChartController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromHelmChartHandlerToHandler(sync HelmChartHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.HelmChart
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.HelmChart))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *helmChartController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.HelmChart))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateHelmChartDeepCopyOnChange(client HelmChartClient, obj *v1.HelmChart, handler func(obj *v1.HelmChart) (*v1.HelmChart, error)) (*v1.HelmChart, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *helmChartController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *helmChartController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *helmChartController) OnChange(ctx context.Context, name string, sync HelmChartHandler) {
	c.AddGenericHandler(ctx, name, FromHelmChartHandlerToHandler(sync))
}

func (c *helmChartController) OnRemove(ctx context.Context, name string, sync HelmChartHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromHelmChartHandlerToHandler(sync)))
}

func (c *helmChartController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *helmChartController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *helmChartController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *helmChartController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *helmChartController) Cache() HelmChartCache {
	return &helmChartCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *helmChartController) Create(obj *v1.HelmChart) (*v1.HelmChart, error) {
	result := &v1.HelmChart{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *helmChartController) Update(obj *v1.HelmChart) (*v1.HelmChart, error) {
	result := &v1.HelmChart{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *helmChartController) UpdateStatus(obj *v1.HelmChart) (*v1.HelmChart, error) {
	result := &v1.HelmChart{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *helmChartController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *helmChartController) Get(namespace, name string, options metav1.GetOptions) (*v1.HelmChart, error) {
	result := &v1.HelmChart{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *helmChartController) List(namespace string, opts metav1.ListOptions) (*v1.HelmChartList, error) {
	result := &v1.HelmChartList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *helmChartController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *helmChartController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.HelmChart, error) {
	result := &v1.HelmChart{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type helmChartCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *helmChartCache) Get(namespace, name string) (*v1.HelmChart, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.HelmChart), nil
}

func (c *helmChartCache) List(namespace string, selector labels.Selector) (ret []*v1.HelmChart, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HelmChart))
	})

	return ret, err
}

func (c *helmChartCache) AddIndexer(indexName string, indexer HelmChartIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.HelmChart))
		},
	}))
}

func (c *helmChartCache) GetByIndex(indexName, key string) (result []*v1.HelmChart, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.HelmChart, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.HelmChart))
	}
	return result, nil
}

type HelmChartStatusHandler func(obj *v1.HelmChart, status v1.HelmChartStatus) (v1.HelmChartStatus, error)

type HelmChartGeneratingHandler func(obj *v1.HelmChart, status v1.HelmChartStatus) ([]runtime.Object, v1.HelmChartStatus, error)

func RegisterHelmChartStatusHandler(ctx context.Context, controller HelmChartController, condition condition.Cond, name string, handler HelmChartStatusHandler) {
	statusHandler := &helmChartStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromHelmChartHandlerToHandler(statusHandler.sync))
}

func RegisterHelmChartGeneratingHandler(ctx context.Context, controller HelmChartController, apply apply.Apply,
	condition condition.Cond, name string, handler HelmChartGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &helmChartGeneratingHandler{
		HelmChartGeneratingHandler: handler,
		apply:                      apply,
		name:                       name,
		gvk:                        controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterHelmChartStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type helmChartStatusHandler struct {
	client    HelmChartClient
	condition condition.Cond
	handler   HelmChartStatusHandler
}

func (a *helmChartStatusHandler) sync(key string, obj *v1.HelmChart) (*v1.HelmChart, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type helmChartGeneratingHandler struct {
	HelmChartGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *helmChartGeneratingHandler) Remove(key string, obj *v1.HelmChart) (*v1.HelmChart, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.HelmChart{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *helmChartGeneratingHandler) Handle(obj *v1.HelmChart, status v1.HelmChartStatus) (v1.HelmChartStatus, error) {
	objs, newStatus, err := a.HelmChartGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
