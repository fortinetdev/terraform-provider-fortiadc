// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system isp addr.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIspAddr() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemIspAddrRead,
		Update: resourceSystemIspAddrUpdate,
		Create: resourceSystemIspAddrCreate,
		Delete: resourceSystemIspAddrDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
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
func resourceSystemIspAddrCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspAddr: type error")
	}

	obj, err := getObjectSystemIspAddr(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemIspAddr resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIspAddr(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemIspAddr resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemIspAddrRead(d, m)
}
func resourceSystemIspAddrUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemIspAddr(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIspAddr resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIspAddr(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemIspAddr resource: %v", err)
	}

	return resourceSystemIspAddrRead(d, m)
}
func resourceSystemIspAddrDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemIspAddr(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIspAddr resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemIspAddrRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemIspAddr(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemIspAddr resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIspAddr(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIspAddr resource from API: %v", err)
	}
	return nil
}
func flattenSystemIspAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIspAddrMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIspAddr(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenSystemIspAddrType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemIspAddrMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemIspAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIspAddrMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIspAddr(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemIspAddrType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemIspAddrMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
