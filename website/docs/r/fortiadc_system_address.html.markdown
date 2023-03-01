---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_address"
description: |-
  Configure fortiadc system address.
---

# fortiadc_system_address
Configure fortiadc system address.

## Example Usage
```hcl
resource "fortiadc_system_address" "address1" {
	mkey = "addr_1"
	type = "ip-netmask"
	ip_netmask = "192.0.2.0/24"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `ip_max` - IP max. 
* `ip_netmask` - IP/netmask. 
* `type` - type. Valid values: 1:ip-netmask, 2:ip-range .
* `ip_min` - IP min. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Address can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_address.labelname {{mkey}}
```
