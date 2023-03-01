---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_virtual_server"
description: |-
  Get information on an fortiadc load balance virtual server
---

# Data Source: fortiadc_load_balance_virtual_server
Use this data source to get information on an fortiadc load balance virtual server

## Example Usage

```hcl
 data "fortiadc_load_balance_virtual_server" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_virtual_server.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance virtual server.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - virtual-server name.
* `scripting_list` - virtual server scripting list. 
* `status` - enable/maintain/disable virtual server. 
* `packet_fwd_method` - packet forwarding method. 
* `warmrate` - virtual server warm up rate. (1,86400)
* `http2https` - enable/disable redirect HTTP request to HTTPS. 
* `pool` - pool name. 
* `error_page` - error-page name. 
* `fortiview` - enable/disable fortiview. 
* `ztna_profile` - ZTNA profile name. 
* `auth_policy` - authentication policy name. 
* `adfs_published_service` - AD FS published service. 
* `content_rewriting` - content rewriting. 
* `http2https_port` - HTTP service port list for redirecting HTTP to HTTPS. 
* `domain_name` - virtual server domainname on one click gslb server. 
* `protocol` - virtual server protocol numbers. 
* `error_msg` - error message. 
* `port` - virtual server service port. 
* `scripting_flag` - enable/disable virtual server scripting. 
* `comments` - virtual server comments. 
* `content_routing_list` - content routing list. 
* `wccp` - enable/disable redirect HTTP/HTTPS request to WCCP client. 
* `public_ip_type` - address type. 
* `one_click_gslb_server` - enable/disable setting one click gslb server. 
* `pagespeed` - virtual server pagespeed. 
* `connection_limit` - connection-limit. (0,100000000)
* `host_name` - virtual server hostname on one click gslb server. 
* `ssl_mirror_intf` - interface list to mirror. 
* `source_pool_list` - ip pool name. 
* `type` - virtual-server service type. 
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
* `use_azure_lb_backend_ip` - use azure lb backend ip. 
* `public_ip` - public ip address of virtual server. 
* `clone_traffic_type` - the traffic type to be cloned. 
* `address` - ip address of virtual server. 
* `alone` - enable/disable alone mode. 
* `stream_scripting_flag` - enable/disable virtual server stream scripting. 
* `client_ssl_profile` - client SSL profile (needed for https tcps smtp ftp). 
* `schedule_pool_list` - schedule pool name. 
* `ssl_mirror` - enable/disable SSL mirror. 
* `schedule_list` - enable/disable schedule list. 
* `warmup` - virtual server warm up time. (0,86400)
* `addr_type` - address type. 
* `connection_pool` - connection pool name. 
* `azure_lb_backend` - azure lb backend name. 
* `connection_rate_limit` - virtual server connection rate limit(0 - disable). (0,86400)
* `trans_rate_limit` - virtual server transactions rate limit. (0,1048567)
* `traffic_log` - enable/disable traffic log. 
* `av_profile` - antivirus profile name. 
* `public_ip6` - public ip address of virtual server. 
* `interface` - interface name. 
* `captcha_profile` - Captcha profile. 
* `waf_profile` - web application firewall profile name. 
* `content_routing` - content routing. 

