// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content routing child match condition.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceContentRoutingChildMatchCondition() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceContentRoutingChildMatchConditionRead,
		Schema: map[string]*schema.Schema{
			"ignorecase": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reverse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": &schema.Schema{
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
func dataSourceLoadBalanceContentRoutingChildMatchConditionRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	o, err := c.ReadLoadBalanceContentRoutingChildMatchCondition(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceContentRoutingChildMatchCondition(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionIgnorecase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceContentRoutingChildMatchCondition(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ignorecase", dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionIgnorecase(o["ignorecase"], d, "ignorecase")); err != nil {
		if !fortiAPIPatch(o["ignorecase"]) {
			return fmt.Errorf("Error reading ignorecase: %v", err)
		}
	}

	if err = d.Set("reverse", dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionReverse(o["reverse"], d, "reverse")); err != nil {
		if !fortiAPIPatch(o["reverse"]) {
			return fmt.Errorf("Error reading reverse: %v", err)
		}
	}

	if err = d.Set("object", dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionObject(o["object"], d, "object")); err != nil {
		if !fortiAPIPatch(o["object"]) {
			return fmt.Errorf("Error reading object: %v", err)
		}
	}

	if err = d.Set("content", dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionContent(o["content"], d, "content")); err != nil {
		if !fortiAPIPatch(o["content"]) {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceContentRoutingChildMatchConditionMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
