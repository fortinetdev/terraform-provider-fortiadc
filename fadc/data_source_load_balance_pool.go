// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pool.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePoolRead,
		Schema: map[string]*schema.Schema{
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_check_relationship": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_addr_private": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"direct_route_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"direct_route_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_check_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"direct_route_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rs_profile": &schema.Schema{
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
func dataSourceLoadBalancePoolRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalancePool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePool: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePool(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePool from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePoolPoolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolHealthCheckRelationship(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolSdnConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolSdnAddrPrivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolDirectRouteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolDirectRouteMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolHealthCheckList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolDirectRouteIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolRsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("pool_type", dataSourceFlattenLoadBalancePoolPoolType(o["pool_type"], d, "pool_type")); err != nil {
		if !fortiAPIPatch(o["pool_type"]) {
			return fmt.Errorf("Error reading pool_type: %v", err)
		}
	}

	if err = d.Set("health_check", dataSourceFlattenLoadBalancePoolHealthCheck(o["health_check"], d, "health_check")); err != nil {
		if !fortiAPIPatch(o["health_check"]) {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_relationship", dataSourceFlattenLoadBalancePoolHealthCheckRelationship(o["health_check_relationship"], d, "health_check_relationship")); err != nil {
		if !fortiAPIPatch(o["health_check_relationship"]) {
			return fmt.Errorf("Error reading health_check_relationship: %v", err)
		}
	}

	if err = d.Set("service", dataSourceFlattenLoadBalancePoolService(o["service"], d, "service")); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("sdn_connector", dataSourceFlattenLoadBalancePoolSdnConnector(o["sdn_connector"], d, "sdn_connector")); err != nil {
		if !fortiAPIPatch(o["sdn_connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	if err = d.Set("sdn_addr_private", dataSourceFlattenLoadBalancePoolSdnAddrPrivate(o["sdn_addr_private"], d, "sdn_addr_private")); err != nil {
		if !fortiAPIPatch(o["sdn_addr_private"]) {
			return fmt.Errorf("Error reading sdn_addr_private: %v", err)
		}
	}

	if err = d.Set("direct_route_ip", dataSourceFlattenLoadBalancePoolDirectRouteIp(o["direct_route_ip"], d, "direct_route_ip")); err != nil {
		if !fortiAPIPatch(o["direct_route_ip"]) {
			return fmt.Errorf("Error reading direct_route_ip: %v", err)
		}
	}

	if err = d.Set("direct_route_mode", dataSourceFlattenLoadBalancePoolDirectRouteMode(o["direct_route_mode"], d, "direct_route_mode")); err != nil {
		if !fortiAPIPatch(o["direct_route_mode"]) {
			return fmt.Errorf("Error reading direct_route_mode: %v", err)
		}
	}

	if err = d.Set("health_check_list", dataSourceFlattenLoadBalancePoolHealthCheckList(o["health_check_list"], d, "health_check_list")); err != nil {
		if !fortiAPIPatch(o["health_check_list"]) {
			return fmt.Errorf("Error reading health_check_list: %v", err)
		}
	}

	if err = d.Set("direct_route_ip6", dataSourceFlattenLoadBalancePoolDirectRouteIp6(o["direct_route_ip6"], d, "direct_route_ip6")); err != nil {
		if !fortiAPIPatch(o["direct_route_ip6"]) {
			return fmt.Errorf("Error reading direct_route_ip6: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalancePoolType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("rs_profile", dataSourceFlattenLoadBalancePoolRsProfile(o["rs_profile"], d, "rs_profile")); err != nil {
		if !fortiAPIPatch(o["rs_profile"]) {
			return fmt.Errorf("Error reading rs_profile: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePoolMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
