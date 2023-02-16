---
subcategory: "FortiADC User"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_user_adfs_publish"
description: |-
  Configure fortiadc adfs publish info.
---

# fortiadc_user_adfs_publish
Configure fortiadc adfs publish info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - adfs-proxy name.
* `status` - enable/disable adfs publish. Valid values: 1:enable, 0:disable .
* `backend_server_url` - Backend Server URL. 
* `external_url` - External URL. 
* `preauth` - preauth method. Valid values: 1:ADFS, 0:Pass-through .
* `relying_party` - Relying Party. 
* `proxy` - AD FS proxy. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 User Adfs Publish can be imported using any of these accepted formats:
```
$ terraform import fortiadc_user_adfs_publish.labelname {{mkey}}
```
