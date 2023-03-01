---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence"
description: |-
  Configure fortiadc load-balance persistence info.
---

# fortiadc_load_balance_persistence
Configure fortiadc load-balance persistence info.

## Example Usage
```hcl
resource "fortiadc_load_balance_persistence" "lb_persistence_radius" {
	mkey = "persistence_radius"
	type = "radius-attribute"
	radius_attribute_relation = "OR"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - persistence table name.
* `match_across_servers` - Match Across Servers. Valid values: enable/disable.
* `keyvalue_relation` - relation to store key value. Valid values: 1:OR, 0:AND .
* `cookie_secure` - Enable or disable Secure cookie attribute. Default disable. Valid values: enable/disable.
* `cookie_httponly` - Enable or disable HttpOnly cookie attribute. Default disable. Valid values: enable/disable.
* `sess_kw_type` - keyword type of session cookie name. Valid values: 1:PHPSESSID, 0:auto, 3:CFID+CFTOKEN, 2:JSESSIONID, 5:custom, 4:ASP.NET_SessionId .
* `avpno` - avpno. 
* `ipv6_maskbits` - IPv6 netmask bits. (1,128)

* `cookie_custom_attr` - Enable or disable custom cookie attribute. Default disable. Valid values: enable/disable.
* `type` - persistence type. Valid values: 11:radius-attribute, 10:embedded-cookie, 13:sip-call-id, 12:ssl-session-id, 15:passive-cookie, 14:rdp-cookie, 17:iso8583-bitmap, 1:source-address, 3:hash-source-address-port, 2:consistent-hash-ip, 5:hash-http-request, 4:hash-http-header, 7:persistent-cookie, 6:hash-cookie, 9:rewrite-cookie, 8:insert-cookie .
* `iso8583_bitmap_relation` - bitmap relation to trigger persistence. Valid values: 1:OR, 0:AND .
* `ipv4_maskbits` - IPv4 netmask bits. (1,32)
* `keyword` - persistence keyword. 
* `cookie_domain` - Set the Domain cookie attribute value. Default is empty. 
* `cookie_custom_attr_val` - Set the custom cookie attribute value. Default is empty. 
* `override_connection_limit` - Override Connection Limit. Valid values: enable/disable.
* `cookie_samesite` - Set the SameSite cookie attribute: nothing|None|Lax|Strict. Default is nothing. Valid values: 1:Secure should be set when using samesite none, 0:nothing, 3:strict, 2:lax .
* `timeout` - persistence entry timeout, unit: second, default 300. (1,86400)

* `radius_attribute_relation` - radius attribute relation. Valid values: 1:OR, 0:AND .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Persistence can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_persistence.labelname {{mkey}}
```
