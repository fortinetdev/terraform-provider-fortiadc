// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate ca group.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCaGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCaGroupRead,
		Update: resourceSystemCertificateCaGroupUpdate,
		Create: resourceSystemCertificateCaGroupCreate,
		Delete: resourceSystemCertificateCaGroupDelete,

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
func resourceSystemCertificateCaGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCaGroup: type error")
	}

	obj, err := getObjectSystemCertificateCaGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCaGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCaGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCaGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateCaGroupRead(d, m)
}
func resourceSystemCertificateCaGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateCaGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCaGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCaGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCaGroup resource: %v", err)
	}

	return resourceSystemCertificateCaGroupRead(d, m)
}
func resourceSystemCertificateCaGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateCaGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCaGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateCaGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateCaGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCaGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCaGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCaGroup resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateCaGroupMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCaGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateCaGroupMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCaGroupMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCaGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCaGroupMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
