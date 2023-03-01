---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile"
description: |-
  Get information on an fortiadc load balance profile
---

# Data Source: fortiadc_load_balance_profile
Use this data source to get information on an fortiadc load balance profile

## Example Usage

```hcl
 data "fortiadc_load_balance_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - profile name.
* `msg_encode_type` - encode type of message, default ascii. 
* `whitelist` - geography ip allowlist.. 
* `ram_caching` - caching name. 
* `timeout_tcp_session` - timeout tcp session. (1,86400)
* `dns_authenticate_flag` - authenticate client by redirecting UDP query to TCP. 
* `dns_max_query_length` - maximum query length. (256,4096)
* `decompression` - decompression name. 

* `client_protocol` - transport protocol used. 
* `intermediate_ca_group` - intermediate ca group name. 
* `vendor_id` - vendor id. 

* `dns_cache_ageout_time` - cache ageout time(seconds). (1,65535)
* `ssl` - use ssl to support https. 
* `timeout_ip_session` - timeout ip session. (1,86400)
* `smtp_disable_command` - smtp disable command. 
* `server_keepalive` - wait for a new server side request in a limited time. 

* `ssl_ciphers` - ssl ciphers. 
* `product_name` - product name. 


* `failed_server_type` - action when server side fail. 

* `forward_client_certificate_header` - foward client certificate header. 
* `tune_bufsize` - buffer size of a session. (128,1048576)
* `local_cert` - local cert name. 


* `client_address` - use client address to connect to pool. 
* `customized_ssl_ciphers_flag` - (en|dis)able SSL user-customized ciphers. 
* `smtp_disable_command_status` - smtp disable command status. 
* `client_keepalive` - wait for a new client side request in a limited time. 
* `failed_server_str` - reply status string when server side fail. 

* `timeout_radius_session` - timeout of radius session. (1,3600)
* `dns_malform_query_action` - action on malform query. 
* `client_sni_required` - client SNI required. 
* `http_keepalive_timeout` - maximum allowed time to wait for a new HTTP request to appear. (1,86400)

* `sip_dlg_timeout` - maximum allowed timeout for a dialog. (1,300)
* `client_timeout` - maximum inactivity time on the client side. (1,86400)

* `mysql_mode` - mysql mode. 


* `insert_client_ip` - insert client source IP in SIP header of request. 
* `opt_header_length` - length of optional header before MTI, including the length-indicator, range 0-32. (0,32)
* `http_request_timeout` - maximum allowed time to wait for a complete HTTP request. (1,86400)
* `http_mode` - operating mode for http(s). 
* `forward_client_certificate` - foward client certificate. 
* `dns_cache_response_type` - cache response type. 
* `max_http_headers` - Max HTTP headers limit, notice: even enlarge this limit you may also meet parse failed because the buffer size limit.. (10,10000)
* `origin_realm` - origin realm. 
* `geoip_list` - geography ip block list.. 


* `server_keepalive_timeout` - maximum allowed time to wait for a new server side request to appear. (5,300)
* `client_ssl` - client ssl. 

* `compression` - compression name. 
* `response_half_closed_request` - If enable ADC will continue serve the request in half closed connection until response complete.. 
* `dynamic_auth` - enable/disable RADIUS dynamic authorization (CoA, Disconnect messages). 
* `server_max_size` - maximum connections on the server side. (1,30000)
* `ip_reputation` - use IP Reputation. 
* `type` - protocol type. 
* `failed_client_str` - reply status string when client side fail. 
* `timeout_udp_session` - timeout udp session. (1,86400)
* `local_cert_group` - local cert group name. 


* `http2_profile` - HTTP/2 profile. 
* `customized_ssl_ciphers` - SSL user-customized ciphers (concatenate multiple ciphers by :). 
* `starttls_active_mode` - starttls active mode. 
* `server_age` - maximum inactivity time on the server side. (1,86400)


* `smtp_domain_name` - domain name. 
* `dns_cache_size` - max cache size(M bytes). (1,100)
* `dns_cache_flag` - cache. 
* `new_ssl_ciphers_long` - ssl ciphers including Perfect Forward Secrecy. 
* `http_x_forwarded_for_header` - change X-Forwarded-For header name. 
* `idle_time` - idle time. (60,3600)
* `max_header_size` - max header size. (16,65536)
* `deploy_mode` - deploy mode. 
* `server_protocol` - transport protocol used for server side. 
* `origin_host` - origin host. 
* `media_address` - media address. 
* `server_close_propagation` - server close propagation. 
* `timeout_tcp_session_after_fin` - timeout tcp session_after_FIN. (1,86400)
* `sip_max_size` - maximum allowed size of a message. (1,65535)

* `ssl_proxy` - ssl proxy mode. 
* `use_tls_tickets` - (en|dis)able SSL tls-tickets support. 
* `dns_cache_entry_size` - maximum cache entry size. (256,4096)

* `queue_timeout` - maximum time to wait in the queue for a connection slot to be free. (1,86400)
* `connect_timeout` - maximum time to wait for a connection attempt to a server to succeed. (1,86400)
* `opt_trailer_hex` - hex string of optional trailer, maximum length 16, i.e. 8 bytes in binary. 

* `length_indicator_shift` - bytes to shift from the beginning of payload to read length value, range 0-32. (0,32)
* `ip_reputation_redirect` - redirect url for IP Reputation. 
* `length_indicator_size` - total bytes reading to calculate length, range 0-8. (0,8)
* `http_x_forwarded_for` - insert X-Forwarded-For header to request. 
* `allow_ssl_versions` - SSL/TLS versions allowed. 
* `source_port` - use source port to connect to pool. 
* `idle_timeout` - idle timeout. (1,86400)
* `security_mode` - security mode. 
* `http_send_timeout` - the timeout(in second) of HTTP send out all the buffered data. (0,86400)

* `dynamic_auth_port` - dynamic auth port. (1,65535)
* `ssl_algorithm` - ssl encryption algorithm. 
* `failed_client_type` - action when client side fail. 
* `stateless` - UDP stateless. 
* `server_timeout` - maximum inactivity time on the server side. (1,86400)
* `cert_verify` - cert verify name. 
* `geoip_redirect` - redirect url for IP geography. 
* `length_indicator_type` - encode type of length indicator, default binary. 

