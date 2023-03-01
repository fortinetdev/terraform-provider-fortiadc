// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system health check script.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemHealthCheckScript() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemHealthCheckScriptRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceSystemHealthCheckScriptRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemHealthCheckScript: type error")
	}

	o, err := c.ReadSystemHealthCheckScript(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemHealthCheckScript: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemHealthCheckScript(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemHealthCheckScript from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemHealthCheckScriptMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckScriptFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemHealthCheckScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenSystemHealthCheckScriptMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("file", dataSourceFlattenSystemHealthCheckScriptFile(o["file"], d, "file")); err != nil {
		if !fortiAPIPatch(o["file"]) {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	return nil
}
