// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance clone pool.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceClonePool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceClonePoolRead,
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
func dataSourceLoadBalanceClonePoolRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceClonePool: type error")
	}

	o, err := c.ReadLoadBalanceClonePool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceClonePool: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceClonePool(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceClonePool from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceClonePoolMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceClonePool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceClonePoolMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
