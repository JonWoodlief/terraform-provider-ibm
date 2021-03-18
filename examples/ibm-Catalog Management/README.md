# Example for CatalogManagementV1

This example illustrates how to use the CatalogManagementV1

These types of resources are supported:

* cm_catalog
* cm_offering
* cm_version
* cm_offering_instance

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CatalogManagementV1 resources

cm_catalog resource:

```hcl
resource "cm_catalog" "cm_catalog_instance" {
  rev = var.cm_catalog_rev
  label = var.cm_catalog_label
  short_description = var.cm_catalog_short_description
  catalog_icon_url = var.cm_catalog_catalog_icon_url
  tags = var.cm_catalog_tags
  features = var.cm_catalog_features
  disabled = var.cm_catalog_disabled
  resource_group_id = var.cm_catalog_resource_group_id
  owning_account = var.cm_catalog_owning_account
  catalog_filters = var.cm_catalog_catalog_filters
  syndication_settings = var.cm_catalog_syndication_settings
}
```
cm_offering resource:

```hcl
resource "cm_offering" "cm_offering_instance" {
  catalog_identifier = var.cm_offering_catalog_identifier
  rev = var.cm_offering_rev
  url = var.cm_offering_url
  crn = var.cm_offering_crn
  label = var.cm_offering_label
  name = var.cm_offering_name
  offering_icon_url = var.cm_offering_offering_icon_url
  offering_docs_url = var.cm_offering_offering_docs_url
  offering_support_url = var.cm_offering_offering_support_url
  tags = var.cm_offering_tags
  rating = var.cm_offering_rating
  created = var.cm_offering_created
  updated = var.cm_offering_updated
  short_description = var.cm_offering_short_description
  long_description = var.cm_offering_long_description
  features = var.cm_offering_features
  kinds = var.cm_offering_kinds
  permit_request_ibm_public_publish = var.cm_offering_permit_request_ibm_public_publish
  ibm_publish_approved = var.cm_offering_ibm_publish_approved
  public_publish_approved = var.cm_offering_public_publish_approved
  public_original_crn = var.cm_offering_public_original_crn
  publish_public_crn = var.cm_offering_publish_public_crn
  portal_approval_record = var.cm_offering_portal_approval_record
  portal_ui_url = var.cm_offering_portal_ui_url
  catalog_id = var.cm_offering_catalog_id
  catalog_name = var.cm_offering_catalog_name
  metadata = var.cm_offering_metadata
  disclaimer = var.cm_offering_disclaimer
  hidden = var.cm_offering_hidden
  provider = var.cm_offering_provider
  repo_info = var.cm_offering_repo_info
}
```
cm_version resource:

```hcl
resource "cm_version" "cm_version_instance" {
  catalog_identifier = var.cm_version_catalog_identifier
  offering_id = var.cm_version_offering_id
  tags = var.cm_version_tags
  target_kinds = var.cm_version_target_kinds
  content = var.cm_version_content
  zipurl = var.cm_version_zipurl
  target_version = var.cm_version_target_version
  include_config = var.cm_version_include_config
  repo_type = var.cm_version_repo_type
}
```
cm_offering_instance resource:

```hcl
resource "cm_offering_instance" "cm_offering_instance_instance" {
  x_auth_refresh_token = var.cm_offering_instance_x_auth_refresh_token
  url = var.cm_offering_instance_url
  crn = var.cm_offering_instance_crn
  label = var.cm_offering_instance_label
  catalog_id = var.cm_offering_instance_catalog_id
  offering_id = var.cm_offering_instance_offering_id
  kind_format = var.cm_offering_instance_kind_format
  version = var.cm_offering_instance_version
  cluster_id = var.cm_offering_instance_cluster_id
  cluster_region = var.cm_offering_instance_cluster_region
  cluster_namespaces = var.cm_offering_instance_cluster_namespaces
  cluster_all_namespaces = var.cm_offering_instance_cluster_all_namespaces
}
```

## CatalogManagementV1 Data sources

cm_catalog data source:

```hcl
data "cm_catalog" "cm_catalog_instance" {
  catalog_identifier = var.cm_catalog_catalog_identifier
}
```
cm_offering data source:

```hcl
data "cm_offering" "cm_offering_instance" {
  catalog_identifier = var.cm_offering_catalog_identifier
  offering_id = var.cm_offering_offering_id
}
```
cm_version data source:

```hcl
data "cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}
```
cm_offering_instance data source:

