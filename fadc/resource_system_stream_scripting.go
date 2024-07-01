// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system stream scripting.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemStreamScripting() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemStreamScriptingRead,
		Update: resourceSystemStreamScriptingUpdate,
		Create: resourceSystemStreamScriptingCreate,
		Delete: resourceSystemStreamScriptingDelete,

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
func resourceSystemStreamScriptingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemStreamScripting: type error")
	}

	obj, err := getObjectSystemStreamScripting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemStreamScripting resource while getting object: %v", err)
	}

	_, err = c.CreateSystemStreamScripting(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemStreamScripting resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemStreamScriptingRead(d, m)
}
func resourceSystemStreamScriptingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemStreamScripting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemStreamScripting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStreamScripting(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemStreamScripting resource: %v", err)
	}

	return resourceSystemStreamScriptingRead(d, m)
}
func resourceSystemStreamScriptingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemStreamScripting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStreamScripting resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemStreamScriptingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemStreamScripting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemStreamScripting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStreamScripting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemStreamScripting resource from API: %v", err)
	}
	return nil
}
func flattenSystemStreamScriptingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStreamScriptingScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemStreamScripting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemStreamScriptingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemStreamScriptingScript(o["script"], d, "script", sv)); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	return nil
}

func expandSystemStreamScriptingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStreamScriptingScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStreamScripting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemStreamScriptingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok {
		t, err := expandSystemStreamScriptingScript(d, v, "script", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	return &obj, nil
}
