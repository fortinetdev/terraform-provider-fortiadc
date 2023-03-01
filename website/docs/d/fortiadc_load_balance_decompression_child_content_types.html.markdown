---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_decompression_child_content_types"
description: |-
  Get information on an fortiadc load balance decompression child content types
---

# Data Source: fortiadc_load_balance_decompression_child_content_types
Use this data source to get information on an fortiadc load balance decompression child content types

## Example Usage

```hcl
 data "fortiadc_load_balance_decompression_child_content_types" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_decompression_child_content_types.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance decompression child content types.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - content type id.
* `content_type` - content type string. 
* `custom_content_type` - custom content type. 

