// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp child neighbor.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterBgpChildNeighbor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBgpChildNeighborRead,
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_list_out6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_list_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_list_in6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_list_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_list_out6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_list_in": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_list_in": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_list_in6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_as": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"holdtime": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceRouterBgpChildNeighborRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterBgpChildNeighbor: type error")
	}

	o, err := c.ReadRouterBgpChildNeighbor(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgpChildNeighbor: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterBgpChildNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgpChildNeighbor from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterBgpChildNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpChildNeighborKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBgpChildNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", dataSourceFlattenRouterBgpChildNeighborIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenRouterBgpChildNeighborWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", dataSourceFlattenRouterBgpChildNeighborDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("Error reading distribute-list-out6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", dataSourceFlattenRouterBgpChildNeighborDistributeListOut(o["distribute-list-out"], d, "distribute_list_out")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("Error reading distribute-list-out: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", dataSourceFlattenRouterBgpChildNeighborPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("Error reading prefix-list-in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", dataSourceFlattenRouterBgpChildNeighborPrefixListOut(o["prefix-list-out"], d, "prefix_list_out")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("Error reading prefix-list-out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", dataSourceFlattenRouterBgpChildNeighborPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("Error reading prefix-list-out6: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", dataSourceFlattenRouterBgpChildNeighborPrefixListIn(o["prefix-list-in"], d, "prefix_list_in")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("Error reading prefix-list-in: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenRouterBgpChildNeighborType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", dataSourceFlattenRouterBgpChildNeighborDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute-list-in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", dataSourceFlattenRouterBgpChildNeighborDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("Error reading distribute-list-in6: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterBgpChildNeighborInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip6", dataSourceFlattenRouterBgpChildNeighborIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("remote_as", dataSourceFlattenRouterBgpChildNeighborRemoteAs(o["remote-as"], d, "remote_as")); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("Error reading remote-as: %v", err)
		}
	}

	if err = d.Set("holdtime", dataSourceFlattenRouterBgpChildNeighborHoldtime(o["holdtime"], d, "holdtime")); err != nil {
		if !fortiAPIPatch(o["holdtime"]) {
			return fmt.Errorf("Error reading holdtime: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenRouterBgpChildNeighborPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterBgpChildNeighborMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("keepalive", dataSourceFlattenRouterBgpChildNeighborKeepalive(o["keepalive"], d, "keepalive")); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	return nil
}
