---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pool"
description: |-
  Configure fortiadc load-balance pool info.
---

# fortiadc_load_balance_pool
Configure fortiadc load-balance pool info.

## Example Usage
```hcl
resource "fortiadc_load_balance_pool" "rsp1" {
	mkey = "rsp1"
	pool_type = "ipv4"
	health_check = "enable"
	health_check_relationship = "AND"
	health_check_list = "LB_HLTHCK_ICMP LB_HLTHCK_HTTPS "
	type = "static"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - pool name.

* `pool_type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `health_check` - health check control. Valid values: enable/disable.
* `health_check_relationship` - health check relationship. Valid values: 1:OR, 0:AND .
* `service` - service. 
* `sdn_connector` - sdn connector. 
* `sdn_addr_private` - use private node ip. Valid values: enable/disable.

* `direct_route_ip` - direct route ip. 
* `direct_route_mode` - direct route mode. Valid values: enable/disable.
* `health_check_list` - health check list. 
* `direct_route_ip6` - direct route ip6. 
* `type` - pool type. Valid values: 1:static, 2:dynamic .
* `rs_profile` - real server SSL profile name. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Pool can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_pool.labelname {{mkey}}
```
