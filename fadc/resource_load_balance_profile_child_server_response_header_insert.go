// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child server response header insert.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceProfileChildServerResponseHeaderInsert() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceProfileChildServerResponseHeaderInsertRead,
		Update: resourceLoadBalanceProfileChildServerResponseHeaderInsertUpdate,
		Create: resourceLoadBalanceProfileChildServerResponseHeaderInsertCreate,
		Delete: resourceLoadBalanceProfileChildServerResponseHeaderInsertDelete,

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
			"string": &schema.Schema{
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
func resourceLoadBalanceProfileChildServerResponseHeaderInsertCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderInsert: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderInsert: type error")
	}

	obj, err := getObjectLoadBalanceProfileChildServerResponseHeaderInsert(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildServerResponseHeaderInsert resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceProfileChildServerResponseHeaderInsert(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildServerResponseHeaderInsert resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceProfileChildServerResponseHeaderInsertRead(d, m)
}
func resourceLoadBalanceProfileChildServerResponseHeaderInsertUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderInsert: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceProfileChildServerResponseHeaderInsert(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildServerResponseHeaderInsert resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceProfileChildServerResponseHeaderInsert(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildServerResponseHeaderInsert resource: %v", err)
	}

	return resourceLoadBalanceProfileChildServerResponseHeaderInsertRead(d, m)
}
func resourceLoadBalanceProfileChildServerResponseHeaderInsertDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderInsert: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceProfileChildServerResponseHeaderInsert(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceProfileChildServerResponseHeaderInsert resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceProfileChildServerResponseHeaderInsertRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderInsert: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceProfileChildServerResponseHeaderInsert(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildServerResponseHeaderInsert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceProfileChildServerResponseHeaderInsert(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildServerResponseHeaderInsert resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceProfileChildServerResponseHeaderInsertType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildServerResponseHeaderInsertMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildServerResponseHeaderInsertString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceProfileChildServerResponseHeaderInsert(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenLoadBalanceProfileChildServerResponseHeaderInsertType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceProfileChildServerResponseHeaderInsertMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("string", flattenLoadBalanceProfileChildServerResponseHeaderInsertString(o["string"], d, "string", sv)); err != nil {
		if !fortiAPIPatch(o["string"]) {
			return fmt.Errorf("Error reading string: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceProfileChildServerResponseHeaderInsertType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildServerResponseHeaderInsertMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildServerResponseHeaderInsertString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceProfileChildServerResponseHeaderInsert(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceProfileChildServerResponseHeaderInsertType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceProfileChildServerResponseHeaderInsertMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("string"); ok {
		t, err := expandLoadBalanceProfileChildServerResponseHeaderInsertString(d, v, "string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["string"] = t
		}
	}

	return &obj, nil
}
