// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child mysql sharding.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceProfileChildMysqlSharding() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceProfileChildMysqlShardingRead,
		Update: resourceLoadBalanceProfileChildMysqlShardingUpdate,
		Create: resourceLoadBalanceProfileChildMysqlShardingCreate,
		Delete: resourceLoadBalanceProfileChildMysqlShardingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"table": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalanceProfileChildMysqlShardingCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceProfileChildMysqlSharding(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildMysqlSharding resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceProfileChildMysqlSharding(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildMysqlSharding resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceProfileChildMysqlShardingRead(d, m)
}
func resourceLoadBalanceProfileChildMysqlShardingUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceProfileChildMysqlSharding(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildMysqlSharding resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceProfileChildMysqlSharding(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildMysqlSharding resource: %v", err)
	}

	return resourceLoadBalanceProfileChildMysqlShardingRead(d, m)
}
func resourceLoadBalanceProfileChildMysqlShardingDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceProfileChildMysqlSharding(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceProfileChildMysqlSharding resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceProfileChildMysqlShardingRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlSharding: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceProfileChildMysqlSharding(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildMysqlSharding resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceProfileChildMysqlSharding(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildMysqlSharding resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceProfileChildMysqlShardingDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlShardingGroupList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlShardingKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlShardingTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlShardingType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlShardingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceProfileChildMysqlSharding(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("database", flattenLoadBalanceProfileChildMysqlShardingDatabase(o["database"], d, "database", sv)); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("group_list", flattenLoadBalanceProfileChildMysqlShardingGroupList(o["group_list"], d, "group_list", sv)); err != nil {
		if !fortiAPIPatch(o["group_list"]) {
			return fmt.Errorf("Error reading group_list: %v", err)
		}
	}

	if err = d.Set("key", flattenLoadBalanceProfileChildMysqlShardingKey(o["key"], d, "key", sv)); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("table", flattenLoadBalanceProfileChildMysqlShardingTable(o["table"], d, "table", sv)); err != nil {
		if !fortiAPIPatch(o["table"]) {
			return fmt.Errorf("Error reading table: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceProfileChildMysqlShardingType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceProfileChildMysqlShardingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceProfileChildMysqlShardingDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlShardingGroupList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlShardingKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlShardingTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlShardingType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlShardingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceProfileChildMysqlSharding(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("database"); ok {
		t, err := expandLoadBalanceProfileChildMysqlShardingDatabase(d, v, "database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("group_list"); ok {
		t, err := expandLoadBalanceProfileChildMysqlShardingGroupList(d, v, "group_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group_list"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandLoadBalanceProfileChildMysqlShardingKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("table"); ok {
		t, err := expandLoadBalanceProfileChildMysqlShardingTable(d, v, "table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceProfileChildMysqlShardingType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceProfileChildMysqlShardingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
