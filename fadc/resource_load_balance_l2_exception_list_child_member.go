// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance l2 exception list child member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceL2ExceptionListChildMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceL2ExceptionListChildMemberRead,
		Update: resourceLoadBalanceL2ExceptionListChildMemberUpdate,
		Create: resourceLoadBalanceL2ExceptionListChildMemberCreate,
		Delete: resourceLoadBalanceL2ExceptionListChildMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
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
func resourceLoadBalanceL2ExceptionListChildMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceL2ExceptionListChildMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceL2ExceptionListChildMember: type error")
	}

	obj, err := getObjectLoadBalanceL2ExceptionListChildMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceL2ExceptionListChildMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceL2ExceptionListChildMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceL2ExceptionListChildMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceL2ExceptionListChildMemberRead(d, m)
}
func resourceLoadBalanceL2ExceptionListChildMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceL2ExceptionListChildMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceL2ExceptionListChildMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceL2ExceptionListChildMember resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceL2ExceptionListChildMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceL2ExceptionListChildMember resource: %v", err)
	}

	return resourceLoadBalanceL2ExceptionListChildMemberRead(d, m)
}
func resourceLoadBalanceL2ExceptionListChildMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceL2ExceptionListChildMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceL2ExceptionListChildMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceL2ExceptionListChildMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceL2ExceptionListChildMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceL2ExceptionListChildMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceL2ExceptionListChildMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceL2ExceptionListChildMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceL2ExceptionListChildMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceL2ExceptionListChildMember resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceL2ExceptionListChildMemberIpNetmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenLoadBalanceL2ExceptionListChildMemberComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceL2ExceptionListChildMemberHostPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceL2ExceptionListChildMemberType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceL2ExceptionListChildMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceL2ExceptionListChildMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip_netmask", flattenLoadBalanceL2ExceptionListChildMemberIpNetmask(o["ip_netmask"], d, "ip_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ip_netmask"]) {
			return fmt.Errorf("Error reading ip_netmask: %v", err)
		}
	}

	if err = d.Set("comments", flattenLoadBalanceL2ExceptionListChildMemberComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("host_pattern", flattenLoadBalanceL2ExceptionListChildMemberHostPattern(o["host_pattern"], d, "host_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["host_pattern"]) {
			return fmt.Errorf("Error reading host_pattern: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceL2ExceptionListChildMemberType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceL2ExceptionListChildMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceL2ExceptionListChildMemberIpNetmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceL2ExceptionListChildMemberComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceL2ExceptionListChildMemberHostPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceL2ExceptionListChildMemberType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceL2ExceptionListChildMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceL2ExceptionListChildMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip_netmask"); ok {
		t, err := expandLoadBalanceL2ExceptionListChildMemberIpNetmask(d, v, "ip_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_netmask"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandLoadBalanceL2ExceptionListChildMemberComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("host_pattern"); ok {
		t, err := expandLoadBalanceL2ExceptionListChildMemberHostPattern(d, v, "host_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_pattern"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceL2ExceptionListChildMemberType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceL2ExceptionListChildMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
