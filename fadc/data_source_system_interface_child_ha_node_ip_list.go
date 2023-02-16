// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface child ha node ip list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemInterfaceChildHaNodeIpList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceChildHaNodeIpListRead,
		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fadc_id": &schema.Schema{
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
func dataSourceSystemInterfaceChildHaNodeIpListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("id")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	o, err := c.ReadSystemInterfaceChildHaNodeIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemInterfaceChildHaNodeIpList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemInterfaceChildHaNodeIpListNode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildHaNodeIpListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceChildHaNodeIpListAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildHaNodeIpListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemInterfaceChildHaNodeIpList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("node", dataSourceFlattenSystemInterfaceChildHaNodeIpListNode(o["node"], d, "node")); err != nil {
		if !fortiAPIPatch(o["node"]) {
			return fmt.Errorf("Error reading node: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemInterfaceChildHaNodeIpListIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", dataSourceFlattenSystemInterfaceChildHaNodeIpListAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("fadc_id", dataSourceFlattenSystemInterfaceChildHaNodeIpListId(o["id"], d, "fadc_id")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}

	return nil
}
