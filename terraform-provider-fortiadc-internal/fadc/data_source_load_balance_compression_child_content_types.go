// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance compression child content types.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceCompressionChildContentTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceCompressionChildContentTypesRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_content_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceLoadBalanceCompressionChildContentTypesRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCompressionChildContentTypes: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceCompressionChildContentTypes: type error")
	}

	o, err := c.ReadLoadBalanceCompressionChildContentTypes(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCompressionChildContentTypes: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceCompressionChildContentTypes(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCompressionChildContentTypes from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceCompressionChildContentTypesMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionChildContentTypesContentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCompressionChildContentTypesCustomContentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceCompressionChildContentTypes(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceCompressionChildContentTypesMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("content_type", dataSourceFlattenLoadBalanceCompressionChildContentTypesContentType(o["content_type"], d, "content_type")); err != nil {
		if !fortiAPIPatch(o["content_type"]) {
			return fmt.Errorf("Error reading content_type: %v", err)
		}
	}

	if err = d.Set("custom_content_type", dataSourceFlattenLoadBalanceCompressionChildContentTypesCustomContentType(o["custom_content_type"], d, "custom_content_type")); err != nil {
		if !fortiAPIPatch(o["custom_content_type"]) {
			return fmt.Errorf("Error reading custom_content_type: %v", err)
		}
	}

	return nil
}
