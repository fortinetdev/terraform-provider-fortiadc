// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance clone pool child pool member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceClonePoolChildPoolMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceClonePoolChildPoolMemberRead,
		Schema: map[string]*schema.Schema{
			"dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceLoadBalanceClonePoolChildPoolMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	o, err := c.ReadLoadBalanceClonePoolChildPoolMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceClonePoolChildPoolMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberDstMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberSrcMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClonePoolChildPoolMemberMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceClonePoolChildPoolMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst_mac", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberDstMac(o["dst_mac"], d, "dst_mac")); err != nil {
		if !fortiAPIPatch(o["dst_mac"]) {
			return fmt.Errorf("Error reading dst_mac: %v", err)
		}
	}

	if err = d.Set("src_mac", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberSrcMac(o["src_mac"], d, "src_mac")); err != nil {
		if !fortiAPIPatch(o["src_mac"]) {
			return fmt.Errorf("Error reading src_mac: %v", err)
		}
	}

	if err = d.Set("address", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenLoadBalanceClonePoolChildPoolMemberMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}
