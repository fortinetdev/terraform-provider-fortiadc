---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_overlay_tunnel"
description: |-
  Get information on an fortiadc system overlay tunnel
---

# Data Source: fortiadc_system_overlay_tunnel
Use this data source to get information on an fortiadc system overlay tunnel

## Example Usage

```hcl
 data "fortiadc_system_overlay_tunnel" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_overlay_tunnel.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system overlay tunnel.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - overlay tunnel name.

* `dstip` - the IPv4 addresses of the remote VTEPs. 
* `tunnel_type` - interface type. 
* `vni` - VXLAN Network Identifier. (1,16777215)
* `mttl` - TTL of multicast packet. (1,255)
* `ipversion` - the ip type used in destination-ip-addresses. 

* `interface` - the outgoing interface. 
* `vsid` - Virtual Subnet Identifier. (1,16777215)
* `port` - the VxLAN UDP port in destination. (1,65535)

