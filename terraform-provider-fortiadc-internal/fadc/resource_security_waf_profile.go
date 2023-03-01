// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security waf profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafProfileRead,
		Update: resourceSecurityWafProfileUpdate,
		Create: resourceSecurityWafProfileCreate,
		Delete: resourceSecurityWafProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"data_leak_prevention_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cors_protection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_protocol_constraint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_header_security_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"threshold_based_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"biometrics_based_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"xml_validation_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"csrf_protection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_security": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"brute_force_login_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"credential_stuffing_defense": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"heuristic_sql_xss_injection_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_attack_signature": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rule_match_record": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"openapi_validation_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_protection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"api_gateway_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"bot_detection_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"json_validation_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"advanced_protection_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"input_validation_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSecurityWafProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSecurityWafProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSecurityWafProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafProfileRead(d, m)
}
func resourceSecurityWafProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityWafProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafProfile resource: %v", err)
	}

	return resourceSecurityWafProfileRead(d, m)
}
func resourceSecurityWafProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSecurityWafProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SecurityWafProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSecurityWafProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSecurityWafProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafProfile resource from API: %v", err)
	}
	return nil
}
func flattenSecurityWafProfileDataLeakPreventionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileCorsProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileHttpProtocolConstraint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileHttpHeaderSecurityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileThresholdBasedDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileBiometricsBasedDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileExceptionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileXmlValidationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileCsrfProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileCookieSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileBruteForceLoginName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileCredentialStuffingDefense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileHeuristicSqlXssInjectionDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileWebAttackSignature(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileRuleMatchRecord(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileOpenapiValidationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileUrlProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileDesc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileApiGatewayPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileBotDetectionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileJsonValidationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileAdvancedProtectionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityWafProfileInputValidationPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("data_leak_prevention_name", flattenSecurityWafProfileDataLeakPreventionName(o["data_leak_prevention_name"], d, "data_leak_prevention_name", sv)); err != nil {
		if !fortiAPIPatch(o["data_leak_prevention_name"]) {
			return fmt.Errorf("Error reading data_leak_prevention_name: %v", err)
		}
	}

	if err = d.Set("cors_protection", flattenSecurityWafProfileCorsProtection(o["cors_protection"], d, "cors_protection", sv)); err != nil {
		if !fortiAPIPatch(o["cors_protection"]) {
			return fmt.Errorf("Error reading cors_protection: %v", err)
		}
	}

	if err = d.Set("http_protocol_constraint", flattenSecurityWafProfileHttpProtocolConstraint(o["http_protocol_constraint"], d, "http_protocol_constraint", sv)); err != nil {
		if !fortiAPIPatch(o["http_protocol_constraint"]) {
			return fmt.Errorf("Error reading http_protocol_constraint: %v", err)
		}
	}

	if err = d.Set("http_header_security_name", flattenSecurityWafProfileHttpHeaderSecurityName(o["http_header_security_name"], d, "http_header_security_name", sv)); err != nil {
		if !fortiAPIPatch(o["http_header_security_name"]) {
			return fmt.Errorf("Error reading http_header_security_name: %v", err)
		}
	}

	if err = d.Set("threshold_based_detection", flattenSecurityWafProfileThresholdBasedDetection(o["threshold_based_detection"], d, "threshold_based_detection", sv)); err != nil {
		if !fortiAPIPatch(o["threshold_based_detection"]) {
			return fmt.Errorf("Error reading threshold_based_detection: %v", err)
		}
	}

	if err = d.Set("biometrics_based_detection", flattenSecurityWafProfileBiometricsBasedDetection(o["biometrics_based_detection"], d, "biometrics_based_detection", sv)); err != nil {
		if !fortiAPIPatch(o["biometrics_based_detection"]) {
			return fmt.Errorf("Error reading biometrics_based_detection: %v", err)
		}
	}

	if err = d.Set("exception_name", flattenSecurityWafProfileExceptionName(o["exception_name"], d, "exception_name", sv)); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}

	if err = d.Set("xml_validation_name", flattenSecurityWafProfileXmlValidationName(o["xml_validation_name"], d, "xml_validation_name", sv)); err != nil {
		if !fortiAPIPatch(o["xml_validation_name"]) {
			return fmt.Errorf("Error reading xml_validation_name: %v", err)
		}
	}

	if err = d.Set("csrf_protection", flattenSecurityWafProfileCsrfProtection(o["csrf_protection"], d, "csrf_protection", sv)); err != nil {
		if !fortiAPIPatch(o["csrf_protection"]) {
			return fmt.Errorf("Error reading csrf_protection: %v", err)
		}
	}

	if err = d.Set("cookie_security", flattenSecurityWafProfileCookieSecurity(o["cookie_security"], d, "cookie_security", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_security"]) {
			return fmt.Errorf("Error reading cookie_security: %v", err)
		}
	}

	if err = d.Set("brute_force_login_name", flattenSecurityWafProfileBruteForceLoginName(o["brute_force_login_name"], d, "brute_force_login_name", sv)); err != nil {
		if !fortiAPIPatch(o["brute_force_login_name"]) {
			return fmt.Errorf("Error reading brute_force_login_name: %v", err)
		}
	}

	if err = d.Set("credential_stuffing_defense", flattenSecurityWafProfileCredentialStuffingDefense(o["credential_stuffing_defense"], d, "credential_stuffing_defense", sv)); err != nil {
		if !fortiAPIPatch(o["credential_stuffing_defense"]) {
			return fmt.Errorf("Error reading credential_stuffing_defense: %v", err)
		}
	}

	if err = d.Set("heuristic_sql_xss_injection_detection", flattenSecurityWafProfileHeuristicSqlXssInjectionDetection(o["heuristic_sql_xss_injection_detection"], d, "heuristic_sql_xss_injection_detection", sv)); err != nil {
		if !fortiAPIPatch(o["heuristic_sql_xss_injection_detection"]) {
			return fmt.Errorf("Error reading heuristic_sql_xss_injection_detection: %v", err)
		}
	}

	if err = d.Set("web_attack_signature", flattenSecurityWafProfileWebAttackSignature(o["web_attack_signature"], d, "web_attack_signature", sv)); err != nil {
		if !fortiAPIPatch(o["web_attack_signature"]) {
			return fmt.Errorf("Error reading web_attack_signature: %v", err)
		}
	}

	if err = d.Set("rule_match_record", flattenSecurityWafProfileRuleMatchRecord(o["rule_match_record"], d, "rule_match_record", sv)); err != nil {
		if !fortiAPIPatch(o["rule_match_record"]) {
			return fmt.Errorf("Error reading rule_match_record: %v", err)
		}
	}

	if err = d.Set("openapi_validation_name", flattenSecurityWafProfileOpenapiValidationName(o["openapi_validation_name"], d, "openapi_validation_name", sv)); err != nil {
		if !fortiAPIPatch(o["openapi_validation_name"]) {
			return fmt.Errorf("Error reading openapi_validation_name: %v", err)
		}
	}

	if err = d.Set("url_protection", flattenSecurityWafProfileUrlProtection(o["url_protection"], d, "url_protection", sv)); err != nil {
		if !fortiAPIPatch(o["url_protection"]) {
			return fmt.Errorf("Error reading url_protection: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSecurityWafProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("desc", flattenSecurityWafProfileDesc(o["desc"], d, "desc", sv)); err != nil {
		if !fortiAPIPatch(o["desc"]) {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("api_gateway_policy_name", flattenSecurityWafProfileApiGatewayPolicyName(o["api_gateway_policy_name"], d, "api_gateway_policy_name", sv)); err != nil {
		if !fortiAPIPatch(o["api_gateway_policy_name"]) {
			return fmt.Errorf("Error reading api_gateway_policy_name: %v", err)
		}
	}

	if err = d.Set("bot_detection_name", flattenSecurityWafProfileBotDetectionName(o["bot_detection_name"], d, "bot_detection_name", sv)); err != nil {
		if !fortiAPIPatch(o["bot_detection_name"]) {
			return fmt.Errorf("Error reading bot_detection_name: %v", err)
		}
	}

	if err = d.Set("json_validation_name", flattenSecurityWafProfileJsonValidationName(o["json_validation_name"], d, "json_validation_name", sv)); err != nil {
		if !fortiAPIPatch(o["json_validation_name"]) {
			return fmt.Errorf("Error reading json_validation_name: %v", err)
		}
	}

	if err = d.Set("advanced_protection_name", flattenSecurityWafProfileAdvancedProtectionName(o["advanced_protection_name"], d, "advanced_protection_name", sv)); err != nil {
		if !fortiAPIPatch(o["advanced_protection_name"]) {
			return fmt.Errorf("Error reading advanced_protection_name: %v", err)
		}
	}

	if err = d.Set("input_validation_policy_name", flattenSecurityWafProfileInputValidationPolicyName(o["input_validation_policy_name"], d, "input_validation_policy_name", sv)); err != nil {
		if !fortiAPIPatch(o["input_validation_policy_name"]) {
			return fmt.Errorf("Error reading input_validation_policy_name: %v", err)
		}
	}

	return nil
}

func expandSecurityWafProfileDataLeakPreventionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileCorsProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileHttpProtocolConstraint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileHttpHeaderSecurityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileThresholdBasedDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileBiometricsBasedDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileExceptionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileXmlValidationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileCsrfProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileCookieSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileBruteForceLoginName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileCredentialStuffingDefense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileHeuristicSqlXssInjectionDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileWebAttackSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileRuleMatchRecord(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileOpenapiValidationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileUrlProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileDesc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileApiGatewayPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileBotDetectionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileJsonValidationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileAdvancedProtectionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityWafProfileInputValidationPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("data_leak_prevention_name"); ok {
		t, err := expandSecurityWafProfileDataLeakPreventionName(d, v, "data_leak_prevention_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data_leak_prevention_name"] = t
		}
	}

	if v, ok := d.GetOk("cors_protection"); ok {
		t, err := expandSecurityWafProfileCorsProtection(d, v, "cors_protection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors_protection"] = t
		}
	}

	if v, ok := d.GetOk("http_protocol_constraint"); ok {
		t, err := expandSecurityWafProfileHttpProtocolConstraint(d, v, "http_protocol_constraint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_protocol_constraint"] = t
		}
	}

	if v, ok := d.GetOk("http_header_security_name"); ok {
		t, err := expandSecurityWafProfileHttpHeaderSecurityName(d, v, "http_header_security_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_header_security_name"] = t
		}
	}

	if v, ok := d.GetOk("threshold_based_detection"); ok {
		t, err := expandSecurityWafProfileThresholdBasedDetection(d, v, "threshold_based_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold_based_detection"] = t
		}
	}

	if v, ok := d.GetOk("biometrics_based_detection"); ok {
		t, err := expandSecurityWafProfileBiometricsBasedDetection(d, v, "biometrics_based_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["biometrics_based_detection"] = t
		}
	}

	if v, ok := d.GetOk("exception_name"); ok {
		t, err := expandSecurityWafProfileExceptionName(d, v, "exception_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_name"] = t
		}
	}

	if v, ok := d.GetOk("xml_validation_name"); ok {
		t, err := expandSecurityWafProfileXmlValidationName(d, v, "xml_validation_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xml_validation_name"] = t
		}
	}

	if v, ok := d.GetOk("csrf_protection"); ok {
		t, err := expandSecurityWafProfileCsrfProtection(d, v, "csrf_protection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csrf_protection"] = t
		}
	}

	if v, ok := d.GetOk("cookie_security"); ok {
		t, err := expandSecurityWafProfileCookieSecurity(d, v, "cookie_security", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_security"] = t
		}
	}

	if v, ok := d.GetOk("brute_force_login_name"); ok {
		t, err := expandSecurityWafProfileBruteForceLoginName(d, v, "brute_force_login_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["brute_force_login_name"] = t
		}
	}

	if v, ok := d.GetOk("credential_stuffing_defense"); ok {
		t, err := expandSecurityWafProfileCredentialStuffingDefense(d, v, "credential_stuffing_defense", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["credential_stuffing_defense"] = t
		}
	}

	if v, ok := d.GetOk("heuristic_sql_xss_injection_detection"); ok {
		t, err := expandSecurityWafProfileHeuristicSqlXssInjectionDetection(d, v, "heuristic_sql_xss_injection_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heuristic_sql_xss_injection_detection"] = t
		}
	}

	if v, ok := d.GetOk("web_attack_signature"); ok {
		t, err := expandSecurityWafProfileWebAttackSignature(d, v, "web_attack_signature", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web_attack_signature"] = t
		}
	}

	if v, ok := d.GetOk("rule_match_record"); ok {
		t, err := expandSecurityWafProfileRuleMatchRecord(d, v, "rule_match_record", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule_match_record"] = t
		}
	}

	if v, ok := d.GetOk("openapi_validation_name"); ok {
		t, err := expandSecurityWafProfileOpenapiValidationName(d, v, "openapi_validation_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openapi_validation_name"] = t
		}
	}

	if v, ok := d.GetOk("url_protection"); ok {
		t, err := expandSecurityWafProfileUrlProtection(d, v, "url_protection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url_protection"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandSecurityWafProfileDesc(d, v, "desc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway_policy_name"); ok {
		t, err := expandSecurityWafProfileApiGatewayPolicyName(d, v, "api_gateway_policy_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api_gateway_policy_name"] = t
		}
	}

	if v, ok := d.GetOk("bot_detection_name"); ok {
		t, err := expandSecurityWafProfileBotDetectionName(d, v, "bot_detection_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bot_detection_name"] = t
		}
	}

	if v, ok := d.GetOk("json_validation_name"); ok {
		t, err := expandSecurityWafProfileJsonValidationName(d, v, "json_validation_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["json_validation_name"] = t
		}
	}

	if v, ok := d.GetOk("advanced_protection_name"); ok {
		t, err := expandSecurityWafProfileAdvancedProtectionName(d, v, "advanced_protection_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advanced_protection_name"] = t
		}
	}

	if v, ok := d.GetOk("input_validation_policy_name"); ok {
		t, err := expandSecurityWafProfileInputValidationPolicyName(d, v, "input_validation_policy_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input_validation_policy_name"] = t
		}
	}

	return &obj, nil
}
