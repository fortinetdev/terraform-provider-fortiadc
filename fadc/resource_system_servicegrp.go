// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system servicegrp.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemServicegrp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemServicegrpRead,
		Update: resourceSystemServicegrpUpdate,
		Create: resourceSystemServicegrpCreate,
		Delete: resourceSystemServicegrpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"member_list": &schema.Schema{
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
func resourceSystemServicegrpCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemServicegrp: type error")
	}

	obj, err := getObjectSystemServicegrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemServicegrp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemServicegrp(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemServicegrp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemServicegrpRead(d, m)
}
func resourceSystemServicegrpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemServicegrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemServicegrp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemServicegrp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemServicegrp resource: %v", err)
	}

	return resourceSystemServicegrpRead(d, m)
}
func resourceSystemServicegrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemServicegrp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemServicegrp resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemServicegrpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemServicegrp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemServicegrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemServicegrp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemServicegrp resource from API: %v", err)
	}
	return nil
}
func flattenSystemServicegrpMemberList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServicegrpMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemServicegrp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("member_list", flattenSystemServicegrpMemberList(o["member-list"], d, "member_list", sv)); err != nil {
		if !fortiAPIPatch(o["member-list"]) {
			return fmt.Errorf("Error reading member-list: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemServicegrpMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemServicegrpMemberList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServicegrpMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemServicegrp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member_list"); ok {
		t, err := expandSystemServicegrpMemberList(d, v, "member_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-list"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemServicegrpMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
