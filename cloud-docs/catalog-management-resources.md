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
  rev = "placeholder"
  label = "placeholder"
  short_description = "placeholder"
  catalog_icon_url = "placeholder"
  tags = [ "placeholder" ]
  features = var.cm_catalog_features
  disabled = false
  resource_group_id = "placeholder"
  owning_account = "placeholder"
  catalog_filters = var.cm_catalog_catalog_filters
  syndication_settings = var.cm_catalog_syndication_settings
}
```

### Input parameters
{: #cm_catalog-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`rev`|String|Optional|Cloudant revision.|No|
|`label`|String|Optional|Display Name in the requested language.|No|
|`short_description`|String|Optional|Description in the requested language.|No|
|`catalog_icon_url`|String|Optional|URL for an icon associated with this catalog.|No|
|`tags`|List|Optional|List of tags associated with this catalog.|No|
|`features`|List|Optional|List of features associated with this catalog.|No|
|`disabled`|Boolean|Optional|Denotes whether a catalog is disabled.|No|
|`resource_group_id`|String|Optional|Resource group id the catalog is owned by.|No|
|`owning_account`|String|Optional|Account that owns catalog.|No|
|`catalog_filters`|List|Optional|Filters for account and catalog filters. You can specify one item in this list only.|No|
|`syndication_settings`|List|Optional|Feature information. You can specify one item in this list only.|No|

### Output parameters
{: #cm_catalog-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_catalog.|
|`url`|String|The url for this specific catalog.|
|`crn`|String|CRN associated with the catalog.|
|`offerings_url`|String|URL path to offerings.|
|`created`|String|The date-time this catalog was created.|
|`updated`|String|The date-time this catalog was last updated.|

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
  catalog_identifier = "catalog_identifier"
  tags = [ "placeholder" ]
  target_kinds = [ "placeholder" ]
  content = 
  zipurl = "placeholder"
  offering_id = "placeholder"
  target_version = "placeholder"
  include_config = false
  repo_type = "placeholder"
  x_auth_token = "placeholder"
}
```

