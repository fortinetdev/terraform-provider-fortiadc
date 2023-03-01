---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_ztna_profile"
description: |-
  Get information on an fortiadc security ztna profile
---

# Data Source: fortiadc_security_ztna_profile
Use this data source to get information on an fortiadc security ztna profile

## Example Usage

```hcl
 data "fortiadc_security_ztna_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_security_ztna_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  security ztna profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - profile name.

* `log_status` - log status. 

* `comments` - comments. 

