---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_overlay_tunnel"
description: |-
  Configure fortiadc overlay tunnel.
---

# fortiadc_system_overlay_tunnel
Configure fortiadc overlay tunnel.

## Example Usage
```hcl
resource "fortiadc_system_overlay_tunnel" "overlay_t" {
	mkey = "overlay_tunnel1"
	interface = "port7"
	dstip = "192.0.2.1"
	vni = "11"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - overlay tunnel name.

* `dstip` - the IPv4 addresses of the remote VTEPs. 
* `tunnel_type` - interface type. Valid values: 1:nvgre, 0:vxlan .
* `vni` - VXLAN Network Identifier. (1,16777215)
* `mttl` - TTL of multicast packet. (1,255)
* `ipversion` - the ip type used in destination-ip-addresses. Valid values: 1:ipv4-multicast, 0:ipv4-unicast .

* `interface` - the outgoing interface. 
* `vsid` - Virtual Subnet Identifier. (1,16777215)
* `port` - the VxLAN UDP port in destination. (1,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Overlay Tunnel can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_overlay_tunnel.labelname {{mkey}}
```
