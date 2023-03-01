// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system address6.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemAddress6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAddress6Read,
		Schema: map[string]*schema.Schema{
			"ip6_network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_min": &schema.Schema{
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
func dataSourceSystemAddress6Read(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAddress6: type error")
	}

	o, err := c.ReadSystemAddress6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemAddress6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemAddress6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAddress6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemAddress6Ip6Network(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAddress6Ip6Max(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAddress6Ip6Min(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAddress6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAddress6Mkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAddress6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip6_network", dataSourceFlattenSystemAddress6Ip6Network(o["ip6-network"], d, "ip6_network")); err != nil {
		if !fortiAPIPatch(o["ip6-network"]) {
			return fmt.Errorf("Error reading ip6-network: %v", err)
		}
	}

	if err = d.Set("ip6_max", dataSourceFlattenSystemAddress6Ip6Max(o["ip6-max"], d, "ip6_max")); err != nil {
		if !fortiAPIPatch(o["ip6-max"]) {
			return fmt.Errorf("Error reading ip6-max: %v", err)
		}
	}

	if err = d.Set("ip6_min", dataSourceFlattenSystemAddress6Ip6Min(o["ip6-min"], d, "ip6_min")); err != nil {
		if !fortiAPIPatch(o["ip6-min"]) {
			return fmt.Errorf("Error reading ip6-min: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemAddress6Type(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemAddress6Mkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
