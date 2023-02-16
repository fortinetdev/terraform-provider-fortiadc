// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp child ha router id list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterBgpChildHaRouterIdList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBgpChildHaRouterIdListRead,
		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"router_id": &schema.Schema{
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
func dataSourceRouterBgpChildHaRouterIdListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterBgpChildHaRouterIdList: type error")
	}

	o, err := c.ReadRouterBgpChildHaRouterIdList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgpChildHaRouterIdList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterBgpChildHaRouterIdList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgpChildHaRouterIdList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterBgpChildHaRouterIdListNode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildHaRouterIdListRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildHaRouterIdListMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBgpChildHaRouterIdList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("node", dataSourceFlattenRouterBgpChildHaRouterIdListNode(o["node"], d, "node")); err != nil {
		if !fortiAPIPatch(o["node"]) {
			return fmt.Errorf("Error reading node: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterBgpChildHaRouterIdListRouterId(o["router_id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router_id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterBgpChildHaRouterIdListMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
