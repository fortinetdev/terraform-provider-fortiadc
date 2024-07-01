---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_dns_vdom"
description: |-
  Provides a list of fortiadc system dns per vdom
---

# Data Source: fortiadc_system_dns_vdom
Provides a list of `fortiadc system dns vdom`.

## Example Usage

```hcl
 data "fortiadc_system_dns_vdomlist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_system_dns_vdomlist.sample1.mkey_list
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

* `vdom` - (Required) Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only non-root vdom can be specified.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_system_dns_vdom`.
