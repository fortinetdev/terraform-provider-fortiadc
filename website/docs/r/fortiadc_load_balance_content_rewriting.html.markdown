---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_content_rewriting"
description: |-
  Configure fortiadc load-balance content rewriting info.
---

# fortiadc_load_balance_content_rewriting
Configure fortiadc load-balance content rewriting info.

## Example Usage
```hcl
resource "fortiadc_load_balance_content_rewriting" "cr1" {
	mkey = "cr1"
	action = "redirect"
	redirect = "redirect_test"
	comments = "terraform_test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - content rewriting name.

* `redirect` - redirect content. 
* `header_name` - http header name. 
* `referer_content` - referer content. 
* `referer_status` - replace referer header content. Valid values: enable/disable.
* `comments` - content rewriting comments. 
* `host_status` - replace host header content. Valid values: enable/disable.
* `url_status` - replace URL. Valid values: enable/disable.
* `action_type` - action type. Valid values: 1:request, 2:response .
* `header_value` - http header value. 
* `action` - action. Valid values: 1:rewrite_http_header, 3:send-403-forbidden, 2:redirect, 5:add_http_header, 4:rewrite_http_location, 6:delete_http_header .
* `url_content` - URL content. 
* `host_content` - host content. 

* `location` - location content. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Content Rewriting can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_content_rewriting.labelname {{mkey}}
```
