---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_server_request_header_insert"
description: |-
  Get information on an fortiadc load balance profile child server request header insert
---

# Data Source: fortiadc_load_balance_profile_child_server_request_header_insert
Use this data source to get information on an fortiadc load balance profile child server request header insert

## Example Usage

```hcl
 data "fortiadc_load_balance_profile_child_server_request_header_insert" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_profile_child_server_request_header_insert.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance profile child server request header insert.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - header insert id.
* `type` - header insert type. 
* `string` - field:value pair string to be inserted. 

