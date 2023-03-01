// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  user adfs publish.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserAdfsPublish() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserAdfsPublishRead,
		Update: resourceUserAdfsPublishUpdate,
		Create: resourceUserAdfsPublishCreate,
		Delete: resourceUserAdfsPublishDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_server_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"external_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"preauth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"relying_party": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"proxy": &schema.Schema{
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
func resourceUserAdfsPublishCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserAdfsPublish: type error")
	}

	obj, err := getObjectUserAdfsPublish(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserAdfsPublish resource while getting object: %v", err)
	}

	_, err = c.CreateUserAdfsPublish(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserAdfsPublish resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserAdfsPublishRead(d, m)
}
func resourceUserAdfsPublishUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserAdfsPublish(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserAdfsPublish resource while getting object: %v", err)
	}

	_, err = c.UpdateUserAdfsPublish(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserAdfsPublish resource: %v", err)
	}

	return resourceUserAdfsPublishRead(d, m)
}
func resourceUserAdfsPublishDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserAdfsPublish(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserAdfsPublish resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceUserAdfsPublishRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserAdfsPublish(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserAdfsPublish resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserAdfsPublish(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserAdfsPublish resource from API: %v", err)
	}
	return nil
}
func flattenUserAdfsPublishStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdfsPublishBackendServerUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdfsPublishExternalUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdfsPublishPreauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdfsPublishRelyingParty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdfsPublishMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdfsPublishProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserAdfsPublish(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenUserAdfsPublishStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("backend_server_url", flattenUserAdfsPublishBackendServerUrl(o["backend-server-url"], d, "backend_server_url", sv)); err != nil {
		if !fortiAPIPatch(o["backend-server-url"]) {
			return fmt.Errorf("Error reading backend-server-url: %v", err)
		}
	}

	if err = d.Set("external_url", flattenUserAdfsPublishExternalUrl(o["external-url"], d, "external_url", sv)); err != nil {
		if !fortiAPIPatch(o["external-url"]) {
			return fmt.Errorf("Error reading external-url: %v", err)
		}
	}

	if err = d.Set("preauth", flattenUserAdfsPublishPreauth(o["preauth"], d, "preauth", sv)); err != nil {
		if !fortiAPIPatch(o["preauth"]) {
			return fmt.Errorf("Error reading preauth: %v", err)
		}
	}

	if err = d.Set("relying_party", flattenUserAdfsPublishRelyingParty(o["relying-party"], d, "relying_party", sv)); err != nil {
		if !fortiAPIPatch(o["relying-party"]) {
			return fmt.Errorf("Error reading relying-party: %v", err)
		}
	}

	if err = d.Set("mkey", flattenUserAdfsPublishMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("proxy", flattenUserAdfsPublishProxy(o["proxy"], d, "proxy", sv)); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	return nil
}

func expandUserAdfsPublishStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdfsPublishBackendServerUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdfsPublishExternalUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdfsPublishPreauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdfsPublishRelyingParty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdfsPublishMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdfsPublishProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserAdfsPublish(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandUserAdfsPublishStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("backend_server_url"); ok {
		t, err := expandUserAdfsPublishBackendServerUrl(d, v, "backend_server_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend-server-url"] = t
		}
	}

	if v, ok := d.GetOk("external_url"); ok {
		t, err := expandUserAdfsPublishExternalUrl(d, v, "external_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-url"] = t
		}
	}

	if v, ok := d.GetOk("preauth"); ok {
		t, err := expandUserAdfsPublishPreauth(d, v, "preauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preauth"] = t
		}
	}

	if v, ok := d.GetOk("relying_party"); ok {
		t, err := expandUserAdfsPublishRelyingParty(d, v, "relying_party", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relying-party"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserAdfsPublishMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		t, err := expandUserAdfsPublishProxy(d, v, "proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	return &obj, nil
}
