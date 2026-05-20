---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_allowed_origin_child_allowed_origin_list"
description: |-
  Configure fortiadc security waf allowed origin child allowed origin list.
---

# fortiadc_security_waf_allowed_origin_child_allowed_origin_list
Configure fortiadc security waf allowed origin child allowed origin list.

## Example Usage
```hcl
resource "fortiadc_security_waf_allowed_origin_child_allowed_origin_list" "test" {
	pkey = "test"
	mkey = "1"
	protocol = "HTTP"
	origin_name = "test"
	port = "80"
	include_sub_domains = "enable"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The id of configuration.
* `pkey` - The name of fortiadc security waf allowed origin.
* `protocol` - Select which type of protocols are allowed for the connections between foreign applications and your application. Valid values: HTTP, HTTPS, ANY.
* `origin_name` - Enter the foreign application's domain name or IP address.
* `port` - Specify the TCP port number for the CORS connections.
* `include_sub_domains` - Enable/disable to allow/disallow the Origin Value to match with the domains of its sub level.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_allowed_origin_child_allowed_origin_list.labelname {{mkey}}
```
