// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVdom() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVdomRead,
		Update: resourceVdomUpdate,
		Create: resourceVdomCreate,
		Delete: resourceVdomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceVdomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing Vdom: type error")
	}

	obj, err := getObjectVdom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating Vdom resource while getting object: %v", err)
	}

	_, err = c.CreateVdom(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating Vdom resource: %v", err)
	}

	d.SetId(mkey)

	return resourceVdomRead(d, m)
}
func resourceVdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	obj, err := getObjectVdom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating Vdom resource while getting object: %v", err)
	}

	_, err = c.UpdateVdom(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating Vdom resource: %v", err)
	}

	return resourceVdomRead(d, m)
}
func resourceVdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	err := c.DeleteVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting Vdom resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceVdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	o, err := c.ReadVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading Vdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVdom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading Vdom resource from API: %v", err)
	}
	return nil
}

func flattenVdomMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVdom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenVdomMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandVdomMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVdom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandVdomMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
