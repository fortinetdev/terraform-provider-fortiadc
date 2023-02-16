---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_clone_pool_child_pool_member"
description: |-
  Configure fortiadc load-balance clone pool info.
---

# fortiadc_load_balance_clone_pool_child_pool_member
Configure fortiadc load-balance clone pool info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - server name.
* `dst_mac` - destination hardware address. 
* `src_mac` - source hardware address. 
* `address` - ip address of server. 
* `interface` - interface name. 
* `port` - server port. (0,65535)
* `mode` - mode of update. Valid values: 1:mirror-dst-mac-update, 0:mirror-interface, 3:mirror-src-dst-mac-update, 2:mirror-src-mac-update, 4:mirror-ip-update .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Clone Pool Child Pool Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_clone_pool_child_pool_member.labelname {{mkey}}
```
