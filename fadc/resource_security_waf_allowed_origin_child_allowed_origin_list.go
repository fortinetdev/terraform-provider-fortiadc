// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf allowed origin child allowed origin list

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAllowedOriginChildAllowedOriginList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafAllowedOriginChildAllowedOriginListRead,
		Update: resourceSecurityWafAllowedOriginChildAllowedOriginListUpdate,
		Create: resourceSecurityWafAllowedOriginChildAllowedOriginListCreate,
		Delete: resourceSecurityWafAllowedOriginChildAllowedOriginListDelete,

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
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"origin_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"include_sub_domains": &schema.Schema{
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

func resourceSecurityWafAllowedOriginChildAllowedOriginListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	obj, err := getObjectSecurityWafAllowedOriginChildAllowedOriginList(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAllowedOriginChildAllowedOriginList resource while getting object: %v", err)
	}

	path := "/api/security_waf_allowed_origin_child_allowed_origin_list?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAllowedOriginChildAllowedOriginList resource: %v", err)
	}

	id := "SecurityWafAllowedOriginChildAllowedOriginList"

	d.SetId(id)

	return resourceSecurityWafAllowedOriginChildAllowedOriginListRead(d, m)
}

func resourceSecurityWafAllowedOriginChildAllowedOriginListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	obj, err := getObjectSecurityWafAllowedOriginChildAllowedOriginList(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAllowedOriginChildAllowedOriginList resource while getting object: %v", err)
	}

	path := "/api/security_waf_allowed_origin_child_allowed_origin_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAllowedOriginChildAllowedOriginList resource: %v", err)
	}

	return resourceSecurityWafAllowedOriginChildAllowedOriginListRead(d, m)
}

func resourceSecurityWafAllowedOriginChildAllowedOriginListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	path := "/api/security_waf_allowed_origin_child_allowed_origin_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAllowedOriginChildAllowedOriginList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafAllowedOriginChildAllowedOriginListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAllowedOriginChildAllowedOriginList: type error")
	}

	path := "/api/security_waf_allowed_origin_child_allowed_origin_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAllowedOriginChildAllowedOriginList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAllowedOriginChildAllowedOriginList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAllowedOriginChildAllowedOriginList resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAllowedOriginChildAllowedOriginList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAllowedOriginChildAllowedOriginList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAllowedOriginChildAllowedOriginList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSecurityWafAllowedOriginChildAllowedOriginList(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("origin_name", flattenSecurityWafAllowedOriginChildAllowedOriginList(o["origin_name"], d, "origin_name", sv)); err != nil {
		if !fortiAPIPatch(o["origin_name"]) {
			return fmt.Errorf("Error reading origin_name: %v", err)
		}
	}

	if err = d.Set("port", flattenSecurityWafAllowedOriginChildAllowedOriginList(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("include_sub_domains", flattenSecurityWafAllowedOriginChildAllowedOriginList(o["include_sub_domains"], d, "include_sub_domains", sv)); err != nil {
		if !fortiAPIPatch(o["include_sub_domains"]) {
			return fmt.Errorf("Error reading include_sub_domains: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAllowedOriginChildAllowedOriginList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAllowedOriginChildAllowedOriginList(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafAllowedOriginChildAllowedOriginList(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSecurityWafAllowedOriginChildAllowedOriginList(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("origin_name"); ok {
		t, err := expandSecurityWafAllowedOriginChildAllowedOriginList(d, v, "origin_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin_name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSecurityWafAllowedOriginChildAllowedOriginList(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("include_sub_domains"); ok {
		t, err := expandSecurityWafAllowedOriginChildAllowedOriginList(d, v, "include_sub_domains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include_sub_domains"] = t
		}
	}
	return &obj, nil
}
