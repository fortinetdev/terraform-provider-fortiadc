---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed_profile"
description: |-
  Configure fortiadc pagespeed profile.
---

# fortiadc_load_balance_pagespeed_profile
Configure fortiadc pagespeed profile.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - profile name.
* `move_css_to_head` - move css elements above script tags. Valid values: enable/disable.
* `jpeg_sampling` - reduces the color sampling of jpeg images to 4:2:0. Valid values: enable/disable.
* `image` - enable/disable image optimizer. Valid values: enable/disable.
* `resize_image` - resizes images when the corresponding img tag specifies a smaller width and height. Valid values: enable/disable.
* `html` - enable/disable html optimizer. Valid values: enable/disable.
* `combine_css` - combine multiple CSS elements into one. Valid values: enable/disable.
* `max_combine_css_byte` - maximum combined css bytes. (1,1048576)
* `css` - enable/disable css optimizer. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Pagespeed Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_pagespeed_profile.labelname {{mkey}}
```
