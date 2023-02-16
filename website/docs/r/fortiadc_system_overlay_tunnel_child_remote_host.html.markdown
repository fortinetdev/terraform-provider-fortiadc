---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_overlay_tunnel_child_remote_host"
description: |-
  Configure fortiadc overlay tunnel.
---

# fortiadc_system_overlay_tunnel_child_remote_host
Configure fortiadc overlay tunnel.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - host id.
* `vtep` - the IPv4 address of virtual tunnel endpoint. 
* `host_mac` - the MAC address of the remote host. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Overlay Tunnel Child Remote Host can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_overlay_tunnel_child_remote_host.labelname {{mkey}}
```
