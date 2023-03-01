---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_ca_certificateupload"
description: |-
  Upload fortiadc ca.
---

# fortiadc_ca_certificateupload
Upload fortiadc ca.

## Example Usage
```hcl
resource "fortiadc_ca_certificateupload" "ca_cert_upload_root" {
	mkey = "2"
	vdom = "global"
	cert = "/root/terr_test/host.cert"
}

```

## Argument Reference

The following arguments are supported:

* `mkey` - name.
* `cert` - cert file path.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

