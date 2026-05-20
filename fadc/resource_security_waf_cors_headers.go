// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf cors headers

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafCorsHeaders() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafCorsHeadersRead,
		Update: resourceSecurityWafCorsHeadersUpdate,
		Create: resourceSecurityWafCorsHeadersCreate,
		Delete: resourceSecurityWafCorsHeadersDelete,

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

func resourceSecurityWafCorsHeadersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsHeaders: type error")
	}

	obj, err := getObjectSecurityWafCorsHeaders(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafCorsHeaders resource while getting object: %v", err)
	}

	path := "/api/security_waf_cors_headers"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafCorsHeaders resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafCorsHeadersRead(d, m)
}

func resourceSecurityWafCorsHeadersUpdate(d *schema.ResourceData, m interface{}) error {
	d.SetId("SecurityWafCorsHeaders")
	return resourceSecurityWafCorsHeadersRead(d, m)
}

func resourceSecurityWafCorsHeadersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_cors_headers?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafCorsHeaders resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafCorsHeadersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_cors_headers?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafCorsHeaders resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafCorsHeaders(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafCorsHeaders resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafCorsHeaders(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafCorsHeaders(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafCorsHeaders(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityWafCorsHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafCorsHeaders(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafCorsHeaders(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
