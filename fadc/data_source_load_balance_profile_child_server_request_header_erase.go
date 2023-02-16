// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child server request header erase.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceProfileChildServerRequestHeaderErase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceProfileChildServerRequestHeaderEraseRead,
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"string": &schema.Schema{
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
func dataSourceLoadBalanceProfileChildServerRequestHeaderEraseRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerRequestHeaderErase: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerRequestHeaderErase: type error")
	}

	o, err := c.ReadLoadBalanceProfileChildServerRequestHeaderErase(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerRequestHeaderErase: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceProfileChildServerRequestHeaderErase(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerRequestHeaderErase from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceProfileChildServerRequestHeaderEraseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildServerRequestHeaderEraseMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildServerRequestHeaderEraseString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceProfileChildServerRequestHeaderErase(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("type", dataSourceFlattenLoadBalanceProfileChildServerRequestHeaderEraseType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceProfileChildServerRequestHeaderEraseMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("string", dataSourceFlattenLoadBalanceProfileChildServerRequestHeaderEraseString(o["string"], d, "string")); err != nil {
		if !fortiAPIPatch(o["string"]) {
			return fmt.Errorf("Error reading string: %v", err)
		}
	}

	return nil
}
