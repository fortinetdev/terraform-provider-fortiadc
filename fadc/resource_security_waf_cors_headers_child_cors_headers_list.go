// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf cors headers child cors headers list

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafCorsHeadersChildCorsHeadersList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafCorsHeadersChildCorsHeadersListRead,
		Update: resourceSecurityWafCorsHeadersChildCorsHeadersListUpdate,
		Create: resourceSecurityWafCorsHeadersChildCorsHeadersListCreate,
		Delete: resourceSecurityWafCorsHeadersChildCorsHeadersListDelete,

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
			"header": &schema.Schema{
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

func resourceSecurityWafCorsHeadersChildCorsHeadersListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	obj, err := getObjectSecurityWafCorsHeadersChildCorsHeadersList(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafCorsHeadersChildCorsHeadersList resource while getting object: %v", err)
	}

	path := "/api/security_waf_cors_headers_child_cors_headers_list?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafCorsHeadersChildCorsHeadersList resource: %v", err)
	}

	id := "SecurityWafCorsHeadersChildCorsHeadersList"

	d.SetId(id)

	return resourceSecurityWafCorsHeadersChildCorsHeadersListRead(d, m)
}

func resourceSecurityWafCorsHeadersChildCorsHeadersListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	obj, err := getObjectSecurityWafCorsHeadersChildCorsHeadersList(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafCorsHeadersChildCorsHeadersList resource while getting object: %v", err)
	}

	path := "/api/security_waf_cors_headers_child_cors_headers_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafCorsHeadersChildCorsHeadersList resource: %v", err)
	}

	return resourceSecurityWafCorsHeadersChildCorsHeadersListRead(d, m)
}

func resourceSecurityWafCorsHeadersChildCorsHeadersListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	path := "/api/security_waf_cors_headers_child_cors_headers_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafCorsHeadersChildCorsHeadersList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafCorsHeadersChildCorsHeadersListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafCorsHeadersChildCorsHeadersList: type error")
	}

	path := "/api/security_waf_cors_headers_child_cors_headers_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafCorsHeadersChildCorsHeadersList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafCorsHeadersChildCorsHeadersList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafCorsHeadersChildCorsHeadersList resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafCorsHeadersChildCorsHeadersList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafCorsHeadersChildCorsHeadersList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafCorsHeadersChildCorsHeadersList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("header", flattenSecurityWafCorsHeadersChildCorsHeadersList(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	return nil
}

func expandSecurityWafCorsHeadersChildCorsHeadersList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafCorsHeadersChildCorsHeadersList(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafCorsHeadersChildCorsHeadersList(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("header"); ok {
		t, err := expandSecurityWafCorsHeadersChildCorsHeadersList(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}
	return &obj, nil
}
