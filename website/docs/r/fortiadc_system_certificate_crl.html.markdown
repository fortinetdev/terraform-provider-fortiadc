---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_crl"
description: |-
  Configure fortiadc certificate revokation list.
---

# fortiadc_system_certificate_crl
Configure fortiadc certificate revokation list.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - certificate name.

* `ldap_server` - LDAP server. 
* `crldp_status` - Enable/disable certificate revocation list distribution point. Valid values: enable/disable.
* `use_ca_id` - Enable/Disable Message of CRL in SCEP. Valid values: enable/disable.

* `update_ahead_time` - CRL update start ahead time before nextupd (by default in seconds, at least 1h). 

* `update_interval` - CRL update interval in case of no change or failure (by default in seconds, at least 5m). 
* `ca_id` - Message type of CRL update via SCEP. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Crl can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_crl.labelname {{mkey}}
```
