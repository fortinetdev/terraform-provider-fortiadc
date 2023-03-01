// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate crl.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateCrlRead,
		Schema: map[string]*schema.Schema{
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"crldp_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_ca_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_ahead_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ca_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceSystemCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrl: type error")
	}

	o, err := c.ReadSystemCertificateCrl(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateCrl: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemCertificateCrl(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateCrl from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemCertificateCrlLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCrlCrldpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCrlUseCaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCrlUpdateAheadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCrlMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateCrlCaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCertificateCrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ldap_server", dataSourceFlattenSystemCertificateCrlLdapServer(o["ldap_server"], d, "ldap_server")); err != nil {
		if !fortiAPIPatch(o["ldap_server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("crldp_status", dataSourceFlattenSystemCertificateCrlCrldpStatus(o["crldp_status"], d, "crldp_status")); err != nil {
		if !fortiAPIPatch(o["crldp_status"]) {
			return fmt.Errorf("Error reading crldp_status: %v", err)
		}
	}

	if err = d.Set("use_ca_id", dataSourceFlattenSystemCertificateCrlUseCaId(o["use_ca_id"], d, "use_ca_id")); err != nil {
		if !fortiAPIPatch(o["use_ca_id"]) {
			return fmt.Errorf("Error reading use_ca_id: %v", err)
		}
	}

	if err = d.Set("update_ahead_time", dataSourceFlattenSystemCertificateCrlUpdateAheadTime(o["update_ahead_time"], d, "update_ahead_time")); err != nil {
		if !fortiAPIPatch(o["update_ahead_time"]) {
			return fmt.Errorf("Error reading update_ahead_time: %v", err)
		}
	}

	if err = d.Set("update_interval", dataSourceFlattenSystemCertificateCrlUpdateInterval(o["update_interval"], d, "update_interval")); err != nil {
		if !fortiAPIPatch(o["update_interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemCertificateCrlMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ca_id", dataSourceFlattenSystemCertificateCrlCaId(o["ca_id"], d, "ca_id")); err != nil {
		if !fortiAPIPatch(o["ca_id"]) {
			return fmt.Errorf("Error reading ca_id: %v", err)
		}
	}

	return nil
}
