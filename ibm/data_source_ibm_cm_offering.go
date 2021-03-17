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

func dataSourceIBMCmOffering() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCmOfferingRead,

		Schema: map[string]*schema.Schema{
			"catalog_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Catalog identifier.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Offering identification.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "unique id.",
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
				Computed:    true,
				Description: "Display Name in the requested language.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of tags associated with this catalog.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rating": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Repository info for offerings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"one_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "One start rating.",
						},
						"two_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Two start rating.",
						},
						"three_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Three start rating.",
						},
						"four_star_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
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
				Computed:    true,
				Description: "Short description in the requested language.",
			},
			"long_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
							Computed:    true,
							Description: "Heading.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
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
							Computed:    true,
							Description: "Unique ID.",
						},
						"format_kind": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "content kind, e.g., helm, vm image.",
						},
						"target_kind": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "target cloud to install, e.g., iks, open_shift_iks.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Open ended metadata information.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"install_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Installation instruction.",
						},
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of tags associated with this catalog.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"additional_features": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of features associated with this offering.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Heading.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature description.",
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
						"versions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "list of versions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
							},
						},
						"plans": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "list of plans.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "unique id.",
									},
									"label": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Display Name in the requested language.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The programmatic name of this offering.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Short description in the requested language.",
									},
									"long_description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Long description in the requested language.",
									},
									"metadata": &schema.Schema{
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "open ended metadata information.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"tags": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "list of tags associated with this catalog.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"additional_features": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "list of features associated with this offering.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"title": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Heading.",
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Feature description.",
												},
											},
										},
									},
									"created": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "the date'time this catalog was created.",
									},
									"updated": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "the date'time this catalog was last updated.",
									},
									"deployments": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "list of deployments.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "unique id.",
												},
												"label": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Display Name in the requested language.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The programmatic name of this offering.",
												},
												"short_description": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Short description in the requested language.",
												},
												"long_description": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Long description in the requested language.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "open ended metadata information.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"tags": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "list of tags associated with this catalog.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"created": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "the date'time this catalog was created.",
												},
												"updated": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
				MaxItems:    1,
				Computed:    true,
				Description: "Repository info for offerings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Token for private repos.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public or enterprise GitHub.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMCmOfferingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

	getOfferingOptions.SetCatalogIdentifier(d.Get("catalog_identifier").(string))
	getOfferingOptions.SetOfferingID(d.Get("offering_id").(string))

	offering, response, err := catalogManagementClient.GetOfferingWithContext(context, getOfferingOptions)
	if err != nil {
		log.Printf("[DEBUG] GetOfferingWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", *offering.CatalogID, *offering.ID))
	if err = d.Set("id", offering.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}
	if err = d.Set("rev", offering.Rev); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rev: %s", err))
	}
	if err = d.Set("url", offering.URL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting url: %s", err))
	}
	if err = d.Set("crn", offering.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("label", offering.Label); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting label: %s", err))
	}
	if err = d.Set("name", offering.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("offering_icon_url", offering.OfferingIconURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting offering_icon_url: %s", err))
	}
	if err = d.Set("offering_docs_url", offering.OfferingDocsURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting offering_docs_url: %s", err))
	}
	if err = d.Set("offering_support_url", offering.OfferingSupportURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting offering_support_url: %s", err))
	}
	if err = d.Set("tags", offering.Tags); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
	}

	if offering.Rating != nil {
		err = d.Set("rating", dataSourceOfferingFlattenRating(*offering.Rating))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting rating %s", err))
		}
	}
	if err = d.Set("created", offering.Created); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if err = d.Set("updated", offering.Updated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}
	if err = d.Set("short_description", offering.ShortDescription); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting short_description: %s", err))
	}
	if err = d.Set("long_description", offering.LongDescription); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting long_description: %s", err))
	}

	if offering.Features != nil {
		err = d.Set("features", dataSourceOfferingFlattenFeatures(offering.Features))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting features %s", err))
		}
	}

	if offering.Kinds != nil {
		err = d.Set("kinds", dataSourceOfferingFlattenKinds(offering.Kinds))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting kinds %s", err))
		}
	}
	if err = d.Set("permit_request_ibm_public_publish", offering.PermitRequestIBMPublicPublish); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting permit_request_ibm_public_publish: %s", err))
	}
	if err = d.Set("ibm_publish_approved", offering.IBMPublishApproved); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting ibm_publish_approved: %s", err))
	}
	if err = d.Set("public_publish_approved", offering.PublicPublishApproved); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting public_publish_approved: %s", err))
	}
	if err = d.Set("public_original_crn", offering.PublicOriginalCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting public_original_crn: %s", err))
	}
	if err = d.Set("publish_public_crn", offering.PublishPublicCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting publish_public_crn: %s", err))
	}
	if err = d.Set("portal_approval_record", offering.PortalApprovalRecord); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting portal_approval_record: %s", err))
	}
	if err = d.Set("portal_ui_url", offering.PortalUIURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting portal_ui_url: %s", err))
	}
	if err = d.Set("catalog_id", offering.CatalogID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_id: %s", err))
	}
	if err = d.Set("catalog_name", offering.CatalogName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_name: %s", err))
	}

	if offering.Metadata != nil {
		convertedMap := make(map[string]interface{}, len(offering.Metadata))
		for k, v := range offering.Metadata {
			convertedMap[k] = v
		}

		if err = d.Set("metadata", Flatten(convertedMap)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata: %s", err))
		}
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata %s", err))
		}
	}
	if err = d.Set("disclaimer", offering.Disclaimer); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disclaimer: %s", err))
	}
	if err = d.Set("hidden", offering.Hidden); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting hidden: %s", err))
	}

	if offering.RepoInfo != nil {
		err = d.Set("repo_info", dataSourceOfferingFlattenRepoInfo(*offering.RepoInfo))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting repo_info %s", err))
		}
	}

	return nil
}

