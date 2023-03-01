// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system dns.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemDnsRead,
		Schema: map[string]*schema.Schema{
			"ip_second": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "SystemDns"

	o, err := c.ReadSystemDns(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemDns: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemDns(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemDns from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemDnsIpSecond(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsIpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip_second", dataSourceFlattenSystemDnsIpSecond(o["ip_second"], d, "ip_second")); err != nil {
		if !fortiAPIPatch(o["ip_second"]) {
			return fmt.Errorf("Error reading ip_second: %v", err)
		}
	}

	if err = d.Set("ip_primary", dataSourceFlattenSystemDnsIpPrimary(o["ip_primary"], d, "ip_primary")); err != nil {
		if !fortiAPIPatch(o["ip_primary"]) {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	return nil
}
