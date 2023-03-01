---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_caching_child_dyn_cache_list"
description: |-
  Configure fortiadc load-balance caching info.
---

# fortiadc_load_balance_caching_child_dyn_cache_list
Configure fortiadc load-balance caching info.

## Example Usage
```hcl
resource "fortiadc_load_balance_caching_child_dyn_cache_list" "caching_dyn" {
	mkey = "1"
	pkey = "lb_caching_1"
	uri = "https://cachingdyn.com"
	depends_on = [fortiadc_load_balance_caching.lb_caching]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - dynamic cache rule id.
* `invalid_uri` - invalid uri pattern. 
* `age` - cache age out time. (1,86400)
* `uri` - dynamic cache uri pattern. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Caching Child Dyn Cache List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_caching_child_dyn_cache_list.labelname {{mkey}}
```
