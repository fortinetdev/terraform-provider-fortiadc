---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_mysql_rule"
description: |-
  Configure fortiadc load-balance profile help info.
---

# fortiadc_load_balance_profile_child_mysql_rule
Configure fortiadc load-balance profile help info.

## Example Usage
```hcl
resource "fortiadc_load_balance_profile_child_mysql_rule" "mysql_rule" {
	mkey = "1"
	pkey = "custom_mysql_profile"
	client_ip_list = "1.1.1.1"
	db_list = "db1"
	user_list = "user1"
	type = "primary"
	table_list = "t1"
	sql_list = "sql1"
	depends_on = [fortiadc_load_balance_profile.mysql_profile]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - mysql rule id.
* `client_ip_list` - rule client ip list. 
* `user_list` - rule user list. 
* `db_list` - rule database list. 
* `type` - type. Valid values: 1:secondary, 0:primary .
* `table_list` - rule table list. 
* `sql_list` - rule sql statement list. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Profile Child Mysql Rule can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_profile_child_mysql_rule.labelname {{mkey}}
```
