// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance clone pool child pool member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceClonePoolChildPoolMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceClonePoolChildPoolMemberRead,
		Update: resourceLoadBalanceClonePoolChildPoolMemberUpdate,
		Create: resourceLoadBalanceClonePoolChildPoolMemberCreate,
		Delete: resourceLoadBalanceClonePoolChildPoolMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalanceClonePoolChildPoolMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	obj, err := getObjectLoadBalanceClonePoolChildPoolMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceClonePoolChildPoolMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceClonePoolChildPoolMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceClonePoolChildPoolMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceClonePoolChildPoolMemberRead(d, m)
}
func resourceLoadBalanceClonePoolChildPoolMemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceClonePoolChildPoolMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClonePoolChildPoolMember resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceClonePoolChildPoolMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClonePoolChildPoolMember resource: %v", err)
	}

	return resourceLoadBalanceClonePoolChildPoolMemberRead(d, m)
}
func resourceLoadBalanceClonePoolChildPoolMemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceClonePoolChildPoolMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceClonePoolChildPoolMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceClonePoolChildPoolMemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceClonePoolChildPoolMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceClonePoolChildPoolMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceClonePoolChildPoolMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceClonePoolChildPoolMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceClonePoolChildPoolMember resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceClonePoolChildPoolMemberDstMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClonePoolChildPoolMemberSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClonePoolChildPoolMemberAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClonePoolChildPoolMemberInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClonePoolChildPoolMemberPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClonePoolChildPoolMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClonePoolChildPoolMemberMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceClonePoolChildPoolMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dst_mac", flattenLoadBalanceClonePoolChildPoolMemberDstMac(o["dst_mac"], d, "dst_mac", sv)); err != nil {
		if !fortiAPIPatch(o["dst_mac"]) {
			return fmt.Errorf("Error reading dst_mac: %v", err)
		}
	}

	if err = d.Set("src_mac", flattenLoadBalanceClonePoolChildPoolMemberSrcMac(o["src_mac"], d, "src_mac", sv)); err != nil {
		if !fortiAPIPatch(o["src_mac"]) {
			return fmt.Errorf("Error reading src_mac: %v", err)
		}
	}

	if err = d.Set("address", flattenLoadBalanceClonePoolChildPoolMemberAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("interface", flattenLoadBalanceClonePoolChildPoolMemberInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("port", flattenLoadBalanceClonePoolChildPoolMemberPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceClonePoolChildPoolMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("mode", flattenLoadBalanceClonePoolChildPoolMemberMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceClonePoolChildPoolMemberDstMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClonePoolChildPoolMemberSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClonePoolChildPoolMemberAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClonePoolChildPoolMemberInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClonePoolChildPoolMemberPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClonePoolChildPoolMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClonePoolChildPoolMemberMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceClonePoolChildPoolMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_mac"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberDstMac(d, v, "dst_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst_mac"] = t
		}
	}

	if v, ok := d.GetOk("src_mac"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberSrcMac(d, v, "src_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src_mac"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandLoadBalanceClonePoolChildPoolMemberMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	return &obj, nil
}
