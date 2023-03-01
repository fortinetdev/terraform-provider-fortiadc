---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_l2_exception_list"
description: |-
  Configure fortiadc load-balance layer2 exception list.
---

# fortiadc_load_balance_l2_exception_list
Configure fortiadc load-balance layer2 exception list.

## Example Usage
```hcl
resource "fortiadc_load_balance_l2_exception_list" "lb_l2_exception" {
	mkey = "l2_exception"
	description = "test_l2_exception"
	web_filter_profile = "web_filter_1"
	depends_on = [fortiadc_load_balance_web_filter_profile.web_filt]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `description` - layer2 exception list description. 


* `web_filter_profile` - layer2 exception web filter profile. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance L2 Exception List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_l2_exception_list.labelname {{mkey}}
```
