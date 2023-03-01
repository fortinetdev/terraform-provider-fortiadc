// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance web filter profile child category members.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceWebFilterProfileChildCategoryMembers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceWebFilterProfileChildCategoryMembersRead,
		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
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
func dataSourceLoadBalanceWebFilterProfileChildCategoryMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	o, err := c.ReadLoadBalanceWebFilterProfileChildCategoryMembers(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceWebFilterProfileChildCategoryMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceWebFilterProfileChildCategoryMembersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceWebFilterProfileChildCategoryMembersMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceWebFilterProfileChildCategoryMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", dataSourceFlattenLoadBalanceWebFilterProfileChildCategoryMembersCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceWebFilterProfileChildCategoryMembersMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