func dataSourceOfferingFlattenRating(result catalogmanagementv1.Rating) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceOfferingRatingToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceOfferingRatingToMap(ratingItem catalogmanagementv1.Rating) (ratingMap map[string]interface{}) {
	ratingMap = map[string]interface{}{}

	if ratingItem.OneStarCount != nil {
		ratingMap["one_star_count"] = ratingItem.OneStarCount
	}
	if ratingItem.TwoStarCount != nil {
		ratingMap["two_star_count"] = ratingItem.TwoStarCount
	}
	if ratingItem.ThreeStarCount != nil {
		ratingMap["three_star_count"] = ratingItem.ThreeStarCount
	}
	if ratingItem.FourStarCount != nil {
		ratingMap["four_star_count"] = ratingItem.FourStarCount
	}

	return ratingMap
}

func dataSourceOfferingFlattenFeatures(result []catalogmanagementv1.Feature) (features []map[string]interface{}) {
	for _, featuresItem := range result {
		features = append(features, dataSourceOfferingFeaturesToMap(featuresItem))
	}

	return features
}

func dataSourceOfferingFeaturesToMap(featuresItem catalogmanagementv1.Feature) (featuresMap map[string]interface{}) {
	featuresMap = map[string]interface{}{}

	if featuresItem.Title != nil {
		featuresMap["title"] = featuresItem.Title
	}
	if featuresItem.Description != nil {
		featuresMap["description"] = featuresItem.Description
	}

	return featuresMap
}

func dataSourceOfferingFlattenKinds(result []catalogmanagementv1.Kind) (kinds []map[string]interface{}) {
	for _, kindsItem := range result {
		kinds = append(kinds, dataSourceOfferingKindsToMap(kindsItem))
	}

	return kinds
}

func dataSourceOfferingKindsToMap(kindsItem catalogmanagementv1.Kind) (kindsMap map[string]interface{}) {
	kindsMap = map[string]interface{}{}

	if kindsItem.ID != nil {
		kindsMap["id"] = kindsItem.ID
	}
	if kindsItem.FormatKind != nil {
		kindsMap["format_kind"] = kindsItem.FormatKind
	}
	if kindsItem.TargetKind != nil {
		kindsMap["target_kind"] = kindsItem.TargetKind
	}
	if kindsItem.Metadata != nil {
		convertedMap := make(map[string]interface{}, len(kindsItem.Metadata))
		for k, v := range kindsItem.Metadata {
			convertedMap[k] = v
		}
		kindsMap["metadata"] = Flatten(convertedMap)
	}
	if kindsItem.InstallDescription != nil {
		kindsMap["install_description"] = kindsItem.InstallDescription
	}
	if kindsItem.Tags != nil {
		kindsMap["tags"] = kindsItem.Tags
	}
	if kindsItem.AdditionalFeatures != nil {
		additionalFeaturesList := []map[string]interface{}{}
		for _, additionalFeaturesItem := range kindsItem.AdditionalFeatures {
			additionalFeaturesList = append(additionalFeaturesList, dataSourceOfferingKindsAdditionalFeaturesToMap(additionalFeaturesItem))
		}
		kindsMap["additional_features"] = additionalFeaturesList
	}
	if kindsItem.Created != nil {
		kindsMap["created"] = kindsItem.Created.String()
	}
	if kindsItem.Updated != nil {
		kindsMap["updated"] = kindsItem.Updated.String()
	}
	if kindsItem.Versions != nil {
		versionsList := []map[string]interface{}{}
		for _, versionsItem := range kindsItem.Versions {
			versionsList = append(versionsList, dataSourceOfferingKindsVersionsToMap(versionsItem))
		}
		kindsMap["versions"] = versionsList
	}
	if kindsItem.Plans != nil {
		plansList := []map[string]interface{}{}
		for _, plansItem := range kindsItem.Plans {
			plansList = append(plansList, dataSourceOfferingKindsPlansToMap(plansItem))
		}
		kindsMap["plans"] = plansList
	}

	return kindsMap
}

