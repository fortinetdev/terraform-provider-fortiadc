---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_schedule_group_child_schedule_member"
description: |-
  Get information on an fortiadc system schedule group child schedule member
---

# Data Source: fortiadc_system_schedule_group_child_schedule_member
Use this data source to get information on an fortiadc system schedule group child schedule member

## Example Usage

```hcl
 data "fortiadc_system_schedule_group_child_schedule_member" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_schedule_group_child_schedule_member.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system schedule group child schedule member.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - schedule member name.
* `startdate` - start date: YYYY/MM/DD. 
* `enddate` - end date: YYYY/MM/DD. 
* `day_of_month` - monthly day. (1,31)
* `day_of_week` - weekly day. 
* `type` - member type. 
* `starttime_of_startdate` - start time of start date: HH:MM. 
* `endtime_of_enddate` - end time of end date: HH:MM. 

