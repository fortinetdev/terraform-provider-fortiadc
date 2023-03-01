---
subcategory: "FortiADC Vdom"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_vdom"
description: |-
  Configure fortiadc route configuration.
---

# fortiadc_vdom
Configure fortiadc vdom configuration.

## Example Usage
```hcl
resource "fortiadc_vdom" "vdom_set" {
	mkey = "vdom2"
}

```

## Argument Reference

The following arguments are supported:

* `mkey` - The vdom name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

