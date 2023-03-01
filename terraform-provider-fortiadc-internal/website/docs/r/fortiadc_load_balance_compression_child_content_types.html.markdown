---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_compression_child_content_types"
description: |-
  Configure fortiadc load-balance compression info.
---

# fortiadc_load_balance_compression_child_content_types
Configure fortiadc load-balance compression info.

## Example Usage
```hcl
resource "fortiadc_load_balance_compression_child_content_types" "comp_content_type" {
	mkey = "1"
	pkey = "comp1"
	content_type = "text/html"
	depends_on = [fortiadc_load_balance_compression.comp]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - content type id.
* `content_type` - content type string. Valid values: 10:custom, 1:application/soap+xml, 3:text/html, 2:application/xml, 5:text/css, 4:text/plain, 7:application/javascript, 6:application/x-javascript, 9:text/xml, 8:text/javascript .
* `custom_content_type` - custom content type. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Compression Child Content Types can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_compression_child_content_types.labelname {{mkey}}
```
