---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_servers"
description: |-
  Configure fortiadc global load balance servers.
---

# fortiadc_global_load_balance_servers
Configure fortiadc global load balance servers.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_servers" "load_balance_servers" {
        mkey = "test"
        server_type = "FortiADC-SLB"
        auth_type = "auth_verify"
        password = "test"
        user_defined_cert = "enable"
        cert = "SSLPROXY_LOCAL_CA"
        address_type = "ipv4"
        ip = "192.0.2.1"
        port = "5858"
        data_center = "dc_test"
        auto_sync = "enable"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `server_type` - FortiADC SLB, Generic Host or SDN Connector.
* `auth_type` - No password, TCP MD5SIG or Auth Verify.
* `password` - Enter the password to authenticate the key. This option is available if Auth Type is FortiADC SLB or TCP MD5SIG.
* `user_defined_cert` - Enable/disable to use a self-defined certificate for authentication. This option is available if Auth Type is FortiADC SLB.
* `cert` - Select the local certificate object to use for the GSLB server. This option is available if Auth Type is FortiADC SLB and User Defined Certificate is enabled.
* `address_type` - IPv4 or IPv6. This option is available if Auth Type is FortiADC SLB.
* `ip` - Specify the IPv4 address for the FortiADC management interface. This option is available if Auth Type is FortiADC SLB.
* `ip6` - Specify the IPv6 address for the FortiADC management interface. This option is available if Auth Type is FortiADC SLB.
* `port` - Specify the port. Default: 5858 Range: 1-65535. his option is available if Auth Type is FortiADC SLB.
* `data_center` - Select a data center configuration object.
* `auto_sync` - enable/disable automatic synchronization with the remote server
* `health_check_ctrl` - Enable/disable health checks for the virtual server list. This option is available if Server Type is Generic Host or SDN Connector.
* `health_check_relationship` - AND or OR. This option is available if Server Type is Generic Host and Health Check Control is enabled.
* `health_check_list` - Select one or more health check configuration objects. This option is available if Server Type is Generic Host and Health Check Control is enabled.
* `sdn_connector` - Select the SDN Connector to synchronize to the GSLB server. This option is available if Server Type is SDN Connector.
* `use_sdn_private_ip` - Enable to use the SDN Private IP address. This option is available if Server Type is SDN Connector.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Servers can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_servers.labelname {{mkey}}
```
