---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_certificate_verify_child_group_member"
description: |-
  Configure fortiadc Certificate Verify.
---

# fortiadc_system_certificate_certificate_verify_child_group_member
Configure fortiadc Certificate Verify.

## Example Usage
```hcl
resource "fortiadc_system_certificate_certificate_verify_child_group_member" "cv_member1" {
	pkey = "cv1"
	mkey = "1"
	ca = "Fortinet_CA"
	depends_on = [fortiadc_system_certificate_certificate_verify.cert_verify]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - Verify group member id.
* `ca` - CA certificate reference. 
* `crl` - CRL reference. 
* `ocsp` - OCSP reference. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Certificate Verify Child Group Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_certificate_verify_child_group_member.labelname {{mkey}}
```
