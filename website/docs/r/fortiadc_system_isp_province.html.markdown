---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_isp_province"
description: |-
  Configure fortiadc isp provice info.
---

# fortiadc_system_isp_province
Configure fortiadc isp provice info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of province.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Isp Province can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_isp_province.labelname {{mkey}}
```
