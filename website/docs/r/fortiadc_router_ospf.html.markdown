---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf"
description: |-
  Configure fortiadc Configure OSPF..
---

# fortiadc_router_ospf
Configure fortiadc Configure OSPF..

## Example Usage
```hcl
resource "fortiadc_router_ospf" "ospf" {
	router_id = "10.106.203.254"
	redistribute_static = "enable"
	distance = "100"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `redistribute_connected` - Enable/Disable redistribute connected. Valid values: enable/disable.
* `default_information_metric` - Default information metric.. (-1,16777214)
* `redistribute_static_metric_type` - Redistribute static metric type setting.. Valid values: 1:1, 2:2 .
* `default_information_metric_type` - Default information metric type.. Valid values: 1:1, 2:2 .
* `redistribute_connected_metric` - Redistribute connected metric setting.. (-1,16777214)
* `router_id` - router-id must be set.. 
* `default_information_originate` - Area type setting.. Valid values: 1:enable, 0:disable, 2:always .
* `redistribute_static_metric` - Redistribute static metric setting.. (-1,16777214)


* `redistribute_static` - Enable/Disable redistribute static.. Valid values: enable/disable.
* `default_metric` - Default metric of redistribute routes.. (0,16777214)




* `redistribute_connected_metric_type` - Redistribute connected metric type setting.. Valid values: 1:1, 2:2 .
* `distance` - Distance of the route.. (1,255)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Ospf can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_ospf.labelname RouterOspf
```
