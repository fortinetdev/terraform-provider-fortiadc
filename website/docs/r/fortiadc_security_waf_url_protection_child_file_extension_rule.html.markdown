---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_url_protection_child_file_extension_rule"
description: |-
  Configure security waf url protection child file extension rule.
---

# fortiadc_security_waf_url_protection_child_file_extension_rule
Configure security waf url protection child file extension rule.

## Example Usage
```hcl
resource "fortiadc_security_waf_url_protection_child_file_extension_rule" "security_waf_url_protection_child_file_extension_rule" {
        pkey = "test"
        mkey = "1"
        file_extension_rule_regex = "tmp"
        file_extension_rule_action = "alert"
        file_extension_rule_severity = "low"
        exception_name = "test"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The name of Waf Url Protection.
* `mkey` - The number of configuration.
* `file_extension_rule_regex` - Matching string.
* `file_extension_rule_action` - Select the action profile that you want to apply.
* `file_extension_rule_severity` - High, Medium or Low.
* `exception_name` -  Select an exception configuration object.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Url Protection Child File Extension Rule can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_url_protection_child_file_extension_rule.labelname SecurityWafUrlProtectionChildFileExtensionRule
```
