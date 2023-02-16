// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router access list6 child rule.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterAccessList6ChildRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterAccessList6ChildRuleRead,
		Update: resourceRouterAccessList6ChildRuleUpdate,
		Create: resourceRouterAccessList6ChildRuleCreate,
		Delete: resourceRouterAccessList6ChildRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceRouterAccessList6ChildRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterAccessList6ChildRule: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterAccessList6ChildRule: type error")
	}

	obj, err := getObjectRouterAccessList6ChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList6ChildRule resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateRouterAccessList6ChildRule(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList6ChildRule resource: %v", err)
	}

	d.SetId(new_id)

	return resourceRouterAccessList6ChildRuleRead(d, m)
}
func resourceRouterAccessList6ChildRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterAccessList6ChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectRouterAccessList6ChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6ChildRule resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterAccessList6ChildRule(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6ChildRule resource: %v", err)
	}

	return resourceRouterAccessList6ChildRuleRead(d, m)
}
func resourceRouterAccessList6ChildRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterAccessList6ChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteRouterAccessList6ChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAccessList6ChildRule resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterAccessList6ChildRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterAccessList6ChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadRouterAccessList6ChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList6ChildRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAccessList6ChildRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList6ChildRule resource from API: %v", err)
	}
	return nil
}
func flattenRouterAccessList6ChildRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessList6ChildRuleMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessList6ChildRulePrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func refreshObjectRouterAccessList6ChildRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("action", flattenRouterAccessList6ChildRuleAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterAccessList6ChildRuleMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterAccessList6ChildRulePrefix6(o["prefix6"], d, "prefix6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix6"]) {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func expandRouterAccessList6ChildRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6ChildRuleMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6ChildRulePrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterAccessList6ChildRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandRouterAccessList6ChildRuleAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterAccessList6ChildRuleMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok {
		t, err := expandRouterAccessList6ChildRulePrefix6(d, v, "prefix6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
