---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_ha_child_remote_ip_monitor_list"
description: |-
  Configure fortiadc HA parameters configuration.
---

# fortiadc_system_ha_child_remote_ip_monitor_list
Configure fortiadc HA parameters configuration.

## Example Usage
```hcl
resource "fortiadc_system_ha_child_remote_ip_monitor_list" "remoteip_list" {
	mkey = "port1_monitor"
	remote_address = "10.106.203.254"
	source_port = "port1"
	health_check_timeout = "5"
	health_check_retry = "4"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - monitor name.
* `health_check_interval` - health check interval for remote ip (1-3600 sec). (1,3600)
* `remote_address` - remote ip address. 
* `health_check_retry` - health check retry times for remote ip (1-10). (1,10)
* `health_check_timeout` - health check timeout for remote ip (1-3600 sec). (1,3600)
* `source_port` - source interface. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Ha Child Remote Ip Monitor List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_ha_child_remote_ip_monitor_list.labelname {{mkey}}
```
