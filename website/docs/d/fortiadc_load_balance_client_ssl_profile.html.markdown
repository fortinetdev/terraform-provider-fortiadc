---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_client_ssl_profile"
description: |-
  Get information on an fortiadc load balance client ssl profile
---

# Data Source: fortiadc_load_balance_client_ssl_profile
Use this data source to get information on an fortiadc load balance client ssl profile

## Example Usage

```hcl
 data "fortiadc_load_balance_client_ssl_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_client_ssl_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance client ssl profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - profile name.
* `ssl_allowed_versions` - SSL/TLS versions allowed. 
* `ssl_renegotiate_size` - size (MB) of application data transmitted over SSL when ADC initiate renegotiation, set to 0 to disable ADC initiate renegotiation by size. (0,2147483647)
* `backend_certificate_verify` - backend certificate verify for SSL forward proxy. 
* `backend_customized_ssl_ciphers` - backend SSL user-customized ciphers for SSL forward proxy (concatenate multiple ciphers by :). 
* `ssl_renegotiation` - (en|dis)able renegotiation. 
* `customized_ssl_ciphers_flag` - (en|dis)able SSL user-customized ciphers. 
* `backend_customized_ssl_ciphers_flag` - (en|dis)able backend SSL user-customized ciphers for SSL forward proxy. 
* `ssl_secure_renegotiation` - set secure renegotiation. 
* `http_forward_client_certificate_header` - forward client certificate header (https profile only). 
* `supported_groups` - ssl Supported Groups. 
* `http_forward_client_certificate` - foward client certificate (https profile only). 
* `use_tls_tickets` - (en|dis)able TLS tickets support (When SSL session ID based persistence is used in a VS, this flag is ignored and tls ticket is disabled). 
* `client_sni_required` - client SNI required. 
* `backend_ssl_ciphers` - backend ssl ciphers including Perfect Forward Secrecy for SSL forward proxy. 
* `client_certificate_verify` - certificate verify name. 
* `backend_ssl_sni_forward` - (en|dis)able backend SSL SNI forward for SSL forward proxy. 
* `client_certificate_verify_mode` - certificate verify option. 
* `local_certificate_group` - local certificate group name. 
* `backend_ssl_allowed_versions` - backend SSL/TLS versions allowed for SSL forward proxy. 
* `ssl_dh_param_size` - The maximum size of the Diffie-Hellman parameters used for generating the enphemeral/temporary Diffie-Hellman key for DHE key exchange. Default value is 1024 bits.. 
* `ssl_ciphers_tlsv13` - tlsv1.3 ciphers. 
* `backend_ciphers_tlsv13` - backend tlsv1.3 ciphers. 
* `ssl_renegotiate_period` - period (s|m|h) (by default in seconds) of ADC initiate renegotiation, set to 0 to disable periodical renegotiation. 
* `forward_proxy_certificate_caching` - certificate caching name. 
* `forward_proxy_local_signing_ca` - local signing CA for SSL forward proxy. 
* `customized_ssl_ciphers` - SSL user-customized ciphers (concatenate multiple ciphers by :). 
* `ssl_ciphers` - ssl ciphers including Perfect Forward Secrecy. 
* `reject_ocsp_stapling_with_missing_nextupdate` - reject OCSP stapling with missing nextupdate. 
* `ssl_dynamic_record_sizing` - Adjust the SSL record size dynamically according to the state of the connection. 
* `rfc7919_comply` - (en|dis)able SSL RFC7919 Comply. 
* `backend_ssl_ocsp_stapling_support` - (en|dis)able backend SSL OCSP stapling support for SSL forward proxy. 
* `ssl_session_cache_flag` - (en|dis)able SSL session cache. 
* `forward_proxy` - ssl forward proxy mode (https profile only). 
* `ssl_renegotiation_interval` - minimum renegotiation interval time (s|m|h) (by default in seconds), set to -1 to disable client-initiated renegotiation. 
* `forward_proxy_intermediate_ca_group` - intermediate ca group for SSL forward proxy. 

