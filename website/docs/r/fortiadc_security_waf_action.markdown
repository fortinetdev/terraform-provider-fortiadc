---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_action"
description: |-
  Configure fortiadc security waf action.
---

# fortiadc_security_waf_action
Configure fortiadc security waf action.

## Example Usage
```hcl
resource "fortiadc_security_waf_action" "test" {
	mkey = "test"
	action_type = "deny"
	deny_code = "400"
	log_status = "enable"
	comments = "test"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `action_type` - Select which action FortiADC takes when the conditions are fulfilled for WAF. Valid values: pass, deny, period-block, redirect, captcha.
* `deny_code` - The Deny Code option is available if the Action Type is Deny or Period Block.
* `block_period` - The Period Block option is available if the Action Type is Period Block. Range: 1-3600.
* `redirect_url` - The Redirect URL option is available if the Action Type is Redirect.
* `log_status` - Enable/Disable log of events.
* `comments` - Enter comment or description of the action for your records.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_action.labelname {{mkey}}
```
