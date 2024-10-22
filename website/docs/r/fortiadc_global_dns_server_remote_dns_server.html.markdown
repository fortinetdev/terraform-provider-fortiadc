---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_remote_dns_server"
description: |-
  Configure fortiadc global dns server remote dns server.
---

# fortiadc_global_dns_server_remote_dns_server
Configure fortiadc global dns server remote dns server.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_remote_dns_server" "global_dns_server_remote_dns_server" {
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server Remote Dns Server can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_remote_dns_server.labelname {{mkey}}
```
