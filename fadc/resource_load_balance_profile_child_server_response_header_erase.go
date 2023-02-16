// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child server response header erase.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceProfileChildServerResponseHeaderErase() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceProfileChildServerResponseHeaderEraseRead,
		Update: resourceLoadBalanceProfileChildServerResponseHeaderEraseUpdate,
		Create: resourceLoadBalanceProfileChildServerResponseHeaderEraseCreate,
		Delete: resourceLoadBalanceProfileChildServerResponseHeaderEraseDelete,

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
func resourceLoadBalanceProfileChildServerResponseHeaderEraseCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderErase: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderErase: type error")
	}

	obj, err := getObjectLoadBalanceProfileChildServerResponseHeaderErase(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildServerResponseHeaderErase resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceProfileChildServerResponseHeaderErase(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildServerResponseHeaderErase resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceProfileChildServerResponseHeaderEraseRead(d, m)
}
func resourceLoadBalanceProfileChildServerResponseHeaderEraseUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderErase: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceProfileChildServerResponseHeaderErase(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildServerResponseHeaderErase resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceProfileChildServerResponseHeaderErase(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildServerResponseHeaderErase resource: %v", err)
	}

	return resourceLoadBalanceProfileChildServerResponseHeaderEraseRead(d, m)
}
func resourceLoadBalanceProfileChildServerResponseHeaderEraseDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderErase: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceProfileChildServerResponseHeaderErase(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceProfileChildServerResponseHeaderErase resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceProfileChildServerResponseHeaderEraseRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildServerResponseHeaderErase: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceProfileChildServerResponseHeaderErase(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildServerResponseHeaderErase resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceProfileChildServerResponseHeaderErase(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildServerResponseHeaderErase resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceProfileChildServerResponseHeaderEraseType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildServerResponseHeaderEraseMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildServerResponseHeaderEraseString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceProfileChildServerResponseHeaderErase(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenLoadBalanceProfileChildServerResponseHeaderEraseType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceProfileChildServerResponseHeaderEraseMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("string", flattenLoadBalanceProfileChildServerResponseHeaderEraseString(o["string"], d, "string", sv)); err != nil {
		if !fortiAPIPatch(o["string"]) {
			return fmt.Errorf("Error reading string: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceProfileChildServerResponseHeaderEraseType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildServerResponseHeaderEraseMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildServerResponseHeaderEraseString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceProfileChildServerResponseHeaderErase(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceProfileChildServerResponseHeaderEraseType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceProfileChildServerResponseHeaderEraseMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("string"); ok {
		t, err := expandLoadBalanceProfileChildServerResponseHeaderEraseString(d, v, "string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["string"] = t
		}
	}

	return &obj, nil
}
