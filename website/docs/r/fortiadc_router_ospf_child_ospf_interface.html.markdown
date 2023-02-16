---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf_child_ospf_interface"
description: |-
  Configure fortiadc Configure OSPF..
---

# fortiadc_router_ospf_child_ospf_interface
Configure fortiadc Configure OSPF..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Interface entry name..
* `transmit_delay` - Transmit delay.. (1,65535)
* `retransmit_interval` - Retransmit interval.. (3,65535)
* `hello_interval` - Hello interval.. (1,65535)
* `text` - Authentication text key.. 
* `mtu_ignore` - Enable/disable ignore MTU.. Valid values: enable/disable.
* `priority` - Priority.. (0,255)
* `authentication` - Authentication type.. Valid values: 1:text, 0:none, 2:md5 .
* `dead_interval` - Dead interval.. (1,65535)
* `interface` - Configuration interface name.. 
* `md5` - Authentication md5 key list.. 
* `network_type` - Network type.. Valid values: 1:point-to-point, 0:broadcast, 2:point-to-multipoint .
* `cost` - Cost of the interface.. (0,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Ospf Child Ospf Interface can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_ospf_child_ospf_interface.labelname {{mkey}}
```
