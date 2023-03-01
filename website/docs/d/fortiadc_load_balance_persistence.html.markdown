---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence"
description: |-
  Get information on an fortiadc load balance persistence
---

# Data Source: fortiadc_load_balance_persistence
Use this data source to get information on an fortiadc load balance persistence

## Example Usage

```hcl
 data "fortiadc_load_balance_persistence" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_persistence.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance persistence.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - persistence table name.
* `match_across_servers` - Match Across Servers. 
* `keyvalue_relation` - relation to store key value. 
* `cookie_secure` - Enable or disable Secure cookie attribute. Default disable. 
* `cookie_httponly` - Enable or disable HttpOnly cookie attribute. Default disable. 
* `sess_kw_type` - keyword type of session cookie name. 
* `avpno` - avpno. 
* `ipv6_maskbits` - IPv6 netmask bits. (1,128)

* `cookie_custom_attr` - Enable or disable custom cookie attribute. Default disable. 
* `type` - persistence type. 
* `iso8583_bitmap_relation` - bitmap relation to trigger persistence. 
* `ipv4_maskbits` - IPv4 netmask bits. (1,32)
* `keyword` - persistence keyword. 
* `cookie_domain` - Set the Domain cookie attribute value. Default is empty. 
* `cookie_custom_attr_val` - Set the custom cookie attribute value. Default is empty. 
* `override_connection_limit` - Override Connection Limit. 
* `cookie_samesite` - Set the SameSite cookie attribute: nothing|None|Lax|Strict. Default is nothing. 
* `timeout` - persistence entry timeout, unit: second, default 300. (1,86400)

* `radius_attribute_relation` - radius attribute relation. 

