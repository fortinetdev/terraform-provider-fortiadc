// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate certificate verify child group member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCertificateVerifyChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCertificateVerifyChildGroupMemberRead,
		Update: resourceSystemCertificateCertificateVerifyChildGroupMemberUpdate,
		Create: resourceSystemCertificateCertificateVerifyChildGroupMemberCreate,
		Delete: resourceSystemCertificateCertificateVerifyChildGroupMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ocsp": &schema.Schema{
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
func resourceSystemCertificateCertificateVerifyChildGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember: type error")
	}

	obj, err := getObjectSystemCertificateCertificateVerifyChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCertificateVerifyChildGroupMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemCertificateCertificateVerifyChildGroupMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCertificateVerifyChildGroupMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemCertificateCertificateVerifyChildGroupMemberRead(d, m)
}
func resourceSystemCertificateCertificateVerifyChildGroupMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemCertificateCertificateVerifyChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCertificateVerifyChildGroupMember resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCertificateVerifyChildGroupMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCertificateVerifyChildGroupMember resource: %v", err)
	}

	return resourceSystemCertificateCertificateVerifyChildGroupMemberRead(d, m)
}
func resourceSystemCertificateCertificateVerifyChildGroupMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemCertificateCertificateVerifyChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCertificateVerifyChildGroupMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateCertificateVerifyChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemCertificateCertificateVerifyChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCertificateVerifyChildGroupMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCertificateVerifyChildGroupMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCertificateVerifyChildGroupMember resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateCertificateVerifyChildGroupMemberCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCertificateVerifyChildGroupMemberCrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCertificateVerifyChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCertificateVerifyChildGroupMemberOcsp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCertificateVerifyChildGroupMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ca", flattenSystemCertificateCertificateVerifyChildGroupMemberCa(o["ca"], d, "ca", sv)); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("crl", flattenSystemCertificateCertificateVerifyChildGroupMemberCrl(o["crl"], d, "crl", sv)); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemCertificateCertificateVerifyChildGroupMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ocsp", flattenSystemCertificateCertificateVerifyChildGroupMemberOcsp(o["ocsp"], d, "ocsp", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp"]) {
			return fmt.Errorf("Error reading ocsp: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCertificateVerifyChildGroupMemberCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCertificateVerifyChildGroupMemberCrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCertificateVerifyChildGroupMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCertificateVerifyChildGroupMemberOcsp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCertificateVerifyChildGroupMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandSystemCertificateCertificateVerifyChildGroupMemberCa(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("crl"); ok {
		t, err := expandSystemCertificateCertificateVerifyChildGroupMemberCrl(d, v, "crl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCertificateVerifyChildGroupMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("ocsp"); ok {
		t, err := expandSystemCertificateCertificateVerifyChildGroupMemberOcsp(d, v, "ocsp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp"] = t
		}
	}

	return &obj, nil
}
