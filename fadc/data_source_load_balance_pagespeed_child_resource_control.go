// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pagespeed child resource control.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePagespeedChildResourceControl() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePagespeedChildResourceControlRead,
		Schema: map[string]*schema.Schema{
			"rewrite_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fetch_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"origin_domain_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceLoadBalancePagespeedChildResourceControlRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	o, err := c.ReadLoadBalancePagespeedChildResourceControl(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePagespeedChildResourceControl(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePagespeedChildResourceControlRewriteDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedChildResourceControlFetchDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedChildResourceControlMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePagespeedChildResourceControlOriginDomainPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePagespeedChildResourceControl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("rewrite_domain", dataSourceFlattenLoadBalancePagespeedChildResourceControlRewriteDomain(o["rewrite_domain"], d, "rewrite_domain")); err != nil {
		if !fortiAPIPatch(o["rewrite_domain"]) {
			return fmt.Errorf("Error reading rewrite_domain: %v", err)
		}
	}

	if err = d.Set("fetch_domain", dataSourceFlattenLoadBalancePagespeedChildResourceControlFetchDomain(o["fetch_domain"], d, "fetch_domain")); err != nil {
		if !fortiAPIPatch(o["fetch_domain"]) {
			return fmt.Errorf("Error reading fetch_domain: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePagespeedChildResourceControlMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("origin_domain_pattern", dataSourceFlattenLoadBalancePagespeedChildResourceControlOriginDomainPattern(o["origin_domain_pattern"], d, "origin_domain_pattern")); err != nil {
		if !fortiAPIPatch(o["origin_domain_pattern"]) {
			return fmt.Errorf("Error reading origin_domain_pattern: %v", err)
		}
	}

	return nil
}
