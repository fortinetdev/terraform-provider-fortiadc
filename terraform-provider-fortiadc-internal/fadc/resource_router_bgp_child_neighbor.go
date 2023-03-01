// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router bgp child neighbor.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpChildNeighbor() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterBgpChildNeighborRead,
		Update: resourceRouterBgpChildNeighborUpdate,
		Create: resourceRouterBgpChildNeighborCreate,
		Delete: resourceRouterBgpChildNeighborDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distribute_list_out6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distribute_list_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix_list_in6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix_list_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix_list_out6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"prefix_list_in": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distribute_list_in": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distribute_list_in6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_as": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"holdtime": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keepalive": &schema.Schema{
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
func resourceRouterBgpChildNeighborCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterBgpChildNeighbor: type error")
	}

	obj, err := getObjectRouterBgpChildNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpChildNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpChildNeighbor(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpChildNeighbor resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterBgpChildNeighborRead(d, m)
}
func resourceRouterBgpChildNeighborUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterBgpChildNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpChildNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpChildNeighbor(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpChildNeighbor resource: %v", err)
	}

	return resourceRouterBgpChildNeighborRead(d, m)
}
func resourceRouterBgpChildNeighborDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterBgpChildNeighbor(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpChildNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterBgpChildNeighborRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterBgpChildNeighbor(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpChildNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpChildNeighbor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpChildNeighbor resource from API: %v", err)
	}
	return nil
}
func flattenRouterBgpChildNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborHoldtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpChildNeighborKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBgpChildNeighbor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip", flattenRouterBgpChildNeighborIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterBgpChildNeighborWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", flattenRouterBgpChildNeighborDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("Error reading distribute-list-out6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", flattenRouterBgpChildNeighborDistributeListOut(o["distribute-list-out"], d, "distribute_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("Error reading distribute-list-out: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", flattenRouterBgpChildNeighborPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("Error reading prefix-list-in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", flattenRouterBgpChildNeighborPrefixListOut(o["prefix-list-out"], d, "prefix_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("Error reading prefix-list-out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", flattenRouterBgpChildNeighborPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("Error reading prefix-list-out6: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", flattenRouterBgpChildNeighborPrefixListIn(o["prefix-list-in"], d, "prefix_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("Error reading prefix-list-in: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterBgpChildNeighborType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterBgpChildNeighborDistributeListIn(o["distribute-list-in"], d, "distribute_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute-list-in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", flattenRouterBgpChildNeighborDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("Error reading distribute-list-in6: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterBgpChildNeighborInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip6", flattenRouterBgpChildNeighborIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("remote_as", flattenRouterBgpChildNeighborRemoteAs(o["remote-as"], d, "remote_as", sv)); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("Error reading remote-as: %v", err)
		}
	}

	if err = d.Set("holdtime", flattenRouterBgpChildNeighborHoldtime(o["holdtime"], d, "holdtime", sv)); err != nil {
		if !fortiAPIPatch(o["holdtime"]) {
			return fmt.Errorf("Error reading holdtime: %v", err)
		}
	}

	if err = d.Set("port", flattenRouterBgpChildNeighborPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterBgpChildNeighborMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenRouterBgpChildNeighborKeepalive(o["keepalive"], d, "keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	return nil
}

func expandRouterBgpChildNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborHoldtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpChildNeighborKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpChildNeighbor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandRouterBgpChildNeighborIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandRouterBgpChildNeighborWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out6"); ok {
		t, err := expandRouterBgpChildNeighborDistributeListOut6(d, v, "distribute_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out"); ok {
		t, err := expandRouterBgpChildNeighborDistributeListOut(d, v, "distribute_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in6"); ok {
		t, err := expandRouterBgpChildNeighborPrefixListIn6(d, v, "prefix_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out"); ok {
		t, err := expandRouterBgpChildNeighborPrefixListOut(d, v, "prefix_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out6"); ok {
		t, err := expandRouterBgpChildNeighborPrefixListOut6(d, v, "prefix_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in"); ok {
		t, err := expandRouterBgpChildNeighborPrefixListIn(d, v, "prefix_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandRouterBgpChildNeighborType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok {
		t, err := expandRouterBgpChildNeighborDistributeListIn(d, v, "distribute_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in6"); ok {
		t, err := expandRouterBgpChildNeighborDistributeListIn6(d, v, "distribute_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterBgpChildNeighborInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandRouterBgpChildNeighborIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("remote_as"); ok {
		t, err := expandRouterBgpChildNeighborRemoteAs(d, v, "remote_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-as"] = t
		}
	}

	if v, ok := d.GetOk("holdtime"); ok {
		t, err := expandRouterBgpChildNeighborHoldtime(d, v, "holdtime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holdtime"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandRouterBgpChildNeighborPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterBgpChildNeighborMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok {
		t, err := expandRouterBgpChildNeighborKeepalive(d, v, "keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	return &obj, nil
}
