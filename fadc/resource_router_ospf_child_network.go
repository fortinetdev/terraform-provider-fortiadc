// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child network.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfChildNetwork() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterOspfChildNetworkRead,
		Update: resourceRouterOspfChildNetworkUpdate,
		Create: resourceRouterOspfChildNetworkCreate,
		Delete: resourceRouterOspfChildNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"area_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix": &schema.Schema{
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
func resourceRouterOspfChildNetworkCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterOspfChildNetwork: type error")
	}

	obj, err := getObjectRouterOspfChildNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildNetwork resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfChildNetwork(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildNetwork resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterOspfChildNetworkRead(d, m)
}
func resourceRouterOspfChildNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterOspfChildNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildNetwork resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfChildNetwork(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildNetwork resource: %v", err)
	}

	return resourceRouterOspfChildNetworkRead(d, m)
}
func resourceRouterOspfChildNetworkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterOspfChildNetwork(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfChildNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterOspfChildNetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterOspfChildNetwork(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfChildNetwork(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildNetwork resource from API: %v", err)
	}
	return nil
}
func flattenRouterOspfChildNetworkAreaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfChildNetworkMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterOspfChildNetwork(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("area_id", flattenRouterOspfChildNetworkAreaId(o["area_id"], d, "area_id", sv)); err != nil {
		if !fortiAPIPatch(o["area_id"]) {
			return fmt.Errorf("Error reading area_id: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterOspfChildNetworkPrefix(o["prefix"], d, "prefix", sv)); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterOspfChildNetworkMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterOspfChildNetworkAreaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildNetworkMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfChildNetwork(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("area_id"); ok {
		t, err := expandRouterOspfChildNetworkAreaId(d, v, "area_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area_id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok {
		t, err := expandRouterOspfChildNetworkPrefix(d, v, "prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterOspfChildNetworkMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
