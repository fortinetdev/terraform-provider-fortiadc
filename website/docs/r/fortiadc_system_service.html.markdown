---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_service"
description: |-
  Configure fortiadc system service.
---

# fortiadc_system_service
Configure fortiadc system service.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name.
* `destination_port_max` - destination port max. (1,65535)
* `protocol` - protocol number. (1,255)
* `source_port_min` - source port min. (1,65535)
* `protocol_type` - protocol type. Valid values: 1:ip, 3:tcp, 2:icmp, 5:tcp-and-udp, 4:udp, 6:sctp .
* `destination_port_min` - destination port min. (1,65535)
* `specify_source_port` - specify source port or not. Valid values: enable/disable.
* `source_port_max` - source port max. (1,65535)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Service can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_service.labelname {{mkey}}
```
