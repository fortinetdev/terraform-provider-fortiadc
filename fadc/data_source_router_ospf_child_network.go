// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child network.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterOspfChildNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspfChildNetworkRead,
		Schema: map[string]*schema.Schema{
			"area_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix": &schema.Schema{
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
func dataSourceRouterOspfChildNetworkRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterOspfChildNetwork: type error")
	}

	o, err := c.ReadRouterOspfChildNetwork(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspfChildNetwork: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterOspfChildNetwork(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspfChildNetwork from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterOspfChildNetworkAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfChildNetworkMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspfChildNetwork(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("area_id", dataSourceFlattenRouterOspfChildNetworkAreaId(o["area_id"], d, "area_id")); err != nil {
		if !fortiAPIPatch(o["area_id"]) {
			return fmt.Errorf("Error reading area_id: %v", err)
		}
	}

	if err = d.Set("prefix", dataSourceFlattenRouterOspfChildNetworkPrefix(o["prefix"], d, "prefix")); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterOspfChildNetworkMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
