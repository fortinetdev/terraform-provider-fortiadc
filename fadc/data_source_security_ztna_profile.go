// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security ztna profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSecurityZtnaProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityZtnaProfileRead,
		Schema: map[string]*schema.Schema{
			"log_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
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
func dataSourceSecurityZtnaProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityZtnaProfile: type error")
	}

	o, err := c.ReadSecurityZtnaProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SecurityZtnaProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSecurityZtnaProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SecurityZtnaProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSecurityZtnaProfileLogStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityZtnaProfileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityZtnaProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSecurityZtnaProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("log_status", dataSourceFlattenSecurityZtnaProfileLogStatus(o["log_status"], d, "log_status")); err != nil {
		if !fortiAPIPatch(o["log_status"]) {
			return fmt.Errorf("Error reading log_status: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSecurityZtnaProfileComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSecurityZtnaProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
