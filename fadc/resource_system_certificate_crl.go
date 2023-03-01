// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate crl.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCrlRead,
		Update: resourceSystemCertificateCrlUpdate,
		Create: resourceSystemCertificateCrlCreate,
		Delete: resourceSystemCertificateCrlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crldp_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_ca_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"update_ahead_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ca_id": &schema.Schema{
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
func resourceSystemCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrl: type error")
	}

	obj, err := getObjectSystemCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCrl(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateCrlRead(d, m)
}
func resourceSystemCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCrl(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrl resource: %v", err)
	}

	return resourceSystemCertificateCrlRead(d, m)
}
func resourceSystemCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateCrl(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateCrl(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCrl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateCrlLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlCrldpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlUseCaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlUpdateAheadTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlCaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCrl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ldap_server", flattenSystemCertificateCrlLdapServer(o["ldap_server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap_server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("crldp_status", flattenSystemCertificateCrlCrldpStatus(o["crldp_status"], d, "crldp_status", sv)); err != nil {
		if !fortiAPIPatch(o["crldp_status"]) {
			return fmt.Errorf("Error reading crldp_status: %v", err)
		}
	}

	if err = d.Set("use_ca_id", flattenSystemCertificateCrlUseCaId(o["use_ca_id"], d, "use_ca_id", sv)); err != nil {
		if !fortiAPIPatch(o["use_ca_id"]) {
			return fmt.Errorf("Error reading use_ca_id: %v", err)
		}
	}

	if err = d.Set("update_ahead_time", flattenSystemCertificateCrlUpdateAheadTime(o["update_ahead_time"], d, "update_ahead_time", sv)); err != nil {
		if !fortiAPIPatch(o["update_ahead_time"]) {
			return fmt.Errorf("Error reading update_ahead_time: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemCertificateCrlUpdateInterval(o["update_interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update_interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemCertificateCrlMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ca_id", flattenSystemCertificateCrlCaId(o["ca_id"], d, "ca_id", sv)); err != nil {
		if !fortiAPIPatch(o["ca_id"]) {
			return fmt.Errorf("Error reading ca_id: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCrlLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlCrldpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlUseCaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlUpdateAheadTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlCaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCrl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandSystemCertificateCrlLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap_server"] = t
		}
	}

	if v, ok := d.GetOk("crldp_status"); ok {
		t, err := expandSystemCertificateCrlCrldpStatus(d, v, "crldp_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crldp_status"] = t
		}
	}

	if v, ok := d.GetOk("use_ca_id"); ok {
		t, err := expandSystemCertificateCrlUseCaId(d, v, "use_ca_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use_ca_id"] = t
		}
	}

	if v, ok := d.GetOk("update_ahead_time"); ok {
		t, err := expandSystemCertificateCrlUpdateAheadTime(d, v, "update_ahead_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update_ahead_time"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {
		t, err := expandSystemCertificateCrlUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update_interval"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCrlMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("ca_id"); ok {
		t, err := expandSystemCertificateCrlCaId(d, v, "ca_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca_id"] = t
		}
	}

	return &obj, nil
}
