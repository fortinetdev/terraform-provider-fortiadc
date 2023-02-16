// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security waf profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSecurityWafProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityWafProfileRead,
		Schema: map[string]*schema.Schema{
			"data_leak_prevention_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cors_protection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_protocol_constraint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_header_security_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"threshold_based_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"biometrics_based_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"exception_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"xml_validation_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"csrf_protection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_security": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"brute_force_login_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"credential_stuffing_defense": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"heuristic_sql_xss_injection_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_attack_signature": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rule_match_record": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"openapi_validation_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url_protection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_gateway_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bot_detection_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"json_validation_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"advanced_protection_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"input_validation_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceSecurityWafProfileRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafProfile: type error")
	}

	o, err := c.ReadSecurityWafProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SecurityWafProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSecurityWafProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SecurityWafProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSecurityWafProfileDataLeakPreventionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileCorsProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileHttpProtocolConstraint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileHttpHeaderSecurityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileThresholdBasedDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileBiometricsBasedDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileExceptionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileXmlValidationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileCsrfProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileCookieSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileBruteForceLoginName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileCredentialStuffingDefense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileHeuristicSqlXssInjectionDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileWebAttackSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileRuleMatchRecord(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileOpenapiValidationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileUrlProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileApiGatewayPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileBotDetectionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileJsonValidationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileAdvancedProtectionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityWafProfileInputValidationPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSecurityWafProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("data_leak_prevention_name", dataSourceFlattenSecurityWafProfileDataLeakPreventionName(o["data_leak_prevention_name"], d, "data_leak_prevention_name")); err != nil {
		if !fortiAPIPatch(o["data_leak_prevention_name"]) {
			return fmt.Errorf("Error reading data_leak_prevention_name: %v", err)
		}
	}

	if err = d.Set("cors_protection", dataSourceFlattenSecurityWafProfileCorsProtection(o["cors_protection"], d, "cors_protection")); err != nil {
		if !fortiAPIPatch(o["cors_protection"]) {
			return fmt.Errorf("Error reading cors_protection: %v", err)
		}
	}

	if err = d.Set("http_protocol_constraint", dataSourceFlattenSecurityWafProfileHttpProtocolConstraint(o["http_protocol_constraint"], d, "http_protocol_constraint")); err != nil {
		if !fortiAPIPatch(o["http_protocol_constraint"]) {
			return fmt.Errorf("Error reading http_protocol_constraint: %v", err)
		}
	}

	if err = d.Set("http_header_security_name", dataSourceFlattenSecurityWafProfileHttpHeaderSecurityName(o["http_header_security_name"], d, "http_header_security_name")); err != nil {
		if !fortiAPIPatch(o["http_header_security_name"]) {
			return fmt.Errorf("Error reading http_header_security_name: %v", err)
		}
	}

	if err = d.Set("threshold_based_detection", dataSourceFlattenSecurityWafProfileThresholdBasedDetection(o["threshold_based_detection"], d, "threshold_based_detection")); err != nil {
		if !fortiAPIPatch(o["threshold_based_detection"]) {
			return fmt.Errorf("Error reading threshold_based_detection: %v", err)
		}
	}

	if err = d.Set("biometrics_based_detection", dataSourceFlattenSecurityWafProfileBiometricsBasedDetection(o["biometrics_based_detection"], d, "biometrics_based_detection")); err != nil {
		if !fortiAPIPatch(o["biometrics_based_detection"]) {
			return fmt.Errorf("Error reading biometrics_based_detection: %v", err)
		}
	}

	if err = d.Set("exception_name", dataSourceFlattenSecurityWafProfileExceptionName(o["exception_name"], d, "exception_name")); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}

	if err = d.Set("xml_validation_name", dataSourceFlattenSecurityWafProfileXmlValidationName(o["xml_validation_name"], d, "xml_validation_name")); err != nil {
		if !fortiAPIPatch(o["xml_validation_name"]) {
			return fmt.Errorf("Error reading xml_validation_name: %v", err)
		}
	}

	if err = d.Set("csrf_protection", dataSourceFlattenSecurityWafProfileCsrfProtection(o["csrf_protection"], d, "csrf_protection")); err != nil {
		if !fortiAPIPatch(o["csrf_protection"]) {
			return fmt.Errorf("Error reading csrf_protection: %v", err)
		}
	}

	if err = d.Set("cookie_security", dataSourceFlattenSecurityWafProfileCookieSecurity(o["cookie_security"], d, "cookie_security")); err != nil {
		if !fortiAPIPatch(o["cookie_security"]) {
			return fmt.Errorf("Error reading cookie_security: %v", err)
		}
	}

	if err = d.Set("brute_force_login_name", dataSourceFlattenSecurityWafProfileBruteForceLoginName(o["brute_force_login_name"], d, "brute_force_login_name")); err != nil {
		if !fortiAPIPatch(o["brute_force_login_name"]) {
			return fmt.Errorf("Error reading brute_force_login_name: %v", err)
		}
	}

	if err = d.Set("credential_stuffing_defense", dataSourceFlattenSecurityWafProfileCredentialStuffingDefense(o["credential_stuffing_defense"], d, "credential_stuffing_defense")); err != nil {
		if !fortiAPIPatch(o["credential_stuffing_defense"]) {
			return fmt.Errorf("Error reading credential_stuffing_defense: %v", err)
		}
	}

	if err = d.Set("heuristic_sql_xss_injection_detection", dataSourceFlattenSecurityWafProfileHeuristicSqlXssInjectionDetection(o["heuristic_sql_xss_injection_detection"], d, "heuristic_sql_xss_injection_detection")); err != nil {
		if !fortiAPIPatch(o["heuristic_sql_xss_injection_detection"]) {
			return fmt.Errorf("Error reading heuristic_sql_xss_injection_detection: %v", err)
		}
	}

	if err = d.Set("web_attack_signature", dataSourceFlattenSecurityWafProfileWebAttackSignature(o["web_attack_signature"], d, "web_attack_signature")); err != nil {
		if !fortiAPIPatch(o["web_attack_signature"]) {
			return fmt.Errorf("Error reading web_attack_signature: %v", err)
		}
	}

	if err = d.Set("rule_match_record", dataSourceFlattenSecurityWafProfileRuleMatchRecord(o["rule_match_record"], d, "rule_match_record")); err != nil {
		if !fortiAPIPatch(o["rule_match_record"]) {
			return fmt.Errorf("Error reading rule_match_record: %v", err)
		}
	}

	if err = d.Set("openapi_validation_name", dataSourceFlattenSecurityWafProfileOpenapiValidationName(o["openapi_validation_name"], d, "openapi_validation_name")); err != nil {
		if !fortiAPIPatch(o["openapi_validation_name"]) {
			return fmt.Errorf("Error reading openapi_validation_name: %v", err)
		}
	}

	if err = d.Set("url_protection", dataSourceFlattenSecurityWafProfileUrlProtection(o["url_protection"], d, "url_protection")); err != nil {
		if !fortiAPIPatch(o["url_protection"]) {
			return fmt.Errorf("Error reading url_protection: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSecurityWafProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("desc", dataSourceFlattenSecurityWafProfileDesc(o["desc"], d, "desc")); err != nil {
		if !fortiAPIPatch(o["desc"]) {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("api_gateway_policy_name", dataSourceFlattenSecurityWafProfileApiGatewayPolicyName(o["api_gateway_policy_name"], d, "api_gateway_policy_name")); err != nil {
		if !fortiAPIPatch(o["api_gateway_policy_name"]) {
			return fmt.Errorf("Error reading api_gateway_policy_name: %v", err)
		}
	}

	if err = d.Set("bot_detection_name", dataSourceFlattenSecurityWafProfileBotDetectionName(o["bot_detection_name"], d, "bot_detection_name")); err != nil {
		if !fortiAPIPatch(o["bot_detection_name"]) {
			return fmt.Errorf("Error reading bot_detection_name: %v", err)
		}
	}

	if err = d.Set("json_validation_name", dataSourceFlattenSecurityWafProfileJsonValidationName(o["json_validation_name"], d, "json_validation_name")); err != nil {
		if !fortiAPIPatch(o["json_validation_name"]) {
			return fmt.Errorf("Error reading json_validation_name: %v", err)
		}
	}

	if err = d.Set("advanced_protection_name", dataSourceFlattenSecurityWafProfileAdvancedProtectionName(o["advanced_protection_name"], d, "advanced_protection_name")); err != nil {
		if !fortiAPIPatch(o["advanced_protection_name"]) {
			return fmt.Errorf("Error reading advanced_protection_name: %v", err)
		}
	}

	if err = d.Set("input_validation_policy_name", dataSourceFlattenSecurityWafProfileInputValidationPolicyName(o["input_validation_policy_name"], d, "input_validation_policy_name")); err != nil {
		if !fortiAPIPatch(o["input_validation_policy_name"]) {
			return fmt.Errorf("Error reading input_validation_policy_name: %v", err)
		}
	}

	return nil
}
