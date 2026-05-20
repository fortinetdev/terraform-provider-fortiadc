// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf allowed origin

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAllowedOrigin() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafAllowedOriginRead,
		Update: resourceSecurityWafAllowedOriginUpdate,
		Create: resourceSecurityWafAllowedOriginCreate,
		Delete: resourceSecurityWafAllowedOriginDelete,

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

func resourceSecurityWafAllowedOriginCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAllowedOrigin: type error")
	}

	obj, err := getObjectSecurityWafAllowedOrigin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAllowedOrigin resource while getting object: %v", err)
	}

	path := "/api/security_waf_allowed_origin"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAllowedOrigin resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafAllowedOriginRead(d, m)
}

func resourceSecurityWafAllowedOriginUpdate(d *schema.ResourceData, m interface{}) error {
	d.SetId("SecurityWafAllowedOrigin")
	return resourceSecurityWafAllowedOriginRead(d, m)
}

func resourceSecurityWafAllowedOriginDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_allowed_origin?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAllowedOrigin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafAllowedOriginRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_allowed_origin?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAllowedOrigin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAllowedOrigin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAllowedOrigin resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAllowedOrigin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAllowedOrigin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAllowedOrigin(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAllowedOrigin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAllowedOrigin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafAllowedOrigin(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
