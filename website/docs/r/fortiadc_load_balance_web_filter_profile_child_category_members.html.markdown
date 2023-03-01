---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_filter_profile_child_category_members"
description: |-
  Configure fortiadc Web Filter Profile.
---

# fortiadc_load_balance_web_filter_profile_child_category_members
Configure fortiadc Web Filter Profile.

## Example Usage
```hcl
resource "fortiadc_load_balance_web_filter_profile_child_category_members" "web_filt_child" {
	mkey = "1"
	pkey = "web_filter_1"
	category = "Security Risk"
	depends_on = [fortiadc_load_balance_web_filter_profile.web_filt]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - Web Filter Profile Category Member Id.
* `category` - Web Categroy / Web Sub Category. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Web Filter Profile Child Category Members can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_web_filter_profile_child_category_members.labelname {{mkey}}
```
