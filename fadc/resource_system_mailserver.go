// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system mailserver.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMailserver() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemMailserverRead,
		Update: resourceSystemMailserverUpdate,
		Create: resourceSystemMailserverUpdate,
		Delete: resourceSystemMailserverDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemMailserverUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemMailserver(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemMailserver resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMailserver(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemMailserver resource: %v", err)
	}

	d.SetId("SystemMailserver")
	return resourceSystemMailserverRead(d, m)
}
func resourceSystemMailserverDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemMailserver(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemMailserver resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMailserver(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemMailserver resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemMailserverRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemMailserver(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemMailserver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMailserver(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemMailserver resource from API: %v", err)
	}
	return nil
}
func flattenSystemMailserverUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMailserverAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMailserverAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMailserverSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMailserverPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMailserverPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemMailserver(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("username", flattenSystemMailserverUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("auth", flattenSystemMailserverAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemMailserverAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("security", flattenSystemMailserverSecurity(o["security"], d, "security", sv)); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemMailserverPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemMailserverPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func expandSystemMailserverUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMailserverAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMailserverAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMailserverSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMailserverPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMailserverPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMailserver(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("username"); ok {
		if setArgNil {
			obj["username"] = nil
		} else {
			t, err := expandSystemMailserverUsername(d, v, "username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		if setArgNil {
			obj["auth"] = nil
		} else {
			t, err := expandSystemMailserverAuth(d, v, "auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("address"); ok {
		if setArgNil {
			obj["address"] = nil
		} else {
			t, err := expandSystemMailserverAddress(d, v, "address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address"] = t
			}
		}
	}

	if v, ok := d.GetOk("security"); ok {
		if setArgNil {
			obj["security"] = nil
		} else {
			t, err := expandSystemMailserverSecurity(d, v, "security", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["security"] = t
			}
		}
	}

	if v, ok := d.GetOk("password"); ok {
		if setArgNil {
			obj["password"] = nil
		} else {
			t, err := expandSystemMailserverPassword(d, v, "password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["password"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {
			t, err := expandSystemMailserverPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	return &obj, nil
}
