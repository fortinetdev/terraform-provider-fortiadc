---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_address_group"
description: |-
  Configure fortiadc global dns server address group.
---

# fortiadc_global_dns_server_address_group
Configure fortiadc global dns server address group.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_address_group" "global_dns_server_address_group" {
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
 Global Dns Server Address Group can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_address_group.labelname {{mkey}}
```
