---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed_child_resource_control"
description: |-
  Get information on an fortiadc load balance pagespeed child resource control
---

# Data Source: fortiadc_load_balance_pagespeed_child_resource_control
Use this data source to get information on an fortiadc load balance pagespeed child resource control

## Example Usage

```hcl
 data "fortiadc_load_balance_pagespeed_child_resource_control" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_pagespeed_child_resource_control.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance pagespeed child resource control.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - ID.
* `rewrite_domain` - fetch domain string, such as http://www.example.com. 
* `fetch_domain` - fetch domain string, such as http://www.example.com. 
* `origin_domain_pattern` - origin domain wildcards, such as (http(s)://)*.example.com. 

