---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_mailserver"
description: |-
  Get information on an fortiadc system mailserver
---

# Data Source: fortiadc_system_mailserver
Use this data source to get information on an fortiadc system mailserver

## Example Usage

```hcl
 data "fortiadc_system_mailserver" sample1 {
}

output output1 {
  value = data.fortiadc_system_mailserver.sample1
}
```

## Argument Reference
* `` - (Required) Specify the mkey of the desired  system mailserver.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `username` - username for SMTP authentication. 
* `auth` - enable/disable SMTP authentication. 
* `address` - SMTP server address. 
* `security` - set security option. 
* `password` - password for SMTP authentication. 
* `port` - SMTP server port. (1,65535)

