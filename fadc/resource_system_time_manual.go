// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system time manual.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTimeManual() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemTimeManualRead,
		Update: resourceSystemTimeManualUpdate,
		Create: resourceSystemTimeManualUpdate,
		Delete: resourceSystemTimeManualDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tz": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemTimeManualUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemTimeManual(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemTimeManual resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemTimeManual(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemTimeManual resource: %v", err)
	}

	d.SetId("SystemTimeManual")
	return resourceSystemTimeManualRead(d, m)
}
func resourceSystemTimeManualDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemTimeManual(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemTimeManual resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemTimeManual(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemTimeManual resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemTimeManualRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemTimeManual(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemTimeManual resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemTimeManual(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemTimeManual resource from API: %v", err)
	}
	return nil
}
func flattenSystemTimeManualTz(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemTimeManualDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemTimeManual(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("tz", flattenSystemTimeManualTz(o["tz"], d, "tz", sv)); err != nil {
		if !fortiAPIPatch(o["tz"]) {
			return fmt.Errorf("Error reading tz: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemTimeManualDst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	return nil
}

func expandSystemTimeManualTz(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemTimeManualDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemTimeManual(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("tz"); ok {
		if setArgNil {
			obj["tz"] = nil
		} else {
			t, err := expandSystemTimeManualTz(d, v, "tz", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tz"] = t
			}
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		if setArgNil {
			obj["dst"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "dst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dst"] = t
			}
		}
	}

	return &obj, nil
}
