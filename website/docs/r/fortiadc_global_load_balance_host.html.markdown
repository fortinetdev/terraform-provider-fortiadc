---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_host"
description: |-
  Configure fortiadc global load balance host.
---

# fortiadc_global_load_balance_host
Configure fortiadc global load balance host.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_host" "global_load_balance_host" {
        mkey = "test1234567"
        host_name = "www"
        domain_name = "test.com."
        dns_policy = "DEFAULT_DNS_POLICY"
        respond_single_record = "enable"
        persistence = "enable"
        load_balance_method = "weight"
        default_feedback_ip = "192.1.2.1"
        default_feedback_ip6 = "2002:db8::1"
        fortiview = "enable"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `host_name` - The hostname part of the FQDN, such as www.
* `domain_name` - The domain name must end with a period. For example: example.com.
* `dns_policy` - Select the DNS policy you want the host to use.
* `respond_single_record` - enable/disable an option to send a single record in response to a query
* `persistence` - enable/disable the persistence table
* `load_balance_method` - Weight, DNS Query Origin or Global Availability.
* `default_feedback_ip` - Specify an IP address to return in the DNS answer if no virtual servers are available.
* `default_feedback_ip6` - Specify an IPv6 address to return in the DNS answer if no virtual servers are available.
* `fortiview` - enable/disable the fortiview option


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Host can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_host.labelname {{mkey}}
```
