---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_sig_profile_sub_category_id_list_group"
description: |-
  Configure security waf sig profile sub category id list group.
---

# fortiadc_security_waf_sig_profile_sub_category_id_list_group
Configure security waf sig profile sub category id list group.

## Example Usage
```hcl
resource "fortiadc_security_waf_sig_profile_sub_category_id_list_group" "test_security_waf_sig_profile_sub_category_id_list_group" {
        pkey = "test3"
        type = "subCategory"
        values  {
                mkey = "41"
                parent_id = "6"
                status = "disable"
        }
        values {
                mkey = "43"
                parent_id = "2"
                status = "disable"
        }
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - Name of Web Attack Signature.
* `type` - Sub-category.
* `values` - Specify the Sub-category, action and status in a list values
* `mkey` - Sub-category (ID). Example: Generic SQL Injection (43) in SQL Injection (2).
* `parent_id` - Category (ID)
* `status` - Enable/disable this Sub-category.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Sig Profile Sub Category Id List Group can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_sig_profile_sub_category_id_list_group.labelname SecurityWafSigProfileSubCategoryIdListGroup
```
