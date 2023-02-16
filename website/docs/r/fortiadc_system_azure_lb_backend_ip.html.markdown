---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_azure_lb_backend_ip"
description: |-
  Configure fortiadc azure lb backend ip list.
---

# fortiadc_system_azure_lb_backend_ip
Configure fortiadc azure lb backend ip list.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `ip` - backend ip. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Azure Lb Backend Ip can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_azure_lb_backend_ip.labelname {{mkey}}
```
