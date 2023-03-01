// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security dos dos protection profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSecurityDosDosProtectionProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityDosDosProtectionProfileRead,
		Schema: map[string]*schema.Schema{
			"tcp_slow_data_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_conn_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_conn_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_req_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_access_limit": &schema.Schema{
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
func dataSourceSecurityDosDosProtectionProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityDosDosProtectionProfile: type error")
	}

	o, err := c.ReadSecurityDosDosProtectionProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SecurityDosDosProtectionProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSecurityDosDosProtectionProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SecurityDosDosProtectionProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSecurityDosDosProtectionProfileTcpSlowDataLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityDosDosProtectionProfileHttpConnLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityDosDosProtectionProfileTcpConnLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityDosDosProtectionProfileHttpReqLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityDosDosProtectionProfileHttpAccessLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityDosDosProtectionProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSecurityDosDosProtectionProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("tcp_slow_data_limit", dataSourceFlattenSecurityDosDosProtectionProfileTcpSlowDataLimit(o["tcp_slow_data_limit"], d, "tcp_slow_data_limit")); err != nil {
		if !fortiAPIPatch(o["tcp_slow_data_limit"]) {
			return fmt.Errorf("Error reading tcp_slow_data_limit: %v", err)
		}
	}

	if err = d.Set("http_conn_limit", dataSourceFlattenSecurityDosDosProtectionProfileHttpConnLimit(o["http_conn_limit"], d, "http_conn_limit")); err != nil {
		if !fortiAPIPatch(o["http_conn_limit"]) {
			return fmt.Errorf("Error reading http_conn_limit: %v", err)
		}
	}

	if err = d.Set("tcp_conn_limit", dataSourceFlattenSecurityDosDosProtectionProfileTcpConnLimit(o["tcp_conn_limit"], d, "tcp_conn_limit")); err != nil {
		if !fortiAPIPatch(o["tcp_conn_limit"]) {
			return fmt.Errorf("Error reading tcp_conn_limit: %v", err)
		}
	}

	if err = d.Set("http_req_limit", dataSourceFlattenSecurityDosDosProtectionProfileHttpReqLimit(o["http_req_limit"], d, "http_req_limit")); err != nil {
		if !fortiAPIPatch(o["http_req_limit"]) {
			return fmt.Errorf("Error reading http_req_limit: %v", err)
		}
	}

	if err = d.Set("http_access_limit", dataSourceFlattenSecurityDosDosProtectionProfileHttpAccessLimit(o["http_access_limit"], d, "http_access_limit")); err != nil {
		if !fortiAPIPatch(o["http_access_limit"]) {
			return fmt.Errorf("Error reading http_access_limit: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSecurityDosDosProtectionProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
