// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Attach load balance virtual server content routing

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceVirtualServerContentRoutingAttach() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceVirtualServerContentRoutingAttachRead,
		Update: resourceLoadBalanceVirtualServerContentRoutingAttachUpdate,
		Create: resourceLoadBalanceVirtualServerContentRoutingAttachCreate,
		Delete: resourceLoadBalanceVirtualServerContentRoutingAttachDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"add_routing_list": &schema.Schema{
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

func resourceLoadBalanceVirtualServerContentRoutingAttachCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceVirtualServerContentRoutingAttach: type error")
	}

	o, err := c.ReadLoadBalanceVirtualServerContentRoutingAttach(mkey, vdom)
	current_routing_list, ok := o["content-routing-list"].(string)
	if !ok {
		fmt.Println("Error: 'name' is not a string")
		return fmt.Errorf("Error read current virtual server %v error", mkey)
	}

	obj, err := getObjectLoadBalanceVirtualServerContentRoutingAttach(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceVirtualServerContentRoutingAttach resource while getting object: %v", err)
	}

	items1 := ""
	if current_routing_list != "" {
		items1 = current_routing_list
	}

	items2 := ""
	if (*obj)["content-routing-list"] != nil {
		if val, ok := (*obj)["content-routing-list"].(string); ok {
			items2 = val
		}
	}

	list1 := strings.Split(items1, " ")
	list2 := strings.Split(items2, " ")
	list := append(list1, list2...)

	(*obj)["content-routing-list"] = strings.Join(list, " ")

	_, err = c.UpdateLoadBalanceVirtualServerContentRoutingAttach(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceVirtualServerContentRoutingAttach resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceVirtualServerContentRoutingAttachRead(d, m)
}

func resourceLoadBalanceVirtualServerContentRoutingAttachUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceVirtualServerContentRoutingAttach(mkey, vdom)
	current_routing_list, ok := o["content-routing-list"].(string)
	if !ok {
		fmt.Println("Error: 'name' is not a string")
		return fmt.Errorf("Error read current virtual server %v error", mkey)
	}

	obj, err := getObjectLoadBalanceVirtualServerContentRoutingAttach(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceVirtualServerContentRoutingAttach resource while getting object: %v ", err)
	}

	items1 := ""
	if current_routing_list != "" {
		items1 = current_routing_list
	}

	items2 := ""
	if (*obj)["content-routing-list"] != nil {
		if val, ok := (*obj)["content-routing-list"].(string); ok {
			items2 = val
		}
	}

	list1 := strings.Split(items1, " ")
	list2 := strings.Split(items2, " ")
	list := append(list1, list2...)

	(*obj)["content-routing-list"] = strings.Join(list, " ")

	_, err = c.UpdateLoadBalanceVirtualServerContentRoutingAttach(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceVirtualServerContentRoutingAttach resource: %v", err)
	}

	return resourceLoadBalanceVirtualServerContentRoutingAttachRead(d, m)
}

func resourceLoadBalanceVirtualServerContentRoutingAttachDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceVirtualServer(mkey, vdom)
	current_routing_list, ok := o["content-routing-list"].(string)
	if !ok {
		fmt.Println("Error: 'name' is not a string")
		return fmt.Errorf("Error read current virtual server %v error", mkey)
	}

	obj, err := getObjectLoadBalanceVirtualServerContentRoutingAttach(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceVirtualServerContentRoutingAttach resource while getting object: %v", err)
	}

	items1 := ""
	if current_routing_list != "" {
		items1 = current_routing_list
	}

	items2 := ""
	if (*obj)["content-routing-list"] != nil {
		if val, ok := (*obj)["content-routing-list"].(string); ok {
			items2 = val
		}
	}

	list1 := strings.Split(items1, " ")
	list2 := strings.Split(items2, " ")
	list2Map := make(map[string]bool)
	for _, item := range list2 {
		list2Map[item] = true
	}

	var list []string
	for _, item := range list1 {
		if item != "" && !list2Map[item] {
			list = append(list, item)
		}
	}

	(*obj)["content-routing-list"] = strings.Join(list, " ")

	_, err = c.UpdateLoadBalanceVirtualServer(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceVirtualServerContentRoutingAttach resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLoadBalanceVirtualServerContentRoutingAttachRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceVirtualServerContentRoutingAttach(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceVirtualServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceVirtualServerContentRoutingAttach(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceVirtualServer resource from API: %v", err)
	}
	return nil
}

func flattenLoadBalanceVirtualServerContentRoutingAttach(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceVirtualServerContentRoutingAttach(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	/*if err = d.Set("content_routing_list", flattenLoadBalanceVirtualServerContentRoutingAttach(o["content-routing-list"], d, "content_routing_list", sv)); err != nil {
		if !fortiAPIPatch(o["content-routing-list"]) {
			return fmt.Errorf("Error reading content-routing-list: %v", err)
		}
	}*/

	if err = d.Set("mkey", flattenLoadBalanceVirtualServerContentRoutingAttach(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceVirtualServerContentRoutingAttach(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceVirtualServerContentRoutingAttach(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_routing_list"); ok {
		t, err := expandLoadBalanceVirtualServerContentRoutingAttach(d, v, "add_routing_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-routing-list"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceVirtualServerContentRoutingAttach(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
