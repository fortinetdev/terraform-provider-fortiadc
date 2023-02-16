// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router md5 ospf child md5 member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterMd5OspfChildMd5Member() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterMd5OspfChildMd5MemberRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceRouterMd5OspfChildMd5MemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	o, err := c.ReadRouterMd5OspfChildMd5Member(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterMd5OspfChildMd5Member(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterMd5OspfChildMd5MemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMd5OspfChildMd5MemberKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterMd5OspfChildMd5Member(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenRouterMd5OspfChildMd5MemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("key", dataSourceFlattenRouterMd5OspfChildMd5MemberKey(o["key"], d, "key")); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	return nil
}
