---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_ocsp"
description: |-
  Configure fortiadc ocsp setup.
---

# fortiadc_system_certificate_ocsp
Configure fortiadc ocsp setup.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `verify_others` - verify OCSP response signing certificates using trusted certificates or CA chain.
* `reject_ocsp_response_with_missing_nextupdate` - reject OCSP response with missing nextupdate. Valid values: enable/disable.
* `max_age` - OCSP response thisupd max-age (in second, set to -1 to disable max-age check). (-1,2147483647)
* `nonce_check` - enable/disable OCSP response nonce check. Valid values: enable/disable.
* `tunnel_username` - username for web proxy authentication. 
* `url` - URL of OCSP server. 
* `tunnel_status` - enable/disable web proxy tunneling for OCSP. Valid values: enable/disable.
* `accept_trusted_root_ca` - accept trusted root CA or not. Valid values: enable/disable.
* `caching_extra_max_age_check` - OCSP caching thisupd extra max-age check (in second, set to -1 to disable it). (-1,2147483647)
* `tunnel_password` - password for web proxy authentication. 
* `caching_expire_ahead_time` - OCSP caching nextupd ahead time (in second, set to -1 to disable it). (-1,2147483647)
* `tunnel_port` - web proxy port. (1,65535)
* `timeout` - OCSP inquery timeout (in millisecond). (1,2147483647)
* `caching` - enable/disable OCSP result caching. Valid values: enable/disable.
* `ca_chain` - CA group reference for verify OCSP signing certificate. 
* `tunnel_ip` - web proxy ip address. 
* `criteria_check` - enable/disable OCSP signing certificate verify against the OCSP issuer criteria. Valid values: enable/disable.
* `remote_certificates` - remote certificates reference. 
* `mkey` - ocsp name. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Ocsp can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_ocsp.labelname {{mkey}}
```
