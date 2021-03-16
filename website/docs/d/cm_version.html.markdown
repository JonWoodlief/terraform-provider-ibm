---
layout: "ibm"
page_title: "IBM : cm_version"
sidebar_current: "docs-ibm-datasource-cm-version"
description: |-
  Get information about cm_version
---

# ibm\_cm_version

Provides a read-only data source for cm_version. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_version" "cm_version" {
	version_loc_id = "version_loc_id"
}
```

## Argument Reference

The following arguments are supported:

* `version_loc_id` - (Required, string) A dotted value of `catalogID`.`versionID`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_version.
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

