---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf_child_area"
description: |-
  Configure fortiadc Configure OSPF..
---

# fortiadc_router_ospf_child_area
Configure fortiadc Configure OSPF..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Area entry IP address..
* `stub_type` - Stub type setting.. Valid values: 1:no-summary, 0:summary .
* `authentication` - Authentication type.. Valid values: 1:text, 0:none, 2:md5 .
* `type` - Area type setting.. Valid values: 1:stub, 0:regular .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Ospf Child Area can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_ospf_child_area.labelname {{mkey}}
```
