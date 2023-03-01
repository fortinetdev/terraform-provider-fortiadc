// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate intermediate ca group child group member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemCertificateIntermediateCaGroupChildGroupMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateIntermediateCaGroupChildGroupMemberRead,
		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"intmed_ca": &schema.Schema{
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
func dataSourceSystemCertificateIntermediateCaGroupChildGroupMemberRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCertificateIntermediateCaGroupChildGroupMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemCertificateIntermediateCaGroupChildGroupMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateIntermediateCaGroupChildGroupMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemCertificateIntermediateCaGroupChildGroupMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateIntermediateCaGroupChildGroupMemberIntmedCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCertificateIntermediateCaGroupChildGroupMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mkey", dataSourceFlattenSystemCertificateIntermediateCaGroupChildGroupMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("intmed_ca", dataSourceFlattenSystemCertificateIntermediateCaGroupChildGroupMemberIntmedCa(o["intmed_ca"], d, "intmed_ca")); err != nil {
		if !fortiAPIPatch(o["intmed_ca"]) {
			return fmt.Errorf("Error reading intmed_ca: %v", err)
		}
	}

	return nil
}
