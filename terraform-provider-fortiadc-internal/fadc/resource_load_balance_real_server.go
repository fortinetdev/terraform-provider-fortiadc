// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance real server.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceRealServer() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceRealServerRead,
		Update: resourceLoadBalanceRealServerUpdate,
		Create: resourceLoadBalanceRealServerCreate,
		Delete: resourceLoadBalanceRealServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_addr_private": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
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
func resourceLoadBalanceRealServerCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceRealServer: type error")
	}

	obj, err := getObjectLoadBalanceRealServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceRealServer resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceRealServer(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceRealServer resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceRealServerRead(d, m)
}
func resourceLoadBalanceRealServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceRealServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceRealServer resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceRealServer(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceRealServer resource: %v", err)
	}

	return resourceLoadBalanceRealServerRead(d, m)
}
func resourceLoadBalanceRealServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceRealServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceRealServer resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceRealServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceRealServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceRealServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceRealServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceRealServer resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceRealServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSdnConnector(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSdnAddrPrivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerInstance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceRealServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLoadBalanceRealServerStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server_type", flattenLoadBalanceRealServerServerType(o["server_type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server_type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("sdn_connector", flattenLoadBalanceRealServerSdnConnector(o["sdn_connector"], d, "sdn_connector", sv)); err != nil {
		if !fortiAPIPatch(o["sdn_connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	if err = d.Set("sdn_addr_private", flattenLoadBalanceRealServerSdnAddrPrivate(o["sdn_addr_private"], d, "sdn_addr_private", sv)); err != nil {
		if !fortiAPIPatch(o["sdn_addr_private"]) {
			return fmt.Errorf("Error reading sdn_addr_private: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenLoadBalanceRealServerFqdn(o["FQDN"], d, "fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["FQDN"]) {
			return fmt.Errorf("Error reading FQDN: %v", err)
		}
	}

	if err = d.Set("address6", flattenLoadBalanceRealServerAddress6(o["address6"], d, "address6", sv)); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("instance", flattenLoadBalanceRealServerInstance(o["instance"], d, "instance", sv)); err != nil {
		if !fortiAPIPatch(o["instance"]) {
			return fmt.Errorf("Error reading instance: %v", err)
		}
	}

	if err = d.Set("address", flattenLoadBalanceRealServerAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceRealServerType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceRealServerMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceRealServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSdnConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSdnAddrPrivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerInstance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceRealServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLoadBalanceRealServerStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandLoadBalanceRealServerServerType(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_type"] = t
		}
	}

	if v, ok := d.GetOk("sdn_connector"); ok {
		t, err := expandLoadBalanceRealServerSdnConnector(d, v, "sdn_connector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn_connector"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_private"); ok {
		t, err := expandLoadBalanceRealServerSdnAddrPrivate(d, v, "sdn_addr_private", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn_addr_private"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {
		t, err := expandLoadBalanceRealServerFqdn(d, v, "fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FQDN"] = t
		}
	}

	if v, ok := d.GetOk("address6"); ok {
		t, err := expandLoadBalanceRealServerAddress6(d, v, "address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("instance"); ok {
		t, err := expandLoadBalanceRealServerInstance(d, v, "instance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["instance"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandLoadBalanceRealServerAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceRealServerType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceRealServerMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
