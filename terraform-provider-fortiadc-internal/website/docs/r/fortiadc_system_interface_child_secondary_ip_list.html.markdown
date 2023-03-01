---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface_child_secondary_ip_list"
description: |-
  Configure fortiadc interface configuration.
---

# fortiadc_system_interface_child_secondary_ip_list
Configure fortiadc interface configuration.

## Example Usage
```hcl
resource "fortiadc_system_interface_child_secondary_ip_list" "vlan_sec_ip" {
	vdom = "root"
	pkey = "vlan99"
	fadc_id = "1"
	allowaccess = "ping https http"
	ip = "9.9.9.10/24"
	floating = "enable"
	floating_ip = "9.9.9.11"
	traffic_group = "tg1"
	depends_on = [fortiadc_system_interface.vlan99]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `fadc_id` - secondary ip number.
* `traffic_group` - traffic group name. 
* `ip` - secondary ip address of interface. 
* `floating_ip` - floating ip address of interface. 
* `allowaccess` - Scondary Allow access with the interface. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .
* `floating` - enable/disable to set floating ip. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Interface Child Secondary Ip List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_interface_child_secondary_ip_list.labelname {{mkey}}
```
