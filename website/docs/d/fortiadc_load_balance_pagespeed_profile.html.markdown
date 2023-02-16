---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed_profile"
description: |-
  Get information on an fortiadc load balance pagespeed profile
---

# Data Source: fortiadc_load_balance_pagespeed_profile
Use this data source to get information on an fortiadc load balance pagespeed profile

## Example Usage

```hcl
 data "fortiadc_load_balance_pagespeed_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_pagespeed_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance pagespeed profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - profile name.
* `move_css_to_head` - move css elements above script tags. 
* `jpeg_sampling` - reduces the color sampling of jpeg images to 4:2:0. 
* `image` - enable/disable image optimizer. 
* `resize_image` - resizes images when the corresponding img tag specifies a smaller width and height. 
* `html` - enable/disable html optimizer. 
* `combine_css` - combine multiple CSS elements into one. 
* `max_combine_css_byte` - maximum combined css bytes. (1,1048576)
* `css` - enable/disable css optimizer. 

