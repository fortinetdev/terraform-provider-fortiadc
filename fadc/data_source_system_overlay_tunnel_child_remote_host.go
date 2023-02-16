// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system overlay tunnel child remote host.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemOverlayTunnelChildRemoteHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemOverlayTunnelChildRemoteHostRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vtep": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceSystemOverlayTunnelChildRemoteHostRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemOverlayTunnelChildRemoteHost: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemOverlayTunnelChildRemoteHost: type error")
	}

	o, err := c.ReadSystemOverlayTunnelChildRemoteHost(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemOverlayTunnelChildRemoteHost: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemOverlayTunnelChildRemoteHost(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemOverlayTunnelChildRemoteHost from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemOverlayTunnelChildRemoteHostMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelChildRemoteHostVtep(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemOverlayTunnelChildRemoteHostHostMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemOverlayTunnelChildRemoteHost(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenSystemOverlayTunnelChildRemoteHostMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("vtep", dataSourceFlattenSystemOverlayTunnelChildRemoteHostVtep(o["vtep"], d, "vtep")); err != nil {
		if !fortiAPIPatch(o["vtep"]) {
			return fmt.Errorf("Error reading vtep: %v", err)
		}
	}

	if err = d.Set("host_mac", dataSourceFlattenSystemOverlayTunnelChildRemoteHostHostMac(o["host_mac"], d, "host_mac")); err != nil {
		if !fortiAPIPatch(o["host_mac"]) {
			return fmt.Errorf("Error reading host_mac: %v", err)
		}
	}

	return nil
}
