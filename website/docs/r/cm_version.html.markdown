---
layout: "ibm"
page_title: "IBM : cm_version"
sidebar_current: "docs-ibm-resource-cm-version"
description: |-
  Manages cm_version.
---

# ibm\_cm_version

Provides a resource for cm_version. This allows cm_version to be created, updated and deleted.

## Example Usage

```hcl
resource "cm_version" "cm_version" {
  catalog_identifier = "catalog_identifier"
  offering_id = "offering_id"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, Forces new resource, string) Catalog identifier.
* `offering_id` - (Required, Forces new resource, string) Offering identification.
* `tags` - (Optional, Forces new resource, List) Tags array.
* `target_kinds` - (Optional, Forces new resource, List) Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.
* `content` - (Optional, Forces new resource, TypeString) byte array representing the content to be imported.  Only supported for OVA images at this time.
* `zipurl` - (Optional, Forces new resource, string) URL path to zip location.  If not specified, must provide content in the body of this call.
* `target_version` - (Optional, Forces new resource, string) The semver value for this new version, if not found in the zip url package content.
* `include_config` - (Optional, Forces new resource, bool) Add all possible configuration values to this version when importing.
* `repo_type` - (Optional, Forces new resource, string) The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_version.
* `rev` - Cloudant revision.
* `crn` - Version's CRN.
* `version` - Version of content type.
* `sha` - hash of the content.
* `created` - The date and time this version was created.
* `updated` - The date and time this version was last updated.
* `catalog_id` - Catalog ID.
* `kind_id` - Kind ID.
* `repo_url` - Content's repo URL.
* `source_url` - Content's source URL (e.g git repo).
* `tgz_url` - File used to on-board this version.
* `configuration` - List of user solicited overrides.
* `metadata` - Open ended metadata information.
* `validation` - Validation response.
* `required_resources` - Resource requirments for installation.
* `single_instance` - Denotes if single instance can be deployed to a given cluster.
* `install` - Script information.
* `pre_install` - Optional pre-install instructions.
* `entitlement` - Entitlement license info.
* `licenses` - List of licenses the product was built with.
* `image_manifest_url` - If set, denotes a url to a YAML file with list of container images used by this version.
* `deprecated` - read only field, indicating if this version is deprecated.
* `package_version` - Version of the package used to create this version.
* `state` - Offering state.
* `version_locator` - A dotted value of `catalogID`.`versionID`.
* `console_url` - Console URL.
* `long_description` - Long description for version.
* `whitelisted_accounts` - Whitelisted accounts for version.
