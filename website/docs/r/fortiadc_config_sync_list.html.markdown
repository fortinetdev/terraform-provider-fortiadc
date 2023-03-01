---
subcategory: "FortiADC Config"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_config_sync_list"
description: |-
  Configure fortiadc config sync list.
---

# fortiadc_config_sync_list
Configure fortiadc config sync list.

## Example Usage
```hcl
resource "fortiadc_config_sync_list" "conf_sync_list" {
	mkey = "sync_list1"
	port = "995"
	password = "fortinet"
	server_ip = "10.106.203.91"
	type = "system networking load-balance log"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - config sync list name.
* `comment` - comment for sync-list. 
* `status` - status. 
* `filename` - filename. 
* `server_ip` - ip address of server. 
* `password` - peer admin password. 
* `type` - synchronization type. Valid values: 8192:share-resource, 1024:security, 32:log, 1:system, 4:networking, 128:link-load-balance, 8:load-balance, 512:global-load-balance, 16384:user .
* `port` - port of server. (1,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Config Sync List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_config_sync_list.labelname {{mkey}}
```
