// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system scripting.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemScripting() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemScriptingRead,
		Update: resourceSystemScriptingUpdate,
		Create: resourceSystemScriptingCreate,
		Delete: resourceSystemScriptingDelete,

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
func resourceSystemScriptingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemScripting: type error")
	}

	obj, err := getObjectSystemScripting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemScripting resource while getting object: %v", err)
	}

	_, err = c.CreateSystemScripting(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemScripting resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemScriptingRead(d, m)
}
func resourceSystemScriptingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemScripting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemScripting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemScripting(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemScripting resource: %v", err)
	}

	return resourceSystemScriptingRead(d, m)
}
func resourceSystemScriptingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemScripting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemScripting resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemScriptingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemScripting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemScripting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemScripting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemScripting resource from API: %v", err)
	}
	return nil
}
func flattenSystemScriptingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScriptingScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemScripting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemScriptingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemScriptingScript(o["script"], d, "script", sv)); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	return nil
}

func expandSystemScriptingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScriptingScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemScripting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemScriptingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok {
		t, err := expandSystemScriptingScript(d, v, "script", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	return &obj, nil
}
