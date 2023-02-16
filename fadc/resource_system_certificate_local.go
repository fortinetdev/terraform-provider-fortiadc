// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate local.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLocalRead,
		Update: resourceSystemCertificateLocalUpdate,
		Create: resourceSystemCertificateLocalCreate,
		Delete: resourceSystemCertificateLocalDelete,

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
func resourceSystemCertificateLocalCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocal: type error")
	}

	obj, err := getObjectSystemCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocal resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateLocal(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocal resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateLocalRead(d, m)
}
func resourceSystemCertificateLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateLocal(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocal resource: %v", err)
	}

	return resourceSystemCertificateLocalRead(d, m)
}
func resourceSystemCertificateLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateLocal(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateLocal(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLocal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocal resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateLocalMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLocal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateLocalMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLocalMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLocalMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
