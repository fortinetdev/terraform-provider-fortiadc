// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load balance host child virtual server pool list

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceHostChildVirtualServerPoolList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceHostChildVirtualServerPoolListRead,
		Update: resourceGlobalLoadBalanceHostChildVirtualServerPoolListUpdate,
		Create: resourceGlobalLoadBalanceHostChildVirtualServerPoolListCreate,
		Delete: resourceGlobalLoadBalanceHostChildVirtualServerPoolListDelete,

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
			"virtual_server_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"topology": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"isp": &schema.Schema{
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

func resourceGlobalLoadBalanceHostChildVirtualServerPoolListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceHostChildVirtualServerPoolList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalLoadBalanceHostChildVirtualServerPoolList: type error")
	}

	obj, err := getObjectGlobalLoadBalanceHostChildVirtualServerPoolList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceHostChildVirtualServerPoolList resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_host_child_virtual_server_pool_list?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceHostChildVirtualServerPoolList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalLoadBalanceHostChildVirtualServerPoolListRead(d, m)
}

func resourceGlobalLoadBalanceHostChildVirtualServerPoolListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

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
		return fmt.Errorf("Error describing GlobalLoadBalanceHostChildVirtualServerPoolList: type error")
	}

	obj, err := getObjectGlobalLoadBalanceHostChildVirtualServerPoolList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceHostChildVirtualServerPoolList resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_host_child_virtual_server_pool_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceHostChildVirtualServerPoolList resource: %v", err)
	}

	return resourceGlobalLoadBalanceHostChildVirtualServerPoolListRead(d, m)
}

func resourceGlobalLoadBalanceHostChildVirtualServerPoolListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

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
		return fmt.Errorf("Error describing GlobalLoadBalanceHostChildVirtualServerPoolList: type error")
	}

	path := "/api/global_load_balance_host_child_virtual_server_pool_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceHostChildVirtualServerPoolList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceHostChildVirtualServerPoolListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

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
		return fmt.Errorf("Error describing GlobalLoadBalanceHostChildVirtualServerPoolList: type error")
	}

	path := "/api/global_load_balance_host_child_virtual_server_pool_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceHostChildVirtualServerPoolList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceHostChildVirtualServerPoolList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceHostChildVirtualServerPoolList resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceHostChildVirtualServerPoolList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceHostChildVirtualServerPoolList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalLoadBalanceHostChildVirtualServerPoolList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("virtual_server_pool", flattenGlobalLoadBalanceHostChildVirtualServerPoolList(o["virtual-server-pool"], d, "virtual_server_pool", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-server-pool"]) {
			return fmt.Errorf("Error reading virtual_server_pool: %v", err)
		}
	}

	if err = d.Set("weight", flattenGlobalLoadBalanceHostChildVirtualServerPoolList(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("topology", flattenGlobalLoadBalanceHostChildVirtualServerPoolList(o["topology"], d, "topology", sv)); err != nil {
		if !fortiAPIPatch(o["topology"]) {
			return fmt.Errorf("Error reading topology: %v", err)
		}
	}

	if err = d.Set("isp", flattenGlobalLoadBalanceHostChildVirtualServerPoolList(o["isp"], d, "isp", sv)); err != nil {
		if !fortiAPIPatch(o["isp"]) {
			return fmt.Errorf("Error reading isp: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceHostChildVirtualServerPoolList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceHostChildVirtualServerPoolList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalLoadBalanceHostChildVirtualServerPoolList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("virtual_server_pool"); ok {
		t, err := expandGlobalLoadBalanceHostChildVirtualServerPoolList(d, v, "virtual_server_pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-server-pool"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandGlobalLoadBalanceHostChildVirtualServerPoolList(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("topology"); ok {
		t, err := expandGlobalLoadBalanceHostChildVirtualServerPoolList(d, v, "topology", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["topology"] = t
		}
	}

	if v, ok := d.GetOk("isp"); ok {
		t, err := expandGlobalLoadBalanceHostChildVirtualServerPoolList(d, v, "isp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isp"] = t
		}
	}

	return &obj, nil
}
