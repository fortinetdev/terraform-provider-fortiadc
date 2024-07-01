---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_certificate_crlupload"
description: |-
  Upload fortiadc certificate crl.
---

# fortiadc_certificate_crlupload
Upload fortiadc certificate crl.

## Example Usage
```hcl
resource "fortiadc_certificate_crlupload" "crl_upload_global" {
	mkey = "crl_global"
	cert = "/root/FortiADCTerraform/terr_test/gtsr1.crl"
	vdom = "global"
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

