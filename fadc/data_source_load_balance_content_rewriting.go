// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content rewriting.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceContentRewriting() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceContentRewritingRead,
		Schema: map[string]*schema.Schema{
			"redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"referer_content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"referer_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url_content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": &schema.Schema{
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
func dataSourceLoadBalanceContentRewritingRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRewriting: type error")
	}

	o, err := c.ReadLoadBalanceContentRewriting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceContentRewriting: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceContentRewriting(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceContentRewriting from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceContentRewritingRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingRefererContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingRefererStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingHostStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingUrlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingActionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingHeaderValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingUrlContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingHostContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceContentRewritingLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceContentRewriting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("redirect", dataSourceFlattenLoadBalanceContentRewritingRedirect(o["redirect"], d, "redirect")); err != nil {
		if !fortiAPIPatch(o["redirect"]) {
			return fmt.Errorf("Error reading redirect: %v", err)
		}
	}

	if err = d.Set("header_name", dataSourceFlattenLoadBalanceContentRewritingHeaderName(o["header_name"], d, "header_name")); err != nil {
		if !fortiAPIPatch(o["header_name"]) {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("referer_content", dataSourceFlattenLoadBalanceContentRewritingRefererContent(o["referer_content"], d, "referer_content")); err != nil {
		if !fortiAPIPatch(o["referer_content"]) {
			return fmt.Errorf("Error reading referer_content: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceContentRewritingMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("referer_status", dataSourceFlattenLoadBalanceContentRewritingRefererStatus(o["referer_status"], d, "referer_status")); err != nil {
		if !fortiAPIPatch(o["referer_status"]) {
			return fmt.Errorf("Error reading referer_status: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenLoadBalanceContentRewritingComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("host_status", dataSourceFlattenLoadBalanceContentRewritingHostStatus(o["host_status"], d, "host_status")); err != nil {
		if !fortiAPIPatch(o["host_status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}

	if err = d.Set("url_status", dataSourceFlattenLoadBalanceContentRewritingUrlStatus(o["url_status"], d, "url_status")); err != nil {
		if !fortiAPIPatch(o["url_status"]) {
			return fmt.Errorf("Error reading url_status: %v", err)
		}
	}

	if err = d.Set("action_type", dataSourceFlattenLoadBalanceContentRewritingActionType(o["action_type"], d, "action_type")); err != nil {
		if !fortiAPIPatch(o["action_type"]) {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if err = d.Set("header_value", dataSourceFlattenLoadBalanceContentRewritingHeaderValue(o["header_value"], d, "header_value")); err != nil {
		if !fortiAPIPatch(o["header_value"]) {
			return fmt.Errorf("Error reading header_value: %v", err)
		}
	}

	if err = d.Set("action", dataSourceFlattenLoadBalanceContentRewritingAction(o["action"], d, "action")); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("url_content", dataSourceFlattenLoadBalanceContentRewritingUrlContent(o["url_content"], d, "url_content")); err != nil {
		if !fortiAPIPatch(o["url_content"]) {
			return fmt.Errorf("Error reading url_content: %v", err)
		}
	}

	if err = d.Set("host_content", dataSourceFlattenLoadBalanceContentRewritingHostContent(o["host_content"], d, "host_content")); err != nil {
		if !fortiAPIPatch(o["host_content"]) {
			return fmt.Errorf("Error reading host_content: %v", err)
		}
	}

	if err = d.Set("location", dataSourceFlattenLoadBalanceContentRewritingLocation(o["location"], d, "location")); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	return nil
}
