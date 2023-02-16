// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspfRead,
		Schema: map[string]*schema.Schema{
			"redistribute_connected": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_static_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_connected_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_static_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_static": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_connected_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "RouterOspf"

	o, err := c.ReadRouterOspf(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterOspf(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterOspfRedistributeConnected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeStaticMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeConnectedMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeStaticMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeStatic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeConnectedMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("redistribute_connected", dataSourceFlattenRouterOspfRedistributeConnected(o["redistribute_connected"], d, "redistribute_connected")); err != nil {
		if !fortiAPIPatch(o["redistribute_connected"]) {
			return fmt.Errorf("Error reading redistribute_connected: %v", err)
		}
	}

	if err = d.Set("default_information_metric", dataSourceFlattenRouterOspfDefaultInformationMetric(o["default_information_metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default_information_metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("redistribute_static_metric_type", dataSourceFlattenRouterOspfRedistributeStaticMetricType(o["redistribute_static_metric_type"], d, "redistribute_static_metric_type")); err != nil {
		if !fortiAPIPatch(o["redistribute_static_metric_type"]) {
			return fmt.Errorf("Error reading redistribute_static_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", dataSourceFlattenRouterOspfDefaultInformationMetricType(o["default_information_metric_type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default_information_metric_type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("redistribute_connected_metric", dataSourceFlattenRouterOspfRedistributeConnectedMetric(o["redistribute_connected_metric"], d, "redistribute_connected_metric")); err != nil {
		if !fortiAPIPatch(o["redistribute_connected_metric"]) {
			return fmt.Errorf("Error reading redistribute_connected_metric: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterOspfRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router-id: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterOspfDefaultInformationOriginate(o["default_information_originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default_information_originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("redistribute_static_metric", dataSourceFlattenRouterOspfRedistributeStaticMetric(o["redistribute_static_metric"], d, "redistribute_static_metric")); err != nil {
		if !fortiAPIPatch(o["redistribute_static_metric"]) {
			return fmt.Errorf("Error reading redistribute_static_metric: %v", err)
		}
	}

	if err = d.Set("redistribute_static", dataSourceFlattenRouterOspfRedistributeStatic(o["redistribute_static"], d, "redistribute_static")); err != nil {
		if !fortiAPIPatch(o["redistribute_static"]) {
			return fmt.Errorf("Error reading redistribute_static: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterOspfDefaultMetric(o["default_metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default_metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("redistribute_connected_metric_type", dataSourceFlattenRouterOspfRedistributeConnectedMetricType(o["redistribute_connected_metric_type"], d, "redistribute_connected_metric_type")); err != nil {
		if !fortiAPIPatch(o["redistribute_connected_metric_type"]) {
			return fmt.Errorf("Error reading redistribute_connected_metric_type: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterOspfDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	return nil
}
