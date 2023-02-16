// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance virtual server.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceVirtualServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceVirtualServerRead,
		Schema: map[string]*schema.Schema{
			"scripting_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"packet_fwd_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"warmrate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http2https": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_page": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiview": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ztna_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"adfs_published_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_rewriting": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http2https_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_msg": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scripting_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_routing_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"one_click_gslb_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pagespeed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_mirror_intf": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dos_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_rewriting_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"l2_exception_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ips_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stream_scripting_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"clone_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_azure_lb_backend_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"clone_traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stream_scripting_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_ssl_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"warmup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_lb_backend": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trans_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"captcha_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_routing": &schema.Schema{
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
func dataSourceLoadBalanceVirtualServerRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceVirtualServer: type error")
	}

	o, err := c.ReadLoadBalanceVirtualServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceVirtualServer: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceVirtualServer(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceVirtualServer from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceVirtualServerScriptingList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPacketFwdMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerWarmrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerHttp2Https(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerErrorPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerZtnaProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAuthPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAdfsPublishedService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerContentRewriting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerHttp2HttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerErrorMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerScriptingFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerContentRoutingList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPublicIpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerOneClickGslbServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPagespeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerConnectionLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerHostName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerSslMirrorIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerSourcePoolList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerDosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerContentRewritingList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerL2ExceptionList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerIpsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerTrafficGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerStreamScriptingList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerClonePool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerUseAzureLbBackendIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerCloneTrafficType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAlone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerStreamScriptingFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerClientSslProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerSchedulePoolList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerSslMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerScheduleList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerWarmup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerConnectionPool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAzureLbBackend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerConnectionRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerTransRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerTrafficLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerPublicIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerCaptchaProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceVirtualServerContentRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceVirtualServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("scripting_list", dataSourceFlattenLoadBalanceVirtualServerScriptingList(o["scripting_list"], d, "scripting_list")); err != nil {
		if !fortiAPIPatch(o["scripting_list"]) {
			return fmt.Errorf("Error reading scripting_list: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenLoadBalanceVirtualServerStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("packet_fwd_method", dataSourceFlattenLoadBalanceVirtualServerPacketFwdMethod(o["packet-fwd-method"], d, "packet_fwd_method")); err != nil {
		if !fortiAPIPatch(o["packet-fwd-method"]) {
			return fmt.Errorf("Error reading packet-fwd-method: %v", err)
		}
	}

	if err = d.Set("warmrate", dataSourceFlattenLoadBalanceVirtualServerWarmrate(o["warmrate"], d, "warmrate")); err != nil {
		if !fortiAPIPatch(o["warmrate"]) {
			return fmt.Errorf("Error reading warmrate: %v", err)
		}
	}

	if err = d.Set("http2https", dataSourceFlattenLoadBalanceVirtualServerHttp2Https(o["http2https"], d, "http2https")); err != nil {
		if !fortiAPIPatch(o["http2https"]) {
			return fmt.Errorf("Error reading http2https: %v", err)
		}
	}

	if err = d.Set("pool", dataSourceFlattenLoadBalanceVirtualServerPool(o["pool"], d, "pool")); err != nil {
		if !fortiAPIPatch(o["pool"]) {
			return fmt.Errorf("Error reading pool: %v", err)
		}
	}

	if err = d.Set("error_page", dataSourceFlattenLoadBalanceVirtualServerErrorPage(o["error-page"], d, "error_page")); err != nil {
		if !fortiAPIPatch(o["error-page"]) {
			return fmt.Errorf("Error reading error-page: %v", err)
		}
	}

	if err = d.Set("fortiview", dataSourceFlattenLoadBalanceVirtualServerFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if !fortiAPIPatch(o["fortiview"]) {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("ztna_profile", dataSourceFlattenLoadBalanceVirtualServerZtnaProfile(o["ztna_profile"], d, "ztna_profile")); err != nil {
		if !fortiAPIPatch(o["ztna_profile"]) {
			return fmt.Errorf("Error reading ztna_profile: %v", err)
		}
	}

	if err = d.Set("auth_policy", dataSourceFlattenLoadBalanceVirtualServerAuthPolicy(o["auth_policy"], d, "auth_policy")); err != nil {
		if !fortiAPIPatch(o["auth_policy"]) {
			return fmt.Errorf("Error reading auth_policy: %v", err)
		}
	}

	if err = d.Set("adfs_published_service", dataSourceFlattenLoadBalanceVirtualServerAdfsPublishedService(o["adfs-published-service"], d, "adfs_published_service")); err != nil {
		if !fortiAPIPatch(o["adfs-published-service"]) {
			return fmt.Errorf("Error reading adfs-published-service: %v", err)
		}
	}

	if err = d.Set("content_rewriting", dataSourceFlattenLoadBalanceVirtualServerContentRewriting(o["content-rewriting"], d, "content_rewriting")); err != nil {
		if !fortiAPIPatch(o["content-rewriting"]) {
			return fmt.Errorf("Error reading content-rewriting: %v", err)
		}
	}

	if err = d.Set("http2https_port", dataSourceFlattenLoadBalanceVirtualServerHttp2HttpsPort(o["http2https-port"], d, "http2https_port")); err != nil {
		if !fortiAPIPatch(o["http2https-port"]) {
			return fmt.Errorf("Error reading http2https-port: %v", err)
		}
	}

	if err = d.Set("domain_name", dataSourceFlattenLoadBalanceVirtualServerDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain-name: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenLoadBalanceVirtualServerProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("error_msg", dataSourceFlattenLoadBalanceVirtualServerErrorMsg(o["error-msg"], d, "error_msg")); err != nil {
		if !fortiAPIPatch(o["error-msg"]) {
			return fmt.Errorf("Error reading error-msg: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenLoadBalanceVirtualServerPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("scripting_flag", dataSourceFlattenLoadBalanceVirtualServerScriptingFlag(o["scripting_flag"], d, "scripting_flag")); err != nil {
		if !fortiAPIPatch(o["scripting_flag"]) {
			return fmt.Errorf("Error reading scripting_flag: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenLoadBalanceVirtualServerComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("content_routing_list", dataSourceFlattenLoadBalanceVirtualServerContentRoutingList(o["content-routing-list"], d, "content_routing_list")); err != nil {
		if !fortiAPIPatch(o["content-routing-list"]) {
			return fmt.Errorf("Error reading content-routing-list: %v", err)
		}
	}

	if err = d.Set("wccp", dataSourceFlattenLoadBalanceVirtualServerWccp(o["wccp"], d, "wccp")); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("public_ip_type", dataSourceFlattenLoadBalanceVirtualServerPublicIpType(o["public-ip-type"], d, "public_ip_type")); err != nil {
		if !fortiAPIPatch(o["public-ip-type"]) {
			return fmt.Errorf("Error reading public-ip-type: %v", err)
		}
	}

	if err = d.Set("one_click_gslb_server", dataSourceFlattenLoadBalanceVirtualServerOneClickGslbServer(o["one-click-gslb-server"], d, "one_click_gslb_server")); err != nil {
		if !fortiAPIPatch(o["one-click-gslb-server"]) {
			return fmt.Errorf("Error reading one-click-gslb-server: %v", err)
		}
	}

	if err = d.Set("pagespeed", dataSourceFlattenLoadBalanceVirtualServerPagespeed(o["pagespeed"], d, "pagespeed")); err != nil {
		if !fortiAPIPatch(o["pagespeed"]) {
			return fmt.Errorf("Error reading pagespeed: %v", err)
		}
	}

	if err = d.Set("connection_limit", dataSourceFlattenLoadBalanceVirtualServerConnectionLimit(o["connection-limit"], d, "connection_limit")); err != nil {
		if !fortiAPIPatch(o["connection-limit"]) {
			return fmt.Errorf("Error reading connection-limit: %v", err)
		}
	}

	if err = d.Set("host_name", dataSourceFlattenLoadBalanceVirtualServerHostName(o["host-name"], d, "host_name")); err != nil {
		if !fortiAPIPatch(o["host-name"]) {
			return fmt.Errorf("Error reading host-name: %v", err)
		}
	}

	if err = d.Set("ssl_mirror_intf", dataSourceFlattenLoadBalanceVirtualServerSslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf")); err != nil {
		if !fortiAPIPatch(o["ssl-mirror-intf"]) {
			return fmt.Errorf("Error reading ssl-mirror-intf: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceVirtualServerMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("source_pool_list", dataSourceFlattenLoadBalanceVirtualServerSourcePoolList(o["source-pool-list"], d, "source_pool_list")); err != nil {
		if !fortiAPIPatch(o["source-pool-list"]) {
			return fmt.Errorf("Error reading source-pool-list: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceVirtualServerType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("method", dataSourceFlattenLoadBalanceVirtualServerMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("persistence", dataSourceFlattenLoadBalanceVirtualServerPersistence(o["persistence"], d, "persistence")); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("dos_profile", dataSourceFlattenLoadBalanceVirtualServerDosProfile(o["dos_profile"], d, "dos_profile")); err != nil {
		if !fortiAPIPatch(o["dos_profile"]) {
			return fmt.Errorf("Error reading dos_profile: %v", err)
		}
	}

	if err = d.Set("profile", dataSourceFlattenLoadBalanceVirtualServerProfile(o["profile"], d, "profile")); err != nil {
		if !fortiAPIPatch(o["profile"]) {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("content_rewriting_list", dataSourceFlattenLoadBalanceVirtualServerContentRewritingList(o["content-rewriting-list"], d, "content_rewriting_list")); err != nil {
		if !fortiAPIPatch(o["content-rewriting-list"]) {
			return fmt.Errorf("Error reading content-rewriting-list: %v", err)
		}
	}

	if err = d.Set("l2_exception_list", dataSourceFlattenLoadBalanceVirtualServerL2ExceptionList(o["l2-exception-list"], d, "l2_exception_list")); err != nil {
		if !fortiAPIPatch(o["l2-exception-list"]) {
			return fmt.Errorf("Error reading l2-exception-list: %v", err)
		}
	}

	if err = d.Set("ips_profile", dataSourceFlattenLoadBalanceVirtualServerIpsProfile(o["ips_profile"], d, "ips_profile")); err != nil {
		if !fortiAPIPatch(o["ips_profile"]) {
			return fmt.Errorf("Error reading ips_profile: %v", err)
		}
	}

	if err = d.Set("traffic_group", dataSourceFlattenLoadBalanceVirtualServerTrafficGroup(o["traffic-group"], d, "traffic_group")); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("stream_scripting_list", dataSourceFlattenLoadBalanceVirtualServerStreamScriptingList(o["stream_scripting_list"], d, "stream_scripting_list")); err != nil {
		if !fortiAPIPatch(o["stream_scripting_list"]) {
			return fmt.Errorf("Error reading stream_scripting_list: %v", err)
		}
	}

	if err = d.Set("address6", dataSourceFlattenLoadBalanceVirtualServerAddress6(o["address6"], d, "address6")); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("clone_pool", dataSourceFlattenLoadBalanceVirtualServerClonePool(o["clone-pool"], d, "clone_pool")); err != nil {
		if !fortiAPIPatch(o["clone-pool"]) {
			return fmt.Errorf("Error reading clone-pool: %v", err)
		}
	}

	if err = d.Set("use_azure_lb_backend_ip", dataSourceFlattenLoadBalanceVirtualServerUseAzureLbBackendIp(o["use-azure-lb-backend-ip"], d, "use_azure_lb_backend_ip")); err != nil {
		if !fortiAPIPatch(o["use-azure-lb-backend-ip"]) {
			return fmt.Errorf("Error reading use-azure-lb-backend-ip: %v", err)
		}
	}

	if err = d.Set("public_ip", dataSourceFlattenLoadBalanceVirtualServerPublicIp(o["public-ip"], d, "public_ip")); err != nil {
		if !fortiAPIPatch(o["public-ip"]) {
			return fmt.Errorf("Error reading public-ip: %v", err)
		}
	}

	if err = d.Set("clone_traffic_type", dataSourceFlattenLoadBalanceVirtualServerCloneTrafficType(o["clone-traffic-type"], d, "clone_traffic_type")); err != nil {
		if !fortiAPIPatch(o["clone-traffic-type"]) {
			return fmt.Errorf("Error reading clone-traffic-type: %v", err)
		}
	}

	if err = d.Set("address", dataSourceFlattenLoadBalanceVirtualServerAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("alone", dataSourceFlattenLoadBalanceVirtualServerAlone(o["alone"], d, "alone")); err != nil {
		if !fortiAPIPatch(o["alone"]) {
			return fmt.Errorf("Error reading alone: %v", err)
		}
	}

	if err = d.Set("stream_scripting_flag", dataSourceFlattenLoadBalanceVirtualServerStreamScriptingFlag(o["stream_scripting_flag"], d, "stream_scripting_flag")); err != nil {
		if !fortiAPIPatch(o["stream_scripting_flag"]) {
			return fmt.Errorf("Error reading stream_scripting_flag: %v", err)
		}
	}

	if err = d.Set("client_ssl_profile", dataSourceFlattenLoadBalanceVirtualServerClientSslProfile(o["client_ssl_profile"], d, "client_ssl_profile")); err != nil {
		if !fortiAPIPatch(o["client_ssl_profile"]) {
			return fmt.Errorf("Error reading client_ssl_profile: %v", err)
		}
	}

	if err = d.Set("schedule_pool_list", dataSourceFlattenLoadBalanceVirtualServerSchedulePoolList(o["schedule-pool-list"], d, "schedule_pool_list")); err != nil {
		if !fortiAPIPatch(o["schedule-pool-list"]) {
			return fmt.Errorf("Error reading schedule-pool-list: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", dataSourceFlattenLoadBalanceVirtualServerSslMirror(o["ssl-mirror"], d, "ssl_mirror")); err != nil {
		if !fortiAPIPatch(o["ssl-mirror"]) {
			return fmt.Errorf("Error reading ssl-mirror: %v", err)
		}
	}

	if err = d.Set("schedule_list", dataSourceFlattenLoadBalanceVirtualServerScheduleList(o["schedule-list"], d, "schedule_list")); err != nil {
		if !fortiAPIPatch(o["schedule-list"]) {
			return fmt.Errorf("Error reading schedule-list: %v", err)
		}
	}

	if err = d.Set("warmup", dataSourceFlattenLoadBalanceVirtualServerWarmup(o["warmup"], d, "warmup")); err != nil {
		if !fortiAPIPatch(o["warmup"]) {
			return fmt.Errorf("Error reading warmup: %v", err)
		}
	}

	if err = d.Set("addr_type", dataSourceFlattenLoadBalanceVirtualServerAddrType(o["addr-type"], d, "addr_type")); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr-type: %v", err)
		}
	}

	if err = d.Set("connection_pool", dataSourceFlattenLoadBalanceVirtualServerConnectionPool(o["connection-pool"], d, "connection_pool")); err != nil {
		if !fortiAPIPatch(o["connection-pool"]) {
			return fmt.Errorf("Error reading connection-pool: %v", err)
		}
	}

	if err = d.Set("azure_lb_backend", dataSourceFlattenLoadBalanceVirtualServerAzureLbBackend(o["azure-lb-backend"], d, "azure_lb_backend")); err != nil {
		if !fortiAPIPatch(o["azure-lb-backend"]) {
			return fmt.Errorf("Error reading azure-lb-backend: %v", err)
		}
	}

	if err = d.Set("connection_rate_limit", dataSourceFlattenLoadBalanceVirtualServerConnectionRateLimit(o["connection-rate-limit"], d, "connection_rate_limit")); err != nil {
		if !fortiAPIPatch(o["connection-rate-limit"]) {
			return fmt.Errorf("Error reading connection-rate-limit: %v", err)
		}
	}

	if err = d.Set("trans_rate_limit", dataSourceFlattenLoadBalanceVirtualServerTransRateLimit(o["trans-rate-limit"], d, "trans_rate_limit")); err != nil {
		if !fortiAPIPatch(o["trans-rate-limit"]) {
			return fmt.Errorf("Error reading trans-rate-limit: %v", err)
		}
	}

	if err = d.Set("traffic_log", dataSourceFlattenLoadBalanceVirtualServerTrafficLog(o["traffic-log"], d, "traffic_log")); err != nil {
		if !fortiAPIPatch(o["traffic-log"]) {
			return fmt.Errorf("Error reading traffic-log: %v", err)
		}
	}

	if err = d.Set("av_profile", dataSourceFlattenLoadBalanceVirtualServerAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av-profile: %v", err)
		}
	}

	if err = d.Set("public_ip6", dataSourceFlattenLoadBalanceVirtualServerPublicIp6(o["public-ip6"], d, "public_ip6")); err != nil {
		if !fortiAPIPatch(o["public-ip6"]) {
			return fmt.Errorf("Error reading public-ip6: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenLoadBalanceVirtualServerInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("captcha_profile", dataSourceFlattenLoadBalanceVirtualServerCaptchaProfile(o["captcha-profile"], d, "captcha_profile")); err != nil {
		if !fortiAPIPatch(o["captcha-profile"]) {
			return fmt.Errorf("Error reading captcha-profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", dataSourceFlattenLoadBalanceVirtualServerWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf-profile: %v", err)
		}
	}

	if err = d.Set("content_routing", dataSourceFlattenLoadBalanceVirtualServerContentRouting(o["content-routing"], d, "content_routing")); err != nil {
		if !fortiAPIPatch(o["content-routing"]) {
			return fmt.Errorf("Error reading content-routing: %v", err)
		}
	}

	return nil
}
