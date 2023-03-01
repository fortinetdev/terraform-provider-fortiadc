---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_ippool_child_node_member"
description: |-
  Get information on an fortiadc load balance ippool child node member
---

# Data Source: fortiadc_load_balance_ippool_child_node_member
Use this data source to get information on an fortiadc load balance ippool child node member

## Example Usage

```hcl
 data "fortiadc_load_balance_ippool_child_node_member" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_ippool_child_node_member.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance ippool child node member.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - pool name.
* `ip6_max` - IPv6 max. 
* `ip6_min` - IPv6 min. 
* `ha_node_num` - ha-node id. (0,7)
* `ip_min` - IP min. 
* `pool_type` - address type. 
* `ip_max` - IP max. 
* `interface` - interface name. 

