---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp"
description: |-
  Configure fortiadc Configure BGP..
---

# fortiadc_router_bgp
Configure fortiadc Configure BGP..

## Example Usage
```hcl
resource "fortiadc_router_bgp" "bgp" {
	as = "1"
	router_id = "10.106.203.254"
	redistribute_static = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `redistribute_ospf` - Enable/Disable redistribute ospf. Valid values: enable/disable.
* `router_id` - Router ID.. 
* `redistribute_connected` - Enable/Disable redistribute connected. Valid values: enable/disable.

* `deterministic_med` - Enable/disable enforce deterministic comparison of MED.. Valid values: enable/disable.


* `redistribute_connected6` - Enable/Disable redistribute connected6. Valid values: enable/disable.
* `as` - Router AS number.. 


* `redistribute_static6` - Enable/Disable redistribute static6.. Valid values: enable/disable.
* `always_compare_med` - Enable/disable always compare MED.. Valid values: enable/disable.
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths.. Valid values: enable/disable.
* `redistribute_static` - Enable/Disable redistribute static.. Valid values: enable/disable.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Bgp can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_bgp.labelname RouterBgp
```
