---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_vdom_link"
description: |-
  Configure fortiadc system vdom-link.
---

# fortiadc_system_vdom_link
Configure fortiadc system vdom-link.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - vdom-link name.
* `type` - type of vdom-link. Valid values: 1:ethernet, 0:ppp .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Vdom Link can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_vdom_link.labelname {{mkey}}
```
