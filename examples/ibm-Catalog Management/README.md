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
  tags = var.cm_offering_tags
  target_kinds = var.cm_offering_target_kinds
  content = var.cm_offering_content
  zipurl = var.cm_offering_zipurl
  offering_id = var.cm_offering_offering_id
  target_version = var.cm_offering_target_version
  include_config = var.cm_offering_include_config
  repo_type = var.cm_offering_repo_type
  x_auth_token = var.cm_offering_x_auth_token
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
| tags | Tags array. | `list(string)` | false |
| target_kinds | Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'. | `list(string)` | false |
| content | byte array representing the content to be imported.  Only supported for OVA images at this time. | `` | false |
| zipurl | URL path to zip location.  If not specified, must provide content in this post body. | `string` | false |
| offering_id | Re-use the specified offeringID during import. | `string` | false |
| target_version | The semver value for this new version. | `string` | false |
| include_config | Add all possible configuration items when creating this version. | `bool` | false |
| repo_type | The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'. | `string` | false |
| x_auth_token | Authentication token used to access the specified zip file. | `string` | false |
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
