---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf"
description: |-
  Provides a list of fortiadc router ospf
---

# Data Source: fortiadc_router_ospf
Provides a list of `fortiadc router ospf`.

## Example Usage

```hcl
 data "fortiadc_router_ospflist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_router_ospflist.sample1.mkey_list
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_router_ospf`.
