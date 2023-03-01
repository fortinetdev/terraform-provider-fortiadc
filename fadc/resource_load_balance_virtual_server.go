// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance virtual server.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceVirtualServer() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceVirtualServerRead,
		Update: resourceLoadBalanceVirtualServerUpdate,
		Create: resourceLoadBalanceVirtualServerCreate,
		Delete: resourceLoadBalanceVirtualServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scripting_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"packet_fwd_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"warmrate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2https": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"error_page": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fortiview": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ztna_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"adfs_published_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_rewriting": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2https_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"error_msg": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scripting_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_routing_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"public_ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"one_click_gslb_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pagespeed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connection_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_mirror_intf": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dos_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_rewriting_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"l2_exception_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ips_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"stream_scripting_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"clone_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_azure_lb_backend_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"clone_traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"alone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"stream_scripting_flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_ssl_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"schedule_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"schedule_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"warmup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connection_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"azure_lb_backend": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connection_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trans_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traffic_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"public_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"captcha_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_routing": &schema.Schema{
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
func resourceLoadBalanceVirtualServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceVirtualServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceVirtualServer resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceVirtualServer(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceVirtualServer resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceVirtualServerRead(d, m)
}
func resourceLoadBalanceVirtualServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceVirtualServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceVirtualServer resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceVirtualServer(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceVirtualServer resource: %v", err)
	}

	return resourceLoadBalanceVirtualServerRead(d, m)
}
func resourceLoadBalanceVirtualServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceVirtualServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceVirtualServer resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceVirtualServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceVirtualServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceVirtualServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceVirtualServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceVirtualServer resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceVirtualServerScriptingList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPacketFwdMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerWarmrate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerHttp2Https(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerErrorPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerFortiview(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerZtnaProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAuthPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAdfsPublishedService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerContentRewriting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerHttp2HttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerErrorMsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerScriptingFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerContentRoutingList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerWccp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPublicIpType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerOneClickGslbServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPagespeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerConnectionLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerHostName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerSslMirrorIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerSourcePoolList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerDosProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerContentRewritingList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerL2ExceptionList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerIpsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerTrafficGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerStreamScriptingList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerClonePool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerUseAzureLbBackendIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPublicIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerCloneTrafficType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAlone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerStreamScriptingFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerClientSslProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerSchedulePoolList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerSslMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerScheduleList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerWarmup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerConnectionPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAzureLbBackend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerConnectionRateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerTransRateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerTrafficLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerAvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerPublicIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerCaptchaProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerWafProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceVirtualServerContentRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceVirtualServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("scripting_list", flattenLoadBalanceVirtualServerScriptingList(o["scripting_list"], d, "scripting_list", sv)); err != nil {
		if !fortiAPIPatch(o["scripting_list"]) {
			return fmt.Errorf("Error reading scripting_list: %v", err)
		}
	}

	if err = d.Set("status", flattenLoadBalanceVirtualServerStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("packet_fwd_method", flattenLoadBalanceVirtualServerPacketFwdMethod(o["packet-fwd-method"], d, "packet_fwd_method", sv)); err != nil {
		if !fortiAPIPatch(o["packet-fwd-method"]) {
			return fmt.Errorf("Error reading packet-fwd-method: %v", err)
		}
	}

	if err = d.Set("warmrate", flattenLoadBalanceVirtualServerWarmrate(o["warmrate"], d, "warmrate", sv)); err != nil {
		if !fortiAPIPatch(o["warmrate"]) {
			return fmt.Errorf("Error reading warmrate: %v", err)
		}
	}

	if err = d.Set("http2https", flattenLoadBalanceVirtualServerHttp2Https(o["http2https"], d, "http2https", sv)); err != nil {
		if !fortiAPIPatch(o["http2https"]) {
			return fmt.Errorf("Error reading http2https: %v", err)
		}
	}

	if err = d.Set("pool", flattenLoadBalanceVirtualServerPool(o["pool"], d, "pool", sv)); err != nil {
		if !fortiAPIPatch(o["pool"]) {
			return fmt.Errorf("Error reading pool: %v", err)
		}
	}

	if err = d.Set("error_page", flattenLoadBalanceVirtualServerErrorPage(o["error-page"], d, "error_page", sv)); err != nil {
		if !fortiAPIPatch(o["error-page"]) {
			return fmt.Errorf("Error reading error-page: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenLoadBalanceVirtualServerFortiview(o["fortiview"], d, "fortiview", sv)); err != nil {
		if !fortiAPIPatch(o["fortiview"]) {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("ztna_profile", flattenLoadBalanceVirtualServerZtnaProfile(o["ztna_profile"], d, "ztna_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ztna_profile"]) {
			return fmt.Errorf("Error reading ztna_profile: %v", err)
		}
	}

	if err = d.Set("auth_policy", flattenLoadBalanceVirtualServerAuthPolicy(o["auth_policy"], d, "auth_policy", sv)); err != nil {
		if !fortiAPIPatch(o["auth_policy"]) {
			return fmt.Errorf("Error reading auth_policy: %v", err)
		}
	}

	if err = d.Set("adfs_published_service", flattenLoadBalanceVirtualServerAdfsPublishedService(o["adfs-published-service"], d, "adfs_published_service", sv)); err != nil {
		if !fortiAPIPatch(o["adfs-published-service"]) {
			return fmt.Errorf("Error reading adfs-published-service: %v", err)
		}
	}

	if err = d.Set("content_rewriting", flattenLoadBalanceVirtualServerContentRewriting(o["content-rewriting"], d, "content_rewriting", sv)); err != nil {
		if !fortiAPIPatch(o["content-rewriting"]) {
			return fmt.Errorf("Error reading content-rewriting: %v", err)
		}
	}

	if err = d.Set("http2https_port", flattenLoadBalanceVirtualServerHttp2HttpsPort(o["http2https-port"], d, "http2https_port", sv)); err != nil {
		if !fortiAPIPatch(o["http2https-port"]) {
			return fmt.Errorf("Error reading http2https-port: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenLoadBalanceVirtualServerDomainName(o["domain-name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain-name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenLoadBalanceVirtualServerProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("error_msg", flattenLoadBalanceVirtualServerErrorMsg(o["error-msg"], d, "error_msg", sv)); err != nil {
		if !fortiAPIPatch(o["error-msg"]) {
			return fmt.Errorf("Error reading error-msg: %v", err)
		}
	}

	if err = d.Set("port", flattenLoadBalanceVirtualServerPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("scripting_flag", flattenLoadBalanceVirtualServerScriptingFlag(o["scripting_flag"], d, "scripting_flag", sv)); err != nil {
		if !fortiAPIPatch(o["scripting_flag"]) {
			return fmt.Errorf("Error reading scripting_flag: %v", err)
		}
	}

	if err = d.Set("comments", flattenLoadBalanceVirtualServerComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("content_routing_list", flattenLoadBalanceVirtualServerContentRoutingList(o["content-routing-list"], d, "content_routing_list", sv)); err != nil {
		if !fortiAPIPatch(o["content-routing-list"]) {
			return fmt.Errorf("Error reading content-routing-list: %v", err)
		}
	}

	if err = d.Set("wccp", flattenLoadBalanceVirtualServerWccp(o["wccp"], d, "wccp", sv)); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("public_ip_type", flattenLoadBalanceVirtualServerPublicIpType(o["public-ip-type"], d, "public_ip_type", sv)); err != nil {
		if !fortiAPIPatch(o["public-ip-type"]) {
			return fmt.Errorf("Error reading public-ip-type: %v", err)
		}
	}

	if err = d.Set("one_click_gslb_server", flattenLoadBalanceVirtualServerOneClickGslbServer(o["one-click-gslb-server"], d, "one_click_gslb_server", sv)); err != nil {
		if !fortiAPIPatch(o["one-click-gslb-server"]) {
			return fmt.Errorf("Error reading one-click-gslb-server: %v", err)
		}
	}

	if err = d.Set("pagespeed", flattenLoadBalanceVirtualServerPagespeed(o["pagespeed"], d, "pagespeed", sv)); err != nil {
		if !fortiAPIPatch(o["pagespeed"]) {
			return fmt.Errorf("Error reading pagespeed: %v", err)
		}
	}

	if err = d.Set("connection_limit", flattenLoadBalanceVirtualServerConnectionLimit(o["connection-limit"], d, "connection_limit", sv)); err != nil {
		if !fortiAPIPatch(o["connection-limit"]) {
			return fmt.Errorf("Error reading connection-limit: %v", err)
		}
	}

	if err = d.Set("host_name", flattenLoadBalanceVirtualServerHostName(o["host-name"], d, "host_name", sv)); err != nil {
		if !fortiAPIPatch(o["host-name"]) {
			return fmt.Errorf("Error reading host-name: %v", err)
		}
	}

	if err = d.Set("ssl_mirror_intf", flattenLoadBalanceVirtualServerSslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mirror-intf"]) {
			return fmt.Errorf("Error reading ssl-mirror-intf: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceVirtualServerMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("source_pool_list", flattenLoadBalanceVirtualServerSourcePoolList(o["source-pool-list"], d, "source_pool_list", sv)); err != nil {
		if !fortiAPIPatch(o["source-pool-list"]) {
			return fmt.Errorf("Error reading source-pool-list: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceVirtualServerType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("method", flattenLoadBalanceVirtualServerMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("persistence", flattenLoadBalanceVirtualServerPersistence(o["persistence"], d, "persistence", sv)); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("dos_profile", flattenLoadBalanceVirtualServerDosProfile(o["dos_profile"], d, "dos_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dos_profile"]) {
			return fmt.Errorf("Error reading dos_profile: %v", err)
		}
	}

	if err = d.Set("profile", flattenLoadBalanceVirtualServerProfile(o["profile"], d, "profile", sv)); err != nil {
		if !fortiAPIPatch(o["profile"]) {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("content_rewriting_list", flattenLoadBalanceVirtualServerContentRewritingList(o["content-rewriting-list"], d, "content_rewriting_list", sv)); err != nil {
		if !fortiAPIPatch(o["content-rewriting-list"]) {
			return fmt.Errorf("Error reading content-rewriting-list: %v", err)
		}
	}

	if err = d.Set("l2_exception_list", flattenLoadBalanceVirtualServerL2ExceptionList(o["l2-exception-list"], d, "l2_exception_list", sv)); err != nil {
		if !fortiAPIPatch(o["l2-exception-list"]) {
			return fmt.Errorf("Error reading l2-exception-list: %v", err)
		}
	}

	if err = d.Set("ips_profile", flattenLoadBalanceVirtualServerIpsProfile(o["ips_profile"], d, "ips_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ips_profile"]) {
			return fmt.Errorf("Error reading ips_profile: %v", err)
		}
	}

	if err = d.Set("traffic_group", flattenLoadBalanceVirtualServerTrafficGroup(o["traffic-group"], d, "traffic_group", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("stream_scripting_list", flattenLoadBalanceVirtualServerStreamScriptingList(o["stream_scripting_list"], d, "stream_scripting_list", sv)); err != nil {
		if !fortiAPIPatch(o["stream_scripting_list"]) {
			return fmt.Errorf("Error reading stream_scripting_list: %v", err)
		}
	}

	if err = d.Set("address6", flattenLoadBalanceVirtualServerAddress6(o["address6"], d, "address6", sv)); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("clone_pool", flattenLoadBalanceVirtualServerClonePool(o["clone-pool"], d, "clone_pool", sv)); err != nil {
		if !fortiAPIPatch(o["clone-pool"]) {
			return fmt.Errorf("Error reading clone-pool: %v", err)
		}
	}

	if err = d.Set("use_azure_lb_backend_ip", flattenLoadBalanceVirtualServerUseAzureLbBackendIp(o["use-azure-lb-backend-ip"], d, "use_azure_lb_backend_ip", sv)); err != nil {
		if !fortiAPIPatch(o["use-azure-lb-backend-ip"]) {
			return fmt.Errorf("Error reading use-azure-lb-backend-ip: %v", err)
		}
	}

	if err = d.Set("public_ip", flattenLoadBalanceVirtualServerPublicIp(o["public-ip"], d, "public_ip", sv)); err != nil {
		if !fortiAPIPatch(o["public-ip"]) {
			return fmt.Errorf("Error reading public-ip: %v", err)
		}
	}

	if err = d.Set("clone_traffic_type", flattenLoadBalanceVirtualServerCloneTrafficType(o["clone-traffic-type"], d, "clone_traffic_type", sv)); err != nil {
		if !fortiAPIPatch(o["clone-traffic-type"]) {
			return fmt.Errorf("Error reading clone-traffic-type: %v", err)
		}
	}

	if err = d.Set("address", flattenLoadBalanceVirtualServerAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("alone", flattenLoadBalanceVirtualServerAlone(o["alone"], d, "alone", sv)); err != nil {
		if !fortiAPIPatch(o["alone"]) {
			return fmt.Errorf("Error reading alone: %v", err)
		}
	}

	if err = d.Set("stream_scripting_flag", flattenLoadBalanceVirtualServerStreamScriptingFlag(o["stream_scripting_flag"], d, "stream_scripting_flag", sv)); err != nil {
		if !fortiAPIPatch(o["stream_scripting_flag"]) {
			return fmt.Errorf("Error reading stream_scripting_flag: %v", err)
		}
	}

	if err = d.Set("client_ssl_profile", flattenLoadBalanceVirtualServerClientSslProfile(o["client_ssl_profile"], d, "client_ssl_profile", sv)); err != nil {
		if !fortiAPIPatch(o["client_ssl_profile"]) {
			return fmt.Errorf("Error reading client_ssl_profile: %v", err)
		}
	}

	if err = d.Set("schedule_pool_list", flattenLoadBalanceVirtualServerSchedulePoolList(o["schedule-pool-list"], d, "schedule_pool_list", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-pool-list"]) {
			return fmt.Errorf("Error reading schedule-pool-list: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", flattenLoadBalanceVirtualServerSslMirror(o["ssl-mirror"], d, "ssl_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mirror"]) {
			return fmt.Errorf("Error reading ssl-mirror: %v", err)
		}
	}

	if err = d.Set("schedule_list", flattenLoadBalanceVirtualServerScheduleList(o["schedule-list"], d, "schedule_list", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-list"]) {
			return fmt.Errorf("Error reading schedule-list: %v", err)
		}
	}

	if err = d.Set("warmup", flattenLoadBalanceVirtualServerWarmup(o["warmup"], d, "warmup", sv)); err != nil {
		if !fortiAPIPatch(o["warmup"]) {
			return fmt.Errorf("Error reading warmup: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenLoadBalanceVirtualServerAddrType(o["addr-type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr-type: %v", err)
		}
	}

	if err = d.Set("connection_pool", flattenLoadBalanceVirtualServerConnectionPool(o["connection-pool"], d, "connection_pool", sv)); err != nil {
		if !fortiAPIPatch(o["connection-pool"]) {
			return fmt.Errorf("Error reading connection-pool: %v", err)
		}
	}

	if err = d.Set("azure_lb_backend", flattenLoadBalanceVirtualServerAzureLbBackend(o["azure-lb-backend"], d, "azure_lb_backend", sv)); err != nil {
		if !fortiAPIPatch(o["azure-lb-backend"]) {
			return fmt.Errorf("Error reading azure-lb-backend: %v", err)
		}
	}

	if err = d.Set("connection_rate_limit", flattenLoadBalanceVirtualServerConnectionRateLimit(o["connection-rate-limit"], d, "connection_rate_limit", sv)); err != nil {
		if !fortiAPIPatch(o["connection-rate-limit"]) {
			return fmt.Errorf("Error reading connection-rate-limit: %v", err)
		}
	}

	if err = d.Set("trans_rate_limit", flattenLoadBalanceVirtualServerTransRateLimit(o["trans-rate-limit"], d, "trans_rate_limit", sv)); err != nil {
		if !fortiAPIPatch(o["trans-rate-limit"]) {
			return fmt.Errorf("Error reading trans-rate-limit: %v", err)
		}
	}

	if err = d.Set("traffic_log", flattenLoadBalanceVirtualServerTrafficLog(o["traffic-log"], d, "traffic_log", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-log"]) {
			return fmt.Errorf("Error reading traffic-log: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenLoadBalanceVirtualServerAvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av-profile: %v", err)
		}
	}

	if err = d.Set("public_ip6", flattenLoadBalanceVirtualServerPublicIp6(o["public-ip6"], d, "public_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["public-ip6"]) {
			return fmt.Errorf("Error reading public-ip6: %v", err)
		}
	}

	if err = d.Set("interface", flattenLoadBalanceVirtualServerInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("captcha_profile", flattenLoadBalanceVirtualServerCaptchaProfile(o["captcha-profile"], d, "captcha_profile", sv)); err != nil {
		if !fortiAPIPatch(o["captcha-profile"]) {
			return fmt.Errorf("Error reading captcha-profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenLoadBalanceVirtualServerWafProfile(o["waf-profile"], d, "waf_profile", sv)); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf-profile: %v", err)
		}
	}

	if err = d.Set("content_routing", flattenLoadBalanceVirtualServerContentRouting(o["content-routing"], d, "content_routing", sv)); err != nil {
		if !fortiAPIPatch(o["content-routing"]) {
			return fmt.Errorf("Error reading content-routing: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceVirtualServerScriptingList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPacketFwdMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerWarmrate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerHttp2Https(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerErrorPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerFortiview(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerZtnaProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAuthPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAdfsPublishedService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerContentRewriting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerHttp2HttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerErrorMsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerScriptingFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerContentRoutingList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerWccp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPublicIpType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerOneClickGslbServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPagespeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerConnectionLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerHostName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerSslMirrorIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerSourcePoolList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerDosProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerContentRewritingList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerL2ExceptionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerIpsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerTrafficGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerStreamScriptingList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerClonePool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerUseAzureLbBackendIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPublicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerCloneTrafficType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAlone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerStreamScriptingFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerClientSslProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerSchedulePoolList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerSslMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerScheduleList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerWarmup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerConnectionPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAzureLbBackend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerConnectionRateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerTransRateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerTrafficLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerAvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerPublicIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerCaptchaProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerWafProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceVirtualServerContentRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceVirtualServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("scripting_list"); ok {
		t, err := expandLoadBalanceVirtualServerScriptingList(d, v, "scripting_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scripting_list"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLoadBalanceVirtualServerStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("packet_fwd_method"); ok {
		t, err := expandLoadBalanceVirtualServerPacketFwdMethod(d, v, "packet_fwd_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-fwd-method"] = t
		}
	}

	if v, ok := d.GetOk("warmrate"); ok {
		t, err := expandLoadBalanceVirtualServerWarmrate(d, v, "warmrate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warmrate"] = t
		}
	}

	if v, ok := d.GetOk("http2https"); ok {
		t, err := expandLoadBalanceVirtualServerHttp2Https(d, v, "http2https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2https"] = t
		}
	}

	if v, ok := d.GetOk("pool"); ok {
		t, err := expandLoadBalanceVirtualServerPool(d, v, "pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool"] = t
		}
	}

	if v, ok := d.GetOk("error_page"); ok {
		t, err := expandLoadBalanceVirtualServerErrorPage(d, v, "error_page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["error-page"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok {
		t, err := expandLoadBalanceVirtualServerFortiview(d, v, "fortiview", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("ztna_profile"); ok {
		t, err := expandLoadBalanceVirtualServerZtnaProfile(d, v, "ztna_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna_profile"] = t
		}
	}

	if v, ok := d.GetOk("auth_policy"); ok {
		t, err := expandLoadBalanceVirtualServerAuthPolicy(d, v, "auth_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth_policy"] = t
		}
	}

	if v, ok := d.GetOk("adfs_published_service"); ok {
		t, err := expandLoadBalanceVirtualServerAdfsPublishedService(d, v, "adfs_published_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-published-service"] = t
		}
	}

	if v, ok := d.GetOk("content_rewriting"); ok {
		t, err := expandLoadBalanceVirtualServerContentRewriting(d, v, "content_rewriting", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-rewriting"] = t
		}
	}

	if v, ok := d.GetOk("http2https_port"); ok {
		t, err := expandLoadBalanceVirtualServerHttp2HttpsPort(d, v, "http2https_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2https-port"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandLoadBalanceVirtualServerDomainName(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandLoadBalanceVirtualServerProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("error_msg"); ok {
		t, err := expandLoadBalanceVirtualServerErrorMsg(d, v, "error_msg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["error-msg"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandLoadBalanceVirtualServerPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("scripting_flag"); ok {
		t, err := expandLoadBalanceVirtualServerScriptingFlag(d, v, "scripting_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scripting_flag"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandLoadBalanceVirtualServerComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("content_routing_list"); ok {
		t, err := expandLoadBalanceVirtualServerContentRoutingList(d, v, "content_routing_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-routing-list"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok {
		t, err := expandLoadBalanceVirtualServerWccp(d, v, "wccp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("public_ip_type"); ok {
		t, err := expandLoadBalanceVirtualServerPublicIpType(d, v, "public_ip_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-ip-type"] = t
		}
	}

	if v, ok := d.GetOk("one_click_gslb_server"); ok {
		t, err := expandLoadBalanceVirtualServerOneClickGslbServer(d, v, "one_click_gslb_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["one-click-gslb-server"] = t
		}
	}

	if v, ok := d.GetOk("pagespeed"); ok {
		t, err := expandLoadBalanceVirtualServerPagespeed(d, v, "pagespeed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pagespeed"] = t
		}
	}

	if v, ok := d.GetOk("connection_limit"); ok {
		t, err := expandLoadBalanceVirtualServerConnectionLimit(d, v, "connection_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-limit"] = t
		}
	}

	if v, ok := d.GetOk("host_name"); ok {
		t, err := expandLoadBalanceVirtualServerHostName(d, v, "host_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-name"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror_intf"); ok {
		t, err := expandLoadBalanceVirtualServerSslMirrorIntf(d, v, "ssl_mirror_intf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror-intf"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceVirtualServerMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("source_pool_list"); ok {
		t, err := expandLoadBalanceVirtualServerSourcePoolList(d, v, "source_pool_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-pool-list"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceVirtualServerType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandLoadBalanceVirtualServerMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok {
		t, err := expandLoadBalanceVirtualServerPersistence(d, v, "persistence", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("dos_profile"); ok {
		t, err := expandLoadBalanceVirtualServerDosProfile(d, v, "dos_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dos_profile"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok {
		t, err := expandLoadBalanceVirtualServerProfile(d, v, "profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("content_rewriting_list"); ok {
		t, err := expandLoadBalanceVirtualServerContentRewritingList(d, v, "content_rewriting_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-rewriting-list"] = t
		}
	}

	if v, ok := d.GetOk("l2_exception_list"); ok {
		t, err := expandLoadBalanceVirtualServerL2ExceptionList(d, v, "l2_exception_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2-exception-list"] = t
		}
	}

	if v, ok := d.GetOk("ips_profile"); ok {
		t, err := expandLoadBalanceVirtualServerIpsProfile(d, v, "ips_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips_profile"] = t
		}
	}

	if v, ok := d.GetOk("traffic_group"); ok {
		t, err := expandLoadBalanceVirtualServerTrafficGroup(d, v, "traffic_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-group"] = t
		}
	}

	if v, ok := d.GetOk("stream_scripting_list"); ok {
		t, err := expandLoadBalanceVirtualServerStreamScriptingList(d, v, "stream_scripting_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream_scripting_list"] = t
		}
	}

	if v, ok := d.GetOk("address6"); ok {
		t, err := expandLoadBalanceVirtualServerAddress6(d, v, "address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("clone_pool"); ok {
		t, err := expandLoadBalanceVirtualServerClonePool(d, v, "clone_pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clone-pool"] = t
		}
	}

	if v, ok := d.GetOk("use_azure_lb_backend_ip"); ok {
		t, err := expandLoadBalanceVirtualServerUseAzureLbBackendIp(d, v, "use_azure_lb_backend_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-azure-lb-backend-ip"] = t
		}
	}

	if v, ok := d.GetOk("public_ip"); ok {
		t, err := expandLoadBalanceVirtualServerPublicIp(d, v, "public_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-ip"] = t
		}
	}

	if v, ok := d.GetOk("clone_traffic_type"); ok {
		t, err := expandLoadBalanceVirtualServerCloneTrafficType(d, v, "clone_traffic_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clone-traffic-type"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandLoadBalanceVirtualServerAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("alone"); ok {
		t, err := expandLoadBalanceVirtualServerAlone(d, v, "alone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alone"] = t
		}
	}

	if v, ok := d.GetOk("stream_scripting_flag"); ok {
		t, err := expandLoadBalanceVirtualServerStreamScriptingFlag(d, v, "stream_scripting_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream_scripting_flag"] = t
		}
	}

	if v, ok := d.GetOk("client_ssl_profile"); ok {
		t, err := expandLoadBalanceVirtualServerClientSslProfile(d, v, "client_ssl_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_ssl_profile"] = t
		}
	}

	if v, ok := d.GetOk("schedule_pool_list"); ok {
		t, err := expandLoadBalanceVirtualServerSchedulePoolList(d, v, "schedule_pool_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-pool-list"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror"); ok {
		t, err := expandLoadBalanceVirtualServerSslMirror(d, v, "ssl_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror"] = t
		}
	}

	if v, ok := d.GetOk("schedule_list"); ok {
		t, err := expandLoadBalanceVirtualServerScheduleList(d, v, "schedule_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-list"] = t
		}
	}

	if v, ok := d.GetOk("warmup"); ok {
		t, err := expandLoadBalanceVirtualServerWarmup(d, v, "warmup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warmup"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandLoadBalanceVirtualServerAddrType(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("connection_pool"); ok {
		t, err := expandLoadBalanceVirtualServerConnectionPool(d, v, "connection_pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-pool"] = t
		}
	}

	if v, ok := d.GetOk("azure_lb_backend"); ok {
		t, err := expandLoadBalanceVirtualServerAzureLbBackend(d, v, "azure_lb_backend", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-lb-backend"] = t
		}
	}

	if v, ok := d.GetOk("connection_rate_limit"); ok {
		t, err := expandLoadBalanceVirtualServerConnectionRateLimit(d, v, "connection_rate_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("trans_rate_limit"); ok {
		t, err := expandLoadBalanceVirtualServerTransRateLimit(d, v, "trans_rate_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trans-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("traffic_log"); ok {
		t, err := expandLoadBalanceVirtualServerTrafficLog(d, v, "traffic_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-log"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandLoadBalanceVirtualServerAvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("public_ip6"); ok {
		t, err := expandLoadBalanceVirtualServerPublicIp6(d, v, "public_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-ip6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandLoadBalanceVirtualServerInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("captcha_profile"); ok {
		t, err := expandLoadBalanceVirtualServerCaptchaProfile(d, v, "captcha_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captcha-profile"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok {
		t, err := expandLoadBalanceVirtualServerWafProfile(d, v, "waf_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("content_routing"); ok {
		t, err := expandLoadBalanceVirtualServerContentRouting(d, v, "content_routing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-routing"] = t
		}
	}

	return &obj, nil
}
