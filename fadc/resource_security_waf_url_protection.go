// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf url protection

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafUrlProtection() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafUrlProtectionRead,
		Update: resourceSecurityWafUrlProtectionUpdate,
		Create: resourceSecurityWafUrlProtectionCreate,
		Delete: resourceSecurityWafUrlProtectionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSecurityWafUrlProtectionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafUrlProtection: type error")
	}

	obj, err := getObjectSecurityWafUrlProtection(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafUrlProtection resource while getting object: %v", err)
	}

	path := "/api/security_waf_url_protection"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafUrlProtection resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafUrlProtectionRead(d, m)
}

func resourceSecurityWafUrlProtectionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafUrlProtection(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafUrlProtection resource while getting object: %v", err)
	}

	path := "/api/security_waf_url_protection?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafUrlProtection resource: %v", err)
	}

	return resourceSecurityWafUrlProtectionRead(d, m)
}

func resourceSecurityWafUrlProtectionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_url_protection?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafUrlProtection resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafUrlProtectionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_url_protection?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafUrlProtection resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafUrlProtection(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafUrlProtection resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafUrlProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafUrlProtection(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafUrlProtection(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityWafUrlProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafUrlProtection(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafUrlProtection(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
