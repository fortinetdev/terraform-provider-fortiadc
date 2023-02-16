// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface child trust ip list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemInterfaceChildTrustIpList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceChildTrustIpListRead,
		Schema: map[string]*schema.Schema{
			"ip6_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
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
func dataSourceSystemInterfaceChildTrustIpListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	o, err := c.ReadSystemInterfaceChildTrustIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemInterfaceChildTrustIpList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemInterfaceChildTrustIpListIp6Max(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListIp6Min(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListIpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListIpNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListIp6Netmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListIpMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceChildTrustIpListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemInterfaceChildTrustIpList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip6_max", dataSourceFlattenSystemInterfaceChildTrustIpListIp6Max(o["ip6-max"], d, "ip6_max")); err != nil {
		if !fortiAPIPatch(o["ip6-max"]) {
			return fmt.Errorf("Error reading ip6-max: %v", err)
		}
	}

	if err = d.Set("ip6_min", dataSourceFlattenSystemInterfaceChildTrustIpListIp6Min(o["ip6-min"], d, "ip6_min")); err != nil {
		if !fortiAPIPatch(o["ip6-min"]) {
			return fmt.Errorf("Error reading ip6-min: %v", err)
		}
	}

	if err = d.Set("ip_max", dataSourceFlattenSystemInterfaceChildTrustIpListIpMax(o["ip-max"], d, "ip_max")); err != nil {
		if !fortiAPIPatch(o["ip-max"]) {
			return fmt.Errorf("Error reading ip-max: %v", err)
		}
	}

	if err = d.Set("ip_netmask", dataSourceFlattenSystemInterfaceChildTrustIpListIpNetmask(o["ip-netmask"], d, "ip_netmask")); err != nil {
		if !fortiAPIPatch(o["ip-netmask"]) {
			return fmt.Errorf("Error reading ip-netmask: %v", err)
		}
	}

	if err = d.Set("ip6_netmask", dataSourceFlattenSystemInterfaceChildTrustIpListIp6Netmask(o["ip6-netmask"], d, "ip6_netmask")); err != nil {
		if !fortiAPIPatch(o["ip6-netmask"]) {
			return fmt.Errorf("Error reading ip6-netmask: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemInterfaceChildTrustIpListType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ip_min", dataSourceFlattenSystemInterfaceChildTrustIpListIpMin(o["ip-min"], d, "ip_min")); err != nil {
		if !fortiAPIPatch(o["ip-min"]) {
			return fmt.Errorf("Error reading ip-min: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemInterfaceChildTrustIpListName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}
