// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance schedule pool.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceSchedulePool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceSchedulePoolRead,
		Update: resourceLoadBalanceSchedulePoolUpdate,
		Create: resourceLoadBalanceSchedulePoolCreate,
		Delete: resourceLoadBalanceSchedulePoolDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule": &schema.Schema{
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
func resourceLoadBalanceSchedulePoolCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceSchedulePool: type error")
	}

	obj, err := getObjectLoadBalanceSchedulePool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceSchedulePool resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceSchedulePool(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceSchedulePool resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceSchedulePoolRead(d, m)
}
func resourceLoadBalanceSchedulePoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceSchedulePool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceSchedulePool resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceSchedulePool(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceSchedulePool resource: %v", err)
	}

	return resourceLoadBalanceSchedulePoolRead(d, m)
}
func resourceLoadBalanceSchedulePoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceSchedulePool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceSchedulePool resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceSchedulePoolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceSchedulePool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceSchedulePool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceSchedulePool(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceSchedulePool resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceSchedulePoolPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceSchedulePoolMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceSchedulePoolSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceSchedulePool(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("pool", flattenLoadBalanceSchedulePoolPool(o["pool"], d, "pool", sv)); err != nil {
		if !fortiAPIPatch(o["pool"]) {
			return fmt.Errorf("Error reading pool: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceSchedulePoolMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("schedule", flattenLoadBalanceSchedulePoolSchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceSchedulePoolPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceSchedulePoolMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceSchedulePoolSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceSchedulePool(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pool"); ok {
		t, err := expandLoadBalanceSchedulePoolPool(d, v, "pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceSchedulePoolMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandLoadBalanceSchedulePoolSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	return &obj, nil
}
