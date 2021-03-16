---
layout: "ibm"
page_title: "IBM : cm_version_instance"
sidebar_current: "docs-ibm-resource-cm-version-instance"
description: |-
  Manages cm_version_instance.
---

# ibm\_cm_version_instance

Provides a resource for cm_version_instance. This allows cm_version_instance to be created, updated and deleted.

## Example Usage

```hcl
resource "cm_version_instance" "cm_version_instance" {
  x_auth_refresh_token = "x_auth_refresh_token"
}
```

## Argument Reference

The following arguments are supported:

* `x_auth_refresh_token` - (Required, string) IAM Refresh token.
* `url` - (Optional, string) url reference to this object.
* `crn` - (Optional, string) platform CRN for this instance.
* `label` - (Optional, string) the label for this instance.
* `catalog_id` - (Optional, string) Catalog ID this instance was created from.
* `offering_id` - (Optional, string) Offering ID this instance was created from.
* `kind_format` - (Optional, string) the format this instance has (helm, operator, ova...).
* `version` - (Optional, string) The version this instance was installed from (not version id).
* `cluster_id` - (Optional, string) Cluster ID.
* `cluster_region` - (Optional, string) Cluster region (e.g., us-south).
* `cluster_namespaces` - (Optional, List) List of target namespaces to install into.
* `cluster_all_namespaces` - (Optional, bool) designate to install into all namespaces.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_version_instance.
