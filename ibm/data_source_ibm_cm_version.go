// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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
				ForceNew:    true,
				Description: "Catalog identifier.",
			},
			"catalog_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Catalog identifier.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Offering identification.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tags array.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"target_kinds": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"content": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "byte array representing the content to be imported.  Only supported for OVA images at this time.",
			},
			"zipurl": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL path to zip location.  If not specified, must provide content in the body of this call.",
			},
			"target_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The semver value for this new version, if not found in the zip url package content.",
			},
			"repo_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
			"single_instance": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Denotes if single instance can be deployed to a given cluster.",
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

	d.SetId(*version.VersionLocator)
	if err = d.Set("crn", version.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("version", version.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}
	if err = d.Set("sha", version.Sha); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sha: %s", err))
	}
	if err = d.Set("catalog_id", version.CatalogID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_id: %s", err))
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

	return nil
}
