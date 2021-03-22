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
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func resourceIBMCmVersion() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCmVersionCreate,
		Read:     resourceIBMCmVersionRead,
		Delete:   resourceIBMCmVersionDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"catalog_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Catalog identifier.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Offering identification.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Tags array.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"target_kinds": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"content": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "byte array representing the content to be imported.  Only supported for OVA images at this time.",
			},
			"zipurl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "URL path to zip location.  If not specified, must provide content in the body of this call.",
			},
			"target_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The semver value for this new version, if not found in the zip url package content.",
			},
			"include_config": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Add all possible configuration values to this version when importing.",
			},
			"repo_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'.",
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
							Optional:    true,
							Description: "Configuration key.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Value type (string, boolean, int).",
						},
						"default_value": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.",
						},
						"value_constraint": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Constraint associated with value, e.g., for string type - regx:[a-z].",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Key description.",
						},
						"required": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Is key required to install.",
						},
						"options": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of options of type.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"hidden": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Hide values.",
						},
					},
				},
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Open ended metadata information.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"validation": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Validation response.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"validated": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time of last successful validation.",
						},
						"requested": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time of last validation was requested.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Current validation state - <empty>, in_progress, valid, invalid, expired.",
						},
						"last_operation": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Last operation (e.g. submit_deployment, generate_installer, install_offering.",
						},
						"target": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.",
							Elem:        &schema.Schema{Type: schema.TypeString},
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
							Optional:    true,
							Description: "Type of requirement.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
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
				Computed:    true,
				Description: "Script information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instructions": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.",
						},
						"script": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional script that needs to be run post any pre-condition script.",
						},
						"script_permission": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional iam permissions that are required on the target cluster to run this script.",
						},
						"delete_script": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional script that if run will remove the installed version.",
						},
						"scope": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
							Optional:    true,
							Description: "Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.",
						},
						"script": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional script that needs to be run post any pre-condition script.",
						},
						"script_permission": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional iam permissions that are required on the target cluster to run this script.",
						},
						"delete_script": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional script that if run will remove the installed version.",
						},
						"scope": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
							Optional:    true,
							Description: "Provider name.",
						},
						"provider_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Provider ID.",
						},
						"product_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Product ID.",
						},
						"part_numbers": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"image_repo_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
							Optional:    true,
							Description: "License ID.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "license name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "type of license e.g., Apache xxx.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "URL for the license text.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
				Computed:    true,
				Description: "Offering state.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "one of: new, validated, account-published, ibm-published, public-published.",
						},
						"current_entered": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time of current request.",
						},
						"pending": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "one of: new, validated, account-published, ibm-published, public-published.",
						},
						"pending_requested": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time of pending request.",
						},
						"previous": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceIBMCmVersionCreate(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	importOfferingVersionOptions := &catalogmanagementv1.ImportOfferingVersionOptions{}

	importOfferingVersionOptions.SetCatalogIdentifier(d.Get("catalog_identifier").(string))
	importOfferingVersionOptions.SetOfferingID(d.Get("offering_id").(string))
	if _, ok := d.GetOk("tags"); ok {
		importOfferingVersionOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("target_kinds"); ok {
		// importOfferingVersionOptions.SetTargetKinds(d.Get("target_kinds").([]string))
		list := expandStringList(d.Get("target_kinds").([]interface{}))
		importOfferingVersionOptions.SetTargetKinds(list)

	}
	if _, ok := d.GetOk("content"); ok {

	}
	if _, ok := d.GetOk("zipurl"); ok {
		importOfferingVersionOptions.SetZipurl(d.Get("zipurl").(string))
	}
	if _, ok := d.GetOk("target_version"); ok {
		importOfferingVersionOptions.SetTargetVersion(d.Get("target_version").(string))
	}
	if _, ok := d.GetOk("include_config"); ok {
		importOfferingVersionOptions.SetIncludeConfig(d.Get("include_config").(bool))
	}
	if _, ok := d.GetOk("repo_type"); ok {
		importOfferingVersionOptions.SetRepoType(d.Get("repo_type").(string))
	}

	offering, response, err := catalogManagementClient.ImportOfferingVersion(importOfferingVersionOptions)

	if err != nil {
		log.Printf("[DEBUG] ImportOfferingVersion failed %s\n%s", err, response)
		return err
	}

	version := offering.Kinds[0].Versions[0]

	d.SetId(fmt.Sprintf("%s/%s/%s", *importOfferingVersionOptions.CatalogIdentifier, *importOfferingVersionOptions.OfferingID, *version.ID))

	return resourceIBMCmVersionRead(d, meta)
}

func resourceIBMCmVersionRead(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	getVersionOptions.SetVersionLocID(parts[2])

	offering, response, err := catalogManagementClient.GetVersion(getVersionOptions)
	version := offering.Kinds[0].Versions[0]

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetVersion failed %s\n%s", err, response)
		return err
	}

	if err = d.Set("crn", version.CRN); err != nil {
		return fmt.Errorf("Error setting crn: %s", err)
	}
	if err = d.Set("version", version.Version); err != nil {
		return fmt.Errorf("Error setting version: %s", err)
	}
	if err = d.Set("sha", version.Sha); err != nil {
		return fmt.Errorf("Error setting sha: %s", err)
	}
	if err = d.Set("created", version.Created.String()); err != nil {
		return fmt.Errorf("Error setting created: %s", err)
	}
	if err = d.Set("updated", version.Updated.String()); err != nil {
		return fmt.Errorf("Error setting updated: %s", err)
	}
	if err = d.Set("catalog_id", version.CatalogID); err != nil {
		return fmt.Errorf("Error setting catalog_id: %s", err)
	}
	if err = d.Set("kind_id", version.KindID); err != nil {
		return fmt.Errorf("Error setting kind_id: %s", err)
	}
	if err = d.Set("repo_url", version.RepoURL); err != nil {
		return fmt.Errorf("Error setting repo_url: %s", err)
	}
	if err = d.Set("source_url", version.SourceURL); err != nil {
		return fmt.Errorf("Error setting source_url: %s", err)
	}
	if err = d.Set("tgz_url", version.TgzURL); err != nil {
		return fmt.Errorf("Error setting tgz_url: %s", err)
	}
	if version.Configuration != nil {
		configuration := []map[string]interface{}{}
		for _, configurationItem := range version.Configuration {
			configurationItemMap := resourceIBMCmVersionConfigurationToMap(configurationItem)
			configuration = append(configuration, configurationItemMap)
		}
		if err = d.Set("configuration", configuration); err != nil {
			return fmt.Errorf("Error setting configuration: %s", err)
		}
	}
	if version.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- not primitive type, not list
	}
	if version.Validation != nil {
		validationMap := resourceIBMCmVersionValidationToMap(*version.Validation)
		if err = d.Set("validation", []map[string]interface{}{validationMap}); err != nil {
			return fmt.Errorf("Error setting validation: %s", err)
		}
	}
	if version.RequiredResources != nil {
		requiredResources := []map[string]interface{}{}
		for _, requiredResourcesItem := range version.RequiredResources {
			requiredResourcesItemMap := resourceIBMCmVersionResourceToMap(requiredResourcesItem)
			requiredResources = append(requiredResources, requiredResourcesItemMap)
		}
		if err = d.Set("required_resources", requiredResources); err != nil {
			return fmt.Errorf("Error setting required_resources: %s", err)
		}
	}
	if err = d.Set("single_instance", version.SingleInstance); err != nil {
		return fmt.Errorf("Error setting single_instance: %s", err)
	}
	if version.Install != nil {
		installMap := resourceIBMCmVersionScriptToMap(*version.Install)
		if err = d.Set("install", []map[string]interface{}{installMap}); err != nil {
			return fmt.Errorf("Error setting install: %s", err)
		}
	}
	if version.PreInstall != nil {
		preInstall := []map[string]interface{}{}
		for _, preInstallItem := range version.PreInstall {
			preInstallItemMap := resourceIBMCmVersionScriptToMap(preInstallItem)
			preInstall = append(preInstall, preInstallItemMap)
		}
		if err = d.Set("pre_install", preInstall); err != nil {
			return fmt.Errorf("Error setting pre_install: %s", err)
		}
	}
	if version.Entitlement != nil {
		entitlementMap := resourceIBMCmVersionVersionEntitlementToMap(*version.Entitlement)
		if err = d.Set("entitlement", []map[string]interface{}{entitlementMap}); err != nil {
			return fmt.Errorf("Error setting entitlement: %s", err)
		}
	}
	if version.Licenses != nil {
		licenses := []map[string]interface{}{}
		for _, licensesItem := range version.Licenses {
			licensesItemMap := resourceIBMCmVersionLicenseToMap(licensesItem)
			licenses = append(licenses, licensesItemMap)
		}
		if err = d.Set("licenses", licenses); err != nil {
			return fmt.Errorf("Error setting licenses: %s", err)
		}
	}
	if err = d.Set("image_manifest_url", version.ImageManifestURL); err != nil {
		return fmt.Errorf("Error setting image_manifest_url: %s", err)
	}
	if err = d.Set("deprecated", version.Deprecated); err != nil {
		return fmt.Errorf("Error setting deprecated: %s", err)
	}
	if err = d.Set("package_version", version.PackageVersion); err != nil {
		return fmt.Errorf("Error setting package_version: %s", err)
	}
	if version.State != nil {
		stateMap := resourceIBMCmVersionStateToMap(*version.State)
		if err = d.Set("state", []map[string]interface{}{stateMap}); err != nil {
			return fmt.Errorf("Error setting state: %s", err)
		}
	}
	if err = d.Set("version_locator", version.VersionLocator); err != nil {
		return fmt.Errorf("Error setting version_locator: %s", err)
	}
	if err = d.Set("console_url", version.ConsoleURL); err != nil {
		return fmt.Errorf("Error setting console_url: %s", err)
	}
	if err = d.Set("long_description", version.LongDescription); err != nil {
		return fmt.Errorf("Error setting long_description: %s", err)
	}
	if version.WhitelistedAccounts != nil {
		if err = d.Set("whitelisted_accounts", version.WhitelistedAccounts); err != nil {
			return fmt.Errorf("Error setting whitelisted_accounts: %s", err)
		}
	}

	return nil
}

