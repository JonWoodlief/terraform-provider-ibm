variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cm_catalog
variable "cm_catalog_rev" {
  description = "Cloudant revision."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_label" {
  description = "Display Name in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_short_description" {
  description = "Description in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_catalog_icon_url" {
  description = "URL for an icon associated with this catalog."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_tags" {
  description = "List of tags associated with this catalog."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_catalog_features" {
  description = "List of features associated with this catalog."
  type        = list(object({ example=string }))
  default     = [ [ { example: "object" } ] ]
}
variable "cm_catalog_disabled" {
  description = "Denotes whether a catalog is disabled."
  type        = bool
  default     = false
}
variable "cm_catalog_resource_group_id" {
  description = "Resource group id the catalog is owned by."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_owning_account" {
  description = "Account that owns catalog."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_catalog_filters" {
  description = "Filters for account and catalog filters."
  type        = object({ example=string })
  default     = [ { example: "object" } ]
}
variable "cm_catalog_syndication_settings" {
  description = "Feature information."
  type        = object({ example=string })
  default     = [ { example: "object" } ]
}

// Resource arguments for cm_offering
variable "cm_offering_catalog_identifier" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_identifier"
}
variable "cm_offering_rev" {
  description = "Cloudant revision."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_url" {
  description = "The url for this specific offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_crn" {
  description = "The crn for this specific offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_label" {
  description = "Display Name in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_name" {
  description = "The programmatic name of this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_offering_icon_url" {
  description = "URL for an icon associated with this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_offering_docs_url" {
  description = "URL for an additional docs with this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_offering_support_url" {
  description = "URL to be displayed in the Consumption UI for getting support on this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_tags" {
  description = "List of tags associated with this catalog."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_offering_rating" {
  description = "Repository info for offerings."
  type        = object({ example=string })
  default     = [ { example: "object" } ]
}
variable "cm_offering_created" {
  description = "The date and time this catalog was created."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "cm_offering_updated" {
  description = "The date and time this catalog was last updated."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "cm_offering_short_description" {
  description = "Short description in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_long_description" {
  description = "Long description in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_features" {
  description = "list of features associated with this offering."
  type        = list(object({ example=string }))
  default     = [ [ { example: "object" } ] ]
}
variable "cm_offering_kinds" {
  description = "Array of kind."
  type        = list(object({ example=string }))
  default     = [ [ { example: "object" } ] ]
}
variable "cm_offering_permit_request_ibm_public_publish" {
  description = "Is it permitted to request publishing to IBM or Public."
  type        = bool
  default     = false
}
variable "cm_offering_ibm_publish_approved" {
  description = "Indicates if this offering has been approved for use by all IBMers."
  type        = bool
  default     = false
}
variable "cm_offering_public_publish_approved" {
  description = "Indicates if this offering has been approved for use by all IBM Cloud users."
  type        = bool
  default     = false
}
variable "cm_offering_public_original_crn" {
  description = "The original offering CRN that this publish entry came from."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_publish_public_crn" {
  description = "The crn of the public catalog entry of this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_portal_approval_record" {
  description = "The portal's approval record ID."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_portal_ui_url" {
  description = "The portal UI URL."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_catalog_id" {
  description = "The id of the catalog containing this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_catalog_name" {
  description = "The name of the catalog."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_metadata" {
  description = "Map of metadata values for this offering."
  type        = map()
  default     = { "example":  }
}
variable "cm_offering_disclaimer" {
  description = "A disclaimer for this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_hidden" {
  description = "Determine if this offering should be displayed in the Consumption UI."
  type        = bool
  default     = false
}
variable "cm_offering_provider" {
  description = "Provider of this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_repo_info" {
  description = "Repository info for offerings."
  type        = object({ example=string })
  default     = [ { example: "object" } ]
}

// Resource arguments for cm_version
variable "cm_version_catalog_identifier" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_identifier"
}
variable "cm_version_offering_id" {
  description = "Offering identification."
  type        = string
  default     = "offering_id"
}
variable "cm_version_tags" {
  description = "Tags array."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_version_target_kinds" {
  description = "Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_version_content" {
  description = "byte array representing the content to be imported.  Only supported for OVA images at this time."
  type        = 
  default     = 
}
variable "cm_version_zipurl" {
  description = "URL path to zip location.  If not specified, must provide content in the body of this call."
  type        = string
  default     = "placeholder"
}
variable "cm_version_target_version" {
  description = "The semver value for this new version, if not found in the zip url package content."
  type        = string
  default     = "placeholder"
}
variable "cm_version_include_config" {
  description = "Add all possible configuration values to this version when importing."
  type        = bool
  default     = false
}
variable "cm_version_repo_type" {
  description = "The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'."
  type        = string
  default     = "placeholder"
}

// Resource arguments for cm_offering_instance
variable "cm_offering_instance_x_auth_refresh_token" {
  description = "IAM Refresh token."
  type        = string
  default     = "x_auth_refresh_token"
}
variable "cm_offering_instance_url" {
  description = "url reference to this object."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_crn" {
  description = "platform CRN for this instance."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_label" {
  description = "the label for this instance."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_catalog_id" {
  description = "Catalog ID this instance was created from."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_offering_id" {
  description = "Offering ID this instance was created from."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_kind_format" {
  description = "the format this instance has (helm, operator, ova...)."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_version" {
  description = "The version this instance was installed from (not version id)."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_cluster_id" {
  description = "Cluster ID."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_cluster_region" {
  description = "Cluster region (e.g., us-south)."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_cluster_namespaces" {
  description = "List of target namespaces to install into."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_offering_instance_cluster_all_namespaces" {
  description = "designate to install into all namespaces."
  type        = bool
  default     = false
}

// Data source arguments for cm_catalog
variable "cm_catalog_catalog_identifier" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_identifier"
}

// Data source arguments for cm_offering
variable "cm_offering_catalog_identifier" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_identifier"
}
variable "cm_offering_offering_id" {
  description = "Offering identification."
  type        = string
  default     = "offering_id"
}

// Data source arguments for cm_version
variable "cm_version_version_loc_id" {
  description = "A dotted value of `catalogID`.`versionID`."
  type        = string
  default     = "version_loc_id"
}

// Data source arguments for cm_offering_instance
variable "cm_offering_instance_instance_identifier" {
  description = "Version Instance identifier."
  type        = string
  default     = "instance_identifier"
}
