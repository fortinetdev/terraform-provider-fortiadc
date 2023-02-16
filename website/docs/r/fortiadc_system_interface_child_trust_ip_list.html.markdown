---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface_child_trust_ip_list"
description: |-
  Configure fortiadc interface configuration.
---

# fortiadc_system_interface_child_trust_ip_list
Configure fortiadc interface configuration.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `name` - trust ip name.
* `ip6_max` - end trust IP6. 
* `ip6_min` - start trust IP6. 
* `ip_max` - end trust ip. 
* `ip_netmask` - trust ip/netmask. 
* `ip6_netmask` - trust ip6/netmask. 
* `type` - trust ip type. Valid values: 1:ip-netmask, 3:ip6-netmask, 2:ip-range, 4:ip6-range .
* `ip_min` - start trust IP. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Interface Child Trust Ip List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_interface_child_trust_ip_list.labelname {{mkey}}
```
