// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf advanced protection child advanced protection rule

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleRead,
		Update: resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleUpdate,
		Create: resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleCreate,
		Delete: resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"resp_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"occurrence_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"occurrence_within": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"percentage_match": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"severity": &schema.Schema{
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

func resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	obj, err := getObjectSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_advanced_protection_child_advanced_protection_rule?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource: %v", err)
	}

	id := "SecurityWafAdvancedProtectionChildAdvancedProtectionRule"

	d.SetId(id)

	return resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleRead(d, m)
}

func resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleUpdate(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	obj, err := getObjectSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_advanced_protection_child_advanced_protection_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource: %v", err)
	}

	return resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleRead(d, m)
}

func resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleDelete(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	path := "/api/security_waf_advanced_protection_child_advanced_protection_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafAdvancedProtectionChildAdvancedProtectionRuleRead(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

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
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdvancedProtectionChildAdvancedProtectionRule: type error")
	}

	path := "/api/security_waf_advanced_protection_child_advanced_protection_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdvancedProtectionChildAdvancedProtectionRule resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("type", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("content_type", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["content-type"], d, "content_type", sv)); err != nil {
		if !fortiAPIPatch(o["content-type"]) {
			return fmt.Errorf("Error reading content_type: %v", err)
		}
	}

	if err = d.Set("occurrence_limit", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["occurrence-limit"], d, "occurrence_limit", sv)); err != nil {
		if !fortiAPIPatch(o["occurrence-limit"]) {
			return fmt.Errorf("Error reading occurrence_limit: %v", err)
		}
	}

	if err = d.Set("occurrence_within", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["occurrence-within"], d, "occurrence_within", sv)); err != nil {
		if !fortiAPIPatch(o["occurrence-within"]) {
			return fmt.Errorf("Error reading occurrence_within: %v", err)
		}
	}

	if err = d.Set("percentage_match", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["percentage-match"], d, "percentage_match", sv)); err != nil {
		if !fortiAPIPatch(o["percentage-match"]) {
			return fmt.Errorf("Error reading percentage_match: %v", err)
		}
	}

	if err = d.Set("resp_code", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["resp-code"], d, "resp_code", sv)); err != nil {
		if !fortiAPIPatch(o["resp-code"]) {
			return fmt.Errorf("Error reading resp_code: %v", err)
		}
	}

	if err = d.Set("action", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("severity", flattenSecurityWafAdvancedProtectionChildAdvancedProtectionRule(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("content_type"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "content_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-type"] = t
		}
	}

	if v, ok := d.GetOk("occurrence_limit"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "occurrence_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["occurrence-limit"] = t
		}
	}

	if v, ok := d.GetOk("occurrence_within"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "occurrence_within", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["occurrence-within"] = t
		}
	}

	if v, ok := d.GetOk("percentage_match"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "percentage_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["percentage-match"] = t
		}
	}

	if v, ok := d.GetOk("resp_code"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "resp_code", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resp-code"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandSecurityWafAdvancedProtectionChildAdvancedProtectionRule(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	return &obj, nil
}
