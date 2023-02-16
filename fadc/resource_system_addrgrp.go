// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system addrgrp.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAddrgrp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAddrgrpRead,
		Update: resourceSystemAddrgrpUpdate,
		Create: resourceSystemAddrgrpCreate,
		Delete: resourceSystemAddrgrpDelete,

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
func resourceSystemAddrgrpCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAddrgrp: type error")
	}

	obj, err := getObjectSystemAddrgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAddrgrp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAddrgrp(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAddrgrp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAddrgrpRead(d, m)
}
func resourceSystemAddrgrpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAddrgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAddrgrp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAddrgrp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAddrgrp resource: %v", err)
	}

	return resourceSystemAddrgrpRead(d, m)
}
func resourceSystemAddrgrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemAddrgrp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAddrgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemAddrgrpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAddrgrp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAddrgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAddrgrp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAddrgrp resource from API: %v", err)
	}
	return nil
}
func flattenSystemAddrgrpMemberList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAddrgrpMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAddrgrp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("member_list", flattenSystemAddrgrpMemberList(o["member-list"], d, "member_list", sv)); err != nil {
		if !fortiAPIPatch(o["member-list"]) {
			return fmt.Errorf("Error reading member-list: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemAddrgrpMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemAddrgrpMemberList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAddrgrpMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAddrgrp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member_list"); ok {
		t, err := expandSystemAddrgrpMemberList(d, v, "member_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-list"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemAddrgrpMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
