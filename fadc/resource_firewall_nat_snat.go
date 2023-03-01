// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  firewall nat snat.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallNatSnat() *schema.Resource {
	return &schema.Resource{
		Read:   resourceFirewallNatSnatRead,
		Update: resourceFirewallNatSnatUpdate,
		Create: resourceFirewallNatSnatCreate,
		Delete: resourceFirewallNatSnatDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"out_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trans_to_ip_end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trans_to_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trans_to_ip_start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trans_to_ip": &schema.Schema{
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
func resourceFirewallNatSnatCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallNatSnat: type error")
	}

	obj, err := getObjectFirewallNatSnat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallNatSnat resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallNatSnat(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating FirewallNatSnat resource: %v", err)
	}

	d.SetId(mkey)

	return resourceFirewallNatSnatRead(d, m)
}
func resourceFirewallNatSnatUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectFirewallNatSnat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallNatSnat resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallNatSnat(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating FirewallNatSnat resource: %v", err)
	}

	return resourceFirewallNatSnatRead(d, m)
}
func resourceFirewallNatSnatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteFirewallNatSnat(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallNatSnat resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceFirewallNatSnatRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadFirewallNatSnat(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading FirewallNatSnat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallNatSnat(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallNatSnat resource from API: %v", err)
	}
	return nil
}
func flattenFirewallNatSnatStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatOutInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatFrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenFirewallNatSnatTrafficGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatTransToIpEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenFirewallNatSnatTransToType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatTransToIpStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatTransToIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallNatSnatMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallNatSnat(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenFirewallNatSnatStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("out_interface", flattenFirewallNatSnatOutInterface(o["out-interface"], d, "out_interface", sv)); err != nil {
		if !fortiAPIPatch(o["out-interface"]) {
			return fmt.Errorf("Error reading out-interface: %v", err)
		}
	}

	if err = d.Set("from", flattenFirewallNatSnatFrom(o["from"], d, "from", sv)); err != nil {
		if !fortiAPIPatch(o["from"]) {
			return fmt.Errorf("Error reading from: %v", err)
		}
	}

	if err = d.Set("traffic_group", flattenFirewallNatSnatTrafficGroup(o["traffic-group"], d, "traffic_group", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("trans_to_ip_end", flattenFirewallNatSnatTransToIpEnd(o["trans-to-ip-end"], d, "trans_to_ip_end", sv)); err != nil {
		if !fortiAPIPatch(o["trans-to-ip-end"]) {
			return fmt.Errorf("Error reading trans-to-ip-end: %v", err)
		}
	}

	if err = d.Set("to", flattenFirewallNatSnatTo(o["to"], d, "to", sv)); err != nil {
		if !fortiAPIPatch(o["to"]) {
			return fmt.Errorf("Error reading to: %v", err)
		}
	}

	if err = d.Set("trans_to_type", flattenFirewallNatSnatTransToType(o["trans-to-type"], d, "trans_to_type", sv)); err != nil {
		if !fortiAPIPatch(o["trans-to-type"]) {
			return fmt.Errorf("Error reading trans-to-type: %v", err)
		}
	}

	if err = d.Set("trans_to_ip_start", flattenFirewallNatSnatTransToIpStart(o["trans-to-ip-start"], d, "trans_to_ip_start", sv)); err != nil {
		if !fortiAPIPatch(o["trans-to-ip-start"]) {
			return fmt.Errorf("Error reading trans-to-ip-start: %v", err)
		}
	}

	if err = d.Set("trans_to_ip", flattenFirewallNatSnatTransToIp(o["trans-to-ip"], d, "trans_to_ip", sv)); err != nil {
		if !fortiAPIPatch(o["trans-to-ip"]) {
			return fmt.Errorf("Error reading trans-to-ip: %v", err)
		}
	}

	if err = d.Set("mkey", flattenFirewallNatSnatMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandFirewallNatSnatStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatOutInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatFrom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatTrafficGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatTransToIpEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatTransToType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatTransToIpStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatTransToIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallNatSnatMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallNatSnat(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallNatSnatStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("out_interface"); ok {
		t, err := expandFirewallNatSnatOutInterface(d, v, "out_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-interface"] = t
		}
	}

	if v, ok := d.GetOk("from"); ok {
		t, err := expandFirewallNatSnatFrom(d, v, "from", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from"] = t
		}
	}

	if v, ok := d.GetOk("traffic_group"); ok {
		t, err := expandFirewallNatSnatTrafficGroup(d, v, "traffic_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-group"] = t
		}
	}

	if v, ok := d.GetOk("trans_to_ip_end"); ok {
		t, err := expandFirewallNatSnatTransToIpEnd(d, v, "trans_to_ip_end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trans-to-ip-end"] = t
		}
	}

	if v, ok := d.GetOk("to"); ok {
		t, err := expandFirewallNatSnatTo(d, v, "to", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["to"] = t
		}
	}

	if v, ok := d.GetOk("trans_to_type"); ok {
		t, err := expandFirewallNatSnatTransToType(d, v, "trans_to_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trans-to-type"] = t
		}
	}

	if v, ok := d.GetOk("trans_to_ip_start"); ok {
		t, err := expandFirewallNatSnatTransToIpStart(d, v, "trans_to_ip_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trans-to-ip-start"] = t
		}
	}

	if v, ok := d.GetOk("trans_to_ip"); ok {
		t, err := expandFirewallNatSnatTransToIp(d, v, "trans_to_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trans-to-ip"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandFirewallNatSnatMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
