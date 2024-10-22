// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load balance servers child virtual server list

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceServersChildVirtualServerList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceServersChildVirtualServerListRead,
		Update: resourceGlobalLoadBalanceServersChildVirtualServerListUpdate,
		Create: resourceGlobalLoadBalanceServersChildVirtualServerListCreate,
		Delete: resourceGlobalLoadBalanceServersChildVirtualServerListDelete,

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
			"address_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gateway_select": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"description": &schema.Schema{
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

func resourceGlobalLoadBalanceServersChildVirtualServerListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceServersChildVirtualServerList: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalLoadBalanceServersChildVirtualServerList: type error")
	}

	obj, err := getObjectGlobalLoadBalanceServers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceServersChildVirtualServerList resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_servers_child_virtual_server_list?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceServersChildVirtualServerList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalLoadBalanceServersChildVirtualServerListRead(d, m)
}

func resourceGlobalLoadBalanceServersChildVirtualServerListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceServersChildVirtualServerList: type error")
	}

	obj, err := getObjectGlobalLoadBalanceServersChildVirtualServerList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceServersChildVirtualServerList resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_servers_child_virtual_server_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceServersChildVirtualServerList resource: %v", err)
	}

	return resourceGlobalLoadBalanceServersChildVirtualServerListRead(d, m)
}

func resourceGlobalLoadBalanceServersChildVirtualServerListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceServersChildVirtualServerList: type error")
	}

	path := "/api/global_load_balance_servers_child_virtual_server_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceServersChildVirtualServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceServersChildVirtualServerListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceServersChildVirtualServerList: type error")
	}

	path := "/api/global_load_balance_servers_child_virtual_server_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceServersChildVirtualServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceServersChildVirtualServerList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceServersChildVirtualServerList resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceServersChildVirtualServerList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceServersChildVirtualServerList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalLoadBalanceServersChildVirtualServerList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("address_type", flattenGlobalLoadBalanceServersChildVirtualServerList(o["address-type"], d, "address_type", sv)); err != nil {
		if !fortiAPIPatch(o["address-type"]) {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenGlobalLoadBalanceServersChildVirtualServerList(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenGlobalLoadBalanceServersChildVirtualServerList(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("gateway", flattenGlobalLoadBalanceServersChildVirtualServerList(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway_select", flattenGlobalLoadBalanceServersChildVirtualServerList(o["gateway_select"], d, "gateway_select", sv)); err != nil {
		if !fortiAPIPatch(o["gateway_select"]) {
			return fmt.Errorf("Error reading gateway_select: %v", err)
		}
	}

	if err = d.Set("description", flattenGlobalLoadBalanceServersChildVirtualServerList(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceServersChildVirtualServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceServersChildVirtualServerList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("address_type"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway_select"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "gateway_select", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway_select"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandGlobalLoadBalanceServersChildVirtualServerList(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
