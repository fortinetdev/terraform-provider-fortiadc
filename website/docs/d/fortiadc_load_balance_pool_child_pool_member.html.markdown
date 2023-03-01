---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pool_child_pool_member"
description: |-
  Get information on an fortiadc load balance pool child pool member
---

# Data Source: fortiadc_load_balance_pool_child_pool_member
Use this data source to get information on an fortiadc load balance pool child pool member

## Example Usage

```hcl
 data "fortiadc_load_balance_pool_child_pool_member" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_pool_child_pool_member.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance pool child pool member.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `pkey` - The parent key.
* `mkey` - pool member id.
* `proxy_protocol` - pool member support proxy protocol version. 
* `hc_status` - pool member hc status. 
* `connlimit` - connection limit. (0,1048576)
* `recover` - pool member recover time. (0,86400)
* `port` - pool member service port. (0,65535)
* `rs_profile_inherit` - real server SSL profile inherit. 
* `warmrate` - pool member warm up rate. (1,86400)
* `weight` - pool member weight. (1,256)
* `server_name` - pool member server name. 
* `mssql_read_only` - set as read only mssql server. 
* `m_health_check_relationship` - health check relationship. 
* `mysql_read_only` - set as read only mysql server. 
* `status` - (en|dis)able/maintain pool member. 
* `real_server_id` - real server. 
* `health_check_inherit` - health check inherit. 
* `address6` - ipv6 address of interface. 
* `mysql_group_id` - pool member mysql group id. 
* `ssl` - (en|dis)able SSL. 
* `host` - host content that modify to. 
* `cookie` - pool member cookie. 
* `address` - ip address of interface. 
* `rs_profile` - real server SSL profile name. 
* `warmup` - pool member warm up time. (0,86400)
* `m_health_check` - health check control. 
* `connection_rate_limit` - pool member connection rate limit(0 - disable), only work for L4VS. (0,86400)
* `m_health_check_list` - health check list. 
* `modify_host` - modify the Host HTTP header. 
* `backup` - set as backup server. 

