---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_ha_child_remote_ip_monitor_list"
description: |-
  Get information on an fortiadc system ha child remote ip monitor list
---

# Data Source: fortiadc_system_ha_child_remote_ip_monitor_list
Use this data source to get information on an fortiadc system ha child remote ip monitor list

## Example Usage

```hcl
 data "fortiadc_system_ha_child_remote_ip_monitor_list" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_ha_child_remote_ip_monitor_list.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system ha child remote ip monitor list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - monitor name.
* `health_check_interval` - health check interval for remote ip (1-3600 sec). (1,3600)
* `remote_address` - remote ip address. 
* `health_check_retry` - health check retry times for remote ip (1-10). (1,10)
* `health_check_timeout` - health check timeout for remote ip (1-3600 sec). (1,3600)
* `source_port` - source interface. 

