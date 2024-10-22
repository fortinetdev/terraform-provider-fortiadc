// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load balance virtual server pool

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceVirtualServerPool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceVirtualServerPoolRead,
		Update: resourceGlobalLoadBalanceVirtualServerPoolUpdate,
		Create: resourceGlobalLoadBalanceVirtualServerPoolCreate,
		Delete: resourceGlobalLoadBalanceVirtualServerPoolDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"preferred": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"preferred_cmr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"alternate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"alternate_cmr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"load_balance_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"check_server_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"check_virtual_server_existent": &schema.Schema{
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

func resourceGlobalLoadBalanceVirtualServerPoolCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPool: type error")
	}

	obj, err := getObjectGlobalLoadBalanceVirtualServerPool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceVirtualServerPool resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_virtual_server_pool"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceVirtualServerPool resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalLoadBalanceVirtualServerPoolRead(d, m)
}

func resourceGlobalLoadBalanceVirtualServerPoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalLoadBalanceVirtualServerPool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceVirtualServerPool resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_virtual_server_pool?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceVirtualServerPool resource: %v", err)
	}

	return resourceGlobalLoadBalanceVirtualServerPoolRead(d, m)
}

func resourceGlobalLoadBalanceVirtualServerPoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_virtual_server_pool?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceVirtualServerPool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceVirtualServerPoolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_virtual_server_pool?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceVirtualServerPool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceVirtualServerPool(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceVirtualServerPool resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceVirtualServerPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceVirtualServerPool(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalLoadBalanceVirtualServerPool(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("preferred", flattenGlobalLoadBalanceVirtualServerPool(o["preferred"], d, "preferred", sv)); err != nil {
		if !fortiAPIPatch(o["preferred"]) {
			return fmt.Errorf("Error reading preferred: %v", err)
		}
	}

	if err = d.Set("preferred_cmr", flattenGlobalLoadBalanceVirtualServerPool(o["preferred-cmr"], d, "preferred_cmr", sv)); err != nil {
		if !fortiAPIPatch(o["preferred-cmr"]) {
			return fmt.Errorf("Error reading preferred_cmr: %v", err)
		}
	}

	if err = d.Set("alternate", flattenGlobalLoadBalanceVirtualServerPool(o["alternate"], d, "alternate", sv)); err != nil {
		if !fortiAPIPatch(o["alternate"]) {
			return fmt.Errorf("Error reading alternate: %v", err)
		}
	}

	if err = d.Set("alternate_cmr", flattenGlobalLoadBalanceVirtualServerPool(o["alternate-cmr"], d, "alternate_cmr", sv)); err != nil {
		if !fortiAPIPatch(o["alternate-cmr"]) {
			return fmt.Errorf("Error reading alternate_cmr: %v", err)
		}
	}

	if err = d.Set("load_balance_method", flattenGlobalLoadBalanceVirtualServerPool(o["load-balance-method"], d, "load_balance_method", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-method"]) {
			return fmt.Errorf("Error reading load_balance_method: %v", err)
		}
	}

	if err = d.Set("check_server_status", flattenGlobalLoadBalanceVirtualServerPool(o["check-server-status"], d, "check_server_status", sv)); err != nil {
		if !fortiAPIPatch(o["check-server-status"]) {
			return fmt.Errorf("Error reading check_server_status: %v", err)
		}
	}

	if err = d.Set("check_virtual_server_existent", flattenGlobalLoadBalanceVirtualServerPool(o["check-virtual-server-existent"], d, "check_virtual_server_existent", sv)); err != nil {
		if !fortiAPIPatch(o["check-virtual-server-existent"]) {
			return fmt.Errorf("Error reading check_virtual_server_existent: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceVirtualServerPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceVirtualServerPool(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("preferred"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "preferred", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred"] = t
		}
	}

	if v, ok := d.GetOk("preferred_cmr"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "preferred_cmr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-cmr"] = t
		}
	}

	if v, ok := d.GetOk("alternate"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "alternate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alternate"] = t
		}
	}

	if v, ok := d.GetOk("alternate_cmr"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "alternate_cmr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alternate-cmr"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_method"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "load_balance_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-method"] = t
		}
	}

	if v, ok := d.GetOk("check_server_status"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "check_server_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-server-status"] = t
		}
	}

	if v, ok := d.GetOk("check_virtual_server_existent"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPool(d, v, "check_virtual_server_existent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-virtual-server-existent"] = t
		}
	}

	return &obj, nil
}
