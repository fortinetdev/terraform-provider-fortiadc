// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system dns vdom.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDnsVdom() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemDnsVdomRead,
		Update: resourceSystemDnsVdomUpdate,
		Create: resourceSystemDnsVdomUpdate,
		Delete: resourceSystemDnsVdomDelete,

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
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_overide": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemDnsVdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemDnsVdom(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDnsVdom(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	d.SetId("SystemDns")
	return resourceSystemDnsRead(d, m)
}
func resourceSystemDnsVdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemDnsVdom(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDnsVdom(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemDnsVdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemDnsVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDnsVdom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}
func flattenSystemDnsVdomIpSecond(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsVdomIpPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsVdomDnsOveride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDnsVdom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip_second", flattenSystemDnsVdomIpSecond(o["ip_second"], d, "ip_second", sv)); err != nil {
		if !fortiAPIPatch(o["ip_second"]) {
			return fmt.Errorf("Error reading ip_second: %v", err)
		}
	}

	if err = d.Set("ip_primary", flattenSystemDnsVdomIpPrimary(o["ip_primary"], d, "ip_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip_primary"]) {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	if err = d.Set("dns_overide", flattenSystemDnsVdomDnsOveride(o["dns_overide"], d, "dns_overide", sv)); err != nil {
		if !fortiAPIPatch(o["dns_overide"]) {
			return fmt.Errorf("Error reading dns_overide: %v", err)
		}
	}

	return nil
}

func expandSystemDnsVdomIpSecond(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsVdomIpPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsVdomDnsOveride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDnsVdom(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip_second"); ok {
		if setArgNil {
			obj["ip_second"] = nil
		} else {
			t, err := expandSystemDnsVdomIpSecond(d, v, "ip_second", sv)
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
			t, err := expandSystemDnsVdomIpPrimary(d, v, "ip_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip_primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_overide"); ok {
		if setArgNil {
			obj["dns_overide"] = nil
		} else {
			t, err := expandSystemDnsVdomDnsOveride(d, v, "dns_overide", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns_overide"] = t
			}
		}
	}

	return &obj, nil
}
