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

package internalversion

import (
	helium "github.com/nervanasystems/nuas/pkg/apis/helium"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JobtypeLister helps list Jobtypes.
type JobtypeLister interface {
	// List lists all Jobtypes in the indexer.
	List(selector labels.Selector) (ret []*helium.Jobtype, err error)
	// Jobtypes returns an object that can list and get Jobtypes.
	Jobtypes(namespace string) JobtypeNamespaceLister
	JobtypeListerExpansion
}

// jobtypeLister implements the JobtypeLister interface.
type jobtypeLister struct {
	indexer cache.Indexer
}

// NewJobtypeLister returns a new JobtypeLister.
func NewJobtypeLister(indexer cache.Indexer) JobtypeLister {
	return &jobtypeLister{indexer: indexer}
}

// List lists all Jobtypes in the indexer.
func (s *jobtypeLister) List(selector labels.Selector) (ret []*helium.Jobtype, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*helium.Jobtype))
	})
	return ret, err
}

// Jobtypes returns an object that can list and get Jobtypes.
func (s *jobtypeLister) Jobtypes(namespace string) JobtypeNamespaceLister {
	return jobtypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JobtypeNamespaceLister helps list and get Jobtypes.
type JobtypeNamespaceLister interface {
	// List lists all Jobtypes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*helium.Jobtype, err error)
	// Get retrieves the Jobtype from the indexer for a given namespace and name.
	Get(name string) (*helium.Jobtype, error)
	JobtypeNamespaceListerExpansion
}

// jobtypeNamespaceLister implements the JobtypeNamespaceLister
// interface.
type jobtypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Jobtypes in the indexer for a given namespace.
func (s jobtypeNamespaceLister) List(selector labels.Selector) (ret []*helium.Jobtype, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*helium.Jobtype))
	})
	return ret, err
}

// Get retrieves the Jobtype from the indexer for a given namespace and name.
func (s jobtypeNamespaceLister) Get(name string) (*helium.Jobtype, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(helium.Resource("jobtype"), name)
	}
	return obj.(*helium.Jobtype), nil
}
