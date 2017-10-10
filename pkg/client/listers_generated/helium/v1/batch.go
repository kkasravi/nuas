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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/nervanasystems/nuas/pkg/apis/helium/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BatchLister helps list Batchs.
type BatchLister interface {
	// List lists all Batchs in the indexer.
	List(selector labels.Selector) (ret []*v1.Batch, err error)
	// Batchs returns an object that can list and get Batchs.
	Batchs(namespace string) BatchNamespaceLister
	BatchListerExpansion
}

// batchLister implements the BatchLister interface.
type batchLister struct {
	indexer cache.Indexer
}

// NewBatchLister returns a new BatchLister.
func NewBatchLister(indexer cache.Indexer) BatchLister {
	return &batchLister{indexer: indexer}
}

// List lists all Batchs in the indexer.
func (s *batchLister) List(selector labels.Selector) (ret []*v1.Batch, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Batch))
	})
	return ret, err
}

// Batchs returns an object that can list and get Batchs.
func (s *batchLister) Batchs(namespace string) BatchNamespaceLister {
	return batchNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BatchNamespaceLister helps list and get Batchs.
type BatchNamespaceLister interface {
	// List lists all Batchs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Batch, err error)
	// Get retrieves the Batch from the indexer for a given namespace and name.
	Get(name string) (*v1.Batch, error)
	BatchNamespaceListerExpansion
}

// batchNamespaceLister implements the BatchNamespaceLister
// interface.
type batchNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Batchs in the indexer for a given namespace.
func (s batchNamespaceLister) List(selector labels.Selector) (ret []*v1.Batch, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Batch))
	})
	return ret, err
}

// Get retrieves the Batch from the indexer for a given namespace and name.
func (s batchNamespaceLister) Get(name string) (*v1.Batch, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("batch"), name)
	}
	return obj.(*v1.Batch), nil
}
