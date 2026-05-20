---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_advanced_protection_child_advanced_protection_rule"
description: |-
  Configure fortiadc security waf advanced protection child advanced protection rule.
---

# fortiadc_security_waf_advanced_protection_child_advanced_protection_rule
Configure fortiadc security waf advanced protection child advanced protection rule.

## Example Usage
```hcl
resource "fortiadc_security_waf_advanced_protection_child_advanced_protection_rule" "test" {
	pkey = "test"
	mkey = "1"
	type = "content-scraping"
	content_type = "text/html application/json"
	occurrence_limit = "99"
	occurrence_within = "66"
	percentage_match = "0"
	action = "alert"
	severity = "medium"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The id of configuration.
* `pkey` - The name of fortiadc security waf advanced protection.
* `type` - Content Scraping or HTTP Response Code.
* `content_type` - Specify a Content Type for the Content Scraping rule.
* `resp_code` - Specify a Response Code for the HTTP Response Code rule.
* `occurrence_limit` - Sets the condition for the limit of the number of responses received from the specified type.
* `occurrence_within` - Sets the time span during which to count how many times a response is received from the specified type.
* `percentage_match` - Sets the condition for what percentage of the traffic received is from the specified type, during the given time frame.
* `action` - Select which action profile that you want to apply.
* `severity` - When FortiADC records violations of this rule in the attack log, each log message contains a Severity Level field. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_advanced_protection_child_advanced_protection_rule.labelname {{mkey}}
```
