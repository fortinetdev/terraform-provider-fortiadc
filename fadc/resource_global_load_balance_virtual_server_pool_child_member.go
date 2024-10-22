// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load virtual server pool child member

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceVirtualServerPoolChildMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceVirtualServerPoolChildMemberRead,
		Update: resourceGlobalLoadBalanceVirtualServerPoolChildMemberUpdate,
		Create: resourceGlobalLoadBalanceVirtualServerPoolChildMemberCreate,
		Delete: resourceGlobalLoadBalanceVirtualServerPoolChildMemberDelete,

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
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_member_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backup": &schema.Schema{
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

func resourceGlobalLoadBalanceVirtualServerPoolChildMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	obj, err := getObjectGlobalLoadBalanceVirtualServerPoolChildMember(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceVirtualServerPoolChildMember resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_virtual_server_pool_child_member?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceVirtualServerPoolChildMember resource: %v", err)
	}

	id := "GlobalLoadBalanceVirtualServerPoolChildMember"

	d.SetId(id)

	return resourceGlobalLoadBalanceVirtualServerPoolChildMemberRead(d, m)
}

func resourceGlobalLoadBalanceVirtualServerPoolChildMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	obj, err := getObjectGlobalLoadBalanceVirtualServerPoolChildMember(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceVirtualServerPoolChildMember resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_virtual_server_pool_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceVirtualServerPoolChildMember resource: %v", err)
	}

	return resourceGlobalLoadBalanceVirtualServerPoolChildMemberRead(d, m)
}

func resourceGlobalLoadBalanceVirtualServerPoolChildMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	path := "/api/global_load_balance_virtual_server_pool_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceVirtualServerPoolChildMember resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceVirtualServerPoolChildMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalLoadBalanceVirtualServerPoolChildMember: type error")
	}

	path := "/api/global_load_balance_virtual_server_pool_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceVirtualServerPoolChildMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceVirtualServerPoolChildMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceVirtualServerPoolChildMember resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceVirtualServerPoolChildMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceVirtualServerPoolChildMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalLoadBalanceVirtualServerPoolChildMember(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("server", flattenGlobalLoadBalanceVirtualServerPoolChildMember(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_member_name", flattenGlobalLoadBalanceVirtualServerPoolChildMember(o["server-member-name"], d, "server_member_name", sv)); err != nil {
		if !fortiAPIPatch(o["server-member-name"]) {
			return fmt.Errorf("Error reading server_member_name: %v", err)
		}
	}

	if err = d.Set("ttl", flattenGlobalLoadBalanceVirtualServerPoolChildMember(o["ttl"], d, "ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ttl"]) {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("weight", flattenGlobalLoadBalanceVirtualServerPoolChildMember(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("backup", flattenGlobalLoadBalanceVirtualServerPoolChildMember(o["backup"], d, "backup", sv)); err != nil {
		if !fortiAPIPatch(o["backup"]) {
			return fmt.Errorf("Error reading backup: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceVirtualServerPoolChildMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceVirtualServerPoolChildMember(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandGlobalLoadBalanceVirtualServerPoolChildMember(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPoolChildMember(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_member_name"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPoolChildMember(d, v, "server_member_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-member-name"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPoolChildMember(d, v, "ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPoolChildMember(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("backup"); ok {
		t, err := expandGlobalLoadBalanceVirtualServerPoolChildMember(d, v, "backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup"] = t
		}
	}

	return &obj, nil
}
