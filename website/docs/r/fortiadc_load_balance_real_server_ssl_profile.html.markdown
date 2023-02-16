---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_real_server_ssl_profile"
description: |-
  Configure fortiadc load-balance real-server-ssl-profile info.
---

# fortiadc_load_balance_real_server_ssl_profile
Configure fortiadc load-balance real-server-ssl-profile info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - real-server-ssl-profile table name.
* `new_ssl_ciphers_long` - ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .
* `local_cert` - Local certificate reference. 
* `server_ocsp_stapling` - Enable/disable real-server OCSP stapling support. Valid values: enable/disable.
* `customized_ssl_ciphers_flag` - (en|dis)able SSL user-customized ciphers. Valid values: enable/disable.
* `secure_renegotiation` - set backend secure renegotiation. Valid values: 1:request, 3:require-strict, 2:require .
* `supported_groups` - ssl Supported Groups. Valid values: 1:secp256r1, 0:x25519, 3:secp521r1, 2:x448, 5:ffdhe2048, 4:secp384r1, 7:ffdhe4096, 6:ffdhe3072, 9:ffdhe8192, 8:ffdhe6144 .
* `renegotiation` - (en|dis)able backend renegotiation. Valid values: enable/disable.
* `renegotiation_deny_action` - set renegotiation deny action. Valid values: 1:ignore, 2:terminate .
* `renegotiate_period` - period (s|m|h) (by default in seconds) of ADC initiate renegotiation, set to 0 to disable periodical renegotiation. 
* `allow_ssl_versions` - SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .
* `sni_forward_flag` - (en|dis)able SSL SNI forward. Valid values: enable/disable.
* `ssl` - (en|dis)able SSL. Valid values: enable/disable.
* `sni` - Set SSL SNI value. 
* `cert_verify` - server cert verify name. 
* `customized_ssl_ciphers` - SSL customized ciphers (concatenate multiple ciphers by :). 
* `rfc7919_comply` - (en|dis)able SSL RFC7919 Comply. Valid values: enable/disable.
* `renegotiate_size` - size (MB) of application data transmitted over SSL when ADC initiate renegotiation, set to 0 to disable ADC initiate renegotiation by size. (0,2147483647)
* `ciphers_tlsv13` - tlsv1.3 ciphers. Valid values: 1:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher .
* `tls_ticket_flag` - (en|dis)able SSL session reuse using tls-ticket. Valid values: enable/disable.
* `session_reuse_limit` - SSL session reuse limit (0 - disable). (0,1048576)
* `session_reuse_flag` - (en|dis)able SSL session reuse. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Real Server Ssl Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_real_server_ssl_profile.labelname {{mkey}}
```
