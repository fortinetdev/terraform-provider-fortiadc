// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance real server ssl profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceRealServerSslProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceRealServerSslProfileRead,
		Update: resourceLoadBalanceRealServerSslProfileUpdate,
		Create: resourceLoadBalanceRealServerSslProfileCreate,
		Delete: resourceLoadBalanceRealServerSslProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"new_ssl_ciphers_long": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_ocsp_stapling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secure_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"supported_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"renegotiation_deny_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"renegotiate_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sni_forward_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cert_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rfc7919_comply": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"renegotiate_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ciphers_tlsv13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_ticket_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_reuse_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_reuse_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceLoadBalanceRealServerSslProfileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceRealServerSslProfile: type error")
	}

	obj, err := getObjectLoadBalanceRealServerSslProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceRealServerSslProfile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceRealServerSslProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceRealServerSslProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceRealServerSslProfileRead(d, m)
}
func resourceLoadBalanceRealServerSslProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceRealServerSslProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceRealServerSslProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceRealServerSslProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceRealServerSslProfile resource: %v", err)
	}

	return resourceLoadBalanceRealServerSslProfileRead(d, m)
}
func resourceLoadBalanceRealServerSslProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceRealServerSslProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceRealServerSslProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceRealServerSslProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceRealServerSslProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceRealServerSslProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceRealServerSslProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceRealServerSslProfile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceRealServerSslProfileNewSslCiphersLong(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileLocalCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileServerOcspStapling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSecureRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSupportedGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileRenegotiationDenyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileRenegotiatePeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileAllowSslVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSniForwardFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileCertVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileRfc7919Comply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileRenegotiateSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileCiphersTlsv13(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileTlsTicketFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSessionReuseLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceRealServerSslProfileSessionReuseFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceRealServerSslProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("new_ssl_ciphers_long", flattenLoadBalanceRealServerSslProfileNewSslCiphersLong(o["new_ssl_ciphers_long"], d, "new_ssl_ciphers_long", sv)); err != nil {
		if !fortiAPIPatch(o["new_ssl_ciphers_long"]) {
			return fmt.Errorf("Error reading new_ssl_ciphers_long: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenLoadBalanceRealServerSslProfileLocalCert(o["local_cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local_cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("server_ocsp_stapling", flattenLoadBalanceRealServerSslProfileServerOcspStapling(o["server_OCSP_stapling"], d, "server_ocsp_stapling", sv)); err != nil {
		if !fortiAPIPatch(o["server_OCSP_stapling"]) {
			return fmt.Errorf("Error reading server_OCSP_stapling: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers_flag", flattenLoadBalanceRealServerSslProfileCustomizedSslCiphersFlag(o["customized_ssl_ciphers_flag"], d, "customized_ssl_ciphers_flag", sv)); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("secure_renegotiation", flattenLoadBalanceRealServerSslProfileSecureRenegotiation(o["secure_renegotiation"], d, "secure_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["secure_renegotiation"]) {
			return fmt.Errorf("Error reading secure_renegotiation: %v", err)
		}
	}

	if err = d.Set("supported_groups", flattenLoadBalanceRealServerSslProfileSupportedGroups(o["supported_groups"], d, "supported_groups", sv)); err != nil {
		if !fortiAPIPatch(o["supported_groups"]) {
			return fmt.Errorf("Error reading supported_groups: %v", err)
		}
	}

	if err = d.Set("renegotiation", flattenLoadBalanceRealServerSslProfileRenegotiation(o["renegotiation"], d, "renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["renegotiation"]) {
			return fmt.Errorf("Error reading renegotiation: %v", err)
		}
	}

	if err = d.Set("renegotiation_deny_action", flattenLoadBalanceRealServerSslProfileRenegotiationDenyAction(o["renegotiation_deny_action"], d, "renegotiation_deny_action", sv)); err != nil {
		if !fortiAPIPatch(o["renegotiation_deny_action"]) {
			return fmt.Errorf("Error reading renegotiation_deny_action: %v", err)
		}
	}

	if err = d.Set("renegotiate_period", flattenLoadBalanceRealServerSslProfileRenegotiatePeriod(o["renegotiate_period"], d, "renegotiate_period", sv)); err != nil {
		if !fortiAPIPatch(o["renegotiate_period"]) {
			return fmt.Errorf("Error reading renegotiate_period: %v", err)
		}
	}

	if err = d.Set("allow_ssl_versions", flattenLoadBalanceRealServerSslProfileAllowSslVersions(o["allow_ssl_versions"], d, "allow_ssl_versions", sv)); err != nil {
		if !fortiAPIPatch(o["allow_ssl_versions"]) {
			return fmt.Errorf("Error reading allow_ssl_versions: %v", err)
		}
	}

	if err = d.Set("sni_forward_flag", flattenLoadBalanceRealServerSslProfileSniForwardFlag(o["sni_forward_flag"], d, "sni_forward_flag", sv)); err != nil {
		if !fortiAPIPatch(o["sni_forward_flag"]) {
			return fmt.Errorf("Error reading sni_forward_flag: %v", err)
		}
	}

	if err = d.Set("ssl", flattenLoadBalanceRealServerSslProfileSsl(o["ssl"], d, "ssl", sv)); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("sni", flattenLoadBalanceRealServerSslProfileSni(o["sni"], d, "sni", sv)); err != nil {
		if !fortiAPIPatch(o["sni"]) {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceRealServerSslProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("cert_verify", flattenLoadBalanceRealServerSslProfileCertVerify(o["cert_verify"], d, "cert_verify", sv)); err != nil {
		if !fortiAPIPatch(o["cert_verify"]) {
			return fmt.Errorf("Error reading cert_verify: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers", flattenLoadBalanceRealServerSslProfileCustomizedSslCiphers(o["customized_ssl_ciphers"], d, "customized_ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("rfc7919_comply", flattenLoadBalanceRealServerSslProfileRfc7919Comply(o["rfc7919_comply"], d, "rfc7919_comply", sv)); err != nil {
		if !fortiAPIPatch(o["rfc7919_comply"]) {
			return fmt.Errorf("Error reading rfc7919_comply: %v", err)
		}
	}

	if err = d.Set("renegotiate_size", flattenLoadBalanceRealServerSslProfileRenegotiateSize(o["renegotiate_size"], d, "renegotiate_size", sv)); err != nil {
		if !fortiAPIPatch(o["renegotiate_size"]) {
			return fmt.Errorf("Error reading renegotiate_size: %v", err)
		}
	}

	if err = d.Set("ciphers_tlsv13", flattenLoadBalanceRealServerSslProfileCiphersTlsv13(o["ciphers_tlsv13"], d, "ciphers_tlsv13", sv)); err != nil {
		if !fortiAPIPatch(o["ciphers_tlsv13"]) {
			return fmt.Errorf("Error reading ciphers_tlsv13: %v", err)
		}
	}

	if err = d.Set("tls_ticket_flag", flattenLoadBalanceRealServerSslProfileTlsTicketFlag(o["tls_ticket_flag"], d, "tls_ticket_flag", sv)); err != nil {
		if !fortiAPIPatch(o["tls_ticket_flag"]) {
			return fmt.Errorf("Error reading tls_ticket_flag: %v", err)
		}
	}

	if err = d.Set("session_reuse_limit", flattenLoadBalanceRealServerSslProfileSessionReuseLimit(o["session_reuse_limit"], d, "session_reuse_limit", sv)); err != nil {
		if !fortiAPIPatch(o["session_reuse_limit"]) {
			return fmt.Errorf("Error reading session_reuse_limit: %v", err)
		}
	}

	if err = d.Set("session_reuse_flag", flattenLoadBalanceRealServerSslProfileSessionReuseFlag(o["session_reuse_flag"], d, "session_reuse_flag", sv)); err != nil {
		if !fortiAPIPatch(o["session_reuse_flag"]) {
			return fmt.Errorf("Error reading session_reuse_flag: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceRealServerSslProfileNewSslCiphersLong(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileLocalCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileServerOcspStapling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileCustomizedSslCiphersFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSecureRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSupportedGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileRenegotiationDenyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileRenegotiatePeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileAllowSslVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSniForwardFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileCertVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileCustomizedSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileRfc7919Comply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileRenegotiateSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileCiphersTlsv13(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileTlsTicketFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSessionReuseLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceRealServerSslProfileSessionReuseFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceRealServerSslProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("new_ssl_ciphers_long"); ok {
		t, err := expandLoadBalanceRealServerSslProfileNewSslCiphersLong(d, v, "new_ssl_ciphers_long", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["new_ssl_ciphers_long"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandLoadBalanceRealServerSslProfileLocalCert(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local_cert"] = t
		}
	}

	if v, ok := d.GetOk("server_ocsp_stapling"); ok {
		t, err := expandLoadBalanceRealServerSslProfileServerOcspStapling(d, v, "server_ocsp_stapling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_OCSP_stapling"] = t
		}
	}

	if v, ok := d.GetOk("customized_ssl_ciphers_flag"); ok {
		t, err := expandLoadBalanceRealServerSslProfileCustomizedSslCiphersFlag(d, v, "customized_ssl_ciphers_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customized_ssl_ciphers_flag"] = t
		}
	}

	if v, ok := d.GetOk("secure_renegotiation"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSecureRenegotiation(d, v, "secure_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure_renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("supported_groups"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSupportedGroups(d, v, "supported_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supported_groups"] = t
		}
	}

	if v, ok := d.GetOk("renegotiation"); ok {
		t, err := expandLoadBalanceRealServerSslProfileRenegotiation(d, v, "renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("renegotiation_deny_action"); ok {
		t, err := expandLoadBalanceRealServerSslProfileRenegotiationDenyAction(d, v, "renegotiation_deny_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renegotiation_deny_action"] = t
		}
	}

	if v, ok := d.GetOk("renegotiate_period"); ok {
		t, err := expandLoadBalanceRealServerSslProfileRenegotiatePeriod(d, v, "renegotiate_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renegotiate_period"] = t
		}
	}

	if v, ok := d.GetOk("allow_ssl_versions"); ok {
		t, err := expandLoadBalanceRealServerSslProfileAllowSslVersions(d, v, "allow_ssl_versions", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow_ssl_versions"] = t
		}
	}

	if v, ok := d.GetOk("sni_forward_flag"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSniForwardFlag(d, v, "sni_forward_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni_forward_flag"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("sni"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSni(d, v, "sni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceRealServerSslProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("cert_verify"); ok {
		t, err := expandLoadBalanceRealServerSslProfileCertVerify(d, v, "cert_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert_verify"] = t
		}
	}

	if v, ok := d.GetOk("customized_ssl_ciphers"); ok {
		t, err := expandLoadBalanceRealServerSslProfileCustomizedSslCiphers(d, v, "customized_ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customized_ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("rfc7919_comply"); ok {
		t, err := expandLoadBalanceRealServerSslProfileRfc7919Comply(d, v, "rfc7919_comply", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc7919_comply"] = t
		}
	}

	if v, ok := d.GetOk("renegotiate_size"); ok {
		t, err := expandLoadBalanceRealServerSslProfileRenegotiateSize(d, v, "renegotiate_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renegotiate_size"] = t
		}
	}

	if v, ok := d.GetOk("ciphers_tlsv13"); ok {
		t, err := expandLoadBalanceRealServerSslProfileCiphersTlsv13(d, v, "ciphers_tlsv13", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ciphers_tlsv13"] = t
		}
	}

	if v, ok := d.GetOk("tls_ticket_flag"); ok {
		t, err := expandLoadBalanceRealServerSslProfileTlsTicketFlag(d, v, "tls_ticket_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls_ticket_flag"] = t
		}
	}

	if v, ok := d.GetOk("session_reuse_limit"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSessionReuseLimit(d, v, "session_reuse_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session_reuse_limit"] = t
		}
	}

	if v, ok := d.GetOk("session_reuse_flag"); ok {
		t, err := expandLoadBalanceRealServerSslProfileSessionReuseFlag(d, v, "session_reuse_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session_reuse_flag"] = t
		}
	}

	return &obj, nil
}
