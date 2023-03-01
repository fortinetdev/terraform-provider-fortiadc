---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_local_cert_group"
description: |-
  Get information on an fortiadc system certificate local cert group
---

# Data Source: fortiadc_system_certificate_local_cert_group
Use this data source to get information on an fortiadc system certificate local cert group

## Example Usage

```hcl
 data "fortiadc_system_certificate_local_cert_group" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_certificate_local_cert_group.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system certificate local cert group.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - Local certificate group name.


