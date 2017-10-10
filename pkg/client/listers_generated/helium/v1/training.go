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

// TrainingLister helps list Trainings.
type TrainingLister interface {
	// List lists all Trainings in the indexer.
	List(selector labels.Selector) (ret []*v1.Training, err error)
	// Trainings returns an object that can list and get Trainings.
	Trainings(namespace string) TrainingNamespaceLister
	TrainingListerExpansion
}

// trainingLister implements the TrainingLister interface.
type trainingLister struct {
	indexer cache.Indexer
}

// NewTrainingLister returns a new TrainingLister.
func NewTrainingLister(indexer cache.Indexer) TrainingLister {
	return &trainingLister{indexer: indexer}
}

// List lists all Trainings in the indexer.
func (s *trainingLister) List(selector labels.Selector) (ret []*v1.Training, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Training))
	})
	return ret, err
}

// Trainings returns an object that can list and get Trainings.
func (s *trainingLister) Trainings(namespace string) TrainingNamespaceLister {
	return trainingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TrainingNamespaceLister helps list and get Trainings.
type TrainingNamespaceLister interface {
	// List lists all Trainings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Training, err error)
	// Get retrieves the Training from the indexer for a given namespace and name.
	Get(name string) (*v1.Training, error)
	TrainingNamespaceListerExpansion
}

// trainingNamespaceLister implements the TrainingNamespaceLister
// interface.
type trainingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Trainings in the indexer for a given namespace.
func (s trainingNamespaceLister) List(selector labels.Selector) (ret []*v1.Training, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Training))
	})
	return ret, err
}

// Get retrieves the Training from the indexer for a given namespace and name.
func (s trainingNamespaceLister) Get(name string) (*v1.Training, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("training"), name)
	}
	return obj.(*v1.Training), nil
}
