// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pool child pool member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePoolChildPoolMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePoolChildPoolMemberRead,
		Update: resourceLoadBalancePoolChildPoolMemberUpdate,
		Create: resourceLoadBalancePoolChildPoolMemberCreate,
		Delete: resourceLoadBalancePoolChildPoolMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"proxy_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hc_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connlimit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"recover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rs_profile_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"warmrate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mssql_read_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"m_health_check_relationship": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mysql_read_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"real_server_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mysql_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rs_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"warmup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"m_health_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connection_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"m_health_check_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"modify_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalancePoolChildPoolMemberCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalancePoolChildPoolMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePoolChildPoolMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalancePoolChildPoolMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePoolChildPoolMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalancePoolChildPoolMemberRead(d, m)
}
func resourceLoadBalancePoolChildPoolMemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalancePoolChildPoolMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePoolChildPoolMember resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePoolChildPoolMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePoolChildPoolMember resource: %v", err)
	}

	return resourceLoadBalancePoolChildPoolMemberRead(d, m)
}
func resourceLoadBalancePoolChildPoolMemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalancePoolChildPoolMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePoolChildPoolMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePoolChildPoolMemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
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

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalancePoolChildPoolMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePoolChildPoolMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePoolChildPoolMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePoolChildPoolMember resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePoolChildPoolMemberProxyProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberHcStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberConnlimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberRecover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberRsProfileInherit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberWarmrate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMssqlReadOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMHealthCheckRelationship(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMysqlReadOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberRealServerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberHealthCheckInherit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMysqlGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberRsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberWarmup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberConnectionRateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberMHealthCheckList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberModifyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePoolChildPoolMemberBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePoolChildPoolMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("proxy_protocol", flattenLoadBalancePoolChildPoolMemberProxyProtocol(o["proxy_protocol"], d, "proxy_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["proxy_protocol"]) {
			return fmt.Errorf("Error reading proxy_protocol: %v", err)
		}
	}

	if err = d.Set("hc_status", flattenLoadBalancePoolChildPoolMemberHcStatus(o["hc_status"], d, "hc_status", sv)); err != nil {
		if !fortiAPIPatch(o["hc_status"]) {
			return fmt.Errorf("Error reading hc_status: %v", err)
		}
	}

	if err = d.Set("connlimit", flattenLoadBalancePoolChildPoolMemberConnlimit(o["connlimit"], d, "connlimit", sv)); err != nil {
		if !fortiAPIPatch(o["connlimit"]) {
			return fmt.Errorf("Error reading connlimit: %v", err)
		}
	}

	if err = d.Set("recover", flattenLoadBalancePoolChildPoolMemberRecover(o["recover"], d, "recover", sv)); err != nil {
		if !fortiAPIPatch(o["recover"]) {
			return fmt.Errorf("Error reading recover: %v", err)
		}
	}

	if err = d.Set("port", flattenLoadBalancePoolChildPoolMemberPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("rs_profile_inherit", flattenLoadBalancePoolChildPoolMemberRsProfileInherit(o["rs_profile_inherit"], d, "rs_profile_inherit", sv)); err != nil {
		if !fortiAPIPatch(o["rs_profile_inherit"]) {
			return fmt.Errorf("Error reading rs_profile_inherit: %v", err)
		}
	}

	if err = d.Set("warmrate", flattenLoadBalancePoolChildPoolMemberWarmrate(o["warmrate"], d, "warmrate", sv)); err != nil {
		if !fortiAPIPatch(o["warmrate"]) {
			return fmt.Errorf("Error reading warmrate: %v", err)
		}
	}

	if err = d.Set("weight", flattenLoadBalancePoolChildPoolMemberWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("server_name", flattenLoadBalancePoolChildPoolMemberServerName(o["server_name"], d, "server_name", sv)); err != nil {
		if !fortiAPIPatch(o["server_name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("mssql_read_only", flattenLoadBalancePoolChildPoolMemberMssqlReadOnly(o["mssql_read_only"], d, "mssql_read_only", sv)); err != nil {
		if !fortiAPIPatch(o["mssql_read_only"]) {
			return fmt.Errorf("Error reading mssql_read_only: %v", err)
		}
	}

	if err = d.Set("m_health_check_relationship", flattenLoadBalancePoolChildPoolMemberMHealthCheckRelationship(o["m_health_check_relationship"], d, "m_health_check_relationship", sv)); err != nil {
		if !fortiAPIPatch(o["m_health_check_relationship"]) {
			return fmt.Errorf("Error reading m_health_check_relationship: %v", err)
		}
	}

	if err = d.Set("mysql_read_only", flattenLoadBalancePoolChildPoolMemberMysqlReadOnly(o["mysql_read_only"], d, "mysql_read_only", sv)); err != nil {
		if !fortiAPIPatch(o["mysql_read_only"]) {
			return fmt.Errorf("Error reading mysql_read_only: %v", err)
		}
	}

	if err = d.Set("status", flattenLoadBalancePoolChildPoolMemberStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("real_server_id", flattenLoadBalancePoolChildPoolMemberRealServerId(o["real_server_id"], d, "real_server_id", sv)); err != nil {
		if !fortiAPIPatch(o["real_server_id"]) {
			return fmt.Errorf("Error reading real_server_id: %v", err)
		}
	}

	if err = d.Set("health_check_inherit", flattenLoadBalancePoolChildPoolMemberHealthCheckInherit(o["health_check_inherit"], d, "health_check_inherit", sv)); err != nil {
		if !fortiAPIPatch(o["health_check_inherit"]) {
			return fmt.Errorf("Error reading health_check_inherit: %v", err)
		}
	}

	if err = d.Set("address6", flattenLoadBalancePoolChildPoolMemberAddress6(o["address6"], d, "address6", sv)); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("mysql_group_id", flattenLoadBalancePoolChildPoolMemberMysqlGroupId(o["mysql_group_id"], d, "mysql_group_id", sv)); err != nil {
		if !fortiAPIPatch(o["mysql_group_id"]) {
			return fmt.Errorf("Error reading mysql_group_id: %v", err)
		}
	}

	if err = d.Set("ssl", flattenLoadBalancePoolChildPoolMemberSsl(o["ssl"], d, "ssl", sv)); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("host", flattenLoadBalancePoolChildPoolMemberHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("cookie", flattenLoadBalancePoolChildPoolMemberCookie(o["cookie"], d, "cookie", sv)); err != nil {
		if !fortiAPIPatch(o["cookie"]) {
			return fmt.Errorf("Error reading cookie: %v", err)
		}
	}

	if err = d.Set("address", flattenLoadBalancePoolChildPoolMemberAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("rs_profile", flattenLoadBalancePoolChildPoolMemberRsProfile(o["rs_profile"], d, "rs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["rs_profile"]) {
			return fmt.Errorf("Error reading rs_profile: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePoolChildPoolMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("warmup", flattenLoadBalancePoolChildPoolMemberWarmup(o["warmup"], d, "warmup", sv)); err != nil {
		if !fortiAPIPatch(o["warmup"]) {
			return fmt.Errorf("Error reading warmup: %v", err)
		}
	}

	if err = d.Set("m_health_check", flattenLoadBalancePoolChildPoolMemberMHealthCheck(o["m_health_check"], d, "m_health_check", sv)); err != nil {
		if !fortiAPIPatch(o["m_health_check"]) {
			return fmt.Errorf("Error reading m_health_check: %v", err)
		}
	}

	if err = d.Set("connection_rate_limit", flattenLoadBalancePoolChildPoolMemberConnectionRateLimit(o["connection-rate-limit"], d, "connection_rate_limit", sv)); err != nil {
		if !fortiAPIPatch(o["connection-rate-limit"]) {
			return fmt.Errorf("Error reading connection-rate-limit: %v", err)
		}
	}

	if err = d.Set("m_health_check_list", flattenLoadBalancePoolChildPoolMemberMHealthCheckList(o["m_health_check_list"], d, "m_health_check_list", sv)); err != nil {
		if !fortiAPIPatch(o["m_health_check_list"]) {
			return fmt.Errorf("Error reading m_health_check_list: %v", err)
		}
	}

	if err = d.Set("modify_host", flattenLoadBalancePoolChildPoolMemberModifyHost(o["modify_host"], d, "modify_host", sv)); err != nil {
		if !fortiAPIPatch(o["modify_host"]) {
			return fmt.Errorf("Error reading modify_host: %v", err)
		}
	}

	if err = d.Set("backup", flattenLoadBalancePoolChildPoolMemberBackup(o["backup"], d, "backup", sv)); err != nil {
		if !fortiAPIPatch(o["backup"]) {
			return fmt.Errorf("Error reading backup: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePoolChildPoolMemberProxyProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberHcStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberConnlimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberRecover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberRsProfileInherit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberWarmrate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMssqlReadOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMHealthCheckRelationship(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMysqlReadOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberRealServerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberHealthCheckInherit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMysqlGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberRsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberWarmup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberConnectionRateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberMHealthCheckList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberModifyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePoolChildPoolMemberBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePoolChildPoolMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("proxy_protocol"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberProxyProtocol(d, v, "proxy_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy_protocol"] = t
		}
	}

	if v, ok := d.GetOk("hc_status"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberHcStatus(d, v, "hc_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hc_status"] = t
		}
	}

	if v, ok := d.GetOk("connlimit"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberConnlimit(d, v, "connlimit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connlimit"] = t
		}
	}

	if v, ok := d.GetOk("recover"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberRecover(d, v, "recover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recover"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("rs_profile_inherit"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberRsProfileInherit(d, v, "rs_profile_inherit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rs_profile_inherit"] = t
		}
	}

	if v, ok := d.GetOk("warmrate"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberWarmrate(d, v, "warmrate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warmrate"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberServerName(d, v, "server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server_name"] = t
		}
	}

	if v, ok := d.GetOk("mssql_read_only"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMssqlReadOnly(d, v, "mssql_read_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mssql_read_only"] = t
		}
	}

	if v, ok := d.GetOk("m_health_check_relationship"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMHealthCheckRelationship(d, v, "m_health_check_relationship", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["m_health_check_relationship"] = t
		}
	}

	if v, ok := d.GetOk("mysql_read_only"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMysqlReadOnly(d, v, "mysql_read_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mysql_read_only"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("real_server_id"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberRealServerId(d, v, "real_server_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["real_server_id"] = t
		}
	}

	if v, ok := d.GetOk("health_check_inherit"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberHealthCheckInherit(d, v, "health_check_inherit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health_check_inherit"] = t
		}
	}

	if v, ok := d.GetOk("address6"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberAddress6(d, v, "address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("mysql_group_id"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMysqlGroupId(d, v, "mysql_group_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mysql_group_id"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("cookie"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberCookie(d, v, "cookie", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("rs_profile"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberRsProfile(d, v, "rs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rs_profile"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("warmup"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberWarmup(d, v, "warmup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warmup"] = t
		}
	}

	if v, ok := d.GetOk("m_health_check"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMHealthCheck(d, v, "m_health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["m_health_check"] = t
		}
	}

	if v, ok := d.GetOk("connection_rate_limit"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberConnectionRateLimit(d, v, "connection_rate_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("m_health_check_list"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberMHealthCheckList(d, v, "m_health_check_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["m_health_check_list"] = t
		}
	}

	if v, ok := d.GetOk("modify_host"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberModifyHost(d, v, "modify_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modify_host"] = t
		}
	}

	if v, ok := d.GetOk("backup"); ok {
		t, err := expandLoadBalancePoolChildPoolMemberBackup(d, v, "backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup"] = t
		}
	}

	return &obj, nil
}
