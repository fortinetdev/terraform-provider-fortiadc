---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_virtual_server"
description: |-
  Configure fortiadc load-balance virtual-server info.
---

# fortiadc_load_balance_virtual_server
Configure fortiadc load-balance virtual-server info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - virtual-server name.
* `scripting_list` - virtual server scripting list. 
* `status` - enable/maintain/disable virtual server. Valid values: 1:enable, 0:disable, 2:maintain .
* `packet_fwd_method` - packet forwarding method. Valid values: 1:direct_routing, 3:FullNAT, 2:NAT, 5:NAT64, 4:NAT46, 6:tunneling .
* `warmrate` - virtual server warm up rate. (1,86400)
* `http2https` - enable/disable redirect HTTP request to HTTPS. Valid values: enable/disable.
* `pool` - pool name. 
* `error_page` - error-page name. 
* `fortiview` - enable/disable fortiview. Valid values: enable/disable.
* `ztna_profile` - ZTNA profile name. 
* `auth_policy` - authentication policy name. 
* `adfs_published_service` - AD FS published service. 
* `content_rewriting` - content rewriting. Valid values: enable/disable.
* `http2https_port` - HTTP service port list for redirecting HTTP to HTTPS. 
* `domain_name` - virtual server domainname on one click gslb server. 
* `protocol` - virtual server protocol numbers. 
* `error_msg` - error message. 
* `port` - virtual server service port. 
* `scripting_flag` - enable/disable virtual server scripting. Valid values: enable/disable.
* `comments` - virtual server comments. 
* `content_routing_list` - content routing list. 
* `wccp` - enable/disable redirect HTTP/HTTPS request to WCCP client. Valid values: enable/disable.
* `public_ip_type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `one_click_gslb_server` - enable/disable setting one click gslb server. Valid values: enable/disable.
* `pagespeed` - virtual server pagespeed. 
* `connection_limit` - connection-limit. (0,100000000)
* `host_name` - virtual server hostname on one click gslb server. 
* `ssl_mirror_intf` - interface list to mirror. 
* `source_pool_list` - ip pool name. 
* `type` - virtual-server service type. Valid values: 1:l4-load-balance, 3:l2-load-balance, 2:l7-load-balance .
* `method` - method name. 
* `persistence` - persistence name. 
* `dos_profile` - DoS profile name. 
* `profile` - profile name. 
* `content_rewriting_list` - content rewriting list. 
* `l2_exception_list` - layer2 exception list. 
* `ips_profile` - IPS profile name. 
* `traffic_group` - traffic group name. 
* `stream_scripting_list` - virtual server scripting list. 
* `address6` - ipv6 address of virtual server. 
* `clone_pool` - clone pool name. 
* `use_azure_lb_backend_ip` - use azure lb backend ip. Valid values: enable/disable.
* `public_ip` - public ip address of virtual server. 
* `clone_traffic_type` - the traffic type to be cloned. Valid values: 1:client-side, 0:both-sides, 2:server-side .
* `address` - ip address of virtual server. 
* `alone` - enable/disable alone mode. Valid values: enable/disable.
* `stream_scripting_flag` - enable/disable virtual server stream scripting. Valid values: enable/disable.
* `client_ssl_profile` - client SSL profile (needed for https tcps smtp ftp). 
* `schedule_pool_list` - schedule pool name. 
* `ssl_mirror` - enable/disable SSL mirror. Valid values: enable/disable.
* `schedule_list` - enable/disable schedule list. Valid values: enable/disable.
* `warmup` - virtual server warm up time. (0,86400)
* `addr_type` - address type. Valid values: 1:ipv4, 2:ipv6 .
* `connection_pool` - connection pool name. 
* `azure_lb_backend` - azure lb backend name. 
* `connection_rate_limit` - virtual server connection rate limit(0 - disable). (0,86400)
* `trans_rate_limit` - virtual server transactions rate limit. (0,1048567)
* `traffic_log` - enable/disable traffic log. Valid values: enable/disable.
* `av_profile` - antivirus profile name. 
* `public_ip6` - public ip address of virtual server. 
* `interface` - interface name. 
* `captcha_profile` - Captcha profile. 
* `waf_profile` - web application firewall profile name. 
* `content_routing` - content routing. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Virtual Server can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_virtual_server.labelname {{mkey}}
```
