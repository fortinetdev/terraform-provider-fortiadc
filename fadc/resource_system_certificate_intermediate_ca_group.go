// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate intermediate ca group.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateIntermediateCaGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateIntermediateCaGroupRead,
		Update: resourceSystemCertificateIntermediateCaGroupUpdate,
		Create: resourceSystemCertificateIntermediateCaGroupCreate,
		Delete: resourceSystemCertificateIntermediateCaGroupDelete,

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
func resourceSystemCertificateIntermediateCaGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroup: type error")
	}

	obj, err := getObjectSystemCertificateIntermediateCaGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCaGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateIntermediateCaGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCaGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateIntermediateCaGroupRead(d, m)
}
func resourceSystemCertificateIntermediateCaGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateIntermediateCaGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCaGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateIntermediateCaGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCaGroup resource: %v", err)
	}

	return resourceSystemCertificateIntermediateCaGroupRead(d, m)
}
func resourceSystemCertificateIntermediateCaGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateIntermediateCaGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateIntermediateCaGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateIntermediateCaGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateIntermediateCaGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCaGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateIntermediateCaGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCaGroup resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateIntermediateCaGroupMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateIntermediateCaGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateIntermediateCaGroupMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateIntermediateCaGroupMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateIntermediateCaGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateIntermediateCaGroupMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
