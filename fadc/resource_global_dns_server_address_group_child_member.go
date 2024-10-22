// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server address group child member

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerAddressGroupChildMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerAddressGroupChildMemberRead,
		Update: resourceGlobalDnsServerAddressGroupChildMemberUpdate,
		Create: resourceGlobalDnsServerAddressGroupChildMemberCreate,
		Delete: resourceGlobalDnsServerAddressGroupChildMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
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

func resourceGlobalDnsServerAddressGroupChildMemberCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	obj, err := getObjectGlobalDnsServerAddressGroupChildMember(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerAddressGroupChildMember resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_address_group_child_member?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerAddressGroupChildMember resource: %v", err)
	}

	id := "GlobalDnsServerAddressGroupChildMember"

	d.SetId(id)

	return resourceGlobalDnsServerAddressGroupChildMemberRead(d, m)
}

func resourceGlobalDnsServerAddressGroupChildMemberUpdate(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()
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
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	obj, err := getObjectGlobalDnsServerAddressGroupChildMember(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerAddressGroupChildMember resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_address_group_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerAddressGroupChildMember resource: %v", err)
	}

	return resourceGlobalDnsServerAddressGroupChildMemberRead(d, m)
}

func resourceGlobalDnsServerAddressGroupChildMemberDelete(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

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
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	path := "/api/global_dns_server_address_group_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerAddressGroupChildMember resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalDnsServerAddressGroupChildMemberRead(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

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
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerAddressGroupChildMember: type error")
	}

	path := "/api/global_dns_server_address_group_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerAddressGroupChildMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerAddressGroupChildMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerAddressGroupChildMember resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerAddressGroupChildMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerAddressGroupChildMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalDnsServerAddressGroupChildMember(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenGlobalDnsServerAddressGroupChildMember(o["addr_type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr_type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("ip_network", flattenGlobalDnsServerAddressGroupChildMember(o["ip_network"], d, "ip_network", sv)); err != nil {
		if !fortiAPIPatch(o["ip_network"]) {
			return fmt.Errorf("Error reading ip_network: %v", err)
		}
	}

	if err = d.Set("ip6_network", flattenGlobalDnsServerAddressGroupChildMember(o["ip6_network"], d, "ip6_network", sv)); err != nil {
		if !fortiAPIPatch(o["ip6_network"]) {
			return fmt.Errorf("Error reading ip6_network: %v", err)
		}
	}

	if err = d.Set("action", flattenGlobalDnsServerAddressGroupChildMember(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerAddressGroupChildMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerAddressGroupChildMember(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandGlobalDnsServerAddressGroupChildMember(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandGlobalDnsServerAddressGroupChildMember(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr_type"] = t
		}
	}

	if v, ok := d.GetOk("ip_network"); ok {
		t, err := expandGlobalDnsServerAddressGroupChildMember(d, v, "ip_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_network"] = t
		}
	}

	if v, ok := d.GetOk("ip6_network"); ok {
		t, err := expandGlobalDnsServerAddressGroupChildMember(d, v, "ip6_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6_network"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandGlobalDnsServerAddressGroupChildMember(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	return &obj, nil
}
