---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_server_request_header_insert"
description: |-
  Configure fortiadc load-balance profile help info.
---

# fortiadc_load_balance_profile_child_server_request_header_insert
Configure fortiadc load-balance profile help info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - header insert id.
* `type` - header insert type. Valid values: 1:insert-if-not-exist, 3:append-if-not-exist, 2:insert-always, 4:append-always .
* `string` - field:value pair string to be inserted. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Profile Child Server Request Header Insert can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_profile_child_server_request_header_insert.labelname {{mkey}}
```
