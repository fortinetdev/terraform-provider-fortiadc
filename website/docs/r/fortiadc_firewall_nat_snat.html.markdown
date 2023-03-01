---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_nat_snat"
description: |-
  Configure fortiadc snat.
---

# fortiadc_firewall_nat_snat
Configure fortiadc snat.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the snat.
* `status` - (en|dis)able snat status. Valid values: enable/disable.
* `out_interface` - Output interface name. 
* `from` - Source network. 
* `traffic_group` - traffic group name. 
* `trans_to_ip_end` - Trans-to-pool ip max. 
* `to` - Destination network. 
* `trans_to_type` - Trans to type ip or ip pool or no-nat. Valid values: 1:pool, 0:ip, 2:no-nat .
* `trans_to_ip_start` - Trans-to-pool ip min. 
* `trans_to_ip` - Trans-to ip. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Firewall Nat Snat can be imported using any of these accepted formats:
```
$ terraform import fortiadc_firewall_nat_snat.labelname {{mkey}}
```
