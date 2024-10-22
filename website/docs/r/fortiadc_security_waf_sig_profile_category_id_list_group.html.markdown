---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_sig_profile_category_id_list_group"
description: |-
  Configure security waf sig profile category id list group
---

# fortiadc_security_waf_sig_profile_category_id_list_group
Configure security waf sig profile category id list group.

## Example Usage
```hcl
resource "fortiadc_security_waf_sig_profile_category_id_list_group" "test_security_waf_sig_profile_category_id_list_group" {
        pkey = "test3"
        type = "category"
        values  {
                mkey = "2"
                action = "block"
                status = "disable"
        }
        values {
                mkey = "3"
                action = "block"
                status = "disable"
        }
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - Name of Web Attack Signature.
* `type` - Category.
* `values` - Specify the Category, action and status in a list values
* `mkey` - Category (ID). Example: SQL Injection (2), Known Exploits (6).
* `action` - Alert, Deny, Block, Silent-deny, Captcha or Test.
* `status` - Enable/disable this Category.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Sig Profile Category Id List Group can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_sig_profile_category_id_list_group.labelname SecurityWafSigProfileCategoryIdListGroup
```
