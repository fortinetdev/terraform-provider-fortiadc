---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_auto_backup"
description: |-
  Configure fortiadc Auto configuration backup.
---

# fortiadc_system_auto_backup
Configure fortiadc Auto configuration backup.

## Example Usage
```hcl
resource "fortiadc_system_auto_backup" "autobackup" {
	storage = "disk"
	scheduled_backup_status = "enable"
	scheduled_backup_day = "Thursday"
	scheduled_backup_frequency = "every week"
	scheduled_backup_time = "23:15"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `username` - username for local server. 
* `overwrite_config` - overwrite previous configuration when backup space is full. Valid values: enable/disable.
* `config_password` - password for config. 
* `scheduled_backup_status` - enable/disable auto backup. Valid values: enable/disable.
* `scheduled_backup_frequency` - scheduled backup frequency. Valid values: 1:every day, 0:time interval, 2:every week .
* `storage` - save configuration to ADC disk or local server. Valid values: 1:sftp, 0:disk .
* `scheduled_backup_time` - hour and minute, hh: 0-23, mm: {00|15|30|45}. 
* `address` - local server ip address. 
* `scheduled_backup_day` - day of the week: Sunday to Saturday. Valid values: 1:Monday, 0:Sunday, 3:Wednesday, 2:Tuesday, 5:Friday, 4:Thursday, 6:Saturday .
* `folder` - local directory. 
* `password` - password for local server. 
* `port` - local server port. (1,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Auto Backup can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_auto_backup.labelname SystemAutoBackup
```
