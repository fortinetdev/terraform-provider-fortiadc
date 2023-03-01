---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_l2_exception_list"
description: |-
  Get information on an fortiadc load balance l2 exception list
---

# Data Source: fortiadc_load_balance_l2_exception_list
Use this data source to get information on an fortiadc load balance l2 exception list

## Example Usage

```hcl
 data "fortiadc_load_balance_l2_exception_list" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_l2_exception_list.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance l2 exception list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - name.
* `description` - layer2 exception list description. 


* `web_filter_profile` - layer2 exception web filter profile. 

