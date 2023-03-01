---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_certificate_caching"
description: |-
  Get information on an fortiadc load balance certificate caching
---

# Data Source: fortiadc_load_balance_certificate_caching
Use this data source to get information on an fortiadc load balance certificate caching

## Example Usage

```hcl
 data "fortiadc_load_balance_certificate_caching" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_certificate_caching.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance certificate caching.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - certificate-caching name.
* `max_entries` - max certificate cache entries. (1,1000000)
* `max_certificate_cache_size` - max certificate cache size. 

