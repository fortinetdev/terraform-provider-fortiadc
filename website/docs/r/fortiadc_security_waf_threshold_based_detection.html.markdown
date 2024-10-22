---
subcategory: "FortiADC Security"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_security_waf_threshold_based_detection"
description: |-
  Configure security waf threshold based detection.
---

# fortiadc_security_waf_threshold_based_detection
Configure security waf threshold based detection.

## Example Usage
```hcl
resource "fortiadc_security_waf_threshold_based_detection" "security_waf_threshold_based_detection" {
        mkey = "test"
        comments = "nothing"
        crawler_status = "enable"
        response_code = "403"
        crawler_action ="alert"
        crawler_severity = "high"
        crawler_occurrence_limit = "100"
        crawler_occurrence_within = "60"
        content_scraping_status = "enable"
        content_type = "text/html application/soap+xml"
        content_action = "alert"
        content_severity = "medium"
        content_occurrence_limit = "101"
        content_occurrence_within = "61"
        attack_detection_status = "enable"
        attack_modules = "input-validation sql-xss-injection-detection"
        attack_action = "alert"
        attack_severity = "low"
        attack_occurrence_limit = "102"
        attack_occurrence_within = "62"
}


```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of configuration.
* `comments` - Optionally, enter comments about the Threshold Based Detection policy.
* `crawler_status` - Enable/Disable Crawler Detection.
* `response_code` - Specify the 3 digit HTTP response code(s) to check. Range: 100-599.
* `crawler_action` - Select the action profile to apply when a web crawler bot is detected.
* `crawler_severity` - Select the event severity to log when a crawler bot is detected. High, Medium or Low.
* `crawler_occurrence_limit` - Specify the maximum number of responses that can be received from the specified Response Code within the time frame. Range: 1-100000.
* `crawler_occurrence_within` - Specify the time span during which to count how many times a response is received from the specified Response Code. Range: 1-600 seconds.
* `content_scraping_status` - Enable/disable Content Detection.
* `content_type` - Select one or more content type to monitor for content scraping. Example: Text/HTML, Application/XML, Application/JSON
* `content_action` - Select the action profile to apply when a content scraping bot is detected.
* `content_severity` - Select the event severity to log when a content scraping bot is detected. High, Medium or Low.
* `content_occurrence_limit` - Specify the maximum number of responses that can be received from the specified Content Type within the time frame. Range: 1-100000.
* `content_occurrence_within` - Specify the time span during which to count how many times a response is received from the specified Content Type. Range: 1-600 seconds.
* `attack_detection_status` - Enable/disable Attack Detection.
* `attack_modules` - Select one or more attack modules to monitor for bot attacks. Example: Web Attack Signature, Data Leak Prevention or SQL/XSS Injection Detection.
* `attack_action` - Select the action profile to apply when a bot attack is detected.
* `attack_severity` - Select the event severity to log when a bot attack is detected. High, Medium or Low.
* `attack_occurrence_limit` - Specify the maximum number of responses that can be received from the specified Attack Module within the time frame. Range: 1-100000.
* `attack_occurrence_within` - Specify the time span during which to count how many times a response is received from the specified Attack Module. Range: 1-600 seconds.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Security Waf Threshold Based Detections can be imported using any of these accepted formats:
```
$ terraform import fortiadc_security_waf_threshold_based_detection.labelname {{mkey}}
```
