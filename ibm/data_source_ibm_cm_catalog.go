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
			"catalog_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Catalog identifier.",
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
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Display Name in the requested language.",
			},
			"short_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description in the requested language.",
			},
			"catalog_icon_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for an icon associated with this catalog.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of tags associated with this catalog.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"features": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of features associated with this catalog.",
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
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Denotes whether a catalog is disabled.",
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
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
				Computed:    true,
				Description: "Filters for account and catalog filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_all": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "-> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.",
						},
						"category_filters": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Filter against offering properties.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id_filters": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Offering filter terms.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"exclude": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Offering filter terms.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
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
			},
			"syndication_settings": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Feature information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remove_related_components": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Remove related components.",
						},
						"clusters": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Syndication clusters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster region.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster ID.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster name.",
									},
									"resource_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource group ID.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Syndication type.",
									},
									"namespaces": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Syndicated namespaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"all_namespaces": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Syndicated to all namespaces on cluster.",
									},
								},
							},
						},
						"history": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Feature information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespaces": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Array of syndicated namespaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Array of syndicated namespaces.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cluster region.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cluster ID.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cluster name.",
												},
												"resource_group_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Resource group ID.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Syndication type.",
												},
												"namespaces": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Syndicated namespaces.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"all_namespaces": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Syndicated to all namespaces on cluster.",
												},
											},
										},
									},
									"last_run": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time last syndicated.",
									},
								},
							},
						},
						"authorization": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Feature information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Array of syndicated namespaces.",
									},
									"last_run": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time last updated.",
									},
								},
							},
						},
					},
				},
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

	getCatalogOptions.SetCatalogIdentifier(d.Get("catalog_identifier").(string))

	catalog, response, err := catalogManagementClient.GetCatalogWithContext(context, getCatalogOptions)
	if err != nil {
		log.Printf("[DEBUG] GetCatalogWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*catalog.ID)
	if err = d.Set("id", catalog.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}
	if err = d.Set("rev", catalog.Rev); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rev: %s", err))
	}
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

	if catalog.Features != nil {
		err = d.Set("features", dataSourceCatalogFlattenFeatures(catalog.Features))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting features %s", err))
		}
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
	if err = d.Set("resource_group_id", catalog.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("owning_account", catalog.OwningAccount); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting owning_account: %s", err))
	}

	if catalog.CatalogFilters != nil {
		err = d.Set("catalog_filters", dataSourceCatalogFlattenCatalogFilters(*catalog.CatalogFilters))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting catalog_filters %s", err))
		}
	}

	if catalog.SyndicationSettings != nil {
		err = d.Set("syndication_settings", dataSourceCatalogFlattenSyndicationSettings(*catalog.SyndicationSettings))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting syndication_settings %s", err))
		}
	}

	return nil
}

func dataSourceCatalogFlattenFeatures(result []catalogmanagementv1.Feature) (features []map[string]interface{}) {
	for _, featuresItem := range result {
		features = append(features, dataSourceCatalogFeaturesToMap(featuresItem))
	}

	return features
}

func dataSourceCatalogFeaturesToMap(featuresItem catalogmanagementv1.Feature) (featuresMap map[string]interface{}) {
	featuresMap = map[string]interface{}{}

	if featuresItem.Title != nil {
		featuresMap["title"] = featuresItem.Title
	}
	if featuresItem.Description != nil {
		featuresMap["description"] = featuresItem.Description
	}

	return featuresMap
}

func dataSourceCatalogFlattenCatalogFilters(result catalogmanagementv1.Filters) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceCatalogCatalogFiltersToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceCatalogCatalogFiltersToMap(catalogFiltersItem catalogmanagementv1.Filters) (catalogFiltersMap map[string]interface{}) {
	catalogFiltersMap = map[string]interface{}{}

	if catalogFiltersItem.IncludeAll != nil {
		catalogFiltersMap["include_all"] = catalogFiltersItem.IncludeAll
	}
	if catalogFiltersItem.CategoryFilters != nil {
		convertedMap := make(map[string]interface{}, len(catalogFiltersItem.CategoryFilters))
		for k, v := range catalogFiltersItem.CategoryFilters {
			convertedMap[k] = v
		}
		catalogFiltersMap["category_filters"] = Flatten(convertedMap)
	}
	if catalogFiltersItem.IDFilters != nil {
		idFiltersList := []map[string]interface{}{}
		idFiltersMap := dataSourceCatalogCatalogFiltersIDFiltersToMap(*catalogFiltersItem.IDFilters)
		idFiltersList = append(idFiltersList, idFiltersMap)
		catalogFiltersMap["id_filters"] = idFiltersList
	}

	return catalogFiltersMap
}

func dataSourceCatalogCatalogFiltersIDFiltersToMap(idFiltersItem catalogmanagementv1.IDFilter) (idFiltersMap map[string]interface{}) {
	idFiltersMap = map[string]interface{}{}

	if idFiltersItem.Include != nil {
		includeList := []map[string]interface{}{}
		includeMap := dataSourceCatalogIDFiltersIncludeToMap(*idFiltersItem.Include)
		includeList = append(includeList, includeMap)
		idFiltersMap["include"] = includeList
	}
	if idFiltersItem.Exclude != nil {
		excludeList := []map[string]interface{}{}
		excludeMap := dataSourceCatalogIDFiltersExcludeToMap(*idFiltersItem.Exclude)
		excludeList = append(excludeList, excludeMap)
		idFiltersMap["exclude"] = excludeList
	}

	return idFiltersMap
}

