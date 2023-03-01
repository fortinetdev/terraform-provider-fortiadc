---
subcategory: "FortiADC Vdom"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_vdom"
description: |-
  Provides a list of fortiadc vdom
---

# Data Source: fortiadc_vdom
Provides a list of `fortiadc vdom`.

## Example Usage

```hcl
 data "fortiadc_vdomlist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_vdomlist.sample1.mkey_list
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_vdom`.
