---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_http2_profile"
description: |-
  Configure fortiadc load-balance HTTP/2 profile.
---

# fortiadc_load_balance_http2_profile
Configure fortiadc load-balance HTTP/2 profile.

## Example Usage
```hcl
resource "fortiadc_load_balance_http2_profile" "http2_profile" {
	mkey = "http2_profile_test"
	ssl_constraint = "enable"
	max_concurrent_stream = "7"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Profile name.
* `max_concurrent_stream` - Maximum number of concurrent stream. (1,200)
* `max_frame_size` - Maximum size of frame. (16384,131072)
* `priority_mode` - Priority of stream mode. Valid values: 0:best-effort .
* `ssl_constraint` - SSL constraint check. Valid values: 1:enable, 0:disable .
* `max_receive_window` - Maximum size of receive window. (16384,524288)
* `header_table_size` - size of header table for HPACK. (4096,65536)
* `upgrade_mode` - Protocol upgrade to HTTP/2 mode. Valid values: 0:upgradeable .
* `max_header_list_size` - Maximum size of header list. (4096,262144)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Http2 Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_http2_profile.labelname {{mkey}}
```
