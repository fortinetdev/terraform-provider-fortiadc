// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance real server ssl profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceRealServerSslProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceRealServerSslProfileRead,
		Schema: map[string]*schema.Schema{
			"new_ssl_ciphers_long": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_ocsp_stapling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secure_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"renegotiation_deny_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"renegotiate_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sni_forward_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cert_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc7919_comply": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"renegotiate_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ciphers_tlsv13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tls_ticket_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_reuse_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_reuse_flag": &schema.Schema{
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
func dataSourceLoadBalanceRealServerSslProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalanceRealServerSslProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceRealServerSslProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceRealServerSslProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceRealServerSslProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceRealServerSslProfileNewSslCiphersLong(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileServerOcspStapling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSecureRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSupportedGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiationDenyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiatePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileAllowSslVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSniForwardFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileCertVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileRfc7919Comply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiateSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileCiphersTlsv13(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileTlsTicketFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSessionReuseLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSslProfileSessionReuseFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceRealServerSslProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("new_ssl_ciphers_long", dataSourceFlattenLoadBalanceRealServerSslProfileNewSslCiphersLong(o["new_ssl_ciphers_long"], d, "new_ssl_ciphers_long")); err != nil {
		if !fortiAPIPatch(o["new_ssl_ciphers_long"]) {
			return fmt.Errorf("Error reading new_ssl_ciphers_long: %v", err)
		}
	}

	if err = d.Set("local_cert", dataSourceFlattenLoadBalanceRealServerSslProfileLocalCert(o["local_cert"], d, "local_cert")); err != nil {
		if !fortiAPIPatch(o["local_cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("server_ocsp_stapling", dataSourceFlattenLoadBalanceRealServerSslProfileServerOcspStapling(o["server_OCSP_stapling"], d, "server_ocsp_stapling")); err != nil {
		if !fortiAPIPatch(o["server_OCSP_stapling"]) {
			return fmt.Errorf("Error reading server_OCSP_stapling: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers_flag", dataSourceFlattenLoadBalanceRealServerSslProfileCustomizedSslCiphersFlag(o["customized_ssl_ciphers_flag"], d, "customized_ssl_ciphers_flag")); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("secure_renegotiation", dataSourceFlattenLoadBalanceRealServerSslProfileSecureRenegotiation(o["secure_renegotiation"], d, "secure_renegotiation")); err != nil {
		if !fortiAPIPatch(o["secure_renegotiation"]) {
			return fmt.Errorf("Error reading secure_renegotiation: %v", err)
		}
	}

	if err = d.Set("supported_groups", dataSourceFlattenLoadBalanceRealServerSslProfileSupportedGroups(o["supported_groups"], d, "supported_groups")); err != nil {
		if !fortiAPIPatch(o["supported_groups"]) {
			return fmt.Errorf("Error reading supported_groups: %v", err)
		}
	}

	if err = d.Set("renegotiation", dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiation(o["renegotiation"], d, "renegotiation")); err != nil {
		if !fortiAPIPatch(o["renegotiation"]) {
			return fmt.Errorf("Error reading renegotiation: %v", err)
		}
	}

	if err = d.Set("renegotiation_deny_action", dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiationDenyAction(o["renegotiation_deny_action"], d, "renegotiation_deny_action")); err != nil {
		if !fortiAPIPatch(o["renegotiation_deny_action"]) {
			return fmt.Errorf("Error reading renegotiation_deny_action: %v", err)
		}
	}

	if err = d.Set("renegotiate_period", dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiatePeriod(o["renegotiate_period"], d, "renegotiate_period")); err != nil {
		if !fortiAPIPatch(o["renegotiate_period"]) {
			return fmt.Errorf("Error reading renegotiate_period: %v", err)
		}
	}

	if err = d.Set("allow_ssl_versions", dataSourceFlattenLoadBalanceRealServerSslProfileAllowSslVersions(o["allow_ssl_versions"], d, "allow_ssl_versions")); err != nil {
		if !fortiAPIPatch(o["allow_ssl_versions"]) {
			return fmt.Errorf("Error reading allow_ssl_versions: %v", err)
		}
	}

	if err = d.Set("sni_forward_flag", dataSourceFlattenLoadBalanceRealServerSslProfileSniForwardFlag(o["sni_forward_flag"], d, "sni_forward_flag")); err != nil {
		if !fortiAPIPatch(o["sni_forward_flag"]) {
			return fmt.Errorf("Error reading sni_forward_flag: %v", err)
		}
	}

	if err = d.Set("ssl", dataSourceFlattenLoadBalanceRealServerSslProfileSsl(o["ssl"], d, "ssl")); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("sni", dataSourceFlattenLoadBalanceRealServerSslProfileSni(o["sni"], d, "sni")); err != nil {
		if !fortiAPIPatch(o["sni"]) {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceRealServerSslProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("cert_verify", dataSourceFlattenLoadBalanceRealServerSslProfileCertVerify(o["cert_verify"], d, "cert_verify")); err != nil {
		if !fortiAPIPatch(o["cert_verify"]) {
			return fmt.Errorf("Error reading cert_verify: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers", dataSourceFlattenLoadBalanceRealServerSslProfileCustomizedSslCiphers(o["customized_ssl_ciphers"], d, "customized_ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("rfc7919_comply", dataSourceFlattenLoadBalanceRealServerSslProfileRfc7919Comply(o["rfc7919_comply"], d, "rfc7919_comply")); err != nil {
		if !fortiAPIPatch(o["rfc7919_comply"]) {
			return fmt.Errorf("Error reading rfc7919_comply: %v", err)
		}
	}

	if err = d.Set("renegotiate_size", dataSourceFlattenLoadBalanceRealServerSslProfileRenegotiateSize(o["renegotiate_size"], d, "renegotiate_size")); err != nil {
		if !fortiAPIPatch(o["renegotiate_size"]) {
			return fmt.Errorf("Error reading renegotiate_size: %v", err)
		}
	}

	if err = d.Set("ciphers_tlsv13", dataSourceFlattenLoadBalanceRealServerSslProfileCiphersTlsv13(o["ciphers_tlsv13"], d, "ciphers_tlsv13")); err != nil {
		if !fortiAPIPatch(o["ciphers_tlsv13"]) {
			return fmt.Errorf("Error reading ciphers_tlsv13: %v", err)
		}
	}

	if err = d.Set("tls_ticket_flag", dataSourceFlattenLoadBalanceRealServerSslProfileTlsTicketFlag(o["tls_ticket_flag"], d, "tls_ticket_flag")); err != nil {
		if !fortiAPIPatch(o["tls_ticket_flag"]) {
			return fmt.Errorf("Error reading tls_ticket_flag: %v", err)
		}
	}

	if err = d.Set("session_reuse_limit", dataSourceFlattenLoadBalanceRealServerSslProfileSessionReuseLimit(o["session_reuse_limit"], d, "session_reuse_limit")); err != nil {
		if !fortiAPIPatch(o["session_reuse_limit"]) {
			return fmt.Errorf("Error reading session_reuse_limit: %v", err)
		}
	}

	if err = d.Set("session_reuse_flag", dataSourceFlattenLoadBalanceRealServerSslProfileSessionReuseFlag(o["session_reuse_flag"], d, "session_reuse_flag")); err != nil {
		if !fortiAPIPatch(o["session_reuse_flag"]) {
			return fmt.Errorf("Error reading session_reuse_flag: %v", err)
		}
	}

	return nil
}
