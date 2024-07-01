---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_dns_vdom"
description: |-
  Get information on an fortiadc system dns per vdom
---

# Data Source: fortiadc_system_dns_vdom
Use this data source to get information on an fortiadc system dns per vdom

## Example Usage

```hcl
 data "fortiadc_system_dns_vdom" sample1 {
}

output output1 {
  value = data.fortiadc_system_dns_vdom.sample1
}
```

## Argument Reference
* `` - (Required) Specify the mkey of the desired system dns.
* `vdom` - (Required) Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only non-root vdom can be specified.


## Attribute Reference

The following attributes are exported:

* `ip_second` - Secondary DNS. 
* `ip_primary` - Primary DNS. 
* `dns_overide` - Enable/Disable DNS.
