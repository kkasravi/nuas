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
	"github.com/nervanasystems/nuas/pkg/client/clientset_generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type HeliumV1Interface interface {
	RESTClient() rest.Interface
	BatchsGetter
	InteractsGetter
	JobtypesGetter
	StreamingsGetter
	TrainingsGetter
}

// HeliumV1Client is used to interact with features provided by the helium.aipg.intel.com group.
type HeliumV1Client struct {
	restClient rest.Interface
}

func (c *HeliumV1Client) Batchs(namespace string) BatchInterface {
	return newBatchs(c, namespace)
}

func (c *HeliumV1Client) Interacts(namespace string) InteractInterface {
	return newInteracts(c, namespace)
}

func (c *HeliumV1Client) Jobtypes(namespace string) JobtypeInterface {
	return newJobtypes(c, namespace)
}

func (c *HeliumV1Client) Streamings(namespace string) StreamingInterface {
	return newStreamings(c, namespace)
}

func (c *HeliumV1Client) Trainings(namespace string) TrainingInterface {
	return newTrainings(c, namespace)
}

// NewForConfig creates a new HeliumV1Client for the given config.
func NewForConfig(c *rest.Config) (*HeliumV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &HeliumV1Client{client}, nil
}

// NewForConfigOrDie creates a new HeliumV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HeliumV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new HeliumV1Client for the given RESTClient.
func New(c rest.Interface) *HeliumV1Client {
	return &HeliumV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HeliumV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
