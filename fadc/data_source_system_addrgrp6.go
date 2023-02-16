// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system addrgrp6.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemAddrgrp6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAddrgrp6Read,
		Schema: map[string]*schema.Schema{
			"member_list": &schema.Schema{
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
		},
	}
}
func dataSourceSystemAddrgrp6Read(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAddrgrp6: type error")
	}

	o, err := c.ReadSystemAddrgrp6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemAddrgrp6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemAddrgrp6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAddrgrp6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemAddrgrp6MemberList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAddrgrp6Mkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAddrgrp6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("member_list", dataSourceFlattenSystemAddrgrp6MemberList(o["member-list"], d, "member_list")); err != nil {
		if !fortiAPIPatch(o["member-list"]) {
			return fmt.Errorf("Error reading member-list: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemAddrgrp6Mkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
