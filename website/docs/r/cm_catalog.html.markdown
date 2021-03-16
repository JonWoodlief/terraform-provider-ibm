---
layout: "ibm"
page_title: "IBM : cm_catalog"
sidebar_current: "docs-ibm-resource-cm-catalog"
description: |-
  Manages cm_catalog.
---

# ibm\_cm_catalog

Provides a resource for cm_catalog. This allows cm_catalog to be created, updated and deleted.

## Example Usage

```hcl
resource "cm_catalog" "cm_catalog" {
}
```

## Argument Reference

The following arguments are supported:

* `rev` - (Optional, string) Cloudant revision.
* `label` - (Optional, string) Display Name in the requested language.
* `short_description` - (Optional, string) Description in the requested language.
* `catalog_icon_url` - (Optional, string) URL for an icon associated with this catalog.
* `tags` - (Optional, List) List of tags associated with this catalog.
* `features` - (Optional, List) List of features associated with this catalog.
  * `title` - (Optional, string) Heading.
  * `description` - (Optional, string) Feature description.
* `disabled` - (Optional, bool) Denotes whether a catalog is disabled.
* `resource_group_id` - (Optional, string) Resource group id the catalog is owned by.
* `owning_account` - (Optional, string) Account that owns catalog.
* `catalog_filters` - (Optional, List) Filters for account and catalog filters.
  * `include_all` - (Optional, bool) -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.
  * `category_filters` - (Optional, map[string]interface{}) Filter against offering properties.
  * `id_filters` - (Optional, IDFilter) Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
* `syndication_settings` - (Optional, List) Feature information.
  * `remove_related_components` - (Optional, bool) Remove related components.
  * `clusters` - (Optional, []interface{}) Syndication clusters.
  * `history` - (Optional, SyndicationHistory) Feature information.
  * `authorization` - (Optional, SyndicationAuthorization) Feature information.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_catalog.
* `url` - The url for this specific catalog.
* `crn` - CRN associated with the catalog.
* `offerings_url` - URL path to offerings.
* `created` - The date-time this catalog was created.
* `updated` - The date-time this catalog was last updated.
