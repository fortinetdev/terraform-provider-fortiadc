// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp child ha router id list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpChildHaRouterIdList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterBgpChildHaRouterIdListRead,
		Update: resourceRouterBgpChildHaRouterIdListUpdate,
		Create: resourceRouterBgpChildHaRouterIdListCreate,
		Delete: resourceRouterBgpChildHaRouterIdListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"router_id": &schema.Schema{
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
func resourceRouterBgpChildHaRouterIdListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterBgpChildHaRouterIdList: type error")
	}

	obj, err := getObjectRouterBgpChildHaRouterIdList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpChildHaRouterIdList resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpChildHaRouterIdList(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpChildHaRouterIdList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterBgpChildHaRouterIdListRead(d, m)
}
func resourceRouterBgpChildHaRouterIdListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterBgpChildHaRouterIdList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpChildHaRouterIdList resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpChildHaRouterIdList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpChildHaRouterIdList resource: %v", err)
	}

	return resourceRouterBgpChildHaRouterIdListRead(d, m)
}
func resourceRouterBgpChildHaRouterIdListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterBgpChildHaRouterIdList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpChildHaRouterIdList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterBgpChildHaRouterIdListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterBgpChildHaRouterIdList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpChildHaRouterIdList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpChildHaRouterIdList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpChildHaRouterIdList resource from API: %v", err)
	}
	return nil
}
func flattenRouterBgpChildHaRouterIdListNode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildHaRouterIdListRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildHaRouterIdListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBgpChildHaRouterIdList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("node", flattenRouterBgpChildHaRouterIdListNode(o["node"], d, "node", sv)); err != nil {
		if !fortiAPIPatch(o["node"]) {
			return fmt.Errorf("Error reading node: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterBgpChildHaRouterIdListRouterId(o["router_id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router_id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterBgpChildHaRouterIdListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterBgpChildHaRouterIdListNode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildHaRouterIdListRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildHaRouterIdListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpChildHaRouterIdList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("node"); ok {
		t, err := expandRouterBgpChildHaRouterIdListNode(d, v, "node", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		t, err := expandRouterBgpChildHaRouterIdListRouterId(d, v, "router_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router_id"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterBgpChildHaRouterIdListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
