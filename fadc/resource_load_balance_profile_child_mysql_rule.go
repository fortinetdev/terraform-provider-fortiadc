// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child mysql rule.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceProfileChildMysqlRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceProfileChildMysqlRuleRead,
		Update: resourceLoadBalanceProfileChildMysqlRuleUpdate,
		Create: resourceLoadBalanceProfileChildMysqlRuleCreate,
		Delete: resourceLoadBalanceProfileChildMysqlRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"client_ip_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"user_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"db_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"table_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sql_list": &schema.Schema{
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
func resourceLoadBalanceProfileChildMysqlRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceProfileChildMysqlRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildMysqlRule resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceProfileChildMysqlRule(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildMysqlRule resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceProfileChildMysqlRuleRead(d, m)
}
func resourceLoadBalanceProfileChildMysqlRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceProfileChildMysqlRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildMysqlRule resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceProfileChildMysqlRule(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildMysqlRule resource: %v", err)
	}

	return resourceLoadBalanceProfileChildMysqlRuleRead(d, m)
}
func resourceLoadBalanceProfileChildMysqlRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceProfileChildMysqlRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceProfileChildMysqlRule resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceProfileChildMysqlRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlRule: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceProfileChildMysqlRule(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildMysqlRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceProfileChildMysqlRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildMysqlRule resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceProfileChildMysqlRuleClientIpList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlRuleUserList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlRuleDbList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlRuleType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlRuleTableList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlRuleMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlRuleSqlList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceProfileChildMysqlRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("client_ip_list", flattenLoadBalanceProfileChildMysqlRuleClientIpList(o["client_ip_list"], d, "client_ip_list", sv)); err != nil {
		if !fortiAPIPatch(o["client_ip_list"]) {
			return fmt.Errorf("Error reading client_ip_list: %v", err)
		}
	}

	if err = d.Set("user_list", flattenLoadBalanceProfileChildMysqlRuleUserList(o["user_list"], d, "user_list", sv)); err != nil {
		if !fortiAPIPatch(o["user_list"]) {
			return fmt.Errorf("Error reading user_list: %v", err)
		}
	}

	if err = d.Set("db_list", flattenLoadBalanceProfileChildMysqlRuleDbList(o["db_list"], d, "db_list", sv)); err != nil {
		if !fortiAPIPatch(o["db_list"]) {
			return fmt.Errorf("Error reading db_list: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceProfileChildMysqlRuleType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("table_list", flattenLoadBalanceProfileChildMysqlRuleTableList(o["table_list"], d, "table_list", sv)); err != nil {
		if !fortiAPIPatch(o["table_list"]) {
			return fmt.Errorf("Error reading table_list: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceProfileChildMysqlRuleMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("sql_list", flattenLoadBalanceProfileChildMysqlRuleSqlList(o["sql_list"], d, "sql_list", sv)); err != nil {
		if !fortiAPIPatch(o["sql_list"]) {
			return fmt.Errorf("Error reading sql_list: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceProfileChildMysqlRuleClientIpList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlRuleUserList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlRuleDbList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlRuleType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlRuleTableList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlRuleMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlRuleSqlList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceProfileChildMysqlRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("client_ip_list"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleClientIpList(d, v, "client_ip_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client_ip_list"] = t
		}
	}

	if v, ok := d.GetOk("user_list"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleUserList(d, v, "user_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user_list"] = t
		}
	}

	if v, ok := d.GetOk("db_list"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleDbList(d, v, "db_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["db_list"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("table_list"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleTableList(d, v, "table_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table_list"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("sql_list"); ok {
		t, err := expandLoadBalanceProfileChildMysqlRuleSqlList(d, v, "sql_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sql_list"] = t
		}
	}

	return &obj, nil
}
