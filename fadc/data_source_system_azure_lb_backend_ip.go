// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system azure lb backend ip.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemAzureLbBackendIp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAzureLbBackendIpRead,
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
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
func dataSourceSystemAzureLbBackendIpRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAzureLbBackendIp: type error")
	}

	o, err := c.ReadSystemAzureLbBackendIp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemAzureLbBackendIp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemAzureLbBackendIp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAzureLbBackendIp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemAzureLbBackendIpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAzureLbBackendIpMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAzureLbBackendIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", dataSourceFlattenSystemAzureLbBackendIpIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemAzureLbBackendIpMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
