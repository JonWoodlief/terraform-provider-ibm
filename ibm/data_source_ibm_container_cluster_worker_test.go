// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerWorkerDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterWorkerDataSourceConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_worker.testacc_ds_worker", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterWorkerDataSourceConfig(clusterName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  wait_till       = "MasterNodeReady"
}

data "ibm_container_cluster" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
}

data "ibm_container_cluster_worker" "testacc_ds_worker" {
  worker_id = data.ibm_container_cluster.testacc_ds_cluster.workers[0]
}
`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}
