---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_exception_child_exception_rule"
description: |-
  Configure fortiadc security waf exception child exception rule.
---

# fortiadc_security_waf_exception_child_exception_rule
Configure fortiadc security waf exception child exception rule.

## Example Usage
```hcl
resource "fortiadc_security_waf_exception_child_exception_rule" "test_security_waf_exception_child_exception_rule" {
        mkey = "1"
        pkey = "test"
        type = "URL"
        exception_host_status = "enable"
        exception_host = "/.test"
        exception_url = "/.test"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The name of waf exception.
* `mkey` - The number of configuration.
* `type` - Element type. URL, Source IP, Source IPv6, HTTP Method, HTTP Header, Cookie or Parameter.
* `exception_host_status` - Enable/disable the setting exceptions by host pattern.
* `exception_host` - Specify the matching string. Regular expressions are supported. Maximum length is 128 characters. This option appears if Exception Host Status is enabled.
* `exception_url` - Specify the matching string. Must begin with a URL path separator (/). Maximum length is 128 characters. This option appears if Element type is URL.
* `ip_netmask` - Specify the IPv4 address and netmask. This option appears if Element type is Source IP. For example: 192.0.2.5/24
* `ip6_netmask` - Specify the IPv6 address and netmask. This option appears if Element type is Source IPv6. For example: 2001:0db8:85a3::8a2e:0370:7334/64
* `exception_method` - HTTP methods. GET, POST, HEAD, TRACE, CONNECT, DELETE, PUT, PATCH, OPTIONS or OTHERS. This option appears if Element type is HTTP Method.
* `exception_name` - Specify the matching string. Regular expressions are supported. Maximum length is 1024 characters. This option appears if Element type is HTTP Header, Cookie or Parameter.
* `exception_value_check` - Enable/disable value checking for the specified element. This option appears if Element type is HTTP Header, Cookie or Parameter.
* `exception_value` - Specify the matching string. Regular expressions are supported. Maximum length is 1024 characters. This optionappears if Check Value of Specified Element is enabled.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception Child Exception Rule can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_exception_child_exception_rule.labelname {{mkey}}
```
