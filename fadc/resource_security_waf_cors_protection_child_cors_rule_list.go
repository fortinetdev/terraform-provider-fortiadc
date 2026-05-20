// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf cors protection child cors rule list

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafCorsProtectionChildCorsRuleList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafCorsProtectionChildCorsRuleListRead,
		Update: resourceSecurityWafCorsProtectionChildCorsRuleListUpdate,
		Create: resourceSecurityWafCorsProtectionChildCorsRuleListCreate,
		Delete: resourceSecurityWafCorsProtectionChildCorsRuleListDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"apply_to_all_cors_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"allowed_origin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"insert_allowed_credentials": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_credentials": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"insert_max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_methods": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"methods": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_headers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_headers_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exposed_headers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exposed_headers_list": &schema.Schema{
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

func resourceSecurityWafCorsProtectionChildCorsRuleListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	obj, err := getObjectSecurityWafCorsProtectionChildCorsRuleList(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafCorsProtectionChildCorsRuleList resource while getting object: %v", err)
	}

	path := "/api/security_waf_cors_protection_child_cors_rule_list?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafCorsProtectionChildCorsRuleList resource: %v", err)
	}

	id := "SecurityWafCorsProtectionChildCorsRuleList"

	d.SetId(id)

	return resourceSecurityWafCorsProtectionChildCorsRuleListRead(d, m)
}

func resourceSecurityWafCorsProtectionChildCorsRuleListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	obj, err := getObjectSecurityWafCorsProtectionChildCorsRuleList(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafCorsProtectionChildCorsRuleList resource while getting object: %v", err)
	}

	path := "/api/security_waf_cors_protection_child_cors_rule_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafCorsProtectionChildCorsRuleList resource: %v", err)
	}

	return resourceSecurityWafCorsProtectionChildCorsRuleListRead(d, m)
}

