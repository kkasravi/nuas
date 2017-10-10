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

// InteractLister helps list Interacts.
type InteractLister interface {
	// List lists all Interacts in the indexer.
	List(selector labels.Selector) (ret []*v1.Interact, err error)
	// Interacts returns an object that can list and get Interacts.
	Interacts(namespace string) InteractNamespaceLister
	InteractListerExpansion
}

// interactLister implements the InteractLister interface.
type interactLister struct {
	indexer cache.Indexer
}

// NewInteractLister returns a new InteractLister.
func NewInteractLister(indexer cache.Indexer) InteractLister {
	return &interactLister{indexer: indexer}
}

// List lists all Interacts in the indexer.
func (s *interactLister) List(selector labels.Selector) (ret []*v1.Interact, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Interact))
	})
	return ret, err
}

// Interacts returns an object that can list and get Interacts.
func (s *interactLister) Interacts(namespace string) InteractNamespaceLister {
	return interactNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InteractNamespaceLister helps list and get Interacts.
type InteractNamespaceLister interface {
	// List lists all Interacts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Interact, err error)
	// Get retrieves the Interact from the indexer for a given namespace and name.
	Get(name string) (*v1.Interact, error)
	InteractNamespaceListerExpansion
}

// interactNamespaceLister implements the InteractNamespaceLister
// interface.
type interactNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Interacts in the indexer for a given namespace.
func (s interactNamespaceLister) List(selector labels.Selector) (ret []*v1.Interact, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Interact))
	})
	return ret, err
}

// Get retrieves the Interact from the indexer for a given namespace and name.
func (s interactNamespaceLister) Get(name string) (*v1.Interact, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("interact"), name)
	}
	return obj.(*v1.Interact), nil
}
