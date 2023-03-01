// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system health check.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemHealthCheck() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemHealthCheckRead,
		Schema: map[string]*schema.Schema{
			"counter_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"basedn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pwd_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dest_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mssql_row": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"row": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mssql_column": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mssql_receive_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_receive_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sip_request_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"acct_appid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connect_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"binddn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_extra_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"column": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_connect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_appid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"agent_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connect_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"method_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dest_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rtsp_method_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mssql_send_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"string_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"down_retry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"origin_realm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"origin_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compare_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dest_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_ip6_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mem_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"receive_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"value_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"radius_reject": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rtsp_describe_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_send_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"up_retry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_id": &schema.Schema{
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
func dataSourceSystemHealthCheckRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHealthCheck(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemHealthCheck: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemHealthCheck(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemHealthCheck from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemHealthCheckCounterValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckBasedn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckPwdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDestAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMssqlRow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMssqlColumn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMssqlReceiveString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckOracleReceiveString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHostAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckSipRequestType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRemotePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckAcctAppid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckAllowSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRemotePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckConnectString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckProductName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckBinddn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHostAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHttpExtraString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckColumn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHttpConnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckAuthAppid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckAgentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRemoteHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckConnectType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMethodType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckSslCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDestAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRtspMethodType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMssqlSendString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckStringValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDownRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckOriginRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckOriginHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckSecretKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckCompareType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMatchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMysqlServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDestAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckCpuWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHostIpAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHostIp6Addr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMemWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckSid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDiskWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckReceiveString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckValueType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRadiusReject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRtspDescribeUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckOracleSendString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckSendString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckUpRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckRemoteUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHealthCheckVendorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemHealthCheck(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("counter_value", dataSourceFlattenSystemHealthCheckCounterValue(o["counter-value"], d, "counter_value")); err != nil {
		if !fortiAPIPatch(o["counter-value"]) {
			return fmt.Errorf("Error reading counter-value: %v", err)
		}
	}

	if err = d.Set("service_name", dataSourceFlattenSystemHealthCheckServiceName(o["service_name"], d, "service_name")); err != nil {
		if !fortiAPIPatch(o["service_name"]) {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("basedn", dataSourceFlattenSystemHealthCheckBasedn(o["baseDN"], d, "basedn")); err != nil {
		if !fortiAPIPatch(o["baseDN"]) {
			return fmt.Errorf("Error reading baseDN: %v", err)
		}
	}

	if err = d.Set("pwd_type", dataSourceFlattenSystemHealthCheckPwdType(o["pwd_type"], d, "pwd_type")); err != nil {
		if !fortiAPIPatch(o["pwd_type"]) {
			return fmt.Errorf("Error reading pwd_type: %v", err)
		}
	}

	if err = d.Set("dest_addr6", dataSourceFlattenSystemHealthCheckDestAddr6(o["dest_addr6"], d, "dest_addr6")); err != nil {
		if !fortiAPIPatch(o["dest_addr6"]) {
			return fmt.Errorf("Error reading dest_addr6: %v", err)
		}
	}

	if err = d.Set("mssql_row", dataSourceFlattenSystemHealthCheckMssqlRow(o["mssql-row"], d, "mssql_row")); err != nil {
		if !fortiAPIPatch(o["mssql-row"]) {
			return fmt.Errorf("Error reading mssql-row: %v", err)
		}
	}

	if err = d.Set("disk", dataSourceFlattenSystemHealthCheckDisk(o["disk"], d, "disk")); err != nil {
		if !fortiAPIPatch(o["disk"]) {
			return fmt.Errorf("Error reading disk: %v", err)
		}
	}

	if err = d.Set("row", dataSourceFlattenSystemHealthCheckRow(o["row"], d, "row")); err != nil {
		if !fortiAPIPatch(o["row"]) {
			return fmt.Errorf("Error reading row: %v", err)
		}
	}

	if err = d.Set("http_version", dataSourceFlattenSystemHealthCheckHttpVersion(o["http_version"], d, "http_version")); err != nil {
		if !fortiAPIPatch(o["http_version"]) {
			return fmt.Errorf("Error reading http_version: %v", err)
		}
	}

	if err = d.Set("passive", dataSourceFlattenSystemHealthCheckPassive(o["passive"], d, "passive")); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("mssql_column", dataSourceFlattenSystemHealthCheckMssqlColumn(o["mssql-column"], d, "mssql_column")); err != nil {
		if !fortiAPIPatch(o["mssql-column"]) {
			return fmt.Errorf("Error reading mssql-column: %v", err)
		}
	}

	if err = d.Set("status_code", dataSourceFlattenSystemHealthCheckStatusCode(o["status_code"], d, "status_code")); err != nil {
		if !fortiAPIPatch(o["status_code"]) {
			return fmt.Errorf("Error reading status_code: %v", err)
		}
	}

	if err = d.Set("folder", dataSourceFlattenSystemHealthCheckFolder(o["folder"], d, "folder")); err != nil {
		if !fortiAPIPatch(o["folder"]) {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("addr_type", dataSourceFlattenSystemHealthCheckAddrType(o["addr_type"], d, "addr_type")); err != nil {
		if !fortiAPIPatch(o["addr_type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("mssql_receive_string", dataSourceFlattenSystemHealthCheckMssqlReceiveString(o["mssql-receive-string"], d, "mssql_receive_string")); err != nil {
		if !fortiAPIPatch(o["mssql-receive-string"]) {
			return fmt.Errorf("Error reading mssql-receive-string: %v", err)
		}
	}

	if err = d.Set("oracle_receive_string", dataSourceFlattenSystemHealthCheckOracleReceiveString(o["oracle-receive-string"], d, "oracle_receive_string")); err != nil {
		if !fortiAPIPatch(o["oracle-receive-string"]) {
			return fmt.Errorf("Error reading oracle-receive-string: %v", err)
		}
	}

	if err = d.Set("host_addr", dataSourceFlattenSystemHealthCheckHostAddr(o["host_addr"], d, "host_addr")); err != nil {
		if !fortiAPIPatch(o["host_addr"]) {
			return fmt.Errorf("Error reading host_addr: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenSystemHealthCheckPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("sip_request_type", dataSourceFlattenSystemHealthCheckSipRequestType(o["sip_request_type"], d, "sip_request_type")); err != nil {
		if !fortiAPIPatch(o["sip_request_type"]) {
			return fmt.Errorf("Error reading sip_request_type: %v", err)
		}
	}

	if err = d.Set("remote_password", dataSourceFlattenSystemHealthCheckRemotePassword(o["remote-password"], d, "remote_password")); err != nil {
		if !fortiAPIPatch(o["remote-password"]) {
			return fmt.Errorf("Error reading remote-password: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenSystemHealthCheckTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("acct_appid", dataSourceFlattenSystemHealthCheckAcctAppid(o["acct-appid"], d, "acct_appid")); err != nil {
		if !fortiAPIPatch(o["acct-appid"]) {
			return fmt.Errorf("Error reading acct-appid: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemHealthCheckPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("allow_ssl_version", dataSourceFlattenSystemHealthCheckAllowSslVersion(o["allow-ssl-version"], d, "allow_ssl_version")); err != nil {
		if !fortiAPIPatch(o["allow-ssl-version"]) {
			return fmt.Errorf("Error reading allow-ssl-version: %v", err)
		}
	}

	if err = d.Set("remote_port", dataSourceFlattenSystemHealthCheckRemotePort(o["remote_port"], d, "remote_port")); err != nil {
		if !fortiAPIPatch(o["remote_port"]) {
			return fmt.Errorf("Error reading remote_port: %v", err)
		}
	}

	if err = d.Set("connect_string", dataSourceFlattenSystemHealthCheckConnectString(o["connect-string"], d, "connect_string")); err != nil {
		if !fortiAPIPatch(o["connect-string"]) {
			return fmt.Errorf("Error reading connect-string: %v", err)
		}
	}

	if err = d.Set("domain_name", dataSourceFlattenSystemHealthCheckDomainName(o["domain_name"], d, "domain_name")); err != nil {
		if !fortiAPIPatch(o["domain_name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("version", dataSourceFlattenSystemHealthCheckVersion(o["version"], d, "version")); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("product_name", dataSourceFlattenSystemHealthCheckProductName(o["product-name"], d, "product_name")); err != nil {
		if !fortiAPIPatch(o["product-name"]) {
			return fmt.Errorf("Error reading product-name: %v", err)
		}
	}

	if err = d.Set("nas_ip", dataSourceFlattenSystemHealthCheckNasIp(o["nas-ip"], d, "nas_ip")); err != nil {
		if !fortiAPIPatch(o["nas-ip"]) {
			return fmt.Errorf("Error reading nas-ip: %v", err)
		}
	}

	if err = d.Set("binddn", dataSourceFlattenSystemHealthCheckBinddn(o["bindDN"], d, "binddn")); err != nil {
		if !fortiAPIPatch(o["bindDN"]) {
			return fmt.Errorf("Error reading bindDN: %v", err)
		}
	}

	if err = d.Set("mem", dataSourceFlattenSystemHealthCheckMem(o["mem"], d, "mem")); err != nil {
		if !fortiAPIPatch(o["mem"]) {
			return fmt.Errorf("Error reading mem: %v", err)
		}
	}

	if err = d.Set("host_addr6", dataSourceFlattenSystemHealthCheckHostAddr6(o["host_addr6"], d, "host_addr6")); err != nil {
		if !fortiAPIPatch(o["host_addr6"]) {
			return fmt.Errorf("Error reading host_addr6: %v", err)
		}
	}

	if err = d.Set("http_extra_string", dataSourceFlattenSystemHealthCheckHttpExtraString(o["http_extra_string"], d, "http_extra_string")); err != nil {
		if !fortiAPIPatch(o["http_extra_string"]) {
			return fmt.Errorf("Error reading http_extra_string: %v", err)
		}
	}

	if err = d.Set("column", dataSourceFlattenSystemHealthCheckColumn(o["column"], d, "column")); err != nil {
		if !fortiAPIPatch(o["column"]) {
			return fmt.Errorf("Error reading column: %v", err)
		}
	}

	if err = d.Set("http_connect", dataSourceFlattenSystemHealthCheckHttpConnect(o["http_connect"], d, "http_connect")); err != nil {
		if !fortiAPIPatch(o["http_connect"]) {
			return fmt.Errorf("Error reading http_connect: %v", err)
		}
	}

	if err = d.Set("auth_appid", dataSourceFlattenSystemHealthCheckAuthAppid(o["auth-appid"], d, "auth_appid")); err != nil {
		if !fortiAPIPatch(o["auth-appid"]) {
			return fmt.Errorf("Error reading auth-appid: %v", err)
		}
	}

	if err = d.Set("agent_type", dataSourceFlattenSystemHealthCheckAgentType(o["agent-type"], d, "agent_type")); err != nil {
		if !fortiAPIPatch(o["agent-type"]) {
			return fmt.Errorf("Error reading agent-type: %v", err)
		}
	}

	if err = d.Set("filter", dataSourceFlattenSystemHealthCheckFilter(o["filter"], d, "filter")); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("remote_host", dataSourceFlattenSystemHealthCheckRemoteHost(o["remote_host"], d, "remote_host")); err != nil {
		if !fortiAPIPatch(o["remote_host"]) {
			return fmt.Errorf("Error reading remote_host: %v", err)
		}
	}

	if err = d.Set("connect_type", dataSourceFlattenSystemHealthCheckConnectType(o["connect-type"], d, "connect_type")); err != nil {
		if !fortiAPIPatch(o["connect-type"]) {
			return fmt.Errorf("Error reading connect-type: %v", err)
		}
	}

	if err = d.Set("method_type", dataSourceFlattenSystemHealthCheckMethodType(o["method_type"], d, "method_type")); err != nil {
		if !fortiAPIPatch(o["method_type"]) {
			return fmt.Errorf("Error reading method_type: %v", err)
		}
	}

	if err = d.Set("ssl_ciphers", dataSourceFlattenSystemHealthCheckSslCiphers(o["ssl-ciphers"], d, "ssl_ciphers")); err != nil {
		if !fortiAPIPatch(o["ssl-ciphers"]) {
			return fmt.Errorf("Error reading ssl-ciphers: %v", err)
		}
	}

	if err = d.Set("community", dataSourceFlattenSystemHealthCheckCommunity(o["community"], d, "community")); err != nil {
		if !fortiAPIPatch(o["community"]) {
			return fmt.Errorf("Error reading community: %v", err)
		}
	}

	if err = d.Set("script", dataSourceFlattenSystemHealthCheckScript(o["script"], d, "script")); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("dest_addr", dataSourceFlattenSystemHealthCheckDestAddr(o["dest_addr"], d, "dest_addr")); err != nil {
		if !fortiAPIPatch(o["dest_addr"]) {
			return fmt.Errorf("Error reading dest_addr: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemHealthCheckType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemHealthCheckUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("rtsp_method_type", dataSourceFlattenSystemHealthCheckRtspMethodType(o["rtsp-method-type"], d, "rtsp_method_type")); err != nil {
		if !fortiAPIPatch(o["rtsp-method-type"]) {
			return fmt.Errorf("Error reading rtsp-method-type: %v", err)
		}
	}

	if err = d.Set("mssql_send_string", dataSourceFlattenSystemHealthCheckMssqlSendString(o["mssql-send-string"], d, "mssql_send_string")); err != nil {
		if !fortiAPIPatch(o["mssql-send-string"]) {
			return fmt.Errorf("Error reading mssql-send-string: %v", err)
		}
	}

	if err = d.Set("string_value", dataSourceFlattenSystemHealthCheckStringValue(o["string-value"], d, "string_value")); err != nil {
		if !fortiAPIPatch(o["string-value"]) {
			return fmt.Errorf("Error reading string-value: %v", err)
		}
	}

	if err = d.Set("down_retry", dataSourceFlattenSystemHealthCheckDownRetry(o["down_retry"], d, "down_retry")); err != nil {
		if !fortiAPIPatch(o["down_retry"]) {
			return fmt.Errorf("Error reading down_retry: %v", err)
		}
	}

	if err = d.Set("origin_realm", dataSourceFlattenSystemHealthCheckOriginRealm(o["origin-realm"], d, "origin_realm")); err != nil {
		if !fortiAPIPatch(o["origin-realm"]) {
			return fmt.Errorf("Error reading origin-realm: %v", err)
		}
	}

	if err = d.Set("origin_host", dataSourceFlattenSystemHealthCheckOriginHost(o["origin-host"], d, "origin_host")); err != nil {
		if !fortiAPIPatch(o["origin-host"]) {
			return fmt.Errorf("Error reading origin-host: %v", err)
		}
	}

	if err = d.Set("secret_key", dataSourceFlattenSystemHealthCheckSecretKey(o["secret_key"], d, "secret_key")); err != nil {
		if !fortiAPIPatch(o["secret_key"]) {
			return fmt.Errorf("Error reading secret_key: %v", err)
		}
	}

	if err = d.Set("compare_type", dataSourceFlattenSystemHealthCheckCompareType(o["compare-type"], d, "compare_type")); err != nil {
		if !fortiAPIPatch(o["compare-type"]) {
			return fmt.Errorf("Error reading compare-type: %v", err)
		}
	}

	if err = d.Set("match_type", dataSourceFlattenSystemHealthCheckMatchType(o["match_type"], d, "match_type")); err != nil {
		if !fortiAPIPatch(o["match_type"]) {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("attribute", dataSourceFlattenSystemHealthCheckAttribute(o["attribute"], d, "attribute")); err != nil {
		if !fortiAPIPatch(o["attribute"]) {
			return fmt.Errorf("Error reading attribute: %v", err)
		}
	}

	if err = d.Set("local_cert", dataSourceFlattenSystemHealthCheckLocalCert(o["local-cert"], d, "local_cert")); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("Error reading local-cert: %v", err)
		}
	}

	if err = d.Set("mysql_server_type", dataSourceFlattenSystemHealthCheckMysqlServerType(o["mysql-server-type"], d, "mysql_server_type")); err != nil {
		if !fortiAPIPatch(o["mysql-server-type"]) {
			return fmt.Errorf("Error reading mysql-server-type: %v", err)
		}
	}

	if err = d.Set("file", dataSourceFlattenSystemHealthCheckFile(o["file"], d, "file")); err != nil {
		if !fortiAPIPatch(o["file"]) {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	if err = d.Set("dest_addr_type", dataSourceFlattenSystemHealthCheckDestAddrType(o["dest_addr_type"], d, "dest_addr_type")); err != nil {
		if !fortiAPIPatch(o["dest_addr_type"]) {
			return fmt.Errorf("Error reading dest_addr_type: %v", err)
		}
	}

	if err = d.Set("cpu_weight", dataSourceFlattenSystemHealthCheckCpuWeight(o["cpu-weight"], d, "cpu_weight")); err != nil {
		if !fortiAPIPatch(o["cpu-weight"]) {
			return fmt.Errorf("Error reading cpu-weight: %v", err)
		}
	}

	if err = d.Set("host_ip_addr", dataSourceFlattenSystemHealthCheckHostIpAddr(o["host-ip-addr"], d, "host_ip_addr")); err != nil {
		if !fortiAPIPatch(o["host-ip-addr"]) {
			return fmt.Errorf("Error reading host-ip-addr: %v", err)
		}
	}

	if err = d.Set("host_ip6_addr", dataSourceFlattenSystemHealthCheckHostIp6Addr(o["host-ip6-addr"], d, "host_ip6_addr")); err != nil {
		if !fortiAPIPatch(o["host-ip6-addr"]) {
			return fmt.Errorf("Error reading host-ip6-addr: %v", err)
		}
	}

	if err = d.Set("hostname", dataSourceFlattenSystemHealthCheckHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("mem_weight", dataSourceFlattenSystemHealthCheckMemWeight(o["mem-weight"], d, "mem_weight")); err != nil {
		if !fortiAPIPatch(o["mem-weight"]) {
			return fmt.Errorf("Error reading mem-weight: %v", err)
		}
	}

	if err = d.Set("sid", dataSourceFlattenSystemHealthCheckSid(o["sid"], d, "sid")); err != nil {
		if !fortiAPIPatch(o["sid"]) {
			return fmt.Errorf("Error reading sid: %v", err)
		}
	}

	if err = d.Set("disk_weight", dataSourceFlattenSystemHealthCheckDiskWeight(o["disk-weight"], d, "disk_weight")); err != nil {
		if !fortiAPIPatch(o["disk-weight"]) {
			return fmt.Errorf("Error reading disk-weight: %v", err)
		}
	}

	if err = d.Set("receive_string", dataSourceFlattenSystemHealthCheckReceiveString(o["receive_string"], d, "receive_string")); err != nil {
		if !fortiAPIPatch(o["receive_string"]) {
			return fmt.Errorf("Error reading receive_string: %v", err)
		}
	}

	if err = d.Set("oid", dataSourceFlattenSystemHealthCheckOid(o["oid"], d, "oid")); err != nil {
		if !fortiAPIPatch(o["oid"]) {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if err = d.Set("value_type", dataSourceFlattenSystemHealthCheckValueType(o["value-type"], d, "value_type")); err != nil {
		if !fortiAPIPatch(o["value-type"]) {
			return fmt.Errorf("Error reading value-type: %v", err)
		}
	}

	if err = d.Set("radius_reject", dataSourceFlattenSystemHealthCheckRadiusReject(o["radius-reject"], d, "radius_reject")); err != nil {
		if !fortiAPIPatch(o["radius-reject"]) {
			return fmt.Errorf("Error reading radius-reject: %v", err)
		}
	}

	if err = d.Set("rtsp_describe_url", dataSourceFlattenSystemHealthCheckRtspDescribeUrl(o["rtsp-describe-url"], d, "rtsp_describe_url")); err != nil {
		if !fortiAPIPatch(o["rtsp-describe-url"]) {
			return fmt.Errorf("Error reading rtsp-describe-url: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemHealthCheckMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("database", dataSourceFlattenSystemHealthCheckDatabase(o["database"], d, "database")); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("oracle_send_string", dataSourceFlattenSystemHealthCheckOracleSendString(o["oracle-send-string"], d, "oracle_send_string")); err != nil {
		if !fortiAPIPatch(o["oracle-send-string"]) {
			return fmt.Errorf("Error reading oracle-send-string: %v", err)
		}
	}

	if err = d.Set("send_string", dataSourceFlattenSystemHealthCheckSendString(o["send_string"], d, "send_string")); err != nil {
		if !fortiAPIPatch(o["send_string"]) {
			return fmt.Errorf("Error reading send_string: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemHealthCheckInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("up_retry", dataSourceFlattenSystemHealthCheckUpRetry(o["up_retry"], d, "up_retry")); err != nil {
		if !fortiAPIPatch(o["up_retry"]) {
			return fmt.Errorf("Error reading up_retry: %v", err)
		}
	}

	if err = d.Set("remote_username", dataSourceFlattenSystemHealthCheckRemoteUsername(o["remote-username"], d, "remote_username")); err != nil {
		if !fortiAPIPatch(o["remote-username"]) {
			return fmt.Errorf("Error reading remote-username: %v", err)
		}
	}

	if err = d.Set("cpu", dataSourceFlattenSystemHealthCheckCpu(o["cpu"], d, "cpu")); err != nil {
		if !fortiAPIPatch(o["cpu"]) {
			return fmt.Errorf("Error reading cpu: %v", err)
		}
	}

	if err = d.Set("vendor_id", dataSourceFlattenSystemHealthCheckVendorId(o["vendor-id"], d, "vendor_id")); err != nil {
		if !fortiAPIPatch(o["vendor-id"]) {
			return fmt.Errorf("Error reading vendor-id: %v", err)
		}
	}

	return nil
}
