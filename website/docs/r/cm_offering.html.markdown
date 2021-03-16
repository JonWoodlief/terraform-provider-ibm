---
layout: "ibm"
page_title: "IBM : cm_offering"
sidebar_current: "docs-ibm-resource-cm-offering"
description: |-
  Manages cm_offering.
---

# ibm\_cm_offering

Provides a resource for cm_offering. This allows cm_offering to be created, updated and deleted.

## Example Usage

```hcl
resource "cm_offering" "cm_offering" {
  catalog_identifier = "catalog_identifier"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, Forces new resource, string) Catalog identifier.
* `tags` - (Optional, Forces new resource, List) Tags array.
* `target_kinds` - (Optional, Forces new resource, List) Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.
* `content` - (Optional, Forces new resource, TypeString) byte array representing the content to be imported.  Only supported for OVA images at this time.
* `zipurl` - (Optional, Forces new resource, string) URL path to zip location.  If not specified, must provide content in this post body.
* `offering_id` - (Optional, Forces new resource, string) Re-use the specified offeringID during import.
* `target_version` - (Optional, Forces new resource, string) The semver value for this new version.
* `include_config` - (Optional, Forces new resource, bool) Add all possible configuration items when creating this version.
* `repo_type` - (Optional, Forces new resource, string) The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'.
* `x_auth_token` - (Optional, Forces new resource, string) Authentication token used to access the specified zip file.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_offering.
* `rev` - Cloudant revision.
* `url` - The url for this specific offering.
* `crn` - The crn for this specific offering.
* `label` - Display Name in the requested language.
* `name` - The programmatic name of this offering.
* `offering_icon_url` - URL for an icon associated with this offering.
* `offering_docs_url` - URL for an additional docs with this offering.
* `offering_support_url` - URL to be displayed in the Consumption UI for getting support on this offering.
* `rating` - Repository info for offerings.
* `created` - The date and time this catalog was created.
* `updated` - The date and time this catalog was last updated.
* `short_description` - Short description in the requested language.
* `long_description` - Long description in the requested language.
* `features` - list of features associated with this offering.
* `kinds` - Array of kind.
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
* `repo_info` - Repository info for offerings.
