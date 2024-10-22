// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server address group

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerAddressGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerAddressGroupRead,
		Update: resourceGlobalDnsServerAddressGroupUpdate,
		Create: resourceGlobalDnsServerAddressGroupCreate,
		Delete: resourceGlobalDnsServerAddressGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceGlobalDnsServerAddressGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroup: type error")
	}

	obj, err := getObjectGlobalDnsServerAddressGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerAddressGroup resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_address_group"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerAddressGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalDnsServerAddressGroupRead(d, m)
}

func resourceGlobalDnsServerAddressGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalDnsServerAddressGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerAddressGroup resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_address_group?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerAddressGroup resource: %v", err)
	}

	return resourceGlobalDnsServerAddressGroupRead(d, m)
}

func resourceGlobalDnsServerAddressGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_address_group?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerAddressGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalDnsServerAddressGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_address_group?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerAddressGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerAddressGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerAddressGroup resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerAddressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerAddressGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalDnsServerAddressGroup(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerAddressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerAddressGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalDnsServerAddressGroup(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
