---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp"
description: |-
  Get information on an fortiadc router bgp
---

# Data Source: fortiadc_router_bgp
Use this data source to get information on an fortiadc router bgp

## Example Usage

```hcl
 data "fortiadc_router_bgp" sample1 {
}

output output1 {
  value = data.fortiadc_router_bgp.sample1
}
```

## Argument Reference
* `` - (Required) Specify the mkey of the desired  router bgp.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `redistribute_ospf` - Enable/Disable redistribute ospf. 
* `router_id` - Router ID.. 
* `redistribute_connected` - Enable/Disable redistribute connected. 

* `deterministic_med` - Enable/disable enforce deterministic comparison of MED.. 


* `redistribute_connected6` - Enable/Disable redistribute connected6. 
* `as` - Router AS number.. 


* `redistribute_static6` - Enable/Disable redistribute static6.. 
* `always_compare_med` - Enable/disable always compare MED.. 
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths.. 
* `redistribute_static` - Enable/Disable redistribute static.. 


