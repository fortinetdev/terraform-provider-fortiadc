// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router md5 ospf.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterMd5Ospf() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterMd5OspfRead,
		Update: resourceRouterMd5OspfUpdate,
		Create: resourceRouterMd5OspfCreate,
		Delete: resourceRouterMd5OspfDelete,

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
func resourceRouterMd5OspfCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterMd5Ospf: type error")
	}

	obj, err := getObjectRouterMd5Ospf(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterMd5Ospf resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMd5Ospf(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterMd5Ospf resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterMd5OspfRead(d, m)
}
func resourceRouterMd5OspfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterMd5Ospf(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMd5Ospf resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMd5Ospf(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterMd5Ospf resource: %v", err)
	}

	return resourceRouterMd5OspfRead(d, m)
}
func resourceRouterMd5OspfDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterMd5Ospf(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMd5Ospf resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterMd5OspfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterMd5Ospf(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterMd5Ospf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMd5Ospf(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMd5Ospf resource from API: %v", err)
	}
	return nil
}
func flattenRouterMd5OspfMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMd5Ospf(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenRouterMd5OspfMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterMd5OspfMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMd5Ospf(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterMd5OspfMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
