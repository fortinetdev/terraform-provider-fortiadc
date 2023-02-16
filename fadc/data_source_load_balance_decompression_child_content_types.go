// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance decompression child content types.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceDecompressionChildContentTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceDecompressionChildContentTypesRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_content_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceLoadBalanceDecompressionChildContentTypesRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("custom_content_type")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceDecompressionChildContentTypes: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceDecompressionChildContentTypes: type error")
	}

	o, err := c.ReadLoadBalanceDecompressionChildContentTypes(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceDecompressionChildContentTypes: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceDecompressionChildContentTypes(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceDecompressionChildContentTypes from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceDecompressionChildContentTypesMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceDecompressionChildContentTypesContentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceDecompressionChildContentTypesCustomContentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceDecompressionChildContentTypes(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceDecompressionChildContentTypesMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("content_type", dataSourceFlattenLoadBalanceDecompressionChildContentTypesContentType(o["content_type"], d, "content_type")); err != nil {
		if !fortiAPIPatch(o["content_type"]) {
			return fmt.Errorf("Error reading content_type: %v", err)
		}
	}

	if err = d.Set("custom_content_type", dataSourceFlattenLoadBalanceDecompressionChildContentTypesCustomContentType(o["custom_content_type"], d, "custom_content_type")); err != nil {
		if !fortiAPIPatch(o["custom_content_type"]) {
			return fmt.Errorf("Error reading custom_content_type: %v", err)
		}
	}

	return nil
}
