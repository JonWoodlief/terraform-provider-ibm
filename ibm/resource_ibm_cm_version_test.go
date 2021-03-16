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

func TestAccIBMCmVersionBasic(t *testing.T) {
	var conf catalogmanagementv1.Version
	catalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	offeringID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	catalogIdentifierUpdate := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	offeringIDUpdate := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfigBasic(catalogIdentifier, offeringID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "catalog_identifier", catalogIdentifier),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "offering_id", offeringID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfigBasic(catalogIdentifierUpdate, offeringIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "catalog_identifier", catalogIdentifierUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "offering_id", offeringIDUpdate),
				),
			},
		},
	})
}

func TestAccIBMCmVersionAllArgs(t *testing.T) {
	var conf catalogmanagementv1.Version
	catalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	offeringID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	zipurl := fmt.Sprintf("zipurl_%d", acctest.RandIntRange(10, 100))
	targetVersion := fmt.Sprintf("targetVersion_%d", acctest.RandIntRange(10, 100))
	includeConfig := "false"
	repoType := fmt.Sprintf("repoType_%d", acctest.RandIntRange(10, 100))
	catalogIdentifierUpdate := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	offeringIDUpdate := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	zipurlUpdate := fmt.Sprintf("zipurl_%d", acctest.RandIntRange(10, 100))
	targetVersionUpdate := fmt.Sprintf("targetVersion_%d", acctest.RandIntRange(10, 100))
	includeConfigUpdate := "true"
	repoTypeUpdate := fmt.Sprintf("repoType_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfig(catalogIdentifier, offeringID, zipurl, targetVersion, includeConfig, repoType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version", conf),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "catalog_identifier", catalogIdentifier),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "offering_id", offeringID),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "zipurl", zipurl),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "targetVersion", targetVersion),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "includeConfig", includeConfig),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "repoType", repoType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfig(catalogIdentifierUpdate, offeringIDUpdate, zipurlUpdate, targetVersionUpdate, includeConfigUpdate, repoTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "catalog_identifier", catalogIdentifierUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "offering_id", offeringIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "zipurl", zipurlUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "targetVersion", targetVersionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "includeConfig", includeConfigUpdate),
					resource.TestCheckResourceAttr("ibm_cm_version.cm_version", "repoType", repoTypeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_version.cm_version",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmVersionConfigBasic(catalogIdentifier string, offeringID string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = "%s"
			offering_id = "%s"
		}
	`, catalogIdentifier, offeringID)
}

func testAccCheckIBMCmVersionConfig(catalogIdentifier string, offeringID string, zipurl string, targetVersion string, includeConfig string, repoType string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = "%s"
			offering_id = "%s"
			tags = "FIXME"
			target_kinds = "FIXME"
			content = "FIXME"
			zipurl = "%s"
			targetVersion = "%s"
			includeConfig = %s
			repoType = "%s"
		}
	`, catalogIdentifier, offeringID, zipurl, targetVersion, includeConfig, repoType)
}

func testAccCheckIBMCmVersionExists(n string, obj catalogmanagementv1.Version) resource.TestCheckFunc {

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

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getVersionOptions.SetVersionLocID(parts[2])

		offering, _, err := catalogManagementClient.GetVersion(getVersionOptions)
		if err != nil {
			return err
		}

		version := offering.Kinds[0].Versions[0]

		obj = version
		return nil
	}
}

func testAccCheckIBMCmVersionDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_version" {
			continue
		}

		getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getVersionOptions.SetVersionLocID(parts[2])

		// Try to find the key
		_, response, err := catalogManagementClient.GetVersion(getVersionOptions)

		if err == nil {
			return fmt.Errorf("cm_version still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_version (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
