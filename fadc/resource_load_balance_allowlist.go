// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance allowlist.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceAllowlist() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceAllowlistRead,
		Update: resourceLoadBalanceAllowlistUpdate,
		Create: resourceLoadBalanceAllowlistCreate,
		Delete: resourceLoadBalanceAllowlistDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
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
func resourceLoadBalanceAllowlistCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceAllowlist: type error")
	}

	obj, err := getObjectLoadBalanceAllowlist(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceAllowlist resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceAllowlist(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceAllowlist resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceAllowlistRead(d, m)
}
func resourceLoadBalanceAllowlistUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceAllowlist(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceAllowlist resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceAllowlist(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceAllowlist resource: %v", err)
	}

	return resourceLoadBalanceAllowlistRead(d, m)
}
func resourceLoadBalanceAllowlistDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceAllowlist(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceAllowlist resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceAllowlistRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceAllowlist(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceAllowlist resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceAllowlist(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceAllowlist resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceAllowlistStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceAllowlistDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceAllowlistMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceAllowlist(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLoadBalanceAllowlistStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("description", flattenLoadBalanceAllowlistDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceAllowlistMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceAllowlistStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceAllowlistDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceAllowlistMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceAllowlist(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLoadBalanceAllowlistStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandLoadBalanceAllowlistDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceAllowlistMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
