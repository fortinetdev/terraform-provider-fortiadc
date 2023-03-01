---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_ztna_profile"
description: |-
  Configure fortiadc security ZTNA profile.
---

# fortiadc_security_ztna_profile
Configure fortiadc security ZTNA profile.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - profile name.

* `log_status` - log status. Valid values: enable/disable.

* `comments` - comments. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Ztna Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_ztna_profile.labelname {{mkey}}
```
