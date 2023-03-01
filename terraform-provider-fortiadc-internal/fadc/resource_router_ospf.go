// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterOspfRead,
		Update: resourceRouterOspfUpdate,
		Create: resourceRouterOspfUpdate,
		Delete: resourceRouterOspfDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"redistribute_connected": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_information_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_static_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_information_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_connected_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_static_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_static": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redistribute_connected_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceRouterOspfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterOspf(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource: %v", err)
	}

	d.SetId("RouterOspf")
	return resourceRouterOspfRead(d, m)
}
func resourceRouterOspfDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterOspf(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing RouterOspf resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterOspf(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource from API: %v", err)
	}
	return nil
}
func flattenRouterOspfRedistributeConnected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeStaticMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeConnectedMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeStaticMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeStatic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeConnectedMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("redistribute_connected", flattenRouterOspfRedistributeConnected(o["redistribute_connected"], d, "redistribute_connected", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute_connected"]) {
			return fmt.Errorf("Error reading redistribute_connected: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspfDefaultInformationMetric(o["default_information_metric"], d, "default_information_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default_information_metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("redistribute_static_metric_type", flattenRouterOspfRedistributeStaticMetricType(o["redistribute_static_metric_type"], d, "redistribute_static_metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute_static_metric_type"]) {
			return fmt.Errorf("Error reading redistribute_static_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspfDefaultInformationMetricType(o["default_information_metric_type"], d, "default_information_metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["default_information_metric_type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("redistribute_connected_metric", flattenRouterOspfRedistributeConnectedMetric(o["redistribute_connected_metric"], d, "redistribute_connected_metric", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute_connected_metric"]) {
			return fmt.Errorf("Error reading redistribute_connected_metric: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspfRouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router-id: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspfDefaultInformationOriginate(o["default_information_originate"], d, "default_information_originate", sv)); err != nil {
		if !fortiAPIPatch(o["default_information_originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("redistribute_static_metric", flattenRouterOspfRedistributeStaticMetric(o["redistribute_static_metric"], d, "redistribute_static_metric", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute_static_metric"]) {
			return fmt.Errorf("Error reading redistribute_static_metric: %v", err)
		}
	}

	if err = d.Set("redistribute_static", flattenRouterOspfRedistributeStatic(o["redistribute_static"], d, "redistribute_static", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute_static"]) {
			return fmt.Errorf("Error reading redistribute_static: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspfDefaultMetric(o["default_metric"], d, "default_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default_metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("redistribute_connected_metric_type", flattenRouterOspfRedistributeConnectedMetricType(o["redistribute_connected_metric_type"], d, "redistribute_connected_metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute_connected_metric_type"]) {
			return fmt.Errorf("Error reading redistribute_connected_metric_type: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterOspfDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	return nil
}

func expandRouterOspfRedistributeConnected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeStaticMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeConnectedMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeStaticMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeStatic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeConnectedMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("redistribute_connected"); ok {
		if setArgNil {
			obj["redistribute_connected"] = nil
		} else {
			t, err := expandRouterOspfRedistributeConnected(d, v, "redistribute_connected", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute_connected"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok {
		if setArgNil {
			obj["default_information_metric"] = nil
		} else {
			t, err := expandRouterOspfDefaultInformationMetric(d, v, "default_information_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default_information_metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_static_metric_type"); ok {
		if setArgNil {
			obj["redistribute_static_metric_type"] = nil
		} else {
			t, err := expandRouterOspfRedistributeStaticMetricType(d, v, "redistribute_static_metric_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute_static_metric_type"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_metric_type"); ok {
		if setArgNil {
			obj["default_information_metric_type"] = nil
		} else {
			t, err := expandRouterOspfDefaultInformationMetricType(d, v, "default_information_metric_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default_information_metric_type"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_connected_metric"); ok {
		if setArgNil {
			obj["redistribute_connected_metric"] = nil
		} else {
			t, err := expandRouterOspfRedistributeConnectedMetric(d, v, "redistribute_connected_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute_connected_metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		if setArgNil {
			obj["router_id"] = nil
		} else {
			t, err := expandRouterOspfRouterId(d, v, "router_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok {
		if setArgNil {
			obj["default_information_originate"] = nil
		} else {
			t, err := expandRouterOspfDefaultInformationOriginate(d, v, "default_information_originate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default_information_originate"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_static_metric"); ok {
		if setArgNil {
			obj["redistribute_static_metric"] = nil
		} else {
			t, err := expandRouterOspfRedistributeStaticMetric(d, v, "redistribute_static_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute_static_metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_static"); ok {
		if setArgNil {
			obj["redistribute_static"] = nil
		} else {
			t, err := expandRouterOspfRedistributeStatic(d, v, "redistribute_static", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute_static"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_metric"); ok {
		if setArgNil {
			obj["default_metric"] = nil
		} else {
			t, err := expandRouterOspfDefaultMetric(d, v, "default_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default_metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_connected_metric_type"); ok {
		if setArgNil {
			obj["redistribute_connected_metric_type"] = nil
		} else {
			t, err := expandRouterOspfRedistributeConnectedMetricType(d, v, "redistribute_connected_metric_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute_connected_metric_type"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		if setArgNil {
			obj["distance"] = nil
		} else {
			t, err := expandRouterOspfDistance(d, v, "distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance"] = t
			}
		}
	}

	return &obj, nil
}
