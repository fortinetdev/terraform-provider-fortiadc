// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system traffic group.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemTrafficGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemTrafficGroupRead,
		Schema: map[string]*schema.Schema{
			"network_failover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"preempt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"failover_order": &schema.Schema{
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
func dataSourceSystemTrafficGroupRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemTrafficGroup: type error")
	}

	o, err := c.ReadSystemTrafficGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemTrafficGroup: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemTrafficGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemTrafficGroup from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemTrafficGroupNetworkFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTrafficGroupPreempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTrafficGroupMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTrafficGroupFailoverOrder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemTrafficGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("network_failover", dataSourceFlattenSystemTrafficGroupNetworkFailover(o["network-failover"], d, "network_failover")); err != nil {
		if !fortiAPIPatch(o["network-failover"]) {
			return fmt.Errorf("Error reading network-failover: %v", err)
		}
	}

	if err = d.Set("preempt", dataSourceFlattenSystemTrafficGroupPreempt(o["preempt"], d, "preempt")); err != nil {
		if !fortiAPIPatch(o["preempt"]) {
			return fmt.Errorf("Error reading preempt: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemTrafficGroupMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("failover_order", dataSourceFlattenSystemTrafficGroupFailoverOrder(o["failover-order"], d, "failover_order")); err != nil {
		if !fortiAPIPatch(o["failover-order"]) {
			return fmt.Errorf("Error reading failover-order: %v", err)
		}
	}

	return nil
}
