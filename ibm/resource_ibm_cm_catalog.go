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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func resourceIBMCmCatalog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCmCatalogCreate,
		ReadContext:   resourceIBMCmCatalogRead,
		UpdateContext: resourceIBMCmCatalogUpdate,
		DeleteContext: resourceIBMCmCatalogDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloudant revision.",
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Display Name in the requested language.",
			},
			"short_description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description in the requested language.",
			},
			"catalog_icon_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL for an icon associated with this catalog.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of tags associated with this catalog.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"features": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
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
				Description: "Denotes whether a catalog is disabled.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource group id the catalog is owned by.",
			},
			"owning_account": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account that owns catalog.",
			},
			"catalog_filters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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

func resourceIBMCmCatalogCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}
	fmt.Printf("client is a nil pointer: %v\n", catalogManagementClient == nil)
	log.Printf("client is a nil pointer: %v\n", catalogManagementClient == nil)

	createCatalogOptions := &catalogmanagementv1.CreateCatalogOptions{}

	if _, ok := d.GetOk("rev"); ok {
		createCatalogOptions.SetRev(d.Get("rev").(string))
	}
	if _, ok := d.GetOk("label"); ok {
		createCatalogOptions.SetLabel(d.Get("label").(string))
	}
	if _, ok := d.GetOk("short_description"); ok {
		createCatalogOptions.SetShortDescription(d.Get("short_description").(string))
	}
	if _, ok := d.GetOk("catalog_icon_url"); ok {
		createCatalogOptions.SetCatalogIconURL(d.Get("catalog_icon_url").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		createCatalogOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("features"); ok {
		var features []catalogmanagementv1.Feature
		for _, e := range d.Get("features").([]interface{}) {
			value := e.(map[string]interface{})
			featuresItem := resourceIBMCmCatalogMapToFeature(value)
			features = append(features, featuresItem)
		}
		createCatalogOptions.SetFeatures(features)
	}
	if _, ok := d.GetOk("disabled"); ok {
		createCatalogOptions.SetDisabled(d.Get("disabled").(bool))
	}
	if _, ok := d.GetOk("resource_group_id"); ok {
		createCatalogOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	}
	if _, ok := d.GetOk("owning_account"); ok {
		createCatalogOptions.SetOwningAccount(d.Get("owning_account").(string))
	}
	if _, ok := d.GetOk("catalog_filters"); ok {
		catalogFilters := resourceIBMCmCatalogMapToFilters(d.Get("catalog_filters.0").(map[string]interface{}))
		createCatalogOptions.SetCatalogFilters(&catalogFilters)
	}
	if _, ok := d.GetOk("syndication_settings"); ok {
		syndicationSettings := resourceIBMCmCatalogMapToSyndicationResource(d.Get("syndication_settings.0").(map[string]interface{}))
		createCatalogOptions.SetSyndicationSettings(&syndicationSettings)
	}

	catalog, response, err := catalogManagementClient.CreateCatalogWithContext(context, createCatalogOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateCatalogWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*catalog.ID)

	return resourceIBMCmCatalogRead(context, d, meta)
}

func resourceIBMCmCatalogMapToFeature(featureMap map[string]interface{}) catalogmanagementv1.Feature {
	feature := catalogmanagementv1.Feature{}

	if featureMap["title"] != nil {
		feature.Title = core.StringPtr(featureMap["title"].(string))
	}
	if featureMap["description"] != nil {
		feature.Description = core.StringPtr(featureMap["description"].(string))
	}

	return feature
}

func resourceIBMCmCatalogMapToFilters(filtersMap map[string]interface{}) catalogmanagementv1.Filters {
	filters := catalogmanagementv1.Filters{}

	if filtersMap["include_all"] != nil {
		filters.IncludeAll = core.BoolPtr(filtersMap["include_all"].(bool))
	}
	if filtersMap["category_filters"] != nil {
		// TODO: handle CategoryFilters of type CategoryFilter -- not primitive type, not list
	}
	if filtersMap["id_filters"] != nil {
		// TODO: handle IDFilters of type IDFilter -- not primitive type, not list
	}

	return filters
}

func resourceIBMCmCatalogMapToCategoryFilter(categoryFilterMap map[string]interface{}) catalogmanagementv1.CategoryFilter {
	categoryFilter := catalogmanagementv1.CategoryFilter{}

	if categoryFilterMap["include"] != nil {
		categoryFilter.Include = core.BoolPtr(categoryFilterMap["include"].(bool))
	}
	if categoryFilterMap["filter"] != nil {
		// TODO: handle Filter of type FilterTerms -- not primitive type, not list
	}

	return categoryFilter
}

/* func resourceIBMCmCatalogMapToFilterTerms(filterTermsMap map[string]interface{}) catalogmanagementv1.FilterTerms {
	filterTerms := catalogmanagementv1.FilterTerms{}

	if filterTermsMap["filter_terms"] != nil {
		filterTerms := []string{}
		for _, filterTermsItem := range filterTermsMap["filter_terms"].([]interface{}) {
			filterTerms = append(filterTerms, filterTermsItem.(string))
		}
		filterTerms.FilterTerms = filterTerms
	}

	return filterTerms
} */

func resourceIBMCmCatalogMapToIDFilter(idFilterMap map[string]interface{}) catalogmanagementv1.IDFilter {
	idFilter := catalogmanagementv1.IDFilter{}

	if idFilterMap["include"] != nil {
		// TODO: handle Include of type FilterTerms -- not primitive type, not list
	}
	if idFilterMap["exclude"] != nil {
		// TODO: handle Exclude of type FilterTerms -- not primitive type, not list
	}

	return idFilter
}

func resourceIBMCmCatalogMapToSyndicationResource(syndicationResourceMap map[string]interface{}) catalogmanagementv1.SyndicationResource {
	syndicationResource := catalogmanagementv1.SyndicationResource{}

	if syndicationResourceMap["remove_related_components"] != nil {
		syndicationResource.RemoveRelatedComponents = core.BoolPtr(syndicationResourceMap["remove_related_components"].(bool))
	}
	if syndicationResourceMap["clusters"] != nil {
		clusters := []catalogmanagementv1.SyndicationCluster{}
		for _, clustersItem := range syndicationResourceMap["clusters"].([]interface{}) {
			clustersItemModel := resourceIBMCmCatalogMapToSyndicationCluster(clustersItem.(map[string]interface{}))
			clusters = append(clusters, clustersItemModel)
		}
		syndicationResource.Clusters = clusters
	}
	if syndicationResourceMap["history"] != nil {
		// TODO: handle History of type SyndicationHistory -- not primitive type, not list
	}
	if syndicationResourceMap["authorization"] != nil {
		// TODO: handle Authorization of type SyndicationAuthorization -- not primitive type, not list
	}

	return syndicationResource
}

func resourceIBMCmCatalogMapToSyndicationCluster(syndicationClusterMap map[string]interface{}) catalogmanagementv1.SyndicationCluster {
	syndicationCluster := catalogmanagementv1.SyndicationCluster{}

	if syndicationClusterMap["region"] != nil {
		syndicationCluster.Region = core.StringPtr(syndicationClusterMap["region"].(string))
	}
	if syndicationClusterMap["id"] != nil {
		syndicationCluster.ID = core.StringPtr(syndicationClusterMap["id"].(string))
	}
	if syndicationClusterMap["name"] != nil {
		syndicationCluster.Name = core.StringPtr(syndicationClusterMap["name"].(string))
	}
	if syndicationClusterMap["resource_group_name"] != nil {
		syndicationCluster.ResourceGroupName = core.StringPtr(syndicationClusterMap["resource_group_name"].(string))
	}
	if syndicationClusterMap["type"] != nil {
		syndicationCluster.Type = core.StringPtr(syndicationClusterMap["type"].(string))
	}
	if syndicationClusterMap["namespaces"] != nil {
		namespaces := []string{}
		for _, namespacesItem := range syndicationClusterMap["namespaces"].([]interface{}) {
			namespaces = append(namespaces, namespacesItem.(string))
		}
		syndicationCluster.Namespaces = namespaces
	}
	if syndicationClusterMap["all_namespaces"] != nil {
		syndicationCluster.AllNamespaces = core.BoolPtr(syndicationClusterMap["all_namespaces"].(bool))
	}

	return syndicationCluster
}

func resourceIBMCmCatalogMapToSyndicationHistory(syndicationHistoryMap map[string]interface{}) catalogmanagementv1.SyndicationHistory {
	syndicationHistory := catalogmanagementv1.SyndicationHistory{}

	if syndicationHistoryMap["namespaces"] != nil {
		namespaces := []string{}
		for _, namespacesItem := range syndicationHistoryMap["namespaces"].([]interface{}) {
			namespaces = append(namespaces, namespacesItem.(string))
		}
		syndicationHistory.Namespaces = namespaces
	}
	if syndicationHistoryMap["clusters"] != nil {
		clusters := []catalogmanagementv1.SyndicationCluster{}
		for _, clustersItem := range syndicationHistoryMap["clusters"].([]interface{}) {
			clustersItemModel := resourceIBMCmCatalogMapToSyndicationCluster(clustersItem.(map[string]interface{}))
			clusters = append(clusters, clustersItemModel)
		}
		syndicationHistory.Clusters = clusters
	}
	if syndicationHistoryMap["last_run"] != nil {

	}

	return syndicationHistory
}

func resourceIBMCmCatalogMapToSyndicationAuthorization(syndicationAuthorizationMap map[string]interface{}) catalogmanagementv1.SyndicationAuthorization {
	syndicationAuthorization := catalogmanagementv1.SyndicationAuthorization{}

	if syndicationAuthorizationMap["token"] != nil {
		syndicationAuthorization.Token = core.StringPtr(syndicationAuthorizationMap["token"].(string))
	}
	if syndicationAuthorizationMap["last_run"] != nil {

	}

	return syndicationAuthorization
}

func resourceIBMCmCatalogRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

	getCatalogOptions.SetCatalogIdentifier(d.Id())

	catalog, response, err := catalogManagementClient.GetCatalogWithContext(context, getCatalogOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetCatalogWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
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
	if catalog.Tags != nil {
		if err = d.Set("tags", catalog.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	if catalog.Features != nil {
		features := []map[string]interface{}{}
		for _, featuresItem := range catalog.Features {
			featuresItemMap := resourceIBMCmCatalogFeatureToMap(featuresItem)
			features = append(features, featuresItemMap)
		}
		if err = d.Set("features", features); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting features: %s", err))
		}
	}
	if err = d.Set("disabled", catalog.Disabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disabled: %s", err))
	}
	if err = d.Set("resource_group_id", catalog.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("owning_account", catalog.OwningAccount); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting owning_account: %s", err))
	}
	if catalog.CatalogFilters != nil {
		catalogFiltersMap := resourceIBMCmCatalogFiltersToMap(*catalog.CatalogFilters)
		if err = d.Set("catalog_filters", []map[string]interface{}{catalogFiltersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting catalog_filters: %s", err))
		}
	}
	if catalog.SyndicationSettings != nil {
		syndicationSettingsMap := resourceIBMCmCatalogSyndicationResourceToMap(*catalog.SyndicationSettings)
		if err = d.Set("syndication_settings", []map[string]interface{}{syndicationSettingsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting syndication_settings: %s", err))
		}
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
	if err = d.Set("created", catalog.Created.String()); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if err = d.Set("updated", catalog.Updated.String()); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}

	return nil
}

func resourceIBMCmCatalogFeatureToMap(feature catalogmanagementv1.Feature) map[string]interface{} {
	featureMap := map[string]interface{}{}

	featureMap["title"] = feature.Title
	featureMap["description"] = feature.Description

	return featureMap
}

func resourceIBMCmCatalogFiltersToMap(filters catalogmanagementv1.Filters) map[string]interface{} {
	filtersMap := map[string]interface{}{}

	filtersMap["include_all"] = filters.IncludeAll
	if filters.CategoryFilters != nil {
		// TODO: handle CategoryFilters of type TypeMap -- container, not list
	}
	if filters.IDFilters != nil {
		IDFiltersMap := resourceIBMCmCatalogIDFilterToMap(*filters.IDFilters)
		filtersMap["id_filters"] = []map[string]interface{}{IDFiltersMap}
	}

	return filtersMap
}

func resourceIBMCmCatalogCategoryFilterToMap(categoryFilter catalogmanagementv1.CategoryFilter) map[string]interface{} {
	categoryFilterMap := map[string]interface{}{}

	categoryFilterMap["include"] = categoryFilter.Include
	if categoryFilter.Filter != nil {
		FilterMap := resourceIBMCmCatalogFilterTermsToMap(*categoryFilter.Filter)
		categoryFilterMap["filter"] = []map[string]interface{}{FilterMap}
	}

	return categoryFilterMap
}

func resourceIBMCmCatalogFilterTermsToMap(filterTerms catalogmanagementv1.FilterTerms) map[string]interface{} {
	filterTermsMap := map[string]interface{}{}

	if filterTerms.FilterTerms != nil {
		filterTermsMap["filter_terms"] = filterTerms.FilterTerms
	}

	return filterTermsMap
}

func resourceIBMCmCatalogIDFilterToMap(idFilter catalogmanagementv1.IDFilter) map[string]interface{} {
	idFilterMap := map[string]interface{}{}

	if idFilter.Include != nil {
		IncludeMap := resourceIBMCmCatalogFilterTermsToMap(*idFilter.Include)
		idFilterMap["include"] = []map[string]interface{}{IncludeMap}
	}
	if idFilter.Exclude != nil {
		ExcludeMap := resourceIBMCmCatalogFilterTermsToMap(*idFilter.Exclude)
		idFilterMap["exclude"] = []map[string]interface{}{ExcludeMap}
	}

	return idFilterMap
}

func resourceIBMCmCatalogSyndicationResourceToMap(syndicationResource catalogmanagementv1.SyndicationResource) map[string]interface{} {
	syndicationResourceMap := map[string]interface{}{}

	syndicationResourceMap["remove_related_components"] = syndicationResource.RemoveRelatedComponents
	if syndicationResource.Clusters != nil {
		clusters := []map[string]interface{}{}
		for _, clustersItem := range syndicationResource.Clusters {
			clustersItemMap := resourceIBMCmCatalogSyndicationClusterToMap(clustersItem)
			clusters = append(clusters, clustersItemMap)
			// TODO: handle Clusters of type TypeList -- list of non-primitive, not model items
		}
		syndicationResourceMap["clusters"] = clusters
	}
	if syndicationResource.History != nil {
		HistoryMap := resourceIBMCmCatalogSyndicationHistoryToMap(*syndicationResource.History)
		syndicationResourceMap["history"] = []map[string]interface{}{HistoryMap}
	}
	if syndicationResource.Authorization != nil {
		AuthorizationMap := resourceIBMCmCatalogSyndicationAuthorizationToMap(*syndicationResource.Authorization)
		syndicationResourceMap["authorization"] = []map[string]interface{}{AuthorizationMap}
	}

	return syndicationResourceMap
}

func resourceIBMCmCatalogSyndicationClusterToMap(syndicationCluster catalogmanagementv1.SyndicationCluster) map[string]interface{} {
	syndicationClusterMap := map[string]interface{}{}

	syndicationClusterMap["region"] = syndicationCluster.Region
	syndicationClusterMap["id"] = syndicationCluster.ID
	syndicationClusterMap["name"] = syndicationCluster.Name
	syndicationClusterMap["resource_group_name"] = syndicationCluster.ResourceGroupName
	syndicationClusterMap["type"] = syndicationCluster.Type
	if syndicationCluster.Namespaces != nil {
		syndicationClusterMap["namespaces"] = syndicationCluster.Namespaces
	}
	syndicationClusterMap["all_namespaces"] = syndicationCluster.AllNamespaces

	return syndicationClusterMap
}

func resourceIBMCmCatalogSyndicationHistoryToMap(syndicationHistory catalogmanagementv1.SyndicationHistory) map[string]interface{} {
	syndicationHistoryMap := map[string]interface{}{}

	if syndicationHistory.Namespaces != nil {
		syndicationHistoryMap["namespaces"] = syndicationHistory.Namespaces
	}
	if syndicationHistory.Clusters != nil {
		clusters := []map[string]interface{}{}
		for _, clustersItem := range syndicationHistory.Clusters {
			clustersItemMap := resourceIBMCmCatalogSyndicationClusterToMap(clustersItem)
			clusters = append(clusters, clustersItemMap)
			// TODO: handle Clusters of type TypeList -- list of non-primitive, not model items
		}
		syndicationHistoryMap["clusters"] = clusters
	}
	syndicationHistoryMap["last_run"] = syndicationHistory.LastRun.String()

	return syndicationHistoryMap
}

func resourceIBMCmCatalogSyndicationAuthorizationToMap(syndicationAuthorization catalogmanagementv1.SyndicationAuthorization) map[string]interface{} {
	syndicationAuthorizationMap := map[string]interface{}{}

	syndicationAuthorizationMap["token"] = syndicationAuthorization.Token
	syndicationAuthorizationMap["last_run"] = syndicationAuthorization.LastRun.String()

	return syndicationAuthorizationMap
}

func resourceIBMCmCatalogUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceCatalogOptions := &catalogmanagementv1.ReplaceCatalogOptions{}

	replaceCatalogOptions.SetCatalogIdentifier(d.Id())
	if _, ok := d.GetOk("rev"); ok {
		replaceCatalogOptions.SetRev(d.Get("rev").(string))
	}
	if _, ok := d.GetOk("label"); ok {
		replaceCatalogOptions.SetLabel(d.Get("label").(string))
	}
	if _, ok := d.GetOk("short_description"); ok {
		replaceCatalogOptions.SetShortDescription(d.Get("short_description").(string))
	}
	if _, ok := d.GetOk("catalog_icon_url"); ok {
		replaceCatalogOptions.SetCatalogIconURL(d.Get("catalog_icon_url").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		replaceCatalogOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("features"); ok {
		var features []catalogmanagementv1.Feature
		for _, e := range d.Get("features").([]interface{}) {
			value := e.(map[string]interface{})
			featuresItem := resourceIBMCmCatalogMapToFeature(value)
			features = append(features, featuresItem)
		}
		replaceCatalogOptions.SetFeatures(features)
	}
	if _, ok := d.GetOk("disabled"); ok {
		replaceCatalogOptions.SetDisabled(d.Get("disabled").(bool))
	}
	if _, ok := d.GetOk("resource_group_id"); ok {
		replaceCatalogOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	}
	if _, ok := d.GetOk("owning_account"); ok {
		replaceCatalogOptions.SetOwningAccount(d.Get("owning_account").(string))
	}
	if _, ok := d.GetOk("catalog_filters"); ok {
		catalogFilters := resourceIBMCmCatalogMapToFilters(d.Get("catalog_filters.0").(map[string]interface{}))
		replaceCatalogOptions.SetCatalogFilters(&catalogFilters)
	}
	if _, ok := d.GetOk("syndication_settings"); ok {
		syndicationSettings := resourceIBMCmCatalogMapToSyndicationResource(d.Get("syndication_settings.0").(map[string]interface{}))
		replaceCatalogOptions.SetSyndicationSettings(&syndicationSettings)
	}

	_, response, err := catalogManagementClient.ReplaceCatalogWithContext(context, replaceCatalogOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceCatalogWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	return resourceIBMCmCatalogRead(context, d, meta)
}

func resourceIBMCmCatalogDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteCatalogOptions := &catalogmanagementv1.DeleteCatalogOptions{}

	deleteCatalogOptions.SetCatalogIdentifier(d.Id())

	response, err := catalogManagementClient.DeleteCatalogWithContext(context, deleteCatalogOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteCatalogWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
