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
resource "ibm_cm_offering" "cm_offering" {
  catalog_id = "catalog_id"
  label = "placeholder"
}
```

## Argument Reference

The following arguments are supported:

* `label` - (Optional, string) Display Name in the requested language.
* `short_description` - (Optional, string) Description in the requested language.
* `catalog_icon_url` - (Optional, string) URL for an icon associated with this catalog.
* `tags` - (Optional, List) List of tags associated with this catalog.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_catalog.
* `url` - The url for this specific catalog.
* `crn` - CRN associated with the catalog.
* `offerings_url` - URL path to offerings.