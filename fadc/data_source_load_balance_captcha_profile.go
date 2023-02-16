// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance captcha profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceCaptchaProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceCaptchaProfileRead,
		Schema: map[string]*schema.Schema{
			"max_block_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_page": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_valid_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_picture_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_verify_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_retry_times": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_page_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pic_difficulty": &schema.Schema{
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
func dataSourceLoadBalanceCaptchaProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCaptchaProfile: type error")
	}

	o, err := c.ReadLoadBalanceCaptchaProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCaptchaProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceCaptchaProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCaptchaProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceCaptchaProfileMaxBlockPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileCustomPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileMaxValidPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileMaxPictureChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileMaxVerifyPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileMaxRetryTimes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileVpath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileCustomPageFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfilePicDifficulty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCaptchaProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceCaptchaProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_block_period", dataSourceFlattenLoadBalanceCaptchaProfileMaxBlockPeriod(o["max_block_period"], d, "max_block_period")); err != nil {
		if !fortiAPIPatch(o["max_block_period"]) {
			return fmt.Errorf("Error reading max_block_period: %v", err)
		}
	}

	if err = d.Set("custom_page", dataSourceFlattenLoadBalanceCaptchaProfileCustomPage(o["custom_page"], d, "custom_page")); err != nil {
		if !fortiAPIPatch(o["custom_page"]) {
			return fmt.Errorf("Error reading custom_page: %v", err)
		}
	}

	if err = d.Set("max_valid_period", dataSourceFlattenLoadBalanceCaptchaProfileMaxValidPeriod(o["max_valid_period"], d, "max_valid_period")); err != nil {
		if !fortiAPIPatch(o["max_valid_period"]) {
			return fmt.Errorf("Error reading max_valid_period: %v", err)
		}
	}

	if err = d.Set("max_picture_changes", dataSourceFlattenLoadBalanceCaptchaProfileMaxPictureChanges(o["max_picture_changes"], d, "max_picture_changes")); err != nil {
		if !fortiAPIPatch(o["max_picture_changes"]) {
			return fmt.Errorf("Error reading max_picture_changes: %v", err)
		}
	}

	if err = d.Set("max_verify_period", dataSourceFlattenLoadBalanceCaptchaProfileMaxVerifyPeriod(o["max_verify_period"], d, "max_verify_period")); err != nil {
		if !fortiAPIPatch(o["max_verify_period"]) {
			return fmt.Errorf("Error reading max_verify_period: %v", err)
		}
	}

	if err = d.Set("max_retry_times", dataSourceFlattenLoadBalanceCaptchaProfileMaxRetryTimes(o["max_retry_times"], d, "max_retry_times")); err != nil {
		if !fortiAPIPatch(o["max_retry_times"]) {
			return fmt.Errorf("Error reading max_retry_times: %v", err)
		}
	}

	if err = d.Set("vpath", dataSourceFlattenLoadBalanceCaptchaProfileVpath(o["vpath"], d, "vpath")); err != nil {
		if !fortiAPIPatch(o["vpath"]) {
			return fmt.Errorf("Error reading vpath: %v", err)
		}
	}

	if err = d.Set("custom_page_file", dataSourceFlattenLoadBalanceCaptchaProfileCustomPageFile(o["custom_page_file"], d, "custom_page_file")); err != nil {
		if !fortiAPIPatch(o["custom_page_file"]) {
			return fmt.Errorf("Error reading custom_page_file: %v", err)
		}
	}

	if err = d.Set("pic_difficulty", dataSourceFlattenLoadBalanceCaptchaProfilePicDifficulty(o["pic_difficulty"], d, "pic_difficulty")); err != nil {
		if !fortiAPIPatch(o["pic_difficulty"]) {
			return fmt.Errorf("Error reading pic_difficulty: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceCaptchaProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
