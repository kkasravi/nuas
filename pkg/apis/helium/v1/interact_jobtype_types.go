
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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nervanasystems/nuas/pkg/apis/helium"
)

// +genclient=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type Interact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

var _ rest.CreaterUpdater = &InteractJobtypeREST{}
var _ rest.Patcher = &InteractJobtypeREST{}

// +k8s:deepcopy-gen=false
type InteractJobtypeREST struct {
	Registry helium.JobtypeRegistry
}

func (r *InteractJobtypeREST) Create(ctx request.Context, obj runtime.Object, includeUninitialized bool) (runtime.Object, error) {
	sub := obj.(*Interact)
	rec, err := r.Registry.GetJobtype(ctx, sub.Name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	// Modify rec in someway before writing it back to storage

	r.Registry.UpdateJobtype(ctx, rec)
	return rec, nil
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *InteractJobtypeREST) Get(ctx request.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return nil, nil
}

// Update alters the status subset of an object.
func (r *InteractJobtypeREST) Update(ctx request.Context, name string, objInfo rest.UpdatedObjectInfo) (runtime.Object, bool, error) {
	return nil, false, nil
}

func (r *InteractJobtypeREST) New() runtime.Object {
	return &Interact{}
}

