// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf url protection child url access rule

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafUrlProtectionChildUrlAccessRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafUrlProtectionChildUrlAccessRuleRead,
		Update: resourceSecurityWafUrlProtectionChildUrlAccessRuleUpdate,
		Create: resourceSecurityWafUrlProtectionChildUrlAccessRuleCreate,
		Delete: resourceSecurityWafUrlProtectionChildUrlAccessRuleDelete,

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
			"url_access_rule_regex": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_access_rule_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_access_rule_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_name": &schema.Schema{
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

func resourceSecurityWafUrlProtectionChildUrlAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	obj, err := getObjectSecurityWafUrlProtectionChildUrlAccessRule(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafUrlProtectionChildUrlAccessRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_url_protection_child_url_access_rule?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafUrlProtectionChildUrlAccessRule resource: %v", err)
	}

	id := "SecurityWafUrlProtectionChildUrlAccessRule"

	d.SetId(id)

	return resourceSecurityWafUrlProtectionChildUrlAccessRuleRead(d, m)
}

func resourceSecurityWafUrlProtectionChildUrlAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	obj, err := getObjectSecurityWafUrlProtectionChildUrlAccessRule(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafUrlProtectionChildUrlAccessRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_url_protection_child_url_access_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafUrlProtectionChildUrlAccessRule resource: %v", err)
	}

	return resourceSecurityWafUrlProtectionChildUrlAccessRuleRead(d, m)
}

func resourceSecurityWafUrlProtectionChildUrlAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	path := "/api/security_waf_url_protection_child_url_access_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafUrlProtectionChildUrlAccessRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafUrlProtectionChildUrlAccessRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildUrlAccessRule: type error")
	}

	path := "/api/security_waf_url_protection_child_url_access_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafUrlProtectionChildUrlAccessRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafUrlProtectionChildUrlAccessRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafUrlProtectionChildUrlAccessRule resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafUrlProtectionChildUrlAccessRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafUrlProtectionChildUrlAccessRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafUrlProtectionChildUrlAccessRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("url_access_rule_regex", flattenSecurityWafUrlProtectionChildUrlAccessRule(o["url_access_rule_regex"], d, "url_access_rule_regex", sv)); err != nil {
		if !fortiAPIPatch(o["url_access_rule_regex"]) {
			return fmt.Errorf("Error reading url_access_rule_regex: %v", err)
		}
	}

	if err = d.Set("url_access_rule_action", flattenSecurityWafUrlProtectionChildUrlAccessRule(o["url_access_rule_action"], d, "url_access_rule_action", sv)); err != nil {
		if !fortiAPIPatch(o["url_access_rule_action"]) {
			return fmt.Errorf("Error reading url_access_rule_action: %v", err)
		}
	}

	if err = d.Set("url_access_rule_severity", flattenSecurityWafUrlProtectionChildUrlAccessRule(o["url_access_rule_severity"], d, "url_access_rule_severity", sv)); err != nil {
		if !fortiAPIPatch(o["url_access_rule_severity"]) {
			return fmt.Errorf("Error reading url_access_rule_severity: %v", err)
		}
	}

	if err = d.Set("exception_name", flattenSecurityWafUrlProtectionChildUrlAccessRule(o["exception_name"], d, "exception_name", sv)); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}

	return nil
}

func expandSecurityWafUrlProtectionChildUrlAccessRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafUrlProtectionChildUrlAccessRule(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafUrlProtectionChildUrlAccessRule(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("url_access_rule_regex"); ok {
		t, err := expandSecurityWafUrlProtectionChildUrlAccessRule(d, v, "url_access_rule_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url_access_rule_regex"] = t
		}
	}

	if v, ok := d.GetOk("url_access_rule_action"); ok {
		t, err := expandSecurityWafUrlProtectionChildUrlAccessRule(d, v, "url_access_rule_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url_access_rule_action"] = t
		}
	}

	if v, ok := d.GetOk("url_access_rule_severity"); ok {
		t, err := expandSecurityWafUrlProtectionChildUrlAccessRule(d, v, "url_access_rule_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url_access_rule_severity"] = t
		}
	}

	if v, ok := d.GetOk("exception_name"); ok {
		t, err := expandSecurityWafUrlProtectionChildUrlAccessRule(d, v, "exception_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_name"] = t
		}
	}

	return &obj, nil
}
