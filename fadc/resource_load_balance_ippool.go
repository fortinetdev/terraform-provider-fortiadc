// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance ippool.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceIppool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceIppoolRead,
		Update: resourceLoadBalanceIppoolUpdate,
		Create: resourceLoadBalanceIppoolCreate,
		Delete: resourceLoadBalanceIppoolDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6_start": &schema.Schema{
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
func resourceLoadBalanceIppoolCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceIppool: type error")
	}

	obj, err := getObjectLoadBalanceIppool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceIppool resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceIppool(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceIppool resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceIppoolRead(d, m)
}
func resourceLoadBalanceIppoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceIppool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceIppool resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceIppool(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceIppool resource: %v", err)
	}

	return resourceLoadBalanceIppoolRead(d, m)
}
func resourceLoadBalanceIppoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceIppool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceIppool resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceIppoolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceIppool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceIppool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceIppool(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceIppool resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceIppoolPoolType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolIpEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolIp6End(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolIpStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceIppoolIp6Start(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceIppool(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("pool_type", flattenLoadBalanceIppoolPoolType(o["pool_type"], d, "pool_type", sv)); err != nil {
		if !fortiAPIPatch(o["pool_type"]) {
			return fmt.Errorf("Error reading pool_type: %v", err)
		}
	}

	if err = d.Set("ip_end", flattenLoadBalanceIppoolIpEnd(o["ip-end"], d, "ip_end", sv)); err != nil {
		if !fortiAPIPatch(o["ip-end"]) {
			return fmt.Errorf("Error reading ip-end: %v", err)
		}
	}

	if err = d.Set("ip6_end", flattenLoadBalanceIppoolIp6End(o["ip6-end"], d, "ip6_end", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-end"]) {
			return fmt.Errorf("Error reading ip6-end: %v", err)
		}
	}

	if err = d.Set("ip_start", flattenLoadBalanceIppoolIpStart(o["ip-start"], d, "ip_start", sv)); err != nil {
		if !fortiAPIPatch(o["ip-start"]) {
			return fmt.Errorf("Error reading ip-start: %v", err)
		}
	}

	if err = d.Set("interface", flattenLoadBalanceIppoolInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceIppoolMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("ip6_start", flattenLoadBalanceIppoolIp6Start(o["ip6-start"], d, "ip6_start", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-start"]) {
			return fmt.Errorf("Error reading ip6-start: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceIppoolPoolType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolIpEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolIp6End(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolIpStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceIppoolIp6Start(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceIppool(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pool_type"); ok {
		t, err := expandLoadBalanceIppoolPoolType(d, v, "pool_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool_type"] = t
		}
	}

	if v, ok := d.GetOk("ip_end"); ok {
		t, err := expandLoadBalanceIppoolIpEnd(d, v, "ip_end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-end"] = t
		}
	}

	if v, ok := d.GetOk("ip6_end"); ok {
		t, err := expandLoadBalanceIppoolIp6End(d, v, "ip6_end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-end"] = t
		}
	}

	if v, ok := d.GetOk("ip_start"); ok {
		t, err := expandLoadBalanceIppoolIpStart(d, v, "ip_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-start"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandLoadBalanceIppoolInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceIppoolMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("ip6_start"); ok {
		t, err := expandLoadBalanceIppoolIp6Start(d, v, "ip6_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-start"] = t
		}
	}

	return &obj, nil
}
