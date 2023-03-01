---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_access_list"
description: |-
  Get information on an fortiadc router access list
---

# Data Source: fortiadc_router_access_list
Use this data source to get information on an fortiadc router access list

## Example Usage

```hcl
 data "fortiadc_router_access_list" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_router_access_list.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  router access list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - name of ipv4 access list..
* `description` - Comments.. 



