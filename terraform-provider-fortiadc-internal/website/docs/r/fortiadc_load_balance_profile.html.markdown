---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile"
description: |-
  Configure fortiadc load-balance profile help info.
---

# fortiadc_load_balance_profile
Configure fortiadc load-balance profile help info.

## Example Usage
```hcl
resource "fortiadc_load_balance_profile" "mysql_profile" {
	mkey = "custom_mysql_profile"
	type = "mysql"
	mysql_mode = "sharding"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - profile name.
* `msg_encode_type` - encode type of message, default ascii. Valid values: 1:ascii, 2:binary .
* `whitelist` - geography ip allowlist.. 
* `ram_caching` - caching name. 
* `timeout_tcp_session` - timeout tcp session. (1,86400)
* `dns_authenticate_flag` - authenticate client by redirecting UDP query to TCP. Valid values: enable/disable.
* `dns_max_query_length` - maximum query length. (256,4096)
* `decompression` - decompression name. 

* `client_protocol` - transport protocol used. Valid values: 1:udp, 2:tcp .
* `intermediate_ca_group` - intermediate ca group name. 
* `vendor_id` - vendor id. 

* `dns_cache_ageout_time` - cache ageout time(seconds). (1,65535)
* `ssl` - use ssl to support https. Valid values: enable/disable.
* `timeout_ip_session` - timeout ip session. (1,86400)
* `smtp_disable_command` - smtp disable command. Valid values: 1:expn, 2:turn, 4:vrfy .
* `server_keepalive` - wait for a new server side request in a limited time. Valid values: enable/disable.

* `ssl_ciphers` - ssl ciphers. 
* `product_name` - product name. 


* `failed_server_type` - action when server side fail. Valid values: 1:drop, 2:send .

* `forward_client_certificate_header` - foward client certificate header. 
* `tune_bufsize` - buffer size of a session. (128,1048576)
* `local_cert` - local cert name. 


* `client_address` - use client address to connect to pool. Valid values: enable/disable.
* `customized_ssl_ciphers_flag` - (en|dis)able SSL user-customized ciphers. Valid values: enable/disable.
* `smtp_disable_command_status` - smtp disable command status. Valid values: enable/disable.
* `client_keepalive` - wait for a new client side request in a limited time. Valid values: enable/disable.
* `failed_server_str` - reply status string when server side fail. 

* `timeout_radius_session` - timeout of radius session. (1,3600)
* `dns_malform_query_action` - action on malform query. Valid values: 1:Drop the query., 2:Forward the query. .
* `client_sni_required` - client SNI required. Valid values: enable/disable.
* `http_keepalive_timeout` - maximum allowed time to wait for a new HTTP request to appear. (1,86400)

* `sip_dlg_timeout` - maximum allowed timeout for a dialog. (1,300)
* `client_timeout` - maximum inactivity time on the client side. (1,86400)

* `mysql_mode` - mysql mode. Valid values: 0:single-primary, 2:sharding .


* `insert_client_ip` - insert client source IP in SIP header of request. Valid values: enable/disable.
* `opt_header_length` - length of optional header before MTI, including the length-indicator, range 0-32. (0,32)
* `http_request_timeout` - maximum allowed time to wait for a complete HTTP request. (1,86400)
* `http_mode` - operating mode for http(s). Valid values: 1:The server-facing connection is closed after the response., 3:All requests and responses are processed and connection stays keep-alived on both sides., 2:Process only the first request and response. .
* `forward_client_certificate` - foward client certificate. Valid values: enable/disable.
* `dns_cache_response_type` - cache response type. Valid values: 1:all-records, 2:round-robin .
* `max_http_headers` - Max HTTP headers limit, notice: even enlarge this limit you may also meet parse failed because the buffer size limit.. (10,10000)
* `origin_realm` - origin realm. 
* `geoip_list` - geography ip block list.. 


* `server_keepalive_timeout` - maximum allowed time to wait for a new server side request to appear. (5,300)
* `client_ssl` - client ssl. Valid values: enable/disable.

* `compression` - compression name. 
* `response_half_closed_request` - If enable ADC will continue serve the request in half closed connection until response complete.. Valid values: enable/disable.
* `dynamic_auth` - enable/disable RADIUS dynamic authorization (CoA, Disconnect messages). Valid values: enable/disable.
* `server_max_size` - maximum connections on the server side. (1,30000)
* `ip_reputation` - use IP Reputation. Valid values: enable/disable.
* `type` - protocol type. Valid values: 20:explicit_http, 21:l7-tcp, 22:l7-udp, 1:tcp, 3:http, 2:udp, 5:radius, 4:ftp, 7:https, 6:tcps, 9:sip, 8:turbohttp, 11:ip, 10:rdp, 13:smtp, 12:dns, 15:rtmp, 14:rtsp, 17:diameter, 16:mysql, 19:mssql, 18:iso8583 .
* `failed_client_str` - reply status string when client side fail. 
* `timeout_udp_session` - timeout udp session. (1,86400)
* `local_cert_group` - local cert group name. 


* `http2_profile` - HTTP/2 profile. 
* `customized_ssl_ciphers` - SSL user-customized ciphers (concatenate multiple ciphers by :). 
* `starttls_active_mode` - starttls active mode. Valid values: 1:allow, 3:none, 2:require .
* `server_age` - maximum inactivity time on the server side. (1,86400)


* `smtp_domain_name` - domain name. 
* `dns_cache_size` - max cache size(M bytes). (1,100)
* `dns_cache_flag` - cache. Valid values: enable/disable.
* `new_ssl_ciphers_long` - ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 4096:Perfect Forward Secrecy cipher, 134217728:, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher .
* `http_x_forwarded_for_header` - change X-Forwarded-For header name. 
* `idle_time` - idle time. (60,3600)
* `max_header_size` - max header size. (16,65536)
* `deploy_mode` - deploy mode. Valid values: 1:Proxy mode, 2:Direct Server Return mode (DSR mode does not support TCP traffic) .
* `server_protocol` - transport protocol used for server side. Valid values: 1:udp, 2:tcp .
* `origin_host` - origin host. 
* `media_address` - media address. 
* `server_close_propagation` - server close propagation. Valid values: enable/disable.
* `timeout_tcp_session_after_fin` - timeout tcp session_after_FIN. (1,86400)
* `sip_max_size` - maximum allowed size of a message. (1,65535)

* `ssl_proxy` - ssl proxy mode. Valid values: enable/disable.
* `use_tls_tickets` - (en|dis)able SSL tls-tickets support. Valid values: enable/disable.
* `dns_cache_entry_size` - maximum cache entry size. (256,4096)

* `queue_timeout` - maximum time to wait in the queue for a connection slot to be free. (1,86400)
* `connect_timeout` - maximum time to wait for a connection attempt to a server to succeed. (1,86400)
* `opt_trailer_hex` - hex string of optional trailer, maximum length 16, i.e. 8 bytes in binary. 

* `length_indicator_shift` - bytes to shift from the beginning of payload to read length value, range 0-32. (0,32)
* `ip_reputation_redirect` - redirect url for IP Reputation. 
* `length_indicator_size` - total bytes reading to calculate length, range 0-8. (0,8)
* `http_x_forwarded_for` - insert X-Forwarded-For header to request. Valid values: enable/disable.
* `allow_ssl_versions` - SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 16:tlsv1.2 .
* `source_port` - use source port to connect to pool. Valid values: enable/disable.
* `idle_timeout` - idle timeout. (1,86400)
* `security_mode` - security mode. Valid values: 1:Use explicit(AUTH SSL/TLS) security mode., 0:No SSL security., 2:Use implicit security mode. .
* `http_send_timeout` - the timeout(in second) of HTTP send out all the buffered data. (0,86400)

* `dynamic_auth_port` - dynamic auth port. (1,65535)
* `ssl_algorithm` - ssl encryption algorithm. Valid values: 1:All algorithms., 3:High algorithms., 2:High and medium algorithms. .
* `failed_client_type` - action when client side fail. Valid values: 1:drop, 2:send .
* `stateless` - UDP stateless. Valid values: enable/disable.
* `server_timeout` - maximum inactivity time on the server side. (1,86400)
* `cert_verify` - cert verify name. 
* `geoip_redirect` - redirect url for IP geography. 
* `length_indicator_type` - encode type of length indicator, default binary. Valid values: 1:binary, 3:decimal-str, 2:BCD, 4:hex-str .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_profile.labelname {{mkey}}
```