func dataSourceCatalogIDFiltersIncludeToMap(includeItem catalogmanagementv1.FilterTerms) (includeMap map[string]interface{}) {
	includeMap = map[string]interface{}{}

	if includeItem.FilterTerms != nil {
		includeMap["filter_terms"] = includeItem.FilterTerms
	}

	return includeMap
}

func dataSourceCatalogIDFiltersExcludeToMap(excludeItem catalogmanagementv1.FilterTerms) (excludeMap map[string]interface{}) {
	excludeMap = map[string]interface{}{}

	if excludeItem.FilterTerms != nil {
		excludeMap["filter_terms"] = excludeItem.FilterTerms
	}

	return excludeMap
}

func dataSourceCatalogFlattenSyndicationSettings(result catalogmanagementv1.SyndicationResource) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceCatalogSyndicationSettingsToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceCatalogSyndicationSettingsToMap(syndicationSettingsItem catalogmanagementv1.SyndicationResource) (syndicationSettingsMap map[string]interface{}) {
	syndicationSettingsMap = map[string]interface{}{}

	if syndicationSettingsItem.RemoveRelatedComponents != nil {
		syndicationSettingsMap["remove_related_components"] = syndicationSettingsItem.RemoveRelatedComponents
	}
	if syndicationSettingsItem.Clusters != nil {
		clustersList := []map[string]interface{}{}
		for _, clustersItem := range syndicationSettingsItem.Clusters {
			clustersList = append(clustersList, dataSourceCatalogSyndicationSettingsClustersToMap(clustersItem))
		}
		syndicationSettingsMap["clusters"] = clustersList
	}
	if syndicationSettingsItem.History != nil {
		historyList := []map[string]interface{}{}
		historyMap := dataSourceCatalogSyndicationSettingsHistoryToMap(*syndicationSettingsItem.History)
		historyList = append(historyList, historyMap)
		syndicationSettingsMap["history"] = historyList
	}
	if syndicationSettingsItem.Authorization != nil {
		authorizationList := []map[string]interface{}{}
		authorizationMap := dataSourceCatalogSyndicationSettingsAuthorizationToMap(*syndicationSettingsItem.Authorization)
		authorizationList = append(authorizationList, authorizationMap)
		syndicationSettingsMap["authorization"] = authorizationList
	}

	return syndicationSettingsMap
}

func dataSourceCatalogSyndicationSettingsClustersToMap(clustersItem catalogmanagementv1.SyndicationCluster) (clustersMap map[string]interface{}) {
	clustersMap = map[string]interface{}{}

	if clustersItem.Region != nil {
		clustersMap["region"] = clustersItem.Region
	}
	if clustersItem.ID != nil {
		clustersMap["id"] = clustersItem.ID
	}
	if clustersItem.Name != nil {
		clustersMap["name"] = clustersItem.Name
	}
	if clustersItem.ResourceGroupName != nil {
		clustersMap["resource_group_name"] = clustersItem.ResourceGroupName
	}
	if clustersItem.Type != nil {
		clustersMap["type"] = clustersItem.Type
	}
	if clustersItem.Namespaces != nil {
		clustersMap["namespaces"] = clustersItem.Namespaces
	}
	if clustersItem.AllNamespaces != nil {
		clustersMap["all_namespaces"] = clustersItem.AllNamespaces
	}

	return clustersMap
}

func dataSourceCatalogSyndicationSettingsHistoryToMap(historyItem catalogmanagementv1.SyndicationHistory) (historyMap map[string]interface{}) {
	historyMap = map[string]interface{}{}

	if historyItem.Namespaces != nil {
		historyMap["namespaces"] = historyItem.Namespaces
	}
	if historyItem.Clusters != nil {
		clustersList := []map[string]interface{}{}
		for _, clustersItem := range historyItem.Clusters {
			clustersList = append(clustersList, dataSourceCatalogHistoryClustersToMap(clustersItem))
		}
		historyMap["clusters"] = clustersList
	}
	if historyItem.LastRun != nil {
		historyMap["last_run"] = historyItem.LastRun.String()
	}

	return historyMap
}

func dataSourceCatalogHistoryClustersToMap(clustersItem catalogmanagementv1.SyndicationCluster) (clustersMap map[string]interface{}) {
	clustersMap = map[string]interface{}{}

	if clustersItem.Region != nil {
		clustersMap["region"] = clustersItem.Region
	}
	if clustersItem.ID != nil {
		clustersMap["id"] = clustersItem.ID
	}
	if clustersItem.Name != nil {
		clustersMap["name"] = clustersItem.Name
	}
	if clustersItem.ResourceGroupName != nil {
		clustersMap["resource_group_name"] = clustersItem.ResourceGroupName
	}
	if clustersItem.Type != nil {
		clustersMap["type"] = clustersItem.Type
	}
	if clustersItem.Namespaces != nil {
		clustersMap["namespaces"] = clustersItem.Namespaces
	}
	if clustersItem.AllNamespaces != nil {
		clustersMap["all_namespaces"] = clustersItem.AllNamespaces
	}

	return clustersMap
}

func dataSourceCatalogSyndicationSettingsAuthorizationToMap(authorizationItem catalogmanagementv1.SyndicationAuthorization) (authorizationMap map[string]interface{}) {
	authorizationMap = map[string]interface{}{}

	if authorizationItem.Token != nil {
		authorizationMap["token"] = authorizationItem.Token
	}
	if authorizationItem.LastRun != nil {
		authorizationMap["last_run"] = authorizationItem.LastRun.String()
	}

	return authorizationMap
}
