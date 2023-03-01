// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance client ssl profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceClientSslProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceClientSslProfileRead,
		Update: resourceLoadBalanceClientSslProfileUpdate,
		Create: resourceLoadBalanceClientSslProfileCreate,
		Delete: resourceLoadBalanceClientSslProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ssl_allowed_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_renegotiate_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_certificate_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_secure_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_forward_client_certificate_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"supported_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_forward_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_tls_tickets": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_sni_required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_ssl_sni_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_verify_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_certificate_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_ssl_allowed_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_dh_param_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_ciphers_tlsv13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_ciphers_tlsv13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_renegotiate_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward_proxy_certificate_caching": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"forward_proxy_local_signing_ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reject_ocsp_stapling_with_missing_nextupdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_dynamic_record_sizing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rfc7919_comply": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_ssl_ocsp_stapling_support": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_session_cache_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_renegotiation_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward_proxy_intermediate_ca_group": &schema.Schema{
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
func resourceLoadBalanceClientSslProfileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceClientSslProfile: type error")
	}

	obj, err := getObjectLoadBalanceClientSslProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceClientSslProfile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceClientSslProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceClientSslProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceClientSslProfileRead(d, m)
}
func resourceLoadBalanceClientSslProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceClientSslProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClientSslProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceClientSslProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClientSslProfile resource: %v", err)
	}

	return resourceLoadBalanceClientSslProfileRead(d, m)
}
func resourceLoadBalanceClientSslProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceClientSslProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceClientSslProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceClientSslProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceClientSslProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceClientSslProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceClientSslProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceClientSslProfile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceClientSslProfileSslAllowedVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslRenegotiateSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendCertificateVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslSecureRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileHttpForwardClientCertificateHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSupportedGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileHttpForwardClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileUseTlsTickets(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileClientSniRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileClientCertificateVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendSslSniForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileClientCertificateVerifyMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileLocalCertificateGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendSslAllowedVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslDhParamSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslCiphersTlsv13(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendCiphersTlsv13(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslRenegotiatePeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileForwardProxyCertificateCaching(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileForwardProxyLocalSigningCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileRejectOcspStaplingWithMissingNextupdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslDynamicRecordSizing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileRfc7919Comply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileBackendSslOcspStaplingSupport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslSessionCacheFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileForwardProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileSslRenegotiationInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceClientSslProfileForwardProxyIntermediateCaGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceClientSslProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ssl_allowed_versions", flattenLoadBalanceClientSslProfileSslAllowedVersions(o["ssl-allowed_versions"], d, "ssl_allowed_versions", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-allowed_versions"]) {
			return fmt.Errorf("Error reading ssl-allowed_versions: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiate_size", flattenLoadBalanceClientSslProfileSslRenegotiateSize(o["ssl_renegotiate_size"], d, "ssl_renegotiate_size", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiate_size"]) {
			return fmt.Errorf("Error reading ssl_renegotiate_size: %v", err)
		}
	}

	if err = d.Set("backend_certificate_verify", flattenLoadBalanceClientSslProfileBackendCertificateVerify(o["backend_certificate_verify"], d, "backend_certificate_verify", sv)); err != nil {
		if !fortiAPIPatch(o["backend_certificate_verify"]) {
			return fmt.Errorf("Error reading backend_certificate_verify: %v", err)
		}
	}

	if err = d.Set("backend_customized_ssl_ciphers", flattenLoadBalanceClientSslProfileBackendCustomizedSslCiphers(o["backend_customized_ssl_ciphers"], d, "backend_customized_ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["backend_customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading backend_customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation", flattenLoadBalanceClientSslProfileSslRenegotiation(o["ssl_renegotiation"], d, "ssl_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiation"]) {
			return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers_flag", flattenLoadBalanceClientSslProfileCustomizedSslCiphersFlag(o["customized_ssl_ciphers_flag"], d, "customized_ssl_ciphers_flag", sv)); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("backend_customized_ssl_ciphers_flag", flattenLoadBalanceClientSslProfileBackendCustomizedSslCiphersFlag(o["backend_customized_ssl_ciphers_flag"], d, "backend_customized_ssl_ciphers_flag", sv)); err != nil {
		if !fortiAPIPatch(o["backend_customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading backend_customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("ssl_secure_renegotiation", flattenLoadBalanceClientSslProfileSslSecureRenegotiation(o["ssl_secure_renegotiation"], d, "ssl_secure_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_secure_renegotiation"]) {
			return fmt.Errorf("Error reading ssl_secure_renegotiation: %v", err)
		}
	}

	if err = d.Set("http_forward_client_certificate_header", flattenLoadBalanceClientSslProfileHttpForwardClientCertificateHeader(o["http-forward_client_certificate_header"], d, "http_forward_client_certificate_header", sv)); err != nil {
		if !fortiAPIPatch(o["http-forward_client_certificate_header"]) {
			return fmt.Errorf("Error reading http-forward_client_certificate_header: %v", err)
		}
	}

	if err = d.Set("supported_groups", flattenLoadBalanceClientSslProfileSupportedGroups(o["supported_groups"], d, "supported_groups", sv)); err != nil {
		if !fortiAPIPatch(o["supported_groups"]) {
			return fmt.Errorf("Error reading supported_groups: %v", err)
		}
	}

	if err = d.Set("http_forward_client_certificate", flattenLoadBalanceClientSslProfileHttpForwardClientCertificate(o["http-forward_client_certificate"], d, "http_forward_client_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["http-forward_client_certificate"]) {
			return fmt.Errorf("Error reading http-forward_client_certificate: %v", err)
		}
	}

	if err = d.Set("use_tls_tickets", flattenLoadBalanceClientSslProfileUseTlsTickets(o["use_tls_tickets"], d, "use_tls_tickets", sv)); err != nil {
		if !fortiAPIPatch(o["use_tls_tickets"]) {
			return fmt.Errorf("Error reading use_tls_tickets: %v", err)
		}
	}

	if err = d.Set("client_sni_required", flattenLoadBalanceClientSslProfileClientSniRequired(o["client_sni_required"], d, "client_sni_required", sv)); err != nil {
		if !fortiAPIPatch(o["client_sni_required"]) {
			return fmt.Errorf("Error reading client_sni_required: %v", err)
		}
	}

	if err = d.Set("backend_ssl_ciphers", flattenLoadBalanceClientSslProfileBackendSslCiphers(o["backend_ssl_ciphers"], d, "backend_ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["backend_ssl_ciphers"]) {
			return fmt.Errorf("Error reading backend_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("client_certificate_verify", flattenLoadBalanceClientSslProfileClientCertificateVerify(o["client_certificate_verify"], d, "client_certificate_verify", sv)); err != nil {
		if !fortiAPIPatch(o["client_certificate_verify"]) {
			return fmt.Errorf("Error reading client_certificate_verify: %v", err)
		}
	}

	if err = d.Set("backend_ssl_sni_forward", flattenLoadBalanceClientSslProfileBackendSslSniForward(o["backend_ssl_sni_forward"], d, "backend_ssl_sni_forward", sv)); err != nil {
		if !fortiAPIPatch(o["backend_ssl_sni_forward"]) {
			return fmt.Errorf("Error reading backend_ssl_sni_forward: %v", err)
		}
	}

	if err = d.Set("client_certificate_verify_mode", flattenLoadBalanceClientSslProfileClientCertificateVerifyMode(o["client_certificate_verify_mode"], d, "client_certificate_verify_mode", sv)); err != nil {
		if !fortiAPIPatch(o["client_certificate_verify_mode"]) {
			return fmt.Errorf("Error reading client_certificate_verify_mode: %v", err)
		}
	}

	if err = d.Set("local_certificate_group", flattenLoadBalanceClientSslProfileLocalCertificateGroup(o["local_certificate_group"], d, "local_certificate_group", sv)); err != nil {
		if !fortiAPIPatch(o["local_certificate_group"]) {
			return fmt.Errorf("Error reading local_certificate_group: %v", err)
		}
	}

	if err = d.Set("backend_ssl_allowed_versions", flattenLoadBalanceClientSslProfileBackendSslAllowedVersions(o["backend_ssl_allowed_versions"], d, "backend_ssl_allowed_versions", sv)); err != nil {
		if !fortiAPIPatch(o["backend_ssl_allowed_versions"]) {
			return fmt.Errorf("Error reading backend_ssl_allowed_versions: %v", err)
		}
	}

	if err = d.Set("ssl_dh_param_size", flattenLoadBalanceClientSslProfileSslDhParamSize(o["ssl_dh_param_size"], d, "ssl_dh_param_size", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_dh_param_size"]) {
			return fmt.Errorf("Error reading ssl_dh_param_size: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers_tlsv13", flattenLoadBalanceClientSslProfileSslCiphersTlsv13(o["ssl_ciphers_tlsv13"], d, "ssl_ciphers_tlsv13", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_ciphers_tlsv13"]) {
			return fmt.Errorf("Error reading ssl_ciphers_tlsv13: %v", err)
		}
	}

	if err = d.Set("backend_ciphers_tlsv13", flattenLoadBalanceClientSslProfileBackendCiphersTlsv13(o["backend_ciphers_tlsv13"], d, "backend_ciphers_tlsv13", sv)); err != nil {
		if !fortiAPIPatch(o["backend_ciphers_tlsv13"]) {
			return fmt.Errorf("Error reading backend_ciphers_tlsv13: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiate_period", flattenLoadBalanceClientSslProfileSslRenegotiatePeriod(o["ssl_renegotiate_period"], d, "ssl_renegotiate_period", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiate_period"]) {
			return fmt.Errorf("Error reading ssl_renegotiate_period: %v", err)
		}
	}

	if err = d.Set("forward_proxy_certificate_caching", flattenLoadBalanceClientSslProfileForwardProxyCertificateCaching(o["forward_proxy_certificate_caching"], d, "forward_proxy_certificate_caching", sv)); err != nil {
		if !fortiAPIPatch(o["forward_proxy_certificate_caching"]) {
			return fmt.Errorf("Error reading forward_proxy_certificate_caching: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceClientSslProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("forward_proxy_local_signing_ca", flattenLoadBalanceClientSslProfileForwardProxyLocalSigningCa(o["forward_proxy_local_signing_ca"], d, "forward_proxy_local_signing_ca", sv)); err != nil {
		if !fortiAPIPatch(o["forward_proxy_local_signing_ca"]) {
			return fmt.Errorf("Error reading forward_proxy_local_signing_ca: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers", flattenLoadBalanceClientSslProfileCustomizedSslCiphers(o["customized_ssl_ciphers"], d, "customized_ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers", flattenLoadBalanceClientSslProfileSslCiphers(o["ssl_ciphers"], d, "ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_ciphers"]) {
			return fmt.Errorf("Error reading ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("reject_ocsp_stapling_with_missing_nextupdate", flattenLoadBalanceClientSslProfileRejectOcspStaplingWithMissingNextupdate(o["reject-ocsp-stapling-with-missing-nextupdate"], d, "reject_ocsp_stapling_with_missing_nextupdate", sv)); err != nil {
		if !fortiAPIPatch(o["reject-ocsp-stapling-with-missing-nextupdate"]) {
			return fmt.Errorf("Error reading reject-ocsp-stapling-with-missing-nextupdate: %v", err)
		}
	}

	if err = d.Set("ssl_dynamic_record_sizing", flattenLoadBalanceClientSslProfileSslDynamicRecordSizing(o["ssl_dynamic_record_sizing"], d, "ssl_dynamic_record_sizing", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_dynamic_record_sizing"]) {
			return fmt.Errorf("Error reading ssl_dynamic_record_sizing: %v", err)
		}
	}

	if err = d.Set("rfc7919_comply", flattenLoadBalanceClientSslProfileRfc7919Comply(o["rfc7919_comply"], d, "rfc7919_comply", sv)); err != nil {
		if !fortiAPIPatch(o["rfc7919_comply"]) {
			return fmt.Errorf("Error reading rfc7919_comply: %v", err)
		}
	}

	if err = d.Set("backend_ssl_ocsp_stapling_support", flattenLoadBalanceClientSslProfileBackendSslOcspStaplingSupport(o["backend_ssl_OCSP_stapling_support"], d, "backend_ssl_ocsp_stapling_support", sv)); err != nil {
		if !fortiAPIPatch(o["backend_ssl_OCSP_stapling_support"]) {
			return fmt.Errorf("Error reading backend_ssl_OCSP_stapling_support: %v", err)
		}
	}

	if err = d.Set("ssl_session_cache_flag", flattenLoadBalanceClientSslProfileSslSessionCacheFlag(o["ssl_session_cache_flag"], d, "ssl_session_cache_flag", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_session_cache_flag"]) {
			return fmt.Errorf("Error reading ssl_session_cache_flag: %v", err)
		}
	}

	if err = d.Set("forward_proxy", flattenLoadBalanceClientSslProfileForwardProxy(o["forward_proxy"], d, "forward_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["forward_proxy"]) {
			return fmt.Errorf("Error reading forward_proxy: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation_interval", flattenLoadBalanceClientSslProfileSslRenegotiationInterval(o["ssl_renegotiation_interval"], d, "ssl_renegotiation_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiation_interval"]) {
			return fmt.Errorf("Error reading ssl_renegotiation_interval: %v", err)
		}
	}

	if err = d.Set("forward_proxy_intermediate_ca_group", flattenLoadBalanceClientSslProfileForwardProxyIntermediateCaGroup(o["forward_proxy_intermediate_ca_group"], d, "forward_proxy_intermediate_ca_group", sv)); err != nil {
		if !fortiAPIPatch(o["forward_proxy_intermediate_ca_group"]) {
			return fmt.Errorf("Error reading forward_proxy_intermediate_ca_group: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceClientSslProfileSslAllowedVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslRenegotiateSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendCertificateVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendCustomizedSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileCustomizedSslCiphersFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendCustomizedSslCiphersFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslSecureRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileHttpForwardClientCertificateHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSupportedGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileHttpForwardClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileUseTlsTickets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileClientSniRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileClientCertificateVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendSslSniForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileClientCertificateVerifyMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileLocalCertificateGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendSslAllowedVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslDhParamSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslCiphersTlsv13(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendCiphersTlsv13(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslRenegotiatePeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileForwardProxyCertificateCaching(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileForwardProxyLocalSigningCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileCustomizedSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileRejectOcspStaplingWithMissingNextupdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslDynamicRecordSizing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileRfc7919Comply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileBackendSslOcspStaplingSupport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslSessionCacheFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileForwardProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileSslRenegotiationInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceClientSslProfileForwardProxyIntermediateCaGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceClientSslProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ssl_allowed_versions"); ok {
		t, err := expandLoadBalanceClientSslProfileSslAllowedVersions(d, v, "ssl_allowed_versions", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-allowed_versions"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiate_size"); ok {
		t, err := expandLoadBalanceClientSslProfileSslRenegotiateSize(d, v, "ssl_renegotiate_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_renegotiate_size"] = t
		}
	}

	if v, ok := d.GetOk("backend_certificate_verify"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendCertificateVerify(d, v, "backend_certificate_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_certificate_verify"] = t
		}
	}

	if v, ok := d.GetOk("backend_customized_ssl_ciphers"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendCustomizedSslCiphers(d, v, "backend_customized_ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_customized_ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiation"); ok {
		t, err := expandLoadBalanceClientSslProfileSslRenegotiation(d, v, "ssl_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("customized_ssl_ciphers_flag"); ok {
		t, err := expandLoadBalanceClientSslProfileCustomizedSslCiphersFlag(d, v, "customized_ssl_ciphers_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customized_ssl_ciphers_flag"] = t
		}
	}

	if v, ok := d.GetOk("backend_customized_ssl_ciphers_flag"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendCustomizedSslCiphersFlag(d, v, "backend_customized_ssl_ciphers_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_customized_ssl_ciphers_flag"] = t
		}
	}

	if v, ok := d.GetOk("ssl_secure_renegotiation"); ok {
		t, err := expandLoadBalanceClientSslProfileSslSecureRenegotiation(d, v, "ssl_secure_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_secure_renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("http_forward_client_certificate_header"); ok {
		t, err := expandLoadBalanceClientSslProfileHttpForwardClientCertificateHeader(d, v, "http_forward_client_certificate_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-forward_client_certificate_header"] = t
		}
	}

	if v, ok := d.GetOk("supported_groups"); ok {
		t, err := expandLoadBalanceClientSslProfileSupportedGroups(d, v, "supported_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supported_groups"] = t
		}
	}

	if v, ok := d.GetOk("http_forward_client_certificate"); ok {
		t, err := expandLoadBalanceClientSslProfileHttpForwardClientCertificate(d, v, "http_forward_client_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-forward_client_certificate"] = t
		}
	}

	if v, ok := d.GetOk("use_tls_tickets"); ok {
		t, err := expandLoadBalanceClientSslProfileUseTlsTickets(d, v, "use_tls_tickets", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use_tls_tickets"] = t
		}
	}

	if v, ok := d.GetOk("client_sni_required"); ok {
		t, err := expandLoadBalanceClientSslProfileClientSniRequired(d, v, "client_sni_required", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_sni_required"] = t
		}
	}

	if v, ok := d.GetOk("backend_ssl_ciphers"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendSslCiphers(d, v, "backend_ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("client_certificate_verify"); ok {
		t, err := expandLoadBalanceClientSslProfileClientCertificateVerify(d, v, "client_certificate_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_certificate_verify"] = t
		}
	}

	if v, ok := d.GetOk("backend_ssl_sni_forward"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendSslSniForward(d, v, "backend_ssl_sni_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_ssl_sni_forward"] = t
		}
	}

	if v, ok := d.GetOk("client_certificate_verify_mode"); ok {
		t, err := expandLoadBalanceClientSslProfileClientCertificateVerifyMode(d, v, "client_certificate_verify_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_certificate_verify_mode"] = t
		}
	}

	if v, ok := d.GetOk("local_certificate_group"); ok {
		t, err := expandLoadBalanceClientSslProfileLocalCertificateGroup(d, v, "local_certificate_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local_certificate_group"] = t
		}
	}

	if v, ok := d.GetOk("backend_ssl_allowed_versions"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendSslAllowedVersions(d, v, "backend_ssl_allowed_versions", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_ssl_allowed_versions"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_param_size"); ok {
		t, err := expandLoadBalanceClientSslProfileSslDhParamSize(d, v, "ssl_dh_param_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_dh_param_size"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ciphers_tlsv13"); ok {
		t, err := expandLoadBalanceClientSslProfileSslCiphersTlsv13(d, v, "ssl_ciphers_tlsv13", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_ciphers_tlsv13"] = t
		}
	}

	if v, ok := d.GetOk("backend_ciphers_tlsv13"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendCiphersTlsv13(d, v, "backend_ciphers_tlsv13", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_ciphers_tlsv13"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiate_period"); ok {
		t, err := expandLoadBalanceClientSslProfileSslRenegotiatePeriod(d, v, "ssl_renegotiate_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_renegotiate_period"] = t
		}
	}

	if v, ok := d.GetOk("forward_proxy_certificate_caching"); ok {
		t, err := expandLoadBalanceClientSslProfileForwardProxyCertificateCaching(d, v, "forward_proxy_certificate_caching", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_proxy_certificate_caching"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceClientSslProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("forward_proxy_local_signing_ca"); ok {
		t, err := expandLoadBalanceClientSslProfileForwardProxyLocalSigningCa(d, v, "forward_proxy_local_signing_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_proxy_local_signing_ca"] = t
		}
	}

	if v, ok := d.GetOk("customized_ssl_ciphers"); ok {
		t, err := expandLoadBalanceClientSslProfileCustomizedSslCiphers(d, v, "customized_ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customized_ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ciphers"); ok {
		t, err := expandLoadBalanceClientSslProfileSslCiphers(d, v, "ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("reject_ocsp_stapling_with_missing_nextupdate"); ok {
		t, err := expandLoadBalanceClientSslProfileRejectOcspStaplingWithMissingNextupdate(d, v, "reject_ocsp_stapling_with_missing_nextupdate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reject-ocsp-stapling-with-missing-nextupdate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dynamic_record_sizing"); ok {
		t, err := expandLoadBalanceClientSslProfileSslDynamicRecordSizing(d, v, "ssl_dynamic_record_sizing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_dynamic_record_sizing"] = t
		}
	}

	if v, ok := d.GetOk("rfc7919_comply"); ok {
		t, err := expandLoadBalanceClientSslProfileRfc7919Comply(d, v, "rfc7919_comply", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc7919_comply"] = t
		}
	}

	if v, ok := d.GetOk("backend_ssl_ocsp_stapling_support"); ok {
		t, err := expandLoadBalanceClientSslProfileBackendSslOcspStaplingSupport(d, v, "backend_ssl_ocsp_stapling_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backend_ssl_OCSP_stapling_support"] = t
		}
	}

	if v, ok := d.GetOk("ssl_session_cache_flag"); ok {
		t, err := expandLoadBalanceClientSslProfileSslSessionCacheFlag(d, v, "ssl_session_cache_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_session_cache_flag"] = t
		}
	}

	if v, ok := d.GetOk("forward_proxy"); ok {
		t, err := expandLoadBalanceClientSslProfileForwardProxy(d, v, "forward_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_proxy"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiation_interval"); ok {
		t, err := expandLoadBalanceClientSslProfileSslRenegotiationInterval(d, v, "ssl_renegotiation_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_renegotiation_interval"] = t
		}
	}

	if v, ok := d.GetOk("forward_proxy_intermediate_ca_group"); ok {
		t, err := expandLoadBalanceClientSslProfileForwardProxyIntermediateCaGroup(d, v, "forward_proxy_intermediate_ca_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_proxy_intermediate_ca_group"] = t
		}
	}

	return &obj, nil
}
