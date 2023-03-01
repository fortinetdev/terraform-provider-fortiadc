---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_sdn_connector"
description: |-
  Configure fortiadc sdn connector.
---

# fortiadc_system_sdn_connector
Configure fortiadc sdn connector.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - sdn connector name.
* `status` - enable/disable sdn connector. Valid values: enable/disable.
* `oci_region_type` - sdn connector oci region type. Valid values: 1:GOVERNMENT_REGION, 0:COMMERCIAL_REGION .
* `port` - server port. (1,65535)
* `oci_cert` - sdn connector oci cert. 
* `sap_ms_http_port` - sap message server http port. (1,65535)
* `aws_secretkey` - sdn connector secret token. 
* `update_interval` - interval (seconds). (30,3600)
* `type` - sdn connector type. Valid values: 10:kubernetes, 13:sap, 2:aws, 7:oci .
* `sap_icm_http_port` - sap icm http port. (1,65535)
* `secret_token` - sdn connector secret token. 
* `oci_tenant_id` - sdn connector oci tenant id. 
* `oci_region` - sdn connector oci region. 
* `aws_region` - sdn connector aws region. 
* `oci_user_id` - sdn connector oci user id. 
* `aws_accesskey` - sdn connector secret token. 
* `use_metadata_iam` - enable/disable using iam role. Valid values: enable/disable.
* `sap_dns_suffix` - dns name for sap system. 
* `oci_ha_status` - enable/disable use HA service. Valid values: enable/disable.
* `sap_password` - sap soap sid admin password. 
* `sap_sidadm` - sap soap sid admin. 
* `oci_compartment_id` - sdn connector oci compartment id. 
* `server` - server ip/FQDN/host(host is only for sap). 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Sdn Connector can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_sdn_connector.labelname {{mkey}}
```
