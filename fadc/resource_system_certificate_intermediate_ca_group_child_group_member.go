// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate intermediate ca group child group member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateIntermediateCaGroupChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateIntermediateCaGroupChildGroupMemberRead,
		Update: resourceSystemCertificateIntermediateCaGroupChildGroupMemberUpdate,
		Create: resourceSystemCertificateIntermediateCaGroupChildGroupMemberCreate,
		Delete: resourceSystemCertificateIntermediateCaGroupChildGroupMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"intmed_ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceSystemCertificateIntermediateCaGroupChildGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember: type error")
	}

	obj, err := getObjectSystemCertificateIntermediateCaGroupChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCaGroupChildGroupMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemCertificateIntermediateCaGroupChildGroupMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCaGroupChildGroupMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemCertificateIntermediateCaGroupChildGroupMemberRead(d, m)
}
func resourceSystemCertificateIntermediateCaGroupChildGroupMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemCertificateIntermediateCaGroupChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCaGroupChildGroupMember resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateIntermediateCaGroupChildGroupMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCaGroupChildGroupMember resource: %v", err)
	}

	return resourceSystemCertificateIntermediateCaGroupChildGroupMemberRead(d, m)
}
func resourceSystemCertificateIntermediateCaGroupChildGroupMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemCertificateIntermediateCaGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateIntermediateCaGroupChildGroupMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateIntermediateCaGroupChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemCertificateIntermediateCaGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCaGroupChildGroupMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateIntermediateCaGroupChildGroupMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCaGroupChildGroupMember resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateIntermediateCaGroupChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateIntermediateCaGroupChildGroupMemberIntmedCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateIntermediateCaGroupChildGroupMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateIntermediateCaGroupChildGroupMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("intmed_ca", flattenSystemCertificateIntermediateCaGroupChildGroupMemberIntmedCa(o["intmed_ca"], d, "intmed_ca", sv)); err != nil {
		if !fortiAPIPatch(o["intmed_ca"]) {
			return fmt.Errorf("Error reading intmed_ca: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateIntermediateCaGroupChildGroupMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateIntermediateCaGroupChildGroupMemberIntmedCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateIntermediateCaGroupChildGroupMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateIntermediateCaGroupChildGroupMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("intmed_ca"); ok {
		t, err := expandSystemCertificateIntermediateCaGroupChildGroupMemberIntmedCa(d, v, "intmed_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intmed_ca"] = t
		}
	}

	return &obj, nil
}
