---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_servers_child_virtual_server_list"
description: |-
  Configure fortiadc global load balance servers child virtual server list.
---

# fortiadc_global_load_balance_servers_child_virtual_server_list
Configure fortiadc global load balance servers child virtual server list.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_servers_child_virtual_server_list" "global_load_balance_servers_child_virtual_server_list" {
	pkey = "test"
  mkey = "test123"
  address_type = "ipv4"
  ip = "192.0.2.1"
  gateway = "123"
  gateway_select = ""
  description = "nothing"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The name of global load balance server.
* `mkey` - The name of configuration.
* `address_type` - IPv4 or IPv6.
* `ip` - Virtual server IPv4 address.
* `ip6` - Virtual server IPv6 address.
* `gateway` - Specify the name of a gateway.
* `gateway_select` - Select a gateway configuration object.
* `description` - Optional description to help administrators know the purpose or usage of the configuration.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Servers Child Virtual Server List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_servers_child_virtual_server_list.labelname {{mkey}}
```
