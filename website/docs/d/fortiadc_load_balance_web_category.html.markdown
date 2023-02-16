---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_category"
description: |-
  Get information on an fortiadc load balance web category
---

# Data Source: fortiadc_load_balance_web_category
Use this data source to get information on an fortiadc load balance web category

## Example Usage

```hcl
 data "fortiadc_load_balance_web_category" sample1 {
  fadc_id = "1"
}

output output1 {
  value = data.fortiadc_load_balance_web_category.sample1
}
```

## Argument Reference
* `id` - (Required) Specify the mkey of the desired  load balance web category.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fadc_id` - Web Category Id.

* `description` - Web Category Description. 
* `mkey` - Web Category Name. 


