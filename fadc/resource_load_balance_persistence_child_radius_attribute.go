// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance persistence child radius attribute.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePersistenceChildRadiusAttribute() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePersistenceChildRadiusAttributeRead,
		Update: resourceLoadBalancePersistenceChildRadiusAttributeUpdate,
		Create: resourceLoadBalancePersistenceChildRadiusAttributeCreate,
		Delete: resourceLoadBalancePersistenceChildRadiusAttributeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vendor_type": &schema.Schema{
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
func resourceLoadBalancePersistenceChildRadiusAttributeCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	obj, err := getObjectLoadBalancePersistenceChildRadiusAttribute(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePersistenceChildRadiusAttribute resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalancePersistenceChildRadiusAttribute(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePersistenceChildRadiusAttribute resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalancePersistenceChildRadiusAttributeRead(d, m)
}
func resourceLoadBalancePersistenceChildRadiusAttributeUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalancePersistenceChildRadiusAttribute(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePersistenceChildRadiusAttribute resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePersistenceChildRadiusAttribute(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePersistenceChildRadiusAttribute resource: %v", err)
	}

	return resourceLoadBalancePersistenceChildRadiusAttributeRead(d, m)
}
func resourceLoadBalancePersistenceChildRadiusAttributeDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalancePersistenceChildRadiusAttribute(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePersistenceChildRadiusAttribute resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePersistenceChildRadiusAttributeRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildRadiusAttribute: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalancePersistenceChildRadiusAttribute(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePersistenceChildRadiusAttribute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePersistenceChildRadiusAttribute(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePersistenceChildRadiusAttribute resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePersistenceChildRadiusAttributeVendorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceChildRadiusAttributeVendorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceChildRadiusAttributeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceChildRadiusAttributeMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePersistenceChildRadiusAttribute(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vendor_id", flattenLoadBalancePersistenceChildRadiusAttributeVendorId(o["vendor_id"], d, "vendor_id", sv)); err != nil {
		if !fortiAPIPatch(o["vendor_id"]) {
			return fmt.Errorf("Error reading vendor_id: %v", err)
		}
	}

	if err = d.Set("vendor_type", flattenLoadBalancePersistenceChildRadiusAttributeVendorType(o["vendor_type"], d, "vendor_type", sv)); err != nil {
		if !fortiAPIPatch(o["vendor_type"]) {
			return fmt.Errorf("Error reading vendor_type: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalancePersistenceChildRadiusAttributeType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePersistenceChildRadiusAttributeMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePersistenceChildRadiusAttributeVendorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceChildRadiusAttributeVendorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceChildRadiusAttributeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceChildRadiusAttributeMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePersistenceChildRadiusAttribute(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vendor_id"); ok {
		t, err := expandLoadBalancePersistenceChildRadiusAttributeVendorId(d, v, "vendor_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor_id"] = t
		}
	}

	if v, ok := d.GetOk("vendor_type"); ok {
		t, err := expandLoadBalancePersistenceChildRadiusAttributeVendorType(d, v, "vendor_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor_type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalancePersistenceChildRadiusAttributeType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePersistenceChildRadiusAttributeMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
