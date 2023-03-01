---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_mailserver"
description: |-
  Configure fortiadc mail server configuration.
---

# fortiadc_system_mailserver
Configure fortiadc mail server configuration.

## Example Usage
```hcl
resource "fortiadc_system_mailserver" "system_mailserver" {
	address = "10.106.202.12"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `username` - username for SMTP authentication. 
* `auth` - enable/disable SMTP authentication. Valid values: enable/disable.
* `address` - SMTP server address. 
* `security` - set security option. Valid values: 1:starttls, 2:none .
* `password` - password for SMTP authentication. 
* `port` - SMTP server port. (1,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Mailserver can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_mailserver.labelname SystemMailserver
```
