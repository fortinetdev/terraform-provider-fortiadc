// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system azure lb backend ip.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAzureLbBackendIp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAzureLbBackendIpRead,
		Update: resourceSystemAzureLbBackendIpUpdate,
		Create: resourceSystemAzureLbBackendIpCreate,
		Delete: resourceSystemAzureLbBackendIpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
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
func resourceSystemAzureLbBackendIpCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAzureLbBackendIp: type error")
	}

	obj, err := getObjectSystemAzureLbBackendIp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAzureLbBackendIp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAzureLbBackendIp(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAzureLbBackendIp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAzureLbBackendIpRead(d, m)
}
func resourceSystemAzureLbBackendIpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAzureLbBackendIp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAzureLbBackendIp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAzureLbBackendIp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAzureLbBackendIp resource: %v", err)
	}

	return resourceSystemAzureLbBackendIpRead(d, m)
}
func resourceSystemAzureLbBackendIpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemAzureLbBackendIp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAzureLbBackendIp resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemAzureLbBackendIpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAzureLbBackendIp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAzureLbBackendIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAzureLbBackendIp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAzureLbBackendIp resource from API: %v", err)
	}
	return nil
}
func flattenSystemAzureLbBackendIpIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAzureLbBackendIpMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAzureLbBackendIp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip", flattenSystemAzureLbBackendIpIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemAzureLbBackendIpMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemAzureLbBackendIpIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAzureLbBackendIpMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAzureLbBackendIp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemAzureLbBackendIpIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemAzureLbBackendIpMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
