---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed_child_page_control"
description: |-
  Configure fortiadc load-balance pagespeed.
---

# fortiadc_load_balance_pagespeed_child_page_control
Configure fortiadc load-balance pagespeed.

## Example Usage
```hcl
resource "fortiadc_load_balance_pagespeed_child_page_control" "page_control" {
	mkey = "1"
	pkey = "pagespeed_1"
	type = "include"
	uri = "https://pagecontrol.com/test.html"
	depends_on = [fortiadc_load_balance_pagespeed.pagespeed]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - ID.
* `type` - include/exclue type. Valid values: 1:include, 2:exclude .
* `uri` - full uri wildcards, such as http(s)://*example.com/*/htmls/*.html. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Pagespeed Child Page Control can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_pagespeed_child_page_control.labelname {{mkey}}
```
