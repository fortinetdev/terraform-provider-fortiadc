---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_l2_exception_list_child_member"
description: |-
  Configure fortiadc load-balance layer2 exception list.
---

# fortiadc_load_balance_l2_exception_list_child_member
Configure fortiadc load-balance layer2 exception list.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - id.
* `ip_netmask` - ip netmask. 
* `comments` - comments. 
* `host_pattern` - host wildcard pattern. 
* `type` - match type: host wildcard pattern/ip netmask. Valid values: 1:match type ip netmask, 0:match type host wildcard pattern .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance L2 Exception List Child Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_l2_exception_list_child_member.labelname {{mkey}}
```
