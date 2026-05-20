---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_cors_headers_child_cors_headers_list"
description: |-
  Configure fortiadc security waf cors headers child cors headers list.
---

# fortiadc_security_waf_cors_headers_child_cors_headers_list
Configure fortiadc security waf cors headers child cors headers list.

## Example Usage
```hcl
resource "fortiadc_security_waf_cors_headers_child_cors_headers_list" "test" {
	pkey = "test"
	mkey = "1"
	header = "testheader"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The id of configuration.
* `pkey` - The name of fortiadc security waf cors headers.
* `header` - Specify the HTTP header as a string.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_cors_headers_child_cors_headers_list.labelname {{mkey}}
```
