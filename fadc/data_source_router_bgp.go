// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBgpRead,
		Schema: map[string]*schema.Schema{
			"redistribute_ospf": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_connected": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"deterministic_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_connected6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"as": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_static6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"always_compare_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_cmp_routerid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_static": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "RouterBgp"

	o, err := c.ReadRouterBgp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterBgp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterBgpRedistributeOspf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeConnected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeConnected6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeStatic6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeStatic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("redistribute_ospf", dataSourceFlattenRouterBgpRedistributeOspf(o["redistribute-ospf"], d, "redistribute_ospf")); err != nil {
		if !fortiAPIPatch(o["redistribute-ospf"]) {
			return fmt.Errorf("Error reading redistribute-ospf: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterBgpRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router-id: %v", err)
		}
	}

	if err = d.Set("redistribute_connected", dataSourceFlattenRouterBgpRedistributeConnected(o["redistribute-connected"], d, "redistribute_connected")); err != nil {
		if !fortiAPIPatch(o["redistribute-connected"]) {
			return fmt.Errorf("Error reading redistribute-connected: %v", err)
		}
	}

	if err = d.Set("deterministic_med", dataSourceFlattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med")); err != nil {
		if !fortiAPIPatch(o["deterministic-med"]) {
			return fmt.Errorf("Error reading deterministic-med: %v", err)
		}
	}

	if err = d.Set("redistribute_connected6", dataSourceFlattenRouterBgpRedistributeConnected6(o["redistribute-connected6"], d, "redistribute_connected6")); err != nil {
		if !fortiAPIPatch(o["redistribute-connected6"]) {
			return fmt.Errorf("Error reading redistribute-connected6: %v", err)
		}
	}

	if err = d.Set("as", dataSourceFlattenRouterBgpAs(o["as"], d, "as")); err != nil {
		if !fortiAPIPatch(o["as"]) {
			return fmt.Errorf("Error reading as: %v", err)
		}
	}

	if err = d.Set("redistribute_static6", dataSourceFlattenRouterBgpRedistributeStatic6(o["redistribute-static6"], d, "redistribute_static6")); err != nil {
		if !fortiAPIPatch(o["redistribute-static6"]) {
			return fmt.Errorf("Error reading redistribute-static6: %v", err)
		}
	}

	if err = d.Set("always_compare_med", dataSourceFlattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med")); err != nil {
		if !fortiAPIPatch(o["always-compare-med"]) {
			return fmt.Errorf("Error reading always-compare-med: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", dataSourceFlattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid")); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-routerid"]) {
			return fmt.Errorf("Error reading bestpath-cmp-routerid: %v", err)
		}
	}

	if err = d.Set("redistribute_static", dataSourceFlattenRouterBgpRedistributeStatic(o["redistribute-static"], d, "redistribute_static")); err != nil {
		if !fortiAPIPatch(o["redistribute-static"]) {
			return fmt.Errorf("Error reading redistribute-static: %v", err)
		}
	}

	return nil
}
