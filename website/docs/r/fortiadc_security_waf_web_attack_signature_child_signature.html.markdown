---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_web_attack_signature_child_signature"
description: |-
  Configure security waf web attack signature child signature.
---

# fortiadc_security_waf_web_attack_signature_child_signature
Configure security waf web attack signature child signature.

## Example Usage
```hcl
resource "fortiadc_security_waf_web_attack_signature_child_signature" "security_waf_web_attack_signature_child_signature" {
        pkey = "test3"
        mkey = "1002001218"
        status = "disable"
        exception_name = "test"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The name of Web Attack Signature.
* `mkey` - Signature ID.
* `status` - Enable/disable the signature.
* `exception_name` - Select an exception object.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
  Security Waf Web Attack Signature Child Signature. can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_web_attack_signature_child_signature.labelname {{mkey}}
```
