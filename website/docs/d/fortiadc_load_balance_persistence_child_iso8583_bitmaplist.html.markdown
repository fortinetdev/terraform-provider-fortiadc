---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence_child_iso8583_bitmap"
description: |-
  Provides a list of fortiadc load balance persistence child iso8583 bitmap
---

# Data Source: fortiadc_load_balance_persistence_child_iso8583_bitmap
Provides a list of `fortiadc load balance persistence child iso8583 bitmap`.

## Example Usage

```hcl
 data "fortiadc_load_balance_persistence_child_iso8583_bitmaplist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_load_balance_persistence_child_iso8583_bitmaplist.sample1.mkey_list
}
```

## Argument Reference

* `pkey` - The parent key.
* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_load_balance_persistence_child_iso8583_bitmap`.
