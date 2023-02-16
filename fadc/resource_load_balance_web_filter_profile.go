// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance web filter profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceWebFilterProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceWebFilterProfileRead,
		Update: resourceLoadBalanceWebFilterProfileUpdate,
		Create: resourceLoadBalanceWebFilterProfileCreate,
		Delete: resourceLoadBalanceWebFilterProfileDelete,

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
func resourceLoadBalanceWebFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfile: type error")
	}

	obj, err := getObjectLoadBalanceWebFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebFilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceWebFilterProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebFilterProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceWebFilterProfileRead(d, m)
}
func resourceLoadBalanceWebFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceWebFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebFilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceWebFilterProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebFilterProfile resource: %v", err)
	}

	return resourceLoadBalanceWebFilterProfileRead(d, m)
}
func resourceLoadBalanceWebFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceWebFilterProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceWebFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceWebFilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceWebFilterProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceWebFilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebFilterProfile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceWebFilterProfileDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceWebFilterProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceWebFilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenLoadBalanceWebFilterProfileDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceWebFilterProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceWebFilterProfileDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceWebFilterProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceWebFilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandLoadBalanceWebFilterProfileDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceWebFilterProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
