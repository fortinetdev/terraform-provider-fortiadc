---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_dns_vdom"
description: |-
  Configure fortiadc Domain Name Server per vdom.
---

# fortiadc_system_dns_vdom
Configure fortiadc Domain Name Server per vdom.

## Example Usage
```hcl
resource "fortiadc_system_dns_vdom" "dns1" {
  vdom = "ta1"
  dns_overide = "enable"
  ip_primary = "8.8.8.9"
  ip_second = "8.8.8.7"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only non-root vdom can be specified.
* `ip_second` - Secondary DNS. 
* `ip_primary` - Primary DNS.
* `dns_overide` - Enable/Disable DNS. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Dns can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_dns_vdom.labelname SystemDns
```
