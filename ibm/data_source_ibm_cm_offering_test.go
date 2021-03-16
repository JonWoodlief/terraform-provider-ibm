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

func TestAccIBMCmOfferingDataSourceBasic(t *testing.T) {
	offeringCatalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingDataSourceConfigBasic(offeringCatalogIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_docs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_support_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rating.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "permit_request_ibm_public_publish"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "ibm_publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "public_publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "public_original_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "publish_public_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "portal_approval_record"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "portal_ui_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "disclaimer"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "hidden"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "provider"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "repo_info.#"),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingDataSourceAllArgs(t *testing.T) {
	offeringCatalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	offeringZipurl := fmt.Sprintf("zipurl_%d", acctest.RandIntRange(10, 100))
	offeringOfferingID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	offeringTargetVersion := fmt.Sprintf("target_version_%d", acctest.RandIntRange(10, 100))
	offeringIncludeConfig := "true"
	offeringRepoType := fmt.Sprintf("repo_type_%d", acctest.RandIntRange(10, 100))
	offeringXAuthToken := fmt.Sprintf("x_auth_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingDataSourceConfig(offeringCatalogIdentifier, offeringZipurl, offeringOfferingID, offeringTargetVersion, offeringIncludeConfig, offeringRepoType, offeringXAuthToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_docs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_support_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rating.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.0.title"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.format_kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.target_kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.install_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "permit_request_ibm_public_publish"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "ibm_publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "public_publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "public_original_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "publish_public_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "portal_approval_record"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "portal_ui_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "disclaimer"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "hidden"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "provider"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "repo_info.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingDataSourceConfigBasic(offeringCatalogIdentifier string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_offering" "cm_offering" {
			catalog_identifier = "%s"
		}

		data "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_offering.cm_offering.catalog_identifier
			offering_id = ibm_cm_offering.cm_offering.offering_id
		}
	`, offeringCatalogIdentifier)
}

func testAccCheckIBMCmOfferingDataSourceConfig(offeringCatalogIdentifier string, offeringZipurl string, offeringOfferingID string, offeringTargetVersion string, offeringIncludeConfig string, offeringRepoType string, offeringXAuthToken string) string {
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

		data "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_offering.cm_offering.catalog_identifier
			offering_id = ibm_cm_offering.cm_offering.offering_id
		}
	`, offeringCatalogIdentifier, offeringZipurl, offeringOfferingID, offeringTargetVersion, offeringIncludeConfig, offeringRepoType, offeringXAuthToken)
}
