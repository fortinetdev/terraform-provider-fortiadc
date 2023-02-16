// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance decompression child content types.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceDecompressionChildContentTypes() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceDecompressionChildContentTypesRead,
		Update: resourceLoadBalanceDecompressionChildContentTypesUpdate,
		Create: resourceLoadBalanceDecompressionChildContentTypesCreate,
		Delete: resourceLoadBalanceDecompressionChildContentTypesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalanceDecompressionChildContentTypesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceDecompressionChildContentTypes(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceDecompressionChildContentTypes resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceDecompressionChildContentTypes(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceDecompressionChildContentTypes resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceDecompressionChildContentTypesRead(d, m)
}
func resourceLoadBalanceDecompressionChildContentTypesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceDecompressionChildContentTypes(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceDecompressionChildContentTypes resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceDecompressionChildContentTypes(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceDecompressionChildContentTypes resource: %v", err)
	}

	return resourceLoadBalanceDecompressionChildContentTypesRead(d, m)
}
func resourceLoadBalanceDecompressionChildContentTypesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceDecompressionChildContentTypes(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceDecompressionChildContentTypes resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceDecompressionChildContentTypesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceDecompressionChildContentTypes(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceDecompressionChildContentTypes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceDecompressionChildContentTypes(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceDecompressionChildContentTypes resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceDecompressionChildContentTypesMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceDecompressionChildContentTypesContentType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceDecompressionChildContentTypesCustomContentType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceDecompressionChildContentTypes(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenLoadBalanceDecompressionChildContentTypesMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("content_type", flattenLoadBalanceDecompressionChildContentTypesContentType(o["content_type"], d, "content_type", sv)); err != nil {
		if !fortiAPIPatch(o["content_type"]) {
			return fmt.Errorf("Error reading content_type: %v", err)
		}
	}

	if err = d.Set("custom_content_type", flattenLoadBalanceDecompressionChildContentTypesCustomContentType(o["custom_content_type"], d, "custom_content_type", sv)); err != nil {
		if !fortiAPIPatch(o["custom_content_type"]) {
			return fmt.Errorf("Error reading custom_content_type: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceDecompressionChildContentTypesMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceDecompressionChildContentTypesContentType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceDecompressionChildContentTypesCustomContentType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceDecompressionChildContentTypes(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceDecompressionChildContentTypesMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("content_type"); ok {
		t, err := expandLoadBalanceDecompressionChildContentTypesContentType(d, v, "content_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content_type"] = t
		}
	}

	if v, ok := d.GetOk("custom_content_type"); ok {
		t, err := expandLoadBalanceDecompressionChildContentTypesCustomContentType(d, v, "custom_content_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom_content_type"] = t
		}
	}

	return &obj, nil
}
