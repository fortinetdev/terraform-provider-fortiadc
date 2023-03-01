// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system isp addr child address.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemIspAddrChildAddress() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemIspAddrChildAddressRead,
		Schema: map[string]*schema.Schema{
			"province": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_netmask": &schema.Schema{
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
func dataSourceSystemIspAddrChildAddressRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	o, err := c.ReadSystemIspAddrChildAddress(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemIspAddrChildAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemIspAddrChildAddress from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemIspAddrChildAddressProvince(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIspAddrChildAddressIpNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemIspAddrChildAddressMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemIspAddrChildAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("province", dataSourceFlattenSystemIspAddrChildAddressProvince(o["province"], d, "province")); err != nil {
		if !fortiAPIPatch(o["province"]) {
			return fmt.Errorf("Error reading province: %v", err)
		}
	}

	if err = d.Set("ip_netmask", dataSourceFlattenSystemIspAddrChildAddressIpNetmask(o["ip-netmask"], d, "ip_netmask")); err != nil {
		if !fortiAPIPatch(o["ip-netmask"]) {
			return fmt.Errorf("Error reading ip-netmask: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemIspAddrChildAddressMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
