---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_nat_snat"
description: |-
  Get information on an fortiadc firewall nat snat
---

# Data Source: fortiadc_firewall_nat_snat
Use this data source to get information on an fortiadc firewall nat snat

## Example Usage

```hcl
 data "fortiadc_firewall_nat_snat" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_firewall_nat_snat.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  firewall nat snat.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - The name of the snat.
* `status` - (en|dis)able snat status. 
* `out_interface` - Output interface name. 
* `from` - Source network. 
* `traffic_group` - traffic group name. 
* `trans_to_ip_end` - Trans-to-pool ip max. 
* `to` - Destination network. 
* `trans_to_type` - Trans to type ip or ip pool or no-nat. 
* `trans_to_ip_start` - Trans-to-pool ip min. 
* `trans_to_ip` - Trans-to ip. 

