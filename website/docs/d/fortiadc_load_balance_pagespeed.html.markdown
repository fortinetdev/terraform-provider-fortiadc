---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed"
description: |-
  Get information on an fortiadc load balance pagespeed
---

# Data Source: fortiadc_load_balance_pagespeed
Use this data source to get information on an fortiadc load balance pagespeed

## Example Usage

```hcl
 data "fortiadc_load_balance_pagespeed" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_pagespeed.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance pagespeed.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - pagespeed name.


* `pagespeed_profile_id` - pagespeed profile name. 
* `file_cache_max_inode` - Maximum file cache inode. (1,100000)


* `file_cache_max_size` - Maximum file cache size (1M ~512M). (1,512)

