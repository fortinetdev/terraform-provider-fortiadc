// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child mysql rule.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceProfileChildMysqlRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceProfileChildMysqlRuleRead,
		Schema: map[string]*schema.Schema{
			"client_ip_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"table_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sql_list": &schema.Schema{
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
func dataSourceLoadBalanceProfileChildMysqlRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule: type error")
	}

	o, err := c.ReadLoadBalanceProfileChildMysqlRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceProfileChildMysqlRule(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceProfileChildMysqlRuleClientIpList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlRuleUserList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlRuleDbList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlRuleType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlRuleTableList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlRuleMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlRuleSqlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceProfileChildMysqlRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("client_ip_list", dataSourceFlattenLoadBalanceProfileChildMysqlRuleClientIpList(o["client_ip_list"], d, "client_ip_list")); err != nil {
		if !fortiAPIPatch(o["client_ip_list"]) {
			return fmt.Errorf("Error reading client_ip_list: %v", err)
		}
	}

	if err = d.Set("user_list", dataSourceFlattenLoadBalanceProfileChildMysqlRuleUserList(o["user_list"], d, "user_list")); err != nil {
		if !fortiAPIPatch(o["user_list"]) {
			return fmt.Errorf("Error reading user_list: %v", err)
		}
	}

	if err = d.Set("db_list", dataSourceFlattenLoadBalanceProfileChildMysqlRuleDbList(o["db_list"], d, "db_list")); err != nil {
		if !fortiAPIPatch(o["db_list"]) {
			return fmt.Errorf("Error reading db_list: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceProfileChildMysqlRuleType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("table_list", dataSourceFlattenLoadBalanceProfileChildMysqlRuleTableList(o["table_list"], d, "table_list")); err != nil {
		if !fortiAPIPatch(o["table_list"]) {
			return fmt.Errorf("Error reading table_list: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceProfileChildMysqlRuleMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("sql_list", dataSourceFlattenLoadBalanceProfileChildMysqlRuleSqlList(o["sql_list"], d, "sql_list")); err != nil {
		if !fortiAPIPatch(o["sql_list"]) {
			return fmt.Errorf("Error reading sql_list: %v", err)
		}
	}

	return nil
}
