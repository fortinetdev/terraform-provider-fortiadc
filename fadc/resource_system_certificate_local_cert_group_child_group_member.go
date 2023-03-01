// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate local cert group child group member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLocalCertGroupChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLocalCertGroupChildGroupMemberRead,
		Update: resourceSystemCertificateLocalCertGroupChildGroupMemberUpdate,
		Create: resourceSystemCertificateLocalCertGroupChildGroupMemberCreate,
		Delete: resourceSystemCertificateLocalCertGroupChildGroupMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ocsp_stapling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"extra_intermediate_cag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"extra_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"extra_ocsp_stapling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"intermediate_cag": &schema.Schema{
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
func resourceSystemCertificateLocalCertGroupChildGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember: type error")
	}

	obj, err := getObjectSystemCertificateLocalCertGroupChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalCertGroupChildGroupMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemCertificateLocalCertGroupChildGroupMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalCertGroupChildGroupMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemCertificateLocalCertGroupChildGroupMemberRead(d, m)
}
func resourceSystemCertificateLocalCertGroupChildGroupMemberUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemCertificateLocalCertGroupChildGroupMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocalCertGroupChildGroupMember resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateLocalCertGroupChildGroupMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocalCertGroupChildGroupMember resource: %v", err)
	}

	return resourceSystemCertificateLocalCertGroupChildGroupMemberRead(d, m)
}
func resourceSystemCertificateLocalCertGroupChildGroupMemberDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemCertificateLocalCertGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLocalCertGroupChildGroupMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateLocalCertGroupChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemCertificateLocalCertGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalCertGroupChildGroupMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLocalCertGroupChildGroupMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalCertGroupChildGroupMember resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateLocalCertGroupChildGroupMemberDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberLocalCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberOcspStapling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberExtraIntermediateCag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberExtraLocalCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberExtraOcspStapling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberIntermediateCag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertGroupChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLocalCertGroupChildGroupMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("default", flattenSystemCertificateLocalCertGroupChildGroupMemberDefault(o["default"], d, "default", sv)); err != nil {
		if !fortiAPIPatch(o["default"]) {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemCertificateLocalCertGroupChildGroupMemberLocalCert(o["local_cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local_cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("ocsp_stapling", flattenSystemCertificateLocalCertGroupChildGroupMemberOcspStapling(o["OCSP_stapling"], d, "ocsp_stapling", sv)); err != nil {
		if !fortiAPIPatch(o["OCSP_stapling"]) {
			return fmt.Errorf("Error reading OCSP_stapling: %v", err)
		}
	}

	if err = d.Set("extra_intermediate_cag", flattenSystemCertificateLocalCertGroupChildGroupMemberExtraIntermediateCag(o["extra_intermediate_cag"], d, "extra_intermediate_cag", sv)); err != nil {
		if !fortiAPIPatch(o["extra_intermediate_cag"]) {
			return fmt.Errorf("Error reading extra_intermediate_cag: %v", err)
		}
	}

	if err = d.Set("extra_local_cert", flattenSystemCertificateLocalCertGroupChildGroupMemberExtraLocalCert(o["extra_local_cert"], d, "extra_local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["extra_local_cert"]) {
			return fmt.Errorf("Error reading extra_local_cert: %v", err)
		}
	}

	if err = d.Set("extra_ocsp_stapling", flattenSystemCertificateLocalCertGroupChildGroupMemberExtraOcspStapling(o["extra_OCSP_stapling"], d, "extra_ocsp_stapling", sv)); err != nil {
		if !fortiAPIPatch(o["extra_OCSP_stapling"]) {
			return fmt.Errorf("Error reading extra_OCSP_stapling: %v", err)
		}
	}

	if err = d.Set("intermediate_cag", flattenSystemCertificateLocalCertGroupChildGroupMemberIntermediateCag(o["intermediate_cag"], d, "intermediate_cag", sv)); err != nil {
		if !fortiAPIPatch(o["intermediate_cag"]) {
			return fmt.Errorf("Error reading intermediate_cag: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemCertificateLocalCertGroupChildGroupMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberLocalCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberOcspStapling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberExtraIntermediateCag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberExtraLocalCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberExtraOcspStapling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberIntermediateCag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertGroupChildGroupMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLocalCertGroupChildGroupMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberDefault(d, v, "default", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberLocalCert(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local_cert"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_stapling"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberOcspStapling(d, v, "ocsp_stapling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["OCSP_stapling"] = t
		}
	}

	if v, ok := d.GetOk("extra_intermediate_cag"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberExtraIntermediateCag(d, v, "extra_intermediate_cag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra_intermediate_cag"] = t
		}
	}

	if v, ok := d.GetOk("extra_local_cert"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberExtraLocalCert(d, v, "extra_local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra_local_cert"] = t
		}
	}

	if v, ok := d.GetOk("extra_ocsp_stapling"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberExtraOcspStapling(d, v, "extra_ocsp_stapling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra_OCSP_stapling"] = t
		}
	}

	if v, ok := d.GetOk("intermediate_cag"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberIntermediateCag(d, v, "intermediate_cag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intermediate_cag"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLocalCertGroupChildGroupMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
