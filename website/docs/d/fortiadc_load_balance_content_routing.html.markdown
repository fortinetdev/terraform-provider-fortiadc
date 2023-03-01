---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_content_routing"
description: |-
  Get information on an fortiadc load balance content routing
---

# Data Source: fortiadc_load_balance_content_routing
Use this data source to get information on an fortiadc load balance content routing

## Example Usage

```hcl
 data "fortiadc_load_balance_content_routing" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_content_routing.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance content routing.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - content routing name.

* `source_pool_list` - ip pool name. 
* `schedule_list` - enable/disable schedule list. 
* `persistence` - persistence name. 
* `connection_pool` - connection pool name. 
* `ip` - ipv4 subnet address. 
* `packet_fwd_method` - packet forwarding method. 
* `comments` - content routing comments. 
* `connection_pool_inherit` - connection pool inherit. 
* `ip6` - ipv6 subnet address. 
* `type` - content-routing type. 
* `pool` - pool name. 
* `persistence_inherit` - persistence inherit. 
* `schedule_pool_list` - schedule pool name. 
* `method` - method name. 

* `method_inherit` - method inherit. 

