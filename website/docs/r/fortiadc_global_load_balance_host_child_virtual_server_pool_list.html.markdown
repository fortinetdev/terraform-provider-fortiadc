---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_host_child_virtual_server_pool_list"
description: |-
  Configure fortiadc global load balance host child virtual server pool list.
---

# fortiadc_global_load_balance_host_child_virtual_server_pool_list
Configure fortiadc global load balance host child virtual server pool list.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_host_child_virtual_server_pool_list" "host_child_virtual_server_pool_list" {
        pkey = "testqw"
        mkey = "test1234567"
        virtual_server_pool = "testqw"
        weight = "100"
        topology = "any"
        isp = "china-unicom"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The name of load balance host.
* `mkey` - The name of configuration.
* `virtual_server_pool` - Select a virtual server pool.
* `weight` - Assign a weight. Valid values range from 1 to 255.
* `topology` - Select a topology.
* `isp` - Select an ISP. Example: china-mobile or china-unicom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Host Child Virtual Server Pool List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_host_child_virtual_server_pool_list.labelname {{mkey}}
```
