// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance compression.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceCompression() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceCompressionRead,
		Update: resourceLoadBalanceCompressionUpdate,
		Create: resourceLoadBalanceCompressionCreate,
		Delete: resourceLoadBalanceCompressionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_cpu_usage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"uri_list_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gzip_window_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gzip_memory_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cpu_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"min_content_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"method": &schema.Schema{
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
func resourceLoadBalanceCompressionCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceCompression(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCompression resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceCompression(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCompression resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceCompressionRead(d, m)
}
func resourceLoadBalanceCompressionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceCompression(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCompression resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceCompression(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCompression resource: %v", err)
	}

	return resourceLoadBalanceCompressionRead(d, m)
}
func resourceLoadBalanceCompressionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceCompression(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceCompression resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceCompressionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceCompression(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCompression resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceCompression(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCompression resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceCompressionMaxCpuUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionUriListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionGzipWindowSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionGzipMemoryLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionCpuLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionMinContentLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceCompression(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_cpu_usage", flattenLoadBalanceCompressionMaxCpuUsage(o["max_cpu_usage"], d, "max_cpu_usage", sv)); err != nil {
		if !fortiAPIPatch(o["max_cpu_usage"]) {
			return fmt.Errorf("Error reading max_cpu_usage: %v", err)
		}
	}

	if err = d.Set("level", flattenLoadBalanceCompressionLevel(o["level"], d, "level", sv)); err != nil {
		if !fortiAPIPatch(o["level"]) {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("uri_list_type", flattenLoadBalanceCompressionUriListType(o["uri_list_type"], d, "uri_list_type", sv)); err != nil {
		if !fortiAPIPatch(o["uri_list_type"]) {
			return fmt.Errorf("Error reading uri_list_type: %v", err)
		}
	}

	if err = d.Set("gzip_window_size", flattenLoadBalanceCompressionGzipWindowSize(o["gzip_window_size"], d, "gzip_window_size", sv)); err != nil {
		if !fortiAPIPatch(o["gzip_window_size"]) {
			return fmt.Errorf("Error reading gzip_window_size: %v", err)
		}
	}

	if err = d.Set("gzip_memory_level", flattenLoadBalanceCompressionGzipMemoryLevel(o["gzip_memory_level"], d, "gzip_memory_level", sv)); err != nil {
		if !fortiAPIPatch(o["gzip_memory_level"]) {
			return fmt.Errorf("Error reading gzip_memory_level: %v", err)
		}
	}

	if err = d.Set("cpu_limit", flattenLoadBalanceCompressionCpuLimit(o["cpu_limit"], d, "cpu_limit", sv)); err != nil {
		if !fortiAPIPatch(o["cpu_limit"]) {
			return fmt.Errorf("Error reading cpu_limit: %v", err)
		}
	}

	if err = d.Set("min_content_length", flattenLoadBalanceCompressionMinContentLength(o["min_content_length"], d, "min_content_length", sv)); err != nil {
		if !fortiAPIPatch(o["min_content_length"]) {
			return fmt.Errorf("Error reading min_content_length: %v", err)
		}
	}

	if err = d.Set("method", flattenLoadBalanceCompressionMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceCompressionMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceCompressionMaxCpuUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionUriListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionGzipWindowSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionGzipMemoryLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionCpuLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionMinContentLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceCompression(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_cpu_usage"); ok {
		t, err := expandLoadBalanceCompressionMaxCpuUsage(d, v, "max_cpu_usage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_cpu_usage"] = t
		}
	}

	if v, ok := d.GetOk("level"); ok {
		t, err := expandLoadBalanceCompressionLevel(d, v, "level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("uri_list_type"); ok {
		t, err := expandLoadBalanceCompressionUriListType(d, v, "uri_list_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri_list_type"] = t
		}
	}

	if v, ok := d.GetOk("gzip_window_size"); ok {
		t, err := expandLoadBalanceCompressionGzipWindowSize(d, v, "gzip_window_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gzip_window_size"] = t
		}
	}

	if v, ok := d.GetOk("gzip_memory_level"); ok {
		t, err := expandLoadBalanceCompressionGzipMemoryLevel(d, v, "gzip_memory_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gzip_memory_level"] = t
		}
	}

	if v, ok := d.GetOk("cpu_limit"); ok {
		t, err := expandLoadBalanceCompressionCpuLimit(d, v, "cpu_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu_limit"] = t
		}
	}

	if v, ok := d.GetOk("min_content_length"); ok {
		t, err := expandLoadBalanceCompressionMinContentLength(d, v, "min_content_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min_content_length"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandLoadBalanceCompressionMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceCompressionMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
