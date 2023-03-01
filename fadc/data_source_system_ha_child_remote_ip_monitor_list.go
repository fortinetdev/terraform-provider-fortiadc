// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system ha child remote ip monitor list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemHaChildRemoteIpMonitorList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemHaChildRemoteIpMonitorListRead,
		Schema: map[string]*schema.Schema{
			"health_check_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_check_retry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_check_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port": &schema.Schema{
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
func dataSourceSystemHaChildRemoteIpMonitorListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemHaChildRemoteIpMonitorList: type error")
	}

	o, err := c.ReadSystemHaChildRemoteIpMonitorList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemHaChildRemoteIpMonitorList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemHaChildRemoteIpMonitorList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemHaChildRemoteIpMonitorList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemHaChildRemoteIpMonitorListHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaChildRemoteIpMonitorListRemoteAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaChildRemoteIpMonitorListHealthCheckRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaChildRemoteIpMonitorListHealthCheckTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaChildRemoteIpMonitorListSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaChildRemoteIpMonitorListMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemHaChildRemoteIpMonitorList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("health_check_interval", dataSourceFlattenSystemHaChildRemoteIpMonitorListHealthCheckInterval(o["health-check-interval"], d, "health_check_interval")); err != nil {
		if !fortiAPIPatch(o["health-check-interval"]) {
			return fmt.Errorf("Error reading health-check-interval: %v", err)
		}
	}

	if err = d.Set("remote_address", dataSourceFlattenSystemHaChildRemoteIpMonitorListRemoteAddress(o["remote-address"], d, "remote_address")); err != nil {
		if !fortiAPIPatch(o["remote-address"]) {
			return fmt.Errorf("Error reading remote-address: %v", err)
		}
	}

	if err = d.Set("health_check_retry", dataSourceFlattenSystemHaChildRemoteIpMonitorListHealthCheckRetry(o["health-check-retry"], d, "health_check_retry")); err != nil {
		if !fortiAPIPatch(o["health-check-retry"]) {
			return fmt.Errorf("Error reading health-check-retry: %v", err)
		}
	}

	if err = d.Set("health_check_timeout", dataSourceFlattenSystemHaChildRemoteIpMonitorListHealthCheckTimeout(o["health-check-timeout"], d, "health_check_timeout")); err != nil {
		if !fortiAPIPatch(o["health-check-timeout"]) {
			return fmt.Errorf("Error reading health-check-timeout: %v", err)
		}
	}

	if err = d.Set("source_port", dataSourceFlattenSystemHaChildRemoteIpMonitorListSourcePort(o["source-port"], d, "source_port")); err != nil {
		if !fortiAPIPatch(o["source-port"]) {
			return fmt.Errorf("Error reading source-port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemHaChildRemoteIpMonitorListMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
