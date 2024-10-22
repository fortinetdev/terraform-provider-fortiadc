// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure firewall policy child rule.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallPolicyChildRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceFirewallPolicyChildRuleRead,
		Update: resourceFirewallPolicyChildRuleUpdate,
		Create: resourceFirewallPolicyChildRuleCreate,
		Delete: resourceFirewallPolicyChildRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"in_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"out_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_address_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_external_resource_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination_address_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination_external_resource_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
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

func resourceFirewallPolicyChildRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallPolicyChildRule: type error")
	}

	obj, err := getObjectFirewallPolicyChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicyChildRule resource while getting object: %v", err)
	}

	path := "/api/firewall_policy_child_rule"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicyChildRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceFirewallPolicyChildRuleRead(d, m)
}

func resourceFirewallPolicyChildRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectFirewallPolicyChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicyChildRule resource while getting object: %v", err)
	}

	path := "/api/firewall_policy_child_rule?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicyChildRule resource: %v", err)
	}

	return resourceFirewallPolicyChildRuleRead(d, m)
}

func resourceFirewallPolicyChildRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/firewall_policy_child_rule?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallPolicyChildRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicyChildRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/firewall_policy_child_rule?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicyChildRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicyChildRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicyChildRule resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicyChildRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallPolicyChildRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenFirewallPolicyChildRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("in_interface", flattenFirewallPolicyChildRule(o["in-interface"], d, "in_interface", sv)); err != nil {
		if !fortiAPIPatch(o["in-interface"]) {
			return fmt.Errorf("Error reading in_interface: %v", err)
		}
	}

	if err = d.Set("out_interface", flattenFirewallPolicyChildRule(o["out-interface"], d, "out_interface", sv)); err != nil {
		if !fortiAPIPatch(o["out-interface"]) {
			return fmt.Errorf("Error reading out_interface: %v", err)
		}
	}

	if err = d.Set("source_type", flattenFirewallPolicyChildRule(o["source-type"], d, "source_type", sv)); err != nil {
		if !fortiAPIPatch(o["source-type"]) {
			return fmt.Errorf("Error reading source_type: %v", err)
		}
	}

	if err = d.Set("source_address", flattenFirewallPolicyChildRule(o["source-address"], d, "source_address", sv)); err != nil {
		if !fortiAPIPatch(o["source-address"]) {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	if err = d.Set("source_address_group", flattenFirewallPolicyChildRule(o["source-address-group"], d, "source_address_group", sv)); err != nil {
		if !fortiAPIPatch(o["source-address-group"]) {
			return fmt.Errorf("Error reading source_address_group: %v", err)
		}
	}

	if err = d.Set("source_external_resource_address", flattenFirewallPolicyChildRule(o["source-external-resource-address"], d, "source_external_resource_address", sv)); err != nil {
		if !fortiAPIPatch(o["source-external-resource-address"]) {
			return fmt.Errorf("Error reading source_external_resource_address: %v", err)
		}
	}

	if err = d.Set("destination_type", flattenFirewallPolicyChildRule(o["destination-type"], d, "destination_type", sv)); err != nil {
		if !fortiAPIPatch(o["destination-type"]) {
			return fmt.Errorf("Error reading destination_type: %v", err)
		}
	}

	if err = d.Set("destination_address", flattenFirewallPolicyChildRule(o["destination-address"], d, "destination_address", sv)); err != nil {
		if !fortiAPIPatch(o["destination-address"]) {
			return fmt.Errorf("Error reading destination_address: %v", err)
		}
	}

	if err = d.Set("destination_address_group", flattenFirewallPolicyChildRule(o["destination-address-group"], d, "destination_address_group", sv)); err != nil {
		if !fortiAPIPatch(o["destination-address-group"]) {
			return fmt.Errorf("Error reading destination_address_group: %v", err)
		}
	}

	if err = d.Set("destination_external_resource_address", flattenFirewallPolicyChildRule(o["destination-external-resource-address"], d, "destination_external_resource_address", sv)); err != nil {
		if !fortiAPIPatch(o["destination-external-resource-address"]) {
			return fmt.Errorf("Error reading destination_external_resource_address: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallPolicyChildRule(o["service"], d, "service", sv)); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallPolicyChildRule(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallPolicyChildRule(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("deny_log", flattenFirewallPolicyChildRule(o["deny-log"], d, "deny_log", sv)); err != nil {
		if !fortiAPIPatch(o["deny-log"]) {
			return fmt.Errorf("Error reading deny_log: %v", err)
		}
	}

	return nil
}

func expandFirewallPolicyChildRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallPolicyChildRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("in_interface"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "in_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["in-interface"] = t
		}
	}

	if v, ok := d.GetOk("out_interface"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "out_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-interface"] = t
		}
	}

	if v, ok := d.GetOk("source_type"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "source_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-type"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "source_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	if v, ok := d.GetOk("source_address_group"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "source_address_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address-group"] = t
		}
	}

	if v, ok := d.GetOk("source_external_resource_address"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "source_external_resource_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-external-resource-address"] = t
		}
	}

	if v, ok := d.GetOk("destination_type"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "destination_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-type"] = t
		}
	}

	if v, ok := d.GetOk("destination_address"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "destination_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-address"] = t
		}
	}

	if v, ok := d.GetOk("destination_address_group"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "destination_address_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-address-group"] = t
		}
	}

	if v, ok := d.GetOk("destination_external_resource_address"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "destination_external_resource_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-external-resource-address"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("deny_log"); ok {
		t, err := expandFirewallPolicyChildRule(d, v, "deny_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny-log"] = t
		}
	}

	return &obj, nil
}
