---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp_child_network"
description: |-
  Configure fortiadc Configure BGP..
---

# fortiadc_router_bgp_child_network
Configure fortiadc Configure BGP..

## Example Usage
```hcl
resource "fortiadc_router_bgp_child_network" "bgp_network" {
	mkey = "1"
	type = "ipv4"
	prefix = "10.10.10.100"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Network entry ID..
* `prefix6` - Network6 prefix.. 
* `prefix` - Network prefix.. 
* `type` - address type. Valid values: 1:ipv4, 2:ipv6 .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Bgp Child Network can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_bgp_child_network.labelname {{mkey}}
```
