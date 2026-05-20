// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf advanced protection

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAdvancedProtection() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafAdvancedProtectionRead,
		Update: resourceSecurityWafAdvancedProtectionUpdate,
		Create: resourceSecurityWafAdvancedProtectionCreate,
		Delete: resourceSecurityWafAdvancedProtectionDelete,

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

func resourceSecurityWafAdvancedProtectionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAdvancedProtection: type error")
	}

	obj, err := getObjectSecurityWafAdvancedProtection(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdvancedProtection resource while getting object: %v", err)
	}

	path := "/api/security_waf_advanced_protection"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdvancedProtection resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafAdvancedProtectionRead(d, m)
}

func resourceSecurityWafAdvancedProtectionUpdate(d *schema.ResourceData, m interface{}) error {
	d.SetId("SecurityWafAdvancedProtection")
	return resourceSecurityWafAdvancedProtectionRead(d, m)
}

func resourceSecurityWafAdvancedProtectionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_advanced_protection?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAdvancedProtection resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafAdvancedProtectionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_advanced_protection?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdvancedProtection resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAdvancedProtection(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdvancedProtection resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAdvancedProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAdvancedProtection(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAdvancedProtection(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAdvancedProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAdvancedProtection(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
