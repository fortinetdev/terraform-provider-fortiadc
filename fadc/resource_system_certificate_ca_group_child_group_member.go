// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate ca group child group member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCaGroupChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCaGroupChildGroupMemberRead,
		Update: resourceSystemCertificateCaGroupChildGroupMemberUpdate,
		Create: resourceSystemCertificateCaGroupChildGroupMemberCreate,
		Delete: resourceSystemCertificateCaGroupChildGroupMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceSystemCertificateCaGroupChildGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCaGroupChildGroupMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateCaGroupChildGroupMember: type error")
	}

	obj, err := getObjectSystemCertificateCaGroupChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCaGroupChildGroupMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemCertificateCaGroupChildGroupMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCaGroupChildGroupMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemCertificateCaGroupChildGroupMemberRead(d, m)
}
func resourceSystemCertificateCaGroupChildGroupMemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateCaGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemCertificateCaGroupChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCaGroupChildGroupMember resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCaGroupChildGroupMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCaGroupChildGroupMember resource: %v", err)
	}

	return resourceSystemCertificateCaGroupChildGroupMemberRead(d, m)
}
func resourceSystemCertificateCaGroupChildGroupMemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateCaGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemCertificateCaGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCaGroupChildGroupMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateCaGroupChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateCaGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemCertificateCaGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCaGroupChildGroupMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCaGroupChildGroupMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCaGroupChildGroupMember resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateCaGroupChildGroupMemberCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCaGroupChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCaGroupChildGroupMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ca", flattenSystemCertificateCaGroupChildGroupMemberCa(o["ca"], d, "ca", sv)); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemCertificateCaGroupChildGroupMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCaGroupChildGroupMemberCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCaGroupChildGroupMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCaGroupChildGroupMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandSystemCertificateCaGroupChildGroupMemberCa(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCaGroupChildGroupMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
