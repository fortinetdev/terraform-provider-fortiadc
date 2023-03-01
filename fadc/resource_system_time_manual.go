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
			"year": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"month": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"mday": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"hour": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"minute": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"second": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"syncinterval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ntpserver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tz": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

	if err = d.Set("year", flattenSystemTimeManualTz(o["year"], d, "year", sv)); err != nil {
		if !fortiAPIPatch(o["year"]) {
			return fmt.Errorf("Error reading year: %v", err)
		}
	}
	if err = d.Set("month", flattenSystemTimeManualTz(o["month"], d, "month", sv)); err != nil {
		if !fortiAPIPatch(o["month"]) {
			return fmt.Errorf("Error reading month: %v", err)
		}
	}
	if err = d.Set("mday", flattenSystemTimeManualTz(o["mday"], d, "mday", sv)); err != nil {
		if !fortiAPIPatch(o["mday"]) {
			return fmt.Errorf("Error reading mday: %v", err)
		}
	}
	if err = d.Set("hour", flattenSystemTimeManualTz(o["hour"], d, "hour", sv)); err != nil {
		if !fortiAPIPatch(o["hour"]) {
			return fmt.Errorf("Error reading hour: %v", err)
		}
	}
	if err = d.Set("minute", flattenSystemTimeManualTz(o["minute"], d, "minute", sv)); err != nil {
		if !fortiAPIPatch(o["minute"]) {
			return fmt.Errorf("Error reading minute: %v", err)
		}
	}
	if err = d.Set("second", flattenSystemTimeManualTz(o["second"], d, "second", sv)); err != nil {
		if !fortiAPIPatch(o["second"]) {
			return fmt.Errorf("Error reading second: %v", err)
		}
	}
	if err = d.Set("ntpsync", flattenSystemTimeManualTz(o["ntpsync"], d, "ntpsync", sv)); err != nil {
		if !fortiAPIPatch(o["ntpsync"]) {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}
	if err = d.Set("syncinterval", flattenSystemTimeManualTz(o["syncinterval"], d, "syncinterval", sv)); err != nil {
		if !fortiAPIPatch(o["syncinterval"]) {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}
	if err = d.Set("ntpserver", flattenSystemTimeManualTz(o["ntpserver"], d, "ntpserver", sv)); err != nil {
		if !fortiAPIPatch(o["ntpserver"]) {
			return fmt.Errorf("Error reading ntpserver: %v", err)
		}
	}
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

	if v, ok := d.GetOk("year"); ok {
		if setArgNil {
			obj["year"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "year", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["year"] = t
			}
		}
	}
	if v, ok := d.GetOk("month"); ok {
		if setArgNil {
			obj["month"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "month", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["month"] = t
			}
		}
	}
	if v, ok := d.GetOk("mday"); ok {
		if setArgNil {
			obj["mday"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "mday", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mday"] = t
			}
		}
	}
	if v, ok := d.GetOk("hour"); ok {
		if setArgNil {
			obj["hour"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "hour", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hour"] = t
			}
		}
	}
	if v, ok := d.GetOk(""); ok {
		if setArgNil {
			obj[""] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj[""] = t
			}
		}
	}
	if v, ok := d.GetOk("minute"); ok {
		if setArgNil {
			obj["minute"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "minute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["minute"] = t
			}
		}
	}
	if v, ok := d.GetOk("second"); ok {
		if setArgNil {
			obj["second"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "second", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["second"] = t
			}
		}
	}
	if v, ok := d.GetOk("ntpsync"); ok {
		if setArgNil {
			obj["ntpsync"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "ntpsync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpsync"] = t
			}
		}
	}
	if v, ok := d.GetOk("ntpserver"); ok {
		if setArgNil {
			obj["ntpserver"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "ntpserver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpserver"] = t
			}
		}
	}
	if v, ok := d.GetOk("syncinterval"); ok {
		if setArgNil {
			obj["syncinterval"] = nil
		} else {
			t, err := expandSystemTimeManualDst(d, v, "syncinterval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["syncinterval"] = t
			}
		}
	}

	return &obj, nil
}
