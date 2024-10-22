---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_dns_server_response_rate_limit"
description: |-
  Configure fortiadc global dns server response rate limit.
---

# fortiadc_global_dns_server_response_rate_limit
Configure fortiadc global dns server response rate limit.

## Example Usage
```hcl
resource "fortiadc_global_dns_server_response_rate_limit" "global_dns_server_response_rate_limit" {
	mkey = "test"
  per_second = "1000"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `per_second` - Maximum number of responses per second. The valid range is 1-2040. The default is 1000.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Dns Server Response Rate Limit can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_dns_server_response_rate_limit.labelname {{mkey}}
```
