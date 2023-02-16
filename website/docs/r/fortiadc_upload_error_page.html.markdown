---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_upload_error_page"
description: |-
  Upload fortiadc error page.
---

# fortiadc_upload_error_page.
Upload fortiadc error page.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `mkey` - name.
* `vpath` - virtual path for error page.
* `update` - Create or update (true/false).
* `errorpagefile` - error page file path.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

