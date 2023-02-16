---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_sub_category"
description: |-
  Configure fortiadc Web Sub Category.
---

# fortiadc_load_balance_web_sub_category
Configure fortiadc Web Sub Category.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `fadc_id` - Web Sub Category Id.
* `mkey` - Web Sub Category Name. 
* `description` - Web Category Description. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Web Sub Category can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_web_sub_category.labelname {{mkey}}
```
