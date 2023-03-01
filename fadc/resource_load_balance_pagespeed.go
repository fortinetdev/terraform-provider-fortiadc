// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pagespeed.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePagespeed() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePagespeedRead,
		Update: resourceLoadBalancePagespeedUpdate,
		Create: resourceLoadBalancePagespeedCreate,
		Delete: resourceLoadBalancePagespeedDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pagespeed_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file_cache_max_inode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file_cache_max_size": &schema.Schema{
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
func resourceLoadBalancePagespeedCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeed: type error")
	}

	obj, err := getObjectLoadBalancePagespeed(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePagespeed resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalancePagespeed(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePagespeed resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalancePagespeedRead(d, m)
}
func resourceLoadBalancePagespeedUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalancePagespeed(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePagespeed resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePagespeed(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePagespeed resource: %v", err)
	}

	return resourceLoadBalancePagespeedRead(d, m)
}
func resourceLoadBalancePagespeedDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalancePagespeed(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePagespeed resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePagespeedRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalancePagespeed(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePagespeed resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePagespeed(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePagespeed resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePagespeedPagespeedProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedFileCacheMaxInode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedFileCacheMaxSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePagespeed(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("pagespeed_profile_id", flattenLoadBalancePagespeedPagespeedProfileId(o["pagespeed_profile_id"], d, "pagespeed_profile_id", sv)); err != nil {
		if !fortiAPIPatch(o["pagespeed_profile_id"]) {
			return fmt.Errorf("Error reading pagespeed_profile_id: %v", err)
		}
	}

	if err = d.Set("file_cache_max_inode", flattenLoadBalancePagespeedFileCacheMaxInode(o["file_cache_max_inode"], d, "file_cache_max_inode", sv)); err != nil {
		if !fortiAPIPatch(o["file_cache_max_inode"]) {
			return fmt.Errorf("Error reading file_cache_max_inode: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePagespeedMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("file_cache_max_size", flattenLoadBalancePagespeedFileCacheMaxSize(o["file_cache_max_size"], d, "file_cache_max_size", sv)); err != nil {
		if !fortiAPIPatch(o["file_cache_max_size"]) {
			return fmt.Errorf("Error reading file_cache_max_size: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePagespeedPagespeedProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedFileCacheMaxInode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedFileCacheMaxSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePagespeed(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pagespeed_profile_id"); ok {
		t, err := expandLoadBalancePagespeedPagespeedProfileId(d, v, "pagespeed_profile_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pagespeed_profile_id"] = t
		}
	}

	if v, ok := d.GetOk("file_cache_max_inode"); ok {
		t, err := expandLoadBalancePagespeedFileCacheMaxInode(d, v, "file_cache_max_inode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file_cache_max_inode"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePagespeedMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("file_cache_max_size"); ok {
		t, err := expandLoadBalancePagespeedFileCacheMaxSize(d, v, "file_cache_max_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file_cache_max_size"] = t
		}
	}

	return &obj, nil
}
