// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf exception child exception rule

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafExceptionChildExceptionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafExceptionChildExceptionRuleRead,
		Update: resourceSecurityWafExceptionChildExceptionRuleUpdate,
		Create: resourceSecurityWafExceptionChildExceptionRuleCreate,
		Delete: resourceSecurityWafExceptionChildExceptionRuleDelete,

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
			"exception_host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_value": &schema.Schema{
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

func resourceSecurityWafExceptionChildExceptionRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	obj, err := getObjectSecurityWafExceptionChildExceptionRule(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SercurityWafExceptionChildExceptionRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_exception_child_exception_rule?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SercurityWafExceptionChildExceptionRule resource: %v", err)
	}

	id := "resourceSecurityWafExceptionChildExceptionRule"

	d.SetId(id)

	return resourceSecurityWafExceptionChildExceptionRuleRead(d, m)
}

func resourceSecurityWafExceptionChildExceptionRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	obj, err := getObjectSecurityWafExceptionChildExceptionRule(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SercurityWafExceptionChildExceptionRule resource while getting object: %v", err)
	}

	path := "/api/security_waf_exception_child_exception_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SercurityWafExceptionChildExceptionRule resource: %v", err)
	}

	return resourceSecurityWafExceptionChildExceptionRuleRead(d, m)
}

func resourceSecurityWafExceptionChildExceptionRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	path := "/api/security_waf_exception_child_exception_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SercurityWafExceptionChildExceptionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafExceptionChildExceptionRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SercurityWafExceptionChildExceptionRule: type error")
	}

	path := "/api/security_waf_exception_child_exception_rule?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SercurityWafExceptionChildExceptionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafExceptionChildExceptionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SercurityWafExceptionChildExceptionRule resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafExceptionChildExceptionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafExceptionChildExceptionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafExceptionChildExceptionRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("type", flattenSecurityWafExceptionChildExceptionRule(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("exception_host_status", flattenSecurityWafExceptionChildExceptionRule(o["exception_host_status"], d, "exception_host_status", sv)); err != nil {
		if !fortiAPIPatch(o["exception_host_status"]) {
			return fmt.Errorf("Error reading exception_host_status: %v", err)
		}
	}

	if err = d.Set("exception_host", flattenSecurityWafExceptionChildExceptionRule(o["exception_host"], d, "exception_host", sv)); err != nil {
		if !fortiAPIPatch(o["exception_host"]) {
			return fmt.Errorf("Error reading exception_host: %v", err)
		}
	}

	if err = d.Set("exception_url", flattenSecurityWafExceptionChildExceptionRule(o["exception_url"], d, "exception_url", sv)); err != nil {
		if !fortiAPIPatch(o["exception_url"]) {
			return fmt.Errorf("Error reading exception_url: %v", err)
		}
	}

	if err = d.Set("ip_netmask", flattenSecurityWafExceptionChildExceptionRule(o["ip_netmask"], d, "ip_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ip_netmask"]) {
			return fmt.Errorf("Error reading ip_netmask: %v", err)
		}
	}

	if err = d.Set("ip6_netmask", flattenSecurityWafExceptionChildExceptionRule(o["ip6_netmask"], d, "ip6_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ip6_netmask"]) {
			return fmt.Errorf("Error reading ip6_netmask: %v", err)
		}
	}

	if err = d.Set("exception_method", flattenSecurityWafExceptionChildExceptionRule(o["exception_method"], d, "exception_method", sv)); err != nil {
		if !fortiAPIPatch(o["exception_method"]) {
			return fmt.Errorf("Error reading exception_method: %v", err)
		}
	}

	if err = d.Set("exception_name", flattenSecurityWafExceptionChildExceptionRule(o["exception_name"], d, "exception_name", sv)); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}

	if err = d.Set("exception_value_check", flattenSecurityWafExceptionChildExceptionRule(o["exception_value_check"], d, "exception_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["exception_value_check"]) {
			return fmt.Errorf("Error reading exception_value_check: %v", err)
		}
	}

	if err = d.Set("exception_value", flattenSecurityWafExceptionChildExceptionRule(o["exception_value"], d, "exception_value", sv)); err != nil {
		if !fortiAPIPatch(o["exception_value"]) {
			return fmt.Errorf("Error reading exception_value: %v", err)
		}
	}

	return nil
}

func expandSecurityWafExceptionChildExceptionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafExceptionChildExceptionRule(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("exception_host_status"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_host_status"] = t
		}
	}

	if v, ok := d.GetOk("exception_host"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_host"] = t
		}
	}

	if v, ok := d.GetOk("exception_url"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_url"] = t
		}
	}

	if v, ok := d.GetOk("ip_netmask"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "ip_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_netmask"] = t
		}
	}

	if v, ok := d.GetOk("ip6_netmask"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "ip6_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6_netmask"] = t
		}
	}

	if v, ok := d.GetOk("exception_method"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_method"] = t
		}
	}

	if v, ok := d.GetOk("exception_name"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_name"] = t
		}
	}

	if v, ok := d.GetOk("exception_value_check"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_value_check"] = t
		}
	}

	if v, ok := d.GetOk("exception_value"); ok {
		t, err := expandSecurityWafExceptionChildExceptionRule(d, v, "exception_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_value"] = t
		}
	}

	return &obj, nil
}
