// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate ocsp.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateOcsp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateOcspRead,
		Update: resourceSystemCertificateOcspUpdate,
		Create: resourceSystemCertificateOcspCreate,
		Delete: resourceSystemCertificateOcspDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"reject_ocsp_response_with_missing_nextupdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"nonce_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"verify_others": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"accept_trusted_root_ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"caching_extra_max_age_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"caching_expire_ahead_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"caching": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ca_chain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"criteria_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_certificates": &schema.Schema{
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
		},
	}
}
func resourceSystemCertificateOcspCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOcsp: type error")
	}

	obj, err := getObjectSystemCertificateOcsp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcsp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateOcsp(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcsp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateOcspRead(d, m)
}
func resourceSystemCertificateOcspUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateOcsp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcsp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateOcsp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcsp resource: %v", err)
	}

	return resourceSystemCertificateOcspRead(d, m)
}
func resourceSystemCertificateOcspDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateOcsp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateOcsp resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateOcspRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateOcsp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcsp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOcsp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcsp resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateOcspRejectOcspResponseWithMissingNextupdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspNonceCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspTunnelUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspTunnelStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspVerifyOthers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspAcceptTrustedRootCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspCachingExtraMaxAgeCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspTunnelPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspCachingExpireAheadTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspTunnelPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspCaching(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspCaChain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspTunnelIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspCriteriaCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspRemoteCertificates(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOcsp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("reject_ocsp_response_with_missing_nextupdate", flattenSystemCertificateOcspRejectOcspResponseWithMissingNextupdate(o["reject-ocsp-response-with-missing-nextupdate"], d, "reject_ocsp_response_with_missing_nextupdate", sv)); err != nil {
		if !fortiAPIPatch(o["reject-ocsp-response-with-missing-nextupdate"]) {
			return fmt.Errorf("Error reading reject-ocsp-response-with-missing-nextupdate: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSystemCertificateOcspMaxAge(o["max_age"], d, "max_age", sv)); err != nil {
		if !fortiAPIPatch(o["max_age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("nonce_check", flattenSystemCertificateOcspNonceCheck(o["nonce_check"], d, "nonce_check", sv)); err != nil {
		if !fortiAPIPatch(o["nonce_check"]) {
			return fmt.Errorf("Error reading nonce_check: %v", err)
		}
	}

	if err = d.Set("tunnel_username", flattenSystemCertificateOcspTunnelUsername(o["tunnel_username"], d, "tunnel_username", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel_username"]) {
			return fmt.Errorf("Error reading tunnel_username: %v", err)
		}
	}

	if err = d.Set("url", flattenSystemCertificateOcspUrl(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("tunnel_status", flattenSystemCertificateOcspTunnelStatus(o["tunnel_status"], d, "tunnel_status", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel_status"]) {
			return fmt.Errorf("Error reading tunnel_status: %v", err)
		}
	}

	if err = d.Set("verify_others", flattenSystemCertificateOcspVerifyOthers(o["verify_others"], d, "verify_others", sv)); err != nil {
		if !fortiAPIPatch(o["verify_others"]) {
			return fmt.Errorf("Error reading verify_others: %v", err)
		}
	}

	if err = d.Set("accept_trusted_root_ca", flattenSystemCertificateOcspAcceptTrustedRootCa(o["accept_trusted_root_ca"], d, "accept_trusted_root_ca", sv)); err != nil {
		if !fortiAPIPatch(o["accept_trusted_root_ca"]) {
			return fmt.Errorf("Error reading accept_trusted_root_ca: %v", err)
		}
	}

	if err = d.Set("caching_extra_max_age_check", flattenSystemCertificateOcspCachingExtraMaxAgeCheck(o["caching_extra_max_age_check"], d, "caching_extra_max_age_check", sv)); err != nil {
		if !fortiAPIPatch(o["caching_extra_max_age_check"]) {
			return fmt.Errorf("Error reading caching_extra_max_age_check: %v", err)
		}
	}

	if err = d.Set("tunnel_password", flattenSystemCertificateOcspTunnelPassword(o["tunnel_password"], d, "tunnel_password", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel_password"]) {
			return fmt.Errorf("Error reading tunnel_password: %v", err)
		}
	}

	if err = d.Set("caching_expire_ahead_time", flattenSystemCertificateOcspCachingExpireAheadTime(o["caching_expire_ahead_time"], d, "caching_expire_ahead_time", sv)); err != nil {
		if !fortiAPIPatch(o["caching_expire_ahead_time"]) {
			return fmt.Errorf("Error reading caching_expire_ahead_time: %v", err)
		}
	}

	if err = d.Set("tunnel_port", flattenSystemCertificateOcspTunnelPort(o["tunnel_port"], d, "tunnel_port", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel_port"]) {
			return fmt.Errorf("Error reading tunnel_port: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemCertificateOcspTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("caching", flattenSystemCertificateOcspCaching(o["caching"], d, "caching", sv)); err != nil {
		if !fortiAPIPatch(o["caching"]) {
			return fmt.Errorf("Error reading caching: %v", err)
		}
	}

	if err = d.Set("ca_chain", flattenSystemCertificateOcspCaChain(o["ca_chain"], d, "ca_chain", sv)); err != nil {
		if !fortiAPIPatch(o["ca_chain"]) {
			return fmt.Errorf("Error reading ca_chain: %v", err)
		}
	}

	if err = d.Set("tunnel_ip", flattenSystemCertificateOcspTunnelIp(o["tunnel_ip"], d, "tunnel_ip", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel_ip"]) {
			return fmt.Errorf("Error reading tunnel_ip: %v", err)
		}
	}

	if err = d.Set("criteria_check", flattenSystemCertificateOcspCriteriaCheck(o["criteria_check"], d, "criteria_check", sv)); err != nil {
		if !fortiAPIPatch(o["criteria_check"]) {
			return fmt.Errorf("Error reading criteria_check: %v", err)
		}
	}

	if err = d.Set("remote_certificates", flattenSystemCertificateOcspRemoteCertificates(o["remote_certificates"], d, "remote_certificates", sv)); err != nil {
		if !fortiAPIPatch(o["remote_certificates"]) {
			return fmt.Errorf("Error reading remote_certificates: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemCertificateOcspMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateOcspRejectOcspResponseWithMissingNextupdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspNonceCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspTunnelUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspTunnelStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspVerifyOthers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspAcceptTrustedRootCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspCachingExtraMaxAgeCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspTunnelPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspCachingExpireAheadTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspTunnelPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspCaching(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspCaChain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspTunnelIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspCriteriaCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspRemoteCertificates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOcsp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("reject_ocsp_response_with_missing_nextupdate"); ok {
		t, err := expandSystemCertificateOcspRejectOcspResponseWithMissingNextupdate(d, v, "reject_ocsp_response_with_missing_nextupdate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reject-ocsp-response-with-missing-nextupdate"] = t
		}
	}

	if v, ok := d.GetOk("max_age"); ok {
		t, err := expandSystemCertificateOcspMaxAge(d, v, "max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_age"] = t
		}
	}

	if v, ok := d.GetOk("nonce_check"); ok {
		t, err := expandSystemCertificateOcspNonceCheck(d, v, "nonce_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nonce_check"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_username"); ok {
		t, err := expandSystemCertificateOcspTunnelUsername(d, v, "tunnel_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel_username"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {
		t, err := expandSystemCertificateOcspUrl(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_status"); ok {
		t, err := expandSystemCertificateOcspTunnelStatus(d, v, "tunnel_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel_status"] = t
		}
	}

	if v, ok := d.GetOk("verify_others"); ok {
		t, err := expandSystemCertificateOcspVerifyOthers(d, v, "verify_others", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify_others"] = t
		}
	}

	if v, ok := d.GetOk("accept_trusted_root_ca"); ok {
		t, err := expandSystemCertificateOcspAcceptTrustedRootCa(d, v, "accept_trusted_root_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept_trusted_root_ca"] = t
		}
	}

	if v, ok := d.GetOk("caching_extra_max_age_check"); ok {
		t, err := expandSystemCertificateOcspCachingExtraMaxAgeCheck(d, v, "caching_extra_max_age_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caching_extra_max_age_check"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_password"); ok {
		t, err := expandSystemCertificateOcspTunnelPassword(d, v, "tunnel_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel_password"] = t
		}
	}

	if v, ok := d.GetOk("caching_expire_ahead_time"); ok {
		t, err := expandSystemCertificateOcspCachingExpireAheadTime(d, v, "caching_expire_ahead_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caching_expire_ahead_time"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_port"); ok {
		t, err := expandSystemCertificateOcspTunnelPort(d, v, "tunnel_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel_port"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandSystemCertificateOcspTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("caching"); ok {
		t, err := expandSystemCertificateOcspCaching(d, v, "caching", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caching"] = t
		}
	}

	if v, ok := d.GetOk("ca_chain"); ok {
		t, err := expandSystemCertificateOcspCaChain(d, v, "ca_chain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca_chain"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_ip"); ok {
		t, err := expandSystemCertificateOcspTunnelIp(d, v, "tunnel_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel_ip"] = t
		}
	}

	if v, ok := d.GetOk("criteria_check"); ok {
		t, err := expandSystemCertificateOcspCriteriaCheck(d, v, "criteria_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["criteria_check"] = t
		}
	}

	if v, ok := d.GetOk("remote_certificates"); ok {
		t, err := expandSystemCertificateOcspRemoteCertificates(d, v, "remote_certificates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote_certificates"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateOcspMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
