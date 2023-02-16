// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate ocsp.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemCertificateOcsp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateOcspRead,
		Schema: map[string]*schema.Schema{
			"reject_ocsp_response_with_missing_nextupdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ca_chain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"criteria_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nonce_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"verify_others": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"accept_trusted_root_ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"caching_extra_max_age_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"caching_expire_ahead_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"caching": &schema.Schema{
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
func dataSourceSystemCertificateOcspRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("verify_others")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemCertificateOcsp: type error")
	}

	o, err := c.ReadSystemCertificateOcsp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateOcsp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemCertificateOcsp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCertificateOcsp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemCertificateOcspRejectOcspResponseWithMissingNextupdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspTunnelUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspCaChain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspCriteriaCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspRemoteCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspTunnelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspTunnelPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspTunnelIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspNonceCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspVerifyOthers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspAcceptTrustedRootCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspCachingExtraMaxAgeCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspCachingExpireAheadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspTunnelPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCertificateOcspCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCertificateOcsp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("reject_ocsp_response_with_missing_nextupdate", dataSourceFlattenSystemCertificateOcspRejectOcspResponseWithMissingNextupdate(o["reject-ocsp-response-with-missing-nextupdate"], d, "reject_ocsp_response_with_missing_nextupdate")); err != nil {
		if !fortiAPIPatch(o["reject-ocsp-response-with-missing-nextupdate"]) {
			return fmt.Errorf("Error reading reject-ocsp-response-with-missing-nextupdate: %v", err)
		}
	}

	if err = d.Set("tunnel_username", dataSourceFlattenSystemCertificateOcspTunnelUsername(o["tunnel_username"], d, "tunnel_username")); err != nil {
		if !fortiAPIPatch(o["tunnel_username"]) {
			return fmt.Errorf("Error reading tunnel_username: %v", err)
		}
	}

	if err = d.Set("ca_chain", dataSourceFlattenSystemCertificateOcspCaChain(o["ca_chain"], d, "ca_chain")); err != nil {
		if !fortiAPIPatch(o["ca_chain"]) {
			return fmt.Errorf("Error reading ca_chain: %v", err)
		}
	}

	if err = d.Set("criteria_check", dataSourceFlattenSystemCertificateOcspCriteriaCheck(o["criteria_check"], d, "criteria_check")); err != nil {
		if !fortiAPIPatch(o["criteria_check"]) {
			return fmt.Errorf("Error reading criteria_check: %v", err)
		}
	}

	if err = d.Set("remote_certificates", dataSourceFlattenSystemCertificateOcspRemoteCertificates(o["remote_certificates"], d, "remote_certificates")); err != nil {
		if !fortiAPIPatch(o["remote_certificates"]) {
			return fmt.Errorf("Error reading remote_certificates: %v", err)
		}
	}

	if err = d.Set("tunnel_status", dataSourceFlattenSystemCertificateOcspTunnelStatus(o["tunnel_status"], d, "tunnel_status")); err != nil {
		if !fortiAPIPatch(o["tunnel_status"]) {
			return fmt.Errorf("Error reading tunnel_status: %v", err)
		}
	}

	if err = d.Set("max_age", dataSourceFlattenSystemCertificateOcspMaxAge(o["max_age"], d, "max_age")); err != nil {
		if !fortiAPIPatch(o["max_age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("tunnel_port", dataSourceFlattenSystemCertificateOcspTunnelPort(o["tunnel_port"], d, "tunnel_port")); err != nil {
		if !fortiAPIPatch(o["tunnel_port"]) {
			return fmt.Errorf("Error reading tunnel_port: %v", err)
		}
	}

	if err = d.Set("tunnel_ip", dataSourceFlattenSystemCertificateOcspTunnelIp(o["tunnel_ip"], d, "tunnel_ip")); err != nil {
		if !fortiAPIPatch(o["tunnel_ip"]) {
			return fmt.Errorf("Error reading tunnel_ip: %v", err)
		}
	}

	if err = d.Set("nonce_check", dataSourceFlattenSystemCertificateOcspNonceCheck(o["nonce_check"], d, "nonce_check")); err != nil {
		if !fortiAPIPatch(o["nonce_check"]) {
			return fmt.Errorf("Error reading nonce_check: %v", err)
		}
	}

	if err = d.Set("verify_others", dataSourceFlattenSystemCertificateOcspVerifyOthers(o["verify_others"], d, "verify_others")); err != nil {
		if !fortiAPIPatch(o["verify_others"]) {
			return fmt.Errorf("Error reading verify_others: %v", err)
		}
	}

	if err = d.Set("accept_trusted_root_ca", dataSourceFlattenSystemCertificateOcspAcceptTrustedRootCa(o["accept_trusted_root_ca"], d, "accept_trusted_root_ca")); err != nil {
		if !fortiAPIPatch(o["accept_trusted_root_ca"]) {
			return fmt.Errorf("Error reading accept_trusted_root_ca: %v", err)
		}
	}

	if err = d.Set("caching_extra_max_age_check", dataSourceFlattenSystemCertificateOcspCachingExtraMaxAgeCheck(o["caching_extra_max_age_check"], d, "caching_extra_max_age_check")); err != nil {
		if !fortiAPIPatch(o["caching_extra_max_age_check"]) {
			return fmt.Errorf("Error reading caching_extra_max_age_check: %v", err)
		}
	}

	if err = d.Set("caching_expire_ahead_time", dataSourceFlattenSystemCertificateOcspCachingExpireAheadTime(o["caching_expire_ahead_time"], d, "caching_expire_ahead_time")); err != nil {
		if !fortiAPIPatch(o["caching_expire_ahead_time"]) {
			return fmt.Errorf("Error reading caching_expire_ahead_time: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemCertificateOcspMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("tunnel_password", dataSourceFlattenSystemCertificateOcspTunnelPassword(o["tunnel_password"], d, "tunnel_password")); err != nil {
		if !fortiAPIPatch(o["tunnel_password"]) {
			return fmt.Errorf("Error reading tunnel_password: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenSystemCertificateOcspTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("caching", dataSourceFlattenSystemCertificateOcspCaching(o["caching"], d, "caching")); err != nil {
		if !fortiAPIPatch(o["caching"]) {
			return fmt.Errorf("Error reading caching: %v", err)
		}
	}

	return nil
}
