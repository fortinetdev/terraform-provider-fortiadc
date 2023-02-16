---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_isp_addr_child_address"
description: |-
  Configure fortiadc isp address group.
---

# fortiadc_system_isp_addr_child_address
Configure fortiadc isp address group.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - entry id.
* `province` - province name. 
* `ip_netmask` - IP netmask. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Isp Addr Child Address can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_isp_addr_child_address.labelname {{mkey}}
```
