---
layout: "ibm"
page_title: "IBM : cm_offering"
sidebar_current: "docs-ibm-datasource-cm-offering"
description: |-
  Get information about cm_offering
---

# ibm\_cm_offering

Provides a read-only data source for cm_offering. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_offering" "cm_offering" {
	catalog_identifier = "catalog_identifier"
	offering_id = "offering_id"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, string) Catalog identifier.
* `offering_id` - (Required, string) Offering identification.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_offering.
* `id` - unique id.

* `rev` - Cloudant revision.

* `url` - The url for this specific offering.

* `crn` - The crn for this specific offering.

* `label` - Display Name in the requested language.

* `name` - The programmatic name of this offering.

* `offering_icon_url` - URL for an icon associated with this offering.

* `offering_docs_url` - URL for an additional docs with this offering.

* `offering_support_url` - URL to be displayed in the Consumption UI for getting support on this offering.

* `tags` - List of tags associated with this catalog.

* `rating` - Repository info for offerings. Nested `rating` blocks have the following structure:
	* `one_star_count` - One start rating.
	* `two_star_count` - Two start rating.
	* `three_star_count` - Three start rating.
	* `four_star_count` - Four start rating.

* `created` - The date and time this catalog was created.

* `updated` - The date and time this catalog was last updated.

* `short_description` - Short description in the requested language.

* `long_description` - Long description in the requested language.

* `features` - list of features associated with this offering. Nested `features` blocks have the following structure:
	* `title` - Heading.
	* `description` - Feature description.

