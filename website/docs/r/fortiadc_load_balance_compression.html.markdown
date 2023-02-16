---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_compression"
description: |-
  Configure fortiadc load-balance compression info.
---

# fortiadc_load_balance_compression
Configure fortiadc load-balance compression info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - compression name.
* `max_cpu_usage` - max cpu usage. (1,100)
* `level` - compression level. (1,9)
* `uri_list_type` - uri list type. Valid values: 1:include, 2:exclude .

* `gzip_window_size` - gzip window size in Kbytes. Valid values: 11:8, 10:4, 13:32, 12:16, 15:128, 14:64, 9:2, 8:1 .
* `gzip_memory_level` - gzip memory level in Kbytes. Valid values: 1:1, 3:4, 2:2, 5:16, 4:8, 7:64, 6:32, 9:256, 8:128 .
* `cpu_limit` - (en|dis)able cpu limit. Valid values: enable/disable.
* `min_content_length` - minimum content length. (1,2147483647)
* `method` - compression method. Valid values: 1:gzip, 3:all, 2:deflate .


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Compression can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_compression.labelname {{mkey}}
```