func resourceIBMCmVersionConfigurationToMap(configuration catalogmanagementv1.Configuration) map[string]interface{} {
	configurationMap := map[string]interface{}{}

	configurationMap["key"] = configuration.Key
	configurationMap["type"] = configuration.Type
	configurationMap["default_value"] = configuration.DefaultValue
	configurationMap["value_constraint"] = configuration.ValueConstraint
	configurationMap["description"] = configuration.Description
	configurationMap["required"] = configuration.Required
	/* 	if configuration.Options != nil {
		options := []map[string]interface{}{}
		for _, optionsItem := range configuration.Options {
			options = append(options, optionsItem)
		}
		configurationMap["options"] = options
	} */
	configurationMap["hidden"] = configuration.Hidden

	return configurationMap
}

func resourceIBMCmVersionValidationToMap(validation catalogmanagementv1.Validation) map[string]interface{} {
	validationMap := map[string]interface{}{}

	validationMap["validated"] = validation.Validated.String()
	validationMap["requested"] = validation.Requested.String()
	validationMap["state"] = validation.State
	validationMap["last_operation"] = validation.LastOperation
	if validation.Target != nil {
		// TODO: handle Target of type TypeMap -- container, not list
	}

	return validationMap
}

func resourceIBMCmVersionResourceToMap(resource catalogmanagementv1.Resource) map[string]interface{} {
	resourceMap := map[string]interface{}{}

	resourceMap["type"] = resource.Type
	resourceMap["value"] = resource.Value

	return resourceMap
}

