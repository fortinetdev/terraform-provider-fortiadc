---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_virtual_server_pool_child_member"
description: |-
  Configure fortiadc global load balance virtual server pool child member.
---

# fortiadc_global_load_balance_virtual_server_pool_child_member
Configure fortiadc global load balance virtual server pool child member.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_virtual_server_pool_child_member" "load_balance_virtual_server_pool_child_member" {
        pkey = "testqwerqw"
        mkey = "1"
        server = "test"
        server_member_name = "testmeber"
        ttl = "100"
        weight = "100"
        backup = "enable"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The name of virtual server pool.
* `mkey` - The number of configuration.
* `server` - Select a GLB Servers configuration object.
* `server_member_name` - Select the name of the virtual server that is in the servers virtual server list configuration.
* `ttl` -
* `weight` - Assigns relative preference among members—higher values are more preferred and are assigned connections more frequently. The default is 1. The valid range is 1-255.
* `backup` - Enable/disable to designate the member as a backup. Backup members are inactive until all main members are down.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Virtual Server Pool Child Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_virtual_server_pool_child_member.labelname GlobalLoadBalanceVirtualServerPoolChildMember
```
