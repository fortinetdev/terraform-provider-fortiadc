// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system isp province.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIspProvince() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemIspProvinceRead,
		Update: resourceSystemIspProvinceUpdate,
		Create: resourceSystemIspProvinceCreate,
		Delete: resourceSystemIspProvinceDelete,

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
func resourceSystemIspProvinceCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspProvince: type error")
	}

	obj, err := getObjectSystemIspProvince(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemIspProvince resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIspProvince(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemIspProvince resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemIspProvinceRead(d, m)
}
func resourceSystemIspProvinceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemIspProvince(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIspProvince resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIspProvince(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemIspProvince resource: %v", err)
	}

	return resourceSystemIspProvinceRead(d, m)
}
func resourceSystemIspProvinceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemIspProvince(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIspProvince resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemIspProvinceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemIspProvince(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemIspProvince resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIspProvince(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIspProvince resource from API: %v", err)
	}
	return nil
}
func flattenSystemIspProvinceMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIspProvince(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemIspProvinceMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemIspProvinceMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIspProvince(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemIspProvinceMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
