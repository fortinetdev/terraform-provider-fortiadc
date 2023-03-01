// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterStaticRead,
		Schema: map[string]*schema.Schema{
			"gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dest": &schema.Schema{
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
func dataSourceRouterStaticRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterStatic: type error")
	}

	o, err := c.ReadRouterStatic(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterStatic: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterStatic(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterStatic from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterStaticGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterStaticMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterStatic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("gw", dataSourceFlattenRouterStaticGw(o["gw"], d, "gw")); err != nil {
		if !fortiAPIPatch(o["gw"]) {
			return fmt.Errorf("Error reading gw: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterStaticDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dest", dataSourceFlattenRouterStaticDest(o["dest"], d, "dest")); err != nil {
		if !fortiAPIPatch(o["dest"]) {
			return fmt.Errorf("Error reading dest: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterStaticMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
