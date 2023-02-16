---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_ippool_child_node_member"
description: |-
  Configure fortiadc load-balance ip pool info.
---

# fortiadc_load_balance_ippool_child_node_member
Configure fortiadc load-balance ip pool info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - pool name.
* `ip6_max` - IPv6 max. 
* `ip6_min` - IPv6 min. 
* `ha_node_num` - ha-node id. (0,7)
* `ip_min` - IP min. 
* `pool_type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `ip_max` - IP max. 
* `interface` - interface name. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Ippool Child Node Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_ippool_child_node_member.labelname {{mkey}}
```
