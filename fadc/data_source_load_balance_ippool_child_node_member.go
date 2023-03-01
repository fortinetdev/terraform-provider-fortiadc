// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance ippool child node member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceIppoolChildNodeMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceIppoolChildNodeMemberRead,
		Schema: map[string]*schema.Schema{
			"ip6_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_node_num": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceLoadBalanceIppoolChildNodeMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	o, err := c.ReadLoadBalanceIppoolChildNodeMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceIppoolChildNodeMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceIppoolChildNodeMemberIp6Max(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberIp6Min(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberHaNodeNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberIpMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberPoolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberIpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolChildNodeMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceIppoolChildNodeMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip6_max", dataSourceFlattenLoadBalanceIppoolChildNodeMemberIp6Max(o["ip6_max"], d, "ip6_max")); err != nil {
		if !fortiAPIPatch(o["ip6_max"]) {
			return fmt.Errorf("Error reading ip6_max: %v", err)
		}
	}

	if err = d.Set("ip6_min", dataSourceFlattenLoadBalanceIppoolChildNodeMemberIp6Min(o["ip6_min"], d, "ip6_min")); err != nil {
		if !fortiAPIPatch(o["ip6_min"]) {
			return fmt.Errorf("Error reading ip6_min: %v", err)
		}
	}

	if err = d.Set("ha_node_num", dataSourceFlattenLoadBalanceIppoolChildNodeMemberHaNodeNum(o["ha_node_num"], d, "ha_node_num")); err != nil {
		if !fortiAPIPatch(o["ha_node_num"]) {
			return fmt.Errorf("Error reading ha_node_num: %v", err)
		}
	}

	if err = d.Set("ip_min", dataSourceFlattenLoadBalanceIppoolChildNodeMemberIpMin(o["ip_min"], d, "ip_min")); err != nil {
		if !fortiAPIPatch(o["ip_min"]) {
			return fmt.Errorf("Error reading ip_min: %v", err)
		}
	}

	if err = d.Set("pool_type", dataSourceFlattenLoadBalanceIppoolChildNodeMemberPoolType(o["pool_type"], d, "pool_type")); err != nil {
		if !fortiAPIPatch(o["pool_type"]) {
			return fmt.Errorf("Error reading pool_type: %v", err)
		}
	}

	if err = d.Set("ip_max", dataSourceFlattenLoadBalanceIppoolChildNodeMemberIpMax(o["ip_max"], d, "ip_max")); err != nil {
		if !fortiAPIPatch(o["ip_max"]) {
			return fmt.Errorf("Error reading ip_max: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenLoadBalanceIppoolChildNodeMemberInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceIppoolChildNodeMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
