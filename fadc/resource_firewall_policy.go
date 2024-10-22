// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure firewall policy.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceFirewallPolicyRead,
		Update: resourceFirewallPolicyUpdate,
		Create: resourceFirewallPolicyUpdate,
		Delete: resourceFirewallPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"deny_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceFirewallPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectFirewallPolicy(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource while getting object: %v", err)
	}

	path := "/api/firewall_policy?mkey=-1"
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource: %v", err)
	}

	d.SetId("FirewallPolicy")
	return resourceFirewallPolicyRead(d, m)
}
func resourceFirewallPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectFirewallPolicy(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource while getting object: %v", err)
	}

	path := "/api/firewall_policy?mkey=-1"
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceFirewallPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/firewall_policy"
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("default_action", flattenFirewallPolicy(o["default-action"], d, "default_action", sv)); err != nil {
		if !fortiAPIPatch(o["default-action"]) {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}
	if err = d.Set("deny_log", flattenFirewallPolicy(o["deny-log"], d, "deny_log", sv)); err != nil {
		if !fortiAPIPatch(o["deny-log"]) {
			return fmt.Errorf("Error reading deny_log: %v", err)
		}
	}

	return nil
}

func expandFirewallPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallPolicy(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_action"); ok {
		if setArgNil {
			obj["default-action"] = nil
		} else {
			t, err := expandFirewallPolicy(d, v, "default_action", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-action"] = t
			}
		}
	}

	if v, ok := d.GetOk("deny_log"); ok {
		if setArgNil {
			obj["deny-log"] = nil
		} else {
			t, err := expandFirewallPolicy(d, v, "deny_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["deny-log"] = t
			}
		}
	}

	return &obj, nil
}
