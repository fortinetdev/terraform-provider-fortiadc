// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance ippool.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceIppool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceIppoolRead,
		Schema: map[string]*schema.Schema{
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_start": &schema.Schema{
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
			"ip6_start": &schema.Schema{
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
func dataSourceLoadBalanceIppoolRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppool: type error")
	}

	o, err := c.ReadLoadBalanceIppool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceIppool: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceIppool(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceIppool from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceIppoolPoolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolIpEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolIp6End(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolIpStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceIppoolIp6Start(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceIppool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("pool_type", dataSourceFlattenLoadBalanceIppoolPoolType(o["pool_type"], d, "pool_type")); err != nil {
		if !fortiAPIPatch(o["pool_type"]) {
			return fmt.Errorf("Error reading pool_type: %v", err)
		}
	}

	if err = d.Set("ip_end", dataSourceFlattenLoadBalanceIppoolIpEnd(o["ip-end"], d, "ip_end")); err != nil {
		if !fortiAPIPatch(o["ip-end"]) {
			return fmt.Errorf("Error reading ip-end: %v", err)
		}
	}

	if err = d.Set("ip6_end", dataSourceFlattenLoadBalanceIppoolIp6End(o["ip6-end"], d, "ip6_end")); err != nil {
		if !fortiAPIPatch(o["ip6-end"]) {
			return fmt.Errorf("Error reading ip6-end: %v", err)
		}
	}

	if err = d.Set("ip_start", dataSourceFlattenLoadBalanceIppoolIpStart(o["ip-start"], d, "ip_start")); err != nil {
		if !fortiAPIPatch(o["ip-start"]) {
			return fmt.Errorf("Error reading ip-start: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenLoadBalanceIppoolInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceIppoolMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ip6_start", dataSourceFlattenLoadBalanceIppoolIp6Start(o["ip6-start"], d, "ip6_start")); err != nil {
		if !fortiAPIPatch(o["ip6-start"]) {
			return fmt.Errorf("Error reading ip6-start: %v", err)
		}
	}

	return nil
}
