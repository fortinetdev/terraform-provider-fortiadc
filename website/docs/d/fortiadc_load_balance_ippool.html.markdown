---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_ippool"
description: |-
  Get information on an fortiadc load balance ippool
---

# Data Source: fortiadc_load_balance_ippool
Use this data source to get information on an fortiadc load balance ippool

## Example Usage

```hcl
 data "fortiadc_load_balance_ippool" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_ippool.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance ippool.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - pool name.
* `pool_type` - address type. 
* `ip_end` - IP max. 

* `ip6_end` - IPv6 max. 
* `ip_start` - IP min. 
* `interface` - interface name. 
* `ip6_start` - IPv6 min. 

