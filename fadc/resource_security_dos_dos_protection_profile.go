// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security dos dos protection profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityDosDosProtectionProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityDosDosProtectionProfileRead,
		Update: resourceSecurityDosDosProtectionProfileUpdate,
		Create: resourceSecurityDosDosProtectionProfileCreate,
		Delete: resourceSecurityDosDosProtectionProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tcp_slow_data_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_conn_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tcp_conn_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_req_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_access_limit": &schema.Schema{
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
func resourceSecurityDosDosProtectionProfileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityDosDosProtectionProfile: type error")
	}

	obj, err := getObjectSecurityDosDosProtectionProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityDosDosProtectionProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSecurityDosDosProtectionProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SecurityDosDosProtectionProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityDosDosProtectionProfileRead(d, m)
}
func resourceSecurityDosDosProtectionProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityDosDosProtectionProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityDosDosProtectionProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityDosDosProtectionProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SecurityDosDosProtectionProfile resource: %v", err)
	}

	return resourceSecurityDosDosProtectionProfileRead(d, m)
}
func resourceSecurityDosDosProtectionProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSecurityDosDosProtectionProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SecurityDosDosProtectionProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSecurityDosDosProtectionProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSecurityDosDosProtectionProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SecurityDosDosProtectionProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityDosDosProtectionProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityDosDosProtectionProfile resource from API: %v", err)
	}
	return nil
}
func flattenSecurityDosDosProtectionProfileTcpSlowDataLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityDosDosProtectionProfileHttpConnLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityDosDosProtectionProfileTcpConnLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityDosDosProtectionProfileHttpReqLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityDosDosProtectionProfileHttpAccessLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityDosDosProtectionProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityDosDosProtectionProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("tcp_slow_data_limit", flattenSecurityDosDosProtectionProfileTcpSlowDataLimit(o["tcp_slow_data_limit"], d, "tcp_slow_data_limit", sv)); err != nil {
		if !fortiAPIPatch(o["tcp_slow_data_limit"]) {
			return fmt.Errorf("Error reading tcp_slow_data_limit: %v", err)
		}
	}

	if err = d.Set("http_conn_limit", flattenSecurityDosDosProtectionProfileHttpConnLimit(o["http_conn_limit"], d, "http_conn_limit", sv)); err != nil {
		if !fortiAPIPatch(o["http_conn_limit"]) {
			return fmt.Errorf("Error reading http_conn_limit: %v", err)
		}
	}

	if err = d.Set("tcp_conn_limit", flattenSecurityDosDosProtectionProfileTcpConnLimit(o["tcp_conn_limit"], d, "tcp_conn_limit", sv)); err != nil {
		if !fortiAPIPatch(o["tcp_conn_limit"]) {
			return fmt.Errorf("Error reading tcp_conn_limit: %v", err)
		}
	}

	if err = d.Set("http_req_limit", flattenSecurityDosDosProtectionProfileHttpReqLimit(o["http_req_limit"], d, "http_req_limit", sv)); err != nil {
		if !fortiAPIPatch(o["http_req_limit"]) {
			return fmt.Errorf("Error reading http_req_limit: %v", err)
		}
	}

	if err = d.Set("http_access_limit", flattenSecurityDosDosProtectionProfileHttpAccessLimit(o["http_access_limit"], d, "http_access_limit", sv)); err != nil {
		if !fortiAPIPatch(o["http_access_limit"]) {
			return fmt.Errorf("Error reading http_access_limit: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSecurityDosDosProtectionProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityDosDosProtectionProfileTcpSlowDataLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityDosDosProtectionProfileHttpConnLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityDosDosProtectionProfileTcpConnLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityDosDosProtectionProfileHttpReqLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityDosDosProtectionProfileHttpAccessLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityDosDosProtectionProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityDosDosProtectionProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("tcp_slow_data_limit"); ok {
		t, err := expandSecurityDosDosProtectionProfileTcpSlowDataLimit(d, v, "tcp_slow_data_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp_slow_data_limit"] = t
		}
	}

	if v, ok := d.GetOk("http_conn_limit"); ok {
		t, err := expandSecurityDosDosProtectionProfileHttpConnLimit(d, v, "http_conn_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_conn_limit"] = t
		}
	}

	if v, ok := d.GetOk("tcp_conn_limit"); ok {
		t, err := expandSecurityDosDosProtectionProfileTcpConnLimit(d, v, "tcp_conn_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp_conn_limit"] = t
		}
	}

	if v, ok := d.GetOk("http_req_limit"); ok {
		t, err := expandSecurityDosDosProtectionProfileHttpReqLimit(d, v, "http_req_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_req_limit"] = t
		}
	}

	if v, ok := d.GetOk("http_access_limit"); ok {
		t, err := expandSecurityDosDosProtectionProfileHttpAccessLimit(d, v, "http_access_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_access_limit"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityDosDosProtectionProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