```hcl
data "cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| rev | Cloudant revision. | `string` | false |
| label | Display Name in the requested language. | `string` | false |
| short_description | Description in the requested language. | `string` | false |
| catalog_icon_url | URL for an icon associated with this catalog. | `string` | false |
| tags | List of tags associated with this catalog. | `list(string)` | false |
| features | List of features associated with this catalog. | `list()` | false |
| disabled | Denotes whether a catalog is disabled. | `bool` | false |
| resource_group_id | Resource group id the catalog is owned by. | `string` | false |
| owning_account | Account that owns catalog. | `string` | false |
| catalog_filters | Filters for account and catalog filters. | `` | false |
| syndication_settings | Feature information. | `` | false |
| catalog_identifier | Catalog identifier. | `string` | true |
| rev | Cloudant revision. | `string` | false |
| url | The url for this specific offering. | `string` | false |
| crn | The crn for this specific offering. | `string` | false |
| label | Display Name in the requested language. | `string` | false |
| name | The programmatic name of this offering. | `string` | false |
| offering_icon_url | URL for an icon associated with this offering. | `string` | false |
| offering_docs_url | URL for an additional docs with this offering. | `string` | false |
| offering_support_url | URL to be displayed in the Consumption UI for getting support on this offering. | `string` | false |
| tags | List of tags associated with this catalog. | `list(string)` | false |
| rating | Repository info for offerings. | `` | false |
| created | The date and time this catalog was created. | `` | false |
| updated | The date and time this catalog was last updated. | `` | false |
| short_description | Short description in the requested language. | `string` | false |
| long_description | Long description in the requested language. | `string` | false |
| features | list of features associated with this offering. | `list()` | false |
| kinds | Array of kind. | `list()` | false |
| permit_request_ibm_public_publish | Is it permitted to request publishing to IBM or Public. | `bool` | false |
| ibm_publish_approved | Indicates if this offering has been approved for use by all IBMers. | `bool` | false |
| public_publish_approved | Indicates if this offering has been approved for use by all IBM Cloud users. | `bool` | false |
| public_original_crn | The original offering CRN that this publish entry came from. | `string` | false |
| publish_public_crn | The crn of the public catalog entry of this offering. | `string` | false |
| portal_approval_record | The portal's approval record ID. | `string` | false |
| portal_ui_url | The portal UI URL. | `string` | false |
| catalog_id | The id of the catalog containing this offering. | `string` | false |
| catalog_name | The name of the catalog. | `string` | false |
| metadata | Map of metadata values for this offering. | `map()` | false |
| disclaimer | A disclaimer for this offering. | `string` | false |
| hidden | Determine if this offering should be displayed in the Consumption UI. | `bool` | false |
| provider | Provider of this offering. | `string` | false |
| repo_info | Repository info for offerings. | `` | false |
| catalog_identifier | Catalog identifier. | `string` | true |
| offering_id | Offering identification. | `string` | true |
| tags | Tags array. | `list(string)` | false |
| target_kinds | Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'. | `list(string)` | false |
| content | byte array representing the content to be imported.  Only supported for OVA images at this time. | `` | false |
| zipurl | URL path to zip location.  If not specified, must provide content in the body of this call. | `string` | false |
| target_version | The semver value for this new version, if not found in the zip url package content. | `string` | false |
| include_config | Add all possible configuration values to this version when importing. | `bool` | false |
| repo_type | The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'. | `string` | false |
| x_auth_refresh_token | IAM Refresh token. | `string` | true |
| url | url reference to this object. | `string` | false |
| crn | platform CRN for this instance. | `string` | false |
| label | the label for this instance. | `string` | false |
| catalog_id | Catalog ID this instance was created from. | `string` | false |
| offering_id | Offering ID this instance was created from. | `string` | false |
| kind_format | the format this instance has (helm, operator, ova...). | `string` | false |
| version | The version this instance was installed from (not version id). | `string` | false |
| cluster_id | Cluster ID. | `string` | false |
| cluster_region | Cluster region (e.g., us-south). | `string` | false |
| cluster_namespaces | List of target namespaces to install into. | `list(string)` | false |
| cluster_all_namespaces | designate to install into all namespaces. | `bool` | false |
| catalog_identifier | Catalog identifier. | `string` | true |
| catalog_identifier | Catalog identifier. | `string` | true |
| offering_id | Offering identification. | `string` | true |
| version_loc_id | A dotted value of `catalogID`.`versionID`. | `string` | true |
| instance_identifier | Version Instance identifier. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| cm_catalog | cm_catalog object |
| cm_offering | cm_offering object |
| cm_version | cm_version object |
| cm_offering_instance | cm_offering_instance object |
| cm_catalog | cm_catalog object |
| cm_offering | cm_offering object |
| cm_version | cm_version object |
| cm_offering_instance | cm_offering_instance object |
