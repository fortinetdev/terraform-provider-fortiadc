---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_intermediate_ca_group"
description: |-
  Configure fortiadc Intermediate CA certificate group.
---

# fortiadc_system_certificate_intermediate_ca_group
Configure fortiadc Intermediate CA certificate group.

## Example Usage
```hcl
resource "fortiadc_system_certificate_intermediate_ca_group" "ca_int_group_vdom1" {
	mkey = "ca_int_group_vdom1"
	vdom = "vdom1"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - Intermediate CA certificate group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Intermediate Ca Group can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_intermediate_ca_group.labelname {{mkey}}
```
