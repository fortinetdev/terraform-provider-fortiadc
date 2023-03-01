// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterBgpRead,
		Update: resourceRouterBgpUpdate,
		Create: resourceRouterBgpUpdate,
		Delete: resourceRouterBgpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"redistribute_ospf": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_connected": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"deterministic_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_connected6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"as": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_static6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"always_compare_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"bestpath_cmp_routerid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_static": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceRouterBgpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterBgp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource: %v", err)
	}

	d.SetId("RouterBgp")
	return resourceRouterBgpRead(d, m)
}
func resourceRouterBgpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterBgp(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing RouterBgp resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterBgp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgp resource from API: %v", err)
	}
	return nil
}
func flattenRouterBgpRedistributeOspf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeConnected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeConnected6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeStatic6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeStatic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("redistribute_ospf", flattenRouterBgpRedistributeOspf(o["redistribute-ospf"], d, "redistribute_ospf", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-ospf"]) {
			return fmt.Errorf("Error reading redistribute-ospf: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterBgpRouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router-id: %v", err)
		}
	}

	if err = d.Set("redistribute_connected", flattenRouterBgpRedistributeConnected(o["redistribute-connected"], d, "redistribute_connected", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-connected"]) {
			return fmt.Errorf("Error reading redistribute-connected: %v", err)
		}
	}

	if err = d.Set("deterministic_med", flattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med", sv)); err != nil {
		if !fortiAPIPatch(o["deterministic-med"]) {
			return fmt.Errorf("Error reading deterministic-med: %v", err)
		}
	}

	if err = d.Set("redistribute_connected6", flattenRouterBgpRedistributeConnected6(o["redistribute-connected6"], d, "redistribute_connected6", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-connected6"]) {
			return fmt.Errorf("Error reading redistribute-connected6: %v", err)
		}
	}

	if err = d.Set("as", flattenRouterBgpAs(o["as"], d, "as", sv)); err != nil {
		if !fortiAPIPatch(o["as"]) {
			return fmt.Errorf("Error reading as: %v", err)
		}
	}

	if err = d.Set("redistribute_static6", flattenRouterBgpRedistributeStatic6(o["redistribute-static6"], d, "redistribute_static6", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-static6"]) {
			return fmt.Errorf("Error reading redistribute-static6: %v", err)
		}
	}

	if err = d.Set("always_compare_med", flattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med", sv)); err != nil {
		if !fortiAPIPatch(o["always-compare-med"]) {
			return fmt.Errorf("Error reading always-compare-med: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", flattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-routerid"]) {
			return fmt.Errorf("Error reading bestpath-cmp-routerid: %v", err)
		}
	}

	if err = d.Set("redistribute_static", flattenRouterBgpRedistributeStatic(o["redistribute-static"], d, "redistribute_static", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-static"]) {
			return fmt.Errorf("Error reading redistribute-static: %v", err)
		}
	}

	return nil
}

func expandRouterBgpRedistributeOspf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeConnected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDeterministicMed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeConnected6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeStatic6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAlwaysCompareMed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpRouterid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeStatic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("redistribute_ospf"); ok {
		if setArgNil {
			obj["redistribute_ospf"] = nil
		} else {
			t, err := expandRouterBgpRedistributeOspf(d, v, "redistribute_ospf", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-ospf"] = t
			}
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		if setArgNil {
			obj["router_id"] = nil
		} else {
			t, err := expandRouterBgpRouterId(d, v, "router_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_connected"); ok {
		if setArgNil {
			obj["redistribute_connected"] = nil
		} else {
			t, err := expandRouterBgpRedistributeConnected(d, v, "redistribute_connected", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-connected"] = t
			}
		}
	}

	if v, ok := d.GetOk("deterministic_med"); ok {
		if setArgNil {
			obj["deterministic_med"] = nil
		} else {
			t, err := expandRouterBgpDeterministicMed(d, v, "deterministic_med", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["deterministic-med"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_connected6"); ok {
		if setArgNil {
			obj["redistribute_connected6"] = nil
		} else {
			t, err := expandRouterBgpRedistributeConnected6(d, v, "redistribute_connected6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-connected6"] = t
			}
		}
	}

	if v, ok := d.GetOk("as"); ok {
		if setArgNil {
			obj["as"] = nil
		} else {
			t, err := expandRouterBgpAs(d, v, "as", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["as"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_static6"); ok {
		if setArgNil {
			obj["redistribute_static6"] = nil
		} else {
			t, err := expandRouterBgpRedistributeStatic6(d, v, "redistribute_static6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-static6"] = t
			}
		}
	}

	if v, ok := d.GetOk("always_compare_med"); ok {
		if setArgNil {
			obj["always_compare_med"] = nil
		} else {
			t, err := expandRouterBgpAlwaysCompareMed(d, v, "always_compare_med", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["always-compare-med"] = t
			}
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_routerid"); ok {
		if setArgNil {
			obj["bestpath_cmp_routerid"] = nil
		} else {
			t, err := expandRouterBgpBestpathCmpRouterid(d, v, "bestpath_cmp_routerid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bestpath-cmp-routerid"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_static"); ok {
		if setArgNil {
			obj["redistribute_static"] = nil
		} else {
			t, err := expandRouterBgpRedistributeStatic(d, v, "redistribute_static", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-static"] = t
			}
		}
	}

	return &obj, nil
}
