---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_stream_scripting"
description: |-
  Configure fortiadc system stream scripting.
---

# fortiadc_system_stream_scripting
Configure fortiadc system stream scripting.

## Example Usage
```hcl
resource "fortiadc_system_stream_scripting" "stream_script" {
        mkey = "COOKIE_COMMANDS"
        script = "when STREAM_REQUEST_DATA{\ndebug(\"test\")\n}"
}
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - stream scripting name.
* `script` - content of scripting. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Stream Scripting can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_stream_scripting.labelname {{mkey}}
```
