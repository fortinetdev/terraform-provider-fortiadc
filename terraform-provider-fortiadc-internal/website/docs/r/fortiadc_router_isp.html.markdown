---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_isp"
description: |-
  Configure fortiadc route isp.
---

# fortiadc_router_isp
Configure fortiadc route isp.

## Example Usage
```hcl
resource "fortiadc_router_isp" "router_isp" {
	mkey = "1"
	destination = "china-mobile"
	gateway = "10.0.1.254"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The number of route isp.
* `destination` - The name of ISP address group. 
* `gateway` - Gateway for this route. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Isp can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_isp.labelname {{mkey}}
```