func resourceIBMCmVersionScriptToMap(script catalogmanagementv1.Script) map[string]interface{} {
	scriptMap := map[string]interface{}{}

	scriptMap["instructions"] = script.Instructions
	scriptMap["script"] = script.Script
	scriptMap["script_permission"] = script.ScriptPermission
	scriptMap["delete_script"] = script.DeleteScript
	scriptMap["scope"] = script.Scope

	return scriptMap
}

func resourceIBMCmVersionVersionEntitlementToMap(versionEntitlement catalogmanagementv1.VersionEntitlement) map[string]interface{} {
	versionEntitlementMap := map[string]interface{}{}

	versionEntitlementMap["provider_name"] = versionEntitlement.ProviderName
	versionEntitlementMap["provider_id"] = versionEntitlement.ProviderID
	versionEntitlementMap["product_id"] = versionEntitlement.ProductID
	if versionEntitlement.PartNumbers != nil {
		versionEntitlementMap["part_numbers"] = versionEntitlement.PartNumbers
	}
	versionEntitlementMap["image_repo_name"] = versionEntitlement.ImageRepoName

	return versionEntitlementMap
}

func resourceIBMCmVersionLicenseToMap(license catalogmanagementv1.License) map[string]interface{} {
	licenseMap := map[string]interface{}{}

	licenseMap["id"] = license.ID
	licenseMap["name"] = license.Name
	licenseMap["type"] = license.Type
	licenseMap["url"] = license.URL
	licenseMap["description"] = license.Description

	return licenseMap
}

func resourceIBMCmVersionStateToMap(state catalogmanagementv1.State) map[string]interface{} {
	stateMap := map[string]interface{}{}

	stateMap["current"] = state.Current
	stateMap["current_entered"] = state.CurrentEntered.String()
	stateMap["pending"] = state.Pending
	stateMap["pending_requested"] = state.PendingRequested.String()
	stateMap["previous"] = state.Previous

	return stateMap
}

func resourceIBMCmVersionDelete(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	deleteVersionOptions.SetVersionLocID(parts[2])

	response, err := catalogManagementClient.DeleteVersion(deleteVersionOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteVersion failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}
