---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_ospf_child_ospf_interface"
description: |-
  Get information on an fortiadc router ospf child ospf interface
---

# Data Source: fortiadc_router_ospf_child_ospf_interface
Use this data source to get information on an fortiadc router ospf child ospf interface

## Example Usage

```hcl
 data "fortiadc_router_ospf_child_ospf_interface" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_router_ospf_child_ospf_interface.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  router ospf child ospf interface.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - Interface entry name..
* `transmit_delay` - Transmit delay.. (1,65535)
* `retransmit_interval` - Retransmit interval.. (3,65535)
* `hello_interval` - Hello interval.. (1,65535)
* `text` - Authentication text key.. 
* `mtu_ignore` - Enable/disable ignore MTU.. 
* `priority` - Priority.. (0,255)
* `authentication` - Authentication type.. 
* `dead_interval` - Dead interval.. (1,65535)
* `interface` - Configuration interface name.. 
* `md5` - Authentication md5 key list.. 
* `network_type` - Network type.. 
* `cost` - Cost of the interface.. (0,65535)

