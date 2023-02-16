---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_decompression"
description: |-
  Configure fortiadc load-balance decompression info.
---

# fortiadc_load_balance_decompression
Configure fortiadc load-balance decompression info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - decompression name.
* `max_cpu_usage` - max cpu usage. (1,100)
* `uri_list_type` - uri list type. Valid values: 1:include, 2:exclude .

* `cpu_limit` - (en|dis)able cpu limit. Valid values: enable/disable.
* `method` - decompression method. Valid values: 1:gzip, 3:all, 2:deflate .


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Decompression can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_decompression.labelname {{mkey}}
```
