// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child mysql sharding.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceProfileChildMysqlSharding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceProfileChildMysqlShardingRead,
		Schema: map[string]*schema.Schema{
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"table": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
func dataSourceLoadBalanceProfileChildMysqlShardingRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding: type error")
	}

	o, err := c.ReadLoadBalanceProfileChildMysqlSharding(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceProfileChildMysqlSharding(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceProfileChildMysqlShardingDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlShardingGroupList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlShardingKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlShardingTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlShardingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceProfileChildMysqlShardingMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceProfileChildMysqlSharding(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("database", dataSourceFlattenLoadBalanceProfileChildMysqlShardingDatabase(o["database"], d, "database")); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("group_list", dataSourceFlattenLoadBalanceProfileChildMysqlShardingGroupList(o["group_list"], d, "group_list")); err != nil {
		if !fortiAPIPatch(o["group_list"]) {
			return fmt.Errorf("Error reading group_list: %v", err)
		}
	}

	if err = d.Set("key", dataSourceFlattenLoadBalanceProfileChildMysqlShardingKey(o["key"], d, "key")); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("table", dataSourceFlattenLoadBalanceProfileChildMysqlShardingTable(o["table"], d, "table")); err != nil {
		if !fortiAPIPatch(o["table"]) {
			return fmt.Errorf("Error reading table: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceProfileChildMysqlShardingType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceProfileChildMysqlShardingMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
