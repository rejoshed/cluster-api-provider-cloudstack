/*
Copyright 2022 The Kubernetes Authors.

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

package cloud_test

import (
	"os"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/aws/cluster-api-provider-cloudstack-staging/test/unit/dummies"
	"github.com/aws/cluster-api-provider-cloudstack/pkg/cloud"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	realCloudClient cloud.Client
	client          cloud.Client
	realCSClient    *cloudstack.CloudStackClient
)

var _ = BeforeSuite(func() {
	projDir := os.Getenv("PROJECT_DIR")
	var connectionErr error
	realCloudClient, connectionErr = cloud.NewClient(projDir + "/cloud-config")
	Ω(connectionErr).ShouldNot(HaveOccurred())

	Ω().Should(Succeed())
})

func TestCloud(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cloud Suite")
}

// FetchIntegTestResources runs through basic CloudStack Client setup methods needed to test others.
func FetchIntegTestResources() {
	Ω(client.ResolveZone(dummies.CSZone1)).Should(Succeed())
	Ω(dummies.CSZone1.Spec.ID).ShouldNot(BeEmpty())
	dummies.CSMachine1.Status.ZoneID = dummies.CSZone1.Spec.ID
	dummies.CSMachine1.Spec.DiskOffering.Name = ""
	dummies.CSCluster.Spec.ControlPlaneEndpoint.Host = ""
	Ω(client.GetOrCreateIsolatedNetwork(dummies.CSZone1, dummies.CSISONet1, dummies.CSCluster)).Should(Succeed())
}