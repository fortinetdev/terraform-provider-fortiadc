---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_mysql_sharding"
description: |-
  Configure fortiadc load-balance profile help info.
---

# fortiadc_load_balance_profile_child_mysql_sharding
Configure fortiadc load-balance profile help info.

## Example Usage
```hcl
resource "fortiadc_load_balance_profile_child_mysql_sharding" "mysql_sharding" {
	mkey = "1"
	pkey = "custom_mysql_profile"
	database = "1"
	table = "1"
	group_list = "1"
	key = "1"
	depends_on = [fortiadc_load_balance_profile.mysql_profile]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - mysql sharding id.
* `database` - sharding database. 
* `group_list` - sharding group list, for range type, such as: 0:1-10 1:11-20, for hash type, such as: 0 1. 
* `key` - sharding table column. 
* `table` - sharding table. 
* `type` - type. Valid values: 1:hash, 0:range .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Profile Child Mysql Sharding can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_profile_child_mysql_sharding.labelname {{mkey}}
```
