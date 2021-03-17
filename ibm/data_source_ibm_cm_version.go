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
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func dataSourceIBMCmVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCmVersionRead,

		Schema: map[string]*schema.Schema{
			"version_loc_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "A dotted value of `catalogID`.`versionID`.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID.",
			},
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloudant revision.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version's CRN.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of content type.",
			},
			"sha": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "hash of the content.",
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time this version was created.",
			},
			"updated": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time this version was last updated.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Offering ID.",
			},
			"catalog_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Catalog ID.",
			},
			"kind_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kind ID.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of tags associated with this catalog.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"repo_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Content's repo URL.",
			},
			"source_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Content's source URL (e.g git repo).",
			},
			"tgz_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File used to on-board this version.",
			},
			"configuration": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of user solicited overrides.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configuration key.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value type (string, boolean, int).",
						},
						"default_value": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.",
						},
						"value_constraint": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Constraint associated with value, e.g., for string type - regx:[a-z].",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key description.",
						},
						"required": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is key required to install.",
						},
						"options": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of options of type.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"hidden": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Hide values.",
						},
					},
				},
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Open ended metadata information.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"validation": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Validation response.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"validated": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of last successful validation.",
						},
						"requested": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of last validation was requested.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current validation state - <empty>, in_progress, valid, invalid, expired.",
						},
						"last_operation": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last operation (e.g. submit_deployment, generate_installer, install_offering.",
						},
						"target": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"required_resources": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Resource requirments for installation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of requirement.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.",
						},
					},
				},
			},
			"single_instance": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Denotes if single instance can be deployed to a given cluster.",
			},
			"install": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Script information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instructions": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.",
						},
						"script": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional script that needs to be run post any pre-condition script.",
						},
						"script_permission": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional iam permissions that are required on the target cluster to run this script.",
						},
						"delete_script": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional script that if run will remove the installed version.",
						},
						"scope": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional value indicating if this script is scoped to a namespace or the entire cluster.",
						},
					},
				},
			},
			"pre_install": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Optional pre-install instructions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instructions": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.",
						},
						"script": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional script that needs to be run post any pre-condition script.",
						},
						"script_permission": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional iam permissions that are required on the target cluster to run this script.",
						},
						"delete_script": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional script that if run will remove the installed version.",
						},
						"scope": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional value indicating if this script is scoped to a namespace or the entire cluster.",
						},
					},
				},
			},
			"entitlement": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Entitlement license info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Provider name.",
						},
						"provider_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Provider ID.",
						},
						"product_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product ID.",
						},
						"part_numbers": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"image_repo_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Image repository name.",
						},
					},
				},
			},
			"licenses": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of licenses the product was built with.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "License ID.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "license name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "type of license e.g., Apache xxx.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL for the license text.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "License description.",
						},
					},
				},
			},
			"image_manifest_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set, denotes a url to a YAML file with list of container images used by this version.",
			},
			"deprecated": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "read only field, indicating if this version is deprecated.",
			},
			"package_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the package used to create this version.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Offering state.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "one of: new, validated, account-published, ibm-published, public-published.",
						},
						"current_entered": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of current request.",
						},
						"pending": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "one of: new, validated, account-published, ibm-published, public-published.",
						},
						"pending_requested": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time of pending request.",
						},
						"previous": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "one of: new, validated, account-published, ibm-published, public-published.",
						},
					},
				},
			},
			"version_locator": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A dotted value of `catalogID`.`versionID`.",
			},
			"console_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Console URL.",
			},
			"long_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Long description for version.",
			},
			"whitelisted_accounts": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whitelisted accounts for version.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIBMCmVersionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

	getVersionOptions.SetVersionLocID(d.Get("version_loc_id").(string))

	offering, response, err := catalogManagementClient.GetVersionWithContext(context, getVersionOptions)
	version := offering.Kinds[0].Versions[0]

	if err != nil {
		log.Printf("[DEBUG] GetVersionWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *version.CatalogID, *offering.ID, *version.ID))

	if err = d.Set("id", version.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}
	if err = d.Set("rev", version.Rev); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rev: %s", err))
	}
	if err = d.Set("crn", version.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("version", version.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}
	if err = d.Set("sha", version.Sha); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sha: %s", err))
	}
	if err = d.Set("created", version.Created); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if err = d.Set("updated", version.Updated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}
	if err = d.Set("catalog_id", version.CatalogID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_id: %s", err))
	}
	if err = d.Set("tags", version.Tags); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
	}
	if err = d.Set("repo_url", version.RepoURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting repo_url: %s", err))
	}
	if err = d.Set("source_url", version.SourceURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_url: %s", err))
	}
	if err = d.Set("tgz_url", version.TgzURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tgz_url: %s", err))
	}

	if version.Configuration != nil {
		err = d.Set("configuration", dataSourceVersionFlattenConfiguration(version.Configuration))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting configuration %s", err))
		}
	}

	if version.Metadata != nil {
		convertedMap := make(map[string]interface{}, len(version.Metadata))
		for k, v := range version.Metadata {
			convertedMap[k] = v
		}

		if err = d.Set("metadata", Flatten(convertedMap)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata: %s", err))
		}
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata %s", err))
		}
	}

	if version.Validation != nil {
		err = d.Set("validation", dataSourceVersionFlattenValidation(*version.Validation))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting validation %s", err))
		}
	}

	if version.RequiredResources != nil {
		err = d.Set("required_resources", dataSourceVersionFlattenRequiredResources(version.RequiredResources))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting required_resources %s", err))
		}
	}
	if err = d.Set("single_instance", version.SingleInstance); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting single_instance: %s", err))
	}

	if version.Install != nil {
		err = d.Set("install", dataSourceVersionFlattenInstall(*version.Install))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting install %s", err))
		}
	}

	if version.PreInstall != nil {
		err = d.Set("pre_install", dataSourceVersionFlattenPreInstall(version.PreInstall))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting pre_install %s", err))
		}
	}

	if version.Entitlement != nil {
		err = d.Set("entitlement", dataSourceVersionFlattenEntitlement(*version.Entitlement))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting entitlement %s", err))
		}
	}

	if version.Licenses != nil {
		err = d.Set("licenses", dataSourceVersionFlattenLicenses(version.Licenses))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting licenses %s", err))
		}
	}
	if err = d.Set("image_manifest_url", version.ImageManifestURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting image_manifest_url: %s", err))
	}
	if err = d.Set("deprecated", version.Deprecated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting deprecated: %s", err))
	}
	if err = d.Set("package_version", version.PackageVersion); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting package_version: %s", err))
	}

	if version.State != nil {
		err = d.Set("state", dataSourceVersionFlattenState(*version.State))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting state %s", err))
		}
	}
	if err = d.Set("version_locator", version.VersionLocator); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version_locator: %s", err))
	}
	if err = d.Set("console_url", version.ConsoleURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting console_url: %s", err))
	}
	if err = d.Set("long_description", version.LongDescription); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting long_description: %s", err))
	}
	if err = d.Set("whitelisted_accounts", version.WhitelistedAccounts); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting whitelisted_accounts: %s", err))
	}

	return nil
}

