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

func TestAccIBMCmCatalogDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "offerings_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "owning_account"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "syndication_settings.#"),
				),
			},
		},
	})
}

func TestAccIBMCmCatalogDataSourceAllArgs(t *testing.T) {
	catalogRev := fmt.Sprintf("rev_%d", acctest.RandIntRange(10, 100))
	catalogLabel := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	catalogShortDescription := fmt.Sprintf("short_description_%d", acctest.RandIntRange(10, 100))
	catalogCatalogIconURL := fmt.Sprintf("catalog_icon_url_%d", acctest.RandIntRange(10, 100))
	catalogDisabled := "true"
	catalogResourceGroupID := fmt.Sprintf("resource_group_id_%d", acctest.RandIntRange(10, 100))
	catalogOwningAccount := fmt.Sprintf("owning_account_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogDataSourceConfig(catalogRev, catalogLabel, catalogShortDescription, catalogCatalogIconURL, catalogDisabled, catalogResourceGroupID, catalogOwningAccount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "offerings_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.0.title"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "owning_account"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "syndication_settings.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCmCatalogDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		data "ibm_cm_catalog" "cm_catalog" {
			catalog_identifier = "catalog_identifier"
		}
	`)
}

func testAccCheckIBMCmCatalogDataSourceConfig(catalogRev string, catalogLabel string, catalogShortDescription string, catalogCatalogIconURL string, catalogDisabled string, catalogResourceGroupID string, catalogOwningAccount string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			_rev = "%s"
			label = "%s"
			short_description = "%s"
			catalog_icon_url = "%s"
			tags = "FIXME"
			features = { example: "object" }
			disabled = %s
			resource_group_id = "%s"
			owning_account = "%s"
			catalog_filters = { example: "object" }
			syndication_settings = { example: "object" }
		}

		data "ibm_cm_catalog" "cm_catalog" {
			catalog_identifier = "catalog_identifier"
		}
	`, catalogRev, catalogLabel, catalogShortDescription, catalogCatalogIconURL, catalogDisabled, catalogResourceGroupID, catalogOwningAccount)
}
