---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_caching"
description: |-
  Configure fortiadc load-balance caching info.
---

# fortiadc_load_balance_caching
Configure fortiadc load-balance caching info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - caching name.
* `max_age` - max age of cache entries. (60,86400)
* `max_cache_size` - max cache size. 
* `max_entries` - max cache entries. (1,262144)

* `max_object_size` - maximum object size. 


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Caching can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_caching.labelname {{mkey}}
```
