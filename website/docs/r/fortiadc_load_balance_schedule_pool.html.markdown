---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_schedule_pool"
description: |-
  Configure fortiadc load-balance schedule pool info.
---

# fortiadc_load_balance_schedule_pool
Configure fortiadc load-balance schedule pool info.

## Example Usage
```hcl
resource "fortiadc_load_balance_schedule_pool" "sch_pool1" {
	mkey = "schedule_pool_test1"
	pool = "rsp1"
	schedule = "schedule_group_test1"
	depends_on = [fortiadc_system_schedule_group.sg1, fortiadc_load_balance_pool.rsp1]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `pool` - pool name. 
* `schedule` - schedule. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Schedule Pool can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_schedule_pool.labelname {{mkey}}
```
