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

func resourceGlobalLoadBalanceDataCenter() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceDataCenterRead,
		Update: resourceGlobalLoadBalanceDataCenterUpdate,
		Create: resourceGlobalLoadBalanceDataCenterCreate,
		Delete: resourceGlobalLoadBalanceDataCenterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"location": &schema.Schema{
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

func resourceGlobalLoadBalanceDataCenterCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalLoadBalanceDataCenter: type error")
	}

	obj, err := getObjectGlobalLoadBalanceDataCenter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceDataCenter resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_data_center"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalLoadBalanceDataCenter resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalLoadBalanceDataCenterRead(d, m)
}
func resourceGlobalLoadBalanceDataCenterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalLoadBalanceDataCenter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceDataCenter resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_data_center" + "?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceDataCenter resource: %v", err)
	}

	return resourceGlobalLoadBalanceDataCenterRead(d, m)
}
func resourceGlobalLoadBalanceDataCenterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_data_center" + "?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error deleting GlobalLoadBalanceDataCenter resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceGlobalLoadBalanceDataCenterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_data_center" + "?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceDataCenter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceDataCenter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceDataCenter resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceDataCenterDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceDataCenterLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceDataCenterMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceDataCenter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenGlobalLoadBalanceDataCenterDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("location", flattenGlobalLoadBalanceDataCenterLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("mkey", flattenGlobalLoadBalanceDataCenterMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceDataCenterDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceDataCenterLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceDataCenterMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceDataCenter(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandGlobalLoadBalanceDataCenterDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {
		t, err := expandGlobalLoadBalanceDataCenterLocation(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalLoadBalanceDataCenterMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
