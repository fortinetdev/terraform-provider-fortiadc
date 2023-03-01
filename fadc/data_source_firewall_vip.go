// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  firewall vip.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceFirewallVip() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallVipRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mappedip_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mappedport_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"no_nat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extif": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mappedip_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mappedport_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extip": &schema.Schema{
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
func dataSourceFirewallVipRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallVip: type error")
	}

	o, err := c.ReadFirewallVip(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing FirewallVip: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectFirewallVip(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallVip from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenFirewallVipStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipMappedipMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipMappedportMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipNoNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipTrafficGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipExtif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipMappedipMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipMappedportMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallVipMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallVip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenFirewallVipStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mappedip_max", dataSourceFlattenFirewallVipMappedipMax(o["mappedip_max"], d, "mappedip_max")); err != nil {
		if !fortiAPIPatch(o["mappedip_max"]) {
			return fmt.Errorf("Error reading mappedip_max: %v", err)
		}
	}

	if err = d.Set("mappedport_min", dataSourceFlattenFirewallVipMappedportMin(o["mappedport_min"], d, "mappedport_min")); err != nil {
		if !fortiAPIPatch(o["mappedport_min"]) {
			return fmt.Errorf("Error reading mappedport_min: %v", err)
		}
	}

	if err = d.Set("no_nat", dataSourceFlattenFirewallVipNoNat(o["no_nat"], d, "no_nat")); err != nil {
		if !fortiAPIPatch(o["no_nat"]) {
			return fmt.Errorf("Error reading no_nat: %v", err)
		}
	}

	if err = d.Set("proto", dataSourceFlattenFirewallVipProto(o["proto"], d, "proto")); err != nil {
		if !fortiAPIPatch(o["proto"]) {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("traffic_group", dataSourceFlattenFirewallVipTrafficGroup(o["traffic-group"], d, "traffic_group")); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("extif", dataSourceFlattenFirewallVipExtif(o["extif"], d, "extif")); err != nil {
		if !fortiAPIPatch(o["extif"]) {
			return fmt.Errorf("Error reading extif: %v", err)
		}
	}

	if err = d.Set("mappedip_min", dataSourceFlattenFirewallVipMappedipMin(o["mappedip_min"], d, "mappedip_min")); err != nil {
		if !fortiAPIPatch(o["mappedip_min"]) {
			return fmt.Errorf("Error reading mappedip_min: %v", err)
		}
	}

	if err = d.Set("mappedport_max", dataSourceFlattenFirewallVipMappedportMax(o["mappedport_max"], d, "mappedport_max")); err != nil {
		if !fortiAPIPatch(o["mappedport_max"]) {
			return fmt.Errorf("Error reading mappedport_max: %v", err)
		}
	}

	if err = d.Set("extport", dataSourceFlattenFirewallVipExtport(o["extport"], d, "extport")); err != nil {
		if !fortiAPIPatch(o["extport"]) {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("portforward", dataSourceFlattenFirewallVipPortforward(o["portforward"], d, "portforward")); err != nil {
		if !fortiAPIPatch(o["portforward"]) {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("extip", dataSourceFlattenFirewallVipExtip(o["extip"], d, "extip")); err != nil {
		if !fortiAPIPatch(o["extip"]) {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenFirewallVipMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
