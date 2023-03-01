---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_captcha_profile"
description: |-
  Configure fortiadc load-balance captcha profile.
---

# fortiadc_load_balance_captcha_profile
Configure fortiadc load-balance captcha profile.

## Example Usage
```hcl
resource "fortiadc_load_balance_captcha_profile" "captcha" {
	mkey = "lb_captcha"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - captcha profile name.
* `max_block_period` - User blocked time after verification failure, unit second. (10,864000)
* `custom_page` - If use customize captcha page file.. Valid values: enable/disable.
* `max_valid_period` - Maximum valid time for verified users, unit second. (60,864000)
* `max_picture_changes` - Maximum number of times the user can change the captcha picture. (1,100)
* `max_verify_period` - Maximum user verification process, unit second. (20,86400)
* `max_retry_times` - The maximum number of times a user can try. (1,100)
* `vpath` - virtual path of captcha service. 
* `custom_page_file` - Customize captcha page package filename. 
* `pic_difficulty` - Verification picture difficulty. Valid values: 1:hard, 0:easy .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Captcha Profile can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_captcha_profile.labelname {{mkey}}
```
