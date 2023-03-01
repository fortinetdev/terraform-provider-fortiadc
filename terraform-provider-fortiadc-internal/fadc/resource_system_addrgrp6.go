// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system addrgrp6.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAddrgrp6() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAddrgrp6Read,
		Update: resourceSystemAddrgrp6Update,
		Create: resourceSystemAddrgrp6Create,
		Delete: resourceSystemAddrgrp6Delete,

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
func resourceSystemAddrgrp6Create(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAddrgrp6: type error")
	}

	obj, err := getObjectSystemAddrgrp6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAddrgrp6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAddrgrp6(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAddrgrp6 resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAddrgrp6Read(d, m)
}
func resourceSystemAddrgrp6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAddrgrp6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAddrgrp6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAddrgrp6(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAddrgrp6 resource: %v", err)
	}

	return resourceSystemAddrgrp6Read(d, m)
}
func resourceSystemAddrgrp6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemAddrgrp6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAddrgrp6 resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemAddrgrp6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAddrgrp6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAddrgrp6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAddrgrp6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAddrgrp6 resource from API: %v", err)
	}
	return nil
}
func flattenSystemAddrgrp6MemberList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAddrgrp6Mkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAddrgrp6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("member_list", flattenSystemAddrgrp6MemberList(o["member-list"], d, "member_list", sv)); err != nil {
		if !fortiAPIPatch(o["member-list"]) {
			return fmt.Errorf("Error reading member-list: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemAddrgrp6Mkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemAddrgrp6MemberList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAddrgrp6Mkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAddrgrp6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member_list"); ok {
		t, err := expandSystemAddrgrp6MemberList(d, v, "member_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-list"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemAddrgrp6Mkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
