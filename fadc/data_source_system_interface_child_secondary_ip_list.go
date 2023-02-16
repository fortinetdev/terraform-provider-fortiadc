// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface child secondary ip list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemInterfaceChildSecondaryIpList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceChildSecondaryIpListRead,
		Schema: map[string]*schema.Schema{
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"floating_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"floating": &schema.Schema{
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
func dataSourceSystemInterfaceChildSecondaryIpListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	o, err := c.ReadSystemInterfaceChildSecondaryIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemInterfaceChildSecondaryIpList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemInterfaceChildSecondaryIpListTrafficGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildSecondaryIpListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceChildSecondaryIpListFloatingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildSecondaryIpListAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildSecondaryIpListFloating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildSecondaryIpListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemInterfaceChildSecondaryIpList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("traffic_group", dataSourceFlattenSystemInterfaceChildSecondaryIpListTrafficGroup(o["traffic-group"], d, "traffic_group")); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemInterfaceChildSecondaryIpListIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("floating_ip", dataSourceFlattenSystemInterfaceChildSecondaryIpListFloatingIp(o["floating-ip"], d, "floating_ip")); err != nil {
		if !fortiAPIPatch(o["floating-ip"]) {
			return fmt.Errorf("Error reading floating-ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", dataSourceFlattenSystemInterfaceChildSecondaryIpListAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("floating", dataSourceFlattenSystemInterfaceChildSecondaryIpListFloating(o["floating"], d, "floating")); err != nil {
		if !fortiAPIPatch(o["floating"]) {
			return fmt.Errorf("Error reading floating: %v", err)
		}
	}

	if err = d.Set("fadc_id", dataSourceFlattenSystemInterfaceChildSecondaryIpListId(o["id"], d, "fadc_id")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}

	return nil
}
