---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface_child_secondary_ip_list"
description: |-
  Get information on an fortiadc system interface child secondary ip list
---

# Data Source: fortiadc_system_interface_child_secondary_ip_list
Use this data source to get information on an fortiadc system interface child secondary ip list

## Example Usage

```hcl
 data "fortiadc_system_interface_child_secondary_ip_list" sample1 {
  fadc_id = "1"
}

output output1 {
  value = data.fortiadc_system_interface_child_secondary_ip_list.sample1
}
```

## Argument Reference
* `id` - (Required) Specify the mkey of the desired  system interface child secondary ip list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `fadc_id` - secondary ip number.
* `traffic_group` - traffic group name. 
* `ip` - secondary ip address of interface. 
* `floating_ip` - floating ip address of interface. 
* `allowaccess` - Scondary Allow access with the interface. 
* `floating` - enable/disable to set floating ip. 

