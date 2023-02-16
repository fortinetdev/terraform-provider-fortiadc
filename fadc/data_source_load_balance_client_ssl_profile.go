// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance client ssl profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceClientSslProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceClientSslProfileRead,
		Schema: map[string]*schema.Schema{
			"ssl_allowed_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_renegotiate_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_certificate_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_secure_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_forward_client_certificate_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_forward_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_tls_tickets": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_sni_required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_certificate_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_ssl_sni_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_certificate_verify_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_certificate_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_ssl_allowed_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_dh_param_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_ciphers_tlsv13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_ciphers_tlsv13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_renegotiate_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_proxy_certificate_caching": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"forward_proxy_local_signing_ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reject_ocsp_stapling_with_missing_nextupdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_dynamic_record_sizing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc7919_comply": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_ssl_ocsp_stapling_support": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_session_cache_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_renegotiation_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_proxy_intermediate_ca_group": &schema.Schema{
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
func dataSourceLoadBalanceClientSslProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalanceClientSslProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceClientSslProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceClientSslProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceClientSslProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceClientSslProfileSslAllowedVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiateSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendCertificateVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslSecureRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileHttpForwardClientCertificateHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSupportedGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileHttpForwardClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileUseTlsTickets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileClientSniRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileClientCertificateVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendSslSniForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileClientCertificateVerifyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileLocalCertificateGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendSslAllowedVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslDhParamSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslCiphersTlsv13(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendCiphersTlsv13(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiatePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileForwardProxyCertificateCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileForwardProxyLocalSigningCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileRejectOcspStaplingWithMissingNextupdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslDynamicRecordSizing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileRfc7919Comply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileBackendSslOcspStaplingSupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslSessionCacheFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileForwardProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiationInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceClientSslProfileForwardProxyIntermediateCaGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceClientSslProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ssl_allowed_versions", dataSourceFlattenLoadBalanceClientSslProfileSslAllowedVersions(o["ssl-allowed_versions"], d, "ssl_allowed_versions")); err != nil {
		if !fortiAPIPatch(o["ssl-allowed_versions"]) {
			return fmt.Errorf("Error reading ssl-allowed_versions: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiate_size", dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiateSize(o["ssl_renegotiate_size"], d, "ssl_renegotiate_size")); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiate_size"]) {
			return fmt.Errorf("Error reading ssl_renegotiate_size: %v", err)
		}
	}

	if err = d.Set("backend_certificate_verify", dataSourceFlattenLoadBalanceClientSslProfileBackendCertificateVerify(o["backend_certificate_verify"], d, "backend_certificate_verify")); err != nil {
		if !fortiAPIPatch(o["backend_certificate_verify"]) {
			return fmt.Errorf("Error reading backend_certificate_verify: %v", err)
		}
	}

	if err = d.Set("backend_customized_ssl_ciphers", dataSourceFlattenLoadBalanceClientSslProfileBackendCustomizedSslCiphers(o["backend_customized_ssl_ciphers"], d, "backend_customized_ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["backend_customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading backend_customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation", dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiation(o["ssl_renegotiation"], d, "ssl_renegotiation")); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiation"]) {
			return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers_flag", dataSourceFlattenLoadBalanceClientSslProfileCustomizedSslCiphersFlag(o["customized_ssl_ciphers_flag"], d, "customized_ssl_ciphers_flag")); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("backend_customized_ssl_ciphers_flag", dataSourceFlattenLoadBalanceClientSslProfileBackendCustomizedSslCiphersFlag(o["backend_customized_ssl_ciphers_flag"], d, "backend_customized_ssl_ciphers_flag")); err != nil {
		if !fortiAPIPatch(o["backend_customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading backend_customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("ssl_secure_renegotiation", dataSourceFlattenLoadBalanceClientSslProfileSslSecureRenegotiation(o["ssl_secure_renegotiation"], d, "ssl_secure_renegotiation")); err != nil {
		if !fortiAPIPatch(o["ssl_secure_renegotiation"]) {
			return fmt.Errorf("Error reading ssl_secure_renegotiation: %v", err)
		}
	}

	if err = d.Set("http_forward_client_certificate_header", dataSourceFlattenLoadBalanceClientSslProfileHttpForwardClientCertificateHeader(o["http-forward_client_certificate_header"], d, "http_forward_client_certificate_header")); err != nil {
		if !fortiAPIPatch(o["http-forward_client_certificate_header"]) {
			return fmt.Errorf("Error reading http-forward_client_certificate_header: %v", err)
		}
	}

	if err = d.Set("supported_groups", dataSourceFlattenLoadBalanceClientSslProfileSupportedGroups(o["supported_groups"], d, "supported_groups")); err != nil {
		if !fortiAPIPatch(o["supported_groups"]) {
			return fmt.Errorf("Error reading supported_groups: %v", err)
		}
	}

	if err = d.Set("http_forward_client_certificate", dataSourceFlattenLoadBalanceClientSslProfileHttpForwardClientCertificate(o["http-forward_client_certificate"], d, "http_forward_client_certificate")); err != nil {
		if !fortiAPIPatch(o["http-forward_client_certificate"]) {
			return fmt.Errorf("Error reading http-forward_client_certificate: %v", err)
		}
	}

	if err = d.Set("use_tls_tickets", dataSourceFlattenLoadBalanceClientSslProfileUseTlsTickets(o["use_tls_tickets"], d, "use_tls_tickets")); err != nil {
		if !fortiAPIPatch(o["use_tls_tickets"]) {
			return fmt.Errorf("Error reading use_tls_tickets: %v", err)
		}
	}

	if err = d.Set("client_sni_required", dataSourceFlattenLoadBalanceClientSslProfileClientSniRequired(o["client_sni_required"], d, "client_sni_required")); err != nil {
		if !fortiAPIPatch(o["client_sni_required"]) {
			return fmt.Errorf("Error reading client_sni_required: %v", err)
		}
	}

	if err = d.Set("backend_ssl_ciphers", dataSourceFlattenLoadBalanceClientSslProfileBackendSslCiphers(o["backend_ssl_ciphers"], d, "backend_ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["backend_ssl_ciphers"]) {
			return fmt.Errorf("Error reading backend_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("client_certificate_verify", dataSourceFlattenLoadBalanceClientSslProfileClientCertificateVerify(o["client_certificate_verify"], d, "client_certificate_verify")); err != nil {
		if !fortiAPIPatch(o["client_certificate_verify"]) {
			return fmt.Errorf("Error reading client_certificate_verify: %v", err)
		}
	}

	if err = d.Set("backend_ssl_sni_forward", dataSourceFlattenLoadBalanceClientSslProfileBackendSslSniForward(o["backend_ssl_sni_forward"], d, "backend_ssl_sni_forward")); err != nil {
		if !fortiAPIPatch(o["backend_ssl_sni_forward"]) {
			return fmt.Errorf("Error reading backend_ssl_sni_forward: %v", err)
		}
	}

	if err = d.Set("client_certificate_verify_mode", dataSourceFlattenLoadBalanceClientSslProfileClientCertificateVerifyMode(o["client_certificate_verify_mode"], d, "client_certificate_verify_mode")); err != nil {
		if !fortiAPIPatch(o["client_certificate_verify_mode"]) {
			return fmt.Errorf("Error reading client_certificate_verify_mode: %v", err)
		}
	}

	if err = d.Set("local_certificate_group", dataSourceFlattenLoadBalanceClientSslProfileLocalCertificateGroup(o["local_certificate_group"], d, "local_certificate_group")); err != nil {
		if !fortiAPIPatch(o["local_certificate_group"]) {
			return fmt.Errorf("Error reading local_certificate_group: %v", err)
		}
	}

	if err = d.Set("backend_ssl_allowed_versions", dataSourceFlattenLoadBalanceClientSslProfileBackendSslAllowedVersions(o["backend_ssl_allowed_versions"], d, "backend_ssl_allowed_versions")); err != nil {
		if !fortiAPIPatch(o["backend_ssl_allowed_versions"]) {
			return fmt.Errorf("Error reading backend_ssl_allowed_versions: %v", err)
		}
	}

	if err = d.Set("ssl_dh_param_size", dataSourceFlattenLoadBalanceClientSslProfileSslDhParamSize(o["ssl_dh_param_size"], d, "ssl_dh_param_size")); err != nil {
		if !fortiAPIPatch(o["ssl_dh_param_size"]) {
			return fmt.Errorf("Error reading ssl_dh_param_size: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers_tlsv13", dataSourceFlattenLoadBalanceClientSslProfileSslCiphersTlsv13(o["ssl_ciphers_tlsv13"], d, "ssl_ciphers_tlsv13")); err != nil {
		if !fortiAPIPatch(o["ssl_ciphers_tlsv13"]) {
			return fmt.Errorf("Error reading ssl_ciphers_tlsv13: %v", err)
		}
	}

	if err = d.Set("backend_ciphers_tlsv13", dataSourceFlattenLoadBalanceClientSslProfileBackendCiphersTlsv13(o["backend_ciphers_tlsv13"], d, "backend_ciphers_tlsv13")); err != nil {
		if !fortiAPIPatch(o["backend_ciphers_tlsv13"]) {
			return fmt.Errorf("Error reading backend_ciphers_tlsv13: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiate_period", dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiatePeriod(o["ssl_renegotiate_period"], d, "ssl_renegotiate_period")); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiate_period"]) {
			return fmt.Errorf("Error reading ssl_renegotiate_period: %v", err)
		}
	}

	if err = d.Set("forward_proxy_certificate_caching", dataSourceFlattenLoadBalanceClientSslProfileForwardProxyCertificateCaching(o["forward_proxy_certificate_caching"], d, "forward_proxy_certificate_caching")); err != nil {
		if !fortiAPIPatch(o["forward_proxy_certificate_caching"]) {
			return fmt.Errorf("Error reading forward_proxy_certificate_caching: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceClientSslProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("forward_proxy_local_signing_ca", dataSourceFlattenLoadBalanceClientSslProfileForwardProxyLocalSigningCa(o["forward_proxy_local_signing_ca"], d, "forward_proxy_local_signing_ca")); err != nil {
		if !fortiAPIPatch(o["forward_proxy_local_signing_ca"]) {
			return fmt.Errorf("Error reading forward_proxy_local_signing_ca: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers", dataSourceFlattenLoadBalanceClientSslProfileCustomizedSslCiphers(o["customized_ssl_ciphers"], d, "customized_ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers", dataSourceFlattenLoadBalanceClientSslProfileSslCiphers(o["ssl_ciphers"], d, "ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["ssl_ciphers"]) {
			return fmt.Errorf("Error reading ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("reject_ocsp_stapling_with_missing_nextupdate", dataSourceFlattenLoadBalanceClientSslProfileRejectOcspStaplingWithMissingNextupdate(o["reject-ocsp-stapling-with-missing-nextupdate"], d, "reject_ocsp_stapling_with_missing_nextupdate")); err != nil {
		if !fortiAPIPatch(o["reject-ocsp-stapling-with-missing-nextupdate"]) {
			return fmt.Errorf("Error reading reject-ocsp-stapling-with-missing-nextupdate: %v", err)
		}
	}

	if err = d.Set("ssl_dynamic_record_sizing", dataSourceFlattenLoadBalanceClientSslProfileSslDynamicRecordSizing(o["ssl_dynamic_record_sizing"], d, "ssl_dynamic_record_sizing")); err != nil {
		if !fortiAPIPatch(o["ssl_dynamic_record_sizing"]) {
			return fmt.Errorf("Error reading ssl_dynamic_record_sizing: %v", err)
		}
	}

	if err = d.Set("rfc7919_comply", dataSourceFlattenLoadBalanceClientSslProfileRfc7919Comply(o["rfc7919_comply"], d, "rfc7919_comply")); err != nil {
		if !fortiAPIPatch(o["rfc7919_comply"]) {
			return fmt.Errorf("Error reading rfc7919_comply: %v", err)
		}
	}

	if err = d.Set("backend_ssl_ocsp_stapling_support", dataSourceFlattenLoadBalanceClientSslProfileBackendSslOcspStaplingSupport(o["backend_ssl_OCSP_stapling_support"], d, "backend_ssl_ocsp_stapling_support")); err != nil {
		if !fortiAPIPatch(o["backend_ssl_OCSP_stapling_support"]) {
			return fmt.Errorf("Error reading backend_ssl_OCSP_stapling_support: %v", err)
		}
	}

	if err = d.Set("ssl_session_cache_flag", dataSourceFlattenLoadBalanceClientSslProfileSslSessionCacheFlag(o["ssl_session_cache_flag"], d, "ssl_session_cache_flag")); err != nil {
		if !fortiAPIPatch(o["ssl_session_cache_flag"]) {
			return fmt.Errorf("Error reading ssl_session_cache_flag: %v", err)
		}
	}

	if err = d.Set("forward_proxy", dataSourceFlattenLoadBalanceClientSslProfileForwardProxy(o["forward_proxy"], d, "forward_proxy")); err != nil {
		if !fortiAPIPatch(o["forward_proxy"]) {
			return fmt.Errorf("Error reading forward_proxy: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation_interval", dataSourceFlattenLoadBalanceClientSslProfileSslRenegotiationInterval(o["ssl_renegotiation_interval"], d, "ssl_renegotiation_interval")); err != nil {
		if !fortiAPIPatch(o["ssl_renegotiation_interval"]) {
			return fmt.Errorf("Error reading ssl_renegotiation_interval: %v", err)
		}
	}

	if err = d.Set("forward_proxy_intermediate_ca_group", dataSourceFlattenLoadBalanceClientSslProfileForwardProxyIntermediateCaGroup(o["forward_proxy_intermediate_ca_group"], d, "forward_proxy_intermediate_ca_group")); err != nil {
		if !fortiAPIPatch(o["forward_proxy_intermediate_ca_group"]) {
			return fmt.Errorf("Error reading forward_proxy_intermediate_ca_group: %v", err)
		}
	}

	return nil
}
