---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_category"
description: |-
  Provides a list of fortiadc load balance web category
---

# Data Source: fortiadc_load_balance_web_category
Provides a list of `fortiadc load balance web category`.

## Example Usage

```hcl
 data "fortiadc_load_balance_web_categorylist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_load_balance_web_categorylist.sample1.mkey_list
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_load_balance_web_category`.
