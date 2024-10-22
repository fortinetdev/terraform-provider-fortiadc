// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server policy

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerPolicyRead,
		Update: resourceGlobalDnsServerPolicyUpdate,
		Create: resourceGlobalDnsServerPolicyCreate,
		Delete: resourceGlobalDnsServerPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"zone_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns64_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"recursion_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dnssec_validate_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forwarders": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rrlimit": &schema.Schema{
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

func resourceGlobalDnsServerPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerPolicy: type error")
	}

	obj, err := getObjectGlobalDnsServerPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerPolicy resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_policy"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerPolicy resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceGlobalDnsServerPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalDnsServerPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerPolicy resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_policy?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerPolicy resource: %v", err)
	}

	return resourceGlobalDnsServerPolicyRead(d, m)
}

func resourceGlobalDnsServerPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_policy?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalDnsServerPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_policy?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerPolicy resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalDnsServerPolicy(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("source_address", flattenGlobalDnsServerPolicy(o["source_address"], d, "source_address", sv)); err != nil {
		if !fortiAPIPatch(o["source_address"]) {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	if err = d.Set("destination_address", flattenGlobalDnsServerPolicy(o["destination_address"], d, "destination_address", sv)); err != nil {
		if !fortiAPIPatch(o["destination_address"]) {
			return fmt.Errorf("Error reading destination_address: %v", err)
		}
	}

	if err = d.Set("zone_list", flattenGlobalDnsServerPolicy(o["zone_list"], d, "zone_list", sv)); err != nil {
		if !fortiAPIPatch(o["zone_list"]) {
			return fmt.Errorf("Error reading zone_list: %v", err)
		}
	}

	if err = d.Set("dns64_list", flattenGlobalDnsServerPolicy(o["dns64_list"], d, "dns64_list", sv)); err != nil {
		if !fortiAPIPatch(o["dns64_list"]) {
			return fmt.Errorf("Error reading dns64_list: %v", err)
		}
	}

	if err = d.Set("recursion_status", flattenGlobalDnsServerPolicy(o["recurision_status"], d, "recursion_status", sv)); err != nil {
		if !fortiAPIPatch(o["recurision_status"]) {
			return fmt.Errorf("Error reading recursion_status: %v", err)
		}
	}

	if err = d.Set("dnssec_validate_status", flattenGlobalDnsServerPolicy(o["dnssec_validate_status"], d, "dnssec_validate_status", sv)); err != nil {
		if !fortiAPIPatch(o["dnssec_validate_status"]) {
			return fmt.Errorf("Error reading dnssec_validate_status: %v", err)
		}
	}

	if err = d.Set("forward", flattenGlobalDnsServerPolicy(o["forward"], d, "forward", sv)); err != nil {
		if !fortiAPIPatch(o["forward"]) {
			return fmt.Errorf("Error reading forward: %v", err)
		}
	}

	if err = d.Set("forwarders", flattenGlobalDnsServerPolicy(o["forwarders"], d, "forwarders", sv)); err != nil {
		if !fortiAPIPatch(o["forwarders"]) {
			return fmt.Errorf("Error reading forwarders: %v", err)
		}
	}

	if err = d.Set("rrlimit", flattenGlobalDnsServerPolicy(o["rrlimit"], d, "rrlimit", sv)); err != nil {
		if !fortiAPIPatch(o["rrlimit"]) {
			return fmt.Errorf("Error reading rrlimit: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "source_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source_address"] = t
		}
	}

	if v, ok := d.GetOk("destination_address"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "destination_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination_address"] = t
		}
	}

	if v, ok := d.GetOk("zone_list"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "zone_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone_list"] = t
		}
	}

	if v, ok := d.GetOk("dns64_list"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "dns64_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns64_list"] = t
		}
	}

	if v, ok := d.GetOk("recursion_status"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "recursion_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recurision_status"] = t
		}
	}

	if v, ok := d.GetOk("dnssec_validate_status"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "dnssec_validate_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssec_validate_status"] = t
		}
	}

	if v, ok := d.GetOk("forward"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward"] = t
		}
	}

	if v, ok := d.GetOk("forwarders"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "forwarders", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarders"] = t
		}
	}

	if v, ok := d.GetOk("rrlimit"); ok {
		t, err := expandGlobalDnsServerPolicy(d, v, "rrlimit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rrlimit"] = t
		}
	}

	return &obj, nil
}
