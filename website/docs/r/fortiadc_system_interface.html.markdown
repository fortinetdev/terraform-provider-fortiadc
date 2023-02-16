---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface"
description: |-
  Configure fortiadc interface configuration.
---

# fortiadc_system_interface
Configure fortiadc interface configuration.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - interface name.
* `status` - interface status. Valid values: 16:up .
* `aggregate_mode` - aggregate mode. Valid values: 0:balance-rr, 3:broadcast, 2:balance-xor, 5:balance-tlb, 4:802.3ad, 6:balance-alb .
* `dhcp_gw_override` - enable/disable to override gateway. Valid values: enable/disable.
* `secondary_ip` - enable/disable to set secondary ip. Valid values: enable/disable.
* `floating_ip` - floating ip address of interface. 
* `floating` - enable/disable to set floating ip. Valid values: enable/disable.
* `speed` - speed. Valid values: 1:10half, 0:auto, 3:100half, 2:10full, 5:1000half, 4:100full, 7:10000full, 6:1000full, 9:100000full, 8:40000full .
* `redundant_member` - redundant/aggregate interface slaves. 
* `dhcp_ip_overlap` - dhcp_ip_is_overlap. 
* `dhcp_gw_distance` - distance for dhcp_gateway. (1,255)
* `wccp` - enable/disable WCCP. Valid values: enable/disable.
* `default_gw` - retrieve default gateway from pppoe server. Valid values: enable/disable.
* `trust_ip` - enable/disable to set trust ip. Valid values: enable/disable.

* `vlanid` - Vlan ID. (1,4094)


* `vdom` - vdom name. 
* `allowaccess` - Allow access with the interface. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .
* `dns_server_override` - use pppoe server DNS or not. Valid values: enable/disable.
* `type` - interface type. Valid values: 10:vdom-link, 1:vlan, 0:physical, 2:aggregate, 7:soft-switch, 6:loopback, 9:nvgre, 8:vxlan .
* `ip` - ip address of interface. 
* `traffic_group` - traffic group name. 
* `disc_retry_timeout` - pppoe discovery retry timeout. (1,255)
* `dhcp_gateway` - gateway. 
* `ip6` - ipv6 address of interface. 
* `dedicate_to_management` - dedicate to managemnet interface. Valid values: enable/disable.
* `interface` - physical interface name. 
* `password` - pppoe account password. 



* `mtu` - maximum transportation unit. (68,9000)
* `username` - pppoe account user name. 
* `floating_ip6` - floating ipv6 address of interface. 
* `mode` - static,pppoe or dhcp. Valid values: 1:dhcp, 0:static, 2:pppoe .
* `aggregate_algorithm` - aggregate interface slave selection algorithm. Valid values: 1:layer3-4, 0:layer2, 2:layer2-3 .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Interface can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_interface.labelname {{mkey}}
```
