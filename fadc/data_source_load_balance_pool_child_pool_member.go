// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pool child pool member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePoolChildPoolMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePoolChildPoolMemberRead,
		Schema: map[string]*schema.Schema{
			"proxy_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hc_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connlimit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"recover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rs_profile_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"warmrate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mssql_read_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"m_health_check_relationship": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_read_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"real_server_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_check_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rs_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"warmup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"m_health_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"m_health_check_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"modify_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceLoadBalancePoolChildPoolMemberRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePoolChildPoolMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalancePoolChildPoolMember: type error")
	}

	o, err := c.ReadLoadBalancePoolChildPoolMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePoolChildPoolMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePoolChildPoolMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePoolChildPoolMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePoolChildPoolMemberProxyProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberHcStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberConnlimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberRecover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberRsProfileInherit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberWarmrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMssqlReadOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMHealthCheckRelationship(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMysqlReadOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberRealServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberHealthCheckInherit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMysqlGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberRsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberWarmup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberConnectionRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberMHealthCheckList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberModifyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePoolChildPoolMemberBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePoolChildPoolMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("proxy_protocol", dataSourceFlattenLoadBalancePoolChildPoolMemberProxyProtocol(o["proxy_protocol"], d, "proxy_protocol")); err != nil {
		if !fortiAPIPatch(o["proxy_protocol"]) {
			return fmt.Errorf("Error reading proxy_protocol: %v", err)
		}
	}

	if err = d.Set("hc_status", dataSourceFlattenLoadBalancePoolChildPoolMemberHcStatus(o["hc_status"], d, "hc_status")); err != nil {
		if !fortiAPIPatch(o["hc_status"]) {
			return fmt.Errorf("Error reading hc_status: %v", err)
		}
	}

	if err = d.Set("connlimit", dataSourceFlattenLoadBalancePoolChildPoolMemberConnlimit(o["connlimit"], d, "connlimit")); err != nil {
		if !fortiAPIPatch(o["connlimit"]) {
			return fmt.Errorf("Error reading connlimit: %v", err)
		}
	}

	if err = d.Set("recover", dataSourceFlattenLoadBalancePoolChildPoolMemberRecover(o["recover"], d, "recover")); err != nil {
		if !fortiAPIPatch(o["recover"]) {
			return fmt.Errorf("Error reading recover: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenLoadBalancePoolChildPoolMemberPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("rs_profile_inherit", dataSourceFlattenLoadBalancePoolChildPoolMemberRsProfileInherit(o["rs_profile_inherit"], d, "rs_profile_inherit")); err != nil {
		if !fortiAPIPatch(o["rs_profile_inherit"]) {
			return fmt.Errorf("Error reading rs_profile_inherit: %v", err)
		}
	}

	if err = d.Set("warmrate", dataSourceFlattenLoadBalancePoolChildPoolMemberWarmrate(o["warmrate"], d, "warmrate")); err != nil {
		if !fortiAPIPatch(o["warmrate"]) {
			return fmt.Errorf("Error reading warmrate: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenLoadBalancePoolChildPoolMemberWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("server_name", dataSourceFlattenLoadBalancePoolChildPoolMemberServerName(o["server_name"], d, "server_name")); err != nil {
		if !fortiAPIPatch(o["server_name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("mssql_read_only", dataSourceFlattenLoadBalancePoolChildPoolMemberMssqlReadOnly(o["mssql_read_only"], d, "mssql_read_only")); err != nil {
		if !fortiAPIPatch(o["mssql_read_only"]) {
			return fmt.Errorf("Error reading mssql_read_only: %v", err)
		}
	}

	if err = d.Set("m_health_check_relationship", dataSourceFlattenLoadBalancePoolChildPoolMemberMHealthCheckRelationship(o["m_health_check_relationship"], d, "m_health_check_relationship")); err != nil {
		if !fortiAPIPatch(o["m_health_check_relationship"]) {
			return fmt.Errorf("Error reading m_health_check_relationship: %v", err)
		}
	}

	if err = d.Set("mysql_read_only", dataSourceFlattenLoadBalancePoolChildPoolMemberMysqlReadOnly(o["mysql_read_only"], d, "mysql_read_only")); err != nil {
		if !fortiAPIPatch(o["mysql_read_only"]) {
			return fmt.Errorf("Error reading mysql_read_only: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenLoadBalancePoolChildPoolMemberStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("real_server_id", dataSourceFlattenLoadBalancePoolChildPoolMemberRealServerId(o["real_server_id"], d, "real_server_id")); err != nil {
		if !fortiAPIPatch(o["real_server_id"]) {
			return fmt.Errorf("Error reading real_server_id: %v", err)
		}
	}

	if err = d.Set("health_check_inherit", dataSourceFlattenLoadBalancePoolChildPoolMemberHealthCheckInherit(o["health_check_inherit"], d, "health_check_inherit")); err != nil {
		if !fortiAPIPatch(o["health_check_inherit"]) {
			return fmt.Errorf("Error reading health_check_inherit: %v", err)
		}
	}

	if err = d.Set("address6", dataSourceFlattenLoadBalancePoolChildPoolMemberAddress6(o["address6"], d, "address6")); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("mysql_group_id", dataSourceFlattenLoadBalancePoolChildPoolMemberMysqlGroupId(o["mysql_group_id"], d, "mysql_group_id")); err != nil {
		if !fortiAPIPatch(o["mysql_group_id"]) {
			return fmt.Errorf("Error reading mysql_group_id: %v", err)
		}
	}

	if err = d.Set("ssl", dataSourceFlattenLoadBalancePoolChildPoolMemberSsl(o["ssl"], d, "ssl")); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("host", dataSourceFlattenLoadBalancePoolChildPoolMemberHost(o["host"], d, "host")); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("cookie", dataSourceFlattenLoadBalancePoolChildPoolMemberCookie(o["cookie"], d, "cookie")); err != nil {
		if !fortiAPIPatch(o["cookie"]) {
			return fmt.Errorf("Error reading cookie: %v", err)
		}
	}

	if err = d.Set("address", dataSourceFlattenLoadBalancePoolChildPoolMemberAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("rs_profile", dataSourceFlattenLoadBalancePoolChildPoolMemberRsProfile(o["rs_profile"], d, "rs_profile")); err != nil {
		if !fortiAPIPatch(o["rs_profile"]) {
			return fmt.Errorf("Error reading rs_profile: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePoolChildPoolMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("warmup", dataSourceFlattenLoadBalancePoolChildPoolMemberWarmup(o["warmup"], d, "warmup")); err != nil {
		if !fortiAPIPatch(o["warmup"]) {
			return fmt.Errorf("Error reading warmup: %v", err)
		}
	}

	if err = d.Set("m_health_check", dataSourceFlattenLoadBalancePoolChildPoolMemberMHealthCheck(o["m_health_check"], d, "m_health_check")); err != nil {
		if !fortiAPIPatch(o["m_health_check"]) {
			return fmt.Errorf("Error reading m_health_check: %v", err)
		}
	}

	if err = d.Set("connection_rate_limit", dataSourceFlattenLoadBalancePoolChildPoolMemberConnectionRateLimit(o["connection-rate-limit"], d, "connection_rate_limit")); err != nil {
		if !fortiAPIPatch(o["connection-rate-limit"]) {
			return fmt.Errorf("Error reading connection-rate-limit: %v", err)
		}
	}

	if err = d.Set("m_health_check_list", dataSourceFlattenLoadBalancePoolChildPoolMemberMHealthCheckList(o["m_health_check_list"], d, "m_health_check_list")); err != nil {
		if !fortiAPIPatch(o["m_health_check_list"]) {
			return fmt.Errorf("Error reading m_health_check_list: %v", err)
		}
	}

	if err = d.Set("modify_host", dataSourceFlattenLoadBalancePoolChildPoolMemberModifyHost(o["modify_host"], d, "modify_host")); err != nil {
		if !fortiAPIPatch(o["modify_host"]) {
			return fmt.Errorf("Error reading modify_host: %v", err)
		}
	}

	if err = d.Set("backup", dataSourceFlattenLoadBalancePoolChildPoolMemberBackup(o["backup"], d, "backup")); err != nil {
		if !fortiAPIPatch(o["backup"]) {
			return fmt.Errorf("Error reading backup: %v", err)
		}
	}

	return nil
}
