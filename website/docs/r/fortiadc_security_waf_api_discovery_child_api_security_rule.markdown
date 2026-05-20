---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_api_discovery_child_api_security_rule"
description: |-
  Configure fortiadc security waf api discovery child api security rule.
---

# fortiadc_security_waf_api_discovery_child_api_security_rule
Configure fortiadc security waf api discovery child api security rule.

## Example Usage
```hcl
resource "fortiadc_security_waf_api_discovery_child_api_security_rule" "test" {
	pkey = "test"
	mkey = "1"
	base_url = "demo.fortinet.com"
	path = "/login"
	req_rate = "1000"
	action = "alert"
	severity = "low"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The id of configuration.
* `pkey` - The name of fortiadc security waf api discovery.
* `base_url` - Specify the HTTP Host header.
* `path` - Specify the API resource path.
* `req_rate` - Specify the allowable requests per second. 
* `action` - Specify a WAF action object to apply when a bot is detected.
* `severity` - Select the event severity to log when a bot is detected.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_api_discovery_child_api_security_rule.labelname {{mkey}}
```
