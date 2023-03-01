---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_intermediate_ca_group_child_group_member"
description: |-
  Configure fortiadc Intermediate CA certificate group.
---

# fortiadc_system_certificate_intermediate_ca_group_child_group_member
Configure fortiadc Intermediate CA certificate group.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - Intermediate CA certificate group member id.
* `intmed_ca` - Intermediate CA certificate reference. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Intermediate Ca Group Child Group Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_intermediate_ca_group_child_group_member.labelname {{mkey}}
```
