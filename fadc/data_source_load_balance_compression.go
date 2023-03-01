// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance compression.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceCompression() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceCompressionRead,
		Schema: map[string]*schema.Schema{
			"max_cpu_usage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri_list_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gzip_window_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gzip_memory_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_content_length": &schema.Schema{
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
func dataSourceLoadBalanceCompressionRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCompression: type error")
	}

	o, err := c.ReadLoadBalanceCompression(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCompression: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceCompression(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCompression from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceCompressionMaxCpuUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionUriListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionGzipWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionGzipMemoryLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionCpuLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionMinContentLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceCompression(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_cpu_usage", dataSourceFlattenLoadBalanceCompressionMaxCpuUsage(o["max_cpu_usage"], d, "max_cpu_usage")); err != nil {
		if !fortiAPIPatch(o["max_cpu_usage"]) {
			return fmt.Errorf("Error reading max_cpu_usage: %v", err)
		}
	}

	if err = d.Set("level", dataSourceFlattenLoadBalanceCompressionLevel(o["level"], d, "level")); err != nil {
		if !fortiAPIPatch(o["level"]) {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("uri_list_type", dataSourceFlattenLoadBalanceCompressionUriListType(o["uri_list_type"], d, "uri_list_type")); err != nil {
		if !fortiAPIPatch(o["uri_list_type"]) {
			return fmt.Errorf("Error reading uri_list_type: %v", err)
		}
	}

	if err = d.Set("gzip_window_size", dataSourceFlattenLoadBalanceCompressionGzipWindowSize(o["gzip_window_size"], d, "gzip_window_size")); err != nil {
		if !fortiAPIPatch(o["gzip_window_size"]) {
			return fmt.Errorf("Error reading gzip_window_size: %v", err)
		}
	}

	if err = d.Set("gzip_memory_level", dataSourceFlattenLoadBalanceCompressionGzipMemoryLevel(o["gzip_memory_level"], d, "gzip_memory_level")); err != nil {
		if !fortiAPIPatch(o["gzip_memory_level"]) {
			return fmt.Errorf("Error reading gzip_memory_level: %v", err)
		}
	}

	if err = d.Set("cpu_limit", dataSourceFlattenLoadBalanceCompressionCpuLimit(o["cpu_limit"], d, "cpu_limit")); err != nil {
		if !fortiAPIPatch(o["cpu_limit"]) {
			return fmt.Errorf("Error reading cpu_limit: %v", err)
		}
	}

	if err = d.Set("min_content_length", dataSourceFlattenLoadBalanceCompressionMinContentLength(o["min_content_length"], d, "min_content_length")); err != nil {
		if !fortiAPIPatch(o["min_content_length"]) {
			return fmt.Errorf("Error reading min_content_length: %v", err)
		}
	}

	if err = d.Set("method", dataSourceFlattenLoadBalanceCompressionMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceCompressionMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
