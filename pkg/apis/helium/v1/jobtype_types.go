
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
	"log"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"github.com/nervanasystems/nuas/pkg/apis/helium"
)

// +genclient=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Jobtype
// +k8s:openapi-gen=true
// +resource:path=jobtypes,strategy=JobtypeStrategy
// +subresource:request=Batch,path=batch,rest=BatchJobtypeREST
// +subresource:request=Interact,path=interact,rest=InteractJobtypeREST
// +subresource:request=Training,path=training,rest=TrainingJobtypeREST
// +subresource:request=Streaming,path=streaming,rest=StreamingJobtypeREST
type Jobtype struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   JobtypeSpec   `json:"spec,omitempty"`
	Status JobtypeStatus `json:"status,omitempty"`
}

// JobtypeSpec defines the desired state of Jobtype
type JobtypeSpec struct {
	Template *corev1.PodSpec `json:"template,omitempty"`
	IngressSpec extensionsv1beta1.IngressSpec `json:"ingress,omitempty"`
	ServiceSpec corev1.ServiceSpec `json:"service,omitempty"`
	ScaleSpec extensionsv1beta1.ScaleSpec `json:"scale,omitempty"`
	ConfigSpec ConfigSpec `json:"config,omitempty"`
}

type ConfigSpec struct {
	NeonRepoSpec    NeonRepoSpec `json:"neonrepo,omitempty"`
	SecuritySpec    SecuritySpec `json:"security,omitempty"`
	StreamDataSpec  StreamDataSpec `json:"streamdata,omitempty"`
	KryptonRepoSpec KryptonRepoSpec `json:"kryptonrepo,omitempty"`
}

type KryptonRepoSpec struct {
	RepoURL             string `json:"RepoURL"`
	Commit              string `json:"Commit"`
	KryptonImage        string `json:"kryptonImage"`
	KryptonSidecarImage string `json:"kryptonSidecarImage"`
}

type NeonRepoSpec struct {
	RepoUrl string `json:"RepoUrl"`
	Commit  string `json:"Commit"`
}

type StreamDataSpec struct {
	ModelPRM         string `json:"modelPRM"`
	ModelPath        string `json:"modelPath"`
	DatasetPath      string `json:"datasetPath"`
	ExtraFilename    string `json:"extraFilename"`
	CustomCodeURL    string `json:"customCodeURL"`
	CustomCommit     string `json:"customCommit"`
	AWSPath          string `json:"awsPath"`
	AWSDefaultRegion string `json:"aWSDefaultRegion"`
	StreamID         int    `json:"streamID"`
	StreamName       string `json:"streamName"`
}

type SecuritySpec struct {
	PresignedToken string `json:"presignedToken"`
	JWTToken       string `json:"jwtToken"`
}

// JobtypeStatus defines the observed state of JobType
type JobtypeStatus struct {
	State int `json:"state,omitempty"`
}

// Validate checks that an instance of Jobtype is well formed
func (JobtypeStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*helium.Jobtype)
	log.Printf("Validating fields for Jobtype %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Jobtype field values
func (JobtypeSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Jobtype)
	// set default field values here
	log.Printf("Defaulting fields for Jobtype %s\n", obj.Name)
}
