// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pagespeed.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePagespeed() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePagespeedRead,
		Schema: map[string]*schema.Schema{
			"pagespeed_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_cache_max_inode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file_cache_max_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceLoadBalancePagespeedRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalancePagespeed(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePagespeed: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePagespeed(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePagespeed from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePagespeedPagespeedProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedFileCacheMaxInode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedFileCacheMaxSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePagespeed(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("pagespeed_profile_id", dataSourceFlattenLoadBalancePagespeedPagespeedProfileId(o["pagespeed_profile_id"], d, "pagespeed_profile_id")); err != nil {
		if !fortiAPIPatch(o["pagespeed_profile_id"]) {
			return fmt.Errorf("Error reading pagespeed_profile_id: %v", err)
		}
	}

	if err = d.Set("file_cache_max_inode", dataSourceFlattenLoadBalancePagespeedFileCacheMaxInode(o["file_cache_max_inode"], d, "file_cache_max_inode")); err != nil {
		if !fortiAPIPatch(o["file_cache_max_inode"]) {
			return fmt.Errorf("Error reading file_cache_max_inode: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePagespeedMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("file_cache_max_size", dataSourceFlattenLoadBalancePagespeedFileCacheMaxSize(o["file_cache_max_size"], d, "file_cache_max_size")); err != nil {
		if !fortiAPIPatch(o["file_cache_max_size"]) {
			return fmt.Errorf("Error reading file_cache_max_size: %v", err)
		}
	}

	return nil
}
