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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/kyverno/kyverno/api/reports/v1"
	scheme "github.com/kyverno/kyverno/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EphemeralReportsGetter has a method to return a EphemeralReportInterface.
// A group's client should implement this interface.
type EphemeralReportsGetter interface {
	EphemeralReports(namespace string) EphemeralReportInterface
}

// EphemeralReportInterface has methods to work with EphemeralReport resources.
type EphemeralReportInterface interface {
	Create(ctx context.Context, ephemeralReport *v1.EphemeralReport, opts metav1.CreateOptions) (*v1.EphemeralReport, error)
	Update(ctx context.Context, ephemeralReport *v1.EphemeralReport, opts metav1.UpdateOptions) (*v1.EphemeralReport, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.EphemeralReport, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.EphemeralReportList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EphemeralReport, err error)
	EphemeralReportExpansion
}

// ephemeralReports implements EphemeralReportInterface
type ephemeralReports struct {
	client rest.Interface
	ns     string
}

// newEphemeralReports returns a EphemeralReports
func newEphemeralReports(c *ReportsV1Client, namespace string) *ephemeralReports {
	return &ephemeralReports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ephemeralReport, and returns the corresponding ephemeralReport object, and an error if there is any.
func (c *ephemeralReports) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EphemeralReport, err error) {
	result = &v1.EphemeralReport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ephemeralreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EphemeralReports that match those selectors.
func (c *ephemeralReports) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EphemeralReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EphemeralReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ephemeralreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ephemeralReports.
func (c *ephemeralReports) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ephemeralreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ephemeralReport and creates it.  Returns the server's representation of the ephemeralReport, and an error, if there is any.
func (c *ephemeralReports) Create(ctx context.Context, ephemeralReport *v1.EphemeralReport, opts metav1.CreateOptions) (result *v1.EphemeralReport, err error) {
	result = &v1.EphemeralReport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ephemeralreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ephemeralReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ephemeralReport and updates it. Returns the server's representation of the ephemeralReport, and an error, if there is any.
func (c *ephemeralReports) Update(ctx context.Context, ephemeralReport *v1.EphemeralReport, opts metav1.UpdateOptions) (result *v1.EphemeralReport, err error) {
	result = &v1.EphemeralReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ephemeralreports").
		Name(ephemeralReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ephemeralReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ephemeralReport and deletes it. Returns an error if one occurs.
func (c *ephemeralReports) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ephemeralreports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ephemeralReports) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ephemeralreports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ephemeralReport.
func (c *ephemeralReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EphemeralReport, err error) {
	result = &v1.EphemeralReport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ephemeralreports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}