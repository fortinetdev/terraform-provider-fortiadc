// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceProfileRead,
		Schema: map[string]*schema.Schema{
			"msg_encode_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ram_caching": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_tcp_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_authenticate_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_max_query_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"decompression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"intermediate_ca_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_ageout_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_disable_command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failed_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_client_certificate_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tune_bufsize": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_disable_command_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failed_server_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_radius_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_malform_query_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_sni_required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sip_dlg_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"insert_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"opt_header_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_request_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_response_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_http_headers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"origin_realm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"geoip_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"response_half_closed_request": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_max_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_reputation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failed_client_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_udp_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_cert_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http2_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"starttls_active_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"new_ssl_ciphers_long": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_x_forwarded_for_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idle_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_header_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"deploy_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"origin_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"media_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_close_propagation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_tcp_session_after_fin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sip_max_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_tls_tickets": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_entry_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"queue_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connect_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"opt_trailer_hex": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"length_indicator_shift": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_reputation_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"length_indicator_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_send_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_auth_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failed_client_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stateless": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cert_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"geoip_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"length_indicator_type": &schema.Schema{
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
func dataSourceLoadBalanceProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfile: type error")
	}

	o, err := c.ReadLoadBalanceProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceProfileMsgEncodeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileRamCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileTimeoutTcpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsAuthenticateFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsMaxQueryLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDecompression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileClientProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileIntermediateCaGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileVendorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsCacheAgeoutTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileTimeoutIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSmtpDisableCommand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileProductName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileFailedServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileForwardClientCertificateHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileTuneBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileClientAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSmtpDisableCommandStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileClientKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileFailedServerStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileTimeoutRadiusSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsMalformQueryAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileClientSniRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttpKeepaliveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSipDlgTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileClientTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileMysqlMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileInsertClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileOptHeaderLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttpRequestTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileForwardClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsCacheResponseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileMaxHttpHeaders(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileOriginRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileGeoipList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerKeepaliveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileClientSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileCompression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileResponseHalfClosedRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDynamicAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerMaxSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileIpReputation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileFailedClientStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileTimeoutUdpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileLocalCertGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttp2Profile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileStarttlsActiveMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSmtpDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsCacheSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsCacheFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileNewSslCiphersLong(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttpXForwardedForHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileIdleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileMaxHeaderSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDeployMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileOriginHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileMediaAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerClosePropagation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileTimeoutTcpSessionAfterFin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSipMaxSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSslProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileUseTlsTickets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDnsCacheEntrySize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileQueueTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileConnectTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileOptTrailerHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileLengthIndicatorShift(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileIpReputationRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileLengthIndicatorSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttpXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileAllowSslVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileHttpSendTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileDynamicAuthPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileFailedClientType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileStateless(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileServerTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileCertVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileGeoipRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileLengthIndicatorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("msg_encode_type", dataSourceFlattenLoadBalanceProfileMsgEncodeType(o["msg_encode_type"], d, "msg_encode_type")); err != nil {
		if !fortiAPIPatch(o["msg_encode_type"]) {
			return fmt.Errorf("Error reading msg_encode_type: %v", err)
		}
	}

	if err = d.Set("whitelist", dataSourceFlattenLoadBalanceProfileWhitelist(o["whitelist"], d, "whitelist")); err != nil {
		if !fortiAPIPatch(o["whitelist"]) {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if err = d.Set("ram_caching", dataSourceFlattenLoadBalanceProfileRamCaching(o["ram_caching"], d, "ram_caching")); err != nil {
		if !fortiAPIPatch(o["ram_caching"]) {
			return fmt.Errorf("Error reading ram_caching: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_session", dataSourceFlattenLoadBalanceProfileTimeoutTcpSession(o["timeout_tcp_session"], d, "timeout_tcp_session")); err != nil {
		if !fortiAPIPatch(o["timeout_tcp_session"]) {
			return fmt.Errorf("Error reading timeout_tcp_session: %v", err)
		}
	}

	if err = d.Set("dns_authenticate_flag", dataSourceFlattenLoadBalanceProfileDnsAuthenticateFlag(o["dns_authenticate_flag"], d, "dns_authenticate_flag")); err != nil {
		if !fortiAPIPatch(o["dns_authenticate_flag"]) {
			return fmt.Errorf("Error reading dns_authenticate_flag: %v", err)
		}
	}

	if err = d.Set("dns_max_query_length", dataSourceFlattenLoadBalanceProfileDnsMaxQueryLength(o["dns_max_query_length"], d, "dns_max_query_length")); err != nil {
		if !fortiAPIPatch(o["dns_max_query_length"]) {
			return fmt.Errorf("Error reading dns_max_query_length: %v", err)
		}
	}

	if err = d.Set("decompression", dataSourceFlattenLoadBalanceProfileDecompression(o["decompression"], d, "decompression")); err != nil {
		if !fortiAPIPatch(o["decompression"]) {
			return fmt.Errorf("Error reading decompression: %v", err)
		}
	}

	if err = d.Set("client_protocol", dataSourceFlattenLoadBalanceProfileClientProtocol(o["client_protocol"], d, "client_protocol")); err != nil {
		if !fortiAPIPatch(o["client_protocol"]) {
			return fmt.Errorf("Error reading client_protocol: %v", err)
		}
	}

	if err = d.Set("intermediate_ca_group", dataSourceFlattenLoadBalanceProfileIntermediateCaGroup(o["intermediate_ca_group"], d, "intermediate_ca_group")); err != nil {
		if !fortiAPIPatch(o["intermediate_ca_group"]) {
			return fmt.Errorf("Error reading intermediate_ca_group: %v", err)
		}
	}

	if err = d.Set("vendor_id", dataSourceFlattenLoadBalanceProfileVendorId(o["vendor_id"], d, "vendor_id")); err != nil {
		if !fortiAPIPatch(o["vendor_id"]) {
			return fmt.Errorf("Error reading vendor_id: %v", err)
		}
	}

	if err = d.Set("dns_cache_ageout_time", dataSourceFlattenLoadBalanceProfileDnsCacheAgeoutTime(o["dns_cache_ageout_time"], d, "dns_cache_ageout_time")); err != nil {
		if !fortiAPIPatch(o["dns_cache_ageout_time"]) {
			return fmt.Errorf("Error reading dns_cache_ageout_time: %v", err)
		}
	}

	if err = d.Set("ssl", dataSourceFlattenLoadBalanceProfileSsl(o["ssl"], d, "ssl")); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("timeout_ip_session", dataSourceFlattenLoadBalanceProfileTimeoutIpSession(o["timeout_ip_session"], d, "timeout_ip_session")); err != nil {
		if !fortiAPIPatch(o["timeout_ip_session"]) {
			return fmt.Errorf("Error reading timeout_ip_session: %v", err)
		}
	}

	if err = d.Set("smtp_disable_command", dataSourceFlattenLoadBalanceProfileSmtpDisableCommand(o["smtp_disable_command"], d, "smtp_disable_command")); err != nil {
		if !fortiAPIPatch(o["smtp_disable_command"]) {
			return fmt.Errorf("Error reading smtp_disable_command: %v", err)
		}
	}

	if err = d.Set("server_keepalive", dataSourceFlattenLoadBalanceProfileServerKeepalive(o["server_keepalive"], d, "server_keepalive")); err != nil {
		if !fortiAPIPatch(o["server_keepalive"]) {
			return fmt.Errorf("Error reading server_keepalive: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers", dataSourceFlattenLoadBalanceProfileSslCiphers(o["ssl_ciphers"], d, "ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["ssl_ciphers"]) {
			return fmt.Errorf("Error reading ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("product_name", dataSourceFlattenLoadBalanceProfileProductName(o["product_name"], d, "product_name")); err != nil {
		if !fortiAPIPatch(o["product_name"]) {
			return fmt.Errorf("Error reading product_name: %v", err)
		}
	}

	if err = d.Set("failed_server_type", dataSourceFlattenLoadBalanceProfileFailedServerType(o["failed_server_type"], d, "failed_server_type")); err != nil {
		if !fortiAPIPatch(o["failed_server_type"]) {
			return fmt.Errorf("Error reading failed_server_type: %v", err)
		}
	}

	if err = d.Set("forward_client_certificate_header", dataSourceFlattenLoadBalanceProfileForwardClientCertificateHeader(o["forward_client_certificate_header"], d, "forward_client_certificate_header")); err != nil {
		if !fortiAPIPatch(o["forward_client_certificate_header"]) {
			return fmt.Errorf("Error reading forward_client_certificate_header: %v", err)
		}
	}

	if err = d.Set("tune_bufsize", dataSourceFlattenLoadBalanceProfileTuneBufsize(o["tune-bufsize"], d, "tune_bufsize")); err != nil {
		if !fortiAPIPatch(o["tune-bufsize"]) {
			return fmt.Errorf("Error reading tune-bufsize: %v", err)
		}
	}

	if err = d.Set("local_cert", dataSourceFlattenLoadBalanceProfileLocalCert(o["local_cert"], d, "local_cert")); err != nil {
		if !fortiAPIPatch(o["local_cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("client_address", dataSourceFlattenLoadBalanceProfileClientAddress(o["client_address"], d, "client_address")); err != nil {
		if !fortiAPIPatch(o["client_address"]) {
			return fmt.Errorf("Error reading client_address: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers_flag", dataSourceFlattenLoadBalanceProfileCustomizedSslCiphersFlag(o["customized_ssl_ciphers_flag"], d, "customized_ssl_ciphers_flag")); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("smtp_disable_command_status", dataSourceFlattenLoadBalanceProfileSmtpDisableCommandStatus(o["smtp_disable_command_status"], d, "smtp_disable_command_status")); err != nil {
		if !fortiAPIPatch(o["smtp_disable_command_status"]) {
			return fmt.Errorf("Error reading smtp_disable_command_status: %v", err)
		}
	}

	if err = d.Set("client_keepalive", dataSourceFlattenLoadBalanceProfileClientKeepalive(o["client_keepalive"], d, "client_keepalive")); err != nil {
		if !fortiAPIPatch(o["client_keepalive"]) {
			return fmt.Errorf("Error reading client_keepalive: %v", err)
		}
	}

	if err = d.Set("failed_server_str", dataSourceFlattenLoadBalanceProfileFailedServerStr(o["failed_server_str"], d, "failed_server_str")); err != nil {
		if !fortiAPIPatch(o["failed_server_str"]) {
			return fmt.Errorf("Error reading failed_server_str: %v", err)
		}
	}

	if err = d.Set("timeout_radius_session", dataSourceFlattenLoadBalanceProfileTimeoutRadiusSession(o["timeout_radius_session"], d, "timeout_radius_session")); err != nil {
		if !fortiAPIPatch(o["timeout_radius_session"]) {
			return fmt.Errorf("Error reading timeout_radius_session: %v", err)
		}
	}

	if err = d.Set("dns_malform_query_action", dataSourceFlattenLoadBalanceProfileDnsMalformQueryAction(o["dns_malform_query_action"], d, "dns_malform_query_action")); err != nil {
		if !fortiAPIPatch(o["dns_malform_query_action"]) {
			return fmt.Errorf("Error reading dns_malform_query_action: %v", err)
		}
	}

	if err = d.Set("client_sni_required", dataSourceFlattenLoadBalanceProfileClientSniRequired(o["client_sni_required"], d, "client_sni_required")); err != nil {
		if !fortiAPIPatch(o["client_sni_required"]) {
			return fmt.Errorf("Error reading client_sni_required: %v", err)
		}
	}

	if err = d.Set("http_keepalive_timeout", dataSourceFlattenLoadBalanceProfileHttpKeepaliveTimeout(o["http_keepalive_timeout"], d, "http_keepalive_timeout")); err != nil {
		if !fortiAPIPatch(o["http_keepalive_timeout"]) {
			return fmt.Errorf("Error reading http_keepalive_timeout: %v", err)
		}
	}

	if err = d.Set("sip_dlg_timeout", dataSourceFlattenLoadBalanceProfileSipDlgTimeout(o["sip_dlg_timeout"], d, "sip_dlg_timeout")); err != nil {
		if !fortiAPIPatch(o["sip_dlg_timeout"]) {
			return fmt.Errorf("Error reading sip_dlg_timeout: %v", err)
		}
	}

	if err = d.Set("client_timeout", dataSourceFlattenLoadBalanceProfileClientTimeout(o["client_timeout"], d, "client_timeout")); err != nil {
		if !fortiAPIPatch(o["client_timeout"]) {
			return fmt.Errorf("Error reading client_timeout: %v", err)
		}
	}

	if err = d.Set("mysql_mode", dataSourceFlattenLoadBalanceProfileMysqlMode(o["mysql_mode"], d, "mysql_mode")); err != nil {
		if !fortiAPIPatch(o["mysql_mode"]) {
			return fmt.Errorf("Error reading mysql_mode: %v", err)
		}
	}

	if err = d.Set("insert_client_ip", dataSourceFlattenLoadBalanceProfileInsertClientIp(o["insert_client_ip"], d, "insert_client_ip")); err != nil {
		if !fortiAPIPatch(o["insert_client_ip"]) {
			return fmt.Errorf("Error reading insert_client_ip: %v", err)
		}
	}

	if err = d.Set("opt_header_length", dataSourceFlattenLoadBalanceProfileOptHeaderLength(o["opt_header_length"], d, "opt_header_length")); err != nil {
		if !fortiAPIPatch(o["opt_header_length"]) {
			return fmt.Errorf("Error reading opt_header_length: %v", err)
		}
	}

	if err = d.Set("http_request_timeout", dataSourceFlattenLoadBalanceProfileHttpRequestTimeout(o["http_request_timeout"], d, "http_request_timeout")); err != nil {
		if !fortiAPIPatch(o["http_request_timeout"]) {
			return fmt.Errorf("Error reading http_request_timeout: %v", err)
		}
	}

	if err = d.Set("http_mode", dataSourceFlattenLoadBalanceProfileHttpMode(o["http_mode"], d, "http_mode")); err != nil {
		if !fortiAPIPatch(o["http_mode"]) {
			return fmt.Errorf("Error reading http_mode: %v", err)
		}
	}

	if err = d.Set("forward_client_certificate", dataSourceFlattenLoadBalanceProfileForwardClientCertificate(o["forward_client_certificate"], d, "forward_client_certificate")); err != nil {
		if !fortiAPIPatch(o["forward_client_certificate"]) {
			return fmt.Errorf("Error reading forward_client_certificate: %v", err)
		}
	}

	if err = d.Set("dns_cache_response_type", dataSourceFlattenLoadBalanceProfileDnsCacheResponseType(o["dns_cache_response_type"], d, "dns_cache_response_type")); err != nil {
		if !fortiAPIPatch(o["dns_cache_response_type"]) {
			return fmt.Errorf("Error reading dns_cache_response_type: %v", err)
		}
	}

	if err = d.Set("max_http_headers", dataSourceFlattenLoadBalanceProfileMaxHttpHeaders(o["max_http_headers"], d, "max_http_headers")); err != nil {
		if !fortiAPIPatch(o["max_http_headers"]) {
			return fmt.Errorf("Error reading max_http_headers: %v", err)
		}
	}

	if err = d.Set("origin_realm", dataSourceFlattenLoadBalanceProfileOriginRealm(o["origin_realm"], d, "origin_realm")); err != nil {
		if !fortiAPIPatch(o["origin_realm"]) {
			return fmt.Errorf("Error reading origin_realm: %v", err)
		}
	}

	if err = d.Set("geoip_list", dataSourceFlattenLoadBalanceProfileGeoipList(o["geoip_list"], d, "geoip_list")); err != nil {
		if !fortiAPIPatch(o["geoip_list"]) {
			return fmt.Errorf("Error reading geoip_list: %v", err)
		}
	}

	if err = d.Set("server_keepalive_timeout", dataSourceFlattenLoadBalanceProfileServerKeepaliveTimeout(o["server_keepalive_timeout"], d, "server_keepalive_timeout")); err != nil {
		if !fortiAPIPatch(o["server_keepalive_timeout"]) {
			return fmt.Errorf("Error reading server_keepalive_timeout: %v", err)
		}
	}

	if err = d.Set("client_ssl", dataSourceFlattenLoadBalanceProfileClientSsl(o["client_ssl"], d, "client_ssl")); err != nil {
		if !fortiAPIPatch(o["client_ssl"]) {
			return fmt.Errorf("Error reading client_ssl: %v", err)
		}
	}

	if err = d.Set("compression", dataSourceFlattenLoadBalanceProfileCompression(o["compression"], d, "compression")); err != nil {
		if !fortiAPIPatch(o["compression"]) {
			return fmt.Errorf("Error reading compression: %v", err)
		}
	}

	if err = d.Set("response_half_closed_request", dataSourceFlattenLoadBalanceProfileResponseHalfClosedRequest(o["response_half_closed_request"], d, "response_half_closed_request")); err != nil {
		if !fortiAPIPatch(o["response_half_closed_request"]) {
			return fmt.Errorf("Error reading response_half_closed_request: %v", err)
		}
	}

	if err = d.Set("dynamic_auth", dataSourceFlattenLoadBalanceProfileDynamicAuth(o["dynamic_auth"], d, "dynamic_auth")); err != nil {
		if !fortiAPIPatch(o["dynamic_auth"]) {
			return fmt.Errorf("Error reading dynamic_auth: %v", err)
		}
	}

	if err = d.Set("server_max_size", dataSourceFlattenLoadBalanceProfileServerMaxSize(o["server_max_size"], d, "server_max_size")); err != nil {
		if !fortiAPIPatch(o["server_max_size"]) {
			return fmt.Errorf("Error reading server_max_size: %v", err)
		}
	}

	if err = d.Set("ip_reputation", dataSourceFlattenLoadBalanceProfileIpReputation(o["ip_reputation"], d, "ip_reputation")); err != nil {
		if !fortiAPIPatch(o["ip_reputation"]) {
			return fmt.Errorf("Error reading ip_reputation: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceProfileType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("failed_client_str", dataSourceFlattenLoadBalanceProfileFailedClientStr(o["failed_client_str"], d, "failed_client_str")); err != nil {
		if !fortiAPIPatch(o["failed_client_str"]) {
			return fmt.Errorf("Error reading failed_client_str: %v", err)
		}
	}

	if err = d.Set("timeout_udp_session", dataSourceFlattenLoadBalanceProfileTimeoutUdpSession(o["timeout_udp_session"], d, "timeout_udp_session")); err != nil {
		if !fortiAPIPatch(o["timeout_udp_session"]) {
			return fmt.Errorf("Error reading timeout_udp_session: %v", err)
		}
	}

	if err = d.Set("local_cert_group", dataSourceFlattenLoadBalanceProfileLocalCertGroup(o["local_cert_group"], d, "local_cert_group")); err != nil {
		if !fortiAPIPatch(o["local_cert_group"]) {
			return fmt.Errorf("Error reading local_cert_group: %v", err)
		}
	}

	if err = d.Set("http2_profile", dataSourceFlattenLoadBalanceProfileHttp2Profile(o["http2_profile"], d, "http2_profile")); err != nil {
		if !fortiAPIPatch(o["http2_profile"]) {
			return fmt.Errorf("Error reading http2_profile: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers", dataSourceFlattenLoadBalanceProfileCustomizedSslCiphers(o["customized_ssl_ciphers"], d, "customized_ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("starttls_active_mode", dataSourceFlattenLoadBalanceProfileStarttlsActiveMode(o["starttls_active_mode"], d, "starttls_active_mode")); err != nil {
		if !fortiAPIPatch(o["starttls_active_mode"]) {
			return fmt.Errorf("Error reading starttls_active_mode: %v", err)
		}
	}

	if err = d.Set("server_age", dataSourceFlattenLoadBalanceProfileServerAge(o["server_age"], d, "server_age")); err != nil {
		if !fortiAPIPatch(o["server_age"]) {
			return fmt.Errorf("Error reading server_age: %v", err)
		}
	}

	if err = d.Set("smtp_domain_name", dataSourceFlattenLoadBalanceProfileSmtpDomainName(o["smtp_domain_name"], d, "smtp_domain_name")); err != nil {
		if !fortiAPIPatch(o["smtp_domain_name"]) {
			return fmt.Errorf("Error reading smtp_domain_name: %v", err)
		}
	}

	if err = d.Set("dns_cache_size", dataSourceFlattenLoadBalanceProfileDnsCacheSize(o["dns_cache_size"], d, "dns_cache_size")); err != nil {
		if !fortiAPIPatch(o["dns_cache_size"]) {
			return fmt.Errorf("Error reading dns_cache_size: %v", err)
		}
	}

	if err = d.Set("dns_cache_flag", dataSourceFlattenLoadBalanceProfileDnsCacheFlag(o["dns_cache_flag"], d, "dns_cache_flag")); err != nil {
		if !fortiAPIPatch(o["dns_cache_flag"]) {
			return fmt.Errorf("Error reading dns_cache_flag: %v", err)
		}
	}

	if err = d.Set("new_ssl_ciphers_long", dataSourceFlattenLoadBalanceProfileNewSslCiphersLong(o["new_ssl_ciphers_long"], d, "new_ssl_ciphers_long")); err != nil {
		if !fortiAPIPatch(o["new_ssl_ciphers_long"]) {
			return fmt.Errorf("Error reading new_ssl_ciphers_long: %v", err)
		}
	}

	if err = d.Set("http_x_forwarded_for_header", dataSourceFlattenLoadBalanceProfileHttpXForwardedForHeader(o["http_x_forwarded_for_header"], d, "http_x_forwarded_for_header")); err != nil {
		if !fortiAPIPatch(o["http_x_forwarded_for_header"]) {
			return fmt.Errorf("Error reading http_x_forwarded_for_header: %v", err)
		}
	}

	if err = d.Set("idle_time", dataSourceFlattenLoadBalanceProfileIdleTime(o["idle_time"], d, "idle_time")); err != nil {
		if !fortiAPIPatch(o["idle_time"]) {
			return fmt.Errorf("Error reading idle_time: %v", err)
		}
	}

	if err = d.Set("max_header_size", dataSourceFlattenLoadBalanceProfileMaxHeaderSize(o["max_header_size"], d, "max_header_size")); err != nil {
		if !fortiAPIPatch(o["max_header_size"]) {
			return fmt.Errorf("Error reading max_header_size: %v", err)
		}
	}

	if err = d.Set("deploy_mode", dataSourceFlattenLoadBalanceProfileDeployMode(o["deploy_mode"], d, "deploy_mode")); err != nil {
		if !fortiAPIPatch(o["deploy_mode"]) {
			return fmt.Errorf("Error reading deploy_mode: %v", err)
		}
	}

	if err = d.Set("server_protocol", dataSourceFlattenLoadBalanceProfileServerProtocol(o["server_protocol"], d, "server_protocol")); err != nil {
		if !fortiAPIPatch(o["server_protocol"]) {
			return fmt.Errorf("Error reading server_protocol: %v", err)
		}
	}

	if err = d.Set("origin_host", dataSourceFlattenLoadBalanceProfileOriginHost(o["origin_host"], d, "origin_host")); err != nil {
		if !fortiAPIPatch(o["origin_host"]) {
			return fmt.Errorf("Error reading origin_host: %v", err)
		}
	}

	if err = d.Set("media_address", dataSourceFlattenLoadBalanceProfileMediaAddress(o["media_address"], d, "media_address")); err != nil {
		if !fortiAPIPatch(o["media_address"]) {
			return fmt.Errorf("Error reading media_address: %v", err)
		}
	}

	if err = d.Set("server_close_propagation", dataSourceFlattenLoadBalanceProfileServerClosePropagation(o["server_close_propagation"], d, "server_close_propagation")); err != nil {
		if !fortiAPIPatch(o["server_close_propagation"]) {
			return fmt.Errorf("Error reading server_close_propagation: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_session_after_fin", dataSourceFlattenLoadBalanceProfileTimeoutTcpSessionAfterFin(o["timeout_tcp_session_after_fin"], d, "timeout_tcp_session_after_fin")); err != nil {
		if !fortiAPIPatch(o["timeout_tcp_session_after_fin"]) {
			return fmt.Errorf("Error reading timeout_tcp_session_after_fin: %v", err)
		}
	}

	if err = d.Set("sip_max_size", dataSourceFlattenLoadBalanceProfileSipMaxSize(o["sip_max_size"], d, "sip_max_size")); err != nil {
		if !fortiAPIPatch(o["sip_max_size"]) {
			return fmt.Errorf("Error reading sip_max_size: %v", err)
		}
	}

	if err = d.Set("ssl_proxy", dataSourceFlattenLoadBalanceProfileSslProxy(o["ssl_proxy"], d, "ssl_proxy")); err != nil {
		if !fortiAPIPatch(o["ssl_proxy"]) {
			return fmt.Errorf("Error reading ssl_proxy: %v", err)
		}
	}

	if err = d.Set("use_tls_tickets", dataSourceFlattenLoadBalanceProfileUseTlsTickets(o["use_tls_tickets"], d, "use_tls_tickets")); err != nil {
		if !fortiAPIPatch(o["use_tls_tickets"]) {
			return fmt.Errorf("Error reading use_tls_tickets: %v", err)
		}
	}

	if err = d.Set("dns_cache_entry_size", dataSourceFlattenLoadBalanceProfileDnsCacheEntrySize(o["dns_cache_entry_size"], d, "dns_cache_entry_size")); err != nil {
		if !fortiAPIPatch(o["dns_cache_entry_size"]) {
			return fmt.Errorf("Error reading dns_cache_entry_size: %v", err)
		}
	}

	if err = d.Set("queue_timeout", dataSourceFlattenLoadBalanceProfileQueueTimeout(o["queue_timeout"], d, "queue_timeout")); err != nil {
		if !fortiAPIPatch(o["queue_timeout"]) {
			return fmt.Errorf("Error reading queue_timeout: %v", err)
		}
	}

	if err = d.Set("connect_timeout", dataSourceFlattenLoadBalanceProfileConnectTimeout(o["connect_timeout"], d, "connect_timeout")); err != nil {
		if !fortiAPIPatch(o["connect_timeout"]) {
			return fmt.Errorf("Error reading connect_timeout: %v", err)
		}
	}

	if err = d.Set("opt_trailer_hex", dataSourceFlattenLoadBalanceProfileOptTrailerHex(o["opt_trailer_hex"], d, "opt_trailer_hex")); err != nil {
		if !fortiAPIPatch(o["opt_trailer_hex"]) {
			return fmt.Errorf("Error reading opt_trailer_hex: %v", err)
		}
	}

	if err = d.Set("length_indicator_shift", dataSourceFlattenLoadBalanceProfileLengthIndicatorShift(o["length_indicator_shift"], d, "length_indicator_shift")); err != nil {
		if !fortiAPIPatch(o["length_indicator_shift"]) {
			return fmt.Errorf("Error reading length_indicator_shift: %v", err)
		}
	}

	if err = d.Set("ip_reputation_redirect", dataSourceFlattenLoadBalanceProfileIpReputationRedirect(o["ip_reputation_redirect"], d, "ip_reputation_redirect")); err != nil {
		if !fortiAPIPatch(o["ip_reputation_redirect"]) {
			return fmt.Errorf("Error reading ip_reputation_redirect: %v", err)
		}
	}

	if err = d.Set("length_indicator_size", dataSourceFlattenLoadBalanceProfileLengthIndicatorSize(o["length_indicator_size"], d, "length_indicator_size")); err != nil {
		if !fortiAPIPatch(o["length_indicator_size"]) {
			return fmt.Errorf("Error reading length_indicator_size: %v", err)
		}
	}

	if err = d.Set("http_x_forwarded_for", dataSourceFlattenLoadBalanceProfileHttpXForwardedFor(o["http_x_forwarded_for"], d, "http_x_forwarded_for")); err != nil {
		if !fortiAPIPatch(o["http_x_forwarded_for"]) {
			return fmt.Errorf("Error reading http_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("allow_ssl_versions", dataSourceFlattenLoadBalanceProfileAllowSslVersions(o["allow_ssl_versions"], d, "allow_ssl_versions")); err != nil {
		if !fortiAPIPatch(o["allow_ssl_versions"]) {
			return fmt.Errorf("Error reading allow_ssl_versions: %v", err)
		}
	}

	if err = d.Set("source_port", dataSourceFlattenLoadBalanceProfileSourcePort(o["source_port"], d, "source_port")); err != nil {
		if !fortiAPIPatch(o["source_port"]) {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("idle_timeout", dataSourceFlattenLoadBalanceProfileIdleTimeout(o["idle_timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle_timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("security_mode", dataSourceFlattenLoadBalanceProfileSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security-mode: %v", err)
		}
	}

	if err = d.Set("http_send_timeout", dataSourceFlattenLoadBalanceProfileHttpSendTimeout(o["http_send_timeout"], d, "http_send_timeout")); err != nil {
		if !fortiAPIPatch(o["http_send_timeout"]) {
			return fmt.Errorf("Error reading http_send_timeout: %v", err)
		}
	}

	if err = d.Set("dynamic_auth_port", dataSourceFlattenLoadBalanceProfileDynamicAuthPort(o["dynamic_auth_port"], d, "dynamic_auth_port")); err != nil {
		if !fortiAPIPatch(o["dynamic_auth_port"]) {
			return fmt.Errorf("Error reading dynamic_auth_port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", dataSourceFlattenLoadBalanceProfileSslAlgorithm(o["ssl_algorithm"], d, "ssl_algorithm")); err != nil {
		if !fortiAPIPatch(o["ssl_algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("failed_client_type", dataSourceFlattenLoadBalanceProfileFailedClientType(o["failed_client_type"], d, "failed_client_type")); err != nil {
		if !fortiAPIPatch(o["failed_client_type"]) {
			return fmt.Errorf("Error reading failed_client_type: %v", err)
		}
	}

	if err = d.Set("stateless", dataSourceFlattenLoadBalanceProfileStateless(o["stateless"], d, "stateless")); err != nil {
		if !fortiAPIPatch(o["stateless"]) {
			return fmt.Errorf("Error reading stateless: %v", err)
		}
	}

	if err = d.Set("server_timeout", dataSourceFlattenLoadBalanceProfileServerTimeout(o["server_timeout"], d, "server_timeout")); err != nil {
		if !fortiAPIPatch(o["server_timeout"]) {
			return fmt.Errorf("Error reading server_timeout: %v", err)
		}
	}

	if err = d.Set("cert_verify", dataSourceFlattenLoadBalanceProfileCertVerify(o["cert_verify"], d, "cert_verify")); err != nil {
		if !fortiAPIPatch(o["cert_verify"]) {
			return fmt.Errorf("Error reading cert_verify: %v", err)
		}
	}

	if err = d.Set("geoip_redirect", dataSourceFlattenLoadBalanceProfileGeoipRedirect(o["geoip_redirect"], d, "geoip_redirect")); err != nil {
		if !fortiAPIPatch(o["geoip_redirect"]) {
			return fmt.Errorf("Error reading geoip_redirect: %v", err)
		}
	}

	if err = d.Set("length_indicator_type", dataSourceFlattenLoadBalanceProfileLengthIndicatorType(o["length_indicator_type"], d, "length_indicator_type")); err != nil {
		if !fortiAPIPatch(o["length_indicator_type"]) {
			return fmt.Errorf("Error reading length_indicator_type: %v", err)
		}
	}

	return nil
}
