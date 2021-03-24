---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# Catalog Management API data sources
{: #catalog-management-data-sources}

Review the data sources that you can use to retrieve information about your Catalog Management API resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cm_catalog`
{: #cm_catalog}

Retrieve information about cm_catalog.
{: shortdesc}

### Sample Terraform code
{: #cm_catalog-sample}

```
data "ibm_cm_catalog" "cm_catalog" {
  catalog_identifier = "catalog_identifier"
}
```

### Input parameters
{: #cm_catalog-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`catalog_identifier`|String|Required|Catalog identifier.|

### Output parameters
{: #cm_catalog-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|Unique ID.|
|`label`|String|Display Name in the requested language.|
|`short_description`|String|Description in the requested language.|
|`catalog_icon_url`|String|URL for an icon associated with this catalog.|
|`tags`|List|List of tags associated with this catalog.|
|`url`|String|The url for this specific catalog.|
|`crn`|String|CRN associated with the catalog.|
|`offerings_url`|String|URL path to offerings.|

## `ibm_cm_offering`
{: #cm_offering}

Retrieve information about cm_offering.
{: shortdesc}

### Sample Terraform code
{: #cm_offering-sample}

```
data "ibm_cm_offering" "cm_offering" {
  catalog_identifier = "catalog_identifier"
  offering_id = "offering_id"
}
```

### Input parameters
{: #cm_offering-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`catalog_identifier`|String|Required|Catalog identifier.|
|`offering_id`|String|Required|Offering identification.|

### Output parameters
{: #cm_offering-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|unique id.|
|`url`|String|The url for this specific offering.|
|`crn`|String|The crn for this specific offering.|
|`label`|String|Display Name in the requested language.|
|`name`|String|The programmatic name of this offering.|
|`offering_icon_url`|String|URL for an icon associated with this offering.|
|`offering_docs_url`|String|URL for an additional docs with this offering.|
|`offering_support_url`|String|URL to be displayed in the Consumption UI for getting support on this offering.|
|`short_description`|String|Short description in the requested language.|
|`long_description`|String|Long description in the requested language.|
|`permit_request_ibm_public_publish`|Boolean|Is it permitted to request publishing to IBM or Public.|
|`ibm_publish_approved`|Boolean|Indicates if this offering has been approved for use by all IBMers.|
|`public_publish_approved`|Boolean|Indicates if this offering has been approved for use by all IBM Cloud users.|
|`public_original_crn`|String|The original offering CRN that this publish entry came from.|
|`publish_public_crn`|String|The crn of the public catalog entry of this offering.|
|`portal_approval_record`|String|The portal's approval record ID.|
|`portal_ui_url`|String|The portal UI URL.|
|`catalog_id`|String|The id of the catalog containing this offering.|
|`catalog_name`|String|The name of the catalog.|
|`disclaimer`|String|A disclaimer for this offering.|
|`hidden`|Boolean|Determine if this offering should be displayed in the Consumption UI.|
|`repo_info`|List|Repository info for offerings. This list contains only one item.|
|`repo_info.token`|String|Token for private repos.|
|`repo_info.type`|String|Public or enterprise GitHub.|

## `ibm_cm_version`
{: #cm_version}

Retrieve information about cm_version.
{: shortdesc}

### Sample Terraform code
{: #cm_version-sample}

```
data "ibm_cm_version" "cm_version" {
  version_loc_id = "version_loc_id"
}
```

### Input parameters
{: #cm_version-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`version_loc_id`|String|Required|A dotted value of `catalogID`.`versionID`.|

### Output parameters
{: #cm_version-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|Unique ID.|
|`crn`|String|Version's CRN.|
|`version`|String|Version of content type.|
|`sha`|String|hash of the content.|
|`created`|String|The date and time this version was created.|
|`updated`|String|The date and time this version was last updated.|
|`catalog_id`|String|Catalog ID.|
|`repo_url`|String|Content's repo URL.|
|`source_url`|String|Content's source URL (e.g git repo).|
|`tgz_url`|String|File used to on-board this version.|

## `ibm_cm_offering_instance`
{: #cm_offering_instance}

Retrieve information about cm_offering_instance.
{: shortdesc}

### Sample Terraform code
{: #cm_offering_instance-sample}

```
data "ibm_cm_offering_instance" "cm_offering_instance" {
  instance_identifier = "instance_identifier"
}
```

### Input parameters
{: #cm_offering_instance-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`instance_identifier`|String|Required|Version Instance identifier.|

### Output parameters
{: #cm_offering_instance-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|provisioned instance ID (part of the CRN).|
|`url`|String|url reference to this object.|
|`crn`|String|platform CRN for this instance.|
|`label`|String|the label for this instance.|
|`catalog_id`|String|Catalog ID this instance was created from.|
|`offering_id`|String|Offering ID this instance was created from.|
|`kind_format`|String|the format this instance has (helm, operator, ova...).|
|`version`|String|The version this instance was installed from (not version id).|
|`cluster_id`|String|Cluster ID.|
|`cluster_region`|String|Cluster region (e.g., us-south).|
|`cluster_namespaces`|List|List of target namespaces to install into.|
|`cluster_all_namespaces`|Boolean|designate to install into all namespaces.|

