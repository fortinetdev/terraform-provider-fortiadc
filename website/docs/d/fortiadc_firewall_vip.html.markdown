---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_vip"
description: |-
  Get information on an fortiadc firewall vip
---

# Data Source: fortiadc_firewall_vip
Use this data source to get information on an fortiadc firewall vip

## Example Usage

```hcl
 data "fortiadc_firewall_vip" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_firewall_vip.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  firewall vip.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - The name of the virtual IP.
* `status` - (en|dis)able static nat status. 
* `mappedip_max` - max mapped IP address. 
* `mappedport_min` - min mapped port . 
* `no_nat` - (en|dis)able static nat no-nat support. 
* `proto` - protocol type. 
* `traffic_group` - traffic group name. 
* `extif` - external interface name. 
* `mappedip_min` - min mapped IP address. 
* `mappedport_max` - max mapped port . 
* `extport` - external port. 
* `portforward` - (en|dis)able port forwarding. 
* `extip` - external IP address. 

