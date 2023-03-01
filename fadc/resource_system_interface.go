// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemInterfaceRead,
		Update: resourceSystemInterfaceUpdate,
		Create: resourceSystemInterfaceCreate,
		Delete: resourceSystemInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"aggregate_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dhcp_gw_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"floating_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"floating": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_member": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dhcp_ip_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dhcp_gw_distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trust_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dhcp_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dedicate_to_management": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"floating_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"aggregate_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterface: type error")
	}

	obj, err := getObjectSystemInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterface(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemInterfaceRead(d, m)
}
func resourceSystemInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterface(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource: %v", err)
	}

	return resourceSystemInterfaceRead(d, m)
}
func resourceSystemInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemInterface(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemInterface(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource from API: %v", err)
	}
	return nil
}
func flattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceAggregateMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpGwOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceFloatingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceFloating(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceRedundantMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpIpOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpGwDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceWccp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDefaultGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceTrustIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDnsServerOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceTrafficGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceDedicateToManagement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfacePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceFloatingIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceAggregateAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemInterfaceStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("aggregate_mode", flattenSystemInterfaceAggregateMode(o["aggregate-mode"], d, "aggregate_mode", sv)); err != nil {
		if !fortiAPIPatch(o["aggregate-mode"]) {
			return fmt.Errorf("Error reading aggregate-mode: %v", err)
		}
	}

	if err = d.Set("dhcp_gw_override", flattenSystemInterfaceDhcpGwOverride(o["dhcp_gw_override"], d, "dhcp_gw_override", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp_gw_override"]) {
			return fmt.Errorf("Error reading dhcp_gw_override: %v", err)
		}
	}

	if err = d.Set("secondary_ip", flattenSystemInterfaceSecondaryIp(o["secondary-ip"], d, "secondary_ip", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-ip"]) {
			return fmt.Errorf("Error reading secondary-ip: %v", err)
		}
	}

	if err = d.Set("floating_ip", flattenSystemInterfaceFloatingIp(o["floating-ip"], d, "floating_ip", sv)); err != nil {
		if !fortiAPIPatch(o["floating-ip"]) {
			return fmt.Errorf("Error reading floating-ip: %v", err)
		}
	}

	if err = d.Set("floating", flattenSystemInterfaceFloating(o["floating"], d, "floating", sv)); err != nil {
		if !fortiAPIPatch(o["floating"]) {
			return fmt.Errorf("Error reading floating: %v", err)
		}
	}

	if err = d.Set("speed", flattenSystemInterfaceSpeed(o["speed"], d, "speed", sv)); err != nil {
		if !fortiAPIPatch(o["speed"]) {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("redundant_member", flattenSystemInterfaceRedundantMember(o["redundant-member"], d, "redundant_member", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-member"]) {
			return fmt.Errorf("Error reading redundant-member: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_overlap", flattenSystemInterfaceDhcpIpOverlap(o["dhcp_ip_overlap"], d, "dhcp_ip_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp_ip_overlap"]) {
			return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
		}
	}

	if err = d.Set("dhcp_gw_distance", flattenSystemInterfaceDhcpGwDistance(o["dhcp_gw_distance"], d, "dhcp_gw_distance", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp_gw_distance"]) {
			return fmt.Errorf("Error reading dhcp_gw_distance: %v", err)
		}
	}

	if err = d.Set("wccp", flattenSystemInterfaceWccp(o["wccp"], d, "wccp", sv)); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("default_gw", flattenSystemInterfaceDefaultGw(o["default-gw"], d, "default_gw", sv)); err != nil {
		if !fortiAPIPatch(o["default-gw"]) {
			return fmt.Errorf("Error reading default-gw: %v", err)
		}
	}

	if err = d.Set("trust_ip", flattenSystemInterfaceTrustIp(o["trust-ip"], d, "trust_ip", sv)); err != nil {
		if !fortiAPIPatch(o["trust-ip"]) {
			return fmt.Errorf("Error reading trust-ip: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenSystemInterfaceVlanid(o["vlanid"], d, "vlanid", sv)); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemInterfaceVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("dns_server_override", flattenSystemInterfaceDnsServerOverride(o["dns-server-override"], d, "dns_server_override", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server-override"]) {
			return fmt.Errorf("Error reading dns-server-override: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemInterfaceType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("traffic_group", flattenSystemInterfaceTrafficGroup(o["traffic-group"], d, "traffic_group", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", flattenSystemInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["disc-retry-timeout"]) {
			return fmt.Errorf("Error reading disc-retry-timeout: %v", err)
		}
	}

	if err = d.Set("dhcp_gateway", flattenSystemInterfaceDhcpGateway(o["dhcp_gateway"], d, "dhcp_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp_gateway"]) {
			return fmt.Errorf("Error reading dhcp_gateway: %v", err)
		}
	}

	if err = d.Set("ip6", flattenSystemInterfaceIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("dedicate_to_management", flattenSystemInterfaceDedicateToManagement(o["dedicate-to-management"], d, "dedicate_to_management", sv)); err != nil {
		if !fortiAPIPatch(o["dedicate-to-management"]) {
			return fmt.Errorf("Error reading dedicate-to-management: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemInterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemInterfacePassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemInterfaceMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("mtu", flattenSystemInterfaceMtu(o["mtu"], d, "mtu", sv)); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemInterfaceUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("floating_ip6", flattenSystemInterfaceFloatingIp6(o["floating-ip6"], d, "floating_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["floating-ip6"]) {
			return fmt.Errorf("Error reading floating-ip6: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemInterfaceMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("aggregate_algorithm", flattenSystemInterfaceAggregateAlgorithm(o["aggregate-algorithm"], d, "aggregate_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["aggregate-algorithm"]) {
			return fmt.Errorf("Error reading aggregate-algorithm: %v", err)
		}
	}

	return nil
}

func expandSystemInterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAggregateMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpGwOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFloatingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFloating(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceRedundantMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpIpOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpGwDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceWccp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDefaultGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDnsServerOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrafficGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDiscRetryTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDedicateToManagement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFloatingIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAggregateAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemInterfaceStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_mode"); ok {
		t, err := expandSystemInterfaceAggregateMode(d, v, "aggregate_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-mode"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_gw_override"); ok {
		t, err := expandSystemInterfaceDhcpGwOverride(d, v, "dhcp_gw_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp_gw_override"] = t
		}
	}

	if v, ok := d.GetOk("secondary_ip"); ok {
		t, err := expandSystemInterfaceSecondaryIp(d, v, "secondary_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-ip"] = t
		}
	}

	if v, ok := d.GetOk("floating_ip"); ok {
		t, err := expandSystemInterfaceFloatingIp(d, v, "floating_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["floating-ip"] = t
		}
	}

	if v, ok := d.GetOk("floating"); ok {
		t, err := expandSystemInterfaceFloating(d, v, "floating", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["floating"] = t
		}
	}

	if v, ok := d.GetOk("speed"); ok {
		t, err := expandSystemInterfaceSpeed(d, v, "speed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speed"] = t
		}
	}

	if v, ok := d.GetOk("redundant_member"); ok {
		t, err := expandSystemInterfaceRedundantMember(d, v, "redundant_member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-member"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_overlap"); ok {
		t, err := expandSystemInterfaceDhcpIpOverlap(d, v, "dhcp_ip_overlap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp_ip_overlap"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_gw_distance"); ok {
		t, err := expandSystemInterfaceDhcpGwDistance(d, v, "dhcp_gw_distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp_gw_distance"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok {
		t, err := expandSystemInterfaceWccp(d, v, "wccp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("default_gw"); ok {
		t, err := expandSystemInterfaceDefaultGw(d, v, "default_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gw"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip"); ok {
		t, err := expandSystemInterfaceTrustIp(d, v, "trust_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok {
		t, err := expandSystemInterfaceVlanid(d, v, "vlanid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemInterfaceVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandSystemInterfaceAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("dns_server_override"); ok {
		t, err := expandSystemInterfaceDnsServerOverride(d, v, "dns_server_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server-override"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemInterfaceType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemInterfaceIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("traffic_group"); ok {
		t, err := expandSystemInterfaceTrafficGroup(d, v, "traffic_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-group"] = t
		}
	}

	if v, ok := d.GetOk("disc_retry_timeout"); ok {
		t, err := expandSystemInterfaceDiscRetryTimeout(d, v, "disc_retry_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disc-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_gateway"); ok {
		t, err := expandSystemInterfaceDhcpGateway(d, v, "dhcp_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp_gateway"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandSystemInterfaceIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("dedicate_to_management"); ok {
		t, err := expandSystemInterfaceDedicateToManagement(d, v, "dedicate_to_management", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicate-to-management"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemInterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemInterfacePassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemInterfaceMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok {
		t, err := expandSystemInterfaceMtu(d, v, "mtu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemInterfaceUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("floating_ip6"); ok {
		t, err := expandSystemInterfaceFloatingIp6(d, v, "floating_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["floating-ip6"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemInterfaceMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_algorithm"); ok {
		t, err := expandSystemInterfaceAggregateAlgorithm(d, v, "aggregate_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-algorithm"] = t
		}
	}

	return &obj, nil
}
