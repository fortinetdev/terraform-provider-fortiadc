// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  config sync list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigSyncList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceConfigSyncListRead,
		Update: resourceConfigSyncListUpdate,
		Create: resourceConfigSyncListCreate,
		Delete: resourceConfigSyncListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
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
func resourceConfigSyncListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing ConfigSyncList: type error")
	}

	obj, err := getObjectConfigSyncList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ConfigSyncList resource while getting object: %v", err)
	}

	_, err = c.CreateConfigSyncList(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ConfigSyncList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceConfigSyncListRead(d, m)
}
func resourceConfigSyncListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectConfigSyncList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ConfigSyncList resource while getting object: %v", err)
	}

	_, err = c.UpdateConfigSyncList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ConfigSyncList resource: %v", err)
	}

	return resourceConfigSyncListRead(d, m)
}
func resourceConfigSyncListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteConfigSyncList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ConfigSyncList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceConfigSyncListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadConfigSyncList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ConfigSyncList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectConfigSyncList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ConfigSyncList resource from API: %v", err)
	}
	return nil
}
func flattenConfigSyncListComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListFilename(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenConfigSyncListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectConfigSyncList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("comment", flattenConfigSyncListComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("status", flattenConfigSyncListStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("filename", flattenConfigSyncListFilename(o["filename"], d, "filename", sv)); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenConfigSyncListServerIp(o["server-ip"], d, "server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["server-ip"]) {
			return fmt.Errorf("Error reading server-ip: %v", err)
		}
	}

	if err = d.Set("password", flattenConfigSyncListPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("type", flattenConfigSyncListType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("port", flattenConfigSyncListPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenConfigSyncListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandConfigSyncListComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListFilename(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandConfigSyncListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectConfigSyncList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandConfigSyncListComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandConfigSyncListStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok {
		t, err := expandConfigSyncListFilename(d, v, "filename", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok {
		t, err := expandConfigSyncListServerIp(d, v, "server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandConfigSyncListPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandConfigSyncListType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandConfigSyncListPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandConfigSyncListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
