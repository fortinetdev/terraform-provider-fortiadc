// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system traffic group.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTrafficGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemTrafficGroupRead,
		Update: resourceSystemTrafficGroupUpdate,
		Create: resourceSystemTrafficGroupCreate,
		Delete: resourceSystemTrafficGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"network_failover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"preempt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"failover_order": &schema.Schema{
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
func resourceSystemTrafficGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemTrafficGroup: type error")
	}

	obj, err := getObjectSystemTrafficGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemTrafficGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemTrafficGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemTrafficGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemTrafficGroupRead(d, m)
}
func resourceSystemTrafficGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemTrafficGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemTrafficGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemTrafficGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemTrafficGroup resource: %v", err)
	}

	return resourceSystemTrafficGroupRead(d, m)
}
func resourceSystemTrafficGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemTrafficGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemTrafficGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemTrafficGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemTrafficGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemTrafficGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemTrafficGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemTrafficGroup resource from API: %v", err)
	}
	return nil
}
func flattenSystemTrafficGroupNetworkFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemTrafficGroupPreempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemTrafficGroupMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemTrafficGroupFailoverOrder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemTrafficGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("network_failover", flattenSystemTrafficGroupNetworkFailover(o["network-failover"], d, "network_failover", sv)); err != nil {
		if !fortiAPIPatch(o["network-failover"]) {
			return fmt.Errorf("Error reading network-failover: %v", err)
		}
	}

	if err = d.Set("preempt", flattenSystemTrafficGroupPreempt(o["preempt"], d, "preempt", sv)); err != nil {
		if !fortiAPIPatch(o["preempt"]) {
			return fmt.Errorf("Error reading preempt: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemTrafficGroupMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("failover_order", flattenSystemTrafficGroupFailoverOrder(o["failover-order"], d, "failover_order", sv)); err != nil {
		if !fortiAPIPatch(o["failover-order"]) {
			return fmt.Errorf("Error reading failover-order: %v", err)
		}
	}

	return nil
}

func expandSystemTrafficGroupNetworkFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemTrafficGroupPreempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemTrafficGroupMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemTrafficGroupFailoverOrder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemTrafficGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("network_failover"); ok {
		t, err := expandSystemTrafficGroupNetworkFailover(d, v, "network_failover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-failover"] = t
		}
	}

	if v, ok := d.GetOk("preempt"); ok {
		t, err := expandSystemTrafficGroupPreempt(d, v, "preempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preempt"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemTrafficGroupMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("failover_order"); ok {
		t, err := expandSystemTrafficGroupFailoverOrder(d, v, "failover_order", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failover-order"] = t
		}
	}

	return &obj, nil
}
