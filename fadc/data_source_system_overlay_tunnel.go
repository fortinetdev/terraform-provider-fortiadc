// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system overlay tunnel.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemOverlayTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemOverlayTunnelRead,
		Schema: map[string]*schema.Schema{
			"dstip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mttl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipversion": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceSystemOverlayTunnelRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemOverlayTunnel: type error")
	}

	o, err := c.ReadSystemOverlayTunnel(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemOverlayTunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemOverlayTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemOverlayTunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemOverlayTunnelDstip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelTunnelType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelMttl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelIpversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelVsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemOverlayTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dstip", dataSourceFlattenSystemOverlayTunnelDstip(o["dstip"], d, "dstip")); err != nil {
		if !fortiAPIPatch(o["dstip"]) {
			return fmt.Errorf("Error reading dstip: %v", err)
		}
	}

	if err = d.Set("tunnel_type", dataSourceFlattenSystemOverlayTunnelTunnelType(o["tunnel_type"], d, "tunnel_type")); err != nil {
		if !fortiAPIPatch(o["tunnel_type"]) {
			return fmt.Errorf("Error reading tunnel_type: %v", err)
		}
	}

	if err = d.Set("vni", dataSourceFlattenSystemOverlayTunnelVni(o["vni"], d, "vni")); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("mttl", dataSourceFlattenSystemOverlayTunnelMttl(o["mttl"], d, "mttl")); err != nil {
		if !fortiAPIPatch(o["mttl"]) {
			return fmt.Errorf("Error reading mttl: %v", err)
		}
	}

	if err = d.Set("ipversion", dataSourceFlattenSystemOverlayTunnelIpversion(o["ipversion"], d, "ipversion")); err != nil {
		if !fortiAPIPatch(o["ipversion"]) {
			return fmt.Errorf("Error reading ipversion: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemOverlayTunnelInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vsid", dataSourceFlattenSystemOverlayTunnelVsid(o["vsid"], d, "vsid")); err != nil {
		if !fortiAPIPatch(o["vsid"]) {
			return fmt.Errorf("Error reading vsid: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemOverlayTunnelPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemOverlayTunnelMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
