---
subcategory: "FortiADC Config"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_config_sync_list"
description: |-
  Get information on an fortiadc config sync list
---

# Data Source: fortiadc_config_sync_list
Use this data source to get information on an fortiadc config sync list

## Example Usage

```hcl
 data "fortiadc_config_sync_list" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_config_sync_list.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  config sync list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - config sync list name.
* `comment` - comment for sync-list. 
* `status` - status. 
* `filename` - filename. 
* `server_ip` - ip address of server. 
* `password` - peer admin password. 
* `type` - synchronization type. 
* `port` - port of server. (1,65535)

