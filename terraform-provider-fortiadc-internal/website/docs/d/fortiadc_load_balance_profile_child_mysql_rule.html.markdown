---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_mysql_rule"
description: |-
  Get information on an fortiadc load balance profile child mysql rule
---

# Data Source: fortiadc_load_balance_profile_child_mysql_rule
Use this data source to get information on an fortiadc load balance profile child mysql rule

## Example Usage

```hcl
 data "fortiadc_load_balance_profile_child_mysql_rule" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_profile_child_mysql_rule.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance profile child mysql rule.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - mysql rule id.
* `client_ip_list` - rule client ip list. 
* `user_list` - rule user list. 
* `db_list` - rule database list. 
* `type` - type. 
* `table_list` - rule table list. 
* `sql_list` - rule sql statement list. 

