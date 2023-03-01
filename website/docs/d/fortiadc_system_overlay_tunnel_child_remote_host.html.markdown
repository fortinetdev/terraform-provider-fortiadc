---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_overlay_tunnel_child_remote_host"
description: |-
  Get information on an fortiadc system overlay tunnel child remote host
---

# Data Source: fortiadc_system_overlay_tunnel_child_remote_host
Use this data source to get information on an fortiadc system overlay tunnel child remote host

## Example Usage

```hcl
 data "fortiadc_system_overlay_tunnel_child_remote_host" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_overlay_tunnel_child_remote_host.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system overlay tunnel child remote host.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - host id.
* `vtep` - the IPv4 address of virtual tunnel endpoint. 
* `host_mac` - the MAC address of the remote host. 

