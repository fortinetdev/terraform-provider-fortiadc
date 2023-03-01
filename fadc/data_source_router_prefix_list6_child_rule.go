// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router prefix list6 child rule.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterPrefixList6ChildRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterPrefixList6ChildRuleRead,
		Schema: map[string]*schema.Schema{
			"le": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ge": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceRouterPrefixList6ChildRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterPrefixList6ChildRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterPrefixList6ChildRule(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterPrefixList6ChildRule from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterPrefixList6ChildRuleLe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPrefixList6ChildRulePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterPrefixList6ChildRuleGe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPrefixList6ChildRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPrefixList6ChildRuleMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterPrefixList6ChildRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("le", dataSourceFlattenRouterPrefixList6ChildRuleLe(o["le"], d, "le")); err != nil {
		if !fortiAPIPatch(o["le"]) {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix6", dataSourceFlattenRouterPrefixList6ChildRulePrefix6(o["prefix6"], d, "prefix6")); err != nil {
		if !fortiAPIPatch(o["prefix6"]) {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	if err = d.Set("ge", dataSourceFlattenRouterPrefixList6ChildRuleGe(o["ge"], d, "ge")); err != nil {
		if !fortiAPIPatch(o["ge"]) {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("action", dataSourceFlattenRouterPrefixList6ChildRuleAction(o["action"], d, "action")); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterPrefixList6ChildRuleMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
