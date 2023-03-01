---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_method"
description: |-
  Configure fortiadc load-balance method info.
---

# fortiadc_load_balance_method
Configure fortiadc load-balance method info.

## Example Usage
```hcl
resource "fortiadc_load_balance_method" "lb_method" {
	mkey = "method1"
	type = "dest-ip-hash"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - method table name.
* `type` - method type. Valid values: 1:round-robin, 3:fastest-response, 2:least-connection, 5:full-uri-hash, 4:uri-hash, 7:host-domain-hash, 6:host-hash, 9:dynamic-load, 8:dest-ip-hash .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Method can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_method.labelname {{mkey}}
```
