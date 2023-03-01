---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_web_filter_profile"
description: |-
  Get information on an fortiadc load balance web filter profile
---

# Data Source: fortiadc_load_balance_web_filter_profile
Use this data source to get information on an fortiadc load balance web filter profile

## Example Usage

```hcl
 data "fortiadc_load_balance_web_filter_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_web_filter_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance web filter profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - Web Filter Profile Name.
* `description` - Web Filter Profile Description. 



