// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance decompression.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceDecompression() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceDecompressionRead,
		Update: resourceLoadBalanceDecompressionUpdate,
		Create: resourceLoadBalanceDecompressionCreate,
		Delete: resourceLoadBalanceDecompressionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_cpu_usage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"uri_list_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cpu_limit": &schema.Schema{
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
func resourceLoadBalanceDecompressionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceDecompression: type error")
	}

	obj, err := getObjectLoadBalanceDecompression(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceDecompression resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceDecompression(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceDecompression resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceDecompressionRead(d, m)
}
func resourceLoadBalanceDecompressionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceDecompression(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceDecompression resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceDecompression(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceDecompression resource: %v", err)
	}

	return resourceLoadBalanceDecompressionRead(d, m)
}
func resourceLoadBalanceDecompressionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceDecompression(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceDecompression resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceDecompressionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceDecompression(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceDecompression resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceDecompression(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceDecompression resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceDecompressionMaxCpuUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceDecompressionUriListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceDecompressionCpuLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceDecompressionMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceDecompressionMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceDecompression(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_cpu_usage", flattenLoadBalanceDecompressionMaxCpuUsage(o["max_cpu_usage"], d, "max_cpu_usage", sv)); err != nil {
		if !fortiAPIPatch(o["max_cpu_usage"]) {
			return fmt.Errorf("Error reading max_cpu_usage: %v", err)
		}
	}

	if err = d.Set("uri_list_type", flattenLoadBalanceDecompressionUriListType(o["uri_list_type"], d, "uri_list_type", sv)); err != nil {
		if !fortiAPIPatch(o["uri_list_type"]) {
			return fmt.Errorf("Error reading uri_list_type: %v", err)
		}
	}

	if err = d.Set("cpu_limit", flattenLoadBalanceDecompressionCpuLimit(o["cpu_limit"], d, "cpu_limit", sv)); err != nil {
		if !fortiAPIPatch(o["cpu_limit"]) {
			return fmt.Errorf("Error reading cpu_limit: %v", err)
		}
	}

	if err = d.Set("method", flattenLoadBalanceDecompressionMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceDecompressionMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceDecompressionMaxCpuUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceDecompressionUriListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceDecompressionCpuLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceDecompressionMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceDecompressionMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceDecompression(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_cpu_usage"); ok {
		t, err := expandLoadBalanceDecompressionMaxCpuUsage(d, v, "max_cpu_usage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_cpu_usage"] = t
		}
	}

	if v, ok := d.GetOk("uri_list_type"); ok {
		t, err := expandLoadBalanceDecompressionUriListType(d, v, "uri_list_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri_list_type"] = t
		}
	}

	if v, ok := d.GetOk("cpu_limit"); ok {
		t, err := expandLoadBalanceDecompressionCpuLimit(d, v, "cpu_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu_limit"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandLoadBalanceDecompressionMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceDecompressionMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
