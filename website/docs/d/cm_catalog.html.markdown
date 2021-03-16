---
layout: "ibm"
page_title: "IBM : cm_catalog"
sidebar_current: "docs-ibm-datasource-cm-catalog"
description: |-
  Get information about cm_catalog
---

# ibm\_cm_catalog

Provides a read-only data source for cm_catalog. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_catalog" "cm_catalog" {
	catalog_identifier = "catalog_identifier"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, string) Catalog identifier.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_catalog.
* `id` - Unique ID.

* `rev` - Cloudant revision.

* `label` - Display Name in the requested language.

* `short_description` - Description in the requested language.

* `catalog_icon_url` - URL for an icon associated with this catalog.

* `tags` - List of tags associated with this catalog.

* `url` - The url for this specific catalog.

* `crn` - CRN associated with the catalog.

* `offerings_url` - URL path to offerings.

* `features` - List of features associated with this catalog. Nested `features` blocks have the following structure:
	* `title` - Heading.
	* `description` - Feature description.

* `disabled` - Denotes whether a catalog is disabled.

* `created` - The date-time this catalog was created.

* `updated` - The date-time this catalog was last updated.

* `resource_group_id` - Resource group id the catalog is owned by.

* `owning_account` - Account that owns catalog.

* `catalog_filters` - Filters for account and catalog filters. Nested `catalog_filters` blocks have the following structure:
	* `include_all` - -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.
	* `category_filters` - Filter against offering properties.
	* `id_filters` - Filter on offering ID's. There is an include filter and an exclule filter. Both can be set. Nested `id_filters` blocks have the following structure:
		* `include` - Offering filter terms. Nested `include` blocks have the following structure:
			* `filter_terms` - List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
		* `exclude` - Offering filter terms. Nested `exclude` blocks have the following structure:
			* `filter_terms` - List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.

* `syndication_settings` - Feature information. Nested `syndication_settings` blocks have the following structure:
	* `remove_related_components` - Remove related components.
	* `clusters` - Syndication clusters. Nested `clusters` blocks have the following structure:
		* `region` - Cluster region.
		* `id` - Cluster ID.
		* `name` - Cluster name.
		* `resource_group_name` - Resource group ID.
		* `type` - Syndication type.
		* `namespaces` - Syndicated namespaces.
		* `all_namespaces` - Syndicated to all namespaces on cluster.
	* `history` - Feature information. Nested `history` blocks have the following structure:
		* `namespaces` - Array of syndicated namespaces.
		* `clusters` - Array of syndicated namespaces. Nested `clusters` blocks have the following structure:
			* `region` - Cluster region.
			* `id` - Cluster ID.
			* `name` - Cluster name.
			* `resource_group_name` - Resource group ID.
			* `type` - Syndication type.
			* `namespaces` - Syndicated namespaces.
			* `all_namespaces` - Syndicated to all namespaces on cluster.
		* `last_run` - Date and time last syndicated.
	* `authorization` - Feature information. Nested `authorization` blocks have the following structure:
		* `token` - Array of syndicated namespaces.
		* `last_run` - Date and time last updated.

