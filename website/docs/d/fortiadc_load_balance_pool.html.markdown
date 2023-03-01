---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pool"
description: |-
  Get information on an fortiadc load balance pool
---

# Data Source: fortiadc_load_balance_pool
Use this data source to get information on an fortiadc load balance pool

## Example Usage

```hcl
 data "fortiadc_load_balance_pool" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_pool.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance pool.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - pool name.

* `pool_type` - address type. 
* `health_check` - health check control. 
* `health_check_relationship` - health check relationship. 
* `service` - service. 
* `sdn_connector` - sdn connector. 
* `sdn_addr_private` - use private node ip. 

* `direct_route_ip` - direct route ip. 
* `direct_route_mode` - direct route mode. 
* `health_check_list` - health check list. 
* `direct_route_ip6` - direct route ip6. 
* `type` - pool type. 
* `rs_profile` - real server SSL profile name. 

