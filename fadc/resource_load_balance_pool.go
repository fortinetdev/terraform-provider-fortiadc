// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pool.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePoolRead,
		Update: resourceLoadBalancePoolUpdate,
		Create: resourceLoadBalancePoolCreate,
		Delete: resourceLoadBalancePoolDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_relationship": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_addr_private": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"direct_route_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"direct_route_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"direct_route_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rs_profile": &schema.Schema{
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
func resourceLoadBalancePoolCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePool: type error")
	}

	obj, err := getObjectLoadBalancePool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePool resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalancePool(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePool resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalancePoolRead(d, m)
}
func resourceLoadBalancePoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalancePool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePool resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePool(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePool resource: %v", err)
	}

	return resourceLoadBalancePoolRead(d, m)
}
func resourceLoadBalancePoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalancePool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePool resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePoolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalancePool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePool(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePool resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePoolPoolType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolHealthCheckRelationship(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolSdnConnector(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolSdnAddrPrivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolDirectRouteIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolDirectRouteMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolHealthCheckList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolDirectRouteIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolRsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePool(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("pool_type", flattenLoadBalancePoolPoolType(o["pool_type"], d, "pool_type", sv)); err != nil {
		if !fortiAPIPatch(o["pool_type"]) {
			return fmt.Errorf("Error reading pool_type: %v", err)
		}
	}

	if err = d.Set("health_check", flattenLoadBalancePoolHealthCheck(o["health_check"], d, "health_check", sv)); err != nil {
		if !fortiAPIPatch(o["health_check"]) {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_relationship", flattenLoadBalancePoolHealthCheckRelationship(o["health_check_relationship"], d, "health_check_relationship", sv)); err != nil {
		if !fortiAPIPatch(o["health_check_relationship"]) {
			return fmt.Errorf("Error reading health_check_relationship: %v", err)
		}
	}

	if err = d.Set("service", flattenLoadBalancePoolService(o["service"], d, "service", sv)); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("sdn_connector", flattenLoadBalancePoolSdnConnector(o["sdn_connector"], d, "sdn_connector", sv)); err != nil {
		if !fortiAPIPatch(o["sdn_connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	if err = d.Set("sdn_addr_private", flattenLoadBalancePoolSdnAddrPrivate(o["sdn_addr_private"], d, "sdn_addr_private", sv)); err != nil {
		if !fortiAPIPatch(o["sdn_addr_private"]) {
			return fmt.Errorf("Error reading sdn_addr_private: %v", err)
		}
	}

	if err = d.Set("direct_route_ip", flattenLoadBalancePoolDirectRouteIp(o["direct_route_ip"], d, "direct_route_ip", sv)); err != nil {
		if !fortiAPIPatch(o["direct_route_ip"]) {
			return fmt.Errorf("Error reading direct_route_ip: %v", err)
		}
	}

	if err = d.Set("direct_route_mode", flattenLoadBalancePoolDirectRouteMode(o["direct_route_mode"], d, "direct_route_mode", sv)); err != nil {
		if !fortiAPIPatch(o["direct_route_mode"]) {
			return fmt.Errorf("Error reading direct_route_mode: %v", err)
		}
	}

	if err = d.Set("health_check_list", flattenLoadBalancePoolHealthCheckList(o["health_check_list"], d, "health_check_list", sv)); err != nil {
		if !fortiAPIPatch(o["health_check_list"]) {
			return fmt.Errorf("Error reading health_check_list: %v", err)
		}
	}

	if err = d.Set("direct_route_ip6", flattenLoadBalancePoolDirectRouteIp6(o["direct_route_ip6"], d, "direct_route_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["direct_route_ip6"]) {
			return fmt.Errorf("Error reading direct_route_ip6: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalancePoolType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("rs_profile", flattenLoadBalancePoolRsProfile(o["rs_profile"], d, "rs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["rs_profile"]) {
			return fmt.Errorf("Error reading rs_profile: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePoolMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePoolPoolType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolHealthCheckRelationship(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolSdnConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolSdnAddrPrivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolDirectRouteIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolDirectRouteMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolHealthCheckList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolDirectRouteIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolRsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePool(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pool_type"); ok {
		t, err := expandLoadBalancePoolPoolType(d, v, "pool_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool_type"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok {
		t, err := expandLoadBalancePoolHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_relationship"); ok {
		t, err := expandLoadBalancePoolHealthCheckRelationship(d, v, "health_check_relationship", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check_relationship"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandLoadBalancePoolService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("sdn_connector"); ok {
		t, err := expandLoadBalancePoolSdnConnector(d, v, "sdn_connector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn_connector"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_private"); ok {
		t, err := expandLoadBalancePoolSdnAddrPrivate(d, v, "sdn_addr_private", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn_addr_private"] = t
		}
	}

	if v, ok := d.GetOk("direct_route_ip"); ok {
		t, err := expandLoadBalancePoolDirectRouteIp(d, v, "direct_route_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direct_route_ip"] = t
		}
	}

	if v, ok := d.GetOk("direct_route_mode"); ok {
		t, err := expandLoadBalancePoolDirectRouteMode(d, v, "direct_route_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direct_route_mode"] = t
		}
	}

	if v, ok := d.GetOk("health_check_list"); ok {
		t, err := expandLoadBalancePoolHealthCheckList(d, v, "health_check_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check_list"] = t
		}
	}

	if v, ok := d.GetOk("direct_route_ip6"); ok {
		t, err := expandLoadBalancePoolDirectRouteIp6(d, v, "direct_route_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direct_route_ip6"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalancePoolType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("rs_profile"); ok {
		t, err := expandLoadBalancePoolRsProfile(d, v, "rs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rs_profile"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePoolMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
