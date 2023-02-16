---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_isp_addr_child_address"
description: |-
  Get information on an fortiadc system isp addr child address
---

# Data Source: fortiadc_system_isp_addr_child_address
Use this data source to get information on an fortiadc system isp addr child address

## Example Usage

```hcl
 data "fortiadc_system_isp_addr_child_address" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_isp_addr_child_address.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system isp addr child address.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - entry id.
* `province` - province name. 
* `ip_netmask` - IP netmask. 

