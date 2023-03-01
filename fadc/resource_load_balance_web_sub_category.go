// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance web sub category.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceWebSubCategory() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceWebSubCategoryRead,
		Update: resourceLoadBalanceWebSubCategoryUpdate,
		Create: resourceLoadBalanceWebSubCategoryCreate,
		Delete: resourceLoadBalanceWebSubCategoryDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fadc_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
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
func resourceLoadBalanceWebSubCategoryCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceWebSubCategory: type error")
	}

	obj, err := getObjectLoadBalanceWebSubCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebSubCategory resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceWebSubCategory(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWebSubCategory resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceWebSubCategoryRead(d, m)
}
func resourceLoadBalanceWebSubCategoryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceWebSubCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebSubCategory resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceWebSubCategory(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWebSubCategory resource: %v", err)
	}

	return resourceLoadBalanceWebSubCategoryRead(d, m)
}
func resourceLoadBalanceWebSubCategoryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceWebSubCategory(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceWebSubCategory resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceWebSubCategoryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceWebSubCategory(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebSubCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceWebSubCategory(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWebSubCategory resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceWebSubCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceWebSubCategoryMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceWebSubCategoryDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceWebSubCategory(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fadc_id", flattenLoadBalanceWebSubCategoryId(o["id"], d, "fadc_id", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceWebSubCategoryMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("description", flattenLoadBalanceWebSubCategoryDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceWebSubCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceWebSubCategoryMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceWebSubCategoryDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceWebSubCategory(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fadc_id"); ok {
		t, err := expandLoadBalanceWebSubCategoryId(d, v, "fadc_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceWebSubCategoryMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandLoadBalanceWebSubCategoryDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
