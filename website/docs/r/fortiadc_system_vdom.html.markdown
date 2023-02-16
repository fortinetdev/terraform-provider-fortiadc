---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_vdom"
description: |-
  Configure fortiadc vdom resource configuration.
---

# fortiadc_system_vdom
Configure fortiadc vdom resource configuration.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - virtual domain name.
* `concurrentsession` - CONCURRENT SESSION. (0,1000000)
* `status` - enable/disable VDOM RLIMIT. Valid values: enable/disable.
* `vs` - Virtual Server. (0,1024)
* `l7rps` - L7 RPS. (0,1000000)
* `rs` - Real Server. (0,1024)
* `inbound` - Inbound traffic, Kbps. (0,40000000)
* `sp` - Source Pool. (0,1024)
* `l7cps` - L7 CPS. (0,1000000)
* `l4cps` - L4 CPS. (0,1000000)
* `lu` - Local User. (0,1024)
* `sslthroughput` - SSL THROUGHPUT, Kbps. (0,10000000)
* `sslcps` - SSL CPS. (0,1000000)
* `hc` - Health Check. (0,1024)
* `ug` - User Group. (0,1024)
* `outbound` - Outbound traffic, Kbps. (0,40000000)
* `ep` - Error-page. (0,1024)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Vdom can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_vdom.labelname {{mkey}}
```
