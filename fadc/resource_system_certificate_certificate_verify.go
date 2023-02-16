// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate certificate verify.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCertificateVerify() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCertificateVerifyRead,
		Update: resourceSystemCertificateCertificateVerifyUpdate,
		Create: resourceSystemCertificateCertificateVerifyCreate,
		Delete: resourceSystemCertificateCertificateVerifyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
func resourceSystemCertificateCertificateVerifyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCertificateVerify: type error")
	}

	obj, err := getObjectSystemCertificateCertificateVerify(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCertificateVerify resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCertificateVerify(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCertificateVerify resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateCertificateVerifyRead(d, m)
}
func resourceSystemCertificateCertificateVerifyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateCertificateVerify(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCertificateVerify resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCertificateVerify(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCertificateVerify resource: %v", err)
	}

	return resourceSystemCertificateCertificateVerifyRead(d, m)
}
func resourceSystemCertificateCertificateVerifyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateCertificateVerify(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCertificateVerify resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateCertificateVerifyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateCertificateVerify(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCertificateVerify resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCertificateVerify(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCertificateVerify resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateCertificateVerifyMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCertificateVerify(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateCertificateVerifyMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCertificateVerifyMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCertificateVerify(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCertificateVerifyMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
