---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_real_server_ssl_profile"
description: |-
  Get information on an fortiadc load balance real server ssl profile
---

# Data Source: fortiadc_load_balance_real_server_ssl_profile
Use this data source to get information on an fortiadc load balance real server ssl profile

## Example Usage

```hcl
 data "fortiadc_load_balance_real_server_ssl_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_real_server_ssl_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance real server ssl profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - real-server-ssl-profile table name.
* `new_ssl_ciphers_long` - ssl ciphers including Perfect Forward Secrecy. 
* `local_cert` - Local certificate reference. 
* `server_ocsp_stapling` - Enable/disable real-server OCSP stapling support. 
* `customized_ssl_ciphers_flag` - (en|dis)able SSL user-customized ciphers. 
* `secure_renegotiation` - set backend secure renegotiation. 
* `supported_groups` - ssl Supported Groups. 
* `renegotiation` - (en|dis)able backend renegotiation. 
* `renegotiation_deny_action` - set renegotiation deny action. 
* `renegotiate_period` - period (s|m|h) (by default in seconds) of ADC initiate renegotiation, set to 0 to disable periodical renegotiation. 
* `allow_ssl_versions` - SSL/TLS versions allowed. 
* `sni_forward_flag` - (en|dis)able SSL SNI forward. 
* `ssl` - (en|dis)able SSL. 
* `sni` - Set SSL SNI value. 
* `cert_verify` - server cert verify name. 
* `customized_ssl_ciphers` - SSL customized ciphers (concatenate multiple ciphers by :). 
* `rfc7919_comply` - (en|dis)able SSL RFC7919 Comply. 
* `renegotiate_size` - size (MB) of application data transmitted over SSL when ADC initiate renegotiation, set to 0 to disable ADC initiate renegotiation by size. (0,2147483647)
* `ciphers_tlsv13` - tlsv1.3 ciphers. 
* `tls_ticket_flag` - (en|dis)able SSL session reuse using tls-ticket. 
* `session_reuse_limit` - SSL session reuse limit (0 - disable). (0,1048576)
* `session_reuse_flag` - (en|dis)able SSL session reuse. 

