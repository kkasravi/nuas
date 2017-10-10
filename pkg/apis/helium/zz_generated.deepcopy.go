// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package helium

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Batch).DeepCopyInto(out.(*Batch))
			return nil
		}, InType: reflect.TypeOf(&Batch{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BatchList).DeepCopyInto(out.(*BatchList))
			return nil
		}, InType: reflect.TypeOf(&BatchList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ConfigSpec).DeepCopyInto(out.(*ConfigSpec))
			return nil
		}, InType: reflect.TypeOf(&ConfigSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Interact).DeepCopyInto(out.(*Interact))
			return nil
		}, InType: reflect.TypeOf(&Interact{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InteractList).DeepCopyInto(out.(*InteractList))
			return nil
		}, InType: reflect.TypeOf(&InteractList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Jobtype).DeepCopyInto(out.(*Jobtype))
			return nil
		}, InType: reflect.TypeOf(&Jobtype{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*JobtypeList).DeepCopyInto(out.(*JobtypeList))
			return nil
		}, InType: reflect.TypeOf(&JobtypeList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*JobtypeSpec).DeepCopyInto(out.(*JobtypeSpec))
			return nil
		}, InType: reflect.TypeOf(&JobtypeSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*JobtypeStatus).DeepCopyInto(out.(*JobtypeStatus))
			return nil
		}, InType: reflect.TypeOf(&JobtypeStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KryptonRepoSpec).DeepCopyInto(out.(*KryptonRepoSpec))
			return nil
		}, InType: reflect.TypeOf(&KryptonRepoSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NeonRepoSpec).DeepCopyInto(out.(*NeonRepoSpec))
			return nil
		}, InType: reflect.TypeOf(&NeonRepoSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecuritySpec).DeepCopyInto(out.(*SecuritySpec))
			return nil
		}, InType: reflect.TypeOf(&SecuritySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StreamDataSpec).DeepCopyInto(out.(*StreamDataSpec))
			return nil
		}, InType: reflect.TypeOf(&StreamDataSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Streaming).DeepCopyInto(out.(*Streaming))
			return nil
		}, InType: reflect.TypeOf(&Streaming{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StreamingList).DeepCopyInto(out.(*StreamingList))
			return nil
		}, InType: reflect.TypeOf(&StreamingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Training).DeepCopyInto(out.(*Training))
			return nil
		}, InType: reflect.TypeOf(&Training{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TrainingList).DeepCopyInto(out.(*TrainingList))
			return nil
		}, InType: reflect.TypeOf(&TrainingList{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Batch) DeepCopyInto(out *Batch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Batch.
func (x *Batch) DeepCopy() *Batch {
	if x == nil {
		return nil
	}
	out := new(Batch)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Batch) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchList) DeepCopyInto(out *BatchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Batch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new BatchList.
func (x *BatchList) DeepCopy() *BatchList {
	if x == nil {
		return nil
	}
	out := new(BatchList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *BatchList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	out.NeonRepoSpec = in.NeonRepoSpec
	out.SecuritySpec = in.SecuritySpec
	out.StreamDataSpec = in.StreamDataSpec
	out.KryptonRepoSpec = in.KryptonRepoSpec
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (x *ConfigSpec) DeepCopy() *ConfigSpec {
	if x == nil {
		return nil
	}
	out := new(ConfigSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interact) DeepCopyInto(out *Interact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Interact.
func (x *Interact) DeepCopy() *Interact {
	if x == nil {
		return nil
	}
	out := new(Interact)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Interact) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InteractList) DeepCopyInto(out *InteractList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Interact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new InteractList.
func (x *InteractList) DeepCopy() *InteractList {
	if x == nil {
		return nil
	}
	out := new(InteractList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *InteractList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jobtype) DeepCopyInto(out *Jobtype) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Jobtype.
func (x *Jobtype) DeepCopy() *Jobtype {
	if x == nil {
		return nil
	}
	out := new(Jobtype)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Jobtype) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobtypeList) DeepCopyInto(out *JobtypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jobtype, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new JobtypeList.
func (x *JobtypeList) DeepCopy() *JobtypeList {
	if x == nil {
		return nil
	}
	out := new(JobtypeList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *JobtypeList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobtypeSpec) DeepCopyInto(out *JobtypeSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.PodSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	in.IngressSpec.DeepCopyInto(&out.IngressSpec)
	in.ServiceSpec.DeepCopyInto(&out.ServiceSpec)
	out.ScaleSpec = in.ScaleSpec
	out.ConfigSpec = in.ConfigSpec
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new JobtypeSpec.
func (x *JobtypeSpec) DeepCopy() *JobtypeSpec {
	if x == nil {
		return nil
	}
	out := new(JobtypeSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobtypeStatus) DeepCopyInto(out *JobtypeStatus) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new JobtypeStatus.
func (x *JobtypeStatus) DeepCopy() *JobtypeStatus {
	if x == nil {
		return nil
	}
	out := new(JobtypeStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KryptonRepoSpec) DeepCopyInto(out *KryptonRepoSpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new KryptonRepoSpec.
func (x *KryptonRepoSpec) DeepCopy() *KryptonRepoSpec {
	if x == nil {
		return nil
	}
	out := new(KryptonRepoSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeonRepoSpec) DeepCopyInto(out *NeonRepoSpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new NeonRepoSpec.
func (x *NeonRepoSpec) DeepCopy() *NeonRepoSpec {
	if x == nil {
		return nil
	}
	out := new(NeonRepoSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySpec) DeepCopyInto(out *SecuritySpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySpec.
func (x *SecuritySpec) DeepCopy() *SecuritySpec {
	if x == nil {
		return nil
	}
	out := new(SecuritySpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamDataSpec) DeepCopyInto(out *StreamDataSpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StreamDataSpec.
func (x *StreamDataSpec) DeepCopy() *StreamDataSpec {
	if x == nil {
		return nil
	}
	out := new(StreamDataSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Streaming) DeepCopyInto(out *Streaming) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Streaming.
func (x *Streaming) DeepCopy() *Streaming {
	if x == nil {
		return nil
	}
	out := new(Streaming)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Streaming) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamingList) DeepCopyInto(out *StreamingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Streaming, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StreamingList.
func (x *StreamingList) DeepCopy() *StreamingList {
	if x == nil {
		return nil
	}
	out := new(StreamingList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *StreamingList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Training) DeepCopyInto(out *Training) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Training.
func (x *Training) DeepCopy() *Training {
	if x == nil {
		return nil
	}
	out := new(Training)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Training) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingList) DeepCopyInto(out *TrainingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Training, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TrainingList.
func (x *TrainingList) DeepCopy() *TrainingList {
	if x == nil {
		return nil
	}
	out := new(TrainingList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *TrainingList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}