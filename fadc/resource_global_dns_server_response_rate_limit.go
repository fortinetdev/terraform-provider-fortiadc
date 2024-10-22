// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server response rate limit

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerResponseRateLimit() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerResponseRateLimitRead,
		Update: resourceGlobalDnsServerResponseRateLimitUpdate,
		Create: resourceGlobalDnsServerResponseRateLimitCreate,
		Delete: resourceGlobalDnsServerResponseRateLimitDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"per_second": &schema.Schema{
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

func resourceGlobalDnsServerResponseRateLimitCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerResponseRateLimit: type error")
	}

	obj, err := getObjectGlobalDnsServerResponseRateLimit(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerResponseRateLimit resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_response_rate_limit"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerResponseRateLimit resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalDnsServerResponseRateLimitRead(d, m)
}

func resourceGlobalDnsServerResponseRateLimitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalDnsServerResponseRateLimit(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerResponseRateLimit resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_response_rate_limit?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerResponseRateLimit resource: %v", err)
	}

	return resourceGlobalDnsServerResponseRateLimitRead(d, m)
}

func resourceGlobalDnsServerResponseRateLimitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_response_rate_limit?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerResponseRateLimit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalDnsServerResponseRateLimitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_response_rate_limit?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerResponseRateLimit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerResponseRateLimit(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerResponseRateLimit resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerResponseRateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerResponseRateLimit(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalDnsServerResponseRateLimit(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("per_second", flattenGlobalDnsServerResponseRateLimit(o["per_second"], d, "per_second", sv)); err != nil {
		if !fortiAPIPatch(o["per_second"]) {
			return fmt.Errorf("Error reading per_second: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerResponseRateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerResponseRateLimit(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalDnsServerResponseRateLimit(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("per_second"); ok {
		t, err := expandGlobalDnsServerResponseRateLimit(d, v, "per_second", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per_second"] = t
		}
	}

	return &obj, nil
}
