---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_dos_dos_protection_profile"
description: |-
  Get information on an fortiadc security dos dos protection profile
---

# Data Source: fortiadc_security_dos_dos_protection_profile
Use this data source to get information on an fortiadc security dos dos protection profile

## Example Usage

```hcl
 data "fortiadc_security_dos_dos_protection_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_security_dos_dos_protection_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  security dos dos protection profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - name.
* `tcp_slow_data_limit` - tcp slowdata limit. 
* `http_conn_limit` - HTTP connection limit. 
* `tcp_conn_limit` - tcp connection access limit. 
* `http_req_limit` - HTTP request limit. 
* `http_access_limit` - HTTP access limit. 

