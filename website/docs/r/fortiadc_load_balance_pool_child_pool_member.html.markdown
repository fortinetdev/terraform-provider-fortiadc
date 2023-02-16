---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pool_child_pool_member"
description: |-
  Configure fortiadc load-balance pool info.
---

# fortiadc_load_balance_pool_child_pool_member
Configure fortiadc load-balance pool info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - pool member id.
* `proxy_protocol` - pool member support proxy protocol version. Valid values: 1:v1, 0:none, 2:v2 .
* `hc_status` - pool member hc status. 
* `connlimit` - connection limit. (0,1048576)
* `recover` - pool member recover time. (0,86400)
* `port` - pool member service port. (0,65535)
* `rs_profile_inherit` - real server SSL profile inherit. Valid values: enable/disable.
* `warmrate` - pool member warm up rate. (1,86400)
* `weight` - pool member weight. (1,256)
* `server_name` - pool member server name. 
* `mssql_read_only` - set as read only mssql server. Valid values: enable/disable.
* `m_health_check_relationship` - health check relationship. Valid values: 1:OR, 0:AND .
* `mysql_read_only` - set as read only mysql server. Valid values: enable/disable.
* `status` - (en|dis)able/maintain pool member. Valid values: 1:enable, 0:disable, 2:maintain .
* `real_server_id` - real server. 
* `health_check_inherit` - health check inherit. Valid values: enable/disable.
* `address6` - ipv6 address of interface. 
* `mysql_group_id` - pool member mysql group id. 
* `ssl` - (en|dis)able SSL. Valid values: enable/disable.
* `host` - host content that modify to. 
* `cookie` - pool member cookie. 
* `address` - ip address of interface. 
* `rs_profile` - real server SSL profile name. 
* `warmup` - pool member warm up time. (0,86400)
* `m_health_check` - health check control. Valid values: enable/disable.
* `connection_rate_limit` - pool member connection rate limit(0 - disable), only work for L4VS. (0,86400)
* `m_health_check_list` - health check list. 
* `modify_host` - modify the Host HTTP header. Valid values: enable/disable.
* `backup` - set as backup server. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Pool Child Pool Member can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_pool_child_pool_member.labelname {{mkey}}
```
