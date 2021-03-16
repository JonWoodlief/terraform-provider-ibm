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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmOfferingInstanceDataSourceBasic(t *testing.T) {
	offeringInstanceXAuthRefreshToken := fmt.Sprintf("x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceDataSourceConfigBasic(offeringInstanceXAuthRefreshToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "instance_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "kind_format"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_region"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_namespaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_all_namespaces"),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingInstanceDataSourceAllArgs(t *testing.T) {
	offeringInstanceXAuthRefreshToken := fmt.Sprintf("x_auth_refresh_token_%d", acctest.RandIntRange(10, 100))
	offeringInstanceURL := fmt.Sprintf("url_%d", acctest.RandIntRange(10, 100))
	offeringInstanceCRN := fmt.Sprintf("crn_%d", acctest.RandIntRange(10, 100))
	offeringInstanceLabel := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	offeringInstanceCatalogID := fmt.Sprintf("catalog_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceOfferingID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceKindFormat := fmt.Sprintf("kind_format_%d", acctest.RandIntRange(10, 100))
	offeringInstanceVersion := fmt.Sprintf("version_%d", acctest.RandIntRange(10, 100))
	offeringInstanceClusterID := fmt.Sprintf("cluster_id_%d", acctest.RandIntRange(10, 100))
	offeringInstanceClusterRegion := fmt.Sprintf("cluster_region_%d", acctest.RandIntRange(10, 100))
	offeringInstanceClusterAllNamespaces := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceDataSourceConfig(offeringInstanceXAuthRefreshToken, offeringInstanceURL, offeringInstanceCRN, offeringInstanceLabel, offeringInstanceCatalogID, offeringInstanceOfferingID, offeringInstanceKindFormat, offeringInstanceVersion, offeringInstanceClusterID, offeringInstanceClusterRegion, offeringInstanceClusterAllNamespaces),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "instance_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "kind_format"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_region"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_namespaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance", "cluster_all_namespaces"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingInstanceDataSourceConfigBasic(offeringInstanceXAuthRefreshToken string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			X-Auth-Refresh-Token = "%s"
		}

		data "ibm_cm_offering_instance" "cm_offering_instance" {
			instance_identifier = "instance_identifier"
		}
	`, offeringInstanceXAuthRefreshToken)
}

func testAccCheckIBMCmOfferingInstanceDataSourceConfig(offeringInstanceXAuthRefreshToken string, offeringInstanceURL string, offeringInstanceCRN string, offeringInstanceLabel string, offeringInstanceCatalogID string, offeringInstanceOfferingID string, offeringInstanceKindFormat string, offeringInstanceVersion string, offeringInstanceClusterID string, offeringInstanceClusterRegion string, offeringInstanceClusterAllNamespaces string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			X-Auth-Refresh-Token = "%s"
			url = "%s"
			crn = "%s"
			label = "%s"
			catalog_id = "%s"
			offering_id = "%s"
			kind_format = "%s"
			version = "%s"
			cluster_id = "%s"
			cluster_region = "%s"
			cluster_namespaces = "FIXME"
			cluster_all_namespaces = %s
		}

		data "ibm_cm_offering_instance" "cm_offering_instance" {
			instance_identifier = "instance_identifier"
		}
	`, offeringInstanceXAuthRefreshToken, offeringInstanceURL, offeringInstanceCRN, offeringInstanceLabel, offeringInstanceCatalogID, offeringInstanceOfferingID, offeringInstanceKindFormat, offeringInstanceVersion, offeringInstanceClusterID, offeringInstanceClusterRegion, offeringInstanceClusterAllNamespaces)
}
