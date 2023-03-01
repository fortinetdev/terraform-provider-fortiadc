// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child ospf interface.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRouterOspfChildOspfInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspfChildOspfInterfaceRead,
		Schema: map[string]*schema.Schema{
			"transmit_delay": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"retransmit_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"md5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceRouterOspfChildOspfInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfChildOspfInterface(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspfChildOspfInterface: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectRouterOspfChildOspfInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspfChildOspfInterface from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenRouterOspfChildOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceText(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceMd5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfChildOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspfChildOspfInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("transmit_delay", dataSourceFlattenRouterOspfChildOspfInterfaceTransmitDelay(o["transmit_delay"], d, "transmit_delay")); err != nil {
		if !fortiAPIPatch(o["transmit_delay"]) {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", dataSourceFlattenRouterOspfChildOspfInterfaceRetransmitInterval(o["retransmit_interval"], d, "retransmit_interval")); err != nil {
		if !fortiAPIPatch(o["retransmit_interval"]) {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", dataSourceFlattenRouterOspfChildOspfInterfaceHelloInterval(o["hello-interval"], d, "hello_interval")); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("Error reading hello-interval: %v", err)
		}
	}

	if err = d.Set("text", dataSourceFlattenRouterOspfChildOspfInterfaceText(o["text"], d, "text")); err != nil {
		if !fortiAPIPatch(o["text"]) {
			return fmt.Errorf("Error reading text: %v", err)
		}
	}

	if err = d.Set("mtu_ignore", dataSourceFlattenRouterOspfChildOspfInterfaceMtuIgnore(o["mtu_ignore"], d, "mtu_ignore")); err != nil {
		if !fortiAPIPatch(o["mtu_ignore"]) {
			return fmt.Errorf("Error reading mtu_ignore: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenRouterOspfChildOspfInterfacePriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("authentication", dataSourceFlattenRouterOspfChildOspfInterfaceAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("dead_interval", dataSourceFlattenRouterOspfChildOspfInterfaceDeadInterval(o["dead_interval"], d, "dead_interval")); err != nil {
		if !fortiAPIPatch(o["dead_interval"]) {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterOspfChildOspfInterfaceInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("md5", dataSourceFlattenRouterOspfChildOspfInterfaceMd5(o["md5"], d, "md5")); err != nil {
		if !fortiAPIPatch(o["md5"]) {
			return fmt.Errorf("Error reading md5: %v", err)
		}
	}

	if err = d.Set("network_type", dataSourceFlattenRouterOspfChildOspfInterfaceNetworkType(o["network_type"], d, "network_type")); err != nil {
		if !fortiAPIPatch(o["network_type"]) {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenRouterOspfChildOspfInterfaceMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("cost", dataSourceFlattenRouterOspfChildOspfInterfaceCost(o["cost"], d, "cost")); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	return nil
}
