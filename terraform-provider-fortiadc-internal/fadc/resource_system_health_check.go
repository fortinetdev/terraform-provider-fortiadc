// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system health check.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHealthCheck() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemHealthCheckRead,
		Update: resourceSystemHealthCheckUpdate,
		Create: resourceSystemHealthCheckCreate,
		Delete: resourceSystemHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"counter_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"basedn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pwd_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dest_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mssql_row": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"row": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mssql_column": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mssql_receive_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oracle_receive_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sip_request_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"acct_appid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connect_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"binddn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_extra_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"column": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_connect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth_appid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"agent_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connect_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"method_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dest_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rtsp_method_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mssql_send_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"string_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"down_retry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"origin_realm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"origin_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"compare_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mysql_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dest_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cpu_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_ip6_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mem_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"disk_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"receive_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"value_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"radius_reject": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rtsp_describe_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oracle_send_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"send_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"up_retry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remote_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vendor_id": &schema.Schema{
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
func resourceSystemHealthCheckCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemHealthCheck: type error")
	}

	obj, err := getObjectSystemHealthCheck(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheck resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHealthCheck(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheck resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemHealthCheckRead(d, m)
}
func resourceSystemHealthCheckUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHealthCheck(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheck resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHealthCheck(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheck resource: %v", err)
	}

	return resourceSystemHealthCheckRead(d, m)
}
func resourceSystemHealthCheckDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemHealthCheck(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHealthCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemHealthCheckRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemHealthCheck(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheck resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHealthCheck(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheck resource from API: %v", err)
	}
	return nil
}
func flattenSystemHealthCheckCounterValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckBasedn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckPwdType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDestAddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMssqlRow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDisk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHttpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMssqlColumn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckStatusCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckFolder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMssqlReceiveString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckOracleReceiveString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHostAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckSipRequestType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRemotePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckAcctAppid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckAllowSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRemotePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckConnectString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckProductName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckNasIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckBinddn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHostAddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHttpExtraString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckColumn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHttpConnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckAuthAppid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckAgentType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRemoteHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckConnectType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMethodType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckSslCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDestAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRtspMethodType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMssqlSendString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckStringValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDownRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckOriginRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckOriginHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckSecretKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckCompareType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMatchType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckLocalCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMysqlServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDestAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckCpuWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHostIpAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHostIp6Addr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMemWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckSid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDiskWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckReceiveString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckOid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckValueType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRadiusReject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRtspDescribeUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckOracleSendString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckSendString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckUpRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckRemoteUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckCpu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckVendorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHealthCheck(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("counter_value", flattenSystemHealthCheckCounterValue(o["counter-value"], d, "counter_value", sv)); err != nil {
		if !fortiAPIPatch(o["counter-value"]) {
			return fmt.Errorf("Error reading counter-value: %v", err)
		}
	}

	if err = d.Set("service_name", flattenSystemHealthCheckServiceName(o["service_name"], d, "service_name", sv)); err != nil {
		if !fortiAPIPatch(o["service_name"]) {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("basedn", flattenSystemHealthCheckBasedn(o["baseDN"], d, "basedn", sv)); err != nil {
		if !fortiAPIPatch(o["baseDN"]) {
			return fmt.Errorf("Error reading baseDN: %v", err)
		}
	}

	if err = d.Set("pwd_type", flattenSystemHealthCheckPwdType(o["pwd_type"], d, "pwd_type", sv)); err != nil {
		if !fortiAPIPatch(o["pwd_type"]) {
			return fmt.Errorf("Error reading pwd_type: %v", err)
		}
	}

	if err = d.Set("dest_addr6", flattenSystemHealthCheckDestAddr6(o["dest_addr6"], d, "dest_addr6", sv)); err != nil {
		if !fortiAPIPatch(o["dest_addr6"]) {
			return fmt.Errorf("Error reading dest_addr6: %v", err)
		}
	}

	if err = d.Set("mssql_row", flattenSystemHealthCheckMssqlRow(o["mssql-row"], d, "mssql_row", sv)); err != nil {
		if !fortiAPIPatch(o["mssql-row"]) {
			return fmt.Errorf("Error reading mssql-row: %v", err)
		}
	}

	if err = d.Set("disk", flattenSystemHealthCheckDisk(o["disk"], d, "disk", sv)); err != nil {
		if !fortiAPIPatch(o["disk"]) {
			return fmt.Errorf("Error reading disk: %v", err)
		}
	}

	if err = d.Set("row", flattenSystemHealthCheckRow(o["row"], d, "row", sv)); err != nil {
		if !fortiAPIPatch(o["row"]) {
			return fmt.Errorf("Error reading row: %v", err)
		}
	}

	if err = d.Set("http_version", flattenSystemHealthCheckHttpVersion(o["http_version"], d, "http_version", sv)); err != nil {
		if !fortiAPIPatch(o["http_version"]) {
			return fmt.Errorf("Error reading http_version: %v", err)
		}
	}

	if err = d.Set("passive", flattenSystemHealthCheckPassive(o["passive"], d, "passive", sv)); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("mssql_column", flattenSystemHealthCheckMssqlColumn(o["mssql-column"], d, "mssql_column", sv)); err != nil {
		if !fortiAPIPatch(o["mssql-column"]) {
			return fmt.Errorf("Error reading mssql-column: %v", err)
		}
	}

	if err = d.Set("status_code", flattenSystemHealthCheckStatusCode(o["status_code"], d, "status_code", sv)); err != nil {
		if !fortiAPIPatch(o["status_code"]) {
			return fmt.Errorf("Error reading status_code: %v", err)
		}
	}

	if err = d.Set("folder", flattenSystemHealthCheckFolder(o["folder"], d, "folder", sv)); err != nil {
		if !fortiAPIPatch(o["folder"]) {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenSystemHealthCheckAddrType(o["addr_type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr_type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("mssql_receive_string", flattenSystemHealthCheckMssqlReceiveString(o["mssql-receive-string"], d, "mssql_receive_string", sv)); err != nil {
		if !fortiAPIPatch(o["mssql-receive-string"]) {
			return fmt.Errorf("Error reading mssql-receive-string: %v", err)
		}
	}

	if err = d.Set("oracle_receive_string", flattenSystemHealthCheckOracleReceiveString(o["oracle-receive-string"], d, "oracle_receive_string", sv)); err != nil {
		if !fortiAPIPatch(o["oracle-receive-string"]) {
			return fmt.Errorf("Error reading oracle-receive-string: %v", err)
		}
	}

	if err = d.Set("host_addr", flattenSystemHealthCheckHostAddr(o["host_addr"], d, "host_addr", sv)); err != nil {
		if !fortiAPIPatch(o["host_addr"]) {
			return fmt.Errorf("Error reading host_addr: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemHealthCheckPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("sip_request_type", flattenSystemHealthCheckSipRequestType(o["sip_request_type"], d, "sip_request_type", sv)); err != nil {
		if !fortiAPIPatch(o["sip_request_type"]) {
			return fmt.Errorf("Error reading sip_request_type: %v", err)
		}
	}

	if err = d.Set("remote_password", flattenSystemHealthCheckRemotePassword(o["remote-password"], d, "remote_password", sv)); err != nil {
		if !fortiAPIPatch(o["remote-password"]) {
			return fmt.Errorf("Error reading remote-password: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemHealthCheckTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("acct_appid", flattenSystemHealthCheckAcctAppid(o["acct-appid"], d, "acct_appid", sv)); err != nil {
		if !fortiAPIPatch(o["acct-appid"]) {
			return fmt.Errorf("Error reading acct-appid: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemHealthCheckPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("allow_ssl_version", flattenSystemHealthCheckAllowSslVersion(o["allow-ssl-version"], d, "allow_ssl_version", sv)); err != nil {
		if !fortiAPIPatch(o["allow-ssl-version"]) {
			return fmt.Errorf("Error reading allow-ssl-version: %v", err)
		}
	}

	if err = d.Set("remote_port", flattenSystemHealthCheckRemotePort(o["remote_port"], d, "remote_port", sv)); err != nil {
		if !fortiAPIPatch(o["remote_port"]) {
			return fmt.Errorf("Error reading remote_port: %v", err)
		}
	}

	if err = d.Set("connect_string", flattenSystemHealthCheckConnectString(o["connect-string"], d, "connect_string", sv)); err != nil {
		if !fortiAPIPatch(o["connect-string"]) {
			return fmt.Errorf("Error reading connect-string: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenSystemHealthCheckDomainName(o["domain_name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain_name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("version", flattenSystemHealthCheckVersion(o["version"], d, "version", sv)); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("product_name", flattenSystemHealthCheckProductName(o["product-name"], d, "product_name", sv)); err != nil {
		if !fortiAPIPatch(o["product-name"]) {
			return fmt.Errorf("Error reading product-name: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenSystemHealthCheckNasIp(o["nas-ip"], d, "nas_ip", sv)); err != nil {
		if !fortiAPIPatch(o["nas-ip"]) {
			return fmt.Errorf("Error reading nas-ip: %v", err)
		}
	}

	if err = d.Set("binddn", flattenSystemHealthCheckBinddn(o["bindDN"], d, "binddn", sv)); err != nil {
		if !fortiAPIPatch(o["bindDN"]) {
			return fmt.Errorf("Error reading bindDN: %v", err)
		}
	}

	if err = d.Set("mem", flattenSystemHealthCheckMem(o["mem"], d, "mem", sv)); err != nil {
		if !fortiAPIPatch(o["mem"]) {
			return fmt.Errorf("Error reading mem: %v", err)
		}
	}

	if err = d.Set("host_addr6", flattenSystemHealthCheckHostAddr6(o["host_addr6"], d, "host_addr6", sv)); err != nil {
		if !fortiAPIPatch(o["host_addr6"]) {
			return fmt.Errorf("Error reading host_addr6: %v", err)
		}
	}

	if err = d.Set("http_extra_string", flattenSystemHealthCheckHttpExtraString(o["http_extra_string"], d, "http_extra_string", sv)); err != nil {
		if !fortiAPIPatch(o["http_extra_string"]) {
			return fmt.Errorf("Error reading http_extra_string: %v", err)
		}
	}

	if err = d.Set("column", flattenSystemHealthCheckColumn(o["column"], d, "column", sv)); err != nil {
		if !fortiAPIPatch(o["column"]) {
			return fmt.Errorf("Error reading column: %v", err)
		}
	}

	if err = d.Set("http_connect", flattenSystemHealthCheckHttpConnect(o["http_connect"], d, "http_connect", sv)); err != nil {
		if !fortiAPIPatch(o["http_connect"]) {
			return fmt.Errorf("Error reading http_connect: %v", err)
		}
	}

	if err = d.Set("auth_appid", flattenSystemHealthCheckAuthAppid(o["auth-appid"], d, "auth_appid", sv)); err != nil {
		if !fortiAPIPatch(o["auth-appid"]) {
			return fmt.Errorf("Error reading auth-appid: %v", err)
		}
	}

	if err = d.Set("agent_type", flattenSystemHealthCheckAgentType(o["agent-type"], d, "agent_type", sv)); err != nil {
		if !fortiAPIPatch(o["agent-type"]) {
			return fmt.Errorf("Error reading agent-type: %v", err)
		}
	}

	if err = d.Set("filter", flattenSystemHealthCheckFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("remote_host", flattenSystemHealthCheckRemoteHost(o["remote_host"], d, "remote_host", sv)); err != nil {
		if !fortiAPIPatch(o["remote_host"]) {
			return fmt.Errorf("Error reading remote_host: %v", err)
		}
	}

	if err = d.Set("connect_type", flattenSystemHealthCheckConnectType(o["connect-type"], d, "connect_type", sv)); err != nil {
		if !fortiAPIPatch(o["connect-type"]) {
			return fmt.Errorf("Error reading connect-type: %v", err)
		}
	}

	if err = d.Set("method_type", flattenSystemHealthCheckMethodType(o["method_type"], d, "method_type", sv)); err != nil {
		if !fortiAPIPatch(o["method_type"]) {
			return fmt.Errorf("Error reading method_type: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers", flattenSystemHealthCheckSslCiphers(o["ssl-ciphers"], d, "ssl_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ciphers"]) {
			return fmt.Errorf("Error reading ssl-ciphers: %v", err)
		}
	}

	if err = d.Set("community", flattenSystemHealthCheckCommunity(o["community"], d, "community", sv)); err != nil {
		if !fortiAPIPatch(o["community"]) {
			return fmt.Errorf("Error reading community: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemHealthCheckScript(o["script"], d, "script", sv)); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("dest_addr", flattenSystemHealthCheckDestAddr(o["dest_addr"], d, "dest_addr", sv)); err != nil {
		if !fortiAPIPatch(o["dest_addr"]) {
			return fmt.Errorf("Error reading dest_addr: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemHealthCheckType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemHealthCheckUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("rtsp_method_type", flattenSystemHealthCheckRtspMethodType(o["rtsp-method-type"], d, "rtsp_method_type", sv)); err != nil {
		if !fortiAPIPatch(o["rtsp-method-type"]) {
			return fmt.Errorf("Error reading rtsp-method-type: %v", err)
		}
	}

	if err = d.Set("mssql_send_string", flattenSystemHealthCheckMssqlSendString(o["mssql-send-string"], d, "mssql_send_string", sv)); err != nil {
		if !fortiAPIPatch(o["mssql-send-string"]) {
			return fmt.Errorf("Error reading mssql-send-string: %v", err)
		}
	}

	if err = d.Set("string_value", flattenSystemHealthCheckStringValue(o["string-value"], d, "string_value", sv)); err != nil {
		if !fortiAPIPatch(o["string-value"]) {
			return fmt.Errorf("Error reading string-value: %v", err)
		}
	}

	if err = d.Set("down_retry", flattenSystemHealthCheckDownRetry(o["down_retry"], d, "down_retry", sv)); err != nil {
		if !fortiAPIPatch(o["down_retry"]) {
			return fmt.Errorf("Error reading down_retry: %v", err)
		}
	}

	if err = d.Set("origin_realm", flattenSystemHealthCheckOriginRealm(o["origin-realm"], d, "origin_realm", sv)); err != nil {
		if !fortiAPIPatch(o["origin-realm"]) {
			return fmt.Errorf("Error reading origin-realm: %v", err)
		}
	}

	if err = d.Set("origin_host", flattenSystemHealthCheckOriginHost(o["origin-host"], d, "origin_host", sv)); err != nil {
		if !fortiAPIPatch(o["origin-host"]) {
			return fmt.Errorf("Error reading origin-host: %v", err)
		}
	}

	if err = d.Set("secret_key", flattenSystemHealthCheckSecretKey(o["secret_key"], d, "secret_key", sv)); err != nil {
		if !fortiAPIPatch(o["secret_key"]) {
			return fmt.Errorf("Error reading secret_key: %v", err)
		}
	}

	if err = d.Set("compare_type", flattenSystemHealthCheckCompareType(o["compare-type"], d, "compare_type", sv)); err != nil {
		if !fortiAPIPatch(o["compare-type"]) {
			return fmt.Errorf("Error reading compare-type: %v", err)
		}
	}

	if err = d.Set("match_type", flattenSystemHealthCheckMatchType(o["match_type"], d, "match_type", sv)); err != nil {
		if !fortiAPIPatch(o["match_type"]) {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("attribute", flattenSystemHealthCheckAttribute(o["attribute"], d, "attribute", sv)); err != nil {
		if !fortiAPIPatch(o["attribute"]) {
			return fmt.Errorf("Error reading attribute: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemHealthCheckLocalCert(o["local-cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("Error reading local-cert: %v", err)
		}
	}

	if err = d.Set("mysql_server_type", flattenSystemHealthCheckMysqlServerType(o["mysql-server-type"], d, "mysql_server_type", sv)); err != nil {
		if !fortiAPIPatch(o["mysql-server-type"]) {
			return fmt.Errorf("Error reading mysql-server-type: %v", err)
		}
	}

	if err = d.Set("file", flattenSystemHealthCheckFile(o["file"], d, "file", sv)); err != nil {
		if !fortiAPIPatch(o["file"]) {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	if err = d.Set("dest_addr_type", flattenSystemHealthCheckDestAddrType(o["dest_addr_type"], d, "dest_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["dest_addr_type"]) {
			return fmt.Errorf("Error reading dest_addr_type: %v", err)
		}
	}

	if err = d.Set("cpu_weight", flattenSystemHealthCheckCpuWeight(o["cpu-weight"], d, "cpu_weight", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-weight"]) {
			return fmt.Errorf("Error reading cpu-weight: %v", err)
		}
	}

	if err = d.Set("host_ip_addr", flattenSystemHealthCheckHostIpAddr(o["host-ip-addr"], d, "host_ip_addr", sv)); err != nil {
		if !fortiAPIPatch(o["host-ip-addr"]) {
			return fmt.Errorf("Error reading host-ip-addr: %v", err)
		}
	}

	if err = d.Set("host_ip6_addr", flattenSystemHealthCheckHostIp6Addr(o["host-ip6-addr"], d, "host_ip6_addr", sv)); err != nil {
		if !fortiAPIPatch(o["host-ip6-addr"]) {
			return fmt.Errorf("Error reading host-ip6-addr: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemHealthCheckHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("mem_weight", flattenSystemHealthCheckMemWeight(o["mem-weight"], d, "mem_weight", sv)); err != nil {
		if !fortiAPIPatch(o["mem-weight"]) {
			return fmt.Errorf("Error reading mem-weight: %v", err)
		}
	}

	if err = d.Set("sid", flattenSystemHealthCheckSid(o["sid"], d, "sid", sv)); err != nil {
		if !fortiAPIPatch(o["sid"]) {
			return fmt.Errorf("Error reading sid: %v", err)
		}
	}

	if err = d.Set("disk_weight", flattenSystemHealthCheckDiskWeight(o["disk-weight"], d, "disk_weight", sv)); err != nil {
		if !fortiAPIPatch(o["disk-weight"]) {
			return fmt.Errorf("Error reading disk-weight: %v", err)
		}
	}

	if err = d.Set("receive_string", flattenSystemHealthCheckReceiveString(o["receive_string"], d, "receive_string", sv)); err != nil {
		if !fortiAPIPatch(o["receive_string"]) {
			return fmt.Errorf("Error reading receive_string: %v", err)
		}
	}

	if err = d.Set("oid", flattenSystemHealthCheckOid(o["oid"], d, "oid", sv)); err != nil {
		if !fortiAPIPatch(o["oid"]) {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if err = d.Set("value_type", flattenSystemHealthCheckValueType(o["value-type"], d, "value_type", sv)); err != nil {
		if !fortiAPIPatch(o["value-type"]) {
			return fmt.Errorf("Error reading value-type: %v", err)
		}
	}

	if err = d.Set("radius_reject", flattenSystemHealthCheckRadiusReject(o["radius-reject"], d, "radius_reject", sv)); err != nil {
		if !fortiAPIPatch(o["radius-reject"]) {
			return fmt.Errorf("Error reading radius-reject: %v", err)
		}
	}

	if err = d.Set("rtsp_describe_url", flattenSystemHealthCheckRtspDescribeUrl(o["rtsp-describe-url"], d, "rtsp_describe_url", sv)); err != nil {
		if !fortiAPIPatch(o["rtsp-describe-url"]) {
			return fmt.Errorf("Error reading rtsp-describe-url: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemHealthCheckMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("database", flattenSystemHealthCheckDatabase(o["database"], d, "database", sv)); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("oracle_send_string", flattenSystemHealthCheckOracleSendString(o["oracle-send-string"], d, "oracle_send_string", sv)); err != nil {
		if !fortiAPIPatch(o["oracle-send-string"]) {
			return fmt.Errorf("Error reading oracle-send-string: %v", err)
		}
	}

	if err = d.Set("send_string", flattenSystemHealthCheckSendString(o["send_string"], d, "send_string", sv)); err != nil {
		if !fortiAPIPatch(o["send_string"]) {
			return fmt.Errorf("Error reading send_string: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemHealthCheckInterval(o["interval"], d, "interval", sv)); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("up_retry", flattenSystemHealthCheckUpRetry(o["up_retry"], d, "up_retry", sv)); err != nil {
		if !fortiAPIPatch(o["up_retry"]) {
			return fmt.Errorf("Error reading up_retry: %v", err)
		}
	}

	if err = d.Set("remote_username", flattenSystemHealthCheckRemoteUsername(o["remote-username"], d, "remote_username", sv)); err != nil {
		if !fortiAPIPatch(o["remote-username"]) {
			return fmt.Errorf("Error reading remote-username: %v", err)
		}
	}

	if err = d.Set("cpu", flattenSystemHealthCheckCpu(o["cpu"], d, "cpu", sv)); err != nil {
		if !fortiAPIPatch(o["cpu"]) {
			return fmt.Errorf("Error reading cpu: %v", err)
		}
	}

	if err = d.Set("vendor_id", flattenSystemHealthCheckVendorId(o["vendor-id"], d, "vendor_id", sv)); err != nil {
		if !fortiAPIPatch(o["vendor-id"]) {
			return fmt.Errorf("Error reading vendor-id: %v", err)
		}
	}

	return nil
}

func expandSystemHealthCheckCounterValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckBasedn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckPwdType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDestAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMssqlRow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDisk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHttpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMssqlColumn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckStatusCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFolder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMssqlReceiveString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckOracleReceiveString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHostAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckSipRequestType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRemotePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckAcctAppid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckAllowSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRemotePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckConnectString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckProductName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckNasIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckBinddn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHostAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHttpExtraString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckColumn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHttpConnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckAuthAppid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckAgentType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRemoteHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckConnectType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMethodType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckSslCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDestAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRtspMethodType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMssqlSendString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckStringValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDownRetry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckOriginRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckOriginHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckSecretKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckCompareType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMatchType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckLocalCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMysqlServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDestAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckCpuWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHostIpAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHostIp6Addr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMemWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckSid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDiskWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckReceiveString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckOid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckValueType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRadiusReject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRtspDescribeUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckOracleSendString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckSendString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckUpRetry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckRemoteUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckCpu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckVendorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHealthCheck(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("counter_value"); ok {
		t, err := expandSystemHealthCheckCounterValue(d, v, "counter_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["counter-value"] = t
		}
	}

	if v, ok := d.GetOk("service_name"); ok {
		t, err := expandSystemHealthCheckServiceName(d, v, "service_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service_name"] = t
		}
	}

	if v, ok := d.GetOk("basedn"); ok {
		t, err := expandSystemHealthCheckBasedn(d, v, "basedn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["baseDN"] = t
		}
	}

	if v, ok := d.GetOk("pwd_type"); ok {
		t, err := expandSystemHealthCheckPwdType(d, v, "pwd_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pwd_type"] = t
		}
	}

	if v, ok := d.GetOk("dest_addr6"); ok {
		t, err := expandSystemHealthCheckDestAddr6(d, v, "dest_addr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest_addr6"] = t
		}
	}

	if v, ok := d.GetOk("mssql_row"); ok {
		t, err := expandSystemHealthCheckMssqlRow(d, v, "mssql_row", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mssql-row"] = t
		}
	}

	if v, ok := d.GetOk("disk"); ok {
		t, err := expandSystemHealthCheckDisk(d, v, "disk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disk"] = t
		}
	}

	if v, ok := d.GetOk("row"); ok {
		t, err := expandSystemHealthCheckRow(d, v, "row", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["row"] = t
		}
	}

	if v, ok := d.GetOk("http_version"); ok {
		t, err := expandSystemHealthCheckHttpVersion(d, v, "http_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_version"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok {
		t, err := expandSystemHealthCheckPassive(d, v, "passive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("mssql_column"); ok {
		t, err := expandSystemHealthCheckMssqlColumn(d, v, "mssql_column", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mssql-column"] = t
		}
	}

	if v, ok := d.GetOk("status_code"); ok {
		t, err := expandSystemHealthCheckStatusCode(d, v, "status_code", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status_code"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok {
		t, err := expandSystemHealthCheckFolder(d, v, "folder", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandSystemHealthCheckAddrType(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr_type"] = t
		}
	}

	if v, ok := d.GetOk("mssql_receive_string"); ok {
		t, err := expandSystemHealthCheckMssqlReceiveString(d, v, "mssql_receive_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mssql-receive-string"] = t
		}
	}

	if v, ok := d.GetOk("oracle_receive_string"); ok {
		t, err := expandSystemHealthCheckOracleReceiveString(d, v, "oracle_receive_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oracle-receive-string"] = t
		}
	}

	if v, ok := d.GetOk("host_addr"); ok {
		t, err := expandSystemHealthCheckHostAddr(d, v, "host_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_addr"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemHealthCheckPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("sip_request_type"); ok {
		t, err := expandSystemHealthCheckSipRequestType(d, v, "sip_request_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip_request_type"] = t
		}
	}

	if v, ok := d.GetOk("remote_password"); ok {
		t, err := expandSystemHealthCheckRemotePassword(d, v, "remote_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-password"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandSystemHealthCheckTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("acct_appid"); ok {
		t, err := expandSystemHealthCheckAcctAppid(d, v, "acct_appid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-appid"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemHealthCheckPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("allow_ssl_version"); ok {
		t, err := expandSystemHealthCheckAllowSslVersion(d, v, "allow_ssl_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_port"); ok {
		t, err := expandSystemHealthCheckRemotePort(d, v, "remote_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote_port"] = t
		}
	}

	if v, ok := d.GetOk("connect_string"); ok {
		t, err := expandSystemHealthCheckConnectString(d, v, "connect_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-string"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandSystemHealthCheckDomainName(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain_name"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok {
		t, err := expandSystemHealthCheckVersion(d, v, "version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("product_name"); ok {
		t, err := expandSystemHealthCheckProductName(d, v, "product_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["product-name"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok {
		t, err := expandSystemHealthCheckNasIp(d, v, "nas_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("binddn"); ok {
		t, err := expandSystemHealthCheckBinddn(d, v, "binddn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bindDN"] = t
		}
	}

	if v, ok := d.GetOk("mem"); ok {
		t, err := expandSystemHealthCheckMem(d, v, "mem", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mem"] = t
		}
	}

	if v, ok := d.GetOk("host_addr6"); ok {
		t, err := expandSystemHealthCheckHostAddr6(d, v, "host_addr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_addr6"] = t
		}
	}

	if v, ok := d.GetOk("http_extra_string"); ok {
		t, err := expandSystemHealthCheckHttpExtraString(d, v, "http_extra_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_extra_string"] = t
		}
	}

	if v, ok := d.GetOk("column"); ok {
		t, err := expandSystemHealthCheckColumn(d, v, "column", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column"] = t
		}
	}

	if v, ok := d.GetOk("http_connect"); ok {
		t, err := expandSystemHealthCheckHttpConnect(d, v, "http_connect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_connect"] = t
		}
	}

	if v, ok := d.GetOk("auth_appid"); ok {
		t, err := expandSystemHealthCheckAuthAppid(d, v, "auth_appid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-appid"] = t
		}
	}

	if v, ok := d.GetOk("agent_type"); ok {
		t, err := expandSystemHealthCheckAgentType(d, v, "agent_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-type"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		t, err := expandSystemHealthCheckFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("remote_host"); ok {
		t, err := expandSystemHealthCheckRemoteHost(d, v, "remote_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote_host"] = t
		}
	}

	if v, ok := d.GetOk("connect_type"); ok {
		t, err := expandSystemHealthCheckConnectType(d, v, "connect_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-type"] = t
		}
	}

	if v, ok := d.GetOk("method_type"); ok {
		t, err := expandSystemHealthCheckMethodType(d, v, "method_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method_type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ciphers"); ok {
		t, err := expandSystemHealthCheckSslCiphers(d, v, "ssl_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("community"); ok {
		t, err := expandSystemHealthCheckCommunity(d, v, "community", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["community"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok {
		t, err := expandSystemHealthCheckScript(d, v, "script", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	if v, ok := d.GetOk("dest_addr"); ok {
		t, err := expandSystemHealthCheckDestAddr(d, v, "dest_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest_addr"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemHealthCheckType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemHealthCheckUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("rtsp_method_type"); ok {
		t, err := expandSystemHealthCheckRtspMethodType(d, v, "rtsp_method_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtsp-method-type"] = t
		}
	}

	if v, ok := d.GetOk("mssql_send_string"); ok {
		t, err := expandSystemHealthCheckMssqlSendString(d, v, "mssql_send_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mssql-send-string"] = t
		}
	}

	if v, ok := d.GetOk("string_value"); ok {
		t, err := expandSystemHealthCheckStringValue(d, v, "string_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["string-value"] = t
		}
	}

	if v, ok := d.GetOk("down_retry"); ok {
		t, err := expandSystemHealthCheckDownRetry(d, v, "down_retry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["down_retry"] = t
		}
	}

	if v, ok := d.GetOk("origin_realm"); ok {
		t, err := expandSystemHealthCheckOriginRealm(d, v, "origin_realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin-realm"] = t
		}
	}

	if v, ok := d.GetOk("origin_host"); ok {
		t, err := expandSystemHealthCheckOriginHost(d, v, "origin_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin-host"] = t
		}
	}

	if v, ok := d.GetOk("secret_key"); ok {
		t, err := expandSystemHealthCheckSecretKey(d, v, "secret_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret_key"] = t
		}
	}

	if v, ok := d.GetOk("compare_type"); ok {
		t, err := expandSystemHealthCheckCompareType(d, v, "compare_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compare-type"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok {
		t, err := expandSystemHealthCheckMatchType(d, v, "match_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match_type"] = t
		}
	}

	if v, ok := d.GetOk("attribute"); ok {
		t, err := expandSystemHealthCheckAttribute(d, v, "attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandSystemHealthCheckLocalCert(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("mysql_server_type"); ok {
		t, err := expandSystemHealthCheckMysqlServerType(d, v, "mysql_server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mysql-server-type"] = t
		}
	}

	if v, ok := d.GetOk("file"); ok {
		t, err := expandSystemHealthCheckFile(d, v, "file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file"] = t
		}
	}

	if v, ok := d.GetOk("dest_addr_type"); ok {
		t, err := expandSystemHealthCheckDestAddrType(d, v, "dest_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest_addr_type"] = t
		}
	}

	if v, ok := d.GetOk("cpu_weight"); ok {
		t, err := expandSystemHealthCheckCpuWeight(d, v, "cpu_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-weight"] = t
		}
	}

	if v, ok := d.GetOk("host_ip_addr"); ok {
		t, err := expandSystemHealthCheckHostIpAddr(d, v, "host_ip_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-ip-addr"] = t
		}
	}

	if v, ok := d.GetOk("host_ip6_addr"); ok {
		t, err := expandSystemHealthCheckHostIp6Addr(d, v, "host_ip6_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-ip6-addr"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandSystemHealthCheckHostname(d, v, "hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("mem_weight"); ok {
		t, err := expandSystemHealthCheckMemWeight(d, v, "mem_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mem-weight"] = t
		}
	}

	if v, ok := d.GetOk("sid"); ok {
		t, err := expandSystemHealthCheckSid(d, v, "sid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sid"] = t
		}
	}

	if v, ok := d.GetOk("disk_weight"); ok {
		t, err := expandSystemHealthCheckDiskWeight(d, v, "disk_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disk-weight"] = t
		}
	}

	if v, ok := d.GetOk("receive_string"); ok {
		t, err := expandSystemHealthCheckReceiveString(d, v, "receive_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["receive_string"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok {
		t, err := expandSystemHealthCheckOid(d, v, "oid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("value_type"); ok {
		t, err := expandSystemHealthCheckValueType(d, v, "value_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-type"] = t
		}
	}

	if v, ok := d.GetOk("radius_reject"); ok {
		t, err := expandSystemHealthCheckRadiusReject(d, v, "radius_reject", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-reject"] = t
		}
	}

	if v, ok := d.GetOk("rtsp_describe_url"); ok {
		t, err := expandSystemHealthCheckRtspDescribeUrl(d, v, "rtsp_describe_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtsp-describe-url"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemHealthCheckMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok {
		t, err := expandSystemHealthCheckDatabase(d, v, "database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("oracle_send_string"); ok {
		t, err := expandSystemHealthCheckOracleSendString(d, v, "oracle_send_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oracle-send-string"] = t
		}
	}

	if v, ok := d.GetOk("send_string"); ok {
		t, err := expandSystemHealthCheckSendString(d, v, "send_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send_string"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		t, err := expandSystemHealthCheckInterval(d, v, "interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("up_retry"); ok {
		t, err := expandSystemHealthCheckUpRetry(d, v, "up_retry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["up_retry"] = t
		}
	}

	if v, ok := d.GetOk("remote_username"); ok {
		t, err := expandSystemHealthCheckRemoteUsername(d, v, "remote_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-username"] = t
		}
	}

	if v, ok := d.GetOk("cpu"); ok {
		t, err := expandSystemHealthCheckCpu(d, v, "cpu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu"] = t
		}
	}

	if v, ok := d.GetOk("vendor_id"); ok {
		t, err := expandSystemHealthCheckVendorId(d, v, "vendor_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor-id"] = t
		}
	}

	return &obj, nil
}
