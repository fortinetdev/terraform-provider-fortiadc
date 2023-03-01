// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router prefix list6.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterPrefixList6() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterPrefixList6Read,
		Update: resourceRouterPrefixList6Update,
		Create: resourceRouterPrefixList6Create,
		Delete: resourceRouterPrefixList6Delete,

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
func resourceRouterPrefixList6Create(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixList6: type error")
	}

	obj, err := getObjectRouterPrefixList6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6 resource while getting object: %v", err)
	}

	_, err = c.CreateRouterPrefixList6(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6 resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterPrefixList6Read(d, m)
}
func resourceRouterPrefixList6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterPrefixList6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterPrefixList6(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6 resource: %v", err)
	}

	return resourceRouterPrefixList6Read(d, m)
}
func resourceRouterPrefixList6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterPrefixList6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixList6 resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterPrefixList6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterPrefixList6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixList6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6 resource from API: %v", err)
	}
	return nil
}
func flattenRouterPrefixList6Description(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6Mkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPrefixList6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenRouterPrefixList6Description(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterPrefixList6Mkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterPrefixList6Description(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6Mkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixList6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandRouterPrefixList6Description(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterPrefixList6Mkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
