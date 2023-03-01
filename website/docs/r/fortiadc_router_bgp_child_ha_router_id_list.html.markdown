---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp_child_ha_router_id_list"
description: |-
  Configure fortiadc Configure BGP..
---

# fortiadc_router_bgp_child_ha_router_id_list
Configure fortiadc Configure BGP..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Router-id list entry ID..
* `node` - Node ID of HA Node. (0,7)
* `router_id` - Router-id.. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Bgp Child Ha Router Id List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_bgp_child_ha_router_id_list.labelname {{mkey}}
```
