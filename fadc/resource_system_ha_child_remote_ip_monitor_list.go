// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system ha child remote ip monitor list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHaChildRemoteIpMonitorList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemHaChildRemoteIpMonitorListRead,
		Update: resourceSystemHaChildRemoteIpMonitorListUpdate,
		Create: resourceSystemHaChildRemoteIpMonitorListCreate,
		Delete: resourceSystemHaChildRemoteIpMonitorListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"health_check_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_retry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceSystemHaChildRemoteIpMonitorListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemHaChildRemoteIpMonitorList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaChildRemoteIpMonitorList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaChildRemoteIpMonitorList(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaChildRemoteIpMonitorList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemHaChildRemoteIpMonitorListRead(d, m)
}
func resourceSystemHaChildRemoteIpMonitorListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHaChildRemoteIpMonitorList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaChildRemoteIpMonitorList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaChildRemoteIpMonitorList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaChildRemoteIpMonitorList resource: %v", err)
	}

	return resourceSystemHaChildRemoteIpMonitorListRead(d, m)
}
func resourceSystemHaChildRemoteIpMonitorListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemHaChildRemoteIpMonitorList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaChildRemoteIpMonitorList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemHaChildRemoteIpMonitorListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemHaChildRemoteIpMonitorList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaChildRemoteIpMonitorList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaChildRemoteIpMonitorList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaChildRemoteIpMonitorList resource from API: %v", err)
	}
	return nil
}
func flattenSystemHaChildRemoteIpMonitorListHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaChildRemoteIpMonitorListRemoteAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaChildRemoteIpMonitorListHealthCheckRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaChildRemoteIpMonitorListHealthCheckTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaChildRemoteIpMonitorListSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaChildRemoteIpMonitorListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHaChildRemoteIpMonitorList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("health_check_interval", flattenSystemHaChildRemoteIpMonitorListHealthCheckInterval(o["health-check-interval"], d, "health_check_interval", sv)); err != nil {
		if !fortiAPIPatch(o["health-check-interval"]) {
			return fmt.Errorf("Error reading health-check-interval: %v", err)
		}
	}

	if err = d.Set("remote_address", flattenSystemHaChildRemoteIpMonitorListRemoteAddress(o["remote-address"], d, "remote_address", sv)); err != nil {
		if !fortiAPIPatch(o["remote-address"]) {
			return fmt.Errorf("Error reading remote-address: %v", err)
		}
	}

	if err = d.Set("health_check_retry", flattenSystemHaChildRemoteIpMonitorListHealthCheckRetry(o["health-check-retry"], d, "health_check_retry", sv)); err != nil {
		if !fortiAPIPatch(o["health-check-retry"]) {
			return fmt.Errorf("Error reading health-check-retry: %v", err)
		}
	}

	if err = d.Set("health_check_timeout", flattenSystemHaChildRemoteIpMonitorListHealthCheckTimeout(o["health-check-timeout"], d, "health_check_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["health-check-timeout"]) {
			return fmt.Errorf("Error reading health-check-timeout: %v", err)
		}
	}

	if err = d.Set("source_port", flattenSystemHaChildRemoteIpMonitorListSourcePort(o["source-port"], d, "source_port", sv)); err != nil {
		if !fortiAPIPatch(o["source-port"]) {
			return fmt.Errorf("Error reading source-port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemHaChildRemoteIpMonitorListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemHaChildRemoteIpMonitorListHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaChildRemoteIpMonitorListRemoteAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaChildRemoteIpMonitorListHealthCheckRetry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaChildRemoteIpMonitorListHealthCheckTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaChildRemoteIpMonitorListSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaChildRemoteIpMonitorListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaChildRemoteIpMonitorList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("health_check_interval"); ok {
		t, err := expandSystemHaChildRemoteIpMonitorListHealthCheckInterval(d, v, "health_check_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("remote_address"); ok {
		t, err := expandSystemHaChildRemoteIpMonitorListRemoteAddress(d, v, "remote_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-address"] = t
		}
	}

	if v, ok := d.GetOk("health_check_retry"); ok {
		t, err := expandSystemHaChildRemoteIpMonitorListHealthCheckRetry(d, v, "health_check_retry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-retry"] = t
		}
	}

	if v, ok := d.GetOk("health_check_timeout"); ok {
		t, err := expandSystemHaChildRemoteIpMonitorListHealthCheckTimeout(d, v, "health_check_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-timeout"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok {
		t, err := expandSystemHaChildRemoteIpMonitorListSourcePort(d, v, "source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemHaChildRemoteIpMonitorListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
