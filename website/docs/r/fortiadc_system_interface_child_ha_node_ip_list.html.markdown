---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface_child_ha_node_ip_list"
description: |-
  Configure fortiadc interface configuration.
---

# fortiadc_system_interface_child_ha_node_ip_list
Configure fortiadc interface configuration.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `fadc_id` - Number of IPs in list.
* `node` - Node ID of HA Node. (0,7)
* `ip` - IP. 
* `allowaccess` - Allow access with the IP. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Interface Child Ha Node Ip List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_interface_child_ha_node_ip_list.labelname {{mkey}}
```
