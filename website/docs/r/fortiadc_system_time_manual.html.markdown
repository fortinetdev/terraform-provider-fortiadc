---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_time_manual"
description: |-
  Configure fortiadc set date and time.
---

# fortiadc_system_time_manual
Configure fortiadc set date and time.

## Example Usage
```hcl
resource "fortiadc_system_time_manual" "time" {
	dst = "enable"
	tz = "6"
	ntpsync = "enable"
	ntpserver = "pool.ntp.org"
	syncinterval = "30"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `tz` - time zone. (0,71)
* `dst` - enable daylight saving time. Valid values: enable/disable.
* `ntpsync` - enable/disable synchronization with NTP server. Valid values: enable/disable.
* `syncinterval` - synchronization time interval. (1,1440)
* `ntpserver` - IP address or hostname of NTP server.
* `year` - year (int)
* `month` - month (int)
* `mday` - mday (int)
* `hour` - hour (int)
* `minute` - minute (int)
* `second` - second (int)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Time Manual can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_time_manual.labelname SystemTimeManual
```
