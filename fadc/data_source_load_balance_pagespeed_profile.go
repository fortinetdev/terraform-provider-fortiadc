// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pagespeed profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePagespeedProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePagespeedProfileRead,
		Schema: map[string]*schema.Schema{
			"move_css_to_head": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"jpeg_sampling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resize_image": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"html": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"combine_css": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_combine_css_byte": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"css": &schema.Schema{
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
func dataSourceLoadBalancePagespeedProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeedProfile: type error")
	}

	o, err := c.ReadLoadBalancePagespeedProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePagespeedProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePagespeedProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePagespeedProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePagespeedProfileMoveCssToHead(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileJpegSampling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileResizeImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileHtml(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileCombineCss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileMaxCombineCssByte(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedProfileCss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePagespeedProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("move_css_to_head", dataSourceFlattenLoadBalancePagespeedProfileMoveCssToHead(o["move_css_to_head"], d, "move_css_to_head")); err != nil {
		if !fortiAPIPatch(o["move_css_to_head"]) {
			return fmt.Errorf("Error reading move_css_to_head: %v", err)
		}
	}

	if err = d.Set("jpeg_sampling", dataSourceFlattenLoadBalancePagespeedProfileJpegSampling(o["jpeg_sampling"], d, "jpeg_sampling")); err != nil {
		if !fortiAPIPatch(o["jpeg_sampling"]) {
			return fmt.Errorf("Error reading jpeg_sampling: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePagespeedProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("image", dataSourceFlattenLoadBalancePagespeedProfileImage(o["image"], d, "image")); err != nil {
		if !fortiAPIPatch(o["image"]) {
			return fmt.Errorf("Error reading image: %v", err)
		}
	}

	if err = d.Set("resize_image", dataSourceFlattenLoadBalancePagespeedProfileResizeImage(o["resize_image"], d, "resize_image")); err != nil {
		if !fortiAPIPatch(o["resize_image"]) {
			return fmt.Errorf("Error reading resize_image: %v", err)
		}
	}

	if err = d.Set("html", dataSourceFlattenLoadBalancePagespeedProfileHtml(o["html"], d, "html")); err != nil {
		if !fortiAPIPatch(o["html"]) {
			return fmt.Errorf("Error reading html: %v", err)
		}
	}

	if err = d.Set("combine_css", dataSourceFlattenLoadBalancePagespeedProfileCombineCss(o["combine_css"], d, "combine_css")); err != nil {
		if !fortiAPIPatch(o["combine_css"]) {
			return fmt.Errorf("Error reading combine_css: %v", err)
		}
	}

	if err = d.Set("max_combine_css_byte", dataSourceFlattenLoadBalancePagespeedProfileMaxCombineCssByte(o["max_combine_css_byte"], d, "max_combine_css_byte")); err != nil {
		if !fortiAPIPatch(o["max_combine_css_byte"]) {
			return fmt.Errorf("Error reading max_combine_css_byte: %v", err)
		}
	}

	if err = d.Set("css", dataSourceFlattenLoadBalancePagespeedProfileCss(o["css"], d, "css")); err != nil {
		if !fortiAPIPatch(o["css"]) {
			return fmt.Errorf("Error reading css: %v", err)
		}
	}

	return nil
}
