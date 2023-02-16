---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_ha"
description: |-
  Get information on an fortiadc system ha
---

# Data Source: fortiadc_system_ha
Use this data source to get information on an fortiadc system ha

## Example Usage

```hcl
 data "fortiadc_system_ha" sample1 {
}

output output1 {
  value = data.fortiadc_system_ha.sample1
}
```

## Argument Reference
* `mode` - (Required) Specify the mkey of the desired  system ha.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mgmt_interface` - Management interface. 
* `local_node_id` - local node id (0-7). (0,7)

* `peer_address` - Fort unicast peer ip address. 
* `group_name` - group name. 
* `sync_l4_persistent` - sync l4 persistent. 
* `monitor_enable` - remote ip monitor. 
* `hbdev` - heartbeat port. 
* `mgmt_ip_allowaccess` - Allow access with the management ip address. 
* `config_priority` - ha config priority (0-255). (0,255)
* `hb_lost_threshold` - lost heartbeat threshold (1-60). (1,60)
* `override` - override on resurge. 
* `groupid` - group id (0-31). (0,31)
* `node_list` - node id. 
* `local_address` - For unicast local listening ip address. 
* `mgmt_ip` - IP address of management interface. 
* `sync_l4_connection` - sync l4 connection. 
* `datadev` - data port. 
* `arp_num` - num for sending (1-60). (1,60)
* `failover_hold_time` - failover hold time for remote ip (60-86400 sec). (60,86400)
* `interval` - heartbeat interval (1-20 (100*ms)). (1,20)

* `arp_interval` - interval for sending arp (1-20 sec). (1,20)
* `mgmt_status` - Management status enable/disable. 
* `interface_list` - interface list to track. 
* `sync_l7_persistent` - sync http persistent. 
* `hbtype` - heartbeat type. 
* `failover_threshold` - failover threshold for remote ip (1-64). (1,64)
* `priority` - ha priority (0-9). (0,9)

