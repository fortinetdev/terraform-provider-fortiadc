---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_policy"
description: |-
  Configure fortiadc firewall policy.
---

# fortiadc_firewall_policy
Configure fortiadc firewall policy.

## Example Usage
```hcl
resource "fortiadc_firewall_policy" "firewall_policy" {
	default_action = "deny"
	deny_log = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `default_action` - Deny or Accept.
* `deny_log` - enable/disable logging for denied action. This option is available is the action is Deny.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Firewall Policy can be imported using any of these accepted formats:
```
$ terraform import fortiadc_firewall_policy.labelname FirewallPolicy
```
