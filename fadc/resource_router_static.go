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

func resourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterStaticRead,
		Update: resourceRouterStaticUpdate,
		Create: resourceRouterStaticCreate,
		Delete: resourceRouterStaticDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dest": &schema.Schema{
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
func resourceRouterStaticCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterStatic: type error")
	}

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic resource while getting object: %v", err)
	}

	_, err = c.CreateRouterStatic(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterStaticRead(d, m)
}
func resourceRouterStaticUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterStatic(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic resource: %v", err)
	}

	return resourceRouterStaticRead(d, m)
}
func resourceRouterStaticDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterStatic(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterStatic resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterStaticRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterStatic(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic resource from API: %v", err)
	}
	return nil
}
func flattenRouterStaticGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterStaticMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterStatic(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("gw", flattenRouterStaticGw(o["gw"], d, "gw", sv)); err != nil {
		if !fortiAPIPatch(o["gw"]) {
			return fmt.Errorf("Error reading gw: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterStaticDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dest", flattenRouterStaticDest(o["dest"], d, "dest", sv)); err != nil {
		if !fortiAPIPatch(o["dest"]) {
			return fmt.Errorf("Error reading dest: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterStaticMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterStaticGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("gw"); ok {
		t, err := expandRouterStaticGw(d, v, "gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gw"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandRouterStaticDistance(d, v, "distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dest"); ok {
		t, err := expandRouterStaticDest(d, v, "dest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterStaticMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
