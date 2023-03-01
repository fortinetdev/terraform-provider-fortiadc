// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router prefix list6 child rule.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterPrefixList6ChildRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterPrefixList6ChildRuleRead,
		Update: resourceRouterPrefixList6ChildRuleUpdate,
		Create: resourceRouterPrefixList6ChildRuleCreate,
		Delete: resourceRouterPrefixList6ChildRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"le": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ge": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceRouterPrefixList6ChildRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule: type error")
	}

	obj, err := getObjectRouterPrefixList6ChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6ChildRule resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateRouterPrefixList6ChildRule(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6ChildRule resource: %v", err)
	}

	d.SetId(new_id)

	return resourceRouterPrefixList6ChildRuleRead(d, m)
}
func resourceRouterPrefixList6ChildRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectRouterPrefixList6ChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6ChildRule resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterPrefixList6ChildRule(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6ChildRule resource: %v", err)
	}

	return resourceRouterPrefixList6ChildRuleRead(d, m)
}
func resourceRouterPrefixList6ChildRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteRouterPrefixList6ChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixList6ChildRule resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterPrefixList6ChildRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadRouterPrefixList6ChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6ChildRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixList6ChildRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6ChildRule resource from API: %v", err)
	}
	return nil
}
func flattenRouterPrefixList6ChildRuleLe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6ChildRulePrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterPrefixList6ChildRuleGe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6ChildRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6ChildRuleMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPrefixList6ChildRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("le", flattenRouterPrefixList6ChildRuleLe(o["le"], d, "le", sv)); err != nil {
		if !fortiAPIPatch(o["le"]) {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterPrefixList6ChildRulePrefix6(o["prefix6"], d, "prefix6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix6"]) {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	if err = d.Set("ge", flattenRouterPrefixList6ChildRuleGe(o["ge"], d, "ge", sv)); err != nil {
		if !fortiAPIPatch(o["ge"]) {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("action", flattenRouterPrefixList6ChildRuleAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterPrefixList6ChildRuleMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterPrefixList6ChildRuleLe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6ChildRulePrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6ChildRuleGe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6ChildRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6ChildRuleMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixList6ChildRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("le"); ok {
		t, err := expandRouterPrefixList6ChildRuleLe(d, v, "le", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["le"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok {
		t, err := expandRouterPrefixList6ChildRulePrefix6(d, v, "prefix6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	if v, ok := d.GetOk("ge"); ok {
		t, err := expandRouterPrefixList6ChildRuleGe(d, v, "ge", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ge"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandRouterPrefixList6ChildRuleAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterPrefixList6ChildRuleMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
