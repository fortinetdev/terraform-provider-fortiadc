// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp child network.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpChildNetwork() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterBgpChildNetworkRead,
		Update: resourceRouterBgpChildNetworkUpdate,
		Create: resourceRouterBgpChildNetworkCreate,
		Delete: resourceRouterBgpChildNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix": &schema.Schema{
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
		},
	}
}
func resourceRouterBgpChildNetworkCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterBgpChildNetwork: type error")
	}

	obj, err := getObjectRouterBgpChildNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpChildNetwork resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpChildNetwork(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpChildNetwork resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterBgpChildNetworkRead(d, m)
}
func resourceRouterBgpChildNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterBgpChildNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpChildNetwork resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpChildNetwork(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpChildNetwork resource: %v", err)
	}

	return resourceRouterBgpChildNetworkRead(d, m)
}
func resourceRouterBgpChildNetworkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterBgpChildNetwork(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpChildNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterBgpChildNetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterBgpChildNetwork(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpChildNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpChildNetwork(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpChildNetwork resource from API: %v", err)
	}
	return nil
}
func flattenRouterBgpChildNetworkPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBgpChildNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBgpChildNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNetworkMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBgpChildNetwork(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("prefix6", flattenRouterBgpChildNetworkPrefix6(o["prefix6"], d, "prefix6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix6"]) {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterBgpChildNetworkPrefix(o["prefix"], d, "prefix", sv)); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterBgpChildNetworkType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterBgpChildNetworkMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterBgpChildNetworkPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNetworkMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpChildNetwork(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("prefix6"); ok {
		t, err := expandRouterBgpChildNetworkPrefix6(d, v, "prefix6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok {
		t, err := expandRouterBgpChildNetworkPrefix(d, v, "prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandRouterBgpChildNetworkType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterBgpChildNetworkMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
