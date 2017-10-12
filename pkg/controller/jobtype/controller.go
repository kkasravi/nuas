
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


package jobtype

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/nervanasystems/nuas/pkg/apis/helium/v1"
	"github.com/nervanasystems/nuas/pkg/controller/sharedinformers"
	listers "github.com/nervanasystems/nuas/pkg/client/listers_generated/helium/v1"
)

// +controller:group=helium,version=v1,kind=Jobtype,resource=jobtypes
type JobtypeControllerImpl struct {
	// informer listens for events about Jobtype
	informer cache.SharedIndexInformer

	// lister indexes properties about Jobtype
	lister listers.JobtypeLister
}

// Init initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
// config - client configuration for talking to the apiserver
// si - informer factory shared across all controllers for listening to events and indexing resource properties
// queue - message queue for handling new events.  unique to this controller.
func (c *JobtypeControllerImpl) Init(
	config *rest.Config,
	si *sharedinformers.SharedInformers,
	queue workqueue.RateLimitingInterface) {

	// Set the informer and lister for subscribing to events and indexing jobtypes labels
	i := si.Factory.Helium().V1().Jobtypes()
	c.informer = i.Informer()
	c.lister = i.Lister()

	// Add an event handler to enqueue a message for jobtypes adds / updates
	c.informer.AddEventHandler(&controller.QueueingEventHandler{queue})
}

// Reconcile handles enqueued messages
func (c *JobtypeControllerImpl) Reconcile(u *v1.Jobtype) error {
	// Implement controller logic here
	log.Printf("Running reconcile Jobtype for %s\n", u.Name)
	return nil
}

func (c *JobtypeControllerImpl) Get(namespace, name string) (*v1.Jobtype, error) {
	return c.lister.Jobtypes(namespace).Get(name)
}
