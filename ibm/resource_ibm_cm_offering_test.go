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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmOfferingBasic(t *testing.T) {
	var conf catalogmanagementv1.Offering
	catalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	catalogIdentifierUpdate := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfigBasic(catalogIdentifier),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingExists("ibm_cm_offering.cm_offering", conf),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "catalog_identifier", catalogIdentifier),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfigBasic(catalogIdentifierUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "catalog_identifier", catalogIdentifierUpdate),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingAllArgs(t *testing.T) {
	var conf catalogmanagementv1.Offering
	catalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	zipurl := fmt.Sprintf("zipurl_%d", acctest.RandIntRange(10, 100))
	offeringID := fmt.Sprintf("offeringID_%d", acctest.RandIntRange(10, 100))
	targetVersion := fmt.Sprintf("targetVersion_%d", acctest.RandIntRange(10, 100))
	includeConfig := "false"
	repoType := fmt.Sprintf("repoType_%d", acctest.RandIntRange(10, 100))
	xAuthToken := fmt.Sprintf("X-Auth-Token_%d", acctest.RandIntRange(10, 100))
	catalogIdentifierUpdate := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	zipurlUpdate := fmt.Sprintf("zipurl_%d", acctest.RandIntRange(10, 100))
	offeringIDUpdate := fmt.Sprintf("offeringID_%d", acctest.RandIntRange(10, 100))
	targetVersionUpdate := fmt.Sprintf("targetVersion_%d", acctest.RandIntRange(10, 100))
	includeConfigUpdate := "true"
	repoTypeUpdate := fmt.Sprintf("repoType_%d", acctest.RandIntRange(10, 100))
	xAuthTokenUpdate := fmt.Sprintf("X-Auth-Token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfig(catalogIdentifier, zipurl, offeringID, targetVersion, includeConfig, repoType, xAuthToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingExists("ibm_cm_offering.cm_offering", conf),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "catalog_identifier", catalogIdentifier),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "zipurl", zipurl),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offeringID", offeringID),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "targetVersion", targetVersion),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "includeConfig", includeConfig),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "repoType", repoType),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "X-Auth-Token", xAuthToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfig(catalogIdentifierUpdate, zipurlUpdate, offeringIDUpdate, targetVersionUpdate, includeConfigUpdate, repoTypeUpdate, xAuthTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "catalog_identifier", catalogIdentifierUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "zipurl", zipurlUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offeringID", offeringIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "targetVersion", targetVersionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "includeConfig", includeConfigUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "repoType", repoTypeUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "X-Auth-Token", xAuthTokenUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_offering.cm_offering",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmOfferingConfigBasic(catalogIdentifier string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_offering" "cm_offering" {
			catalog_identifier = "%s"
		}
	`, catalogIdentifier)
}

func testAccCheckIBMCmOfferingConfig(catalogIdentifier string, zipurl string, offeringID string, targetVersion string, includeConfig string, repoType string, xAuthToken string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_offering" "cm_offering" {
			catalog_identifier = "%s"
			tags = "FIXME"
			target_kinds = "FIXME"
			content = "FIXME"
			zipurl = "%s"
			offeringID = "%s"
			targetVersion = "%s"
			includeConfig = %s
			repoType = "%s"
			X-Auth-Token = "%s"
		}
	`, catalogIdentifier, zipurl, offeringID, targetVersion, includeConfig, repoType, xAuthToken)
}

func testAccCheckIBMCmOfferingExists(n string, obj catalogmanagementv1.Offering) resource.TestCheckFunc {

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

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getOfferingOptions.SetCatalogIdentifier(parts[0])
		getOfferingOptions.SetOfferingID(parts[1])

		offering, _, err := catalogManagementClient.GetOffering(getOfferingOptions)
		if err != nil {
			return err
		}

		obj = *offering
		return nil
	}
}

func testAccCheckIBMCmOfferingDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_offering" {
			continue
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getOfferingOptions.SetCatalogIdentifier(parts[0])
		getOfferingOptions.SetOfferingID(parts[1])

		// Try to find the key
		_, response, err := catalogManagementClient.GetOffering(getOfferingOptions)

		if err == nil {
			return fmt.Errorf("cm_offering still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_offering (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
