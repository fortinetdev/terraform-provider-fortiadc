---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_global"
description: |-
  Get information on an fortiadc system global
---

# Data Source: fortiadc_system_global
Use this data source to get information on an fortiadc system global

## Example Usage

```hcl
 data "fortiadc_system_global" sample1 {
}

output output1 {
  value = data.fortiadc_system_global.sample1
}
```

## Argument Reference
* `` - (Required) Specify the mkey of the desired  system global.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `vdom_admin` - enable/disable . 
* `telnet_port` - the port number of the telnet service. (1,65535)
* `pre_login_banner` - enable/disable pre-login-banner. 
* `https_server_cert` - appliance's local default certificate. 
* `hostname` - appliance's host name. 
* `config_sync_enable` - config sync enable. 


* `vdom_mode` - set vdom mode. 


* `sys_global_language` - Global GUI display language. 
* `gui_log` - Enable/disable menu log on Web GUI. 
* `ssh_port` - the port number of the ssh service. (1,65535)
* `config_sync_port` - config sync port. (1,65535)

* `gui_device_latitude` - latitude between (-90, 90). 
* `gui_router` - Enable/disable menu router on Web GUI. 
* `default_intermediate_ca_group` - appliance's default ssl certificate chain. 
* `admin_idle_timeout` - the idle time-out for system administration. (1,480)
* `https_port` - the port number of the https service. (1,65535)
* `https_redirect` - HTTPS redirect enable/disable. 
* `gui_system` - Enable/disable menu system on Web GUI. 
* `http_port` - the port number of the http service. (1,65535)
* `use_default_hostname` - use-default-hostname enable/disable. 

* `gui_device_longtitude` - longitude between (-180, 180). 

* `ip_second` - Secondary DNS.
* `ip_primary` - Primary DNS.
