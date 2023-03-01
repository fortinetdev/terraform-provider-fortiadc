// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHaMgmtUpdate() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemHaMgmtUpdateRead,
		Update: resourceSystemHaMgmtUpdateUpdate,
		Create: resourceSystemHaMgmtUpdateUpdate,
		Delete: resourceSystemHaMgmtUpdateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mgmt_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mgmt_ip_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemHaMgmtUpdateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHaMgmtUpdate(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMgmtUpdate resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaMgmtUpdate(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMgmtUpdate resource: %v", err)
	}

	d.SetId("SystemHaMgmtUpdate")
	return resourceSystemHaMgmtUpdateRead(d, m)
}
func flattenSystemHaMgmtUpdateInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMgmtUpdateIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemHaMgmtUpdateStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMgmtUpdateIpAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHaMgmtUpdate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mgmt_interface", flattenSystemHaMgmtUpdateInterface(o["mgmt-interface"], d, "mgmt_interface", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-interface"]) {
			return fmt.Errorf("Error reading mgmt-interface: %v", err)
		}
	}

	if err = d.Set("mgmt_ip_allowaccess", flattenSystemHaMgmtUpdateIpAllowaccess(o["mgmt-ip-allowaccess"], d, "mgmt_ip_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-ip-allowaccess"]) {
			return fmt.Errorf("Error reading mgmt-ip-allowaccess: %v", err)
		}
	}

	if err = d.Set("mgmt_ip", flattenSystemHaMgmtUpdateIp(o["mgmt-ip"], d, "mgmt_ip", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-ip"]) {
			return fmt.Errorf("Error reading mgmt-ip: %v", err)
		}
	}

	if err = d.Set("mgmt_status", flattenSystemHaMgmtUpdateStatus(o["mgmt-status"], d, "mgmt_status", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-status"]) {
			return fmt.Errorf("Error reading mgmt-status: %v", err)
		}
	}

	return nil
}
func resourceSystemHaMgmtUpdateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHaMgmtUpdate(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMgmtUpdate resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaMgmtUpdate(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemHaMgmtUpdateRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error reading SystemHaMgmtUpdate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaMgmtUpdate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMgmtUpdate resource from API: %v", err)
	}
	return nil
}
func expandSystemHaMgmtUpdateMgmtStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMgmtUpdateMgmtInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMgmtUpdateMgmtIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMgmtUpdateMgmtIpAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaMgmtUpdate(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
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

	return &obj, nil
}
