// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security ips profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSecurityIpsProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityIpsProfileRead,
		Schema: map[string]*schema.Schema{
			"ips_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
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
func dataSourceSecurityIpsProfileRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("ips_profile_name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityIpsProfile: type error")
	}

	o, err := c.ReadSecurityIpsProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SecurityIpsProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSecurityIpsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SecurityIpsProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSecurityIpsProfileIpsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityIpsProfileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSecurityIpsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ips_profile_name", dataSourceFlattenSecurityIpsProfileIpsProfileName(o["ips_profile_name"], d, "ips_profile_name")); err != nil {
		if !fortiAPIPatch(o["ips_profile_name"]) {
			return fmt.Errorf("Error reading ips_profile_name: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSecurityIpsProfileComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}
