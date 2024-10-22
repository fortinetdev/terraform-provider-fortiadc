---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_remote_dns_server_child_member"
description: |-
  Configure fortiadc global dns server remote dns server child member.
---

# fortiadc_global_dns_server_remote_dns_server_child_member
Configure fortiadc global dns server remote dns server child member.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_remote_dns_server_child_member" "dns_server_remote_dns_server_child_member" {
	pkey = "test"
	mkey = "1"
	addr_type = "ipv4"
	ip = "192.0.2.0"
	port = "53"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - Remote dn server name.
* `mkey` - Member id.
* `addr_type` - IPv4/IPv6.
* `ip` - IP v4 address of the remote DNS server.
* `ip6` - IP v6 address of the remote DNS server.
* `port` - Port number the remote server uses for DNS. The default is 53.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server Remote Dns Server Child Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_remote_dns_server_child_member.labelname GlobalDnsServerRemoteDnsServerChildMember
```
