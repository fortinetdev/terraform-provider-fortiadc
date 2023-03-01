// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content routing.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceContentRouting() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceContentRoutingRead,
		Schema: map[string]*schema.Schema{
			"source_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"packet_fwd_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_pool_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"persistence_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"method_inherit": &schema.Schema{
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
func dataSourceLoadBalanceContentRoutingRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRouting: type error")
	}

	o, err := c.ReadLoadBalanceContentRouting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceContentRouting: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceContentRouting(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceContentRouting from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceContentRoutingSourcePoolList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingScheduleList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingConnectionPool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenLoadBalanceContentRoutingPacketFwdMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingConnectionPoolInherit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenLoadBalanceContentRoutingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingPool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingPersistenceInherit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingSchedulePoolList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingMethodInherit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceContentRouting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("source_pool_list", dataSourceFlattenLoadBalanceContentRoutingSourcePoolList(o["source-pool-list"], d, "source_pool_list")); err != nil {
		if !fortiAPIPatch(o["source-pool-list"]) {
			return fmt.Errorf("Error reading source-pool-list: %v", err)
		}
	}

	if err = d.Set("schedule_list", dataSourceFlattenLoadBalanceContentRoutingScheduleList(o["schedule-list"], d, "schedule_list")); err != nil {
		if !fortiAPIPatch(o["schedule-list"]) {
			return fmt.Errorf("Error reading schedule-list: %v", err)
		}
	}

	if err = d.Set("persistence", dataSourceFlattenLoadBalanceContentRoutingPersistence(o["persistence"], d, "persistence")); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("connection_pool", dataSourceFlattenLoadBalanceContentRoutingConnectionPool(o["connection-pool"], d, "connection_pool")); err != nil {
		if !fortiAPIPatch(o["connection-pool"]) {
			return fmt.Errorf("Error reading connection-pool: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenLoadBalanceContentRoutingIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("packet_fwd_method", dataSourceFlattenLoadBalanceContentRoutingPacketFwdMethod(o["packet-fwd-method"], d, "packet_fwd_method")); err != nil {
		if !fortiAPIPatch(o["packet-fwd-method"]) {
			return fmt.Errorf("Error reading packet-fwd-method: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenLoadBalanceContentRoutingComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("connection_pool_inherit", dataSourceFlattenLoadBalanceContentRoutingConnectionPoolInherit(o["connection_pool_inherit"], d, "connection_pool_inherit")); err != nil {
		if !fortiAPIPatch(o["connection_pool_inherit"]) {
			return fmt.Errorf("Error reading connection_pool_inherit: %v", err)
		}
	}

	if err = d.Set("ip6", dataSourceFlattenLoadBalanceContentRoutingIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceContentRoutingType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("pool", dataSourceFlattenLoadBalanceContentRoutingPool(o["pool"], d, "pool")); err != nil {
		if !fortiAPIPatch(o["pool"]) {
			return fmt.Errorf("Error reading pool: %v", err)
		}
	}

	if err = d.Set("persistence_inherit", dataSourceFlattenLoadBalanceContentRoutingPersistenceInherit(o["persistence_inherit"], d, "persistence_inherit")); err != nil {
		if !fortiAPIPatch(o["persistence_inherit"]) {
			return fmt.Errorf("Error reading persistence_inherit: %v", err)
		}
	}

	if err = d.Set("schedule_pool_list", dataSourceFlattenLoadBalanceContentRoutingSchedulePoolList(o["schedule-pool-list"], d, "schedule_pool_list")); err != nil {
		if !fortiAPIPatch(o["schedule-pool-list"]) {
			return fmt.Errorf("Error reading schedule-pool-list: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceContentRoutingMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("method", dataSourceFlattenLoadBalanceContentRoutingMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("method_inherit", dataSourceFlattenLoadBalanceContentRoutingMethodInherit(o["method_inherit"], d, "method_inherit")); err != nil {
		if !fortiAPIPatch(o["method_inherit"]) {
			return fmt.Errorf("Error reading method_inherit: %v", err)
		}
	}

	return nil
}
