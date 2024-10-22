---
subcategory: "FortiADC Global"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_global_load_balance_setting"
description: |-
  Configure fortiadc global load balance setting.
---

# fortiadc_global_load_balance_setting
Configure fortiadc global load balance setting.

## Example Usage
```hcl
resource "fortiadc_global_load_balance_setting" "global_load_balance_setting" {
        auth_type = "TCP_MD5SIG"
        password = "test"
        ca_verify = "disable"
        trust_ca_group = ""
        trust_inter_ca_group = ""
        ipv4_accessed_status = "disable"
        ipv6_accessed_status = "disable"
        listen_on_interface_all = "disable"
        listen_on_interface_list = ""
        port = "5858"
        proto = "icmp"
        retry_num = "3"
        retry_interval = "3"
        mask_length = "24"
        mask_length6 = "64"
        aging_period = "86400"
        persist_mask_length = "24"
        persist_mask_length6 = "64"
        persist_aging_period = "86400"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `auth_type` - No password, TCP MD5SIG or Auth Verify.
* `password` - Enter the password to authenticate the key. This option is available if Auth Type is TCP MD5SIG or Auth Verify.
* `ca_verify` - Enable/disable the certificate verification when synchronizing the SLB information to the GSLB server.
* `trust_ca_group` - Select a trusted CA group to verify the peer certificate. This option is available if CA Verify is enabled.
* `trust_inter_ca_group` - Select a trusted intermediate CA group to verify the peer certificate. This option is available if CA Verify is enabled.
* `ipv4_accessed_status` - Enable/disable listening for DNS requests on the interface IPv4 address.
* `ipv6_accessed_status` - Enable/disable listening for DNS requests on the interface IPv6 address.
* `listen_on_interface_all` - Enable/disable IPv4/IPv6 network access status on all interfaces.
* `listen_on_interface_list` - Specify specific interfaces for IPv4/IPv6 network access. This option is available if Listen on All Interfaces is disabled.
* `port` - Specify the port to listen on. Default: 5858 Range: 1-65535.
* `proto` - ICMP, ICMP and TCP
* `retry_num` - 	Specify the number of retries if the probe fails. The valid range is 1-10 times.
* `retry_interval` - Specify the interval between retries if the probe fails. The valid range is 1-3600 seconds.
* `mask_length` - Specify the number of IPv4 netmask bits that define network affinity for the persistence table. The default is 24.
* `mask_length6` - Specify the number of IPv6 netmask bits that define network affinity for the persistence table. The default is 64.
* `aging_period` - This setting specifies the length of time in seconds for which the entry is maintained in the persistence table. The default is 86400. The valid range is 60-2,592,000 seconds.
* `persist_mask_length` - Specify the number of IPv4 netmask bits that define network affinity for the RTT table. The default is 24.
* `persist_mask_length6` - Specify the number of IPv6 netmask bits that define network affinity for the RTT table. The default is 64.
* `persist_aging_period` - Specify the for how long RTT results are cached. This setting specifies the length of time in seconds for which the RTT cache entry is valid. The default is 86400. The valid range is 60-2,592,000 seconds.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Global Load Balance Setting can be imported using any of these accepted formats:
```
$ terraform import fortiadc_global_load_balance_setting.labelname GlobalLoadBalanceSetting
```
