// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance ippool child node member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceIppoolChildNodeMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceIppoolChildNodeMemberRead,
		Update: resourceLoadBalanceIppoolChildNodeMemberUpdate,
		Create: resourceLoadBalanceIppoolChildNodeMemberCreate,
		Delete: resourceLoadBalanceIppoolChildNodeMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip6_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ha_node_num": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalanceIppoolChildNodeMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	obj, err := getObjectLoadBalanceIppoolChildNodeMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceIppoolChildNodeMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceIppoolChildNodeMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceIppoolChildNodeMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceIppoolChildNodeMemberRead(d, m)
}
func resourceLoadBalanceIppoolChildNodeMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceIppoolChildNodeMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceIppoolChildNodeMember resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceIppoolChildNodeMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceIppoolChildNodeMember resource: %v", err)
	}

	return resourceLoadBalanceIppoolChildNodeMemberRead(d, m)
}
func resourceLoadBalanceIppoolChildNodeMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceIppoolChildNodeMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceIppoolChildNodeMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceIppoolChildNodeMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppoolChildNodeMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceIppoolChildNodeMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceIppoolChildNodeMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceIppoolChildNodeMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceIppoolChildNodeMember resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceIppoolChildNodeMemberIp6Max(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberIp6Min(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberHaNodeNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberIpMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberPoolType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberIpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolChildNodeMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceIppoolChildNodeMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip6_max", flattenLoadBalanceIppoolChildNodeMemberIp6Max(o["ip6_max"], d, "ip6_max", sv)); err != nil {
		if !fortiAPIPatch(o["ip6_max"]) {
			return fmt.Errorf("Error reading ip6_max: %v", err)
		}
	}

	if err = d.Set("ip6_min", flattenLoadBalanceIppoolChildNodeMemberIp6Min(o["ip6_min"], d, "ip6_min", sv)); err != nil {
		if !fortiAPIPatch(o["ip6_min"]) {
			return fmt.Errorf("Error reading ip6_min: %v", err)
		}
	}

	if err = d.Set("ha_node_num", flattenLoadBalanceIppoolChildNodeMemberHaNodeNum(o["ha_node_num"], d, "ha_node_num", sv)); err != nil {
		if !fortiAPIPatch(o["ha_node_num"]) {
			return fmt.Errorf("Error reading ha_node_num: %v", err)
		}
	}

	if err = d.Set("ip_min", flattenLoadBalanceIppoolChildNodeMemberIpMin(o["ip_min"], d, "ip_min", sv)); err != nil {
		if !fortiAPIPatch(o["ip_min"]) {
			return fmt.Errorf("Error reading ip_min: %v", err)
		}
	}

	if err = d.Set("pool_type", flattenLoadBalanceIppoolChildNodeMemberPoolType(o["pool_type"], d, "pool_type", sv)); err != nil {
		if !fortiAPIPatch(o["pool_type"]) {
			return fmt.Errorf("Error reading pool_type: %v", err)
		}
	}

	if err = d.Set("ip_max", flattenLoadBalanceIppoolChildNodeMemberIpMax(o["ip_max"], d, "ip_max", sv)); err != nil {
		if !fortiAPIPatch(o["ip_max"]) {
			return fmt.Errorf("Error reading ip_max: %v", err)
		}
	}

	if err = d.Set("interface", flattenLoadBalanceIppoolChildNodeMemberInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceIppoolChildNodeMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceIppoolChildNodeMemberIp6Max(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberIp6Min(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberHaNodeNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberIpMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberPoolType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberIpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolChildNodeMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceIppoolChildNodeMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip6_max"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberIp6Max(d, v, "ip6_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6_max"] = t
		}
	}

	if v, ok := d.GetOk("ip6_min"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberIp6Min(d, v, "ip6_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6_min"] = t
		}
	}

	if v, ok := d.GetOk("ha_node_num"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberHaNodeNum(d, v, "ha_node_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha_node_num"] = t
		}
	}

	if v, ok := d.GetOk("ip_min"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberIpMin(d, v, "ip_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_min"] = t
		}
	}

	if v, ok := d.GetOk("pool_type"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberPoolType(d, v, "pool_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool_type"] = t
		}
	}

	if v, ok := d.GetOk("ip_max"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberIpMax(d, v, "ip_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_max"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceIppoolChildNodeMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
