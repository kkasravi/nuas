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
package fake

import (
	v1 "github.com/nervanasystems/nuas/pkg/apis/helium/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBatchs implements BatchInterface
type FakeBatchs struct {
	Fake *FakeHeliumV1
	ns   string
}

var batchsResource = schema.GroupVersionResource{Group: "helium.aipg.intel.com", Version: "v1", Resource: "batchs"}

var batchsKind = schema.GroupVersionKind{Group: "helium.aipg.intel.com", Version: "v1", Kind: "Batch"}

func (c *FakeBatchs) Create(batch *v1.Batch) (result *v1.Batch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(batchsResource, c.ns, batch), &v1.Batch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Batch), err
}

func (c *FakeBatchs) Update(batch *v1.Batch) (result *v1.Batch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(batchsResource, c.ns, batch), &v1.Batch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Batch), err
}

func (c *FakeBatchs) Delete(name string, options *meta_v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(batchsResource, c.ns, name), &v1.Batch{})

	return err
}

func (c *FakeBatchs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(batchsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1.BatchList{})
	return err
}

func (c *FakeBatchs) Get(name string, options meta_v1.GetOptions) (result *v1.Batch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(batchsResource, c.ns, name), &v1.Batch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Batch), err
}

func (c *FakeBatchs) List(opts meta_v1.ListOptions) (result *v1.BatchList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(batchsResource, batchsKind, c.ns, opts), &v1.BatchList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.BatchList{}
	for _, item := range obj.(*v1.BatchList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchs.
func (c *FakeBatchs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(batchsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched batch.
func (c *FakeBatchs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Batch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(batchsResource, c.ns, name, data, subresources...), &v1.Batch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Batch), err
}
