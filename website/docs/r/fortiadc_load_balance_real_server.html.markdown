---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_real_server"
description: |-
  Configure fortiadc load-balance real server info.
---

# fortiadc_load_balance_real_server
Configure fortiadc load-balance real server info.

## Example Usage
```hcl
resource "fortiadc_load_balance_real_server" "rs1" {
	mkey = "rs1"
	server_type     = "static"
	status    = "enable"
	type = "ip"
	address = "10.10.10.10"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - real server name.
* `status` - (en|dis)able/maintain real server. Valid values: 1:enable, 0:disable, 2:maintain .
* `server_type` - real server type. Valid values: 1:static, 3:dynamic_manual, 2:dynamic_auto .
* `sdn_connector` - sdn connector. 
* `sdn_addr_private` - use private node ip. Valid values: enable/disable.
* `fqdn` - Fully Qualified Domain Name. 
* `address6` - ipv6 address of interface. 
* `instance` - instances filter. 
* `address` - ip address of interface. 
* `type` - ip or fqdn. Valid values: 1:fqdn, 0:ip .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Real Server can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_real_server.labelname {{mkey}}
```
