// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance compression child uri list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceCompressionChildUriList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceCompressionChildUriListRead,
		Update: resourceLoadBalanceCompressionChildUriListUpdate,
		Create: resourceLoadBalanceCompressionChildUriListCreate,
		Delete: resourceLoadBalanceCompressionChildUriListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": &schema.Schema{
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
func resourceLoadBalanceCompressionChildUriListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCompressionChildUriList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceCompressionChildUriList: type error")
	}

	obj, err := getObjectLoadBalanceCompressionChildUriList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCompressionChildUriList resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceCompressionChildUriList(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCompressionChildUriList resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceCompressionChildUriListRead(d, m)
}
func resourceLoadBalanceCompressionChildUriListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCompressionChildUriList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceCompressionChildUriList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCompressionChildUriList resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceCompressionChildUriList(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCompressionChildUriList resource: %v", err)
	}

	return resourceLoadBalanceCompressionChildUriListRead(d, m)
}
func resourceLoadBalanceCompressionChildUriListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCompressionChildUriList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceCompressionChildUriList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceCompressionChildUriList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceCompressionChildUriListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCompressionChildUriList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceCompressionChildUriList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCompressionChildUriList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceCompressionChildUriList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCompressionChildUriList resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceCompressionChildUriListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCompressionChildUriListUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceCompressionChildUriList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenLoadBalanceCompressionChildUriListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("uri", flattenLoadBalanceCompressionChildUriListUri(o["uri"], d, "uri", sv)); err != nil {
		if !fortiAPIPatch(o["uri"]) {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceCompressionChildUriListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCompressionChildUriListUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceCompressionChildUriList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceCompressionChildUriListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("uri"); ok {
		t, err := expandLoadBalanceCompressionChildUriListUri(d, v, "uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri"] = t
		}
	}

	return &obj, nil
}
