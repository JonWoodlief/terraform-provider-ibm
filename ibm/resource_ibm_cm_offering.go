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

func resourceIBMCmOffering() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCmOfferingCreate,
		Read:     resourceIBMCmOfferingRead,
		Delete:   resourceIBMCmOfferingDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"catalog_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Catalog identifier.",
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
				Description: "URL path to zip location.  If not specified, must provide content in this post body.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "Re-use the specified offeringID during import.",
			},
			"target_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The semver value for this new version.",
			},
			"include_config": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Add all possible configuration items when creating this version.",
			},
			"repo_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'.",
			},
			"x_auth_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Authentication token used to access the specified zip file.",
			},
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloudant revision.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The url for this specific offering.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn for this specific offering.",
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display Name in the requested language.",
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The programmatic name of this offering.",
			},
			"offering_icon_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for an icon associated with this offering.",
			},
			"offering_docs_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for an additional docs with this offering.",
			},
			"offering_support_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to be displayed in the Consumption UI for getting support on this offering.",
			},
			"rating": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Repository info for offerings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"one_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "One start rating.",
						},
						"two_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Two start rating.",
						},
						"three_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Three start rating.",
						},
						"four_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Four start rating.",
						},
					},
				},
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time this catalog was created.",
			},
			"updated": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time this catalog was last updated.",
			},
			"short_description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Short description in the requested language.",
			},
			"long_description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Long description in the requested language.",
			},
			"features": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of features associated with this offering.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Heading.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Feature description.",
						},
					},
				},
			},
			"kinds": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of kind.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique ID.",
						},
						"format_kind": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "content kind, e.g., helm, vm image.",
						},
						"target_kind": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "target cloud to install, e.g., iks, open_shift_iks.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Open ended metadata information.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"install_description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Installation instruction.",
						},
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of tags associated with this catalog.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"additional_features": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of features associated with this offering.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Heading.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Feature description.",
									},
								},
							},
						},
						"created": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The date and time this catalog was created.",
						},
						"updated": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The date and time this catalog was last updated.",
						},
						"versions": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "list of versions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique ID.",
									},
									"rev": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloudant revision.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Version's CRN.",
									},
									"version": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Version of content type.",
									},
									"sha": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "hash of the content.",
									},
									"created": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The date and time this version was created.",
									},
									"updated": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The date and time this version was last updated.",
									},
									"offering_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Offering ID.",
									},
									"catalog_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Catalog ID.",
									},
									"kind_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Kind ID.",
									},
									"tags": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of tags associated with this catalog.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Content's repo URL.",
									},
									"source_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Content's source URL (e.g git repo).",
									},
									"tgz_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "File used to on-board this version.",
									},
									"configuration": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
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
										Optional:    true,
										Description: "Open ended metadata information.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"validation": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
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
										Optional:    true,
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
										Optional:    true,
										Description: "Denotes if single instance can be deployed to a given cluster.",
									},
									"install": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
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
										Optional:    true,
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
										MaxItems:    1,
										Optional:    true,
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
										Optional:    true,
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
										Optional:    true,
										Description: "If set, denotes a url to a YAML file with list of container images used by this version.",
									},
									"deprecated": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "read only field, indicating if this version is deprecated.",
									},
									"package_version": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Version of the package used to create this version.",
									},
									"state": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
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
										Optional:    true,
										Description: "A dotted value of `catalogID`.`versionID`.",
									},
									"console_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Console URL.",
									},
									"long_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Long description for version.",
									},
									"whitelisted_accounts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Whitelisted accounts for version.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"plans": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "list of plans.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "unique id.",
									},
									"label": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Display Name in the requested language.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The programmatic name of this offering.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Short description in the requested language.",
									},
									"long_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Long description in the requested language.",
									},
									"metadata": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "open ended metadata information.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"tags": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "list of tags associated with this catalog.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"additional_features": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "list of features associated with this offering.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"title": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Heading.",
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Feature description.",
												},
											},
										},
									},
									"created": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "the date'time this catalog was created.",
									},
									"updated": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "the date'time this catalog was last updated.",
									},
									"deployments": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "list of deployments.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "unique id.",
												},
												"label": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Display Name in the requested language.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The programmatic name of this offering.",
												},
												"short_description": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Short description in the requested language.",
												},
												"long_description": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Long description in the requested language.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "open ended metadata information.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"tags": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "list of tags associated with this catalog.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"created": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "the date'time this catalog was created.",
												},
												"updated": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "the date'time this catalog was last updated.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"permit_request_ibm_public_publish": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is it permitted to request publishing to IBM or Public.",
			},
			"ibm_publish_approved": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if this offering has been approved for use by all IBMers.",
			},
			"public_publish_approved": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if this offering has been approved for use by all IBM Cloud users.",
			},
			"public_original_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The original offering CRN that this publish entry came from.",
			},
			"publish_public_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the public catalog entry of this offering.",
			},
			"portal_approval_record": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The portal's approval record ID.",
			},
			"portal_ui_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The portal UI URL.",
			},
			"catalog_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the catalog containing this offering.",
			},
			"catalog_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the catalog.",
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map of metadata values for this offering.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"disclaimer": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A disclaimer for this offering.",
			},
			"hidden": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determine if this offering should be displayed in the Consumption UI.",
			},
			"repo_info": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Repository info for offerings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Token for private repos.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Public or enterprise GitHub.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMCmOfferingCreate(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	importOfferingOptions := &catalogmanagementv1.ImportOfferingOptions{}

	importOfferingOptions.SetCatalogIdentifier(d.Get("catalog_identifier").(string))
	if _, ok := d.GetOk("tags"); ok {
		importOfferingOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("target_kinds"); ok {
		importOfferingOptions.SetTargetKinds(d.Get("target_kinds").([]string))
	}
	if _, ok := d.GetOk("content"); ok {

	}
	if _, ok := d.GetOk("zipurl"); ok {
		importOfferingOptions.SetZipurl(d.Get("zipurl").(string))
	}
	if _, ok := d.GetOk("offering_id"); ok {
		importOfferingOptions.SetOfferingID(d.Get("offering_id").(string))
	}
	if _, ok := d.GetOk("target_version"); ok {
		importOfferingOptions.SetTargetVersion(d.Get("target_version").(string))
	}
	if _, ok := d.GetOk("include_config"); ok {
		importOfferingOptions.SetIncludeConfig(d.Get("include_config").(bool))
	}
	if _, ok := d.GetOk("repo_type"); ok {
		importOfferingOptions.SetRepoType(d.Get("repo_type").(string))
	}
	if _, ok := d.GetOk("x_auth_token"); ok {
		importOfferingOptions.SetXAuthToken(d.Get("x_auth_token").(string))
	}

	offering, response, err := catalogManagementClient.ImportOffering(importOfferingOptions)
	if err != nil {
		log.Printf("[DEBUG] ImportOfferingWithContext failed %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", *importOfferingOptions.CatalogIdentifier, *offering.ID))

	return resourceIBMCmOfferingRead(d, meta)
}

func resourceIBMCmOfferingRead(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	getOfferingOptions.SetCatalogIdentifier(parts[0])
	getOfferingOptions.SetOfferingID(parts[1])

	offering, response, err := catalogManagementClient.GetOffering(getOfferingOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetOfferingWithContext failed %s\n%s", err, response)
		return err
	}
	if err = d.Set("offering_id", offering.ID); err != nil {
		return fmt.Errorf("Error setting offering_id: %s", err)
	}
	if err = d.Set("url", offering.URL); err != nil {
		return fmt.Errorf("Error setting url: %s", err)
	}
	if err = d.Set("crn", offering.CRN); err != nil {
		return fmt.Errorf("Error setting crn: %s", err)
	}
	if err = d.Set("label", offering.Label); err != nil {
		return fmt.Errorf("Error setting label: %s", err)
	}
	if err = d.Set("name", offering.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err = d.Set("offering_icon_url", offering.OfferingIconURL); err != nil {
		return fmt.Errorf("Error setting offering_icon_url: %s", err)
	}
	if err = d.Set("offering_docs_url", offering.OfferingDocsURL); err != nil {
		return fmt.Errorf("Error setting offering_docs_url: %s", err)
	}
	if err = d.Set("offering_support_url", offering.OfferingSupportURL); err != nil {
		return fmt.Errorf("Error setting offering_support_url: %s", err)
	}
	if offering.Rating != nil {
		ratingMap := resourceIBMCmOfferingRatingToMap(*offering.Rating)
		if err = d.Set("rating", []map[string]interface{}{ratingMap}); err != nil {
			return fmt.Errorf("Error setting rating: %s", err)
		}
	}
	if err = d.Set("created", offering.Created.String()); err != nil {
		return fmt.Errorf("Error setting created: %s", err)
	}
	if err = d.Set("updated", offering.Updated.String()); err != nil {
		return fmt.Errorf("Error setting updated: %s", err)
	}
	if err = d.Set("short_description", offering.ShortDescription); err != nil {
		return fmt.Errorf("Error setting short_description: %s", err)
	}
	if err = d.Set("long_description", offering.LongDescription); err != nil {
		return fmt.Errorf("Error setting long_description: %s", err)
	}
	if offering.Features != nil {
		features := []map[string]interface{}{}
		for _, featuresItem := range offering.Features {
			featuresItemMap := resourceIBMCmOfferingFeatureToMap(featuresItem)
			features = append(features, featuresItemMap)
		}
		if err = d.Set("features", features); err != nil {
			return fmt.Errorf("Error setting features: %s", err)
		}
	}
	if offering.Kinds != nil {
		kinds := []map[string]interface{}{}
		for _, kindsItem := range offering.Kinds {
			kindsItemMap := resourceIBMCmOfferingKindToMap(kindsItem)
			kinds = append(kinds, kindsItemMap)
		}
		if err = d.Set("kinds", kinds); err != nil {
			return fmt.Errorf("Error setting kinds: %s", err)
		}
	}
	if err = d.Set("permit_request_ibm_public_publish", offering.PermitRequestIBMPublicPublish); err != nil {
		return fmt.Errorf("Error setting permit_request_ibm_public_publish: %s", err)
	}
	if err = d.Set("ibm_publish_approved", offering.IBMPublishApproved); err != nil {
		return fmt.Errorf("Error setting ibm_publish_approved: %s", err)
	}
	if err = d.Set("public_publish_approved", offering.PublicPublishApproved); err != nil {
		return fmt.Errorf("Error setting public_publish_approved: %s", err)
	}
	if err = d.Set("public_original_crn", offering.PublicOriginalCRN); err != nil {
		return fmt.Errorf("Error setting public_original_crn: %s", err)
	}
	if err = d.Set("publish_public_crn", offering.PublishPublicCRN); err != nil {
		return fmt.Errorf("Error setting publish_public_crn: %s", err)
	}
	if err = d.Set("portal_approval_record", offering.PortalApprovalRecord); err != nil {
		return fmt.Errorf("Error setting portal_approval_record: %s", err)
	}
	if err = d.Set("portal_ui_url", offering.PortalUIURL); err != nil {
		return fmt.Errorf("Error setting portal_ui_url: %s", err)
	}
	if err = d.Set("catalog_id", offering.CatalogID); err != nil {
		return fmt.Errorf("Error setting catalog_id: %s", err)
	}
	if err = d.Set("catalog_name", offering.CatalogName); err != nil {
		return fmt.Errorf("Error setting catalog_name: %s", err)
	}
	if offering.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- not primitive type, not list
	}
	if err = d.Set("disclaimer", offering.Disclaimer); err != nil {
		return fmt.Errorf("Error setting disclaimer: %s", err)
	}
	if err = d.Set("hidden", offering.Hidden); err != nil {
		return fmt.Errorf("Error setting hidden: %s", err)
	}
	if offering.RepoInfo != nil {
		repoInfoMap := resourceIBMCmOfferingRepoInfoToMap(*offering.RepoInfo)
		if err = d.Set("repo_info", []map[string]interface{}{repoInfoMap}); err != nil {
			return fmt.Errorf("Error setting repo_info: %s", err)
		}
	}

	return nil
}

func resourceIBMCmOfferingRatingToMap(rating catalogmanagementv1.Rating) map[string]interface{} {
	ratingMap := map[string]interface{}{}

	ratingMap["one_star_count"] = intValue(rating.OneStarCount)
	ratingMap["two_star_count"] = intValue(rating.TwoStarCount)
	ratingMap["three_star_count"] = intValue(rating.ThreeStarCount)
	ratingMap["four_star_count"] = intValue(rating.FourStarCount)

	return ratingMap
}

func resourceIBMCmOfferingFeatureToMap(feature catalogmanagementv1.Feature) map[string]interface{} {
	featureMap := map[string]interface{}{}

	featureMap["title"] = feature.Title
	featureMap["description"] = feature.Description

	return featureMap
}

func resourceIBMCmOfferingKindToMap(kind catalogmanagementv1.Kind) map[string]interface{} {
	kindMap := map[string]interface{}{}

	kindMap["id"] = kind.ID
	kindMap["format_kind"] = kind.FormatKind
	kindMap["target_kind"] = kind.TargetKind
	if kind.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- container, not list
	}
	kindMap["install_description"] = kind.InstallDescription
	if kind.Tags != nil {
		kindMap["tags"] = kind.Tags
	}
	if kind.AdditionalFeatures != nil {
		additionalFeatures := []map[string]interface{}{}
		for _, additionalFeaturesItem := range kind.AdditionalFeatures {
			additionalFeaturesItemMap := resourceIBMCmOfferingFeatureToMap(additionalFeaturesItem)
			additionalFeatures = append(additionalFeatures, additionalFeaturesItemMap)
			// TODO: handle AdditionalFeatures of type TypeList -- list of non-primitive, not model items
		}
		kindMap["additional_features"] = additionalFeatures
	}
	kindMap["created"] = kind.Created.String()
	kindMap["updated"] = kind.Updated.String()
	if kind.Versions != nil {
		versions := []map[string]interface{}{}
		for _, versionsItem := range kind.Versions {
			versionsItemMap := resourceIBMCmOfferingVersionToMap(versionsItem)
			versions = append(versions, versionsItemMap)
			// TODO: handle Versions of type TypeList -- list of non-primitive, not model items
		}
		kindMap["versions"] = versions
	}
	if kind.Plans != nil {
		plans := []map[string]interface{}{}
		for _, plansItem := range kind.Plans {
			plansItemMap := resourceIBMCmOfferingPlanToMap(plansItem)
			plans = append(plans, plansItemMap)
			// TODO: handle Plans of type TypeList -- list of non-primitive, not model items
		}
		kindMap["plans"] = plans
	}

	return kindMap
}

