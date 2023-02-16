// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router access list6.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterAccessList6() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterAccessList6Read,
		Update: resourceRouterAccessList6Update,
		Create: resourceRouterAccessList6Create,
		Delete: resourceRouterAccessList6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
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
func resourceRouterAccessList6Create(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterAccessList6: type error")
	}

	obj, err := getObjectRouterAccessList6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList6 resource while getting object: %v", err)
	}

	_, err = c.CreateRouterAccessList6(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList6 resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterAccessList6Read(d, m)
}
func resourceRouterAccessList6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterAccessList6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterAccessList6(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6 resource: %v", err)
	}

	return resourceRouterAccessList6Read(d, m)
}
func resourceRouterAccessList6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterAccessList6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAccessList6 resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterAccessList6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterAccessList6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAccessList6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList6 resource from API: %v", err)
	}
	return nil
}
func flattenRouterAccessList6Description(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessList6Mkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterAccessList6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenRouterAccessList6Description(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterAccessList6Mkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterAccessList6Description(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6Mkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterAccessList6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandRouterAccessList6Description(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterAccessList6Mkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
