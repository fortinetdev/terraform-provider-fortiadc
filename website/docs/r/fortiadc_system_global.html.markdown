---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_global"
description: |-
  Configure fortiadc system global configuration.
---

# fortiadc_system_global
Configure fortiadc system global configuration.

## Example Usage
```hcl
resource "fortiadc_system_global" "system_settings" {
	config_sync_enable = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `vdom_admin` - enable/disable . Valid values: enable/disable.
* `telnet_port` - the port number of the telnet service. (1,65535)
* `pre_login_banner` - enable/disable pre-login-banner. Valid values: enable/disable.
* `https_server_cert` - appliance's local default certificate. 
* `hostname` - appliance's host name. 
* `config_sync_enable` - config sync enable. Valid values: enable/disable.
* `vdom_mode` - set vdom mode. Valid values: 1:share-network, 0:independent-network .
* `sys_global_language` - Global GUI display language. Valid values: 1:english, 3:japanese, 2:chinese-simplified, 5:spanish, 7:portuguese, 6:chinese-traditional .
* `gui_log` - Enable/disable menu log on Web GUI. Valid values: enable/disable.
* `ssh_port` - the port number of the ssh service. (1,65535)
* `config_sync_port` - config sync port. (1,65535)
* `gui_device_latitude` - latitude between (-90, 90). 
* `gui_router` - Enable/disable menu router on Web GUI. Valid values: enable/disable.
* `default_intermediate_ca_group` - appliance's default ssl certificate chain. 
* `admin_idle_timeout` - the idle time-out for system administration. (1,480)
* `https_port` - the port number of the https service. (1,65535)
* `https_redirect` - HTTPS redirect enable/disable. Valid values: enable/disable.
* `gui_system` - Enable/disable menu system on Web GUI. Valid values: enable/disable.
* `http_port` - the port number of the http service. (1,65535)
* `use_default_hostname` - use-default-hostname enable/disable. Valid values: enable/disable.
* `gui_device_longtitude` - longitude between (-180, 180).
* `ip_second` - Secondary DNS.
* `ip_primary` - Primary DNS. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Global can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_global.labelname SystemGlobal
```
