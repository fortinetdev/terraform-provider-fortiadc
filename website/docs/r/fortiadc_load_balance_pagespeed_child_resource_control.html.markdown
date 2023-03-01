---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed_child_resource_control"
description: |-
  Configure fortiadc load-balance pagespeed.
---

# fortiadc_load_balance_pagespeed_child_resource_control
Configure fortiadc load-balance pagespeed.

## Example Usage
```hcl
resource "fortiadc_load_balance_pagespeed_child_resource_control" "resource_control" {
	mkey = "1"
	pkey = "pagespeed_1"
	origin_domain_pattern = "https://www.example.com"
	fetch_domain = "https://www.exmaple.com"
	rewrite_domain = "https://www.example.com"
	depends_on = [fortiadc_load_balance_pagespeed.pagespeed]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - ID.
* `rewrite_domain` - fetch domain string, such as http://www.example.com. 
* `fetch_domain` - fetch domain string, such as http://www.example.com. 
* `origin_domain_pattern` - origin domain wildcards, such as (http(s)://)*.example.com. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Pagespeed Child Resource Control can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_pagespeed_child_resource_control.labelname {{mkey}}
```
