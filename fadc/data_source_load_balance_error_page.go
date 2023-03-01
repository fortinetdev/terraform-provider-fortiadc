// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance error page.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceErrorPage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceErrorPageRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vpath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceLoadBalanceErrorPageRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceErrorPage: type error")
	}

	o, err := c.ReadLoadBalanceErrorPage(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceErrorPage: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceErrorPage(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceErrorPage from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceErrorPageMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceErrorPageVpath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceErrorPageFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceErrorPage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceErrorPageMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("vpath", dataSourceFlattenLoadBalanceErrorPageVpath(o["vpath"], d, "vpath")); err != nil {
		if !fortiAPIPatch(o["vpath"]) {
			return fmt.Errorf("Error reading vpath: %v", err)
		}
	}

	if err = d.Set("file", dataSourceFlattenLoadBalanceErrorPageFile(o["file"], d, "file")); err != nil {
		if !fortiAPIPatch(o["file"]) {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	return nil
}