func resourceIBMCmOfferingVersionToMap(version catalogmanagementv1.Version) map[string]interface{} {
	versionMap := map[string]interface{}{}

	versionMap["id"] = version.ID
	versionMap["_rev"] = version.Rev
	versionMap["crn"] = version.CRN
	versionMap["version"] = version.Version
	versionMap["sha"] = version.Sha
	versionMap["created"] = version.Created.String()
	versionMap["updated"] = version.Updated.String()
	versionMap["offering_id"] = version.OfferingID
	versionMap["catalog_id"] = version.CatalogID
	versionMap["kind_id"] = version.KindID
	if version.Tags != nil {
		versionMap["tags"] = version.Tags
	}
	versionMap["repo_url"] = version.RepoURL
	versionMap["source_url"] = version.SourceURL
	versionMap["tgz_url"] = version.TgzURL
	if version.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- container, not list
	}
	if version.Validation != nil {
		ValidationMap := resourceIBMCmOfferingValidationToMap(*version.Validation)
		versionMap["validation"] = []map[string]interface{}{ValidationMap}
	}
	if version.RequiredResources != nil {
		requiredResources := []map[string]interface{}{}
		for _, requiredResourcesItem := range version.RequiredResources {
			requiredResourcesItemMap := resourceIBMCmOfferingResourceToMap(requiredResourcesItem)
			requiredResources = append(requiredResources, requiredResourcesItemMap)
			// TODO: handle RequiredResources of type TypeList -- list of non-primitive, not model items
		}
		versionMap["required_resources"] = requiredResources
	}
	versionMap["single_instance"] = version.SingleInstance
	if version.Install != nil {
		InstallMap := resourceIBMCmOfferingScriptToMap(*version.Install)
		versionMap["install"] = []map[string]interface{}{InstallMap}
	}
	if version.PreInstall != nil {
		preInstall := []map[string]interface{}{}
		for _, preInstallItem := range version.PreInstall {
			preInstallItemMap := resourceIBMCmOfferingScriptToMap(preInstallItem)
			preInstall = append(preInstall, preInstallItemMap)
			// TODO: handle PreInstall of type TypeList -- list of non-primitive, not model items
		}
		versionMap["pre_install"] = preInstall
	}
	if version.Entitlement != nil {
		EntitlementMap := resourceIBMCmOfferingVersionEntitlementToMap(*version.Entitlement)
		versionMap["entitlement"] = []map[string]interface{}{EntitlementMap}
	}
	if version.Licenses != nil {
		licenses := []map[string]interface{}{}
		for _, licensesItem := range version.Licenses {
			licensesItemMap := resourceIBMCmOfferingLicenseToMap(licensesItem)
			licenses = append(licenses, licensesItemMap)
			// TODO: handle Licenses of type TypeList -- list of non-primitive, not model items
		}
		versionMap["licenses"] = licenses
	}
	versionMap["image_manifest_url"] = version.ImageManifestURL
	versionMap["deprecated"] = version.Deprecated
	versionMap["package_version"] = version.PackageVersion
	if version.State != nil {
		StateMap := resourceIBMCmOfferingStateToMap(*version.State)
		versionMap["state"] = []map[string]interface{}{StateMap}
	}
	versionMap["version_locator"] = version.VersionLocator
	versionMap["console_url"] = version.ConsoleURL
	versionMap["long_description"] = version.LongDescription
	if version.WhitelistedAccounts != nil {
		versionMap["whitelisted_accounts"] = version.WhitelistedAccounts
	}

	return versionMap
}

