---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_filter_profile"
description: |-
  Configure fortiadc Web Filter Profile.
---

# fortiadc_load_balance_web_filter_profile
Configure fortiadc Web Filter Profile.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Web Filter Profile Name.
* `description` - Web Filter Profile Description. 



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Web Filter Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_web_filter_profile.labelname {{mkey}}
```
