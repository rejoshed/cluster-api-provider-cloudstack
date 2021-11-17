/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloud

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	_ "github.com/golang/mock/gomock"
	_ "github.com/onsi/gomega"
)

func newCSClient() *cloudstack.CloudStackClient {
	apiUrl, apiKey, secretKey := readAPIConfig()
	return cloudstack.NewAsyncClient(apiUrl, apiKey, secretKey, false)
}

// // From start to finish this test takes a while as it actually instantiates and destroys a VM instance.
// func TestCreateInstance(t *testing.T) {
// 	t.Run("Create Instance", func(t *testing.T) {
// 		cs := newCSClient()

// 		// This cluster should exist.
// 		// None of the tests should ever delete this cluster.
// 		c := CSCluster: &infrav1.CloudStackCluster{
// 				Spec: infrav1.CloudStackClusterSpec{
// 					Zone:    "zone1",
// 					Network: "guestNet1",
// 				},
// 			},
// 		}
// 		err := c.Fetch(cs)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		// Create instance using cluster fetched above.
// 		I := Instance{
// 			Cluster: &infrav1.CloudStackCluster{Status: c.CSCluster.Status},
// 			Machine: &clusterv1.Machine{},
// 			CSMachine: &infrav1.CloudStackMachine{
// 				Spec: infrav1.CloudStackMachineSpec{
// 					Offering: "KubeServiceOffering",
// 					Template: "kube-v1.20.10/ubuntu-1804",
// 					SSHKey:   "CAPCKeyPair6",
// 				},
// 				Status: infrav1.CloudStackMachineStatus{}},
// 		}
// 		err = I.Create(cs)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		err = I.Delete(cs)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	})
// }