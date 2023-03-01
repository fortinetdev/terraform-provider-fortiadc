---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_dos_dos_protection_profile"
description: |-
  Configure fortiadc DoS protect profile.
---

# fortiadc_security_dos_dos_protection_profile
Configure fortiadc DoS protect profile.

## Example Usage
```hcl
resource "fortiadc_security_dos_dos_protection_profile" "ddos" {
	mkey = "custom_dos"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `tcp_slow_data_limit` - tcp slowdata limit. 
* `http_conn_limit` - HTTP connection limit. 
* `tcp_conn_limit` - tcp connection access limit. 
* `http_req_limit` - HTTP request limit. 
* `http_access_limit` - HTTP access limit. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Dos Dos Protection Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_dos_dos_protection_profile.labelname {{mkey}}
```
