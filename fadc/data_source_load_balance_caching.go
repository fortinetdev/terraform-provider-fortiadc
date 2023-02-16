// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance caching.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceCaching() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceCachingRead,
		Schema: map[string]*schema.Schema{
			"max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_cache_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_entries": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_object_size": &schema.Schema{
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
func dataSourceLoadBalanceCachingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalanceCaching(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCaching: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceCaching(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCaching from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceCachingMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingMaxCacheSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingMaxEntries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingMaxObjectSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceCaching(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_age", dataSourceFlattenLoadBalanceCachingMaxAge(o["max_age"], d, "max_age")); err != nil {
		if !fortiAPIPatch(o["max_age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_cache_size", dataSourceFlattenLoadBalanceCachingMaxCacheSize(o["max_cache_size"], d, "max_cache_size")); err != nil {
		if !fortiAPIPatch(o["max_cache_size"]) {
			return fmt.Errorf("Error reading max_cache_size: %v", err)
		}
	}

	if err = d.Set("max_entries", dataSourceFlattenLoadBalanceCachingMaxEntries(o["max_entries"], d, "max_entries")); err != nil {
		if !fortiAPIPatch(o["max_entries"]) {
			return fmt.Errorf("Error reading max_entries: %v", err)
		}
	}

	if err = d.Set("max_object_size", dataSourceFlattenLoadBalanceCachingMaxObjectSize(o["max_object_size"], d, "max_object_size")); err != nil {
		if !fortiAPIPatch(o["max_object_size"]) {
			return fmt.Errorf("Error reading max_object_size: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceCachingMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
