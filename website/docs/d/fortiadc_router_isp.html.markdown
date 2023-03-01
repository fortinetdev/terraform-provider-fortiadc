---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_isp"
description: |-
  Get information on an fortiadc router isp
---

# Data Source: fortiadc_router_isp
Use this data source to get information on an fortiadc router isp

## Example Usage

```hcl
 data "fortiadc_router_isp" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_router_isp.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  router isp.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - The number of route isp.
* `destination` - The name of ISP address group. 
* `gateway` - Gateway for this route. 

