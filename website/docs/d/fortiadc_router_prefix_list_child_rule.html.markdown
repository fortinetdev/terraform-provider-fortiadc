---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_prefix_list_child_rule"
description: |-
  Get information on an fortiadc router prefix list child rule
---

# Data Source: fortiadc_router_prefix_list_child_rule
Use this data source to get information on an fortiadc router prefix list child rule

## Example Usage

```hcl
 data "fortiadc_router_prefix_list_child_rule" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_router_prefix_list_child_rule.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  router prefix list child rule.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - id..
* `le` - Maximum prefix length to be matched.. 
* `prefix` - Prefix to define regular filter criteria.. 
* `ge` - Minimum prefix length to be matched.. 
* `action` - Action.. 

