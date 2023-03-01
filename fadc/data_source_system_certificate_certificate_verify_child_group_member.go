// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate certificate verify child group member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemCertificateCertificateVerifyChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateCertificateVerifyChildGroupMemberRead,
		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ocsp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
func dataSourceSystemCertificateCertificateVerifyChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCertificateCertificateVerifyChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemCertificateCertificateVerifyChildGroupMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateCertificateVerifyChildGroupMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberCrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberOcsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCertificateCertificateVerifyChildGroupMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ca", dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberCa(o["ca"], d, "ca")); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("crl", dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberCrl(o["crl"], d, "crl")); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ocsp", dataSourceFlattenSystemCertificateCertificateVerifyChildGroupMemberOcsp(o["ocsp"], d, "ocsp")); err != nil {
		if !fortiAPIPatch(o["ocsp"]) {
			return fmt.Errorf("Error reading ocsp: %v", err)
		}
	}

	return nil
}
