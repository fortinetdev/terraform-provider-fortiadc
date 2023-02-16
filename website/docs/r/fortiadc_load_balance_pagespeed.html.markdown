---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed"
description: |-
  Configure fortiadc load-balance pagespeed.
---

# fortiadc_load_balance_pagespeed
Configure fortiadc load-balance pagespeed.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - pagespeed name.


* `pagespeed_profile_id` - pagespeed profile name. 
* `file_cache_max_inode` - Maximum file cache inode. (1,100000)


* `file_cache_max_size` - Maximum file cache size (1M ~512M). (1,512)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Pagespeed can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_pagespeed.labelname {{mkey}}
```
