---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_clone_pool_child_pool_member"
description: |-
  Get information on an fortiadc load balance clone pool child pool member
---

# Data Source: fortiadc_load_balance_clone_pool_child_pool_member
Use this data source to get information on an fortiadc load balance clone pool child pool member

## Example Usage

```hcl
 data "fortiadc_load_balance_clone_pool_child_pool_member" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_clone_pool_child_pool_member.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance clone pool child pool member.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - server name.
* `dst_mac` - destination hardware address. 
* `src_mac` - source hardware address. 
* `address` - ip address of server. 
* `interface` - interface name. 
* `port` - server port. (0,65535)
* `mode` - mode of update. 

