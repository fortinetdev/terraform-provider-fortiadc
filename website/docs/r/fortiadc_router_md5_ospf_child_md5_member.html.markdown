---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_md5_ospf_child_md5_member"
description: |-
  Configure fortiadc Configure md5 authentication list for ospf..
---

# fortiadc_router_md5_ospf_child_md5_member
Configure fortiadc Configure md5 authentication list for ospf..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - id..
* `key` - md5 key. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Md5 Ospf Child Md5 Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_md5_ospf_child_md5_member.labelname {{mkey}}
```
