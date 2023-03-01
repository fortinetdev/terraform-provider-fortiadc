// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance auth policy.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceAuthPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceAuthPolicyRead,
		Update: resourceLoadBalanceAuthPolicyUpdate,
		Create: resourceLoadBalanceAuthPolicyCreate,
		Delete: resourceLoadBalanceAuthPolicyDelete,

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
func resourceLoadBalanceAuthPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceAuthPolicy: type error")
	}

	obj, err := getObjectLoadBalanceAuthPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceAuthPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceAuthPolicy(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceAuthPolicy resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceAuthPolicyRead(d, m)
}
func resourceLoadBalanceAuthPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceAuthPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceAuthPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceAuthPolicy(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceAuthPolicy resource: %v", err)
	}

	return resourceLoadBalanceAuthPolicyRead(d, m)
}
func resourceLoadBalanceAuthPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceAuthPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceAuthPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceAuthPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceAuthPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceAuthPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceAuthPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceAuthPolicy resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceAuthPolicyMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceAuthPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenLoadBalanceAuthPolicyMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceAuthPolicyMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceAuthPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceAuthPolicyMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
