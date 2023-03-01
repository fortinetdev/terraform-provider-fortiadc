// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance decompression.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceDecompression() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceDecompressionRead,
		Schema: map[string]*schema.Schema{
			"max_cpu_usage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri_list_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"method": &schema.Schema{
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
func dataSourceLoadBalanceDecompressionRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceDecompression: type error")
	}

	o, err := c.ReadLoadBalanceDecompression(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceDecompression: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceDecompression(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceDecompression from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceDecompressionMaxCpuUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceDecompressionUriListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceDecompressionCpuLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceDecompressionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceDecompressionMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceDecompression(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_cpu_usage", dataSourceFlattenLoadBalanceDecompressionMaxCpuUsage(o["max_cpu_usage"], d, "max_cpu_usage")); err != nil {
		if !fortiAPIPatch(o["max_cpu_usage"]) {
			return fmt.Errorf("Error reading max_cpu_usage: %v", err)
		}
	}

	if err = d.Set("uri_list_type", dataSourceFlattenLoadBalanceDecompressionUriListType(o["uri_list_type"], d, "uri_list_type")); err != nil {
		if !fortiAPIPatch(o["uri_list_type"]) {
			return fmt.Errorf("Error reading uri_list_type: %v", err)
		}
	}

	if err = d.Set("cpu_limit", dataSourceFlattenLoadBalanceDecompressionCpuLimit(o["cpu_limit"], d, "cpu_limit")); err != nil {
		if !fortiAPIPatch(o["cpu_limit"]) {
			return fmt.Errorf("Error reading cpu_limit: %v", err)
		}
	}

	if err = d.Set("method", dataSourceFlattenLoadBalanceDecompressionMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceDecompressionMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
