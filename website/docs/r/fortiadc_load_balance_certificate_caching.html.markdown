---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_certificate_caching"
description: |-
  Configure fortiadc load-balance certificate-caching info.
---

# fortiadc_load_balance_certificate_caching
Configure fortiadc load-balance certificate-caching info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - certificate-caching name.
* `max_entries` - max certificate cache entries. (1,1000000)
* `max_certificate_cache_size` - max certificate cache size. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Certificate Caching can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_certificate_caching.labelname {{mkey}}
```
