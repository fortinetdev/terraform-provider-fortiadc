---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_adaptive_learning_child_url_list"
description: |-
  Configure fortiadc security waf adaptive learning child url list.
---

# fortiadc_security_waf_adaptive_learning_child_url_list
Configure fortiadc security waf adaptive learning child url list.

## Example Usage
```hcl
resource "fortiadc_security_waf_adaptive_learning_child_url_list" "test" {
	mkey = "1"
	pkey = "test"
	host_status = "enable"
	host = "0.0.0.0"
	url = "/"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The id of configuration.
* `pkey` - The name of security waf adaptive learning.
* `host_status` - If enabled, require authorization only for the specified host. 
* `host` - Specify the HTTP Host header.
* `url` - The literal URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_adaptive_learning_child_url_list.labelname {{mkey}}
```
