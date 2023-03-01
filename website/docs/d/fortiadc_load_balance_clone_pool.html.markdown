---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_clone_pool"
description: |-
  Get information on an fortiadc load balance clone pool
---

# Data Source: fortiadc_load_balance_clone_pool
Use this data source to get information on an fortiadc load balance clone pool

## Example Usage

```hcl
 data "fortiadc_load_balance_clone_pool" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_clone_pool.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance clone pool.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - pool name.



