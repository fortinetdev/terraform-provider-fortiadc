// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system dns vdom.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemDnsVdom() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemDnsVdomRead,
		Schema: map[string]*schema.Schema{
			"ip_second": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_overide": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemDnsVdomRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "SystemDns"

	o, err := c.ReadSystemDnsVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemDns: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemDnsVdom(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemDns from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemDnsVdomIpSecond(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsVdomIpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsVdomDnsOveride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemDnsVdom(d *schema.ResourceData, o map[string]interface{}) error {
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

	if err = d.Set("dns_overide", dataSourceFlattenSystemDnsVdomDnsOveride(o["dns_overide"], d, "dns_overide")); err != nil {
		if !fortiAPIPatch(o["dns_overide"]) {
			return fmt.Errorf("Error reading dns_overide: %v", err)
		}
	}

	return nil
}
