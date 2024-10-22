---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_general"
description: |-
  Configure fortiadc global dns server general.
---

# fortiadc_global_dns_server_general
Configure fortiadc global dns server general.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_general" "global_dns_server_general" {
	gds_status = "enable"
	recurision_status = "enable"
	dnssec_validate_status = "enable"
	minimal_responses = "enable"
	ipv4_accessed_status = "enable"
	ipv6_accessed_status = "enable"
	listen_on_all_interface = "disable"
	interface_list = "port5"
	dohs_status = "enable"
	dohs_port = "443"
	dohs_interface_list = "port3"
	doh_status = "enable"
	doh_port = "80"
	doh_interface_list = "port4"
	dot_status = "enable"
	dot_port = "853"
	dot_interface_list = "port2"
	certificate = "SSLPROXY_LOCAL_CA"
	forward = "only"
	use_system_dns_server = "disable"
	forwarders = "test"
	traffic_log = "enable"
	response_ratelimit = "testtrr"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `gds_status` - enable/disable the Global DNS Configuration.
* `recurision_status` - enable/disable recursion. If enabled, the DNS server attempts to do all the work required to answer the query. If not enabled, the server returns a referral response when it does not already know the answer.
* `dnssec_validate_status` - enable/disables DNSSEC validation.
* `minimal_responses` - enables/disables Minimal Responses to hide the Authority Section and Additional Section of DNS queries.
* `ipv4_accessed_status` - enables/disables listening for DNS requests on the interface IPv4 address.
* `ipv6_accessed_status` - enables/disables listening for DNS requests on the interface IPv6 address.
* `listen_on_all_interface` - enables listening on all interfaces
* `interface_list` - If not listening to all interfaces, select one or more ports to listen on. This option is available if Listen on All Interface is disabled.
* `dohs_status` - enables/disables DNS over HTTPS to encrypt DNS queries using the HTTPS protocol.
* `dohs_port` - Specify the port to listen on DNS over HTTPS. Default: 443 Range: 1-65535. This option is available if DNS over HTTPS is enabled.
* `dohs_interface_list` - Select the interface(s) to listen on for DNS over HTTPS. This option is available if DNS over HTTPS is enabled.
* `doh_status` - enables/disables DNS over HTTP to encrypt DNS queries using the HTTP protocol.
* `doh_port` - Specify the port to listen on DNS over HTTP. Default: 80 Range: 1-65535. This option is available if DNS over HTTP is enabled.
* `doh_interface_list` - Select the interface(s) to listen on for DNS over HTTP. This option is available if DNS over HTTP is enabled.
* `dot_status` - enables/disables DNS over TLS to encrypt DNS queries using the TLS protocol.
* `dot_port` - Specify the port to listen on DNS over TLS. Default: 853 Range: 1-65535. This option is available if DNS over TLS is enabled.
* `dot_interface_list` - Select the interface(s) to listen on for DNS queries for DNS over TLS. This option is available if DNS over TLS is enabled.
* `certificate` - Select the certificate object to apply for DNS over HTTPS or DNS over TLS. This option is available if DNS over HTTPS or DNS over TLS is enabled.
* `forward` - Queries the forwarder before doing its own DNS lookup or does not perform its own DNS lookups. First or Only.
* `use_system_dns_server` - Forwards DNS requests to the system DNS server instead of the forwarders list.
* `forwarders` - Specify the forward(s).
* `traffic_log` - enables/disables traffic log.
* `response_ratelimit` - Selects a rate limit configuration object.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server General can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_general.labelname GlobalDnsServerGeneral
```
