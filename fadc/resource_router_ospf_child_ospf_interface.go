// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child ospf interface.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfChildOspfInterface() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterOspfChildOspfInterfaceRead,
		Update: resourceRouterOspfChildOspfInterfaceUpdate,
		Create: resourceRouterOspfChildOspfInterfaceCreate,
		Delete: resourceRouterOspfChildOspfInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"transmit_delay": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"retransmit_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mtu_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"md5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cost": &schema.Schema{
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
func resourceRouterOspfChildOspfInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterOspfChildOspfInterface: type error")
	}

	obj, err := getObjectRouterOspfChildOspfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildOspfInterface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfChildOspfInterface(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildOspfInterface resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterOspfChildOspfInterfaceRead(d, m)
}
func resourceRouterOspfChildOspfInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterOspfChildOspfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildOspfInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfChildOspfInterface(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildOspfInterface resource: %v", err)
	}

	return resourceRouterOspfChildOspfInterfaceRead(d, m)
}
func resourceRouterOspfChildOspfInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterOspfChildOspfInterface(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfChildOspfInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterOspfChildOspfInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterOspfChildOspfInterface(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildOspfInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfChildOspfInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildOspfInterface resource from API: %v", err)
	}
	return nil
}
func flattenRouterOspfChildOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceText(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceMd5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterOspfChildOspfInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("transmit_delay", flattenRouterOspfChildOspfInterfaceTransmitDelay(o["transmit_delay"], d, "transmit_delay", sv)); err != nil {
		if !fortiAPIPatch(o["transmit_delay"]) {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterOspfChildOspfInterfaceRetransmitInterval(o["retransmit_interval"], d, "retransmit_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retransmit_interval"]) {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterOspfChildOspfInterfaceHelloInterval(o["hello-interval"], d, "hello_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("Error reading hello-interval: %v", err)
		}
	}

	if err = d.Set("text", flattenRouterOspfChildOspfInterfaceText(o["text"], d, "text", sv)); err != nil {
		if !fortiAPIPatch(o["text"]) {
			return fmt.Errorf("Error reading text: %v", err)
		}
	}

	if err = d.Set("mtu_ignore", flattenRouterOspfChildOspfInterfaceMtuIgnore(o["mtu_ignore"], d, "mtu_ignore", sv)); err != nil {
		if !fortiAPIPatch(o["mtu_ignore"]) {
			return fmt.Errorf("Error reading mtu_ignore: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterOspfChildOspfInterfacePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("authentication", flattenRouterOspfChildOspfInterfaceAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterOspfChildOspfInterfaceDeadInterval(o["dead_interval"], d, "dead_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dead_interval"]) {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterOspfChildOspfInterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("md5", flattenRouterOspfChildOspfInterfaceMd5(o["md5"], d, "md5", sv)); err != nil {
		if !fortiAPIPatch(o["md5"]) {
			return fmt.Errorf("Error reading md5: %v", err)
		}
	}

	if err = d.Set("network_type", flattenRouterOspfChildOspfInterfaceNetworkType(o["network_type"], d, "network_type", sv)); err != nil {
		if !fortiAPIPatch(o["network_type"]) {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterOspfChildOspfInterfaceMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterOspfChildOspfInterfaceCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	return nil
}

func expandRouterOspfChildOspfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceText(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceMd5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildOspfInterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfChildOspfInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("transmit_delay"); ok {
		t, err := expandRouterOspfChildOspfInterfaceTransmitDelay(d, v, "transmit_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit_delay"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok {
		t, err := expandRouterOspfChildOspfInterfaceRetransmitInterval(d, v, "retransmit_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit_interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok {
		t, err := expandRouterOspfChildOspfInterfaceHelloInterval(d, v, "hello_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("text"); ok {
		t, err := expandRouterOspfChildOspfInterfaceText(d, v, "text", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["text"] = t
		}
	}

	if v, ok := d.GetOk("mtu_ignore"); ok {
		t, err := expandRouterOspfChildOspfInterfaceMtuIgnore(d, v, "mtu_ignore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu_ignore"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandRouterOspfChildOspfInterfacePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandRouterOspfChildOspfInterfaceAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok {
		t, err := expandRouterOspfChildOspfInterfaceDeadInterval(d, v, "dead_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead_interval"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterOspfChildOspfInterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("md5"); ok {
		t, err := expandRouterOspfChildOspfInterfaceMd5(d, v, "md5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok {
		t, err := expandRouterOspfChildOspfInterfaceNetworkType(d, v, "network_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network_type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterOspfChildOspfInterfaceMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok {
		t, err := expandRouterOspfChildOspfInterfaceCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	return &obj, nil
}
