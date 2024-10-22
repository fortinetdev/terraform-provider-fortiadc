// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load balance host

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceHost() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceHostRead,
		Update: resourceGlobalLoadBalanceHostUpdate,
		Create: resourceGlobalLoadBalanceHostCreate,
		Delete: resourceGlobalLoadBalanceHostDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"respond_single_record": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"load_balance_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_feedback_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_feedback_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fortiview": &schema.Schema{
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

func resourceGlobalLoadBalanceHostCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceHost: type error")
	}

	obj, err := getObjectGlobalLoadBalanceHost(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceHost resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_host"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceHost resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalLoadBalanceHostRead(d, m)
}

func resourceGlobalLoadBalanceHostUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalLoadBalanceHost(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceHost resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_host?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceHost resource: %v", err)
	}

	return resourceGlobalLoadBalanceHostRead(d, m)
}

func resourceGlobalLoadBalanceHostDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_host?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceHost resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceHostRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_host?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceHost resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceHost(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceHost resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceHost(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalLoadBalanceHost(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("host_name", flattenGlobalLoadBalanceHost(o["host-name"], d, "host_name", sv)); err != nil {
		if !fortiAPIPatch(o["host-name"]) {
			return fmt.Errorf("Error reading host_name: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenGlobalLoadBalanceHost(o["domain-name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("dns_policy", flattenGlobalLoadBalanceHost(o["dns_policy"], d, "dns_policy", sv)); err != nil {
		if !fortiAPIPatch(o["dns_policy"]) {
			return fmt.Errorf("Error reading dns_policy: %v", err)
		}
	}

	if err = d.Set("respond_single_record", flattenGlobalLoadBalanceHost(o["respond-single-record"], d, "respond_single_record", sv)); err != nil {
		if !fortiAPIPatch(o["respond-single-record"]) {
			return fmt.Errorf("Error reading respond_single_record: %v", err)
		}
	}

	if err = d.Set("persistence", flattenGlobalLoadBalanceHost(o["persistence"], d, "persistence", sv)); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("load_balance_method", flattenGlobalLoadBalanceHost(o["load-balance-method"], d, "load_balance_method", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-method"]) {
			return fmt.Errorf("Error reading load_balance_method: %v", err)
		}
	}

	if err = d.Set("default_feedback_ip", flattenGlobalLoadBalanceHost(o["default-feedback-ip"], d, "default_feedback_ip", sv)); err != nil {
		if !fortiAPIPatch(o["default-feedback-ip"]) {
			return fmt.Errorf("Error reading default_feedback_ip: %v", err)
		}
	}

	if err = d.Set("default_feedback_ip6", flattenGlobalLoadBalanceHost(o["default-feedback-ip6"], d, "default_feedback_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["default-feedback-ip6"]) {
			return fmt.Errorf("Error reading default_feedback_ip6: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenGlobalLoadBalanceHost(o["fortiview"], d, "fortiview", sv)); err != nil {
		if !fortiAPIPatch(o["fortiview"]) {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceHost(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("host_name"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "host_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-name"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("dns_policy"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "dns_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_policy"] = t
		}
	}

	if v, ok := d.GetOk("respond_single_record"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "respond_single_record", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respond-single-record"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "persistence", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_method"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "load_balance_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-method"] = t
		}
	}

	if v, ok := d.GetOk("default_feedback_ip"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "default_feedback_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-feedback-ip"] = t
		}
	}

	if v, ok := d.GetOk("default_feedback_ip6"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "default_feedback_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-feedback-ip6"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok {
		t, err := expandGlobalLoadBalanceHost(d, v, "fortiview", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	return &obj, nil
}
