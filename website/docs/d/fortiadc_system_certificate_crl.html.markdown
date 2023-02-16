---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_crl"
description: |-
  Get information on an fortiadc system certificate crl
---

# Data Source: fortiadc_system_certificate_crl
Use this data source to get information on an fortiadc system certificate crl

## Example Usage

```hcl
 data "fortiadc_system_certificate_crl" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_certificate_crl.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system certificate crl.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - certificate name.

* `ldap_server` - LDAP server. 
* `crldp_status` - Enable/disable certificate revocation list distribution point. 
* `use_ca_id` - Enable/Disable Message of CRL in SCEP. 

* `update_ahead_time` - CRL update start ahead time before nextupd (by default in seconds, at least 1h). 

* `update_interval` - CRL update interval in case of no change or failure (by default in seconds, at least 5m). 
* `ca_id` - Message type of CRL update via SCEP. 

