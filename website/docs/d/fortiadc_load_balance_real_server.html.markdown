---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_real_server"
description: |-
  Get information on an fortiadc load balance real server
---

# Data Source: fortiadc_load_balance_real_server
Use this data source to get information on an fortiadc load balance real server

## Example Usage

```hcl
 data "fortiadc_load_balance_real_server" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_real_server.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance real server.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - real server name.
* `status` - (en|dis)able/maintain real server. 
* `server_type` - real server type. 
* `sdn_connector` - sdn connector. 
* `sdn_addr_private` - use private node ip. 
* `fqdn` - Fully Qualified Domain Name. 
* `address6` - ipv6 address of interface. 
* `instance` - instances filter. 
* `address` - ip address of interface. 
* `type` - ip or fqdn. 

