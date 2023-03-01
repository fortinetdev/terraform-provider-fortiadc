---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_mysql_sharding"
description: |-
  Get information on an fortiadc load balance profile child mysql sharding
---

# Data Source: fortiadc_load_balance_profile_child_mysql_sharding
Use this data source to get information on an fortiadc load balance profile child mysql sharding

## Example Usage

```hcl
 data "fortiadc_load_balance_profile_child_mysql_sharding" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_profile_child_mysql_sharding.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance profile child mysql sharding.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - mysql sharding id.
* `database` - sharding database. 
* `group_list` - sharding group list, for range type, such as: 0:1-10 1:11-20, for hash type, such as: 0 1. 
* `key` - sharding table column. 
* `table` - sharding table. 
* `type` - type. 

