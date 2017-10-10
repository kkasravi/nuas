/*
Copyright (c) 2016-2017 Intel Corporation.

Redistribution.  Redistribution and use in binary form, without
modification, are permitted provided that the following conditions are
met:
  Redistributions must reproduce the above copyright notice and the
  following disclaimer in the documentation and/or other materials
  provided with the distribution.
  Neither the name of Intel Corporation nor the names of its suppliers
  may be used to endorse or promote products derived from this software
  without specific prior written permission.
  No reverse engineering, decompilation, or disassembly of this software
  is permitted.

Limited patent license.  Intel Corporation grants a world-wide,
royalty-free, non-exclusive license under patents it now or hereafter
owns or controls to make, have made, use, import, offer to sell and
sell ("Utilize") this software, but solely to the extent that any
such patent is necessary to Utilize the software alone, or in
combination with an operating system licensed under an approved Open
Source license as listed by the Open Source Initiative at
http://opensource.org/licenses.  The patent license shall not apply to
any other combinations which include this software.  No hardware per
se is licensed hereunder.

DISCLAIMER.  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING,
BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR
TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE
USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
DAMAGE.
*/
package v1

import (
	v1 "github.com/nervanasystems/nuas/pkg/apis/helium/v1"
	scheme "github.com/nervanasystems/nuas/pkg/client/clientset_generated/clientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JobtypesGetter has a method to return a JobtypeInterface.
// A group's client should implement this interface.
type JobtypesGetter interface {
	Jobtypes(namespace string) JobtypeInterface
}

// JobtypeInterface has methods to work with Jobtype resources.
type JobtypeInterface interface {
	Create(*v1.Jobtype) (*v1.Jobtype, error)
	Update(*v1.Jobtype) (*v1.Jobtype, error)
	UpdateStatus(*v1.Jobtype) (*v1.Jobtype, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Jobtype, error)
	List(opts meta_v1.ListOptions) (*v1.JobtypeList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Jobtype, err error)
	JobtypeExpansion
}

// jobtypes implements JobtypeInterface
type jobtypes struct {
	client rest.Interface
	ns     string
}

// newJobtypes returns a Jobtypes
func newJobtypes(c *HeliumV1Client, namespace string) *jobtypes {
	return &jobtypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a jobtype and creates it.  Returns the server's representation of the jobtype, and an error, if there is any.
func (c *jobtypes) Create(jobtype *v1.Jobtype) (result *v1.Jobtype, err error) {
	result = &v1.Jobtype{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jobtypes").
		Body(jobtype).
		Do().
		Into(result)
	return
}

// Update takes the representation of a jobtype and updates it. Returns the server's representation of the jobtype, and an error, if there is any.
func (c *jobtypes) Update(jobtype *v1.Jobtype) (result *v1.Jobtype, err error) {
	result = &v1.Jobtype{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jobtypes").
		Name(jobtype.Name).
		Body(jobtype).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *jobtypes) UpdateStatus(jobtype *v1.Jobtype) (result *v1.Jobtype, err error) {
	result = &v1.Jobtype{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jobtypes").
		Name(jobtype.Name).
		SubResource("status").
		Body(jobtype).
		Do().
		Into(result)
	return
}

// Delete takes name of the jobtype and deletes it. Returns an error if one occurs.
func (c *jobtypes) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jobtypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jobtypes) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jobtypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the jobtype, and returns the corresponding jobtype object, and an error if there is any.
func (c *jobtypes) Get(name string, options meta_v1.GetOptions) (result *v1.Jobtype, err error) {
	result = &v1.Jobtype{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jobtypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Jobtypes that match those selectors.
func (c *jobtypes) List(opts meta_v1.ListOptions) (result *v1.JobtypeList, err error) {
	result = &v1.JobtypeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jobtypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jobtypes.
func (c *jobtypes) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jobtypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched jobtype.
func (c *jobtypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Jobtype, err error) {
	result = &v1.Jobtype{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jobtypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
