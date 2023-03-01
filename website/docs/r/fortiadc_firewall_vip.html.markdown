---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_vip"
description: |-
  Configure fortiadc virtual IP.
---

# fortiadc_firewall_vip
Configure fortiadc virtual IP.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the virtual IP.
* `status` - (en|dis)able static nat status. Valid values: enable/disable.
* `mappedip_max` - max mapped IP address. 
* `mappedport_min` - min mapped port . 
* `no_nat` - (en|dis)able static nat no-nat support. Valid values: enable/disable.
* `proto` - protocol type. Valid values: 17:udp, 6:tcp .
* `traffic_group` - traffic group name. 
* `extif` - external interface name. 
* `mappedip_min` - min mapped IP address. 
* `mappedport_max` - max mapped port . 
* `extport` - external port. 
* `portforward` - (en|dis)able port forwarding. Valid values: enable/disable.
* `extip` - external IP address. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Firewall Vip can be imported using any of these accepted formats:
```
$ terraform import fortiadc_firewall_vip.labelname {{mkey}}
```
