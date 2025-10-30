---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_virtual_server_content_routing_attach"
description: |-
  Attach routing list to existing content routing list by virtual-server name.
---

# fortiadc_load_balance_virtual_server_content_routing_attach
Attach routing list to existing content routing list by virtual-server name.

## Example Usage
```hcl
resource "fortiadc_load_balance_virtual_server_content_routing_attach" "cr1" {
	mkey = "vs1"
	add_routing_list = "cr1"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - virtual-server name.
* `add_routing_list` - routing list. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Virtual Server Content Routing Attach can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_virtual_server_content_routing_attach.labelname {{mkey}}
```
