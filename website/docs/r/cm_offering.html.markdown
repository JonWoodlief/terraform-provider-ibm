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

* `catalog_identifier` - (Required, string) Catalog identifier.
* `rev` - (Optional, string) Cloudant revision.
* `url` - (Optional, string) The url for this specific offering.
* `crn` - (Optional, string) The crn for this specific offering.
* `label` - (Optional, string) Display Name in the requested language.
* `name` - (Optional, string) The programmatic name of this offering.
* `offering_icon_url` - (Optional, string) URL for an icon associated with this offering.
* `offering_docs_url` - (Optional, string) URL for an additional docs with this offering.
* `offering_support_url` - (Optional, string) URL to be displayed in the Consumption UI for getting support on this offering.
* `tags` - (Optional, List) List of tags associated with this catalog.
* `rating` - (Optional, List) Repository info for offerings.
  * `one_star_count` - (Optional, int) One start rating.
  * `two_star_count` - (Optional, int) Two start rating.
  * `three_star_count` - (Optional, int) Three start rating.
  * `four_star_count` - (Optional, int) Four start rating.
* `created` - (Optional, TypeString) The date and time this catalog was created.
* `updated` - (Optional, TypeString) The date and time this catalog was last updated.
* `short_description` - (Optional, string) Short description in the requested language.
* `long_description` - (Optional, string) Long description in the requested language.
* `features` - (Optional, List) list of features associated with this offering.
  * `title` - (Optional, string) Heading.
  * `description` - (Optional, string) Feature description.
* `kinds` - (Optional, List) Array of kind.
  * `id` - (Optional, string) Unique ID.
  * `format_kind` - (Optional, string) content kind, e.g., helm, vm image.
  * `target_kind` - (Optional, string) target cloud to install, e.g., iks, open_shift_iks.
  * `metadata` - (Optional, map[string]interface{}) Open ended metadata information.
  * `install_description` - (Optional, string) Installation instruction.
  * `tags` - (Optional, []interface{}) List of tags associated with this catalog.
  * `additional_features` - (Optional, []interface{}) List of features associated with this offering.
  * `created` - (Optional, TypeString) The date and time this catalog was created.
  * `updated` - (Optional, TypeString) The date and time this catalog was last updated.
  * `versions` - (Optional, []interface{}) list of versions.
  * `plans` - (Optional, []interface{}) list of plans.
* `permit_request_ibm_public_publish` - (Optional, bool) Is it permitted to request publishing to IBM or Public.
* `ibm_publish_approved` - (Optional, bool) Indicates if this offering has been approved for use by all IBMers.
* `public_publish_approved` - (Optional, bool) Indicates if this offering has been approved for use by all IBM Cloud users.
* `public_original_crn` - (Optional, string) The original offering CRN that this publish entry came from.
* `publish_public_crn` - (Optional, string) The crn of the public catalog entry of this offering.
* `portal_approval_record` - (Optional, string) The portal's approval record ID.
* `portal_ui_url` - (Optional, string) The portal UI URL.
* `catalog_id` - (Optional, string) The id of the catalog containing this offering.
* `catalog_name` - (Optional, string) The name of the catalog.
* `disclaimer` - (Optional, string) A disclaimer for this offering.
* `hidden` - (Optional, bool) Determine if this offering should be displayed in the Consumption UI.
* `provider` - (Optional, string) Provider of this offering.
* `repo_info` - (Optional, List) Repository info for offerings.
  * `token` - (Optional, string) Token for private repos.
  * `type` - (Optional, string) Public or enterprise GitHub.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_offering.
