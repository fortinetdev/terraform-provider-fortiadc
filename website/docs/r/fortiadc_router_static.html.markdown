---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_static"
description: |-
  Configure fortiadc route configuration.
---

# fortiadc_router_static
Configure fortiadc route configuration.

## Example Usage
```hcl
resource "fortiadc_router_static" "router_test" {
	mkey = "2"
	dest = "110.2.2.122/32"
	gw = "2.2.2.2"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The number of the route.
* `gw` - Gateway for this route. 
* `distance` - distance for this route. (1,255)
* `dest` - Destination for this route. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Static can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_static.labelname {{mkey}}
```