func resourceSecurityWafCorsProtectionChildCorsRuleListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	path := "/api/security_waf_cors_protection_child_cors_rule_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafCorsProtectionChildCorsRuleList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafCorsProtectionChildCorsRuleListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafCorsProtectionChildCorsRuleList: type error")
	}

	path := "/api/security_waf_cors_protection_child_cors_rule_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafCorsProtectionChildCorsRuleList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafCorsProtectionChildCorsRuleList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafCorsProtectionChildCorsRuleList resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafCorsProtectionChildCorsRuleList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafCorsProtectionChildCorsRuleList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafCorsProtectionChildCorsRuleList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("action", flattenSecurityWafCorsProtectionChildCorsRuleList(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("host_status", flattenSecurityWafCorsProtectionChildCorsRuleList(o["host_status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host_status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}

	if err = d.Set("host", flattenSecurityWafCorsProtectionChildCorsRuleList(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("request_url", flattenSecurityWafCorsProtectionChildCorsRuleList(o["request_url"], d, "request_url", sv)); err != nil {
		if !fortiAPIPatch(o["request_url"]) {
			return fmt.Errorf("Error reading request_url: %v", err)
		}
	}

	if err = d.Set("apply_to_all_cors_traffic", flattenSecurityWafCorsProtectionChildCorsRuleList(o["apply_to_all_cors_traffic"], d, "apply_to_all_cors_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["apply_to_all_cors_traffic"]) {
			return fmt.Errorf("Error reading apply_to_all_cors_traffic: %v", err)
		}
	}

	if err = d.Set("allowed_origin", flattenSecurityWafCorsProtectionChildCorsRuleList(o["allowed_origin"], d, "allowed_origin", sv)); err != nil {
		if !fortiAPIPatch(o["allowed_origin"]) {
			return fmt.Errorf("Error reading allowed_origin: %v", err)
		}
	}

	if err = d.Set("insert_allowed_credentials", flattenSecurityWafCorsProtectionChildCorsRuleList(o["insert_allowed_credentials"], d, "insert_allowed_credentials", sv)); err != nil {
		if !fortiAPIPatch(o["insert_allowed_credentials"]) {
			return fmt.Errorf("Error reading insert_allowed_credentials: %v", err)
		}
	}

	if err = d.Set("allowed_credentials", flattenSecurityWafCorsProtectionChildCorsRuleList(o["allowed_credentials"], d, "allowed_credentials", sv)); err != nil {
		if !fortiAPIPatch(o["allowed_credentials"]) {
			return fmt.Errorf("Error reading allowed_credentials: %v", err)
		}
	}

	if err = d.Set("insert_max_age", flattenSecurityWafCorsProtectionChildCorsRuleList(o["insert_max_age"], d, "insert_max_age", sv)); err != nil {
		if !fortiAPIPatch(o["insert_max_age"]) {
			return fmt.Errorf("Error reading insert_max_age: %v", err)
		}
	}

	if err = d.Set("allowed_max_age", flattenSecurityWafCorsProtectionChildCorsRuleList(o["allowed_max_age"], d, "allowed_max_age", sv)); err != nil {
		if !fortiAPIPatch(o["allowed_max_age"]) {
			return fmt.Errorf("Error reading allowed_max_age: %v", err)
		}
	}

	if err = d.Set("allowed_methods", flattenSecurityWafCorsProtectionChildCorsRuleList(o["allowed_methods"], d, "allowed_methods", sv)); err != nil {
		if !fortiAPIPatch(o["allowed_methods"]) {
			return fmt.Errorf("Error reading allowed_methods: %v", err)
		}
	}

	if err = d.Set("methods", flattenSecurityWafCorsProtectionChildCorsRuleList(o["methods"], d, "methods", sv)); err != nil {
		if !fortiAPIPatch(o["methods"]) {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("allowed_headers", flattenSecurityWafCorsProtectionChildCorsRuleList(o["allowed_headers"], d, "allowed_headers", sv)); err != nil {
		if !fortiAPIPatch(o["allowed_headers"]) {
			return fmt.Errorf("Error reading allowed_headers: %v", err)
		}
	}

	if err = d.Set("allowed_headers_list", flattenSecurityWafCorsProtectionChildCorsRuleList(o["allowed_headers_list"], d, "allowed_headers_list", sv)); err != nil {
		if !fortiAPIPatch(o["allowed_headers_list"]) {
			return fmt.Errorf("Error reading allowed_headers_list: %v", err)
		}
	}

	if err = d.Set("exposed_headers", flattenSecurityWafCorsProtectionChildCorsRuleList(o["exposed_headers"], d, "exposed_headers", sv)); err != nil {
		if !fortiAPIPatch(o["exposed_headers"]) {
			return fmt.Errorf("Error reading exposed_headers: %v", err)
		}
	}

	if err = d.Set("exposed_headers_list", flattenSecurityWafCorsProtectionChildCorsRuleList(o["exposed_headers_list"], d, "exposed_headers_list", sv)); err != nil {
		if !fortiAPIPatch(o["exposed_headers_list"]) {
			return fmt.Errorf("Error reading exposed_headers_list: %v", err)
		}
	}

	return nil
}

func expandSecurityWafCorsProtectionChildCorsRuleList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafCorsProtectionChildCorsRuleList(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_status"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("request_url"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "request_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request_url"] = t
		}
	}

	if v, ok := d.GetOk("apply_to_all_cors_traffic"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "apply_to_all_cors_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apply_to_all_cors_traffic"] = t
		}
	}

	if v, ok := d.GetOk("allowed_origin"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "allowed_origin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed_origin"] = t
		}
	}

	if v, ok := d.GetOk("insert_allowed_credentials"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "insert_allowed_credentials", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["insert_allowed_credentials"] = t
		}
	}

	if v, ok := d.GetOk("allowed_credentials"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "allowed_credentials", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed_credentials"] = t
		}
	}

	if v, ok := d.GetOk("insert_max_age"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "insert_max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["insert_max_age"] = t
		}
	}

	if v, ok := d.GetOk("allowed_max_age"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "allowed_max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed_max_age"] = t
		}
	}

	if v, ok := d.GetOk("allowed_methods"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "allowed_methods", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed_methods"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "methods", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("allowed_headers"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "allowed_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed_headers"] = t
		}
	}

	if v, ok := d.GetOk("allowed_headers_list"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "allowed_headers_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed_headers_list"] = t
		}
	}

	if v, ok := d.GetOk("exposed_headers"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "exposed_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exposed_headers"] = t
		}
	}

	if v, ok := d.GetOk("exposed_headers_list"); ok {
		t, err := expandSecurityWafCorsProtectionChildCorsRuleList(d, v, "exposed_headers_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exposed_headers_list"] = t
		}
	}
	return &obj, nil
}
