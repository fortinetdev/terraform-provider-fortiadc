---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf"
description: |-
  Get information on an fortiadc router ospf
---

# Data Source: fortiadc_router_ospf
Use this data source to get information on an fortiadc router ospf

## Example Usage

```hcl
 data "fortiadc_router_ospf" sample1 {
}

output output1 {
  value = data.fortiadc_router_ospf.sample1
}
```

## Argument Reference
* `` - (Required) Specify the mkey of the desired  router ospf.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:


* `redistribute_connected` - Enable/Disable redistribute connected. 
* `default_information_metric` - Default information metric.. (-1,16777214)
* `redistribute_static_metric_type` - Redistribute static metric type setting.. 
* `default_information_metric_type` - Default information metric type.. 
* `redistribute_connected_metric` - Redistribute connected metric setting.. (-1,16777214)
* `router_id` - router-id must be set.. 
* `default_information_originate` - Area type setting.. 
* `redistribute_static_metric` - Redistribute static metric setting.. (-1,16777214)


* `redistribute_static` - Enable/Disable redistribute static.. 
* `default_metric` - Default metric of redistribute routes.. (0,16777214)




* `redistribute_connected_metric_type` - Redistribute connected metric type setting.. 
* `distance` - Distance of the route.. (1,255)


