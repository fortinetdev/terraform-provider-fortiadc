---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_api_discovery"
description: |-
  Configure fortiadc security waf api discovery.
---

# fortiadc_security_waf_api_discovery
Configure fortiadc security waf api discovery.

## Example Usage
```hcl
resource "fortiadc_security_waf_api_discovery" "test" {
	mkey = "test"
	api_discovery = "enable"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `api_discovery` - Enable or disable the API Discovery configuration.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_api_discovery.labelname {{mkey}}
```
