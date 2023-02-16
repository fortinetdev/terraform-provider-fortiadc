---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp_child_neighbor"
description: |-
  Configure fortiadc Configure BGP..
---

# fortiadc_router_bgp_child_neighbor
Configure fortiadc Configure BGP..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Neighbor entry ID..
* `ip` - IP address of neighbor.. 
* `weight` - Default weight for routes from this neighbor.. (0,65535)
* `distribute_list_out6` - Filter for IP updates to this neighbor.. 
* `distribute_list_out` - Filter for IP updates to this neighbor.. 
* `prefix_list_in6` - IP Inbound filter for updates from this neighbor.. 
* `prefix_list_out` - IP Outbound filter for updates to this neighbor.. 
* `prefix_list_out6` - IP Outbound filter for updates to this neighbor.. 
* `prefix_list_in` - IP Inbound filter for updates from this neighbor.. 
* `type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `distribute_list_in` - Filter for IP updates from this neighbor.. 
* `distribute_list_in6` - Filter for IP updates from this neighbor.. 
* `interface` - Interface.. 
* `ip6` - IPv6 address of neighbor.. 
* `remote_as` - AS number of neighbor.. 
* `holdtime` - Number of seconds to mark peer as dead.. (0,65535)
* `port` - Port.. (0,65535)
* `keepalive` - Frequency to send keep alive requests.. (0,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Bgp Child Neighbor can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_bgp_child_neighbor.labelname {{mkey}}
```
