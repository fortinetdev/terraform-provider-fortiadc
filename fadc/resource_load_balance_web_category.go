// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance web category.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceWebCategory() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceWebCategoryRead,
		Update: resourceLoadBalanceWebCategoryUpdate,
		Create: resourceLoadBalanceWebCategoryCreate,
		Delete: resourceLoadBalanceWebCategoryDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fadc_id": &schema.Schema{
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
func resourceLoadBalanceWebCategoryCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebCategory: type error")
	}

	obj, err := getObjectLoadBalanceWebCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebCategory resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceWebCategory(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebCategory resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceWebCategoryRead(d, m)
}
func resourceLoadBalanceWebCategoryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceWebCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebCategory resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceWebCategory(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebCategory resource: %v", err)
	}

	return resourceLoadBalanceWebCategoryRead(d, m)
}
func resourceLoadBalanceWebCategoryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceWebCategory(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceWebCategory resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceWebCategoryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceWebCategory(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceWebCategory(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebCategory resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceWebCategoryDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceWebCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceWebCategoryMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceWebCategory(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenLoadBalanceWebCategoryDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fadc_id", flattenLoadBalanceWebCategoryId(o["id"], d, "fadc_id", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceWebCategoryMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceWebCategoryDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceWebCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceWebCategoryMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceWebCategory(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandLoadBalanceWebCategoryDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fadc_id"); ok {
		t, err := expandLoadBalanceWebCategoryId(d, v, "fadc_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceWebCategoryMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
