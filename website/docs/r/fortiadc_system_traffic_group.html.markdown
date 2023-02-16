---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_traffic_group"
description: |-
  Configure fortiadc traffic group parameters configuration.
---

# fortiadc_system_traffic_group
Configure fortiadc traffic group parameters configuration.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - traffic group name.
* `network_failover` - network failover track. Valid values: enable/disable.
* `preempt` - preempt mode. Valid values: enable/disable.
* `failover_order` - node id. (0,7)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Traffic Group can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_traffic_group.labelname {{mkey}}
```
