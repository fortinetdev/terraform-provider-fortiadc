// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf api discovery child api security rule

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafApiDiscoveryChildApiSecurityRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafApiDiscoveryChildApiSecurityRuleRead,
		Update: resourceSecurityWafApiDiscoveryChildApiSecurityRuleUpdate,
		Create: resourceSecurityWafApiDiscoveryChildApiSecurityRuleCreate,
		Delete: resourceSecurityWafApiDiscoveryChildApiSecurityRuleDelete,

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
			"base_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"req_rate": &schema.Schema{
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

func resourceSecurityWafApiDiscoveryChildApiSecurityRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	obj, err := getObjectSecurityWafApiDiscoveryChildApiSecurityRule(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafApiDiscoveryChildApiSecurityRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_api_discovery_child_api_security_rule?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafApiDiscoveryChildApiSecurityRule resource: %v", err)
	}

	id := "SecurityWafApiDiscoveryChildApiSecurityRule"

	d.SetId(id)

	return resourceSecurityWafApiDiscoveryChildApiSecurityRuleRead(d, m)
}

func resourceSecurityWafApiDiscoveryChildApiSecurityRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	obj, err := getObjectSecurityWafApiDiscoveryChildApiSecurityRule(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafApiDiscoveryChildApiSecurityRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_api_discovery_child_api_security_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafApiDiscoveryChildApiSecurityRule resource: %v", err)
	}

	return resourceSecurityWafApiDiscoveryChildApiSecurityRuleRead(d, m)
}

func resourceSecurityWafApiDiscoveryChildApiSecurityRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	path := "/api/security_waf_api_discovery_child_api_security_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafApiDiscoveryChildApiSecurityRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafApiDiscoveryChildApiSecurityRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafApiDiscoveryChildApiSecurityRule: type error")
	}

	path := "/api/security_waf_api_discovery_child_api_security_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafApiDiscoveryChildApiSecurityRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafApiDiscoveryChildApiSecurityRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafApiDiscoveryChildApiSecurityRule resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafApiDiscoveryChildApiSecurityRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafApiDiscoveryChildApiSecurityRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafApiDiscoveryChildApiSecurityRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("base_url", flattenSecurityWafApiDiscoveryChildApiSecurityRule(o["base_url"], d, "base_url", sv)); err != nil {
		if !fortiAPIPatch(o["base_url"]) {
			return fmt.Errorf("Error reading base_url: %v", err)
		}
	}

	if err = d.Set("path", flattenSecurityWafApiDiscoveryChildApiSecurityRule(o["path"], d, "path", sv)); err != nil {
		if !fortiAPIPatch(o["path"]) {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("req_rate", flattenSecurityWafApiDiscoveryChildApiSecurityRule(o["req_rate"], d, "req_rate", sv)); err != nil {
		if !fortiAPIPatch(o["req_rate"]) {
			return fmt.Errorf("Error reading req_rate: %v", err)
		}
	}

	if err = d.Set("action", flattenSecurityWafApiDiscoveryChildApiSecurityRule(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("severity", flattenSecurityWafApiDiscoveryChildApiSecurityRule(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	return nil
}

func expandSecurityWafApiDiscoveryChildApiSecurityRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafApiDiscoveryChildApiSecurityRule(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafApiDiscoveryChildApiSecurityRule(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("base_url"); ok {
		t, err := expandSecurityWafApiDiscoveryChildApiSecurityRule(d, v, "base_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base_url"] = t
		}
	}

	if v, ok := d.GetOk("path"); ok {
		t, err := expandSecurityWafApiDiscoveryChildApiSecurityRule(d, v, "path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	}

	if v, ok := d.GetOk("req_rate"); ok {
		t, err := expandSecurityWafApiDiscoveryChildApiSecurityRule(d, v, "req_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["req_rate"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandSecurityWafApiDiscoveryChildApiSecurityRule(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandSecurityWafApiDiscoveryChildApiSecurityRule(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	return &obj, nil
}
