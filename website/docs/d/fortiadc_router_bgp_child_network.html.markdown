---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp_child_network"
description: |-
  Get information on an fortiadc router bgp child network
---

# Data Source: fortiadc_router_bgp_child_network
Use this data source to get information on an fortiadc router bgp child network

## Example Usage

```hcl
 data "fortiadc_router_bgp_child_network" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_router_bgp_child_network.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  router bgp child network.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - Network entry ID..
* `prefix6` - Network6 prefix.. 
* `prefix` - Network prefix.. 
* `type` - address type. 

