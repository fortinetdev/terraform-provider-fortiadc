---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_policy_child_rule"
description: |-
  Configure fortiadc firewall policy rule.
---

# fortiadc_firewall_policy_child_rule
Configure fortiadc firewall policy rule.

## Example Usage
```hcl
resource "fortiadc_firewall_policy_child_rule" "firewall_policy_rule" {
	mkey = "test2"
	in_interface = "port4"
	out_interface = "port5"
	source_type = "address"
	source_address = "Any"
	destination_type = "address"
	destination_address = "Any"
	service = "HTTPS"
	action = "deny"
	status = "enable"
	deny_log = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `in_interface` - Ingress Interface. Select the interface that receives traffic.
* `out_interface` - Egress Interface.
Select an outgoing interface from the drop-down list if your FortiADC is configured for link load-balancing and/or traffic routing.
* `source_type` - Select the source type to use to form the matching tuple.
* `source_address` - Specify the Address object to use as the source. The option is available if the source type is Address.
* `source_address_group` - Specify the Address Group object to use as the source. The option is available if the source type is Address Group.
* `source_external_resource_address` - Specify the external IP address list imported through the IP Address connector to use as the source. The option is available if the source type is External Resource.
* `destination_type` - Select the destination type to use to form the matching tuple.
* `destination_address` - Specify the Address object to use as the destination. The option is available if the destination type is Address.
* `destination_address_group` - Specify the Address Group object to use as the destination. The option is available if the destination type is Address Group.
* `destination_external_resource_address` - Specify the external IP address list imported through the IP Address connector to use as the destination. The option is available if the destination type is External Resource.
* `service` - Select a service object to use to form the matching tuple.
* `action` - Deny or Accept.
* `status` - enable/disable the firewall policy rule.
* `deny_log` -  enable/disable logging for denied action. This option is available is the action is Deny.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Firewall Policy Child Rule can be imported using any of these accepted formats:
```
$ terraform import fortiadc_firewall_policy_child_rule.labelname {{mkey}}
```
