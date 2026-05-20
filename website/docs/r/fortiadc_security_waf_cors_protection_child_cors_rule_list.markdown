---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_cors_protection_child_cors_rule_list"
description: |-
  Configure fortiadc security waf cors protection child cors rule list.
---

# fortiadc_security_waf_cors_protection_child_cors_rule_list
Configure fortiadc security waf cors protection child cors rule list.

## Example Usage
```hcl
resource "fortiadc_security_waf_cors_protection_child_cors_rule_list" "test" {
	pkey = "test"
	mkey = "1"
	action = "block"
	host_status = "enable"
	host = "test"
	request_url = "/."
	apply_to_all_cors_traffic = "disable"
	allowed_origin = "test"
	insert_allowed_credentials = "enable"
	allowed_credentials = "false"
	insert_max_age = "enable"
	allowed_max_age = "0"
	allowed_methods = "enable"
	methods = "GET POST"
	allowed_headers = "enable"
	allowed_headers_list = "test"
	exposed_headers = "enable"
	exposed_headers_list = "test"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The id of configuration.
* `pkey` - The name of fortiadc security waf cors protection.
* `action` - Specify the WAF action.
* `host_status` - Enable/disable to allow this rule to protect a specific domain name or IP address.
* `host` - Specify the host name.
* `request_url` - Specify the request URL as a regular expression.
* `apply_to_all_cors_traffic` - Enable/disable to apply the CORS Protection Rule to all CORS traffic.
* `allowed_origin` - Specify the name of the Allowed Origin.
* `insert_allowed_credentials` - Enable/disable to allow whether the CORS requests from foreign applications can include user credentials.
* `allowed_credentials` - This option appears if Insert Allow Credentials is enabled. Ture or False.
* `insert_max_age` - Enable/disable to specify a maximum time period before the result of the preflight request expires.
* `allowed_max_age` - Specify the maximum time period in seconds.
* `allowed_methods` - Enable/disable to allow FortiADC to use the Methods specified to verify whether the methods used in the CORS requests are legitimate.
* `methods` - Specify the methods.
* `allowed_headers` - Enable/disable to allow FortiADC to use the CORS Headers List to verify whether the headers used in the CORS requests are legitimate.
* `allowed_headers_list` - Specify the name of the CORS Headers List to allow.
* `exposed_headers` - Enable/disable to allow FortiADC to expose the specified headers in the CORS Headers List in JavaScript and share with foreign applications.
* `exposed_headers_list` - Specify the name of the CORS Headers List to expose.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_cors_protection_child_cors_rule_list.labelname {{mkey}}
```
