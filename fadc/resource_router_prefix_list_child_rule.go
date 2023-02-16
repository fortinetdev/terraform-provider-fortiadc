// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router prefix list child rule.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterPrefixListChildRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterPrefixListChildRuleRead,
		Update: resourceRouterPrefixListChildRuleUpdate,
		Create: resourceRouterPrefixListChildRuleCreate,
		Delete: resourceRouterPrefixListChildRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"le": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix": &schema.Schema{
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
func resourceRouterPrefixListChildRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixListChildRule: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterPrefixListChildRule: type error")
	}

	obj, err := getObjectRouterPrefixListChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixListChildRule resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateRouterPrefixListChildRule(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixListChildRule resource: %v", err)
	}

	d.SetId(new_id)

	return resourceRouterPrefixListChildRuleRead(d, m)
}
func resourceRouterPrefixListChildRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixListChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectRouterPrefixListChildRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixListChildRule resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterPrefixListChildRule(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixListChildRule resource: %v", err)
	}

	return resourceRouterPrefixListChildRuleRead(d, m)
}
func resourceRouterPrefixListChildRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixListChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteRouterPrefixListChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixListChildRule resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterPrefixListChildRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPrefixListChildRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadRouterPrefixListChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixListChildRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixListChildRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixListChildRule resource from API: %v", err)
	}
	return nil
}
func flattenRouterPrefixListChildRuleLe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixListChildRulePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterPrefixListChildRuleGe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixListChildRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixListChildRuleMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPrefixListChildRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("le", flattenRouterPrefixListChildRuleLe(o["le"], d, "le", sv)); err != nil {
		if !fortiAPIPatch(o["le"]) {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterPrefixListChildRulePrefix(o["prefix"], d, "prefix", sv)); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("ge", flattenRouterPrefixListChildRuleGe(o["ge"], d, "ge", sv)); err != nil {
		if !fortiAPIPatch(o["ge"]) {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("action", flattenRouterPrefixListChildRuleAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterPrefixListChildRuleMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterPrefixListChildRuleLe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListChildRulePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListChildRuleGe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListChildRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListChildRuleMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixListChildRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("le"); ok {
		t, err := expandRouterPrefixListChildRuleLe(d, v, "le", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["le"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok {
		t, err := expandRouterPrefixListChildRulePrefix(d, v, "prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("ge"); ok {
		t, err := expandRouterPrefixListChildRuleGe(d, v, "ge", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ge"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandRouterPrefixListChildRuleAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterPrefixListChildRuleMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
