// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf api discovery

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafApiDiscovery() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafApiDiscoveryRead,
		Update: resourceSecurityWafApiDiscoveryUpdate,
		Create: resourceSecurityWafApiDiscoveryCreate,
		Delete: resourceSecurityWafApiDiscoveryDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"api_discovery": &schema.Schema{
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

func resourceSecurityWafApiDiscoveryCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafApiDiscovery: type error")
	}

	obj, err := getObjectSecurityWafApiDiscovery(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafApiDiscovery resource while getting object: %v", err)
	}

	path := "/api/security_waf_api_discovery"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafApiDiscovery resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafApiDiscoveryRead(d, m)
}

func resourceSecurityWafApiDiscoveryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafApiDiscovery(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafApiDiscovery resource while getting object: %v", err)
	}

	path := "/api/security_waf_api_discovery?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafApiDiscovery resource: %v", err)
	}

	return resourceSecurityWafApiDiscoveryRead(d, m)
}

func resourceSecurityWafApiDiscoveryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_api_discovery?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafApiDiscovery resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafApiDiscoveryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_api_discovery?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafApiDiscovery resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafApiDiscovery(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafApiDiscovery resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafApiDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafApiDiscovery(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafApiDiscovery(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("api_discovery", flattenSecurityWafApiDiscovery(o["api_discovery"], d, "api_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["api_discovery"]) {
			return fmt.Errorf("Error reading api_discovery: %v", err)
		}
	}

	return nil
}

func expandSecurityWafApiDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafApiDiscovery(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafApiDiscovery(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("api_discovery"); ok {
		t, err := expandSecurityWafApiDiscovery(d, v, "api_discovery", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api_discovery"] = t
		}
	}

	return &obj, nil
}
