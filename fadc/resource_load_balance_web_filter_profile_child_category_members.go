// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance web filter profile child category members.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceWebFilterProfileChildCategoryMembers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceWebFilterProfileChildCategoryMembersRead,
		Update: resourceLoadBalanceWebFilterProfileChildCategoryMembersUpdate,
		Create: resourceLoadBalanceWebFilterProfileChildCategoryMembersCreate,
		Delete: resourceLoadBalanceWebFilterProfileChildCategoryMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
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
func resourceLoadBalanceWebFilterProfileChildCategoryMembersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	obj, err := getObjectLoadBalanceWebFilterProfileChildCategoryMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebFilterProfileChildCategoryMembers resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceWebFilterProfileChildCategoryMembers(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebFilterProfileChildCategoryMembers resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceWebFilterProfileChildCategoryMembersRead(d, m)
}
func resourceLoadBalanceWebFilterProfileChildCategoryMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceWebFilterProfileChildCategoryMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebFilterProfileChildCategoryMembers resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceWebFilterProfileChildCategoryMembers(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebFilterProfileChildCategoryMembers resource: %v", err)
	}

	return resourceLoadBalanceWebFilterProfileChildCategoryMembersRead(d, m)
}
func resourceLoadBalanceWebFilterProfileChildCategoryMembersDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceWebFilterProfileChildCategoryMembers(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceWebFilterProfileChildCategoryMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceWebFilterProfileChildCategoryMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebFilterProfileChildCategoryMembers: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceWebFilterProfileChildCategoryMembers(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebFilterProfileChildCategoryMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceWebFilterProfileChildCategoryMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebFilterProfileChildCategoryMembers resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceWebFilterProfileChildCategoryMembersCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceWebFilterProfileChildCategoryMembersMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceWebFilterProfileChildCategoryMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("category", flattenLoadBalanceWebFilterProfileChildCategoryMembersCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceWebFilterProfileChildCategoryMembersMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceWebFilterProfileChildCategoryMembersCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceWebFilterProfileChildCategoryMembersMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceWebFilterProfileChildCategoryMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok {
		t, err := expandLoadBalanceWebFilterProfileChildCategoryMembersCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceWebFilterProfileChildCategoryMembersMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
