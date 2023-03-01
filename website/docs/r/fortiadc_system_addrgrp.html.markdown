---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_addrgrp"
description: |-
  Configure fortiadc address group info.
---

# fortiadc_system_addrgrp
Configure fortiadc address group info.

## Example Usage
```hcl
resource "fortiadc_system_addrgrp" "addressgrp1" {
	mkey = "addrgrp_1"
	member_list = "addr_1"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - address group name.
* `member_list` - member list. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Addrgrp can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_addrgrp.labelname {{mkey}}
```
