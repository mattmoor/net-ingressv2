/*
Copyright 2020 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "knative.dev/net-ingressv2/pkg/client/clientset/versioned/scheme"
	v1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
)

// HTTPRoutesGetter has a method to return a HTTPRouteInterface.
// A group's client should implement this interface.
type HTTPRoutesGetter interface {
	HTTPRoutes(namespace string) HTTPRouteInterface
}

// HTTPRouteInterface has methods to work with HTTPRoute resources.
type HTTPRouteInterface interface {
	Create(ctx context.Context, hTTPRoute *v1alpha1.HTTPRoute, opts v1.CreateOptions) (*v1alpha1.HTTPRoute, error)
	Update(ctx context.Context, hTTPRoute *v1alpha1.HTTPRoute, opts v1.UpdateOptions) (*v1alpha1.HTTPRoute, error)
	UpdateStatus(ctx context.Context, hTTPRoute *v1alpha1.HTTPRoute, opts v1.UpdateOptions) (*v1alpha1.HTTPRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HTTPRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HTTPRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPRoute, err error)
	HTTPRouteExpansion
}

// hTTPRoutes implements HTTPRouteInterface
type hTTPRoutes struct {
	client rest.Interface
	ns     string
}

// newHTTPRoutes returns a HTTPRoutes
func newHTTPRoutes(c *NetworkingV1alpha1Client, namespace string) *hTTPRoutes {
	return &hTTPRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPRoute, and returns the corresponding hTTPRoute object, and an error if there is any.
func (c *hTTPRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HTTPRoute, err error) {
	result = &v1alpha1.HTTPRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httproutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPRoutes that match those selectors.
func (c *hTTPRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HTTPRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HTTPRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httproutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPRoutes.
func (c *hTTPRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httproutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hTTPRoute and creates it.  Returns the server's representation of the hTTPRoute, and an error, if there is any.
func (c *hTTPRoutes) Create(ctx context.Context, hTTPRoute *v1alpha1.HTTPRoute, opts v1.CreateOptions) (result *v1alpha1.HTTPRoute, err error) {
	result = &v1alpha1.HTTPRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httproutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPRoute).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hTTPRoute and updates it. Returns the server's representation of the hTTPRoute, and an error, if there is any.
func (c *hTTPRoutes) Update(ctx context.Context, hTTPRoute *v1alpha1.HTTPRoute, opts v1.UpdateOptions) (result *v1alpha1.HTTPRoute, err error) {
	result = &v1alpha1.HTTPRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httproutes").
		Name(hTTPRoute.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPRoute).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hTTPRoutes) UpdateStatus(ctx context.Context, hTTPRoute *v1alpha1.HTTPRoute, opts v1.UpdateOptions) (result *v1alpha1.HTTPRoute, err error) {
	result = &v1alpha1.HTTPRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httproutes").
		Name(hTTPRoute.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPRoute).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hTTPRoute and deletes it. Returns an error if one occurs.
func (c *hTTPRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httproutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httproutes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hTTPRoute.
func (c *hTTPRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPRoute, err error) {
	result = &v1alpha1.HTTPRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httproutes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
