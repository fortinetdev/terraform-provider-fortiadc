// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate local cert group.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLocalCertGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLocalCertGroupRead,
		Update: resourceSystemCertificateLocalCertGroupUpdate,
		Create: resourceSystemCertificateLocalCertGroupCreate,
		Delete: resourceSystemCertificateLocalCertGroupDelete,

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
func resourceSystemCertificateLocalCertGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroup: type error")
	}

	obj, err := getObjectSystemCertificateLocalCertGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalCertGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateLocalCertGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalCertGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateLocalCertGroupRead(d, m)
}
func resourceSystemCertificateLocalCertGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateLocalCertGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocalCertGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateLocalCertGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocalCertGroup resource: %v", err)
	}

	return resourceSystemCertificateLocalCertGroupRead(d, m)
}
func resourceSystemCertificateLocalCertGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateLocalCertGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLocalCertGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateLocalCertGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateLocalCertGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalCertGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLocalCertGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalCertGroup resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateLocalCertGroupMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLocalCertGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateLocalCertGroupMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLocalCertGroupMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLocalCertGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLocalCertGroupMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