* `kinds` - Array of kind. Nested `kinds` blocks have the following structure:
	* `id` - Unique ID.
	* `format_kind` - content kind, e.g., helm, vm image.
	* `target_kind` - target cloud to install, e.g., iks, open_shift_iks.
	* `metadata` - Open ended metadata information.
	* `install_description` - Installation instruction.
	* `tags` - List of tags associated with this catalog.
	* `additional_features` - List of features associated with this offering. Nested `additional_features` blocks have the following structure:
		* `title` - Heading.
		* `description` - Feature description.
	* `created` - The date and time this catalog was created.
	* `updated` - The date and time this catalog was last updated.
	* `versions` - list of versions. Nested `versions` blocks have the following structure:
		* `id` - Unique ID.
		* `rev` - Cloudant revision.
		* `crn` - Version's CRN.
		* `version` - Version of content type.
		* `sha` - hash of the content.
		* `created` - The date and time this version was created.
		* `updated` - The date and time this version was last updated.
		* `offering_id` - Offering ID.
		* `catalog_id` - Catalog ID.
		* `kind_id` - Kind ID.
		* `tags` - List of tags associated with this catalog.
		* `repo_url` - Content's repo URL.
		* `source_url` - Content's source URL (e.g git repo).
		* `tgz_url` - File used to on-board this version.
		* `configuration` - List of user solicited overrides. Nested `configuration` blocks have the following structure:
			* `key` - Configuration key.
			* `type` - Value type (string, boolean, int).
			* `default_value` - The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.
			* `value_constraint` - Constraint associated with value, e.g., for string type - regx:[a-z].
			* `description` - Key description.
			* `required` - Is key required to install.
			* `options` - List of options of type.
			* `hidden` - Hide values.
		* `metadata` - Open ended metadata information.
		* `validation` - Validation response. Nested `validation` blocks have the following structure:
			* `validated` - Date and time of last successful validation.
			* `requested` - Date and time of last validation was requested.
			* `state` - Current validation state - <empty>, in_progress, valid, invalid, expired.
			* `last_operation` - Last operation (e.g. submit_deployment, generate_installer, install_offering.
			* `target` - Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.
		* `required_resources` - Resource requirments for installation. Nested `required_resources` blocks have the following structure:
			* `type` - Type of requirement.
			* `value` - mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.
		* `single_instance` - Denotes if single instance can be deployed to a given cluster.
		* `install` - Script information. Nested `install` blocks have the following structure:
			* `instructions` - Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.
			* `script` - Optional script that needs to be run post any pre-condition script.
			* `script_permission` - Optional iam permissions that are required on the target cluster to run this script.
			* `delete_script` - Optional script that if run will remove the installed version.
			* `scope` - Optional value indicating if this script is scoped to a namespace or the entire cluster.
		* `pre_install` - Optional pre-install instructions. Nested `pre_install` blocks have the following structure:
			* `instructions` - Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.
			* `script` - Optional script that needs to be run post any pre-condition script.
			* `script_permission` - Optional iam permissions that are required on the target cluster to run this script.
			* `delete_script` - Optional script that if run will remove the installed version.
			* `scope` - Optional value indicating if this script is scoped to a namespace or the entire cluster.
		* `entitlement` - Entitlement license info. Nested `entitlement` blocks have the following structure:
			* `provider_name` - Provider name.
			* `provider_id` - Provider ID.
			* `product_id` - Product ID.
			* `part_numbers` - list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.
			* `image_repo_name` - Image repository name.
		* `licenses` - List of licenses the product was built with. Nested `licenses` blocks have the following structure:
			* `id` - License ID.
			* `name` - license name.
			* `type` - type of license e.g., Apache xxx.
			* `url` - URL for the license text.
			* `description` - License description.
		* `image_manifest_url` - If set, denotes a url to a YAML file with list of container images used by this version.
		* `deprecated` - read only field, indicating if this version is deprecated.
		* `package_version` - Version of the package used to create this version.
		* `state` - Offering state. Nested `state` blocks have the following structure:
			* `current` - one of: new, validated, account-published, ibm-published, public-published.
			* `current_entered` - Date and time of current request.
			* `pending` - one of: new, validated, account-published, ibm-published, public-published.
			* `pending_requested` - Date and time of pending request.
			* `previous` - one of: new, validated, account-published, ibm-published, public-published.
		* `version_locator` - A dotted value of `catalogID`.`versionID`.
		* `console_url` - Console URL.
		* `long_description` - Long description for version.
		* `whitelisted_accounts` - Whitelisted accounts for version.
	* `plans` - list of plans. Nested `plans` blocks have the following structure:
		* `id` - unique id.
		* `label` - Display Name in the requested language.
		* `name` - The programmatic name of this offering.
		* `short_description` - Short description in the requested language.
		* `long_description` - Long description in the requested language.
		* `metadata` - open ended metadata information.
		* `tags` - list of tags associated with this catalog.
		* `additional_features` - list of features associated with this offering. Nested `additional_features` blocks have the following structure:
			* `title` - Heading.
			* `description` - Feature description.
		* `created` - the date'time this catalog was created.
		* `updated` - the date'time this catalog was last updated.
		* `deployments` - list of deployments. Nested `deployments` blocks have the following structure:
			* `id` - unique id.
			* `label` - Display Name in the requested language.
			* `name` - The programmatic name of this offering.
			* `short_description` - Short description in the requested language.
			* `long_description` - Long description in the requested language.
			* `metadata` - open ended metadata information.
			* `tags` - list of tags associated with this catalog.
			* `created` - the date'time this catalog was created.
			* `updated` - the date'time this catalog was last updated.

* `permit_request_ibm_public_publish` - Is it permitted to request publishing to IBM or Public.

* `ibm_publish_approved` - Indicates if this offering has been approved for use by all IBMers.

* `public_publish_approved` - Indicates if this offering has been approved for use by all IBM Cloud users.

* `public_original_crn` - The original offering CRN that this publish entry came from.

* `publish_public_crn` - The crn of the public catalog entry of this offering.

* `portal_approval_record` - The portal's approval record ID.

* `portal_ui_url` - The portal UI URL.

* `catalog_id` - The id of the catalog containing this offering.

* `catalog_name` - The name of the catalog.

* `metadata` - Map of metadata values for this offering.

* `disclaimer` - A disclaimer for this offering.

* `hidden` - Determine if this offering should be displayed in the Consumption UI.

* `provider` - Provider of this offering.

* `repo_info` - Repository info for offerings. Nested `repo_info` blocks have the following structure:
	* `token` - Token for private repos.
	* `type` - Public or enterprise GitHub.

