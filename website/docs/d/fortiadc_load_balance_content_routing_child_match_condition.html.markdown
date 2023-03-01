---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_content_routing_child_match_condition"
description: |-
  Get information on an fortiadc load balance content routing child match condition
---

# Data Source: fortiadc_load_balance_content_routing_child_match_condition
Use this data source to get information on an fortiadc load balance content routing child match condition

## Example Usage

```hcl
 data "fortiadc_load_balance_content_routing_child_match_condition" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_content_routing_child_match_condition.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance content routing child match condition.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - match condition id.
* `ignorecase` - ignore case. 
* `reverse` - reverse match. 
* `object` - match object. 
* `content` - match content. 
* `type` - match type. 

