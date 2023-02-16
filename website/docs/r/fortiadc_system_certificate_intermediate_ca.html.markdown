---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_certificate_intermediate_ca"
description: |-
  Configure fortiadc Intermediate CA certificate.
---

# fortiadc_system_certificate_intermediate_ca
Configure fortiadc Intermediate CA certificate.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - certificate name.
* `scep_url` - scep url.                                  
* `ca_id` - Specify the CA Identifier.
* `host_header` - host geader.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Certificate Intermediate Ca can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_certificate_intermediate_ca.labelname {{mkey}}
```
