// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUploadCaptchaCust() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUploadCaptchaCustRead,
		Update: resourceUploadCaptchaCustUpdate,
		Create: resourceUploadCaptchaCustCreate,
		Delete: resourceUploadCaptchaCustDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"errorpagefile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_retry_times": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_picture_changes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pic_difficulty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_block_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_verify_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_valid_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceUploadCaptchaCustCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UploadCaptchaCust: type error")
	}

	obj, err := getObjectUploadCaptchaCust(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UploadCaptchaCust resource while getting object: %v", err)
	}

	_, err = c.CreateUploadCaptchaCust(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating UploadCaptchaCust resource: %v", err)
	}

	d.SetId(mkey)

	//return resourceUploadCaptchaCustRead(d, m)
	return nil
}
func resourceUploadCaptchaCustUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUploadCaptchaCust(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClientSslProfile resource while getting object: %v", err)
	}

	_, err = c.CreateUploadCaptchaCust(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClientSslProfile resource: %v", err)
	}

	return resourceUploadCaptchaCustRead(d, m)
}
func resourceUploadCaptchaCustDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUploadCaptchaCust(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UploadCaptchaCust resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceUploadCaptchaCustRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUploadCaptchaCust(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UploadCaptchaCust resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUploadCaptchaCust(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UploadCaptchaCust resource from API: %v", err)
	}
	return nil
}

func flattenUploadCaptchaCustMaxBlockPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustCustomPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustMaxValidPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustMaxPictureChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustMaxVerifyPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustMaxRetryTimes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustVpath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustCustomPageFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustPicDifficulty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadCaptchaCustMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUploadCaptchaCust(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_block_period", flattenUploadCaptchaCustMaxBlockPeriod(o["max_block_period"], d, "max_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max_block_period"]) {
			return fmt.Errorf("Error reading max_block_period: %v", err)
		}
	}

	if err = d.Set("custom_page", flattenUploadCaptchaCustCustomPage(o["custom_page"], d, "custom_page", sv)); err != nil {
		if !fortiAPIPatch(o["custom_page"]) {
			return fmt.Errorf("Error reading custom_page: %v", err)
		}
	}

	if err = d.Set("max_valid_period", flattenUploadCaptchaCustMaxValidPeriod(o["max_valid_period"], d, "max_valid_period", sv)); err != nil {
		if !fortiAPIPatch(o["max_valid_period"]) {
			return fmt.Errorf("Error reading max_valid_period: %v", err)
		}
	}

	if err = d.Set("max_picture_changes", flattenUploadCaptchaCustMaxPictureChanges(o["max_picture_changes"], d, "max_picture_changes", sv)); err != nil {
		if !fortiAPIPatch(o["max_picture_changes"]) {
			return fmt.Errorf("Error reading max_picture_changes: %v", err)
		}
	}

	if err = d.Set("max_verify_period", flattenUploadCaptchaCustMaxVerifyPeriod(o["max_verify_period"], d, "max_verify_period", sv)); err != nil {
		if !fortiAPIPatch(o["max_verify_period"]) {
			return fmt.Errorf("Error reading max_verify_period: %v", err)
		}
	}

	if err = d.Set("max_retry_times", flattenUploadCaptchaCustMaxRetryTimes(o["max_retry_times"], d, "max_retry_times", sv)); err != nil {
		if !fortiAPIPatch(o["max_retry_times"]) {
			return fmt.Errorf("Error reading max_retry_times: %v", err)
		}
	}

	if err = d.Set("vpath", flattenUploadCaptchaCustVpath(o["vpath"], d, "vpath", sv)); err != nil {
		if !fortiAPIPatch(o["vpath"]) {
			return fmt.Errorf("Error reading vpath: %v", err)
		}
	}

	if err = d.Set("custom_page_file", flattenUploadCaptchaCustCustomPageFile(o["custom_page_file"], d, "custom_page_file", sv)); err != nil {
		if !fortiAPIPatch(o["custom_page_file"]) {
			return fmt.Errorf("Error reading custom_page_file: %v", err)
		}
	}

	if err = d.Set("pic_difficulty", flattenUploadCaptchaCustPicDifficulty(o["pic_difficulty"], d, "pic_difficulty", sv)); err != nil {
		if !fortiAPIPatch(o["pic_difficulty"]) {
			return fmt.Errorf("Error reading pic_difficulty: %v", err)
		}
	}

	if err = d.Set("mkey", flattenUploadCaptchaCustMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUploadCaptchaCustMaxBlockPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustCustomPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustMaxValidPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustMaxPictureChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustMaxVerifyPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustMaxRetryTimes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustCustomPageFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustPicDifficulty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustVpath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustUpdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustErrorPageFile(d *schema.ResourceData, v interface{}, pre string, sv string, writer *multipart.Writer) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadCaptchaCustVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUploadCaptchaCust(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	body := &bytes.Buffer{}
	obj := make(map[string]interface{})
	writer := multipart.NewWriter(body)

	if v, ok := d.GetOk("max_picture_changes"); ok {
		t, err := expandUploadCaptchaCustMaxPictureChanges(d, v, "max_picture_changes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "max_picture_changes", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("pic_difficulty"); ok {
		t, err := expandUploadCaptchaCustPicDifficulty(d, v, "pic_difficulty", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "pic_difficulty", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("max_block_period"); ok {
		t, err := expandUploadCaptchaCustMaxBlockPeriod(d, v, "max_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "max_block_period", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("max_verify_period"); ok {
		t, err := expandUploadCaptchaCustMaxVerifyPeriod(d, v, "max_verify_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "max_verify_period", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("max_valid_period"); ok {
		t, err := expandUploadCaptchaCustMaxValidPeriod(d, v, "max_valid_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "max_valid_period", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("custom_page"); ok {
		t, err := expandUploadCaptchaCustCustomPage(d, v, "custom_page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "custom_page", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("max_retry_times"); ok {
		t, err := expandUploadCaptchaCustMaxRetryTimes(d, v, "max_retry_times", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "max_retry_times", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("vpath"); ok {
		t, err := expandUploadCaptchaCustVpath(d, v, "vpath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "vpath", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	if v, ok := d.GetOk("update"); ok {
		t, err := expandUploadCaptchaCustUpdate(d, v, "update", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "update", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("errorpagefile"); ok {
		t, err := expandUploadCaptchaCustErrorPageFile(d, v, "errorpagefile", sv, writer)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "errorPageFile", t.(string), true)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUploadCaptchaCustMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "mkey", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandUploadCaptchaCustVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "vdom", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	writer.Close()
	obj["content_type"] = writer.FormDataContentType()
	obj["data"] = body.Bytes()

	return &obj, nil
}
