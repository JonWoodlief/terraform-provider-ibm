provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cm_catalog resource instance
resource "ibm_cm_catalog" "cm_catalog_instance" {
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

// Provision cm_offering resource instance
resource "ibm_cm_offering" "cm_offering_instance" {
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

// Provision cm_version resource instance
resource "ibm_cm_version" "cm_version_instance" {
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

// Provision cm_offering_instance resource instance
resource "ibm_cm_offering_instance" "cm_offering_instance_instance" {
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

// Create cm_catalog data source
data "ibm_cm_catalog" "cm_catalog_instance" {
  catalog_identifier = var.cm_catalog_catalog_identifier
}

// Create cm_offering data source
data "ibm_cm_offering" "cm_offering_instance" {
  catalog_identifier = var.cm_offering_catalog_identifier
  offering_id = var.cm_offering_offering_id
}

// Create cm_version data source
data "ibm_cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}

// Create cm_offering_instance data source
data "ibm_cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
