// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child area.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfChildArea() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterOspfChildAreaRead,
		Update: resourceRouterOspfChildAreaUpdate,
		Create: resourceRouterOspfChildAreaCreate,
		Delete: resourceRouterOspfChildAreaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"stub_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
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
func resourceRouterOspfChildAreaCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterOspfChildArea: type error")
	}

	obj, err := getObjectRouterOspfChildArea(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildArea resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfChildArea(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildArea resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterOspfChildAreaRead(d, m)
}
func resourceRouterOspfChildAreaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterOspfChildArea(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildArea resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfChildArea(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildArea resource: %v", err)
	}

	return resourceRouterOspfChildAreaRead(d, m)
}
func resourceRouterOspfChildAreaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterOspfChildArea(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfChildArea resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterOspfChildAreaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterOspfChildArea(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildArea resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfChildArea(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildArea resource from API: %v", err)
	}
	return nil
}
func flattenRouterOspfChildAreaStubType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildAreaAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildAreaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildAreaMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterOspfChildArea(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("stub_type", flattenRouterOspfChildAreaStubType(o["stub_type"], d, "stub_type", sv)); err != nil {
		if !fortiAPIPatch(o["stub_type"]) {
			return fmt.Errorf("Error reading stub_type: %v", err)
		}
	}

	if err = d.Set("authentication", flattenRouterOspfChildAreaAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterOspfChildAreaType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterOspfChildAreaMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterOspfChildAreaStubType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildAreaAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildAreaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildAreaMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfChildArea(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("stub_type"); ok {
		t, err := expandRouterOspfChildAreaStubType(d, v, "stub_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stub_type"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandRouterOspfChildAreaAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandRouterOspfChildAreaType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterOspfChildAreaMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
