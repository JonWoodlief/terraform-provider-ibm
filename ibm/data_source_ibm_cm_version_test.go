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

func TestAccIBMCmVersionDataSourceBasic(t *testing.T) {
	versionCatalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	versionOfferingID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionDataSourceConfigBasic(versionCatalogIdentifier, versionOfferingID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_loc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "sha"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "kind_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "repo_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "source_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tgz_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "validation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "single_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "install.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "entitlement.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "image_manifest_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "deprecated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "package_version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_locator"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "console_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "whitelisted_accounts.#"),
				),
			},
		},
	})
}

func TestAccIBMCmVersionDataSourceAllArgs(t *testing.T) {
	versionCatalogIdentifier := fmt.Sprintf("catalog_identifier_%d", acctest.RandIntRange(10, 100))
	versionOfferingID := fmt.Sprintf("offering_id_%d", acctest.RandIntRange(10, 100))
	versionZipurl := fmt.Sprintf("zipurl_%d", acctest.RandIntRange(10, 100))
	versionTargetVersion := fmt.Sprintf("target_version_%d", acctest.RandIntRange(10, 100))
	versionIncludeConfig := "true"
	versionRepoType := fmt.Sprintf("repo_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionDataSourceConfig(versionCatalogIdentifier, versionOfferingID, versionZipurl, versionTargetVersion, versionIncludeConfig, versionRepoType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_loc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "sha"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "kind_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "repo_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "source_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tgz_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.key"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.default_value"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.value_constraint"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.required"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.hidden"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "validation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "single_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "install.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.instructions"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.script"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.script_permission"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.delete_script"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.scope"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "entitlement.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "image_manifest_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "deprecated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "package_version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_locator"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "console_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "whitelisted_accounts.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCmVersionDataSourceConfigBasic(versionCatalogIdentifier string, versionOfferingID string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = "%s"
			offering_id = "%s"
		}

		data "ibm_cm_version" "cm_version" {
			version_loc_id = "version_loc_id"
		}
	`, versionCatalogIdentifier, versionOfferingID)
}

func testAccCheckIBMCmVersionDataSourceConfig(versionCatalogIdentifier string, versionOfferingID string, versionZipurl string, versionTargetVersion string, versionIncludeConfig string, versionRepoType string) string {
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

		data "ibm_cm_version" "cm_version" {
			version_loc_id = "version_loc_id"
		}
	`, versionCatalogIdentifier, versionOfferingID, versionZipurl, versionTargetVersion, versionIncludeConfig, versionRepoType)
}
