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

// StreamingsGetter has a method to return a StreamingInterface.
// A group's client should implement this interface.
type StreamingsGetter interface {
	Streamings(namespace string) StreamingInterface
}

// StreamingInterface has methods to work with Streaming resources.
type StreamingInterface interface {
	Create(*v1.Streaming) (*v1.Streaming, error)
	Update(*v1.Streaming) (*v1.Streaming, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Streaming, error)
	List(opts meta_v1.ListOptions) (*v1.StreamingList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Streaming, err error)
	StreamingExpansion
}

// streamings implements StreamingInterface
type streamings struct {
	client rest.Interface
	ns     string
}

// newStreamings returns a Streamings
func newStreamings(c *HeliumV1Client, namespace string) *streamings {
	return &streamings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a streaming and creates it.  Returns the server's representation of the streaming, and an error, if there is any.
func (c *streamings) Create(streaming *v1.Streaming) (result *v1.Streaming, err error) {
	result = &v1.Streaming{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("streamings").
		Body(streaming).
		Do().
		Into(result)
	return
}

// Update takes the representation of a streaming and updates it. Returns the server's representation of the streaming, and an error, if there is any.
func (c *streamings) Update(streaming *v1.Streaming) (result *v1.Streaming, err error) {
	result = &v1.Streaming{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("streamings").
		Name(streaming.Name).
		Body(streaming).
		Do().
		Into(result)
	return
}

// Delete takes name of the streaming and deletes it. Returns an error if one occurs.
func (c *streamings) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("streamings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *streamings) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("streamings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the streaming, and returns the corresponding streaming object, and an error if there is any.
func (c *streamings) Get(name string, options meta_v1.GetOptions) (result *v1.Streaming, err error) {
	result = &v1.Streaming{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("streamings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Streamings that match those selectors.
func (c *streamings) List(opts meta_v1.ListOptions) (result *v1.StreamingList, err error) {
	result = &v1.StreamingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("streamings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested streamings.
func (c *streamings) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("streamings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched streaming.
func (c *streamings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Streaming, err error) {
	result = &v1.Streaming{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("streamings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
