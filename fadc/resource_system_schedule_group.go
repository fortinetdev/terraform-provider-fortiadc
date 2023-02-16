// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system schedule group.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemScheduleGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemScheduleGroupRead,
		Update: resourceSystemScheduleGroupUpdate,
		Create: resourceSystemScheduleGroupCreate,
		Delete: resourceSystemScheduleGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
func resourceSystemScheduleGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemScheduleGroup: type error")
	}

	obj, err := getObjectSystemScheduleGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemScheduleGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemScheduleGroupRead(d, m)
}
func resourceSystemScheduleGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemScheduleGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemScheduleGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleGroup resource: %v", err)
	}

	return resourceSystemScheduleGroupRead(d, m)
}
func resourceSystemScheduleGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemScheduleGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemScheduleGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemScheduleGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemScheduleGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemScheduleGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleGroup resource from API: %v", err)
	}
	return nil
}
func flattenSystemScheduleGroupMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemScheduleGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemScheduleGroupMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemScheduleGroupMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemScheduleGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemScheduleGroupMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
