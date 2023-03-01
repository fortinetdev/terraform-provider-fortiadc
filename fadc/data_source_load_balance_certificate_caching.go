// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance certificate caching.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceCertificateCaching() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceCertificateCachingRead,
		Schema: map[string]*schema.Schema{
			"max_entries": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_certificate_cache_size": &schema.Schema{
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
func dataSourceLoadBalanceCertificateCachingRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCertificateCaching: type error")
	}

	o, err := c.ReadLoadBalanceCertificateCaching(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCertificateCaching: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceCertificateCaching(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceCertificateCaching from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceCertificateCachingMaxEntries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCertificateCachingMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceCertificateCachingMaxCertificateCacheSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceCertificateCaching(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_entries", dataSourceFlattenLoadBalanceCertificateCachingMaxEntries(o["max_entries"], d, "max_entries")); err != nil {
		if !fortiAPIPatch(o["max_entries"]) {
			return fmt.Errorf("Error reading max_entries: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceCertificateCachingMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("max_certificate_cache_size", dataSourceFlattenLoadBalanceCertificateCachingMaxCertificateCacheSize(o["max_certificate_cache_size"], d, "max_certificate_cache_size")); err != nil {
		if !fortiAPIPatch(o["max_certificate_cache_size"]) {
			return fmt.Errorf("Error reading max_certificate_cache_size: %v", err)
		}
	}

	return nil
}
