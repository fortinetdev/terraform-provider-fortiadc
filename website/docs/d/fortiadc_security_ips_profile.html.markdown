---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_ips_profile"
description: |-
  Get information on an fortiadc security ips profile
---

# Data Source: fortiadc_security_ips_profile
Use this data source to get information on an fortiadc security ips profile

## Example Usage

```hcl
 data "fortiadc_security_ips_profile" sample1 {
  ips_profile_name = "1"
}

output output1 {
  value = data.fortiadc_security_ips_profile.sample1
}
```

## Argument Reference
* `ips_profile_name` - (Required) Specify the mkey of the desired  security ips profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `ips_profile_name` - IPS profile name.
* `comments` - ips profile comments. 