func dataSourceVersionFlattenConfiguration(result []catalogmanagementv1.Configuration) (configuration []map[string]interface{}) {
	for _, configurationItem := range result {
		configuration = append(configuration, dataSourceVersionConfigurationToMap(configurationItem))
	}

	return configuration
}

func dataSourceVersionConfigurationToMap(configurationItem catalogmanagementv1.Configuration) (configurationMap map[string]interface{}) {
	configurationMap = map[string]interface{}{}

	if configurationItem.Key != nil {
		configurationMap["key"] = configurationItem.Key
	}
	if configurationItem.Type != nil {
		configurationMap["type"] = configurationItem.Type
	}
	if configurationItem.DefaultValue != nil {
		configurationMap["default_value"] = configurationItem.DefaultValue
	}
	if configurationItem.ValueConstraint != nil {
		configurationMap["value_constraint"] = configurationItem.ValueConstraint
	}
	if configurationItem.Description != nil {
		configurationMap["description"] = configurationItem.Description
	}
	if configurationItem.Required != nil {
		configurationMap["required"] = configurationItem.Required
	}
	if configurationItem.Options != nil {
		configurationMap["options"] = configurationItem.Options
	}
	if configurationItem.Hidden != nil {
		configurationMap["hidden"] = configurationItem.Hidden
	}

	return configurationMap
}

func dataSourceVersionFlattenValidation(result catalogmanagementv1.Validation) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVersionValidationToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVersionValidationToMap(validationItem catalogmanagementv1.Validation) (validationMap map[string]interface{}) {
	validationMap = map[string]interface{}{}

	if validationItem.Validated != nil {
		validationMap["validated"] = validationItem.Validated.String()
	}
	if validationItem.Requested != nil {
		validationMap["requested"] = validationItem.Requested.String()
	}
	if validationItem.State != nil {
		validationMap["state"] = validationItem.State
	}
	if validationItem.LastOperation != nil {
		validationMap["last_operation"] = validationItem.LastOperation
	}
	if validationItem.Target != nil {
		convertedMap := make(map[string]interface{}, len(validationItem.Target))
		for k, v := range validationItem.Target {
			convertedMap[k] = v
		}
		validationMap["target"] = Flatten(convertedMap)
	}

	return validationMap
}

func dataSourceVersionFlattenRequiredResources(result []catalogmanagementv1.Resource) (requiredResources []map[string]interface{}) {
	for _, requiredResourcesItem := range result {
		requiredResources = append(requiredResources, dataSourceVersionRequiredResourcesToMap(requiredResourcesItem))
	}

	return requiredResources
}

