---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_filter_profile_child_category_members"
description: |-
  Get information on an fortiadc load balance web filter profile child category members
---

# Data Source: fortiadc_load_balance_web_filter_profile_child_category_members
Use this data source to get information on an fortiadc load balance web filter profile child category members

## Example Usage

```hcl
 data "fortiadc_load_balance_web_filter_profile_child_category_members" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_web_filter_profile_child_category_members.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance web filter profile child category members.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - Web Filter Profile Category Member Id.
* `category` - Web Categroy / Web Sub Category. 

