// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf action

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAction() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafActionRead,
		Update: resourceSecurityWafActionUpdate,
		Create: resourceSecurityWafActionCreate,
		Delete: resourceSecurityWafActionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"deny_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"log_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
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

func resourceSecurityWafActionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAction: type error")
	}

	obj, err := getObjectSecurityWafAction(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAction resource while getting object: %v", err)
	}

	path := "/api/security_waf_action"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAction resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafActionRead(d, m)
}

func resourceSecurityWafActionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafAction(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAction resource while getting object: %v", err)
	}

	path := "/api/security_waf_action?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAction resource: %v", err)
	}

	return resourceSecurityWafActionRead(d, m)
}

func resourceSecurityWafActionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_action?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafActionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_action?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAction(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAction resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAction(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAction(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("action_type", flattenSecurityWafAction(o["action_type"], d, "action_type", sv)); err != nil {
		if !fortiAPIPatch(o["action_type"]) {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if err = d.Set("deny_code", flattenSecurityWafAction(o["deny_code"], d, "deny_code", sv)); err != nil {
		if !fortiAPIPatch(o["deny_code"]) {
			return fmt.Errorf("Error reading deny_code: %v", err)
		}
	}

	if err = d.Set("block_period", flattenSecurityWafAction(o["block_period"], d, "block_period", sv)); err != nil {
		if !fortiAPIPatch(o["block_period"]) {
			return fmt.Errorf("Error reading block_period: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenSecurityWafAction(o["redirect_url"], d, "redirect_url", sv)); err != nil {
		if !fortiAPIPatch(o["redirect_url"]) {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("log_status", flattenSecurityWafAction(o["log_status"], d, "log_status", sv)); err != nil {
		if !fortiAPIPatch(o["log_status"]) {
			return fmt.Errorf("Error reading log_status: %v", err)
		}
	}

	if err = d.Set("comments", flattenSecurityWafAction(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAction(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafAction(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("action_type"); ok {
		t, err := expandSecurityWafAction(d, v, "action_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action_type"] = t
		}
	}

	if v, ok := d.GetOk("deny_code"); ok {
		t, err := expandSecurityWafAction(d, v, "deny_code", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny_code"] = t
		}
	}

	if v, ok := d.GetOk("block_period"); ok {
		t, err := expandSecurityWafAction(d, v, "block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_period"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok {
		t, err := expandSecurityWafAction(d, v, "redirect_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect_url"] = t
		}
	}

	if v, ok := d.GetOk("log_status"); ok {
		t, err := expandSecurityWafAction(d, v, "log_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSecurityWafAction(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
