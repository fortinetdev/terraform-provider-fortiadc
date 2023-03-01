---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence_child_iso8583_bitmap"
description: |-
  Get information on an fortiadc load balance persistence child iso8583 bitmap
---

# Data Source: fortiadc_load_balance_persistence_child_iso8583_bitmap
Use this data source to get information on an fortiadc load balance persistence child iso8583 bitmap

## Example Usage

```hcl
 data "fortiadc_load_balance_persistence_child_iso8583_bitmap" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_persistence_child_iso8583_bitmap.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance persistence child iso8583 bitmap.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - iso8583 bitmap id.
* `type` - iso8583 bitmap type. 

