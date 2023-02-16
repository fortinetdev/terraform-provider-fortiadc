---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_content_routing_child_match_condition"
description: |-
  Configure fortiadc load-balance content routing info.
---

# fortiadc_load_balance_content_routing_child_match_condition
Configure fortiadc load-balance content routing info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - match condition id.
* `ignorecase` - ignore case. Valid values: enable/disable.
* `reverse` - reverse match. Valid values: enable/disable.
* `object` - match object. Valid values: 1:http-host-header, 3:http-referer-header, 2:http-request-url, 4:ip-source-address, 6:sni .
* `content` - match content. 
* `type` - match type. Valid values: 1:string, 2:regular-expression .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Content Routing Child Match Condition can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_content_routing_child_match_condition.labelname {{mkey}}
```
