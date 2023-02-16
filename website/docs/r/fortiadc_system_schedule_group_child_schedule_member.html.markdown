---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_schedule_group_child_schedule_member"
description: |-
  Configure fortiadc schedule group info.
---

# fortiadc_system_schedule_group_child_schedule_member
Configure fortiadc schedule group info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - schedule member name.
* `startdate` - start date: YYYY/MM/DD. 
* `enddate` - end date: YYYY/MM/DD. 
* `day_of_month` - monthly day. (1,31)
* `day_of_week` - weekly day. Valid values: 1:monday, 0:sunday, 3:wednesday, 2:tuesday, 5:friday, 4:thursday, 6:saturday .
* `type` - member type. Valid values: 1:monthly-recurring, 0:one-time, 3:daily-recurring, 2:weekly-recurring, 4:hourly-recurring .
* `starttime_of_startdate` - start time of start date: HH:MM. 
* `endtime_of_enddate` - end time of end date: HH:MM. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Schedule Group Child Schedule Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_schedule_group_child_schedule_member.labelname {{mkey}}
```
