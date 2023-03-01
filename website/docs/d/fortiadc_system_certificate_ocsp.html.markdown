---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_ocsp"
description: |-
  Get information on an fortiadc system certificate ocsp
---

# Data Source: fortiadc_system_certificate_ocsp
Use this data source to get information on an fortiadc system certificate ocsp

## Example Usage

```hcl
 data "fortiadc_system_certificate_ocsp" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_certificate_ocsp.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system certificate ocsp.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - ocsp name.
* `reject_ocsp_response_with_missing_nextupdate` - reject OCSP response with missing nextupdate. 
* `max_age` - OCSP response thisupd max-age (in second, set to -1 to disable max-age check). (-1,2147483647)
* `nonce_check` - enable/disable OCSP response nonce check. 
* `tunnel_username` - username for web proxy authentication. 
* `url` - URL of OCSP server. 
* `tunnel_status` - enable/disable web proxy tunneling for OCSP. 
* `verify_others` - verify OCSP response signing certificates using trusted certificates or CA chain. 
* `accept_trusted_root_ca` - accept trusted root CA or not. 
* `caching_extra_max_age_check` - OCSP caching thisupd extra max-age check (in second, set to -1 to disable it). (-1,2147483647)
* `tunnel_password` - password for web proxy authentication. 
* `caching_expire_ahead_time` - OCSP caching nextupd ahead time (in second, set to -1 to disable it). (-1,2147483647)
* `tunnel_port` - web proxy port. (1,65535)
* `timeout` - OCSP inquery timeout (in millisecond). (1,2147483647)
* `caching` - enable/disable OCSP result caching. 
* `ca_chain` - CA group reference for verify OCSP signing certificate. 
* `tunnel_ip` - web proxy ip address. 
* `criteria_check` - enable/disable OCSP signing certificate verify against the OCSP issuer criteria. 
* `remote_certificates` - remote certificates reference. 

