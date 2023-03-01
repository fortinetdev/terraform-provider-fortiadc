// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate local cert group child group member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemCertificateLocalCertGroupChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateLocalCertGroupChildGroupMemberRead,
		Schema: map[string]*schema.Schema{
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ocsp_stapling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extra_intermediate_cag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extra_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extra_ocsp_stapling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"intermediate_cag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceSystemCertificateLocalCertGroupChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCertificateLocalCertGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemCertificateLocalCertGroupChildGroupMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateLocalCertGroupChildGroupMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberOcspStapling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberExtraIntermediateCag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberExtraLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberExtraOcspStapling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberIntermediateCag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCertificateLocalCertGroupChildGroupMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberDefault(o["default"], d, "default")); err != nil {
		if !fortiAPIPatch(o["default"]) {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("local_cert", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberLocalCert(o["local_cert"], d, "local_cert")); err != nil {
		if !fortiAPIPatch(o["local_cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("ocsp_stapling", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberOcspStapling(o["OCSP_stapling"], d, "ocsp_stapling")); err != nil {
		if !fortiAPIPatch(o["OCSP_stapling"]) {
			return fmt.Errorf("Error reading OCSP_stapling: %v", err)
		}
	}

	if err = d.Set("extra_intermediate_cag", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberExtraIntermediateCag(o["extra_intermediate_cag"], d, "extra_intermediate_cag")); err != nil {
		if !fortiAPIPatch(o["extra_intermediate_cag"]) {
			return fmt.Errorf("Error reading extra_intermediate_cag: %v", err)
		}
	}

	if err = d.Set("extra_local_cert", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberExtraLocalCert(o["extra_local_cert"], d, "extra_local_cert")); err != nil {
		if !fortiAPIPatch(o["extra_local_cert"]) {
			return fmt.Errorf("Error reading extra_local_cert: %v", err)
		}
	}

	if err = d.Set("extra_ocsp_stapling", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberExtraOcspStapling(o["extra_OCSP_stapling"], d, "extra_ocsp_stapling")); err != nil {
		if !fortiAPIPatch(o["extra_OCSP_stapling"]) {
			return fmt.Errorf("Error reading extra_OCSP_stapling: %v", err)
		}
	}

	if err = d.Set("intermediate_cag", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberIntermediateCag(o["intermediate_cag"], d, "intermediate_cag")); err != nil {
		if !fortiAPIPatch(o["intermediate_cag"]) {
			return fmt.Errorf("Error reading intermediate_cag: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemCertificateLocalCertGroupChildGroupMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
