// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system time manual.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemTimeManual() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemTimeManualRead,
		Schema: map[string]*schema.Schema{
			"tz": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemTimeManualRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "SystemTimeManual"

	o, err := c.ReadSystemTimeManual(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemTimeManual: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemTimeManual(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemTimeManual from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemTimeManualTz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTimeManualDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemTimeManual(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("tz", dataSourceFlattenSystemTimeManualTz(o["tz"], d, "tz")); err != nil {
		if !fortiAPIPatch(o["tz"]) {
			return fmt.Errorf("Error reading tz: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenSystemTimeManualDst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	return nil
}
