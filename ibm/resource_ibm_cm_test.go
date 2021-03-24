/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCm(t *testing.T) {
	clusterId := os.Getenv("CATMGMT_CLUSTERID")
	clusterRegion := os.Getenv("CATMGMT_CLUSTERREGION")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmConfig(clusterId, clusterRegion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog"),
					testAccCheckIBMCmOfferingExists("ibm_cm_offering"),
					testAccCheckIBMCmVersionExists("ibm_cm_version"),
					testAccCheckIBMCmVersionExists("ibm_cm_offering_instance"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_catalog.cm_catalog",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmCatalogConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
		}
	`)
}

func testAccCheckIBMCmConfig(clusterId string, clusterRegion string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_catalog"
			short_description = "testing terraform provider with catalog"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "tf_test_offering
			tags = ["dev_ops", "target_roks", "operator"]
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			zipurl = "https://raw.githubusercontent.com/operator-framework/community-operators/master/community-operators/cockroachdb/5.0.3/manifests/cockroachdb.clusterserviceversion.yaml"
		}
		resource "ibm_cm_offering_instance" "cm_offering_instance_instance" {
			label = "tf_test_offering_label"
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			kind_format = "operator"
			version = ibm_cm_version.cm_version.version.kind_id
			cluster_id = %s
			cluster_region = %s
			cluster_namespaces = ["tf-cm-test"]
			cluster_all_namespaces = false
		}
		  
		data "ibm_cm_catalog" "cm_catalog_data" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}
		  
		data "ibm_cm_offering" "cm_offering_data" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
		}
		  
		data "ibm_cm_version" "cm_version_data" {
			version_loc_id = ibm_cm_version.cm_version.id
		}
		  
		data "ibm_cm_offering_instance" "cm_offering_instance_data" {
			instance_identifier = ibm_cm_offering_instance.cm_offering_instance_instance.id
		}		    
		`, clusterId, clusterRegion)
}

func testAccCheckIBMCmCatalogExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		_, _, err = catalogManagementClient.GetCatalog(getCatalogOptions)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMCmCatalogDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_catalog" {
			continue
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetCatalog(getCatalogOptions)

		if err == nil {
			return fmt.Errorf("cm_catalog still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_catalog (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMCmOfferingExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		getOfferingOptions.SetCatalogIdentifier(rs.Primary.Attributes["catalog_id"])
		getOfferingOptions.SetOfferingID(rs.Primary.ID)

		_, _, err = catalogManagementClient.GetOffering(getOfferingOptions)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMCmVersionExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

		getVersionOptions.SetVersionLocID(rs.Primary.ID)

		_, _, err = catalogManagementClient.GetVersion(getVersionOptions)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMCmOfferingInstanceExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

		getOfferingInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		_, _, err = catalogManagementClient.GetOfferingInstance(getOfferingInstanceOptions)
		if err != nil {
			return err
		}

		return nil
	}
}
