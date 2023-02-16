---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_md5_ospf_child_md5_member"
description: |-
  Provides a list of fortiadc router md5 ospf child md5 member
---

# Data Source: fortiadc_router_md5_ospf_child_md5_member
Provides a list of `fortiadc router md5 ospf child md5 member`.

## Example Usage

```hcl
 data "fortiadc_router_md5_ospf_child_md5_memberlist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_router_md5_ospf_child_md5_memberlist.sample1.mkey_list
}
```

## Argument Reference

* `pkey` - The parent key.
* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_router_md5_ospf_child_md5_member`.
