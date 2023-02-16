---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_time_manual"
description: |-
  Get information on an fortiadc system time manual
---

# Data Source: fortiadc_system_time_manual
Use this data source to get information on an fortiadc system time manual

## Example Usage

```hcl
 data "fortiadc_system_time_manual" sample1 {
}

output output1 {
  value = data.fortiadc_system_time_manual.sample1
}
```

## Argument Reference
* `dst` - (Required) Specify the mkey of the desired  system time manual.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `tz` - time zone. (0,71)

