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
|`rev`|String|Cloudant revision.|
|`label`|String|Display Name in the requested language.|
|`short_description`|String|Description in the requested language.|
|`catalog_icon_url`|String|URL for an icon associated with this catalog.|
|`tags`|List|List of tags associated with this catalog.|
|`url`|String|The url for this specific catalog.|
|`crn`|String|CRN associated with the catalog.|
|`offerings_url`|String|URL path to offerings.|
|`features`|List|List of features associated with this catalog.|
|`features.title`|String|Heading.|
|`features.description`|String|Feature description.|
|`disabled`|Boolean|Denotes whether a catalog is disabled.|
|`created`|String|The date-time this catalog was created.|
|`updated`|String|The date-time this catalog was last updated.|
|`resource_group_id`|String|Resource group id the catalog is owned by.|
|`owning_account`|String|Account that owns catalog.|
|`catalog_filters`|List|Filters for account and catalog filters. This list contains only one item.|
|`catalog_filters.include_all`|Boolean|-> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.|
|`catalog_filters.category_filters`|Map|Filter against offering properties.|
|`catalog_filters.id_filters`|List|Filter on offering ID's. There is an include filter and an exclule filter. Both can be set. This list contains only one item.|
|`catalog_filters.id_filters.include`|List|Offering filter terms. This list contains only one item.|
|`catalog_filters.id_filters.include.filter_terms`|List|List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.|
|`catalog_filters.id_filters.exclude`|List|Offering filter terms. This list contains only one item.|
|`catalog_filters.id_filters.exclude.filter_terms`|List|List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.|
|`syndication_settings`|List|Feature information. This list contains only one item.|
|`syndication_settings.remove_related_components`|Boolean|Remove related components.|
|`syndication_settings.clusters`|List|Syndication clusters.|
|`syndication_settings.clusters.region`|String|Cluster region.|
|`syndication_settings.clusters.id`|String|Cluster ID.|
|`syndication_settings.clusters.name`|String|Cluster name.|
|`syndication_settings.clusters.resource_group_name`|String|Resource group ID.|
|`syndication_settings.clusters.type`|String|Syndication type.|
|`syndication_settings.clusters.namespaces`|List|Syndicated namespaces.|
|`syndication_settings.clusters.all_namespaces`|Boolean|Syndicated to all namespaces on cluster.|
|`syndication_settings.history`|List|Feature information. This list contains only one item.|
|`syndication_settings.history.namespaces`|List|Array of syndicated namespaces.|
|`syndication_settings.history.clusters`|List|Array of syndicated namespaces.|
|`syndication_settings.history.clusters.region`|String|Cluster region.|
|`syndication_settings.history.clusters.id`|String|Cluster ID.|
|`syndication_settings.history.clusters.name`|String|Cluster name.|
|`syndication_settings.history.clusters.resource_group_name`|String|Resource group ID.|
|`syndication_settings.history.clusters.type`|String|Syndication type.|
|`syndication_settings.history.clusters.namespaces`|List|Syndicated namespaces.|
|`syndication_settings.history.clusters.all_namespaces`|Boolean|Syndicated to all namespaces on cluster.|
|`syndication_settings.history.last_run`|String|Date and time last syndicated.|
|`syndication_settings.authorization`|List|Feature information. This list contains only one item.|
|`syndication_settings.authorization.token`|String|Array of syndicated namespaces.|
|`syndication_settings.authorization.last_run`|String|Date and time last updated.|

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
|`rev`|String|Cloudant revision.|
|`url`|String|The url for this specific offering.|
|`crn`|String|The crn for this specific offering.|
|`label`|String|Display Name in the requested language.|
|`name`|String|The programmatic name of this offering.|
|`offering_icon_url`|String|URL for an icon associated with this offering.|
|`offering_docs_url`|String|URL for an additional docs with this offering.|
|`offering_support_url`|String|URL to be displayed in the Consumption UI for getting support on this offering.|
|`tags`|List|List of tags associated with this catalog.|
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
|`rev`|String|Cloudant revision.|
|`crn`|String|Version's CRN.|
|`version`|String|Version of content type.|
|`sha`|String|hash of the content.|
|`created`|String|The date and time this version was created.|
|`updated`|String|The date and time this version was last updated.|
|`offering_id`|String|Offering ID.|
|`catalog_id`|String|Catalog ID.|
|`kind_id`|String|Kind ID.|
|`tags`|List|List of tags associated with this catalog.|
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

