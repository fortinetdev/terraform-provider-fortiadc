// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf web attack signature

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafWebAttackSignature() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafWebAttackSignatureRead,
		Update: resourceSecurityWafWebAttackSignatureUpdate,
		Create: resourceSecurityWafWebAttackSignatureCreate,
		Delete: resourceSecurityWafWebAttackSignatureDelete,

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

func resourceSecurityWafWebAttackSignatureCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafWebAttackSignature: type error")
	}

	obj, err := getObjectSecurityWafWebAttackSignature(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafWebAttackSignature resource while getting object: %v", err)
	}

	path := "/api/security_waf_web_attack_signature"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafWebAttackSignature resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafWebAttackSignatureRead(d, m)
}

func resourceSecurityWafWebAttackSignatureUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafWebAttackSignature(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafWebAttackSignature resource while getting object: %v", err)
	}

	path := "/api/security_waf_web_attack_signature?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafWebAttackSignature resource: %v", err)
	}

	return resourceSecurityWafWebAttackSignatureRead(d, m)
}

func resourceSecurityWafWebAttackSignatureDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_web_attack_signature?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafWebAttackSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafWebAttackSignatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_web_attack_signature?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafWebAttackSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafWebAttackSignature(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafWebAttackSignature resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafWebAttackSignature(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafWebAttackSignature(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafWebAttackSignature(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityWafWebAttackSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafWebAttackSignature(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafWebAttackSignature(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
