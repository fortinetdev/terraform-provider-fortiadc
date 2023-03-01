// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content routing.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceContentRouting() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceContentRoutingRead,
		Update: resourceLoadBalanceContentRoutingUpdate,
		Create: resourceLoadBalanceContentRoutingCreate,
		Delete: resourceLoadBalanceContentRoutingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"source_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"schedule_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connection_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"packet_fwd_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"connection_pool_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persistence_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"schedule_pool_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"method_inherit": &schema.Schema{
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
func resourceLoadBalanceContentRoutingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRouting: type error")
	}

	obj, err := getObjectLoadBalanceContentRouting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRouting resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceContentRouting(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRouting resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceContentRoutingRead(d, m)
}
func resourceLoadBalanceContentRoutingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceContentRouting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRouting resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceContentRouting(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRouting resource: %v", err)
	}

	return resourceLoadBalanceContentRoutingRead(d, m)
}
func resourceLoadBalanceContentRoutingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceContentRouting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceContentRouting resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceContentRoutingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceContentRouting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRouting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceContentRouting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRouting resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceContentRoutingSourcePoolList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingScheduleList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingConnectionPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenLoadBalanceContentRoutingPacketFwdMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingConnectionPoolInherit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenLoadBalanceContentRoutingType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingPersistenceInherit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingSchedulePoolList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingMethodInherit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceContentRouting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("source_pool_list", flattenLoadBalanceContentRoutingSourcePoolList(o["source-pool-list"], d, "source_pool_list", sv)); err != nil {
		if !fortiAPIPatch(o["source-pool-list"]) {
			return fmt.Errorf("Error reading source-pool-list: %v", err)
		}
	}

	if err = d.Set("schedule_list", flattenLoadBalanceContentRoutingScheduleList(o["schedule-list"], d, "schedule_list", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-list"]) {
			return fmt.Errorf("Error reading schedule-list: %v", err)
		}
	}

	if err = d.Set("persistence", flattenLoadBalanceContentRoutingPersistence(o["persistence"], d, "persistence", sv)); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("connection_pool", flattenLoadBalanceContentRoutingConnectionPool(o["connection-pool"], d, "connection_pool", sv)); err != nil {
		if !fortiAPIPatch(o["connection-pool"]) {
			return fmt.Errorf("Error reading connection-pool: %v", err)
		}
	}

	if err = d.Set("ip", flattenLoadBalanceContentRoutingIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("packet_fwd_method", flattenLoadBalanceContentRoutingPacketFwdMethod(o["packet-fwd-method"], d, "packet_fwd_method", sv)); err != nil {
		if !fortiAPIPatch(o["packet-fwd-method"]) {
			return fmt.Errorf("Error reading packet-fwd-method: %v", err)
		}
	}

	if err = d.Set("comments", flattenLoadBalanceContentRoutingComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("connection_pool_inherit", flattenLoadBalanceContentRoutingConnectionPoolInherit(o["connection_pool_inherit"], d, "connection_pool_inherit", sv)); err != nil {
		if !fortiAPIPatch(o["connection_pool_inherit"]) {
			return fmt.Errorf("Error reading connection_pool_inherit: %v", err)
		}
	}

	if err = d.Set("ip6", flattenLoadBalanceContentRoutingIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceContentRoutingType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("pool", flattenLoadBalanceContentRoutingPool(o["pool"], d, "pool", sv)); err != nil {
		if !fortiAPIPatch(o["pool"]) {
			return fmt.Errorf("Error reading pool: %v", err)
		}
	}

	if err = d.Set("persistence_inherit", flattenLoadBalanceContentRoutingPersistenceInherit(o["persistence_inherit"], d, "persistence_inherit", sv)); err != nil {
		if !fortiAPIPatch(o["persistence_inherit"]) {
			return fmt.Errorf("Error reading persistence_inherit: %v", err)
		}
	}

	if err = d.Set("schedule_pool_list", flattenLoadBalanceContentRoutingSchedulePoolList(o["schedule-pool-list"], d, "schedule_pool_list", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-pool-list"]) {
			return fmt.Errorf("Error reading schedule-pool-list: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceContentRoutingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("method", flattenLoadBalanceContentRoutingMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("method_inherit", flattenLoadBalanceContentRoutingMethodInherit(o["method_inherit"], d, "method_inherit", sv)); err != nil {
		if !fortiAPIPatch(o["method_inherit"]) {
			return fmt.Errorf("Error reading method_inherit: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceContentRoutingSourcePoolList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingScheduleList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingConnectionPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingPacketFwdMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingConnectionPoolInherit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingPersistenceInherit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingSchedulePoolList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingMethodInherit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceContentRouting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("source_pool_list"); ok {
		t, err := expandLoadBalanceContentRoutingSourcePoolList(d, v, "source_pool_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-pool-list"] = t
		}
	}

	if v, ok := d.GetOk("schedule_list"); ok {
		t, err := expandLoadBalanceContentRoutingScheduleList(d, v, "schedule_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-list"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok {
		t, err := expandLoadBalanceContentRoutingPersistence(d, v, "persistence", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("connection_pool"); ok {
		t, err := expandLoadBalanceContentRoutingConnectionPool(d, v, "connection_pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-pool"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandLoadBalanceContentRoutingIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("packet_fwd_method"); ok {
		t, err := expandLoadBalanceContentRoutingPacketFwdMethod(d, v, "packet_fwd_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-fwd-method"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandLoadBalanceContentRoutingComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("connection_pool_inherit"); ok {
		t, err := expandLoadBalanceContentRoutingConnectionPoolInherit(d, v, "connection_pool_inherit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection_pool_inherit"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandLoadBalanceContentRoutingIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceContentRoutingType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("pool"); ok {
		t, err := expandLoadBalanceContentRoutingPool(d, v, "pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool"] = t
		}
	}

	if v, ok := d.GetOk("persistence_inherit"); ok {
		t, err := expandLoadBalanceContentRoutingPersistenceInherit(d, v, "persistence_inherit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence_inherit"] = t
		}
	}

	if v, ok := d.GetOk("schedule_pool_list"); ok {
		t, err := expandLoadBalanceContentRoutingSchedulePoolList(d, v, "schedule_pool_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-pool-list"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceContentRoutingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandLoadBalanceContentRoutingMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("method_inherit"); ok {
		t, err := expandLoadBalanceContentRoutingMethodInherit(d, v, "method_inherit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method_inherit"] = t
		}
	}

	return &obj, nil
}
