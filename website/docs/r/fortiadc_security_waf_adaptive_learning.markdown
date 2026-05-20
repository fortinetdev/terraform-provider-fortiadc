---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_adaptive_learning"
description: |-
  Configure fortiadc security waf adaptive learning.
---

# fortiadc_security_waf_adaptive_learning
Configure fortiadc security waf adaptive learning.

## Example Usage
```hcl
resource "fortiadc_security_waf_adaptive_learning" "test" {
	mkey = "test"
	status = "enable"
	sampling_rate = "100"
	fp_threshold = "1"
	fp_exp_id = "test"
	least_time = "10081"
	action = "alert"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `status` - Enable Status to view the Adaptive Learning configuration parameters.
* `sampling_rate` - Specify the percentage of received requests and their responses that will be sampled. Range: 1-100.
* `fp_threshold` - Specify the threshold at which triggered events should be considered a false positive. Range: 0-100000000.
* `fp_exp_id` - Select an existing WAF Exception object from the drop-down list or create a new one to use as the Adaptive Learning False Positive Policy.
* `least_time` - Specify the Learning Time period in minutes. Range: 1-20160.
* `action` - Set the action to take after you accept the recommendation for the WAF policy from Adaptive Learning.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Exception can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_adaptive_learning.labelname {{mkey}}
```