func dataSourceVersionRequiredResourcesToMap(requiredResourcesItem catalogmanagementv1.Resource) (requiredResourcesMap map[string]interface{}) {
	requiredResourcesMap = map[string]interface{}{}

	if requiredResourcesItem.Type != nil {
		requiredResourcesMap["type"] = requiredResourcesItem.Type
	}
	if requiredResourcesItem.Value != nil {
		requiredResourcesMap["value"] = requiredResourcesItem.Value
	}

	return requiredResourcesMap
}

func dataSourceVersionFlattenInstall(result catalogmanagementv1.Script) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVersionInstallToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVersionInstallToMap(installItem catalogmanagementv1.Script) (installMap map[string]interface{}) {
	installMap = map[string]interface{}{}

	if installItem.Instructions != nil {
		installMap["instructions"] = installItem.Instructions
	}
	if installItem.Script != nil {
		installMap["script"] = installItem.Script
	}
	if installItem.ScriptPermission != nil {
		installMap["script_permission"] = installItem.ScriptPermission
	}
	if installItem.DeleteScript != nil {
		installMap["delete_script"] = installItem.DeleteScript
	}
	if installItem.Scope != nil {
		installMap["scope"] = installItem.Scope
	}

	return installMap
}

func dataSourceVersionFlattenPreInstall(result []catalogmanagementv1.Script) (preInstall []map[string]interface{}) {
	for _, preInstallItem := range result {
		preInstall = append(preInstall, dataSourceVersionPreInstallToMap(preInstallItem))
	}

	return preInstall
}

func dataSourceVersionPreInstallToMap(preInstallItem catalogmanagementv1.Script) (preInstallMap map[string]interface{}) {
	preInstallMap = map[string]interface{}{}

	if preInstallItem.Instructions != nil {
		preInstallMap["instructions"] = preInstallItem.Instructions
	}
	if preInstallItem.Script != nil {
		preInstallMap["script"] = preInstallItem.Script
	}
	if preInstallItem.ScriptPermission != nil {
		preInstallMap["script_permission"] = preInstallItem.ScriptPermission
	}
	if preInstallItem.DeleteScript != nil {
		preInstallMap["delete_script"] = preInstallItem.DeleteScript
	}
	if preInstallItem.Scope != nil {
		preInstallMap["scope"] = preInstallItem.Scope
	}

	return preInstallMap
}

func dataSourceVersionFlattenEntitlement(result catalogmanagementv1.VersionEntitlement) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVersionEntitlementToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVersionEntitlementToMap(entitlementItem catalogmanagementv1.VersionEntitlement) (entitlementMap map[string]interface{}) {
	entitlementMap = map[string]interface{}{}

	if entitlementItem.ProviderName != nil {
		entitlementMap["provider_name"] = entitlementItem.ProviderName
	}
	if entitlementItem.ProviderID != nil {
		entitlementMap["provider_id"] = entitlementItem.ProviderID
	}
	if entitlementItem.ProductID != nil {
		entitlementMap["product_id"] = entitlementItem.ProductID
	}
	if entitlementItem.PartNumbers != nil {
		entitlementMap["part_numbers"] = entitlementItem.PartNumbers
	}
	if entitlementItem.ImageRepoName != nil {
		entitlementMap["image_repo_name"] = entitlementItem.ImageRepoName
	}

	return entitlementMap
}

func dataSourceVersionFlattenLicenses(result []catalogmanagementv1.License) (licenses []map[string]interface{}) {
	for _, licensesItem := range result {
		licenses = append(licenses, dataSourceVersionLicensesToMap(licensesItem))
	}

	return licenses
}

func dataSourceVersionLicensesToMap(licensesItem catalogmanagementv1.License) (licensesMap map[string]interface{}) {
	licensesMap = map[string]interface{}{}

	if licensesItem.ID != nil {
		licensesMap["id"] = licensesItem.ID
	}
	if licensesItem.Name != nil {
		licensesMap["name"] = licensesItem.Name
	}
	if licensesItem.Type != nil {
		licensesMap["type"] = licensesItem.Type
	}
	if licensesItem.URL != nil {
		licensesMap["url"] = licensesItem.URL
	}
	if licensesItem.Description != nil {
		licensesMap["description"] = licensesItem.Description
	}

	return licensesMap
}

func dataSourceVersionFlattenState(result catalogmanagementv1.State) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVersionStateToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVersionStateToMap(stateItem catalogmanagementv1.State) (stateMap map[string]interface{}) {
	stateMap = map[string]interface{}{}

	if stateItem.Current != nil {
		stateMap["current"] = stateItem.Current
	}
	if stateItem.CurrentEntered != nil {
		stateMap["current_entered"] = stateItem.CurrentEntered.String()
	}
	if stateItem.Pending != nil {
		stateMap["pending"] = stateItem.Pending
	}
	if stateItem.PendingRequested != nil {
		stateMap["pending_requested"] = stateItem.PendingRequested.String()
	}
	if stateItem.Previous != nil {
		stateMap["previous"] = stateItem.Previous
	}

	return stateMap
}
