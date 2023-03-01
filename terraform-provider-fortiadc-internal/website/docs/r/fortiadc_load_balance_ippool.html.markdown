---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_ippool"
description: |-
  Configure fortiadc load-balance ip pool info.
---

# fortiadc_load_balance_ippool
Configure fortiadc load-balance ip pool info.

## Example Usage
```hcl
resource "fortiadc_load_balance_ippool" "ippool1" {
	mkey = "ippool1"
	pool_type = "ipv4"
	interface = "port2"
	ip_start = "192.168.2.101"
	ip_end = "192.168.2.104"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - pool name.
* `pool_type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `ip_end` - IP max. 

* `ip6_end` - IPv6 max. 
* `ip_start` - IP min. 
* `interface` - interface name. 
* `ip6_start` - IPv6 min. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Ippool can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_ippool.labelname {{mkey}}
```
