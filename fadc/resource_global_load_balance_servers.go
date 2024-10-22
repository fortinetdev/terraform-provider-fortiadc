// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load balance data center

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceServers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceServersRead,
		Update: resourceGlobalLoadBalanceServersUpdate,
		Create: resourceGlobalLoadBalanceServersCreate,
		Delete: resourceGlobalLoadBalanceServersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"user_defined_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"data_center": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auto_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_ctrl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_relationship": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_sdn_private_ip": &schema.Schema{
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

func resourceGlobalLoadBalanceServersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceServers: type error")
	}

	obj, err := getObjectGlobalLoadBalanceServers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceServers resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_servers"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceServers resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalLoadBalanceServersRead(d, m)
}

func resourceGlobalLoadBalanceServersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalLoadBalanceServers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceServers resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_servers?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceServers resource: %v", err)
	}

	return resourceGlobalLoadBalanceServersRead(d, m)
}

func resourceGlobalLoadBalanceServersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_servers?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceServers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceServersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_servers?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceServers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceServers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceServers resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceServers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalLoadBalanceServers(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("server_type", flattenGlobalLoadBalanceServers(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenGlobalLoadBalanceServers(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("password", flattenGlobalLoadBalanceServers(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("user_defined_cert", flattenGlobalLoadBalanceServers(o["user-defined-cert"], d, "user_defined_cert", sv)); err != nil {
		if !fortiAPIPatch(o["user-defined-cert"]) {
			return fmt.Errorf("Error reading user_defined_cert: %v", err)
		}
	}

	if err = d.Set("cert", flattenGlobalLoadBalanceServers(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("address_type", flattenGlobalLoadBalanceServers(o["address-type"], d, "address_type", sv)); err != nil {
		if !fortiAPIPatch(o["address-type"]) {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenGlobalLoadBalanceServers(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenGlobalLoadBalanceServers(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("port", flattenGlobalLoadBalanceServers(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("data_center", flattenGlobalLoadBalanceServers(o["data-center"], d, "data_center", sv)); err != nil {
		if !fortiAPIPatch(o["data-center"]) {
			return fmt.Errorf("Error reading data_center: %v", err)
		}
	}

	if err = d.Set("auto_sync", flattenGlobalLoadBalanceServers(o["auto-sync"], d, "auto_sync", sv)); err != nil {
		if !fortiAPIPatch(o["auto-sync"]) {
			return fmt.Errorf("Error reading auto_sync: %v", err)
		}
	}

	if err = d.Set("health_check_ctrl", flattenGlobalLoadBalanceServers(o["health_check_ctrl"], d, "health_check_ctrl", sv)); err != nil {
		if !fortiAPIPatch(o["health_check_ctrl"]) {
			return fmt.Errorf("Error reading health_check_ctrl: %v", err)
		}
	}

	if err = d.Set("health_check_relationship", flattenGlobalLoadBalanceServers(o["health_check_relationship"], d, "health_check_relationship", sv)); err != nil {
		if !fortiAPIPatch(o["health_check_relationship"]) {
			return fmt.Errorf("Error reading health_check_relationship: %v", err)
		}
	}

	if err = d.Set("health_check_list", flattenGlobalLoadBalanceServers(o["health_check_list"], d, "health_check_list", sv)); err != nil {
		if !fortiAPIPatch(o["health_check_list"]) {
			return fmt.Errorf("Error reading health_check_list: %v", err)
		}
	}

	if err = d.Set("sdn_connector", flattenGlobalLoadBalanceServers(o["sdn-connector"], d, "sdn_connector", sv)); err != nil {
		if !fortiAPIPatch(o["sdn-connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	if err = d.Set("use_sdn_private_ip", flattenGlobalLoadBalanceServers(o["use-sdn-private-ip"], d, "use_sdn_private_ip", sv)); err != nil {
		if !fortiAPIPatch(o["use-sdn-private-ip"]) {
			return fmt.Errorf("Error reading use_sdn_private_ip: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceServers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("user_defined_cert"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "user_defined_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-defined-cert"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("address_type"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("data_center"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "data_center", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-center"] = t
		}
	}

	if v, ok := d.GetOk("auto_sync"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "auto_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-sync"] = t
		}
	}

	if v, ok := d.GetOk("health_check_ctrl"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "health_check_ctrl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check_ctrl"] = t
		}
	}

	if v, ok := d.GetOk("health_check_relationship"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "health_check_relationship", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check_relationship"] = t
		}
	}

	if v, ok := d.GetOk("health_check_list"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "health_check_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check_list"] = t
		}
	}

	if v, ok := d.GetOk("sdn_connector"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "sdn_connector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-connector"] = t
		}
	}

	if v, ok := d.GetOk("use_sdn_private_ip"); ok {
		t, err := expandGlobalLoadBalanceServers(d, v, "use_sdn_private_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdn-private-ip"] = t
		}
	}

	return &obj, nil
}
