// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance captcha profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceCaptchaProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceCaptchaProfileRead,
		Update: resourceLoadBalanceCaptchaProfileUpdate,
		Create: resourceLoadBalanceCaptchaProfileCreate,
		Delete: resourceLoadBalanceCaptchaProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_block_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"custom_page": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_valid_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_picture_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_verify_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_retry_times": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vpath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"custom_page_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pic_difficulty": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalanceCaptchaProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceCaptchaProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCaptchaProfile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceCaptchaProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCaptchaProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceCaptchaProfileRead(d, m)
}
func resourceLoadBalanceCaptchaProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceCaptchaProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCaptchaProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceCaptchaProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCaptchaProfile resource: %v", err)
	}

	return resourceLoadBalanceCaptchaProfileRead(d, m)
}
func resourceLoadBalanceCaptchaProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceCaptchaProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceCaptchaProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceCaptchaProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceCaptchaProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCaptchaProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceCaptchaProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCaptchaProfile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceCaptchaProfileMaxBlockPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileCustomPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileMaxValidPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileMaxPictureChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileMaxVerifyPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileMaxRetryTimes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileVpath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileCustomPageFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfilePicDifficulty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCaptchaProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceCaptchaProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_block_period", flattenLoadBalanceCaptchaProfileMaxBlockPeriod(o["max_block_period"], d, "max_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max_block_period"]) {
			return fmt.Errorf("Error reading max_block_period: %v", err)
		}
	}

	if err = d.Set("custom_page", flattenLoadBalanceCaptchaProfileCustomPage(o["custom_page"], d, "custom_page", sv)); err != nil {
		if !fortiAPIPatch(o["custom_page"]) {
			return fmt.Errorf("Error reading custom_page: %v", err)
		}
	}

	if err = d.Set("max_valid_period", flattenLoadBalanceCaptchaProfileMaxValidPeriod(o["max_valid_period"], d, "max_valid_period", sv)); err != nil {
		if !fortiAPIPatch(o["max_valid_period"]) {
			return fmt.Errorf("Error reading max_valid_period: %v", err)
		}
	}

	if err = d.Set("max_picture_changes", flattenLoadBalanceCaptchaProfileMaxPictureChanges(o["max_picture_changes"], d, "max_picture_changes", sv)); err != nil {
		if !fortiAPIPatch(o["max_picture_changes"]) {
			return fmt.Errorf("Error reading max_picture_changes: %v", err)
		}
	}

	if err = d.Set("max_verify_period", flattenLoadBalanceCaptchaProfileMaxVerifyPeriod(o["max_verify_period"], d, "max_verify_period", sv)); err != nil {
		if !fortiAPIPatch(o["max_verify_period"]) {
			return fmt.Errorf("Error reading max_verify_period: %v", err)
		}
	}

	if err = d.Set("max_retry_times", flattenLoadBalanceCaptchaProfileMaxRetryTimes(o["max_retry_times"], d, "max_retry_times", sv)); err != nil {
		if !fortiAPIPatch(o["max_retry_times"]) {
			return fmt.Errorf("Error reading max_retry_times: %v", err)
		}
	}

	if err = d.Set("vpath", flattenLoadBalanceCaptchaProfileVpath(o["vpath"], d, "vpath", sv)); err != nil {
		if !fortiAPIPatch(o["vpath"]) {
			return fmt.Errorf("Error reading vpath: %v", err)
		}
	}

	if err = d.Set("custom_page_file", flattenLoadBalanceCaptchaProfileCustomPageFile(o["custom_page_file"], d, "custom_page_file", sv)); err != nil {
		if !fortiAPIPatch(o["custom_page_file"]) {
			return fmt.Errorf("Error reading custom_page_file: %v", err)
		}
	}

	if err = d.Set("pic_difficulty", flattenLoadBalanceCaptchaProfilePicDifficulty(o["pic_difficulty"], d, "pic_difficulty", sv)); err != nil {
		if !fortiAPIPatch(o["pic_difficulty"]) {
			return fmt.Errorf("Error reading pic_difficulty: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceCaptchaProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceCaptchaProfileMaxBlockPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileCustomPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileMaxValidPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileMaxPictureChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileMaxVerifyPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileMaxRetryTimes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileVpath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileCustomPageFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfilePicDifficulty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCaptchaProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceCaptchaProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_block_period"); ok {
		t, err := expandLoadBalanceCaptchaProfileMaxBlockPeriod(d, v, "max_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_block_period"] = t
		}
	}

	if v, ok := d.GetOk("custom_page"); ok {
		t, err := expandLoadBalanceCaptchaProfileCustomPage(d, v, "custom_page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom_page"] = t
		}
	}

	if v, ok := d.GetOk("max_valid_period"); ok {
		t, err := expandLoadBalanceCaptchaProfileMaxValidPeriod(d, v, "max_valid_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_valid_period"] = t
		}
	}

	if v, ok := d.GetOk("max_picture_changes"); ok {
		t, err := expandLoadBalanceCaptchaProfileMaxPictureChanges(d, v, "max_picture_changes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_picture_changes"] = t
		}
	}

	if v, ok := d.GetOk("max_verify_period"); ok {
		t, err := expandLoadBalanceCaptchaProfileMaxVerifyPeriod(d, v, "max_verify_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_verify_period"] = t
		}
	}

	if v, ok := d.GetOk("max_retry_times"); ok {
		t, err := expandLoadBalanceCaptchaProfileMaxRetryTimes(d, v, "max_retry_times", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_retry_times"] = t
		}
	}

	if v, ok := d.GetOk("vpath"); ok {
		t, err := expandLoadBalanceCaptchaProfileVpath(d, v, "vpath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpath"] = t
		}
	}

	if v, ok := d.GetOk("custom_page_file"); ok {
		t, err := expandLoadBalanceCaptchaProfileCustomPageFile(d, v, "custom_page_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom_page_file"] = t
		}
	}

	if v, ok := d.GetOk("pic_difficulty"); ok {
		t, err := expandLoadBalanceCaptchaProfilePicDifficulty(d, v, "pic_difficulty", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pic_difficulty"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceCaptchaProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
