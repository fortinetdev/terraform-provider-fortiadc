// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance caching.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceCaching() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceCachingRead,
		Update: resourceLoadBalanceCachingUpdate,
		Create: resourceLoadBalanceCachingCreate,
		Delete: resourceLoadBalanceCachingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cache_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_entries": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_object_size": &schema.Schema{
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
func resourceLoadBalanceCachingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCaching: type error")
	}

	obj, err := getObjectLoadBalanceCaching(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCaching resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceCaching(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCaching resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceCachingRead(d, m)
}
func resourceLoadBalanceCachingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceCaching(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCaching resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceCaching(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCaching resource: %v", err)
	}

	return resourceLoadBalanceCachingRead(d, m)
}
func resourceLoadBalanceCachingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceCaching(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceCaching resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceCachingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceCaching(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCaching resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceCaching(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCaching resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceCachingMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingMaxCacheSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingMaxEntries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingMaxObjectSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceCaching(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_age", flattenLoadBalanceCachingMaxAge(o["max_age"], d, "max_age", sv)); err != nil {
		if !fortiAPIPatch(o["max_age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_cache_size", flattenLoadBalanceCachingMaxCacheSize(o["max_cache_size"], d, "max_cache_size", sv)); err != nil {
		if !fortiAPIPatch(o["max_cache_size"]) {
			return fmt.Errorf("Error reading max_cache_size: %v", err)
		}
	}

	if err = d.Set("max_entries", flattenLoadBalanceCachingMaxEntries(o["max_entries"], d, "max_entries", sv)); err != nil {
		if !fortiAPIPatch(o["max_entries"]) {
			return fmt.Errorf("Error reading max_entries: %v", err)
		}
	}

	if err = d.Set("max_object_size", flattenLoadBalanceCachingMaxObjectSize(o["max_object_size"], d, "max_object_size", sv)); err != nil {
		if !fortiAPIPatch(o["max_object_size"]) {
			return fmt.Errorf("Error reading max_object_size: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceCachingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceCachingMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingMaxCacheSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingMaxEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingMaxObjectSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceCaching(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_age"); ok {
		t, err := expandLoadBalanceCachingMaxAge(d, v, "max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_age"] = t
		}
	}

	if v, ok := d.GetOk("max_cache_size"); ok {
		t, err := expandLoadBalanceCachingMaxCacheSize(d, v, "max_cache_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_cache_size"] = t
		}
	}

	if v, ok := d.GetOk("max_entries"); ok {
		t, err := expandLoadBalanceCachingMaxEntries(d, v, "max_entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_entries"] = t
		}
	}

	if v, ok := d.GetOk("max_object_size"); ok {
		t, err := expandLoadBalanceCachingMaxObjectSize(d, v, "max_object_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_object_size"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceCachingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
