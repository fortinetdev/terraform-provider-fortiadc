// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system ha.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemHaRead,
		Update: resourceSystemHaUpdate,
		Create: resourceSystemHaUpdate,
		Delete: resourceSystemHaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mgmt_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_node_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"peer_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sync_l4_persistent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"monitor_enable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hbdev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mgmt_ip_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"config_priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"groupid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"node_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mgmt_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sync_l4_connection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"datadev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"arp_num": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"failover_hold_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"arp_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sync_l7_persistent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hbtype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"failover_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHa(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	d.SetId("SystemHa")
	return resourceSystemHaRead(d, m)
}
func resourceSystemHaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHa(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemHa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}
	return nil
}
func flattenSystemHaMgmtInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaLocalNodeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPeerAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSyncL4Persistent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMonitorEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHbdev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMgmtIpAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaConfigPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaGroupid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaNodeList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaLocalAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMgmtIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemHaSyncL4Connection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaDatadev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaArpNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaFailoverHoldTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaArpInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMgmtStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaInterfaceList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSyncL7Persistent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHbtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaFailoverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mgmt_interface", flattenSystemHaMgmtInterface(o["mgmt-interface"], d, "mgmt_interface", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-interface"]) {
			return fmt.Errorf("Error reading mgmt-interface: %v", err)
		}
	}

	if err = d.Set("local_node_id", flattenSystemHaLocalNodeId(o["local-node-id"], d, "local_node_id", sv)); err != nil {
		if !fortiAPIPatch(o["local-node-id"]) {
			return fmt.Errorf("Error reading local-node-id: %v", err)
		}
	}

	if err = d.Set("peer_address", flattenSystemHaPeerAddress(o["peer-address"], d, "peer_address", sv)); err != nil {
		if !fortiAPIPatch(o["peer-address"]) {
			return fmt.Errorf("Error reading peer-address: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemHaGroupName(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group-name: %v", err)
		}
	}

	if err = d.Set("sync_l4_persistent", flattenSystemHaSyncL4Persistent(o["sync-l4-persistent"], d, "sync_l4_persistent", sv)); err != nil {
		if !fortiAPIPatch(o["sync-l4-persistent"]) {
			return fmt.Errorf("Error reading sync-l4-persistent: %v", err)
		}
	}

	if err = d.Set("monitor_enable", flattenSystemHaMonitorEnable(o["monitor-enable"], d, "monitor_enable", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-enable"]) {
			return fmt.Errorf("Error reading monitor-enable: %v", err)
		}
	}

	if err = d.Set("hbdev", flattenSystemHaHbdev(o["hbdev"], d, "hbdev", sv)); err != nil {
		if !fortiAPIPatch(o["hbdev"]) {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}

	if err = d.Set("mgmt_ip_allowaccess", flattenSystemHaMgmtIpAllowaccess(o["mgmt-ip-allowaccess"], d, "mgmt_ip_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-ip-allowaccess"]) {
			return fmt.Errorf("Error reading mgmt-ip-allowaccess: %v", err)
		}
	}

	if err = d.Set("config_priority", flattenSystemHaConfigPriority(o["config-priority"], d, "config_priority", sv)); err != nil {
		if !fortiAPIPatch(o["config-priority"]) {
			return fmt.Errorf("Error reading config-priority: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb-lost-threshold: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemHaOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("groupid", flattenSystemHaGroupid(o["groupid"], d, "groupid", sv)); err != nil {
		if !fortiAPIPatch(o["groupid"]) {
			return fmt.Errorf("Error reading groupid: %v", err)
		}
	}

	if err = d.Set("node_list", flattenSystemHaNodeList(o["node-list"], d, "node_list", sv)); err != nil {
		if !fortiAPIPatch(o["node-list"]) {
			return fmt.Errorf("Error reading node-list: %v", err)
		}
	}

	if err = d.Set("local_address", flattenSystemHaLocalAddress(o["local-address"], d, "local_address", sv)); err != nil {
		if !fortiAPIPatch(o["local-address"]) {
			return fmt.Errorf("Error reading local-address: %v", err)
		}
	}

	if err = d.Set("mgmt_ip", flattenSystemHaMgmtIp(o["mgmt-ip"], d, "mgmt_ip", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-ip"]) {
			return fmt.Errorf("Error reading mgmt-ip: %v", err)
		}
	}

	if err = d.Set("sync_l4_connection", flattenSystemHaSyncL4Connection(o["sync-l4-connection"], d, "sync_l4_connection", sv)); err != nil {
		if !fortiAPIPatch(o["sync-l4-connection"]) {
			return fmt.Errorf("Error reading sync-l4-connection: %v", err)
		}
	}

	if err = d.Set("datadev", flattenSystemHaDatadev(o["datadev"], d, "datadev", sv)); err != nil {
		if !fortiAPIPatch(o["datadev"]) {
			return fmt.Errorf("Error reading datadev: %v", err)
		}
	}

	if err = d.Set("arp_num", flattenSystemHaArpNum(o["arp_num"], d, "arp_num", sv)); err != nil {
		if !fortiAPIPatch(o["arp_num"]) {
			return fmt.Errorf("Error reading arp_num: %v", err)
		}
	}

	if err = d.Set("failover_hold_time", flattenSystemHaFailoverHoldTime(o["failover-hold-time"], d, "failover_hold_time", sv)); err != nil {
		if !fortiAPIPatch(o["failover-hold-time"]) {
			return fmt.Errorf("Error reading failover-hold-time: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemHaInterval(o["interval"], d, "interval", sv)); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("arp_interval", flattenSystemHaArpInterval(o["arp_interval"], d, "arp_interval", sv)); err != nil {
		if !fortiAPIPatch(o["arp_interval"]) {
			return fmt.Errorf("Error reading arp_interval: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("mgmt_status", flattenSystemHaMgmtStatus(o["mgmt-status"], d, "mgmt_status", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-status"]) {
			return fmt.Errorf("Error reading mgmt-status: %v", err)
		}
	}

	if err = d.Set("interface_list", flattenSystemHaInterfaceList(o["interface_list"], d, "interface_list", sv)); err != nil {
		if !fortiAPIPatch(o["interface_list"]) {
			return fmt.Errorf("Error reading interface_list: %v", err)
		}
	}

	if err = d.Set("sync_l7_persistent", flattenSystemHaSyncL7Persistent(o["sync-l7-persistent"], d, "sync_l7_persistent", sv)); err != nil {
		if !fortiAPIPatch(o["sync-l7-persistent"]) {
			return fmt.Errorf("Error reading sync-l7-persistent: %v", err)
		}
	}

	if err = d.Set("hbtype", flattenSystemHaHbtype(o["hbtype"], d, "hbtype", sv)); err != nil {
		if !fortiAPIPatch(o["hbtype"]) {
			return fmt.Errorf("Error reading hbtype: %v", err)
		}
	}

	if err = d.Set("failover_threshold", flattenSystemHaFailoverThreshold(o["failover-threshold"], d, "failover_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["failover-threshold"]) {
			return fmt.Errorf("Error reading failover-threshold: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemHaPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func expandSystemHaMgmtInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLocalNodeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncL4Persistent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbdev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMgmtIpAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaConfigPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbLostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaNodeList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLocalAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMgmtIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncL4Connection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaDatadev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArpNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFailoverHoldTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArpInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMgmtStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaInterfaceList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncL7Persistent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFailoverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mgmt_interface"); ok {
		if setArgNil {
			obj["mgmt_interface"] = nil
		} else {
			t, err := expandSystemHaMgmtInterface(d, v, "mgmt_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_node_id"); ok {
		if setArgNil {
			obj["local_node_id"] = nil
		} else {
			t, err := expandSystemHaLocalNodeId(d, v, "local_node_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-node-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("peer_address"); ok {
		if setArgNil {
			obj["peer_address"] = nil
		} else {
			t, err := expandSystemHaPeerAddress(d, v, "peer_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["peer-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("group_name"); ok {
		if setArgNil {
			obj["group_name"] = nil
		} else {
			t, err := expandSystemHaGroupName(d, v, "group_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["group-name"] = t
			}
		}
	}

	if v, ok := d.GetOk("sync_l4_persistent"); ok {
		if setArgNil {
			obj["sync_l4_persistent"] = nil
		} else {
			t, err := expandSystemHaSyncL4Persistent(d, v, "sync_l4_persistent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sync-l4-persistent"] = t
			}
		}
	}

	if v, ok := d.GetOk("monitor_enable"); ok {
		if setArgNil {
			obj["monitor_enable"] = nil
		} else {
			t, err := expandSystemHaMonitorEnable(d, v, "monitor_enable", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["monitor-enable"] = t
			}
		}
	}

	if v, ok := d.GetOk("hbdev"); ok {
		if setArgNil {
			obj["hbdev"] = nil
		} else {
			t, err := expandSystemHaHbdev(d, v, "hbdev", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hbdev"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_ip_allowaccess"); ok {
		if setArgNil {
			obj["mgmt_ip_allowaccess"] = nil
		} else {
			t, err := expandSystemHaMgmtIpAllowaccess(d, v, "mgmt_ip_allowaccess", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-ip-allowaccess"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_priority"); ok {
		if setArgNil {
			obj["config_priority"] = nil
		} else {
			t, err := expandSystemHaConfigPriority(d, v, "config_priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-priority"] = t
			}
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		if setArgNil {
			obj["hb_lost_threshold"] = nil
		} else {
			t, err := expandSystemHaHbLostThreshold(d, v, "hb_lost_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hb-lost-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {
			t, err := expandSystemHaOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("groupid"); ok {
		if setArgNil {
			obj["groupid"] = nil
		} else {
			t, err := expandSystemHaGroupid(d, v, "groupid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["groupid"] = t
			}
		}
	}

	if v, ok := d.GetOk("node_list"); ok {
		if setArgNil {
			obj["node_list"] = nil
		} else {
			t, err := expandSystemHaNodeList(d, v, "node_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["node-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_address"); ok {
		if setArgNil {
			obj["local_address"] = nil
		} else {
			t, err := expandSystemHaLocalAddress(d, v, "local_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_ip"); ok {
		if setArgNil {
			obj["mgmt_ip"] = nil
		} else {
			t, err := expandSystemHaMgmtIp(d, v, "mgmt_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("sync_l4_connection"); ok {
		if setArgNil {
			obj["sync_l4_connection"] = nil
		} else {
			t, err := expandSystemHaSyncL4Connection(d, v, "sync_l4_connection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sync-l4-connection"] = t
			}
		}
	}

	if v, ok := d.GetOk("datadev"); ok {
		if setArgNil {
			obj["datadev"] = nil
		} else {
			t, err := expandSystemHaDatadev(d, v, "datadev", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["datadev"] = t
			}
		}
	}

	if v, ok := d.GetOk("arp_num"); ok {
		if setArgNil {
			obj["arp_num"] = nil
		} else {
			t, err := expandSystemHaArpNum(d, v, "arp_num", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["arp_num"] = t
			}
		}
	}

	if v, ok := d.GetOk("failover_hold_time"); ok {
		if setArgNil {
			obj["failover_hold_time"] = nil
		} else {
			t, err := expandSystemHaFailoverHoldTime(d, v, "failover_hold_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failover-hold-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		if setArgNil {
			obj["interval"] = nil
		} else {
			t, err := expandSystemHaInterval(d, v, "interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("arp_interval"); ok {
		if setArgNil {
			obj["arp_interval"] = nil
		} else {
			t, err := expandSystemHaArpInterval(d, v, "arp_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["arp_interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {
			t, err := expandSystemHaMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_status"); ok {
		if setArgNil {
			obj["mgmt_status"] = nil
		} else {
			t, err := expandSystemHaMgmtStatus(d, v, "mgmt_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_list"); ok {
		if setArgNil {
			obj["interface_list"] = nil
		} else {
			t, err := expandSystemHaInterfaceList(d, v, "interface_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface_list"] = t
			}
		}
	}

	if v, ok := d.GetOk("sync_l7_persistent"); ok {
		if setArgNil {
			obj["sync_l7_persistent"] = nil
		} else {
			t, err := expandSystemHaSyncL7Persistent(d, v, "sync_l7_persistent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sync-l7-persistent"] = t
			}
		}
	}

	if v, ok := d.GetOk("hbtype"); ok {
		if setArgNil {
			obj["hbtype"] = nil
		} else {
			t, err := expandSystemHaHbtype(d, v, "hbtype", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hbtype"] = t
			}
		}
	}

	if v, ok := d.GetOk("failover_threshold"); ok {
		if setArgNil {
			obj["failover_threshold"] = nil
		} else {
			t, err := expandSystemHaFailoverThreshold(d, v, "failover_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failover-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		if setArgNil {
			obj["priority"] = nil
		} else {
			t, err := expandSystemHaPriority(d, v, "priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["priority"] = t
			}
		}
	}

	return &obj, nil
}
