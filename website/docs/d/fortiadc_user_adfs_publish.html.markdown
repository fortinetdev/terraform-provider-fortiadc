---
subcategory: "FortiADC User"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_user_adfs_publish"
description: |-
  Get information on an fortiadc user adfs publish
---

# Data Source: fortiadc_user_adfs_publish
Use this data source to get information on an fortiadc user adfs publish

## Example Usage

```hcl
 data "fortiadc_user_adfs_publish" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_user_adfs_publish.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  user adfs publish.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - adfs-proxy name.
* `status` - enable/disable adfs publish. 
* `backend_server_url` - Backend Server URL. 
* `external_url` - External URL. 
* `preauth` - preauth method. 
* `relying_party` - Relying Party. 
* `proxy` - AD FS proxy. 