### Input parameters
{: #cm_offering-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`catalog_identifier`|String|Required|Catalog identifier.|Yes|
|`tags`|List|Optional|Tags array.|Yes|
|`target_kinds`|List|Optional|Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.|Yes|
|`content`|String|Optional|byte array representing the content to be imported.  Only supported for OVA images at this time.|Yes|
|`zipurl`|String|Optional|URL path to zip location.  If not specified, must provide content in this post body.|Yes|
|`offering_id`|String|Optional|Re-use the specified offeringID during import.|Yes|
|`target_version`|String|Optional|The semver value for this new version.|Yes|
|`include_config`|Boolean|Optional|Add all possible configuration items when creating this version.|Yes|
|`repo_type`|String|Optional|The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'.|Yes|
|`x_auth_token`|String|Optional|Authentication token used to access the specified zip file.|Yes|

### Output parameters
{: #cm_offering-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_offering.|
|`rev`|String|Cloudant revision.|
|`url`|String|The url for this specific offering.|
|`crn`|String|The crn for this specific offering.|
|`label`|String|Display Name in the requested language.|
|`name`|String|The programmatic name of this offering.|
|`offering_icon_url`|String|URL for an icon associated with this offering.|
|`offering_docs_url`|String|URL for an additional docs with this offering.|
|`offering_support_url`|String|URL to be displayed in the Consumption UI for getting support on this offering.|
|`rating`|List|Repository info for offerings. This list contains only one item.|
|`rating.one_star_count`|Integer|One start rating.|
|`rating.two_star_count`|Integer|Two start rating.|
|`rating.three_star_count`|Integer|Three start rating.|
|`rating.four_star_count`|Integer|Four start rating.|
|`created`|String|The date and time this catalog was created.|
|`updated`|String|The date and time this catalog was last updated.|
|`short_description`|String|Short description in the requested language.|
|`long_description`|String|Long description in the requested language.|
|`features`|List|list of features associated with this offering.|
|`features.title`|String|Heading.|
|`features.description`|String|Feature description.|
|`kinds`|List|Array of kind.|
|`kinds.id`|String|Unique ID.|
|`kinds.format_kind`|String|content kind, e.g., helm, vm image.|
|`kinds.target_kind`|String|target cloud to install, e.g., iks, open_shift_iks.|
|`kinds.metadata`|Map|Open ended metadata information.|
|`kinds.install_description`|String|Installation instruction.|
|`kinds.tags`|List|List of tags associated with this catalog.|
|`kinds.additional_features`|List|List of features associated with this offering.|
|`kinds.additional_features.title`|String|Heading.|
|`kinds.additional_features.description`|String|Feature description.|
|`kinds.created`|String|The date and time this catalog was created.|
|`kinds.updated`|String|The date and time this catalog was last updated.|
|`kinds.versions`|List|list of versions.|
|`kinds.versions.id`|String|Unique ID.|
|`kinds.versions.rev`|String|Cloudant revision.|
|`kinds.versions.crn`|String|Version's CRN.|
|`kinds.versions.version`|String|Version of content type.|
|`kinds.versions.sha`|String|hash of the content.|
|`kinds.versions.created`|String|The date and time this version was created.|
|`kinds.versions.updated`|String|The date and time this version was last updated.|
|`kinds.versions.offering_id`|String|Offering ID.|
|`kinds.versions.catalog_id`|String|Catalog ID.|
|`kinds.versions.kind_id`|String|Kind ID.|
|`kinds.versions.tags`|List|List of tags associated with this catalog.|
|`kinds.versions.repo_url`|String|Content's repo URL.|
|`kinds.versions.source_url`|String|Content's source URL (e.g git repo).|
|`kinds.versions.tgz_url`|String|File used to on-board this version.|
|`kinds.versions.configuration`|List|List of user solicited overrides.|
|`kinds.versions.configuration.key`|String|Configuration key.|
|`kinds.versions.configuration.type`|String|Value type (string, boolean, int).|
|`kinds.versions.configuration.default_value`|Map|The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.|
|`kinds.versions.configuration.value_constraint`|String|Constraint associated with value, e.g., for string type - regx:[a-z].|
|`kinds.versions.configuration.description`|String|Key description.|
|`kinds.versions.configuration.required`|Boolean|Is key required to install.|
|`kinds.versions.configuration.options`|List|List of options of type.|
|`kinds.versions.configuration.hidden`|Boolean|Hide values.|
|`kinds.versions.metadata`|Map|Open ended metadata information.|
|`kinds.versions.validation`|List|Validation response. This list contains only one item.|
|`kinds.versions.validation.validated`|String|Date and time of last successful validation.|
|`kinds.versions.validation.requested`|String|Date and time of last validation was requested.|
|`kinds.versions.validation.state`|String|Current validation state - <empty>, in_progress, valid, invalid, expired.|
|`kinds.versions.validation.last_operation`|String|Last operation (e.g. submit_deployment, generate_installer, install_offering.|
|`kinds.versions.validation.target`|Map|Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.|
|`kinds.versions.required_resources`|List|Resource requirments for installation.|
|`kinds.versions.required_resources.type`|String|Type of requirement.|
|`kinds.versions.required_resources.value`|Map|mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.|
|`kinds.versions.single_instance`|Boolean|Denotes if single instance can be deployed to a given cluster.|
|`kinds.versions.install`|List|Script information. This list contains only one item.|
|`kinds.versions.install.instructions`|String|Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.|
|`kinds.versions.install.script`|String|Optional script that needs to be run post any pre-condition script.|
|`kinds.versions.install.script_permission`|String|Optional iam permissions that are required on the target cluster to run this script.|
|`kinds.versions.install.delete_script`|String|Optional script that if run will remove the installed version.|
|`kinds.versions.install.scope`|String|Optional value indicating if this script is scoped to a namespace or the entire cluster.|
|`kinds.versions.pre_install`|List|Optional pre-install instructions.|
|`kinds.versions.pre_install.instructions`|String|Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.|
|`kinds.versions.pre_install.script`|String|Optional script that needs to be run post any pre-condition script.|
|`kinds.versions.pre_install.script_permission`|String|Optional iam permissions that are required on the target cluster to run this script.|
|`kinds.versions.pre_install.delete_script`|String|Optional script that if run will remove the installed version.|
|`kinds.versions.pre_install.scope`|String|Optional value indicating if this script is scoped to a namespace or the entire cluster.|
|`kinds.versions.entitlement`|List|Entitlement license info. This list contains only one item.|
|`kinds.versions.entitlement.provider_name`|String|Provider name.|
|`kinds.versions.entitlement.provider_id`|String|Provider ID.|
|`kinds.versions.entitlement.product_id`|String|Product ID.|
|`kinds.versions.entitlement.part_numbers`|List|list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.|
|`kinds.versions.entitlement.image_repo_name`|String|Image repository name.|
|`kinds.versions.licenses`|List|List of licenses the product was built with.|
|`kinds.versions.licenses.id`|String|License ID.|
|`kinds.versions.licenses.name`|String|license name.|
|`kinds.versions.licenses.type`|String|type of license e.g., Apache xxx.|
|`kinds.versions.licenses.url`|String|URL for the license text.|
|`kinds.versions.licenses.description`|String|License description.|
|`kinds.versions.image_manifest_url`|String|If set, denotes a url to a YAML file with list of container images used by this version.|
|`kinds.versions.deprecated`|Boolean|read only field, indicating if this version is deprecated.|
|`kinds.versions.package_version`|String|Version of the package used to create this version.|
|`kinds.versions.state`|List|Offering state. This list contains only one item.|
|`kinds.versions.state.current`|String|one of: new, validated, account-published, ibm-published, public-published.|
|`kinds.versions.state.current_entered`|String|Date and time of current request.|
|`kinds.versions.state.pending`|String|one of: new, validated, account-published, ibm-published, public-published.|
|`kinds.versions.state.pending_requested`|String|Date and time of pending request.|
|`kinds.versions.state.previous`|String|one of: new, validated, account-published, ibm-published, public-published.|
|`kinds.versions.version_locator`|String|A dotted value of `catalogID`.`versionID`.|
|`kinds.versions.console_url`|String|Console URL.|
|`kinds.versions.long_description`|String|Long description for version.|
|`kinds.versions.whitelisted_accounts`|List|Whitelisted accounts for version.|
|`kinds.plans`|List|list of plans.|
|`kinds.plans.id`|String|unique id.|
|`kinds.plans.label`|String|Display Name in the requested language.|
|`kinds.plans.name`|String|The programmatic name of this offering.|
|`kinds.plans.short_description`|String|Short description in the requested language.|
|`kinds.plans.long_description`|String|Long description in the requested language.|
|`kinds.plans.metadata`|Map|open ended metadata information.|
|`kinds.plans.tags`|List|list of tags associated with this catalog.|
|`kinds.plans.additional_features`|List|list of features associated with this offering.|
|`kinds.plans.additional_features.title`|String|Heading.|
|`kinds.plans.additional_features.description`|String|Feature description.|
|`kinds.plans.created`|String|the date'time this catalog was created.|
|`kinds.plans.updated`|String|the date'time this catalog was last updated.|
|`kinds.plans.deployments`|List|list of deployments.|
|`kinds.plans.deployments.id`|String|unique id.|
|`kinds.plans.deployments.label`|String|Display Name in the requested language.|
|`kinds.plans.deployments.name`|String|The programmatic name of this offering.|
|`kinds.plans.deployments.short_description`|String|Short description in the requested language.|
|`kinds.plans.deployments.long_description`|String|Long description in the requested language.|
|`kinds.plans.deployments.metadata`|Map|open ended metadata information.|
|`kinds.plans.deployments.tags`|List|list of tags associated with this catalog.|
|`kinds.plans.deployments.created`|String|the date'time this catalog was created.|
|`kinds.plans.deployments.updated`|String|the date'time this catalog was last updated.|
|`permit_request_ibm_public_publish`|Boolean|Is it permitted to request publishing to IBM or Public.|
|`ibm_publish_approved`|Boolean|Indicates if this offering has been approved for use by all IBMers.|
|`public_publish_approved`|Boolean|Indicates if this offering has been approved for use by all IBM Cloud users.|
|`public_original_crn`|String|The original offering CRN that this publish entry came from.|
|`publish_public_crn`|String|The crn of the public catalog entry of this offering.|
|`portal_approval_record`|String|The portal's approval record ID.|
|`portal_ui_url`|String|The portal UI URL.|
|`catalog_id`|String|The id of the catalog containing this offering.|
|`catalog_name`|String|The name of the catalog.|
|`metadata`|Map|Map of metadata values for this offering.|
|`disclaimer`|String|A disclaimer for this offering.|
|`hidden`|Boolean|Determine if this offering should be displayed in the Consumption UI.|
|`provider`|String|Provider of this offering.|
|`repo_info`|List|Repository info for offerings. This list contains only one item.|
|`repo_info.token`|String|Token for private repos.|
|`repo_info.type`|String|Public or enterprise GitHub.|

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
  content = 
  zipurl = "placeholder"
  target_version = "placeholder"
  include_config = false
  repo_type = "placeholder"
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
|`content`|String|Optional|byte array representing the content to be imported.  Only supported for OVA images at this time.|Yes|
|`zipurl`|String|Optional|URL path to zip location.  If not specified, must provide content in the body of this call.|Yes|
|`target_version`|String|Optional|The semver value for this new version, if not found in the zip url package content.|Yes|
|`include_config`|Boolean|Optional|Add all possible configuration values to this version when importing.|Yes|
|`repo_type`|String|Optional|The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'.|Yes|

### Output parameters
{: #cm_version-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cm_version.|
|`rev`|String|Cloudant revision.|
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
|`configuration`|List|List of user solicited overrides.|
|`configuration.key`|String|Configuration key.|
|`configuration.type`|String|Value type (string, boolean, int).|
|`configuration.default_value`|Map|The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.|
|`configuration.value_constraint`|String|Constraint associated with value, e.g., for string type - regx:[a-z].|
|`configuration.description`|String|Key description.|
|`configuration.required`|Boolean|Is key required to install.|
|`configuration.options`|List|List of options of type.|
|`configuration.hidden`|Boolean|Hide values.|
|`metadata`|Map|Open ended metadata information.|
|`validation`|List|Validation response. This list contains only one item.|
|`validation.validated`|String|Date and time of last successful validation.|
|`validation.requested`|String|Date and time of last validation was requested.|
|`validation.state`|String|Current validation state - <empty>, in_progress, valid, invalid, expired.|
|`validation.last_operation`|String|Last operation (e.g. submit_deployment, generate_installer, install_offering.|
|`validation.target`|Map|Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.|
|`required_resources`|List|Resource requirments for installation.|
|`required_resources.type`|String|Type of requirement.|
|`required_resources.value`|Map|mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.|
|`single_instance`|Boolean|Denotes if single instance can be deployed to a given cluster.|
|`install`|List|Script information. This list contains only one item.|
|`install.instructions`|String|Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.|
|`install.script`|String|Optional script that needs to be run post any pre-condition script.|
|`install.script_permission`|String|Optional iam permissions that are required on the target cluster to run this script.|
|`install.delete_script`|String|Optional script that if run will remove the installed version.|
|`install.scope`|String|Optional value indicating if this script is scoped to a namespace or the entire cluster.|
|`pre_install`|List|Optional pre-install instructions.|
|`pre_install.instructions`|String|Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.|
|`pre_install.script`|String|Optional script that needs to be run post any pre-condition script.|
|`pre_install.script_permission`|String|Optional iam permissions that are required on the target cluster to run this script.|
|`pre_install.delete_script`|String|Optional script that if run will remove the installed version.|
|`pre_install.scope`|String|Optional value indicating if this script is scoped to a namespace or the entire cluster.|
|`entitlement`|List|Entitlement license info. This list contains only one item.|
|`entitlement.provider_name`|String|Provider name.|
|`entitlement.provider_id`|String|Provider ID.|
|`entitlement.product_id`|String|Product ID.|
|`entitlement.part_numbers`|List|list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.|
|`entitlement.image_repo_name`|String|Image repository name.|
|`licenses`|List|List of licenses the product was built with.|
|`licenses.id`|String|License ID.|
|`licenses.name`|String|license name.|
|`licenses.type`|String|type of license e.g., Apache xxx.|
|`licenses.url`|String|URL for the license text.|
|`licenses.description`|String|License description.|
|`image_manifest_url`|String|If set, denotes a url to a YAML file with list of container images used by this version.|
|`deprecated`|Boolean|read only field, indicating if this version is deprecated.|
|`package_version`|String|Version of the package used to create this version.|
|`state`|List|Offering state. This list contains only one item.|
|`state.current`|String|one of: new, validated, account-published, ibm-published, public-published.|
|`state.current_entered`|String|Date and time of current request.|
|`state.pending`|String|one of: new, validated, account-published, ibm-published, public-published.|
|`state.pending_requested`|String|Date and time of pending request.|
|`state.previous`|String|one of: new, validated, account-published, ibm-published, public-published.|
|`version_locator`|String|A dotted value of `catalogID`.`versionID`.|
|`console_url`|String|Console URL.|
|`long_description`|String|Long description for version.|
|`whitelisted_accounts`|List|Whitelisted accounts for version.|

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
  x_auth_refresh_token = "x_auth_refresh_token"
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
|`x_auth_refresh_token`|String|Required|IAM Refresh token.|No|
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
