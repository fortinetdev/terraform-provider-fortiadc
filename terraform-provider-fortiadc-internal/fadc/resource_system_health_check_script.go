// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system health check script.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHealthCheckScript() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemHealthCheckScriptRead,
		Update: resourceSystemHealthCheckScriptUpdate,
		Create: resourceSystemHealthCheckScriptCreate,
		Delete: resourceSystemHealthCheckScriptDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSystemHealthCheckScriptCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemHealthCheckScript: type error")
	}

	obj, err := getObjectSystemHealthCheckScript(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheckScript resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHealthCheckScript(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheckScript resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemHealthCheckScriptRead(d, m)
}
func resourceSystemHealthCheckScriptUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHealthCheckScript(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheckScript resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHealthCheckScript(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheckScript resource: %v", err)
	}

	return resourceSystemHealthCheckScriptRead(d, m)
}
func resourceSystemHealthCheckScriptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemHealthCheckScript(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHealthCheckScript resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemHealthCheckScriptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemHealthCheckScript(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheckScript resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHealthCheckScript(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheckScript resource from API: %v", err)
	}
	return nil
}
func flattenSystemHealthCheckScriptMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckScriptFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHealthCheckScript(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemHealthCheckScriptMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemHealthCheckScriptFile(o["script"], d, "script", sv)); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	return nil
}

func expandSystemHealthCheckScriptMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckScriptFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHealthCheckScript(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemHealthCheckScriptMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok {
		t, err := expandSystemHealthCheckScriptFile(d, v, "script", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	return &obj, nil
}
