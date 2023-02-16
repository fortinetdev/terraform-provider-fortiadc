// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pagespeed profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePagespeedProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePagespeedProfileRead,
		Update: resourceLoadBalancePagespeedProfileUpdate,
		Create: resourceLoadBalancePagespeedProfileCreate,
		Delete: resourceLoadBalancePagespeedProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"move_css_to_head": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"jpeg_sampling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"resize_image": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"html": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"combine_css": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_combine_css_byte": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"css": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceLoadBalancePagespeedProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalancePagespeedProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePagespeedProfile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalancePagespeedProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePagespeedProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalancePagespeedProfileRead(d, m)
}
func resourceLoadBalancePagespeedProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalancePagespeedProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePagespeedProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePagespeedProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePagespeedProfile resource: %v", err)
	}

	return resourceLoadBalancePagespeedProfileRead(d, m)
}
func resourceLoadBalancePagespeedProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalancePagespeedProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePagespeedProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePagespeedProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalancePagespeedProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePagespeedProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePagespeedProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePagespeedProfile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePagespeedProfileMoveCssToHead(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileJpegSampling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileImage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileResizeImage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileHtml(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileCombineCss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileMaxCombineCssByte(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedProfileCss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePagespeedProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("move_css_to_head", flattenLoadBalancePagespeedProfileMoveCssToHead(o["move_css_to_head"], d, "move_css_to_head", sv)); err != nil {
		if !fortiAPIPatch(o["move_css_to_head"]) {
			return fmt.Errorf("Error reading move_css_to_head: %v", err)
		}
	}

	if err = d.Set("jpeg_sampling", flattenLoadBalancePagespeedProfileJpegSampling(o["jpeg_sampling"], d, "jpeg_sampling", sv)); err != nil {
		if !fortiAPIPatch(o["jpeg_sampling"]) {
			return fmt.Errorf("Error reading jpeg_sampling: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePagespeedProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("image", flattenLoadBalancePagespeedProfileImage(o["image"], d, "image", sv)); err != nil {
		if !fortiAPIPatch(o["image"]) {
			return fmt.Errorf("Error reading image: %v", err)
		}
	}

	if err = d.Set("resize_image", flattenLoadBalancePagespeedProfileResizeImage(o["resize_image"], d, "resize_image", sv)); err != nil {
		if !fortiAPIPatch(o["resize_image"]) {
			return fmt.Errorf("Error reading resize_image: %v", err)
		}
	}

	if err = d.Set("html", flattenLoadBalancePagespeedProfileHtml(o["html"], d, "html", sv)); err != nil {
		if !fortiAPIPatch(o["html"]) {
			return fmt.Errorf("Error reading html: %v", err)
		}
	}

	if err = d.Set("combine_css", flattenLoadBalancePagespeedProfileCombineCss(o["combine_css"], d, "combine_css", sv)); err != nil {
		if !fortiAPIPatch(o["combine_css"]) {
			return fmt.Errorf("Error reading combine_css: %v", err)
		}
	}

	if err = d.Set("max_combine_css_byte", flattenLoadBalancePagespeedProfileMaxCombineCssByte(o["max_combine_css_byte"], d, "max_combine_css_byte", sv)); err != nil {
		if !fortiAPIPatch(o["max_combine_css_byte"]) {
			return fmt.Errorf("Error reading max_combine_css_byte: %v", err)
		}
	}

	if err = d.Set("css", flattenLoadBalancePagespeedProfileCss(o["css"], d, "css", sv)); err != nil {
		if !fortiAPIPatch(o["css"]) {
			return fmt.Errorf("Error reading css: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePagespeedProfileMoveCssToHead(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileJpegSampling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileImage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileResizeImage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileHtml(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileCombineCss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileMaxCombineCssByte(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedProfileCss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePagespeedProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("move_css_to_head"); ok {
		t, err := expandLoadBalancePagespeedProfileMoveCssToHead(d, v, "move_css_to_head", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["move_css_to_head"] = t
		}
	}

	if v, ok := d.GetOk("jpeg_sampling"); ok {
		t, err := expandLoadBalancePagespeedProfileJpegSampling(d, v, "jpeg_sampling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jpeg_sampling"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePagespeedProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("image"); ok {
		t, err := expandLoadBalancePagespeedProfileImage(d, v, "image", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image"] = t
		}
	}

	if v, ok := d.GetOk("resize_image"); ok {
		t, err := expandLoadBalancePagespeedProfileResizeImage(d, v, "resize_image", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resize_image"] = t
		}
	}

	if v, ok := d.GetOk("html"); ok {
		t, err := expandLoadBalancePagespeedProfileHtml(d, v, "html", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["html"] = t
		}
	}

	if v, ok := d.GetOk("combine_css"); ok {
		t, err := expandLoadBalancePagespeedProfileCombineCss(d, v, "combine_css", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["combine_css"] = t
		}
	}

	if v, ok := d.GetOk("max_combine_css_byte"); ok {
		t, err := expandLoadBalancePagespeedProfileMaxCombineCssByte(d, v, "max_combine_css_byte", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_combine_css_byte"] = t
		}
	}

	if v, ok := d.GetOk("css"); ok {
		t, err := expandLoadBalancePagespeedProfileCss(d, v, "css", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["css"] = t
		}
	}

	return &obj, nil
}
