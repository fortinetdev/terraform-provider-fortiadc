---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_local_cert_group_child_group_member"
description: |-
  Configure fortiadc Local certificate group.
---

# fortiadc_system_certificate_local_cert_group_child_group_member
Configure fortiadc Local certificate group.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - Local certificate group member id.
* `default` - As default certificate. Valid values: enable/disable.
* `local_cert` - Local certificate reference. 
* `ocsp_stapling` - OCSP stapling reference. 
* `extra_intermediate_cag` - Intermediate CA group reference. 
* `extra_local_cert` - Local certificate reference. 
* `extra_ocsp_stapling` - OCSP stapling reference. 
* `intermediate_cag` - Intermediate CA group reference. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Local Cert Group Child Group Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_local_cert_group_child_group_member.labelname {{mkey}}
```
