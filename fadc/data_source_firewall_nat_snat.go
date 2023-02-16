// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  firewall nat snat.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceFirewallNatSnat() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallNatSnatRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"out_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trans_to_ip_end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trans_to_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trans_to_ip_start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trans_to_ip": &schema.Schema{
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
func dataSourceFirewallNatSnatRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallNatSnat: type error")
	}

	o, err := c.ReadFirewallNatSnat(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing FirewallNatSnat: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectFirewallNatSnat(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallNatSnat from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenFirewallNatSnatStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatOutInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenFirewallNatSnatTrafficGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatTransToIpEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenFirewallNatSnatTransToType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatTransToIpStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatTransToIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallNatSnatMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallNatSnat(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenFirewallNatSnatStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("out_interface", dataSourceFlattenFirewallNatSnatOutInterface(o["out-interface"], d, "out_interface")); err != nil {
		if !fortiAPIPatch(o["out-interface"]) {
			return fmt.Errorf("Error reading out-interface: %v", err)
		}
	}

	if err = d.Set("from", dataSourceFlattenFirewallNatSnatFrom(o["from"], d, "from")); err != nil {
		if !fortiAPIPatch(o["from"]) {
			return fmt.Errorf("Error reading from: %v", err)
		}
	}

	if err = d.Set("traffic_group", dataSourceFlattenFirewallNatSnatTrafficGroup(o["traffic-group"], d, "traffic_group")); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("trans_to_ip_end", dataSourceFlattenFirewallNatSnatTransToIpEnd(o["trans-to-ip-end"], d, "trans_to_ip_end")); err != nil {
		if !fortiAPIPatch(o["trans-to-ip-end"]) {
			return fmt.Errorf("Error reading trans-to-ip-end: %v", err)
		}
	}

	if err = d.Set("to", dataSourceFlattenFirewallNatSnatTo(o["to"], d, "to")); err != nil {
		if !fortiAPIPatch(o["to"]) {
			return fmt.Errorf("Error reading to: %v", err)
		}
	}

	if err = d.Set("trans_to_type", dataSourceFlattenFirewallNatSnatTransToType(o["trans-to-type"], d, "trans_to_type")); err != nil {
		if !fortiAPIPatch(o["trans-to-type"]) {
			return fmt.Errorf("Error reading trans-to-type: %v", err)
		}
	}

	if err = d.Set("trans_to_ip_start", dataSourceFlattenFirewallNatSnatTransToIpStart(o["trans-to-ip-start"], d, "trans_to_ip_start")); err != nil {
		if !fortiAPIPatch(o["trans-to-ip-start"]) {
			return fmt.Errorf("Error reading trans-to-ip-start: %v", err)
		}
	}

	if err = d.Set("trans_to_ip", dataSourceFlattenFirewallNatSnatTransToIp(o["trans-to-ip"], d, "trans_to_ip")); err != nil {
		if !fortiAPIPatch(o["trans-to-ip"]) {
			return fmt.Errorf("Error reading trans-to-ip: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenFirewallNatSnatMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
