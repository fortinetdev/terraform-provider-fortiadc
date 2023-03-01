---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_mysql_user_password"
description: |-
  Configure fortiadc load-balance profile help info.
---

# fortiadc_load_balance_profile_child_mysql_user_password
Configure fortiadc load-balance profile help info.

## Example Usage
```hcl
resource "fortiadc_load_balance_profile_child_mysql_user_password" "mysql_user" {
	mkey = "1"
	pkey = "custom_mysql_profile"
	username = "test"
	password = "123"
	depends_on = [fortiadc_load_balance_profile.mysql_profile]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - mysql user password id.
* `username` - username. 
* `password` - password. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Profile Child Mysql User Password can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_profile_child_mysql_user_password.labelname {{mkey}}
```
