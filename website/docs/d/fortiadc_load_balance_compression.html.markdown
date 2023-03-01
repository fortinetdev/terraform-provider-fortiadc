---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_compression"
description: |-
  Get information on an fortiadc load balance compression
---

# Data Source: fortiadc_load_balance_compression
Use this data source to get information on an fortiadc load balance compression

## Example Usage

```hcl
 data "fortiadc_load_balance_compression" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_compression.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance compression.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - compression name.
* `max_cpu_usage` - max cpu usage. (1,100)
* `level` - compression level. (1,9)
* `uri_list_type` - uri list type. 

* `gzip_window_size` - gzip window size in Kbytes. 
* `gzip_memory_level` - gzip memory level in Kbytes. 
* `cpu_limit` - (en|dis)able cpu limit. 
* `min_content_length` - minimum content length. (1,2147483647)
* `method` - compression method. 


