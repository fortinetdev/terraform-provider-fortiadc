// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server remote dns server child memnber

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerRemoteDnsServerChildMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerRemoteDnsServerChildMemberRead,
		Update: resourceGlobalDnsServerRemoteDnsServerChildMemberUpdate,
		Create: resourceGlobalDnsServerRemoteDnsServerChildMemberCreate,
		Delete: resourceGlobalDnsServerRemoteDnsServerChildMemberDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
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

func resourceGlobalDnsServerRemoteDnsServerChildMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	obj, err := getObjectGlobalDnsServerRemoteDnsServerChildMember(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerRemoteDnsServerChildMember resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_remote_dns_server_child_member?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerRemoteDnsServerChildMember resource: %v", err)
	}

	id := "GlobalDnsServerRemoteDnsServerChildMember"

	d.SetId(id)

	return resourceGlobalDnsServerRemoteDnsServerChildMemberRead(d, m)
}

func resourceGlobalDnsServerRemoteDnsServerChildMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	obj, err := getObjectGlobalDnsServerRemoteDnsServerChildMember(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerRemoteDnsServerChildMember resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_remote_dns_server_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerRemoteDnsServerChildMember resource: %v", err)
	}

	return resourceGlobalDnsServerRemoteDnsServerChildMemberRead(d, m)
}

func resourceGlobalDnsServerRemoteDnsServerChildMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	path := "/api/global_dns_server_remote_dns_server_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerRemoteDnsServerChildMember resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalDnsServerRemoteDnsServerChildMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing GlobalDnsServerRemoteDnsServerChildMember: type error")
	}

	path := "/api/global_dns_server_remote_dns_server_child_member?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerRemoteDnsServerChildMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerRemoteDnsServerChildMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerRemoteDnsServerChildMember resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerRemoteDnsServerChildMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerRemoteDnsServerChildMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalDnsServerRemoteDnsServerChildMember(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenGlobalDnsServerRemoteDnsServerChildMember(o["addr_type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr_type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenGlobalDnsServerRemoteDnsServerChildMember(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenGlobalDnsServerRemoteDnsServerChildMember(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("port", flattenGlobalDnsServerRemoteDnsServerChildMember(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerRemoteDnsServerChildMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerRemoteDnsServerChildMember(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandGlobalDnsServerRemoteDnsServerChildMember(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandGlobalDnsServerRemoteDnsServerChildMember(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr_type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandGlobalDnsServerRemoteDnsServerChildMember(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandGlobalDnsServerRemoteDnsServerChildMember(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandGlobalDnsServerRemoteDnsServerChildMember(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	return &obj, nil
}
