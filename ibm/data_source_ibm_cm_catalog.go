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

func dataSourceIBMCmCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCmCatalogRead,
		Schema: map[string]*schema.Schema{
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloudant revision.",
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Display Name in the requested language.",
			},
			"short_description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Description in the requested language.",
			},
			"catalog_icon_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "URL for an icon associated with this catalog.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "List of tags associated with this catalog.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"features": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "List of features associated with this catalog.",
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
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Denotes whether a catalog is disabled.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Resource group id the catalog is owned by.",
			},
			"owning_account": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account that owns catalog.",
			},
			"catalog_filters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Filters for account and catalog filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_all": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "-> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.",
						},
						"category_filters": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Filter against offering properties.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"id_filters": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Offering filter terms.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"exclude": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Offering filter terms.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.",
													Elem:        &schema.Schema{Type: schema.TypeString},
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
			"syndication_settings": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Feature information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remove_related_components": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Remove related components.",
						},
						"clusters": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Syndication clusters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cluster region.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cluster ID.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cluster name.",
									},
									"resource_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Resource group ID.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Syndication type.",
									},
									"namespaces": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Syndicated namespaces.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"all_namespaces": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Syndicated to all namespaces on cluster.",
									},
								},
							},
						},
						"history": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Feature information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespaces": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of syndicated namespaces.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"clusters": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of syndicated namespaces.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Cluster region.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Cluster ID.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Cluster name.",
												},
												"resource_group_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Resource group ID.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Syndication type.",
												},
												"namespaces": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Syndicated namespaces.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"all_namespaces": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Syndicated to all namespaces on cluster.",
												},
											},
										},
									},
									"last_run": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Date and time last syndicated.",
									},
								},
							},
						},
						"authorization": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Feature information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Array of syndicated namespaces.",
									},
									"last_run": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Date and time last updated.",
									},
								},
							},
						},
					},
				},
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The url for this specific catalog.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN associated with the catalog.",
			},
			"offerings_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL path to offerings.",
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date-time this catalog was created.",
			},
			"updated": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date-time this catalog was last updated.",
			},
		},
	}
}

func dataSourceIBMCmCatalogRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

	getCatalogOptions.SetCatalogIdentifier(d.Id())

	catalog, response, err := catalogManagementClient.GetCatalogWithContext(context, getCatalogOptions)
	if err != nil {
		log.Printf("[DEBUG] GetCatalogWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*catalog.ID)
	if err = d.Set("label", catalog.Label); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting label: %s", err))
	}
	if err = d.Set("short_description", catalog.ShortDescription); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting short_description: %s", err))
	}
	if err = d.Set("catalog_icon_url", catalog.CatalogIconURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_icon_url: %s", err))
	}
	if err = d.Set("tags", catalog.Tags); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
	}
	if err = d.Set("url", catalog.URL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting url: %s", err))
	}
	if err = d.Set("crn", catalog.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("offerings_url", catalog.OfferingsURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting offerings_url: %s", err))
	}
	if err = d.Set("disabled", catalog.Disabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disabled: %s", err))
	}
	if err = d.Set("created", catalog.Created); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if err = d.Set("updated", catalog.Updated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}
	return nil
}
