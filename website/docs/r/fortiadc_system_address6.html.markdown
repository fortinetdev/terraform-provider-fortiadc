---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_address6"
description: |-
  Configure fortiadc system IPv6 address.
---

# fortiadc_system_address6
Configure fortiadc system IPv6 address.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `ip6_network` - IPv6/Prefix. 
* `ip6_max` - IP6 max. 
* `ip6_min` - IPv6 min. 
* `type` - type. Valid values: 1:ip6-network, 2:ip6-range .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Address6 can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_address6.labelname {{mkey}}
```