func dataSourceOfferingKindsAdditionalFeaturesToMap(additionalFeaturesItem catalogmanagementv1.Feature) (additionalFeaturesMap map[string]interface{}) {
	additionalFeaturesMap = map[string]interface{}{}

	if additionalFeaturesItem.Title != nil {
		additionalFeaturesMap["title"] = additionalFeaturesItem.Title
	}
	if additionalFeaturesItem.Description != nil {
		additionalFeaturesMap["description"] = additionalFeaturesItem.Description
	}

	return additionalFeaturesMap
}

func dataSourceOfferingKindsVersionsToMap(versionsItem catalogmanagementv1.Version) (versionsMap map[string]interface{}) {
	versionsMap = map[string]interface{}{}

	if versionsItem.ID != nil {
		versionsMap["id"] = versionsItem.ID
	}
	if versionsItem.Rev != nil {
		versionsMap["rev"] = versionsItem.Rev
	}
	if versionsItem.CRN != nil {
		versionsMap["crn"] = versionsItem.CRN
	}
	if versionsItem.Version != nil {
		versionsMap["version"] = versionsItem.Version
	}
	return versionsMap
}

func dataSourceOfferingKindsPlansToMap(plansItem catalogmanagementv1.Plan) (plansMap map[string]interface{}) {
	plansMap = map[string]interface{}{}

	if plansItem.ID != nil {
		plansMap["id"] = plansItem.ID
	}
	if plansItem.Label != nil {
		plansMap["label"] = plansItem.Label
	}
	if plansItem.Name != nil {
		plansMap["name"] = plansItem.Name
	}
	if plansItem.ShortDescription != nil {
		plansMap["short_description"] = plansItem.ShortDescription
	}
	if plansItem.LongDescription != nil {
		plansMap["long_description"] = plansItem.LongDescription
	}
	if plansItem.Metadata != nil {
		convertedMap := make(map[string]interface{}, len(plansItem.Metadata))
		for k, v := range plansItem.Metadata {
			convertedMap[k] = v
		}
		plansMap["metadata"] = Flatten(convertedMap)
	}
	if plansItem.Tags != nil {
		plansMap["tags"] = plansItem.Tags
	}
	if plansItem.AdditionalFeatures != nil {
		additionalFeaturesList := []map[string]interface{}{}
		for _, additionalFeaturesItem := range plansItem.AdditionalFeatures {
			additionalFeaturesList = append(additionalFeaturesList, dataSourceOfferingKindsAdditionalFeaturesToMap(additionalFeaturesItem))
		}
		plansMap["additional_features"] = additionalFeaturesList
	}
	if plansItem.Created != nil {
		plansMap["created"] = plansItem.Created.String()
	}
	if plansItem.Updated != nil {
		plansMap["updated"] = plansItem.Updated.String()
	}
	if plansItem.Deployments != nil {
		deploymentsList := []map[string]interface{}{}
		for _, deploymentsItem := range plansItem.Deployments {
			deploymentsList = append(deploymentsList, dataSourceOfferingPlansDeploymentsToMap(deploymentsItem))
		}
		plansMap["deployments"] = deploymentsList
	}

	return plansMap
}

func dataSourceOfferingPlansDeploymentsToMap(deploymentsItem catalogmanagementv1.Deployment) (deploymentsMap map[string]interface{}) {
	deploymentsMap = map[string]interface{}{}

	if deploymentsItem.ID != nil {
		deploymentsMap["id"] = deploymentsItem.ID
	}

	if deploymentsItem.Label != nil {
		deploymentsMap["label"] = deploymentsItem.Label
	}

	if deploymentsItem.Name != nil {
		deploymentsMap["name"] = deploymentsItem.ID
	}

	if deploymentsItem.ShortDescription != nil {
		deploymentsMap["shortDescription"] = deploymentsItem.ShortDescription
	}

	if deploymentsItem.LongDescription != nil {
		deploymentsMap["longDescription"] = deploymentsItem.LongDescription
	}

	return deploymentsMap
}

func dataSourceOfferingFlattenRepoInfo(result catalogmanagementv1.RepoInfo) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceOfferingRepoInfoToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceOfferingRepoInfoToMap(repoInfoItem catalogmanagementv1.RepoInfo) (repoInfoMap map[string]interface{}) {
	repoInfoMap = map[string]interface{}{}

	if repoInfoItem.Token != nil {
		repoInfoMap["token"] = repoInfoItem.Token
	}
	if repoInfoItem.Type != nil {
		repoInfoMap["type"] = repoInfoItem.Type
	}

	return repoInfoMap
}
