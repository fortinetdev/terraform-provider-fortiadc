---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_address6"
description: |-
  Get information on an fortiadc system address6
---

# Data Source: fortiadc_system_address6
Use this data source to get information on an fortiadc system address6

## Example Usage

```hcl
 data "fortiadc_system_address6" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_address6.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system address6.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - name.
* `ip6_network` - IPv6/Prefix. 
* `ip6_max` - IP6 max. 
* `ip6_min` - IPv6 min. 
* `type` - type. 

