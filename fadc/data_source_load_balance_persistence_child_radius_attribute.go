// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance persistence child radius attribute.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePersistenceChildRadiusAttribute() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePersistenceChildRadiusAttributeRead,
		Schema: map[string]*schema.Schema{
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
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
func dataSourceLoadBalancePersistenceChildRadiusAttributeRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	o, err := c.ReadLoadBalancePersistenceChildRadiusAttribute(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePersistenceChildRadiusAttribute(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeVendorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeVendorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePersistenceChildRadiusAttribute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vendor_id", dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeVendorId(o["vendor_id"], d, "vendor_id")); err != nil {
		if !fortiAPIPatch(o["vendor_id"]) {
			return fmt.Errorf("Error reading vendor_id: %v", err)
		}
	}

	if err = d.Set("vendor_type", dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeVendorType(o["vendor_type"], d, "vendor_type")); err != nil {
		if !fortiAPIPatch(o["vendor_type"]) {
			return fmt.Errorf("Error reading vendor_type: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePersistenceChildRadiusAttributeMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
