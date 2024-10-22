---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_zone"
description: |-
  Configure fortiadc global dns server zone.
---

# fortiadc_global_dns_server_zone
Configure fortiadc global dns server zone.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_zone" "global_dns_server_zone" {
	mkey = "tesst"
	type = "primary"
	domain_name = "example13.com."
	dns_policy = ""
	dnssec_status = "enable"
	dnssec_algorithm= "RSASHA512"
	dnssec_keysize = "4096"
	dsset_info_list = ""
	ttl = "86400"
	serial = "10004"
	negative_ttl = "3600"
	responsible_mail = "admin"
	primary_server_name = "admin.example.com."
	primary_server_ip = "192.0.2.1"
	primary_server_ip6 = "2001:db8::1"
	forward_host_enable = "enable"
	forward = "first"
	forwarders = "testre"
	notify_status = "enable"
	also_notify_list = "192.0.2.1 2001:db8::1"
	allow_transfer = "testt"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `type` - Primary - The configuration contains the “primary” copy of data for the zone and is the authoritative server for it. Forward - The configuration allows you to apply DNS forwarding on a per-domain basis, overriding the forwarding settings in the “general” configuration. FQDN Generate - The zone and its resource record is generated from the global load balancing framework.
* `domain_name` - The domain name must end with a period. For example: example.com.
* `dns_policy` - Select the DNS policy you want the zone to use.
* `dnssec_status` - enable/disable DNSSEC
* `dnssec_algorithm` - Select the cryptographic algorithm to use for authenticating DNSSEC. This option is available if DNSSEC is enabled.
* `dnssec_keysize` - Select the key size (number of bits) for the encryption algorithm. This option is available if DNSSEC is enabled. Prior to FortiADC 7.4.0, the DNSSEC key size only supported 512 bit.
* `dsset_info_list` - ISelect a DSSET configuration object. This option is available if DNSSEC is enabled.
* `ttl` - The default is 86,400. The valid range is 0 to 2,147,483,647.
* `serial` - Set the serial number of the zone. Default 10004. Range 1-4294967295.
* `negative_ttl` - The default is 3600 seconds. The valid range is 0 to 2,147,483,647.
* `responsible_mail` - Username of the person responsible for this zone, such as hostprimary.example.com..
* `primary_server_name` - Sets the server name in the SOA record.
* `primary_server_ip` - The IP address of the primary server.
* `primary_server_ip6` - The IP6 address of the primary server.
* `forward_host_enable` - enable/disable forward host
* `forward` - Queries the forwarder before doing its own DNS lookup or does not perform its own DNS lookups. First or Only.
* `forwarders` - Select a remote server configuration object.
* `notify_status` - Enable/Disable notify status. The IP in "also notify IP list" will be notified only when Notify Status is enabled.
* `also_notify_list` - Set a list of IP addresses that will be notified if Notify Status is enabled.
* `allow_transfer` - Defines a list of IP addresses that are allowed to transfer the DNS zone information.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server Zone can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_zone.labelname {{mkey}}
```
