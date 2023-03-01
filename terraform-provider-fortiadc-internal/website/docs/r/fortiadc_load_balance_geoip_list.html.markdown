---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_geoip_list"
description: |-
  Configure fortiadc Configure geography ip block list..
---

# fortiadc_load_balance_geoip_list
Configure fortiadc Configure geography ip block list..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name of block list..
* `status` - status. Valid values: enable/disable.
* `log` - enable/disable blocking log.. Valid values: enable/disable.


* `action` - default action. Valid values: 1:pass, 2:deny, 5:send-403-forbidden, 4:redirect .
* `severity` - High, Medium or Low. Valid values: 1:low, 3:high, 2:medium .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Geoip List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_geoip_list.labelname {{mkey}}
```
