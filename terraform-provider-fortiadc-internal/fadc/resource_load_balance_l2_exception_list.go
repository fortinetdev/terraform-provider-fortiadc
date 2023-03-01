// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance l2 exception list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceL2ExceptionList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceL2ExceptionListRead,
		Update: resourceLoadBalanceL2ExceptionListUpdate,
		Create: resourceLoadBalanceL2ExceptionListCreate,
		Delete: resourceLoadBalanceL2ExceptionListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_filter_profile": &schema.Schema{
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
		},
	}
}
func resourceLoadBalanceL2ExceptionListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceL2ExceptionList: type error")
	}

	obj, err := getObjectLoadBalanceL2ExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceL2ExceptionList resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceL2ExceptionList(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceL2ExceptionList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceL2ExceptionListRead(d, m)
}
func resourceLoadBalanceL2ExceptionListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceL2ExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceL2ExceptionList resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceL2ExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceL2ExceptionList resource: %v", err)
	}

	return resourceLoadBalanceL2ExceptionListRead(d, m)
}
func resourceLoadBalanceL2ExceptionListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceL2ExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceL2ExceptionList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceL2ExceptionListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceL2ExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceL2ExceptionList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceL2ExceptionList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceL2ExceptionList resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceL2ExceptionListDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceL2ExceptionListWebFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceL2ExceptionListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceL2ExceptionList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenLoadBalanceL2ExceptionListDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("web_filter_profile", flattenLoadBalanceL2ExceptionListWebFilterProfile(o["web_filter_profile"], d, "web_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["web_filter_profile"]) {
			return fmt.Errorf("Error reading web_filter_profile: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceL2ExceptionListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceL2ExceptionListDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceL2ExceptionListWebFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceL2ExceptionListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceL2ExceptionList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandLoadBalanceL2ExceptionListDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_profile"); ok {
		t, err := expandLoadBalanceL2ExceptionListWebFilterProfile(d, v, "web_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web_filter_profile"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceL2ExceptionListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
