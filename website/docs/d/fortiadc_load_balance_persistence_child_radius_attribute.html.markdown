---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence_child_radius_attribute"
description: |-
  Get information on an fortiadc load balance persistence child radius attribute
---

# Data Source: fortiadc_load_balance_persistence_child_radius_attribute
Use this data source to get information on an fortiadc load balance persistence child radius attribute

## Example Usage

```hcl
 data "fortiadc_load_balance_persistence_child_radius_attribute" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_persistence_child_radius_attribute.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance persistence child radius attribute.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - radius attribute id.
* `vendor_id` - vendor id, 0 means will use entire attribute as input of persistence. 
* `vendor_type` - vendor type, 0 means will use entire attribute as input of persistence. (0,255)
* `type` - radius attribute type. 

