---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_local_certificateupload"
description: |-
  Upload fortiadc intermediate ca.
---

# fortiadc_certificate_intmed_caupload
Upload fortiadc local cert.

## Example Usage
```hcl
resource "fortiadc_system_certificate_local_upload" "local_upload" {
	mkey = "cert_test"
	type = "CertKey"
	upload = "text"
	vdom = "root"
	cert = "/root/terr_test/host.cert"
	key = "/root/terr_test/host.key"
}

resource "fortiadc_system_certificate_local_upload" "local_upload_vdom" {
	mkey = "cert_test1_vdom"
	type = "CertKey"
	upload = "text"
	vdom = "vdom1"
	cert = "/root/terr_test/host.cert"
	key = "/root/terr_test/host.key"
}

resource "fortiadc_system_certificate_local_upload" "local_upload_global" {
	mkey = "cert_test1_global"
	type = "CertKey"
	upload = "text"
	vdom = "global"
	cert = "/root/terr_test/host.cert"
	key = "/root/terr_test/host.key"
}


```

## Argument Reference

The following arguments are supported:

* `mkey` - name.
* `type` - CertKey/LocalCert/PKCS12.
* `upload` - must set text.
* `passwd` - password.
* `key` - key file path.
* `cert` - cert file path.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

