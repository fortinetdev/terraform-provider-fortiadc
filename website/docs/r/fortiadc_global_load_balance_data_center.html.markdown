---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_data_center"
description: |-
  Configure fortiadc global load balance data center.
---

# fortiadc_global_load_balance_data_center
Configure fortiadc global load balance data center.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_data_center" "global_load_balance_data_center" {
        mkey = "tesst"
        location = "ZZ"
        description = "ZZ"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `location` - Select a location. For example: ZZ, US, or CN65.
* `description` - Optional description to help administrators know the purpose or usage of the configuration.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Data Center can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_data_center.labelname {{mkey}}
```
