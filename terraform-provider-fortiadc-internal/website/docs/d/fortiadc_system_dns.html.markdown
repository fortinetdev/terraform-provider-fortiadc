---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_dns"
description: |-
  Get information on an fortiadc system dns
---

# Data Source: fortiadc_system_dns
Use this data source to get information on an fortiadc system dns

## Example Usage

```hcl
 data "fortiadc_system_dns" sample1 {
}

output output1 {
  value = data.fortiadc_system_dns.sample1
}
```

## Argument Reference
* `` - (Required) Specify the mkey of the desired  system dns.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `ip_second` - Secondary DNS. 
* `ip_primary` - Primary DNS. 

