// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance caching child dyn cache list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceCachingChildDynCacheList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceCachingChildDynCacheListRead,
		Schema: map[string]*schema.Schema{
			"invalid_uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceLoadBalanceCachingChildDynCacheListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	o, err := c.ReadLoadBalanceCachingChildDynCacheList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceCachingChildDynCacheList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceCachingChildDynCacheListInvalidUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingChildDynCacheListAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingChildDynCacheListUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCachingChildDynCacheListMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceCachingChildDynCacheList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("invalid_uri", dataSourceFlattenLoadBalanceCachingChildDynCacheListInvalidUri(o["invalid_uri"], d, "invalid_uri")); err != nil {
		if !fortiAPIPatch(o["invalid_uri"]) {
			return fmt.Errorf("Error reading invalid_uri: %v", err)
		}
	}

	if err = d.Set("age", dataSourceFlattenLoadBalanceCachingChildDynCacheListAge(o["age"], d, "age")); err != nil {
		if !fortiAPIPatch(o["age"]) {
			return fmt.Errorf("Error reading age: %v", err)
		}
	}

	if err = d.Set("uri", dataSourceFlattenLoadBalanceCachingChildDynCacheListUri(o["uri"], d, "uri")); err != nil {
		if !fortiAPIPatch(o["uri"]) {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceCachingChildDynCacheListMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
