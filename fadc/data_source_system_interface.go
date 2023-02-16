// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aggregate_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_gw_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"floating_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"floating": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundant_member": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_ip_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_gw_distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dedicate_to_management": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"floating_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aggregate_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemInterface(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterface: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterface from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAggregateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpGwOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFloatingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFloating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRedundantMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpIpOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpGwDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDefaultGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTrustIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceTrafficGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceDedicateToManagement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFloatingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAggregateAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemInterfaceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("aggregate_mode", dataSourceFlattenSystemInterfaceAggregateMode(o["aggregate-mode"], d, "aggregate_mode")); err != nil {
		if !fortiAPIPatch(o["aggregate-mode"]) {
			return fmt.Errorf("Error reading aggregate-mode: %v", err)
		}
	}

	if err = d.Set("dhcp_gw_override", dataSourceFlattenSystemInterfaceDhcpGwOverride(o["dhcp_gw_override"], d, "dhcp_gw_override")); err != nil {
		if !fortiAPIPatch(o["dhcp_gw_override"]) {
			return fmt.Errorf("Error reading dhcp_gw_override: %v", err)
		}
	}

	if err = d.Set("secondary_ip", dataSourceFlattenSystemInterfaceSecondaryIp(o["secondary-ip"], d, "secondary_ip")); err != nil {
		if !fortiAPIPatch(o["secondary-ip"]) {
			return fmt.Errorf("Error reading secondary-ip: %v", err)
		}
	}

	if err = d.Set("floating_ip", dataSourceFlattenSystemInterfaceFloatingIp(o["floating-ip"], d, "floating_ip")); err != nil {
		if !fortiAPIPatch(o["floating-ip"]) {
			return fmt.Errorf("Error reading floating-ip: %v", err)
		}
	}

	if err = d.Set("floating", dataSourceFlattenSystemInterfaceFloating(o["floating"], d, "floating")); err != nil {
		if !fortiAPIPatch(o["floating"]) {
			return fmt.Errorf("Error reading floating: %v", err)
		}
	}

	if err = d.Set("speed", dataSourceFlattenSystemInterfaceSpeed(o["speed"], d, "speed")); err != nil {
		if !fortiAPIPatch(o["speed"]) {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("redundant_member", dataSourceFlattenSystemInterfaceRedundantMember(o["redundant-member"], d, "redundant_member")); err != nil {
		if !fortiAPIPatch(o["redundant-member"]) {
			return fmt.Errorf("Error reading redundant-member: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_overlap", dataSourceFlattenSystemInterfaceDhcpIpOverlap(o["dhcp_ip_overlap"], d, "dhcp_ip_overlap")); err != nil {
		if !fortiAPIPatch(o["dhcp_ip_overlap"]) {
			return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
		}
	}

	if err = d.Set("dhcp_gw_distance", dataSourceFlattenSystemInterfaceDhcpGwDistance(o["dhcp_gw_distance"], d, "dhcp_gw_distance")); err != nil {
		if !fortiAPIPatch(o["dhcp_gw_distance"]) {
			return fmt.Errorf("Error reading dhcp_gw_distance: %v", err)
		}
	}

	if err = d.Set("wccp", dataSourceFlattenSystemInterfaceWccp(o["wccp"], d, "wccp")); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("default_gw", dataSourceFlattenSystemInterfaceDefaultGw(o["default-gw"], d, "default_gw")); err != nil {
		if !fortiAPIPatch(o["default-gw"]) {
			return fmt.Errorf("Error reading default-gw: %v", err)
		}
	}

	if err = d.Set("trust_ip", dataSourceFlattenSystemInterfaceTrustIp(o["trust-ip"], d, "trust_ip")); err != nil {
		if !fortiAPIPatch(o["trust-ip"]) {
			return fmt.Errorf("Error reading trust-ip: %v", err)
		}
	}

	if err = d.Set("vlanid", dataSourceFlattenSystemInterfaceVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemInterfaceVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("allowaccess", dataSourceFlattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("dns_server_override", dataSourceFlattenSystemInterfaceDnsServerOverride(o["dns-server-override"], d, "dns_server_override")); err != nil {
		if !fortiAPIPatch(o["dns-server-override"]) {
			return fmt.Errorf("Error reading dns-server-override: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemInterfaceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemInterfaceIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("traffic_group", dataSourceFlattenSystemInterfaceTrafficGroup(o["traffic-group"], d, "traffic_group")); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", dataSourceFlattenSystemInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["disc-retry-timeout"]) {
			return fmt.Errorf("Error reading disc-retry-timeout: %v", err)
		}
	}

	if err = d.Set("dhcp_gateway", dataSourceFlattenSystemInterfaceDhcpGateway(o["dhcp_gateway"], d, "dhcp_gateway")); err != nil {
		if !fortiAPIPatch(o["dhcp_gateway"]) {
			return fmt.Errorf("Error reading dhcp_gateway: %v", err)
		}
	}

	if err = d.Set("ip6", dataSourceFlattenSystemInterfaceIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("dedicate_to_management", dataSourceFlattenSystemInterfaceDedicateToManagement(o["dedicate-to-management"], d, "dedicate_to_management")); err != nil {
		if !fortiAPIPatch(o["dedicate-to-management"]) {
			return fmt.Errorf("Error reading dedicate-to-management: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemInterfaceInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenSystemInterfacePassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemInterfaceMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("mtu", dataSourceFlattenSystemInterfaceMtu(o["mtu"], d, "mtu")); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemInterfaceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("floating_ip6", dataSourceFlattenSystemInterfaceFloatingIp6(o["floating-ip6"], d, "floating_ip6")); err != nil {
		if !fortiAPIPatch(o["floating-ip6"]) {
			return fmt.Errorf("Error reading floating-ip6: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemInterfaceMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("aggregate_algorithm", dataSourceFlattenSystemInterfaceAggregateAlgorithm(o["aggregate-algorithm"], d, "aggregate_algorithm")); err != nil {
		if !fortiAPIPatch(o["aggregate-algorithm"]) {
			return fmt.Errorf("Error reading aggregate-algorithm: %v", err)
		}
	}

	return nil
}
