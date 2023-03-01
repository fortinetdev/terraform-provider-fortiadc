// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child area.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterOspfChildArea() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspfChildAreaRead,
		Schema: map[string]*schema.Schema{
			"stub_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
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
func dataSourceRouterOspfChildAreaRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterOspfChildArea: type error")
	}

	o, err := c.ReadRouterOspfChildArea(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspfChildArea: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterOspfChildArea(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspfChildArea from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterOspfChildAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildAreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildAreaMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspfChildArea(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("stub_type", dataSourceFlattenRouterOspfChildAreaStubType(o["stub_type"], d, "stub_type")); err != nil {
		if !fortiAPIPatch(o["stub_type"]) {
			return fmt.Errorf("Error reading stub_type: %v", err)
		}
	}

	if err = d.Set("authentication", dataSourceFlattenRouterOspfChildAreaAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenRouterOspfChildAreaType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterOspfChildAreaMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
