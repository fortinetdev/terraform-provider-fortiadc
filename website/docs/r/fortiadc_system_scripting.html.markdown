---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_scripting"
description: |-
  Configure fortiadc system scripting.
---

# fortiadc_system_scripting
Configure fortiadc system scripting.

## Example Usage
```hcl
resource "fortiadc_system_scripting" "script" {
        mkey = "HTTP_TEST2"
        script = "when HTTP_REQUEST{\ndebug(\"test\")\n}"
}
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - scripting name.
* `script` - content of scripting. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Scripting can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_scripting.labelname {{mkey}}
```
