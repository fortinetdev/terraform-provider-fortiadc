// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceProfileRead,
		Update: resourceLoadBalanceProfileUpdate,
		Create: resourceLoadBalanceProfileCreate,
		Delete: resourceLoadBalanceProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"msg_encode_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ram_caching": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout_tcp_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_authenticate_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_max_query_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"decompression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"intermediate_ca_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_cache_ageout_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"smtp_disable_command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"failed_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward_client_certificate_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tune_bufsize": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"customized_ssl_ciphers_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"smtp_disable_command_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"failed_server_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout_radius_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_malform_query_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_sni_required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sip_dlg_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mysql_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"insert_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"opt_header_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_request_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_cache_response_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_headers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"origin_realm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"geoip_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"compression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"response_half_closed_request": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dynamic_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_max_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_reputation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"failed_client_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout_udp_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"customized_ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"starttls_active_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"smtp_domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_cache_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_cache_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"new_ssl_ciphers_long": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_x_forwarded_for_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"idle_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"deploy_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"origin_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"media_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_close_propagation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout_tcp_session_after_fin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sip_max_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_tls_tickets": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_cache_entry_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"queue_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connect_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"opt_trailer_hex": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"length_indicator_shift": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_reputation_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"length_indicator_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_send_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dynamic_auth_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"failed_client_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"stateless": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cert_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"geoip_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"length_indicator_type": &schema.Schema{
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
func resourceLoadBalanceProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceProfileRead(d, m)
}
func resourceLoadBalanceProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfile resource: %v", err)
	}

	return resourceLoadBalanceProfileRead(d, m)
}
func resourceLoadBalanceProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceProfileMsgEncodeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileWhitelist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileRamCaching(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileTimeoutTcpSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsAuthenticateFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsMaxQueryLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDecompression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileClientProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileIntermediateCaGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileVendorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsCacheAgeoutTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileTimeoutIpSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSmtpDisableCommand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileProductName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileFailedServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileForwardClientCertificateHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileTuneBufsize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileLocalCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileClientAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileCustomizedSslCiphersFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSmtpDisableCommandStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileClientKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileFailedServerStr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileTimeoutRadiusSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsMalformQueryAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileClientSniRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttpKeepaliveTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSipDlgTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileClientTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileMysqlMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileInsertClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileOptHeaderLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttpRequestTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileForwardClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsCacheResponseType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileMaxHttpHeaders(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileOriginRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileGeoipList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerKeepaliveTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileClientSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileCompression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileResponseHalfClosedRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDynamicAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerMaxSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileIpReputation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileFailedClientStr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileTimeoutUdpSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileLocalCertGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttp2Profile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileCustomizedSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileStarttlsActiveMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSmtpDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsCacheSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsCacheFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileNewSslCiphersLong(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttpXForwardedForHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileIdleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileMaxHeaderSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDeployMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileOriginHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileMediaAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerClosePropagation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileTimeoutTcpSessionAfterFin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSipMaxSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSslProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileUseTlsTickets(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDnsCacheEntrySize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileQueueTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileConnectTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileOptTrailerHex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileLengthIndicatorShift(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileIpReputationRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileLengthIndicatorSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttpXForwardedFor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileAllowSslVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSecurityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileHttpSendTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileDynamicAuthPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileSslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileFailedClientType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileStateless(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileServerTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileCertVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileGeoipRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileLengthIndicatorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("msg_encode_type", flattenLoadBalanceProfileMsgEncodeType(o["msg_encode_type"], d, "msg_encode_type", sv)); err != nil {
		if !fortiAPIPatch(o["msg_encode_type"]) {
			return fmt.Errorf("Error reading msg_encode_type: %v", err)
		}
	}

	if err = d.Set("whitelist", flattenLoadBalanceProfileWhitelist(o["whitelist"], d, "whitelist", sv)); err != nil {
		if !fortiAPIPatch(o["whitelist"]) {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if err = d.Set("ram_caching", flattenLoadBalanceProfileRamCaching(o["ram_caching"], d, "ram_caching", sv)); err != nil {
		if !fortiAPIPatch(o["ram_caching"]) {
			return fmt.Errorf("Error reading ram_caching: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_session", flattenLoadBalanceProfileTimeoutTcpSession(o["timeout_tcp_session"], d, "timeout_tcp_session", sv)); err != nil {
		if !fortiAPIPatch(o["timeout_tcp_session"]) {
			return fmt.Errorf("Error reading timeout_tcp_session: %v", err)
		}
	}

	if err = d.Set("dns_authenticate_flag", flattenLoadBalanceProfileDnsAuthenticateFlag(o["dns_authenticate_flag"], d, "dns_authenticate_flag", sv)); err != nil {
		if !fortiAPIPatch(o["dns_authenticate_flag"]) {
			return fmt.Errorf("Error reading dns_authenticate_flag: %v", err)
		}
	}

	if err = d.Set("dns_max_query_length", flattenLoadBalanceProfileDnsMaxQueryLength(o["dns_max_query_length"], d, "dns_max_query_length", sv)); err != nil {
		if !fortiAPIPatch(o["dns_max_query_length"]) {
			return fmt.Errorf("Error reading dns_max_query_length: %v", err)
		}
	}

	if err = d.Set("decompression", flattenLoadBalanceProfileDecompression(o["decompression"], d, "decompression", sv)); err != nil {
		if !fortiAPIPatch(o["decompression"]) {
			return fmt.Errorf("Error reading decompression: %v", err)
		}
	}

	if err = d.Set("client_protocol", flattenLoadBalanceProfileClientProtocol(o["client_protocol"], d, "client_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["client_protocol"]) {
			return fmt.Errorf("Error reading client_protocol: %v", err)
		}
	}

	if err = d.Set("intermediate_ca_group", flattenLoadBalanceProfileIntermediateCaGroup(o["intermediate_ca_group"], d, "intermediate_ca_group", sv)); err != nil {
		if !fortiAPIPatch(o["intermediate_ca_group"]) {
			return fmt.Errorf("Error reading intermediate_ca_group: %v", err)
		}
	}

	if err = d.Set("vendor_id", flattenLoadBalanceProfileVendorId(o["vendor_id"], d, "vendor_id", sv)); err != nil {
		if !fortiAPIPatch(o["vendor_id"]) {
			return fmt.Errorf("Error reading vendor_id: %v", err)
		}
	}

	if err = d.Set("dns_cache_ageout_time", flattenLoadBalanceProfileDnsCacheAgeoutTime(o["dns_cache_ageout_time"], d, "dns_cache_ageout_time", sv)); err != nil {
		if !fortiAPIPatch(o["dns_cache_ageout_time"]) {
			return fmt.Errorf("Error reading dns_cache_ageout_time: %v", err)
		}
	}

	if err = d.Set("ssl", flattenLoadBalanceProfileSsl(o["ssl"], d, "ssl", sv)); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("timeout_ip_session", flattenLoadBalanceProfileTimeoutIpSession(o["timeout_ip_session"], d, "timeout_ip_session", sv)); err != nil {
		if !fortiAPIPatch(o["timeout_ip_session"]) {
			return fmt.Errorf("Error reading timeout_ip_session: %v", err)
		}
	}

	if err = d.Set("smtp_disable_command", flattenLoadBalanceProfileSmtpDisableCommand(o["smtp_disable_command"], d, "smtp_disable_command", sv)); err != nil {
		if !fortiAPIPatch(o["smtp_disable_command"]) {
			return fmt.Errorf("Error reading smtp_disable_command: %v", err)
		}
	}

	if err = d.Set("server_keepalive", flattenLoadBalanceProfileServerKeepalive(o["server_keepalive"], d, "server_keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["server_keepalive"]) {
			return fmt.Errorf("Error reading server_keepalive: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers", flattenLoadBalanceProfileSslCiphers(o["ssl_ciphers"], d, "ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_ciphers"]) {
			return fmt.Errorf("Error reading ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("product_name", flattenLoadBalanceProfileProductName(o["product_name"], d, "product_name", sv)); err != nil {
		if !fortiAPIPatch(o["product_name"]) {
			return fmt.Errorf("Error reading product_name: %v", err)
		}
	}

	if err = d.Set("failed_server_type", flattenLoadBalanceProfileFailedServerType(o["failed_server_type"], d, "failed_server_type", sv)); err != nil {
		if !fortiAPIPatch(o["failed_server_type"]) {
			return fmt.Errorf("Error reading failed_server_type: %v", err)
		}
	}

	if err = d.Set("forward_client_certificate_header", flattenLoadBalanceProfileForwardClientCertificateHeader(o["forward_client_certificate_header"], d, "forward_client_certificate_header", sv)); err != nil {
		if !fortiAPIPatch(o["forward_client_certificate_header"]) {
			return fmt.Errorf("Error reading forward_client_certificate_header: %v", err)
		}
	}

	if err = d.Set("tune_bufsize", flattenLoadBalanceProfileTuneBufsize(o["tune-bufsize"], d, "tune_bufsize", sv)); err != nil {
		if !fortiAPIPatch(o["tune-bufsize"]) {
			return fmt.Errorf("Error reading tune-bufsize: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenLoadBalanceProfileLocalCert(o["local_cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local_cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("client_address", flattenLoadBalanceProfileClientAddress(o["client_address"], d, "client_address", sv)); err != nil {
		if !fortiAPIPatch(o["client_address"]) {
			return fmt.Errorf("Error reading client_address: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers_flag", flattenLoadBalanceProfileCustomizedSslCiphersFlag(o["customized_ssl_ciphers_flag"], d, "customized_ssl_ciphers_flag", sv)); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers_flag"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers_flag: %v", err)
		}
	}

	if err = d.Set("smtp_disable_command_status", flattenLoadBalanceProfileSmtpDisableCommandStatus(o["smtp_disable_command_status"], d, "smtp_disable_command_status", sv)); err != nil {
		if !fortiAPIPatch(o["smtp_disable_command_status"]) {
			return fmt.Errorf("Error reading smtp_disable_command_status: %v", err)
		}
	}

	if err = d.Set("client_keepalive", flattenLoadBalanceProfileClientKeepalive(o["client_keepalive"], d, "client_keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["client_keepalive"]) {
			return fmt.Errorf("Error reading client_keepalive: %v", err)
		}
	}

	if err = d.Set("failed_server_str", flattenLoadBalanceProfileFailedServerStr(o["failed_server_str"], d, "failed_server_str", sv)); err != nil {
		if !fortiAPIPatch(o["failed_server_str"]) {
			return fmt.Errorf("Error reading failed_server_str: %v", err)
		}
	}

	if err = d.Set("timeout_radius_session", flattenLoadBalanceProfileTimeoutRadiusSession(o["timeout_radius_session"], d, "timeout_radius_session", sv)); err != nil {
		if !fortiAPIPatch(o["timeout_radius_session"]) {
			return fmt.Errorf("Error reading timeout_radius_session: %v", err)
		}
	}

	if err = d.Set("dns_malform_query_action", flattenLoadBalanceProfileDnsMalformQueryAction(o["dns_malform_query_action"], d, "dns_malform_query_action", sv)); err != nil {
		if !fortiAPIPatch(o["dns_malform_query_action"]) {
			return fmt.Errorf("Error reading dns_malform_query_action: %v", err)
		}
	}

	if err = d.Set("client_sni_required", flattenLoadBalanceProfileClientSniRequired(o["client_sni_required"], d, "client_sni_required", sv)); err != nil {
		if !fortiAPIPatch(o["client_sni_required"]) {
			return fmt.Errorf("Error reading client_sni_required: %v", err)
		}
	}

	if err = d.Set("http_keepalive_timeout", flattenLoadBalanceProfileHttpKeepaliveTimeout(o["http_keepalive_timeout"], d, "http_keepalive_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["http_keepalive_timeout"]) {
			return fmt.Errorf("Error reading http_keepalive_timeout: %v", err)
		}
	}

	if err = d.Set("sip_dlg_timeout", flattenLoadBalanceProfileSipDlgTimeout(o["sip_dlg_timeout"], d, "sip_dlg_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["sip_dlg_timeout"]) {
			return fmt.Errorf("Error reading sip_dlg_timeout: %v", err)
		}
	}

	if err = d.Set("client_timeout", flattenLoadBalanceProfileClientTimeout(o["client_timeout"], d, "client_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["client_timeout"]) {
			return fmt.Errorf("Error reading client_timeout: %v", err)
		}
	}

	if err = d.Set("mysql_mode", flattenLoadBalanceProfileMysqlMode(o["mysql_mode"], d, "mysql_mode", sv)); err != nil {
		if !fortiAPIPatch(o["mysql_mode"]) {
			return fmt.Errorf("Error reading mysql_mode: %v", err)
		}
	}

	if err = d.Set("insert_client_ip", flattenLoadBalanceProfileInsertClientIp(o["insert_client_ip"], d, "insert_client_ip", sv)); err != nil {
		if !fortiAPIPatch(o["insert_client_ip"]) {
			return fmt.Errorf("Error reading insert_client_ip: %v", err)
		}
	}

	if err = d.Set("opt_header_length", flattenLoadBalanceProfileOptHeaderLength(o["opt_header_length"], d, "opt_header_length", sv)); err != nil {
		if !fortiAPIPatch(o["opt_header_length"]) {
			return fmt.Errorf("Error reading opt_header_length: %v", err)
		}
	}

	if err = d.Set("http_request_timeout", flattenLoadBalanceProfileHttpRequestTimeout(o["http_request_timeout"], d, "http_request_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["http_request_timeout"]) {
			return fmt.Errorf("Error reading http_request_timeout: %v", err)
		}
	}

	if err = d.Set("http_mode", flattenLoadBalanceProfileHttpMode(o["http_mode"], d, "http_mode", sv)); err != nil {
		if !fortiAPIPatch(o["http_mode"]) {
			return fmt.Errorf("Error reading http_mode: %v", err)
		}
	}

	if err = d.Set("forward_client_certificate", flattenLoadBalanceProfileForwardClientCertificate(o["forward_client_certificate"], d, "forward_client_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["forward_client_certificate"]) {
			return fmt.Errorf("Error reading forward_client_certificate: %v", err)
		}
	}

	if err = d.Set("dns_cache_response_type", flattenLoadBalanceProfileDnsCacheResponseType(o["dns_cache_response_type"], d, "dns_cache_response_type", sv)); err != nil {
		if !fortiAPIPatch(o["dns_cache_response_type"]) {
			return fmt.Errorf("Error reading dns_cache_response_type: %v", err)
		}
	}

	if err = d.Set("max_http_headers", flattenLoadBalanceProfileMaxHttpHeaders(o["max_http_headers"], d, "max_http_headers", sv)); err != nil {
		if !fortiAPIPatch(o["max_http_headers"]) {
			return fmt.Errorf("Error reading max_http_headers: %v", err)
		}
	}

	if err = d.Set("origin_realm", flattenLoadBalanceProfileOriginRealm(o["origin_realm"], d, "origin_realm", sv)); err != nil {
		if !fortiAPIPatch(o["origin_realm"]) {
			return fmt.Errorf("Error reading origin_realm: %v", err)
		}
	}

	if err = d.Set("geoip_list", flattenLoadBalanceProfileGeoipList(o["geoip_list"], d, "geoip_list", sv)); err != nil {
		if !fortiAPIPatch(o["geoip_list"]) {
			return fmt.Errorf("Error reading geoip_list: %v", err)
		}
	}

	if err = d.Set("server_keepalive_timeout", flattenLoadBalanceProfileServerKeepaliveTimeout(o["server_keepalive_timeout"], d, "server_keepalive_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["server_keepalive_timeout"]) {
			return fmt.Errorf("Error reading server_keepalive_timeout: %v", err)
		}
	}

	if err = d.Set("client_ssl", flattenLoadBalanceProfileClientSsl(o["client_ssl"], d, "client_ssl", sv)); err != nil {
		if !fortiAPIPatch(o["client_ssl"]) {
			return fmt.Errorf("Error reading client_ssl: %v", err)
		}
	}

	if err = d.Set("compression", flattenLoadBalanceProfileCompression(o["compression"], d, "compression", sv)); err != nil {
		if !fortiAPIPatch(o["compression"]) {
			return fmt.Errorf("Error reading compression: %v", err)
		}
	}

	if err = d.Set("response_half_closed_request", flattenLoadBalanceProfileResponseHalfClosedRequest(o["response_half_closed_request"], d, "response_half_closed_request", sv)); err != nil {
		if !fortiAPIPatch(o["response_half_closed_request"]) {
			return fmt.Errorf("Error reading response_half_closed_request: %v", err)
		}
	}

	if err = d.Set("dynamic_auth", flattenLoadBalanceProfileDynamicAuth(o["dynamic_auth"], d, "dynamic_auth", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic_auth"]) {
			return fmt.Errorf("Error reading dynamic_auth: %v", err)
		}
	}

	if err = d.Set("server_max_size", flattenLoadBalanceProfileServerMaxSize(o["server_max_size"], d, "server_max_size", sv)); err != nil {
		if !fortiAPIPatch(o["server_max_size"]) {
			return fmt.Errorf("Error reading server_max_size: %v", err)
		}
	}

	if err = d.Set("ip_reputation", flattenLoadBalanceProfileIpReputation(o["ip_reputation"], d, "ip_reputation", sv)); err != nil {
		if !fortiAPIPatch(o["ip_reputation"]) {
			return fmt.Errorf("Error reading ip_reputation: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceProfileType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("failed_client_str", flattenLoadBalanceProfileFailedClientStr(o["failed_client_str"], d, "failed_client_str", sv)); err != nil {
		if !fortiAPIPatch(o["failed_client_str"]) {
			return fmt.Errorf("Error reading failed_client_str: %v", err)
		}
	}

	if err = d.Set("timeout_udp_session", flattenLoadBalanceProfileTimeoutUdpSession(o["timeout_udp_session"], d, "timeout_udp_session", sv)); err != nil {
		if !fortiAPIPatch(o["timeout_udp_session"]) {
			return fmt.Errorf("Error reading timeout_udp_session: %v", err)
		}
	}

	if err = d.Set("local_cert_group", flattenLoadBalanceProfileLocalCertGroup(o["local_cert_group"], d, "local_cert_group", sv)); err != nil {
		if !fortiAPIPatch(o["local_cert_group"]) {
			return fmt.Errorf("Error reading local_cert_group: %v", err)
		}
	}

	if err = d.Set("http2_profile", flattenLoadBalanceProfileHttp2Profile(o["http2_profile"], d, "http2_profile", sv)); err != nil {
		if !fortiAPIPatch(o["http2_profile"]) {
			return fmt.Errorf("Error reading http2_profile: %v", err)
		}
	}

	if err = d.Set("customized_ssl_ciphers", flattenLoadBalanceProfileCustomizedSslCiphers(o["customized_ssl_ciphers"], d, "customized_ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["customized_ssl_ciphers"]) {
			return fmt.Errorf("Error reading customized_ssl_ciphers: %v", err)
		}
	}

	if err = d.Set("starttls_active_mode", flattenLoadBalanceProfileStarttlsActiveMode(o["starttls_active_mode"], d, "starttls_active_mode", sv)); err != nil {
		if !fortiAPIPatch(o["starttls_active_mode"]) {
			return fmt.Errorf("Error reading starttls_active_mode: %v", err)
		}
	}

	if err = d.Set("server_age", flattenLoadBalanceProfileServerAge(o["server_age"], d, "server_age", sv)); err != nil {
		if !fortiAPIPatch(o["server_age"]) {
			return fmt.Errorf("Error reading server_age: %v", err)
		}
	}

	if err = d.Set("smtp_domain_name", flattenLoadBalanceProfileSmtpDomainName(o["smtp_domain_name"], d, "smtp_domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["smtp_domain_name"]) {
			return fmt.Errorf("Error reading smtp_domain_name: %v", err)
		}
	}

	if err = d.Set("dns_cache_size", flattenLoadBalanceProfileDnsCacheSize(o["dns_cache_size"], d, "dns_cache_size", sv)); err != nil {
		if !fortiAPIPatch(o["dns_cache_size"]) {
			return fmt.Errorf("Error reading dns_cache_size: %v", err)
		}
	}

	if err = d.Set("dns_cache_flag", flattenLoadBalanceProfileDnsCacheFlag(o["dns_cache_flag"], d, "dns_cache_flag", sv)); err != nil {
		if !fortiAPIPatch(o["dns_cache_flag"]) {
			return fmt.Errorf("Error reading dns_cache_flag: %v", err)
		}
	}

	if err = d.Set("new_ssl_ciphers_long", flattenLoadBalanceProfileNewSslCiphersLong(o["new_ssl_ciphers_long"], d, "new_ssl_ciphers_long", sv)); err != nil {
		if !fortiAPIPatch(o["new_ssl_ciphers_long"]) {
			return fmt.Errorf("Error reading new_ssl_ciphers_long: %v", err)
		}
	}

	if err = d.Set("http_x_forwarded_for_header", flattenLoadBalanceProfileHttpXForwardedForHeader(o["http_x_forwarded_for_header"], d, "http_x_forwarded_for_header", sv)); err != nil {
		if !fortiAPIPatch(o["http_x_forwarded_for_header"]) {
			return fmt.Errorf("Error reading http_x_forwarded_for_header: %v", err)
		}
	}

	if err = d.Set("idle_time", flattenLoadBalanceProfileIdleTime(o["idle_time"], d, "idle_time", sv)); err != nil {
		if !fortiAPIPatch(o["idle_time"]) {
			return fmt.Errorf("Error reading idle_time: %v", err)
		}
	}

	if err = d.Set("max_header_size", flattenLoadBalanceProfileMaxHeaderSize(o["max_header_size"], d, "max_header_size", sv)); err != nil {
		if !fortiAPIPatch(o["max_header_size"]) {
			return fmt.Errorf("Error reading max_header_size: %v", err)
		}
	}

	if err = d.Set("deploy_mode", flattenLoadBalanceProfileDeployMode(o["deploy_mode"], d, "deploy_mode", sv)); err != nil {
		if !fortiAPIPatch(o["deploy_mode"]) {
			return fmt.Errorf("Error reading deploy_mode: %v", err)
		}
	}

	if err = d.Set("server_protocol", flattenLoadBalanceProfileServerProtocol(o["server_protocol"], d, "server_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["server_protocol"]) {
			return fmt.Errorf("Error reading server_protocol: %v", err)
		}
	}

	if err = d.Set("origin_host", flattenLoadBalanceProfileOriginHost(o["origin_host"], d, "origin_host", sv)); err != nil {
		if !fortiAPIPatch(o["origin_host"]) {
			return fmt.Errorf("Error reading origin_host: %v", err)
		}
	}

	if err = d.Set("media_address", flattenLoadBalanceProfileMediaAddress(o["media_address"], d, "media_address", sv)); err != nil {
		if !fortiAPIPatch(o["media_address"]) {
			return fmt.Errorf("Error reading media_address: %v", err)
		}
	}

	if err = d.Set("server_close_propagation", flattenLoadBalanceProfileServerClosePropagation(o["server_close_propagation"], d, "server_close_propagation", sv)); err != nil {
		if !fortiAPIPatch(o["server_close_propagation"]) {
			return fmt.Errorf("Error reading server_close_propagation: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_session_after_fin", flattenLoadBalanceProfileTimeoutTcpSessionAfterFin(o["timeout_tcp_session_after_fin"], d, "timeout_tcp_session_after_fin", sv)); err != nil {
		if !fortiAPIPatch(o["timeout_tcp_session_after_fin"]) {
			return fmt.Errorf("Error reading timeout_tcp_session_after_fin: %v", err)
		}
	}

	if err = d.Set("sip_max_size", flattenLoadBalanceProfileSipMaxSize(o["sip_max_size"], d, "sip_max_size", sv)); err != nil {
		if !fortiAPIPatch(o["sip_max_size"]) {
			return fmt.Errorf("Error reading sip_max_size: %v", err)
		}
	}

	if err = d.Set("ssl_proxy", flattenLoadBalanceProfileSslProxy(o["ssl_proxy"], d, "ssl_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_proxy"]) {
			return fmt.Errorf("Error reading ssl_proxy: %v", err)
		}
	}

	if err = d.Set("use_tls_tickets", flattenLoadBalanceProfileUseTlsTickets(o["use_tls_tickets"], d, "use_tls_tickets", sv)); err != nil {
		if !fortiAPIPatch(o["use_tls_tickets"]) {
			return fmt.Errorf("Error reading use_tls_tickets: %v", err)
		}
	}

	if err = d.Set("dns_cache_entry_size", flattenLoadBalanceProfileDnsCacheEntrySize(o["dns_cache_entry_size"], d, "dns_cache_entry_size", sv)); err != nil {
		if !fortiAPIPatch(o["dns_cache_entry_size"]) {
			return fmt.Errorf("Error reading dns_cache_entry_size: %v", err)
		}
	}

	if err = d.Set("queue_timeout", flattenLoadBalanceProfileQueueTimeout(o["queue_timeout"], d, "queue_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["queue_timeout"]) {
			return fmt.Errorf("Error reading queue_timeout: %v", err)
		}
	}

	if err = d.Set("connect_timeout", flattenLoadBalanceProfileConnectTimeout(o["connect_timeout"], d, "connect_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["connect_timeout"]) {
			return fmt.Errorf("Error reading connect_timeout: %v", err)
		}
	}

	if err = d.Set("opt_trailer_hex", flattenLoadBalanceProfileOptTrailerHex(o["opt_trailer_hex"], d, "opt_trailer_hex", sv)); err != nil {
		if !fortiAPIPatch(o["opt_trailer_hex"]) {
			return fmt.Errorf("Error reading opt_trailer_hex: %v", err)
		}
	}

	if err = d.Set("length_indicator_shift", flattenLoadBalanceProfileLengthIndicatorShift(o["length_indicator_shift"], d, "length_indicator_shift", sv)); err != nil {
		if !fortiAPIPatch(o["length_indicator_shift"]) {
			return fmt.Errorf("Error reading length_indicator_shift: %v", err)
		}
	}

	if err = d.Set("ip_reputation_redirect", flattenLoadBalanceProfileIpReputationRedirect(o["ip_reputation_redirect"], d, "ip_reputation_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["ip_reputation_redirect"]) {
			return fmt.Errorf("Error reading ip_reputation_redirect: %v", err)
		}
	}

	if err = d.Set("length_indicator_size", flattenLoadBalanceProfileLengthIndicatorSize(o["length_indicator_size"], d, "length_indicator_size", sv)); err != nil {
		if !fortiAPIPatch(o["length_indicator_size"]) {
			return fmt.Errorf("Error reading length_indicator_size: %v", err)
		}
	}

	if err = d.Set("http_x_forwarded_for", flattenLoadBalanceProfileHttpXForwardedFor(o["http_x_forwarded_for"], d, "http_x_forwarded_for", sv)); err != nil {
		if !fortiAPIPatch(o["http_x_forwarded_for"]) {
			return fmt.Errorf("Error reading http_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("allow_ssl_versions", flattenLoadBalanceProfileAllowSslVersions(o["allow_ssl_versions"], d, "allow_ssl_versions", sv)); err != nil {
		if !fortiAPIPatch(o["allow_ssl_versions"]) {
			return fmt.Errorf("Error reading allow_ssl_versions: %v", err)
		}
	}

	if err = d.Set("source_port", flattenLoadBalanceProfileSourcePort(o["source_port"], d, "source_port", sv)); err != nil {
		if !fortiAPIPatch(o["source_port"]) {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenLoadBalanceProfileIdleTimeout(o["idle_timeout"], d, "idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["idle_timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenLoadBalanceProfileSecurityMode(o["security-mode"], d, "security_mode", sv)); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security-mode: %v", err)
		}
	}

	if err = d.Set("http_send_timeout", flattenLoadBalanceProfileHttpSendTimeout(o["http_send_timeout"], d, "http_send_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["http_send_timeout"]) {
			return fmt.Errorf("Error reading http_send_timeout: %v", err)
		}
	}

	if err = d.Set("dynamic_auth_port", flattenLoadBalanceProfileDynamicAuthPort(o["dynamic_auth_port"], d, "dynamic_auth_port", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic_auth_port"]) {
			return fmt.Errorf("Error reading dynamic_auth_port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenLoadBalanceProfileSslAlgorithm(o["ssl_algorithm"], d, "ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("failed_client_type", flattenLoadBalanceProfileFailedClientType(o["failed_client_type"], d, "failed_client_type", sv)); err != nil {
		if !fortiAPIPatch(o["failed_client_type"]) {
			return fmt.Errorf("Error reading failed_client_type: %v", err)
		}
	}

	if err = d.Set("stateless", flattenLoadBalanceProfileStateless(o["stateless"], d, "stateless", sv)); err != nil {
		if !fortiAPIPatch(o["stateless"]) {
			return fmt.Errorf("Error reading stateless: %v", err)
		}
	}

	if err = d.Set("server_timeout", flattenLoadBalanceProfileServerTimeout(o["server_timeout"], d, "server_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["server_timeout"]) {
			return fmt.Errorf("Error reading server_timeout: %v", err)
		}
	}

	if err = d.Set("cert_verify", flattenLoadBalanceProfileCertVerify(o["cert_verify"], d, "cert_verify", sv)); err != nil {
		if !fortiAPIPatch(o["cert_verify"]) {
			return fmt.Errorf("Error reading cert_verify: %v", err)
		}
	}

	if err = d.Set("geoip_redirect", flattenLoadBalanceProfileGeoipRedirect(o["geoip_redirect"], d, "geoip_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["geoip_redirect"]) {
			return fmt.Errorf("Error reading geoip_redirect: %v", err)
		}
	}

	if err = d.Set("length_indicator_type", flattenLoadBalanceProfileLengthIndicatorType(o["length_indicator_type"], d, "length_indicator_type", sv)); err != nil {
		if !fortiAPIPatch(o["length_indicator_type"]) {
			return fmt.Errorf("Error reading length_indicator_type: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceProfileMsgEncodeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileWhitelist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileRamCaching(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileTimeoutTcpSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsAuthenticateFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsMaxQueryLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDecompression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileClientProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileIntermediateCaGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileVendorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsCacheAgeoutTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileTimeoutIpSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSmtpDisableCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileProductName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileFailedServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileForwardClientCertificateHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileTuneBufsize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileLocalCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileClientAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileCustomizedSslCiphersFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSmtpDisableCommandStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileClientKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileFailedServerStr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileTimeoutRadiusSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsMalformQueryAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileClientSniRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttpKeepaliveTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSipDlgTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileClientTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileMysqlMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileInsertClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileOptHeaderLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttpRequestTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileForwardClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsCacheResponseType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileMaxHttpHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileOriginRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileGeoipList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerKeepaliveTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileClientSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileCompression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileResponseHalfClosedRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDynamicAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerMaxSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileIpReputation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileFailedClientStr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileTimeoutUdpSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileLocalCertGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttp2Profile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileCustomizedSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileStarttlsActiveMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSmtpDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsCacheSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsCacheFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileNewSslCiphersLong(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttpXForwardedForHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileIdleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileMaxHeaderSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDeployMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileOriginHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileMediaAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerClosePropagation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileTimeoutTcpSessionAfterFin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSipMaxSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSslProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileUseTlsTickets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDnsCacheEntrySize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileQueueTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileConnectTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileOptTrailerHex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileLengthIndicatorShift(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileIpReputationRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileLengthIndicatorSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttpXForwardedFor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileAllowSslVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSecurityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileHttpSendTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileDynamicAuthPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileSslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileFailedClientType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileStateless(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileServerTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileCertVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileGeoipRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileLengthIndicatorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("msg_encode_type"); ok {
		t, err := expandLoadBalanceProfileMsgEncodeType(d, v, "msg_encode_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg_encode_type"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok {
		t, err := expandLoadBalanceProfileWhitelist(d, v, "whitelist", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	if v, ok := d.GetOk("ram_caching"); ok {
		t, err := expandLoadBalanceProfileRamCaching(d, v, "ram_caching", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ram_caching"] = t
		}
	}

	if v, ok := d.GetOk("timeout_tcp_session"); ok {
		t, err := expandLoadBalanceProfileTimeoutTcpSession(d, v, "timeout_tcp_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout_tcp_session"] = t
		}
	}

	if v, ok := d.GetOk("dns_authenticate_flag"); ok {
		t, err := expandLoadBalanceProfileDnsAuthenticateFlag(d, v, "dns_authenticate_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_authenticate_flag"] = t
		}
	}

	if v, ok := d.GetOk("dns_max_query_length"); ok {
		t, err := expandLoadBalanceProfileDnsMaxQueryLength(d, v, "dns_max_query_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_max_query_length"] = t
		}
	}

	if v, ok := d.GetOk("decompression"); ok {
		t, err := expandLoadBalanceProfileDecompression(d, v, "decompression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decompression"] = t
		}
	}

	if v, ok := d.GetOk("client_protocol"); ok {
		t, err := expandLoadBalanceProfileClientProtocol(d, v, "client_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_protocol"] = t
		}
	}

	if v, ok := d.GetOk("intermediate_ca_group"); ok {
		t, err := expandLoadBalanceProfileIntermediateCaGroup(d, v, "intermediate_ca_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intermediate_ca_group"] = t
		}
	}

	if v, ok := d.GetOk("vendor_id"); ok {
		t, err := expandLoadBalanceProfileVendorId(d, v, "vendor_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor_id"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_ageout_time"); ok {
		t, err := expandLoadBalanceProfileDnsCacheAgeoutTime(d, v, "dns_cache_ageout_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_cache_ageout_time"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandLoadBalanceProfileSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("timeout_ip_session"); ok {
		t, err := expandLoadBalanceProfileTimeoutIpSession(d, v, "timeout_ip_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout_ip_session"] = t
		}
	}

	if v, ok := d.GetOk("smtp_disable_command"); ok {
		t, err := expandLoadBalanceProfileSmtpDisableCommand(d, v, "smtp_disable_command", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp_disable_command"] = t
		}
	}

	if v, ok := d.GetOk("server_keepalive"); ok {
		t, err := expandLoadBalanceProfileServerKeepalive(d, v, "server_keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_keepalive"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ciphers"); ok {
		t, err := expandLoadBalanceProfileSslCiphers(d, v, "ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("product_name"); ok {
		t, err := expandLoadBalanceProfileProductName(d, v, "product_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["product_name"] = t
		}
	}

	if v, ok := d.GetOk("failed_server_type"); ok {
		t, err := expandLoadBalanceProfileFailedServerType(d, v, "failed_server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failed_server_type"] = t
		}
	}

	if v, ok := d.GetOk("forward_client_certificate_header"); ok {
		t, err := expandLoadBalanceProfileForwardClientCertificateHeader(d, v, "forward_client_certificate_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_client_certificate_header"] = t
		}
	}

	if v, ok := d.GetOk("tune_bufsize"); ok {
		t, err := expandLoadBalanceProfileTuneBufsize(d, v, "tune_bufsize", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tune-bufsize"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandLoadBalanceProfileLocalCert(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local_cert"] = t
		}
	}

	if v, ok := d.GetOk("client_address"); ok {
		t, err := expandLoadBalanceProfileClientAddress(d, v, "client_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_address"] = t
		}
	}

	if v, ok := d.GetOk("customized_ssl_ciphers_flag"); ok {
		t, err := expandLoadBalanceProfileCustomizedSslCiphersFlag(d, v, "customized_ssl_ciphers_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customized_ssl_ciphers_flag"] = t
		}
	}

	if v, ok := d.GetOk("smtp_disable_command_status"); ok {
		t, err := expandLoadBalanceProfileSmtpDisableCommandStatus(d, v, "smtp_disable_command_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp_disable_command_status"] = t
		}
	}

	if v, ok := d.GetOk("client_keepalive"); ok {
		t, err := expandLoadBalanceProfileClientKeepalive(d, v, "client_keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_keepalive"] = t
		}
	}

	if v, ok := d.GetOk("failed_server_str"); ok {
		t, err := expandLoadBalanceProfileFailedServerStr(d, v, "failed_server_str", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failed_server_str"] = t
		}
	}

	if v, ok := d.GetOk("timeout_radius_session"); ok {
		t, err := expandLoadBalanceProfileTimeoutRadiusSession(d, v, "timeout_radius_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout_radius_session"] = t
		}
	}

	if v, ok := d.GetOk("dns_malform_query_action"); ok {
		t, err := expandLoadBalanceProfileDnsMalformQueryAction(d, v, "dns_malform_query_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_malform_query_action"] = t
		}
	}

	if v, ok := d.GetOk("client_sni_required"); ok {
		t, err := expandLoadBalanceProfileClientSniRequired(d, v, "client_sni_required", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_sni_required"] = t
		}
	}

	if v, ok := d.GetOk("http_keepalive_timeout"); ok {
		t, err := expandLoadBalanceProfileHttpKeepaliveTimeout(d, v, "http_keepalive_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_keepalive_timeout"] = t
		}
	}

	if v, ok := d.GetOk("sip_dlg_timeout"); ok {
		t, err := expandLoadBalanceProfileSipDlgTimeout(d, v, "sip_dlg_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip_dlg_timeout"] = t
		}
	}

	if v, ok := d.GetOk("client_timeout"); ok {
		t, err := expandLoadBalanceProfileClientTimeout(d, v, "client_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_timeout"] = t
		}
	}

	if v, ok := d.GetOk("mysql_mode"); ok {
		t, err := expandLoadBalanceProfileMysqlMode(d, v, "mysql_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mysql_mode"] = t
		}
	}

	if v, ok := d.GetOk("insert_client_ip"); ok {
		t, err := expandLoadBalanceProfileInsertClientIp(d, v, "insert_client_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["insert_client_ip"] = t
		}
	}

	if v, ok := d.GetOk("opt_header_length"); ok {
		t, err := expandLoadBalanceProfileOptHeaderLength(d, v, "opt_header_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opt_header_length"] = t
		}
	}

	if v, ok := d.GetOk("http_request_timeout"); ok {
		t, err := expandLoadBalanceProfileHttpRequestTimeout(d, v, "http_request_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_request_timeout"] = t
		}
	}

	if v, ok := d.GetOk("http_mode"); ok {
		t, err := expandLoadBalanceProfileHttpMode(d, v, "http_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_mode"] = t
		}
	}

	if v, ok := d.GetOk("forward_client_certificate"); ok {
		t, err := expandLoadBalanceProfileForwardClientCertificate(d, v, "forward_client_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_client_certificate"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_response_type"); ok {
		t, err := expandLoadBalanceProfileDnsCacheResponseType(d, v, "dns_cache_response_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_cache_response_type"] = t
		}
	}

	if v, ok := d.GetOk("max_http_headers"); ok {
		t, err := expandLoadBalanceProfileMaxHttpHeaders(d, v, "max_http_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_http_headers"] = t
		}
	}

	if v, ok := d.GetOk("origin_realm"); ok {
		t, err := expandLoadBalanceProfileOriginRealm(d, v, "origin_realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin_realm"] = t
		}
	}

	if v, ok := d.GetOk("geoip_list"); ok {
		t, err := expandLoadBalanceProfileGeoipList(d, v, "geoip_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip_list"] = t
		}
	}

	if v, ok := d.GetOk("server_keepalive_timeout"); ok {
		t, err := expandLoadBalanceProfileServerKeepaliveTimeout(d, v, "server_keepalive_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_keepalive_timeout"] = t
		}
	}

	if v, ok := d.GetOk("client_ssl"); ok {
		t, err := expandLoadBalanceProfileClientSsl(d, v, "client_ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_ssl"] = t
		}
	}

	if v, ok := d.GetOk("compression"); ok {
		t, err := expandLoadBalanceProfileCompression(d, v, "compression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compression"] = t
		}
	}

	if v, ok := d.GetOk("response_half_closed_request"); ok {
		t, err := expandLoadBalanceProfileResponseHalfClosedRequest(d, v, "response_half_closed_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response_half_closed_request"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_auth"); ok {
		t, err := expandLoadBalanceProfileDynamicAuth(d, v, "dynamic_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_auth"] = t
		}
	}

	if v, ok := d.GetOk("server_max_size"); ok {
		t, err := expandLoadBalanceProfileServerMaxSize(d, v, "server_max_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_max_size"] = t
		}
	}

	if v, ok := d.GetOk("ip_reputation"); ok {
		t, err := expandLoadBalanceProfileIpReputation(d, v, "ip_reputation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_reputation"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceProfileType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("failed_client_str"); ok {
		t, err := expandLoadBalanceProfileFailedClientStr(d, v, "failed_client_str", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failed_client_str"] = t
		}
	}

	if v, ok := d.GetOk("timeout_udp_session"); ok {
		t, err := expandLoadBalanceProfileTimeoutUdpSession(d, v, "timeout_udp_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout_udp_session"] = t
		}
	}

	if v, ok := d.GetOk("local_cert_group"); ok {
		t, err := expandLoadBalanceProfileLocalCertGroup(d, v, "local_cert_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local_cert_group"] = t
		}
	}

	if v, ok := d.GetOk("http2_profile"); ok {
		t, err := expandLoadBalanceProfileHttp2Profile(d, v, "http2_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2_profile"] = t
		}
	}

	if v, ok := d.GetOk("customized_ssl_ciphers"); ok {
		t, err := expandLoadBalanceProfileCustomizedSslCiphers(d, v, "customized_ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customized_ssl_ciphers"] = t
		}
	}

	if v, ok := d.GetOk("starttls_active_mode"); ok {
		t, err := expandLoadBalanceProfileStarttlsActiveMode(d, v, "starttls_active_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["starttls_active_mode"] = t
		}
	}

	if v, ok := d.GetOk("server_age"); ok {
		t, err := expandLoadBalanceProfileServerAge(d, v, "server_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_age"] = t
		}
	}

	if v, ok := d.GetOk("smtp_domain_name"); ok {
		t, err := expandLoadBalanceProfileSmtpDomainName(d, v, "smtp_domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp_domain_name"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_size"); ok {
		t, err := expandLoadBalanceProfileDnsCacheSize(d, v, "dns_cache_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_cache_size"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_flag"); ok {
		t, err := expandLoadBalanceProfileDnsCacheFlag(d, v, "dns_cache_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_cache_flag"] = t
		}
	}

	if v, ok := d.GetOk("new_ssl_ciphers_long"); ok {
		t, err := expandLoadBalanceProfileNewSslCiphersLong(d, v, "new_ssl_ciphers_long", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["new_ssl_ciphers_long"] = t
		}
	}

	if v, ok := d.GetOk("http_x_forwarded_for_header"); ok {
		t, err := expandLoadBalanceProfileHttpXForwardedForHeader(d, v, "http_x_forwarded_for_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_x_forwarded_for_header"] = t
		}
	}

	if v, ok := d.GetOk("idle_time"); ok {
		t, err := expandLoadBalanceProfileIdleTime(d, v, "idle_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle_time"] = t
		}
	}

	if v, ok := d.GetOk("max_header_size"); ok {
		t, err := expandLoadBalanceProfileMaxHeaderSize(d, v, "max_header_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_header_size"] = t
		}
	}

	if v, ok := d.GetOk("deploy_mode"); ok {
		t, err := expandLoadBalanceProfileDeployMode(d, v, "deploy_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deploy_mode"] = t
		}
	}

	if v, ok := d.GetOk("server_protocol"); ok {
		t, err := expandLoadBalanceProfileServerProtocol(d, v, "server_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_protocol"] = t
		}
	}

	if v, ok := d.GetOk("origin_host"); ok {
		t, err := expandLoadBalanceProfileOriginHost(d, v, "origin_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin_host"] = t
		}
	}

	if v, ok := d.GetOk("media_address"); ok {
		t, err := expandLoadBalanceProfileMediaAddress(d, v, "media_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["media_address"] = t
		}
	}

	if v, ok := d.GetOk("server_close_propagation"); ok {
		t, err := expandLoadBalanceProfileServerClosePropagation(d, v, "server_close_propagation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_close_propagation"] = t
		}
	}

	if v, ok := d.GetOk("timeout_tcp_session_after_fin"); ok {
		t, err := expandLoadBalanceProfileTimeoutTcpSessionAfterFin(d, v, "timeout_tcp_session_after_fin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout_tcp_session_after_fin"] = t
		}
	}

	if v, ok := d.GetOk("sip_max_size"); ok {
		t, err := expandLoadBalanceProfileSipMaxSize(d, v, "sip_max_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip_max_size"] = t
		}
	}

	if v, ok := d.GetOk("ssl_proxy"); ok {
		t, err := expandLoadBalanceProfileSslProxy(d, v, "ssl_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_proxy"] = t
		}
	}

	if v, ok := d.GetOk("use_tls_tickets"); ok {
		t, err := expandLoadBalanceProfileUseTlsTickets(d, v, "use_tls_tickets", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use_tls_tickets"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_entry_size"); ok {
		t, err := expandLoadBalanceProfileDnsCacheEntrySize(d, v, "dns_cache_entry_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_cache_entry_size"] = t
		}
	}

	if v, ok := d.GetOk("queue_timeout"); ok {
		t, err := expandLoadBalanceProfileQueueTimeout(d, v, "queue_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue_timeout"] = t
		}
	}

	if v, ok := d.GetOk("connect_timeout"); ok {
		t, err := expandLoadBalanceProfileConnectTimeout(d, v, "connect_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect_timeout"] = t
		}
	}

	if v, ok := d.GetOk("opt_trailer_hex"); ok {
		t, err := expandLoadBalanceProfileOptTrailerHex(d, v, "opt_trailer_hex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opt_trailer_hex"] = t
		}
	}

	if v, ok := d.GetOk("length_indicator_shift"); ok {
		t, err := expandLoadBalanceProfileLengthIndicatorShift(d, v, "length_indicator_shift", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["length_indicator_shift"] = t
		}
	}

	if v, ok := d.GetOk("ip_reputation_redirect"); ok {
		t, err := expandLoadBalanceProfileIpReputationRedirect(d, v, "ip_reputation_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip_reputation_redirect"] = t
		}
	}

	if v, ok := d.GetOk("length_indicator_size"); ok {
		t, err := expandLoadBalanceProfileLengthIndicatorSize(d, v, "length_indicator_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["length_indicator_size"] = t
		}
	}

	if v, ok := d.GetOk("http_x_forwarded_for"); ok {
		t, err := expandLoadBalanceProfileHttpXForwardedFor(d, v, "http_x_forwarded_for", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_x_forwarded_for"] = t
		}
	}

	if v, ok := d.GetOk("allow_ssl_versions"); ok {
		t, err := expandLoadBalanceProfileAllowSslVersions(d, v, "allow_ssl_versions", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow_ssl_versions"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok {
		t, err := expandLoadBalanceProfileSourcePort(d, v, "source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source_port"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok {
		t, err := expandLoadBalanceProfileIdleTimeout(d, v, "idle_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle_timeout"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {
		t, err := expandLoadBalanceProfileSecurityMode(d, v, "security_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_send_timeout"); ok {
		t, err := expandLoadBalanceProfileHttpSendTimeout(d, v, "http_send_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_send_timeout"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_auth_port"); ok {
		t, err := expandLoadBalanceProfileDynamicAuthPort(d, v, "dynamic_auth_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_auth_port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {
		t, err := expandLoadBalanceProfileSslAlgorithm(d, v, "ssl_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_algorithm"] = t
		}
	}

	if v, ok := d.GetOk("failed_client_type"); ok {
		t, err := expandLoadBalanceProfileFailedClientType(d, v, "failed_client_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failed_client_type"] = t
		}
	}

	if v, ok := d.GetOk("stateless"); ok {
		t, err := expandLoadBalanceProfileStateless(d, v, "stateless", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stateless"] = t
		}
	}

	if v, ok := d.GetOk("server_timeout"); ok {
		t, err := expandLoadBalanceProfileServerTimeout(d, v, "server_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_timeout"] = t
		}
	}

	if v, ok := d.GetOk("cert_verify"); ok {
		t, err := expandLoadBalanceProfileCertVerify(d, v, "cert_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert_verify"] = t
		}
	}

	if v, ok := d.GetOk("geoip_redirect"); ok {
		t, err := expandLoadBalanceProfileGeoipRedirect(d, v, "geoip_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip_redirect"] = t
		}
	}

	if v, ok := d.GetOk("length_indicator_type"); ok {
		t, err := expandLoadBalanceProfileLengthIndicatorType(d, v, "length_indicator_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["length_indicator_type"] = t
		}
	}

	return &obj, nil
}
