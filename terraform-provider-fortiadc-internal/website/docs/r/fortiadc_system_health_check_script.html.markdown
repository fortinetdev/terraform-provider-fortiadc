---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_health_check_script"
description: |-
  Configure fortiadc the script of health-check.
---

# fortiadc_system_health_check_script
Configure fortiadc the script of health-check.

## Example Usage
```hcl
resource "fortiadc_system_health_check_script" "hc_script1" {
	mkey = "icmp_script1"
	script = "if ping -c 1 $NODE_IP > /dev/null; then\n    exit $UP\nelse\n    exit $DOWN\nfi"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - scripting name.
* `script` - scripting content. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Health Check Script can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_health_check_script.labelname {{mkey}}
```
