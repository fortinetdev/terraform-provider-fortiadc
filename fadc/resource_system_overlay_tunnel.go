// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system overlay tunnel.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemOverlayTunnel() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemOverlayTunnelRead,
		Update: resourceSystemOverlayTunnelUpdate,
		Create: resourceSystemOverlayTunnelCreate,
		Delete: resourceSystemOverlayTunnelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dstip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mttl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipversion": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vsid": &schema.Schema{
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
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSystemOverlayTunnelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemOverlayTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemOverlayTunnel resource while getting object: %v", err)
	}

	_, err = c.CreateSystemOverlayTunnel(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemOverlayTunnel resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemOverlayTunnelRead(d, m)
}
func resourceSystemOverlayTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemOverlayTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemOverlayTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemOverlayTunnel(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemOverlayTunnel resource: %v", err)
	}

	return resourceSystemOverlayTunnelRead(d, m)
}
func resourceSystemOverlayTunnelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemOverlayTunnel(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemOverlayTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemOverlayTunnelRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemOverlayTunnel(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemOverlayTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemOverlayTunnel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemOverlayTunnel resource from API: %v", err)
	}
	return nil
}
func flattenSystemOverlayTunnelDstip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelTunnelType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelVni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelMttl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelIpversion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelVsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemOverlayTunnel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dstip", flattenSystemOverlayTunnelDstip(o["dstip"], d, "dstip", sv)); err != nil {
		if !fortiAPIPatch(o["dstip"]) {
			return fmt.Errorf("Error reading dstip: %v", err)
		}
	}

	if err = d.Set("tunnel_type", flattenSystemOverlayTunnelTunnelType(o["tunnel_type"], d, "tunnel_type", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel_type"]) {
			return fmt.Errorf("Error reading tunnel_type: %v", err)
		}
	}

	if err = d.Set("vni", flattenSystemOverlayTunnelVni(o["vni"], d, "vni", sv)); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("mttl", flattenSystemOverlayTunnelMttl(o["mttl"], d, "mttl", sv)); err != nil {
		if !fortiAPIPatch(o["mttl"]) {
			return fmt.Errorf("Error reading mttl: %v", err)
		}
	}

	if err = d.Set("ipversion", flattenSystemOverlayTunnelIpversion(o["ipversion"], d, "ipversion", sv)); err != nil {
		if !fortiAPIPatch(o["ipversion"]) {
			return fmt.Errorf("Error reading ipversion: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemOverlayTunnelInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vsid", flattenSystemOverlayTunnelVsid(o["vsid"], d, "vsid", sv)); err != nil {
		if !fortiAPIPatch(o["vsid"]) {
			return fmt.Errorf("Error reading vsid: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemOverlayTunnelPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemOverlayTunnelMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemOverlayTunnelDstip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelTunnelType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelVni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelMttl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelIpversion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelVsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemOverlayTunnel(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstip"); ok {
		t, err := expandSystemOverlayTunnelDstip(d, v, "dstip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstip"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_type"); ok {
		t, err := expandSystemOverlayTunnelTunnelType(d, v, "tunnel_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel_type"] = t
		}
	}

	if v, ok := d.GetOk("vni"); ok {
		t, err := expandSystemOverlayTunnelVni(d, v, "vni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	if v, ok := d.GetOk("mttl"); ok {
		t, err := expandSystemOverlayTunnelMttl(d, v, "mttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mttl"] = t
		}
	}

	if v, ok := d.GetOk("ipversion"); ok {
		t, err := expandSystemOverlayTunnelIpversion(d, v, "ipversion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipversion"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemOverlayTunnelInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("vsid"); ok {
		t, err := expandSystemOverlayTunnelVsid(d, v, "vsid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vsid"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemOverlayTunnelPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemOverlayTunnelMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
