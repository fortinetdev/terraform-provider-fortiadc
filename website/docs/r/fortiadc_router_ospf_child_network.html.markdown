---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf_child_network"
description: |-
  Configure fortiadc Configure OSPF..
---

# fortiadc_router_ospf_child_network
Configure fortiadc Configure OSPF..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Network entry ID..
* `area_id` - Attach the network to area. 
* `prefix` - prefix. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Ospf Child Network can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_ospf_child_network.labelname {{mkey}}
```
