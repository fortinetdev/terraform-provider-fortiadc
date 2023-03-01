---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_md5_ospf_child_md5_member"
description: |-
  Get information on an fortiadc router md5 ospf child md5 member
---

# Data Source: fortiadc_router_md5_ospf_child_md5_member
Use this data source to get information on an fortiadc router md5 ospf child md5 member

## Example Usage

```hcl
 data "fortiadc_router_md5_ospf_child_md5_member" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_router_md5_ospf_child_md5_member.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  router md5 ospf child md5 member.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - id..
* `key` - md5 key. 

