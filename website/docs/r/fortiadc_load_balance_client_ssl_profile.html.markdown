---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_client_ssl_profile"
description: |-
  Configure fortiadc load-balance client-ssl-profile info.
---

# fortiadc_load_balance_client_ssl_profile
Configure fortiadc load-balance client-ssl-profile info.

## Example Usage
```hcl
resource "fortiadc_load_balance_client_ssl_profile" "SSL_PROF1" {
	mkey = "SSL_PROFILE_TEST"
	local_certificate_group = "LOCAL_CERT_GROUP"
	client_certificate_verify = "cv1"
	backend_ssl_allowed_versions = "tlsv1.1 tlsv1.2 tlsv1.0"
	depends_on = [fortiadc_system_certificate_certificate_verify.cert_verify]
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - profile name.
* `ssl_allowed_versions` - SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .
* `ssl_renegotiate_size` - size (MB) of application data transmitted over SSL when ADC initiate renegotiation, set to 0 to disable ADC initiate renegotiation by size. (0,2147483647)
* `backend_certificate_verify` - backend certificate verify for SSL forward proxy. 
* `backend_customized_ssl_ciphers` - backend SSL user-customized ciphers for SSL forward proxy (concatenate multiple ciphers by :). 
* `ssl_renegotiation` - (en|dis)able renegotiation. Valid values: enable/disable.
* `customized_ssl_ciphers_flag` - (en|dis)able SSL user-customized ciphers. Valid values: enable/disable.
* `backend_customized_ssl_ciphers_flag` - (en|dis)able backend SSL user-customized ciphers for SSL forward proxy. Valid values: enable/disable.
* `ssl_secure_renegotiation` - set secure renegotiation. Valid values: 1:request, 3:require-strict, 2:require .
* `http_forward_client_certificate_header` - forward client certificate header (https profile only). 
* `supported_groups` - ssl Supported Groups. Valid values: 1:secp256r1, 0:x25519, 3:secp521r1, 2:x448, 5:ffdhe2048, 4:secp384r1, 7:ffdhe4096, 6:ffdhe3072, 9:ffdhe8192, 8:ffdhe6144 .
* `http_forward_client_certificate` - foward client certificate (https profile only). Valid values: enable/disable.
* `use_tls_tickets` - (en|dis)able TLS tickets support (When SSL session ID based persistence is used in a VS, this flag is ignored and tls ticket is disabled). Valid values: enable/disable.
* `client_sni_required` - client SNI required. Valid values: enable/disable.
* `backend_ssl_ciphers` - backend ssl ciphers including Perfect Forward Secrecy for SSL forward proxy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .
* `client_certificate_verify` - certificate verify name. 
* `backend_ssl_sni_forward` - (en|dis)able backend SSL SNI forward for SSL forward proxy. Valid values: enable/disable.
* `client_certificate_verify_mode` - certificate verify option. Valid values: 1:required, 2:optional .
* `local_certificate_group` - local certificate group name. 
* `backend_ssl_allowed_versions` - backend SSL/TLS versions allowed for SSL forward proxy. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .
* `ssl_dh_param_size` - The maximum size of the Diffie-Hellman parameters used for generating the enphemeral/temporary Diffie-Hellman key for DHE key exchange. Default value is 1024 bits.. Valid values: 1024:1024bit, 4096:4096bit, 2048:2048bit .
* `ssl_ciphers_tlsv13` - tlsv1.3 ciphers. Valid values: 1:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher .
* `backend_ciphers_tlsv13` - backend tlsv1.3 ciphers. Valid values: 1:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher .
* `ssl_renegotiate_period` - period (s|m|h) (by default in seconds) of ADC initiate renegotiation, set to 0 to disable periodical renegotiation. 
* `forward_proxy_certificate_caching` - certificate caching name. 
* `forward_proxy_local_signing_ca` - local signing CA for SSL forward proxy. 
* `customized_ssl_ciphers` - SSL user-customized ciphers (concatenate multiple ciphers by :). 
* `ssl_ciphers` - ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .
* `reject_ocsp_stapling_with_missing_nextupdate` - reject OCSP stapling with missing nextupdate. Valid values: enable/disable.
* `ssl_dynamic_record_sizing` - Adjust the SSL record size dynamically according to the state of the connection. Valid values: enable/disable.
* `rfc7919_comply` - (en|dis)able SSL RFC7919 Comply. Valid values: enable/disable.
* `backend_ssl_ocsp_stapling_support` - (en|dis)able backend SSL OCSP stapling support for SSL forward proxy. Valid values: enable/disable.
* `ssl_session_cache_flag` - (en|dis)able SSL session cache. Valid values: enable/disable.
* `forward_proxy` - ssl forward proxy mode (https profile only). Valid values: enable/disable.
* `ssl_renegotiation_interval` - minimum renegotiation interval time (s|m|h) (by default in seconds), set to -1 to disable client-initiated renegotiation. 
* `forward_proxy_intermediate_ca_group` - intermediate ca group for SSL forward proxy. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Client Ssl Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_client_ssl_profile.labelname {{mkey}}
```
