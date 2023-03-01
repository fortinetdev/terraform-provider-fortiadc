---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_allowlist"
description: |-
  Configure fortiadc Configure geography ip allow list..
---

# fortiadc_load_balance_allowlist
Configure fortiadc Configure geography ip allow list..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name of allow list..
* `status` - status. Valid values: enable/disable.
* `description` - simple description.. 



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Allowlist can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_allowlist.labelname {{mkey}}
```
