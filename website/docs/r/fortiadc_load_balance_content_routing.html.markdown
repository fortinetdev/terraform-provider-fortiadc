---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_content_routing"
description: |-
  Configure fortiadc load-balance content routing info.
---

# fortiadc_load_balance_content_routing
Configure fortiadc load-balance content routing info.

## Example Usage
```hcl
resource "fortiadc_load_balance_content_routing" "c_routing1" {
	mkey = "cr1"
	type = "l7-content-routing"
	pool = "rsp1"
	comments = "content_routing_comments"
	depends_on = [fortiadc_load_balance_pool.rsp1]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - content routing name.

* `source_pool_list` - ip pool name. 
* `schedule_list` - enable/disable schedule list. Valid values: enable/disable.
* `persistence` - persistence name. 
* `connection_pool` - connection pool name. 
* `ip` - ipv4 subnet address. 
* `packet_fwd_method` - packet forwarding method. Valid values: 0:inherit, 3:FullNAT .
* `comments` - content routing comments. 
* `connection_pool_inherit` - connection pool inherit. Valid values: enable/disable.
* `ip6` - ipv6 subnet address. 
* `type` - content-routing type. Valid values: 1:l4-content-routing, 2:l7-content-routing .
* `pool` - pool name. 
* `persistence_inherit` - persistence inherit. Valid values: enable/disable.
* `schedule_pool_list` - schedule pool name. 
* `method` - method name. 

* `method_inherit` - method inherit. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Content Routing can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_content_routing.labelname {{mkey}}
```
