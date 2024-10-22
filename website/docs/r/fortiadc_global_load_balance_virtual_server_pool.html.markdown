---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_virtual_server_pool"
description: |-
  Configure fortiadc global load balance virtual server pool.
---

# fortiadc_global_load_balance_virtual_server_pool
Configure fortiadc global load balance virtual server pool.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_virtual_server_pool" "load_balance_virtual_server_pool" {
        mkey = "test1234567"
        preferred = "GEO-ISP"
        alternate = "RTT"
        load_balance_method = "wrr"
        check_server_status = "enable"
        check_virtual_server_existent = "enable"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `preferred` - No preference, Geo, Geo-ISP, RTT, Least, Connection-Limit, Bytes-Per-Second or Server-Performance.
* `alternate` - Same as Preferred.
* `load_balance_method` - Weighted Round Robin.
* `check_server_status` - Enable/disable polling of the local FortiADC SLB.
* `check_virtual_server_existent` - Enable/disable checks on whether the status of the virtual servers in the virtual server list is known.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Virtual Server Pool can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_virtual_server_pool.labelname {{mkey}}
```
