---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_profile"
description: |-
  Get information on an fortiadc security waf profile
---

# Data Source: fortiadc_security_waf_profile
Use this data source to get information on an fortiadc security waf profile

## Example Usage

```hcl
 data "fortiadc_security_waf_profile" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_security_waf_profile.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  security waf profile.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - profile name.
* `data_leak_prevention_name` - Data leak prevention. 
* `cors_protection` - CORS protection. 
* `http_protocol_constraint` - HTTP protocol constraint. 
* `http_header_security_name` - HTTP Header Security. 
* `threshold_based_detection` - Threshold Based Detection. 
* `biometrics_based_detection` - Biometrics Based Detection. 
* `exception_name` - Exceptions. 
* `xml_validation_name` - XML validation. 
* `csrf_protection` - CSRF protection. 
* `cookie_security` - Cookie Security. 
* `brute_force_login_name` - Brute force login. 
* `credential_stuffing_defense` - Credential stuffing defense. 
* `heuristic_sql_xss_injection_detection` - Heuristic sql injection and xss detect. 
* `web_attack_signature` - Web attack signature. 
* `rule_match_record` - record matched rules content enable/disable. 
* `openapi_validation_name` - OPENAPI validation. 
* `url_protection` - URL protection. 
* `desc` - waf profile description. 
* `api_gateway_policy_name` - API Gateway. 
* `bot_detection_name` - Bot detection. 
* `json_validation_name` - JSON validation. 
* `advanced_protection_name` - Advanced protection profile. 
* `input_validation_policy_name` - input validation policy. 

