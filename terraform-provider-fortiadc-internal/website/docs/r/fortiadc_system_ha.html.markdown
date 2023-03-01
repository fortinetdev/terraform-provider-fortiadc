---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_ha"
description: |-
  Configure fortiadc HA parameters configuration.
---

# fortiadc_system_ha
Configure fortiadc HA parameters configuration.

## Example Usage
```hcl
resource "fortiadc_system_ha" "HA" {
	mode = "active-active-vrrp"
	hbdev = "port4"
	group_name = "terraform_ha"
	groupid = "10"
	mgmt_status = "enable"
	mgmt_interface = "port6"
	mgmt_ip_allowaccess = "ping http https ssh"
	mgmt_ip = "6.6.6.10"
	monitor_enable = "enable"
	sync_l4_connection = "enable"
	sync_l4_persistent = "enable"
	sync_l7_persistent = "enable"
	hbtype = "unicast"
	local_address = "202.151.212.1"
	peer_address = "202.151.212.2"
	config_priority = "50"
	priority = "2"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mgmt_interface` - Management interface. 
* `local_node_id` - local node id (0-7). (0,7)

* `peer_address` - Fort unicast peer ip address. 
* `group_name` - group name. 
* `sync_l4_persistent` - sync l4 persistent. Valid values: enable/disable.
* `monitor_enable` - remote ip monitor. Valid values: enable/disable.
* `hbdev` - heartbeat port. 
* `mgmt_ip_allowaccess` - Allow access with the management ip address. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .
* `config_priority` - ha config priority (0-255). (0,255)
* `hb_lost_threshold` - lost heartbeat threshold (1-60). (1,60)
* `override` - override on resurge. Valid values: enable/disable.
* `groupid` - group id (0-31). (0,31)
* `node_list` - node id. Valid values: 1:1, 0:0, 3:3, 2:2, 5:5, 4:4, 7:7, 6:6 .
* `local_address` - For unicast local listening ip address. 
* `mgmt_ip` - IP address of management interface. 
* `sync_l4_connection` - sync l4 connection. Valid values: enable/disable.
* `datadev` - data port. 
* `arp_num` - num for sending (1-60). (1,60)
* `failover_hold_time` - failover hold time for remote ip (60-86400 sec). (60,86400)
* `interval` - heartbeat interval (1-20 (100*ms)). (1,20)

* `arp_interval` - interval for sending arp (1-20 sec). (1,20)
* `mode` - high availability mode. Valid values: 1:active-passive, 0:standalone, 3:active-active-vrrp, 2:active-active .
* `mgmt_status` - Management status enable/disable. Valid values: enable/disable.
* `interface_list` - interface list to track. 
* `sync_l7_persistent` - sync http persistent. Valid values: enable/disable.
* `hbtype` - heartbeat type. Valid values: 1:broadcast, 0:multicast, 2:unicast .
* `failover_threshold` - failover threshold for remote ip (1-64). (1,64)
* `priority` - ha priority (0-9). (0,9)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Ha can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_ha.labelname SystemHa
```
