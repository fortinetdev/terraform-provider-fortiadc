---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_http2_profile"
description: |-
  Get information on an fortiadc load balance http2 profile
---

# Data Source: fortiadc_load_balance_http2_profile
Use this data source to get information on an fortiadc load balance http2 profile

## Example Usage

```hcl
 data "fortiadc_load_balance_http2_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_http2_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance http2 profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - Profile name.
* `max_concurrent_stream` - Maximum number of concurrent stream. (1,200)
* `max_frame_size` - Maximum size of frame. (16384,131072)
* `priority_mode` - Priority of stream mode. 
* `ssl_constraint` - SSL constraint check. 
* `max_receive_window` - Maximum size of receive window. (16384,524288)
* `header_table_size` - size of header table for HPACK. (4096,65536)
* `upgrade_mode` - Protocol upgrade to HTTP/2 mode. 
* `max_header_list_size` - Maximum size of header list. (4096,262144)

