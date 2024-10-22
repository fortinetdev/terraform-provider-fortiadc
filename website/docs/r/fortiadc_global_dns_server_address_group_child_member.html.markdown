---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_address_group_child_member"
description: |-
  Configure fortiadc global dns server address group child member.
---

# fortiadc_global_dns_server_address_group_child_member
Configure fortiadc global dns server address group child member.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_address_group_child_member" "dns_server_address_group_child_member" {
	pkey = "test"
	mkey = "1"
	addr_type = "ipv4"
	ip_network = "192.0.2.0/24"
	action = "include"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - Address group name.
* `mkey` - Member id.
* `addr_type` - IPv4/IPv6.
* `ip_network` - Address/mask notation to match the IP v4 address in the packet header.
* `ip6_network` - Address/mask notation to match the IP v6 address in the packet header.
* `action` - Include/Exclude—The rule logic creates an address object that excludes addresses matching the specified address block.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server Address Group Child Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_address_group_child_member.labelname GlobalDnsServerAddressGroupChildMember
```
