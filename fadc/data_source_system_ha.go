// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system ha.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemHaRead,
		Schema: map[string]*schema.Schema{
			"mgmt_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_node_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_l4_persistent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitor_enable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hbdev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_ip_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"groupid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_l4_connection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"datadev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arp_num": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failover_hold_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arp_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_l7_persistent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hbtype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failover_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "SystemHa"

	o, err := c.ReadSystemHa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemHa: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemHa(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemHa from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemHaMgmtInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaLocalNodeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPeerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSyncL4Persistent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMonitorEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMgmtIpAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaConfigPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaGroupid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaNodeList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaLocalAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMgmtIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemHaSyncL4Connection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaDatadev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaArpNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaFailoverHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaArpInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMgmtStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaInterfaceList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSyncL7Persistent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mgmt_interface", dataSourceFlattenSystemHaMgmtInterface(o["mgmt-interface"], d, "mgmt_interface")); err != nil {
		if !fortiAPIPatch(o["mgmt-interface"]) {
			return fmt.Errorf("Error reading mgmt-interface: %v", err)
		}
	}

	if err = d.Set("local_node_id", dataSourceFlattenSystemHaLocalNodeId(o["local-node-id"], d, "local_node_id")); err != nil {
		if !fortiAPIPatch(o["local-node-id"]) {
			return fmt.Errorf("Error reading local-node-id: %v", err)
		}
	}

	if err = d.Set("peer_address", dataSourceFlattenSystemHaPeerAddress(o["peer-address"], d, "peer_address")); err != nil {
		if !fortiAPIPatch(o["peer-address"]) {
			return fmt.Errorf("Error reading peer-address: %v", err)
		}
	}

	if err = d.Set("group_name", dataSourceFlattenSystemHaGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group-name: %v", err)
		}
	}

	if err = d.Set("sync_l4_persistent", dataSourceFlattenSystemHaSyncL4Persistent(o["sync-l4-persistent"], d, "sync_l4_persistent")); err != nil {
		if !fortiAPIPatch(o["sync-l4-persistent"]) {
			return fmt.Errorf("Error reading sync-l4-persistent: %v", err)
		}
	}

	if err = d.Set("monitor_enable", dataSourceFlattenSystemHaMonitorEnable(o["monitor-enable"], d, "monitor_enable")); err != nil {
		if !fortiAPIPatch(o["monitor-enable"]) {
			return fmt.Errorf("Error reading monitor-enable: %v", err)
		}
	}

	if err = d.Set("hbdev", dataSourceFlattenSystemHaHbdev(o["hbdev"], d, "hbdev")); err != nil {
		if !fortiAPIPatch(o["hbdev"]) {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}

	if err = d.Set("mgmt_ip_allowaccess", dataSourceFlattenSystemHaMgmtIpAllowaccess(o["mgmt-ip-allowaccess"], d, "mgmt_ip_allowaccess")); err != nil {
		if !fortiAPIPatch(o["mgmt-ip-allowaccess"]) {
			return fmt.Errorf("Error reading mgmt-ip-allowaccess: %v", err)
		}
	}

	if err = d.Set("config_priority", dataSourceFlattenSystemHaConfigPriority(o["config-priority"], d, "config_priority")); err != nil {
		if !fortiAPIPatch(o["config-priority"]) {
			return fmt.Errorf("Error reading config-priority: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", dataSourceFlattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb-lost-threshold: %v", err)
		}
	}

	if err = d.Set("override", dataSourceFlattenSystemHaOverride(o["override"], d, "override")); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("groupid", dataSourceFlattenSystemHaGroupid(o["groupid"], d, "groupid")); err != nil {
		if !fortiAPIPatch(o["groupid"]) {
			return fmt.Errorf("Error reading groupid: %v", err)
		}
	}

	if err = d.Set("node_list", dataSourceFlattenSystemHaNodeList(o["node-list"], d, "node_list")); err != nil {
		if !fortiAPIPatch(o["node-list"]) {
			return fmt.Errorf("Error reading node-list: %v", err)
		}
	}

	if err = d.Set("local_address", dataSourceFlattenSystemHaLocalAddress(o["local-address"], d, "local_address")); err != nil {
		if !fortiAPIPatch(o["local-address"]) {
			return fmt.Errorf("Error reading local-address: %v", err)
		}
	}

	if err = d.Set("mgmt_ip", dataSourceFlattenSystemHaMgmtIp(o["mgmt-ip"], d, "mgmt_ip")); err != nil {
		if !fortiAPIPatch(o["mgmt-ip"]) {
			return fmt.Errorf("Error reading mgmt-ip: %v", err)
		}
	}

	if err = d.Set("sync_l4_connection", dataSourceFlattenSystemHaSyncL4Connection(o["sync-l4-connection"], d, "sync_l4_connection")); err != nil {
		if !fortiAPIPatch(o["sync-l4-connection"]) {
			return fmt.Errorf("Error reading sync-l4-connection: %v", err)
		}
	}

	if err = d.Set("datadev", dataSourceFlattenSystemHaDatadev(o["datadev"], d, "datadev")); err != nil {
		if !fortiAPIPatch(o["datadev"]) {
			return fmt.Errorf("Error reading datadev: %v", err)
		}
	}

	if err = d.Set("arp_num", dataSourceFlattenSystemHaArpNum(o["arp_num"], d, "arp_num")); err != nil {
		if !fortiAPIPatch(o["arp_num"]) {
			return fmt.Errorf("Error reading arp_num: %v", err)
		}
	}

	if err = d.Set("failover_hold_time", dataSourceFlattenSystemHaFailoverHoldTime(o["failover-hold-time"], d, "failover_hold_time")); err != nil {
		if !fortiAPIPatch(o["failover-hold-time"]) {
			return fmt.Errorf("Error reading failover-hold-time: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemHaInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("arp_interval", dataSourceFlattenSystemHaArpInterval(o["arp_interval"], d, "arp_interval")); err != nil {
		if !fortiAPIPatch(o["arp_interval"]) {
			return fmt.Errorf("Error reading arp_interval: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemHaMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("mgmt_status", dataSourceFlattenSystemHaMgmtStatus(o["mgmt-status"], d, "mgmt_status")); err != nil {
		if !fortiAPIPatch(o["mgmt-status"]) {
			return fmt.Errorf("Error reading mgmt-status: %v", err)
		}
	}

	if err = d.Set("interface_list", dataSourceFlattenSystemHaInterfaceList(o["interface_list"], d, "interface_list")); err != nil {
		if !fortiAPIPatch(o["interface_list"]) {
			return fmt.Errorf("Error reading interface_list: %v", err)
		}
	}

	if err = d.Set("sync_l7_persistent", dataSourceFlattenSystemHaSyncL7Persistent(o["sync-l7-persistent"], d, "sync_l7_persistent")); err != nil {
		if !fortiAPIPatch(o["sync-l7-persistent"]) {
			return fmt.Errorf("Error reading sync-l7-persistent: %v", err)
		}
	}

	if err = d.Set("hbtype", dataSourceFlattenSystemHaHbtype(o["hbtype"], d, "hbtype")); err != nil {
		if !fortiAPIPatch(o["hbtype"]) {
			return fmt.Errorf("Error reading hbtype: %v", err)
		}
	}

	if err = d.Set("failover_threshold", dataSourceFlattenSystemHaFailoverThreshold(o["failover-threshold"], d, "failover_threshold")); err != nil {
		if !fortiAPIPatch(o["failover-threshold"]) {
			return fmt.Errorf("Error reading failover-threshold: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemHaPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}
