// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system dns.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Create: resourceSystemDnsUpdate,
		Delete: resourceSystemDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip_second": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemDns(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	d.SetId("SystemDns")
	return resourceSystemDnsRead(d, m)
}
func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemDns(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemDns(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}
func flattenSystemDnsIpSecond(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsIpPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip_second", flattenSystemDnsIpSecond(o["ip_second"], d, "ip_second", sv)); err != nil {
		if !fortiAPIPatch(o["ip_second"]) {
			return fmt.Errorf("Error reading ip_second: %v", err)
		}
	}

	if err = d.Set("ip_primary", flattenSystemDnsIpPrimary(o["ip_primary"], d, "ip_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip_primary"]) {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	return nil
}

func expandSystemDnsIpSecond(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIpPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip_second"); ok {
		if setArgNil {
			obj["ip_second"] = nil
		} else {
			t, err := expandSystemDnsIpSecond(d, v, "ip_second", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip_second"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_primary"); ok {
		if setArgNil {
			obj["ip_primary"] = nil
		} else {
			t, err := expandSystemDnsIpPrimary(d, v, "ip_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip_primary"] = t
			}
		}
	}

	return &obj, nil
}
