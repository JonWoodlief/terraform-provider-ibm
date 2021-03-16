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

func TestAccIBMCmCatalogBasic(t *testing.T) {
	var conf catalogmanagementv1.Catalog

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfigBasic(),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIBMCmCatalogAllArgs(t *testing.T) {
	var conf catalogmanagementv1.Catalog
	rev := fmt.Sprintf("_rev_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("short_description_%d", acctest.RandIntRange(10, 100))
	catalogIconURL := fmt.Sprintf("catalog_icon_url_%d", acctest.RandIntRange(10, 100))
	disabled := "true"
	resourceGroupID := fmt.Sprintf("resource_group_id_%d", acctest.RandIntRange(10, 100))
	owningAccount := fmt.Sprintf("owning_account_%d", acctest.RandIntRange(10, 100))
	revUpdate := fmt.Sprintf("_rev_%d", acctest.RandIntRange(10, 100))
	labelUpdate := fmt.Sprintf("label_%d", acctest.RandIntRange(10, 100))
	shortDescriptionUpdate := fmt.Sprintf("short_description_%d", acctest.RandIntRange(10, 100))
	catalogIconURLUpdate := fmt.Sprintf("catalog_icon_url_%d", acctest.RandIntRange(10, 100))
	disabledUpdate := "false"
	resourceGroupIDUpdate := fmt.Sprintf("resource_group_id_%d", acctest.RandIntRange(10, 100))
	owningAccountUpdate := fmt.Sprintf("owning_account_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(rev, label, shortDescription, catalogIconURL, disabled, resourceGroupID, owningAccount),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog", conf),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "_rev", rev),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "catalog_icon_url", catalogIconURL),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "owning_account", owningAccount),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(revUpdate, labelUpdate, shortDescriptionUpdate, catalogIconURLUpdate, disabledUpdate, resourceGroupIDUpdate, owningAccountUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "_rev", revUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "label", labelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "short_description", shortDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "catalog_icon_url", catalogIconURLUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "resource_group_id", resourceGroupIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "owning_account", owningAccountUpdate),
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

func testAccCheckIBMCmCatalogConfig(rev string, label string, shortDescription string, catalogIconURL string, disabled string, resourceGroupID string, owningAccount string) string {
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
	`, rev, label, shortDescription, catalogIconURL, disabled, resourceGroupID, owningAccount)
}

func testAccCheckIBMCmCatalogExists(n string, obj catalogmanagementv1.Catalog) resource.TestCheckFunc {

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

		catalog, _, err := catalogManagementClient.GetCatalog(getCatalogOptions)
		if err != nil {
			return err
		}

		obj = *catalog
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
