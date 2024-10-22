// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf url protection child file extension rule

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafUrlProtectionChildFileExtensionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafUrlProtectionChildFileExtensionRuleRead,
		Update: resourceSecurityWafUrlProtectionChildFileExtensionRuleUpdate,
		Create: resourceSecurityWafUrlProtectionChildFileExtensionRuleCreate,
		Delete: resourceSecurityWafUrlProtectionChildFileExtensionRuleDelete,

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
			"file_extension_rule_regex": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file_extension_rule_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file_extension_rule_severity": &schema.Schema{
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

func resourceSecurityWafUrlProtectionChildFileExtensionRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	obj, err := getObjectSecurityWafUrlProtectionChildFileExtensionRule(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafUrlProtectionChildFileExtensionRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_url_protection_child_file_extension_rule?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafUrlProtectionChildFileExtensionRule resource: %v", err)
	}

	id := "SecurityWafUrlProtectionChildFileExtensionRule"

	d.SetId(id)

	return resourceSecurityWafUrlProtectionChildFileExtensionRuleRead(d, m)
}

func resourceSecurityWafUrlProtectionChildFileExtensionRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	obj, err := getObjectSecurityWafUrlProtectionChildFileExtensionRule(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafUrlProtectionChildFileExtensionRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_url_protection_child_file_extension_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafUrlProtectionChildFileExtensionRule resource: %v", err)
	}

	return resourceSecurityWafUrlProtectionChildFileExtensionRuleRead(d, m)
}

func resourceSecurityWafUrlProtectionChildFileExtensionRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	path := "/api/security_waf_url_protection_child_file_extension_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafUrlProtectionChildFileExtensionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafUrlProtectionChildFileExtensionRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafUrlProtectionChildFileExtensionRule: type error")
	}

	path := "/api/security_waf_url_protection_child_file_extension_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafUrlProtectionChildFileExtensionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafUrlProtectionChildFileExtensionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafUrlProtectionChildFileExtensionRule resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafUrlProtectionChildFileExtensionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafUrlProtectionChildFileExtensionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafUrlProtectionChildFileExtensionRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("file_extensions_rule_regex", flattenSecurityWafUrlProtectionChildFileExtensionRule(o["file_extension_rule_regex"], d, "file_extension_rule_regex", sv)); err != nil {
		if !fortiAPIPatch(o["file_extension_rule_regex"]) {
			return fmt.Errorf("Error reading file_extension_rule_regex: %v", err)
		}
	}

	if err = d.Set("file_extensions_rule_action", flattenSecurityWafUrlProtectionChildFileExtensionRule(o["file_extension_rule_action"], d, "file_extension_rule_action", sv)); err != nil {
		if !fortiAPIPatch(o["file_extension_rule_action"]) {
			return fmt.Errorf("Error reading file_extension_rule_action: %v", err)
		}
	}

	if err = d.Set("file_extension_rule_severity", flattenSecurityWafUrlProtectionChildFileExtensionRule(o["file_extension_rule_severity"], d, "file_extension_rule_severity", sv)); err != nil {
		if !fortiAPIPatch(o["file_extension_rule_severity"]) {
			return fmt.Errorf("Error reading file_extension_rule_severity: %v", err)
		}
	}

	if err = d.Set("exception_name", flattenSecurityWafUrlProtectionChildFileExtensionRule(o["exception_name"], d, "exception_name", sv)); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}

	return nil
}

func expandSecurityWafUrlProtectionChildFileExtensionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafUrlProtectionChildFileExtensionRule(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafUrlProtectionChildFileExtensionRule(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("file_extension_rule_regex"); ok {
		t, err := expandSecurityWafUrlProtectionChildFileExtensionRule(d, v, "file_extension_rule_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file_extension_rule_regex"] = t
		}
	}

	if v, ok := d.GetOk("file_extension_rule_action"); ok {
		t, err := expandSecurityWafUrlProtectionChildFileExtensionRule(d, v, "file_extension_rule_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file_extension_rule_action"] = t
		}
	}

	if v, ok := d.GetOk("file_extension_rule_severity"); ok {
		t, err := expandSecurityWafUrlProtectionChildFileExtensionRule(d, v, "file_extension_rule_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file_extension_rule_severity"] = t
		}
	}

	if v, ok := d.GetOk("exception_name"); ok {
		t, err := expandSecurityWafUrlProtectionChildFileExtensionRule(d, v, "exception_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_name"] = t
		}
	}

	return &obj, nil
}
