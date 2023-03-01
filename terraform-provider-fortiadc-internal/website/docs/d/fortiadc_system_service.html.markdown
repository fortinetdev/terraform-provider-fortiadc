---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_service"
description: |-
  Get information on an fortiadc system service
---

# Data Source: fortiadc_system_service
Use this data source to get information on an fortiadc system service

## Example Usage

```hcl
 data "fortiadc_system_service" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_service.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system service.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - name.
* `destination_port_max` - destination port max. (1,65535)
* `protocol` - protocol number. (1,255)
* `source_port_min` - source port min. (1,65535)
* `protocol_type` - protocol type. 
* `destination_port_min` - destination port min. (1,65535)
* `specify_source_port` - specify source port or not. 
* `source_port_max` - source port max. (1,65535)

