// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system overlay tunnel child remote host.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemOverlayTunnelChildRemoteHost() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemOverlayTunnelChildRemoteHostRead,
		Update: resourceSystemOverlayTunnelChildRemoteHostUpdate,
		Create: resourceSystemOverlayTunnelChildRemoteHostCreate,
		Delete: resourceSystemOverlayTunnelChildRemoteHostDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vtep": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceSystemOverlayTunnelChildRemoteHostCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemOverlayTunnelChildRemoteHost(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemOverlayTunnelChildRemoteHost resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemOverlayTunnelChildRemoteHost(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemOverlayTunnelChildRemoteHost resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemOverlayTunnelChildRemoteHostRead(d, m)
}
func resourceSystemOverlayTunnelChildRemoteHostUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemOverlayTunnelChildRemoteHost(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemOverlayTunnelChildRemoteHost resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemOverlayTunnelChildRemoteHost(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemOverlayTunnelChildRemoteHost resource: %v", err)
	}

	return resourceSystemOverlayTunnelChildRemoteHostRead(d, m)
}
func resourceSystemOverlayTunnelChildRemoteHostDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemOverlayTunnelChildRemoteHost(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemOverlayTunnelChildRemoteHost resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemOverlayTunnelChildRemoteHostRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemOverlayTunnelChildRemoteHost(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemOverlayTunnelChildRemoteHost resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemOverlayTunnelChildRemoteHost(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemOverlayTunnelChildRemoteHost resource from API: %v", err)
	}
	return nil
}
func flattenSystemOverlayTunnelChildRemoteHostMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelChildRemoteHostVtep(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemOverlayTunnelChildRemoteHostHostMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemOverlayTunnelChildRemoteHost(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemOverlayTunnelChildRemoteHostMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("vtep", flattenSystemOverlayTunnelChildRemoteHostVtep(o["vtep"], d, "vtep", sv)); err != nil {
		if !fortiAPIPatch(o["vtep"]) {
			return fmt.Errorf("Error reading vtep: %v", err)
		}
	}

	if err = d.Set("host_mac", flattenSystemOverlayTunnelChildRemoteHostHostMac(o["host_mac"], d, "host_mac", sv)); err != nil {
		if !fortiAPIPatch(o["host_mac"]) {
			return fmt.Errorf("Error reading host_mac: %v", err)
		}
	}

	return nil
}

func expandSystemOverlayTunnelChildRemoteHostMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelChildRemoteHostVtep(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemOverlayTunnelChildRemoteHostHostMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemOverlayTunnelChildRemoteHost(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemOverlayTunnelChildRemoteHostMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("vtep"); ok {
		t, err := expandSystemOverlayTunnelChildRemoteHostVtep(d, v, "vtep", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vtep"] = t
		}
	}

	if v, ok := d.GetOk("host_mac"); ok {
		t, err := expandSystemOverlayTunnelChildRemoteHostHostMac(d, v, "host_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_mac"] = t
		}
	}

	return &obj, nil
}
