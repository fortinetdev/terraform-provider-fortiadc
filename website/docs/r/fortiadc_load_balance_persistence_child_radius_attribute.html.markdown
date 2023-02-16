---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_persistence_child_radius_attribute"
description: |-
  Configure fortiadc load-balance persistence info.
---

# fortiadc_load_balance_persistence_child_radius_attribute
Configure fortiadc load-balance persistence info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `pkey` - The parent key.
* `mkey` - radius attribute id.
* `vendor_id` - vendor id, 0 means will use entire attribute as input of persistence. 
* `vendor_type` - vendor type, 0 means will use entire attribute as input of persistence. (0,255)
* `type` - radius attribute type. Valid values: 50:50-acct-multi-session-id, 60:60-chap-challenge, 61:61-nas-port-type, 62:62-port-limit, 63:63-login-lat-port, 24:24-state, 26:26-vendor-specific, 48:48-acct-output-packets, 49:49-acct-terminate-cause, 46:46-acct-session-time, 47:47-acct-input-packets, 44:44-acct-session-id, 45:45-acct-authentic, 42:42-acct-input-octets, 43:43-acct-output-octets, 40:40-acct-status-type, 41:41-acct-delay-time, 1:1-user-name, 5:5-nas-port, 4:4-nas-ip-address, 7:7-framed-protocol, 6:6-service-type, 9:9-framed-ip-netmask, 8:8-framed-ip-address, 13:13-framed-compression, 12:12-framed-mtu, 14:14-login-ip-host, 19:19-callback-number, 32:32-nas-identifier, 31:31-calling-station-id, 30:30-called-station-id, 51:51-acct-link-count, 36:36-login-lat-group, 35:35-login-lat-node, 34:34-login-lat-service, 33:33-proxy-state .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Persistence Child Radius Attribute can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_persistence_child_radius_attribute.labelname {{mkey}}
```
