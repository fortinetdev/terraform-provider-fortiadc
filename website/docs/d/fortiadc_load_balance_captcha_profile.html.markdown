---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_captcha_profile"
description: |-
  Get information on an fortiadc load balance captcha profile
---

# Data Source: fortiadc_load_balance_captcha_profile
Use this data source to get information on an fortiadc load balance captcha profile

## Example Usage

```hcl
 data "fortiadc_load_balance_captcha_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_captcha_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance captcha profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - captcha profile name.
* `max_block_period` - User blocked time after verification failure, unit second. (10,864000)
* `custom_page` - If use customize captcha page file.. 
* `max_valid_period` - Maximum valid time for verified users, unit second. (60,864000)
* `max_picture_changes` - Maximum number of times the user can change the captcha picture. (1,100)
* `max_verify_period` - Maximum user verification process, unit second. (20,86400)
* `max_retry_times` - The maximum number of times a user can try. (1,100)
* `vpath` - virtual path of captcha service. 
* `custom_page_file` - Customize captcha page package filename. 
* `pic_difficulty` - Verification picture difficulty. 

