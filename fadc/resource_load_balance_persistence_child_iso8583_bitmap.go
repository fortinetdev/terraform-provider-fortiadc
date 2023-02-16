// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance persistence child iso8583 bitmap.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePersistenceChildIso8583Bitmap() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePersistenceChildIso8583BitmapRead,
		Update: resourceLoadBalancePersistenceChildIso8583BitmapUpdate,
		Create: resourceLoadBalancePersistenceChildIso8583BitmapCreate,
		Delete: resourceLoadBalancePersistenceChildIso8583BitmapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
func resourceLoadBalancePersistenceChildIso8583BitmapCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildIso8583Bitmap: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalancePersistenceChildIso8583Bitmap: type error")
	}

	obj, err := getObjectLoadBalancePersistenceChildIso8583Bitmap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePersistenceChildIso8583Bitmap resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalancePersistenceChildIso8583Bitmap(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePersistenceChildIso8583Bitmap resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalancePersistenceChildIso8583BitmapRead(d, m)
}
func resourceLoadBalancePersistenceChildIso8583BitmapUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildIso8583Bitmap: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalancePersistenceChildIso8583Bitmap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePersistenceChildIso8583Bitmap resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePersistenceChildIso8583Bitmap(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePersistenceChildIso8583Bitmap resource: %v", err)
	}

	return resourceLoadBalancePersistenceChildIso8583BitmapRead(d, m)
}
func resourceLoadBalancePersistenceChildIso8583BitmapDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildIso8583Bitmap: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalancePersistenceChildIso8583Bitmap(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePersistenceChildIso8583Bitmap resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePersistenceChildIso8583BitmapRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistenceChildIso8583Bitmap: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalancePersistenceChildIso8583Bitmap(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePersistenceChildIso8583Bitmap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePersistenceChildIso8583Bitmap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePersistenceChildIso8583Bitmap resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePersistenceChildIso8583BitmapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceChildIso8583BitmapMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePersistenceChildIso8583Bitmap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenLoadBalancePersistenceChildIso8583BitmapType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePersistenceChildIso8583BitmapMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePersistenceChildIso8583BitmapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceChildIso8583BitmapMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePersistenceChildIso8583Bitmap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalancePersistenceChildIso8583BitmapType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePersistenceChildIso8583BitmapMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
