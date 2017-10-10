
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


package v1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/nervanasystems/nuas/pkg/apis/helium/v1"
	. "github.com/nervanasystems/nuas/pkg/client/clientset_generated/clientset/typed/helium/v1"
)

var _ = Describe("Jobtype", func() {
	var instance Jobtype
	var expected Jobtype
	var client JobtypeInterface

	BeforeEach(func() {
		instance = Jobtype{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	Describe("when sending a training request", func() {
		It("should return success", func() {
			client = cs.HeliumV1Client.Jobtypes("jobtype-test-training")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			training := &Training{}
			training.Name = instance.Name
			restClient := cs.HeliumV1Client.RESTClient()
			err = restClient.Post().Namespace("jobtype-test-training").
				Name(instance.Name).
				Resource("jobtypes").
				SubResource("training").
				Body(training).Do().Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