func resourceIBMCmOfferingConfigurationToMap(configuration catalogmanagementv1.Configuration) map[string]interface{} {
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

func resourceIBMCmOfferingValidationToMap(validation catalogmanagementv1.Validation) map[string]interface{} {
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

func resourceIBMCmOfferingResourceToMap(resource catalogmanagementv1.Resource) map[string]interface{} {
	resourceMap := map[string]interface{}{}

	resourceMap["type"] = resource.Type
	resourceMap["value"] = resource.Value

	return resourceMap
}

func resourceIBMCmOfferingScriptToMap(script catalogmanagementv1.Script) map[string]interface{} {
	scriptMap := map[string]interface{}{}

	scriptMap["instructions"] = script.Instructions
	scriptMap["script"] = script.Script
	scriptMap["script_permission"] = script.ScriptPermission
	scriptMap["delete_script"] = script.DeleteScript
	scriptMap["scope"] = script.Scope

	return scriptMap
}

func resourceIBMCmOfferingVersionEntitlementToMap(versionEntitlement catalogmanagementv1.VersionEntitlement) map[string]interface{} {
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

func resourceIBMCmOfferingLicenseToMap(license catalogmanagementv1.License) map[string]interface{} {
	licenseMap := map[string]interface{}{}

	licenseMap["id"] = license.ID
	licenseMap["name"] = license.Name
	licenseMap["type"] = license.Type
	licenseMap["url"] = license.URL
	licenseMap["description"] = license.Description

	return licenseMap
}

func resourceIBMCmOfferingStateToMap(state catalogmanagementv1.State) map[string]interface{} {
	stateMap := map[string]interface{}{}

	stateMap["current"] = state.Current
	stateMap["current_entered"] = state.CurrentEntered.String()
	stateMap["pending"] = state.Pending
	stateMap["pending_requested"] = state.PendingRequested.String()
	stateMap["previous"] = state.Previous

	return stateMap
}

func resourceIBMCmOfferingPlanToMap(plan catalogmanagementv1.Plan) map[string]interface{} {
	planMap := map[string]interface{}{}

	planMap["id"] = plan.ID
	planMap["label"] = plan.Label
	planMap["name"] = plan.Name
	planMap["short_description"] = plan.ShortDescription
	planMap["long_description"] = plan.LongDescription
	if plan.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- container, not list
	}
	if plan.Tags != nil {
		planMap["tags"] = plan.Tags
	}
	if plan.AdditionalFeatures != nil {
		additionalFeatures := []map[string]interface{}{}
		for _, additionalFeaturesItem := range plan.AdditionalFeatures {
			additionalFeaturesItemMap := resourceIBMCmOfferingFeatureToMap(additionalFeaturesItem)
			additionalFeatures = append(additionalFeatures, additionalFeaturesItemMap)
			// TODO: handle AdditionalFeatures of type TypeList -- list of non-primitive, not model items
		}
		planMap["additional_features"] = additionalFeatures
	}
	planMap["created"] = plan.Created.String()
	planMap["updated"] = plan.Updated.String()
	if plan.Deployments != nil {
		deployments := []map[string]interface{}{}
		for _, deploymentsItem := range plan.Deployments {
			deploymentsItemMap := resourceIBMCmOfferingDeploymentToMap(deploymentsItem)
			deployments = append(deployments, deploymentsItemMap)
			// TODO: handle Deployments of type TypeList -- list of non-primitive, not model items
		}
		planMap["deployments"] = deployments
	}

	return planMap
}

func resourceIBMCmOfferingDeploymentToMap(deployment catalogmanagementv1.Deployment) map[string]interface{} {
	deploymentMap := map[string]interface{}{}

	deploymentMap["id"] = deployment.ID
	deploymentMap["label"] = deployment.Label
	deploymentMap["name"] = deployment.Name
	deploymentMap["short_description"] = deployment.ShortDescription
	deploymentMap["long_description"] = deployment.LongDescription
	if deployment.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- container, not list
	}
	if deployment.Tags != nil {
		deploymentMap["tags"] = deployment.Tags
	}
	deploymentMap["created"] = deployment.Created.String()
	deploymentMap["updated"] = deployment.Updated.String()

	return deploymentMap
}

func resourceIBMCmOfferingRepoInfoToMap(repoInfo catalogmanagementv1.RepoInfo) map[string]interface{} {
	repoInfoMap := map[string]interface{}{}

	repoInfoMap["token"] = repoInfo.Token
	repoInfoMap["type"] = repoInfo.Type

	return repoInfoMap
}

func resourceIBMCmOfferingDelete(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	deleteOfferingOptions := &catalogmanagementv1.DeleteOfferingOptions{}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	deleteOfferingOptions.SetCatalogIdentifier(parts[0])
	deleteOfferingOptions.SetOfferingID(parts[1])

	response, err := catalogManagementClient.DeleteOffering(deleteOfferingOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteOfferingWithContext failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}
