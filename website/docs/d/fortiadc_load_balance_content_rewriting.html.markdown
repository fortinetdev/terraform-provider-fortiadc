---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_content_rewriting"
description: |-
  Get information on an fortiadc load balance content rewriting
---

# Data Source: fortiadc_load_balance_content_rewriting
Use this data source to get information on an fortiadc load balance content rewriting

## Example Usage

```hcl
 data "fortiadc_load_balance_content_rewriting" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_content_rewriting.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance content rewriting.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - content rewriting name.

* `redirect` - redirect content. 
* `header_name` - http header name. 
* `referer_content` - referer content. 
* `referer_status` - replace referer header content. 
* `comments` - content rewriting comments. 
* `host_status` - replace host header content. 
* `url_status` - replace URL. 
* `action_type` - action type. 
* `header_value` - http header value. 
* `action` - action. 
* `url_content` - URL content. 
* `host_content` - host content. 

* `location` - location content. 

