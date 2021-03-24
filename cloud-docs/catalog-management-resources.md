---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# Catalog Management API resources
{: #catalog-management-resources}

Create, update, or delete Catalog Management API resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cm_catalog`
{: #cm_catalog}

Create, update, or delete an cm_catalog.
{: shortdesc}

### Sample Terraform code
{: #cm_catalog-sample}

```
resource "ibm_cm_catalog" "cm_catalog" {
  label = "placeholder"
  short_description = "placeholder"
  catalog_icon_url = "placeholder"
  tags = [ "placeholder" ]
}
```

### Input parameters
{: #cm_catalog-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`label`|String|Optional|Display Name in the requested language.|No|
|`short_description`|String|Optional|Description in the requested language.|No|
|`catalog_icon_url`|String|Optional|URL for an icon associated with this catalog.|No|
|`tags`|List|Optional|List of tags associated with this catalog.|No|

### Output parameters
{: #cm_catalog-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_catalog.|
|`url`|String|The url for this specific catalog.|
|`crn`|String|CRN associated with the catalog.|
|`offerings_url`|String|URL path to offerings.|

### Import
{: #cm_catalog-import}

`ibm_cm_catalog` can be imported by ID

```
$ terraform import ibm_cm_catalog.example sample-id
```

## `ibm_cm_offering`
{: #cm_offering}

Create, update, or delete an cm_offering.
{: shortdesc}

### Sample Terraform code
{: #cm_offering-sample}

```
resource "ibm_cm_offering" "cm_offering" {
  catalog_id = "catalog_id"
  label = "placeholder"
  offering_icon_url = "placeholder"
  offering_docs_url = "placeholder"
  offering_support_url = "placeholder"
  tags = [ "placeholder" ]
}
```

### Input parameters
{: #cm_offering-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`catalog_id`|String|Required|Catalog identifier.|No|
|`label`|String|Optional|Display Name in the requested language.|No|
|`offering_icon_url`|String|Optional|URL for an icon associated with this offering.|No|
|`offering_docs_url`|String|Optional|URL for an additional docs with this offering.|No|
|`offering_support_url`|String|Optional|URL to be displayed in the Consumption UI for getting support on this offering.|No|
|`tags`|List|Optional|List of tags associated with this catalog.|No|

### Output parameters
{: #cm_offering-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_offering.|

### Import
{: #cm_offering-import}

`ibm_cm_offering` can be imported by ID

```
$ terraform import ibm_cm_offering.example sample-id
```

## `ibm_cm_version`
{: #cm_version}

Create, update, or delete an cm_version.
{: shortdesc}

### Sample Terraform code
{: #cm_version-sample}

```
resource "ibm_cm_version" "cm_version" {
  catalog_identifier = "catalog_identifier"
  offering_id = "offering_id"
  tags = [ "placeholder" ]
  target_kinds = [ "placeholder" ]
  zipurl = "placeholder"
  target_version = "placeholder"
}
```

### Input parameters
{: #cm_version-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`catalog_identifier`|String|Required|Catalog identifier.|Yes|
|`offering_id`|String|Required|Offering identification.|Yes|
|`tags`|List|Optional|Tags array.|Yes|
|`target_kinds`|List|Optional|Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.|Yes|
|`zipurl`|String|Optional|URL path to zip location.  If not specified, must provide content in the body of this call.|Yes|
|`target_version`|String|Optional|The semver value for this new version, if not found in the zip url package content.|Yes|

### Output parameters
{: #cm_version-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_version.|
|`crn`|String|Version's CRN.|
|`version`|String|Version of content type.|
|`sha`|String|hash of the content.|
|`created`|String|The date and time this version was created.|
|`updated`|String|The date and time this version was last updated.|
|`catalog_id`|String|Catalog ID.|
|`kind_id`|String|Kind ID.|
|`repo_url`|String|Content's repo URL.|
|`source_url`|String|Content's source URL (e.g git repo).|
|`tgz_url`|String|File used to on-board this version.|

### Import
{: #cm_version-import}

`ibm_cm_version` can be imported by ID

```
$ terraform import ibm_cm_version.example sample-id
```

## `ibm_cm_offering_instance`
{: #cm_offering_instance}

Create, update, or delete an cm_offering_instance.
{: shortdesc}

### Sample Terraform code
{: #cm_offering_instance-sample}

```
resource "ibm_cm_offering_instance" "cm_offering_instance" {
  url = "placeholder"
  crn = "placeholder"
  label = "placeholder"
  catalog_id = "placeholder"
  offering_id = "placeholder"
  kind_format = "placeholder"
  version = "placeholder"
  cluster_id = "placeholder"
  cluster_region = "placeholder"
  cluster_namespaces = [ "placeholder" ]
  cluster_all_namespaces = false
}
```

### Input parameters
{: #cm_offering_instance-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`url`|String|Optional|url reference to this object.|No|
|`crn`|String|Optional|platform CRN for this instance.|No|
|`label`|String|Optional|the label for this instance.|No|
|`catalog_id`|String|Optional|Catalog ID this instance was created from.|No|
|`offering_id`|String|Optional|Offering ID this instance was created from.|No|
|`kind_format`|String|Optional|the format this instance has (helm, operator, ova...).|No|
|`version`|String|Optional|The version this instance was installed from (not version id).|No|
|`cluster_id`|String|Optional|Cluster ID.|No|
|`cluster_region`|String|Optional|Cluster region (e.g., us-south).|No|
|`cluster_namespaces`|List|Optional|List of target namespaces to install into.|No|
|`cluster_all_namespaces`|Boolean|Optional|designate to install into all namespaces.|No|

### Output parameters
{: #cm_offering_instance-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_offering_instance.|

### Import
{: #cm_offering_instance-import}

`ibm_cm_offering_instance` can be imported by ID

```
$ terraform import ibm_cm_offering_instance.example sample-id
```

