---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_policy"
description: |-
  Configure fortiadc global dns server policy.
---

# fortiadc_global_dns_server_policy
Configure fortiadc global dns server policy.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_policy" "global_dns_server_policy" {
	mkey = "testpolicy"
	destination_address = "testt"
	dns64_list = "testtt"
	dnssec_validate_status = "enable"
	forward = "only"
	forwarders = "testre"
	recursion_status = "enable"
	rrlimit = "testtrr"
	source_address = "test"
	zone_list = "fqdn_generate_example.com."
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `source_address` - Select an address object to specify the source match criteria.
* `destination_address` - Select an address object to specify the destination match criteria.
* `zone_list` - Select one or more zone configurations to serve DNS requests from matching traffic.
* `dns64_list` - Select one or more DNS64 configurations to use when resolving IPv6 requests.
* `recursion_status` - enables/disables recursion
* `dnssec_validate_status` - enables/disables DNSSEC validation
* `forward` - Queries the forwarder before doing its own DNS lookup or does not perform its own DNS lookups. First or Only.
* `forwarders` - If the DNS server zone has been configured as a forwarder, select the remote DNS server to which it forwards requests.
* `rrlimit` - Select a rate limit configuration object.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server Policy can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_policy.labelname {{mkey}}
```
