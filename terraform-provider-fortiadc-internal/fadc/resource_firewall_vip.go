// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  firewall vip.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallVip() *schema.Resource {
	return &schema.Resource{
		Read:   resourceFirewallVipRead,
		Update: resourceFirewallVipUpdate,
		Create: resourceFirewallVipCreate,
		Delete: resourceFirewallVipDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mappedip_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mappedport_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"no_nat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"extif": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mappedip_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mappedport_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"extip": &schema.Schema{
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
		},
	}
}
func resourceFirewallVipCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallVip: type error")
	}

	obj, err := getObjectFirewallVip(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallVip(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip resource: %v", err)
	}

	d.SetId(mkey)

	return resourceFirewallVipRead(d, m)
}
func resourceFirewallVipUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectFirewallVip(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallVip(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip resource: %v", err)
	}

	return resourceFirewallVipRead(d, m)
}
func resourceFirewallVipDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteFirewallVip(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceFirewallVipRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadFirewallVip(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip resource from API: %v", err)
	}
	return nil
}
func flattenFirewallVipStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedipMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedportMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipNoNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipTrafficGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtif(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedipMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedportMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipPortforward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallVip(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenFirewallVipStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mappedip_max", flattenFirewallVipMappedipMax(o["mappedip_max"], d, "mappedip_max", sv)); err != nil {
		if !fortiAPIPatch(o["mappedip_max"]) {
			return fmt.Errorf("Error reading mappedip_max: %v", err)
		}
	}

	if err = d.Set("mappedport_min", flattenFirewallVipMappedportMin(o["mappedport_min"], d, "mappedport_min", sv)); err != nil {
		if !fortiAPIPatch(o["mappedport_min"]) {
			return fmt.Errorf("Error reading mappedport_min: %v", err)
		}
	}

	if err = d.Set("no_nat", flattenFirewallVipNoNat(o["no_nat"], d, "no_nat", sv)); err != nil {
		if !fortiAPIPatch(o["no_nat"]) {
			return fmt.Errorf("Error reading no_nat: %v", err)
		}
	}

	if err = d.Set("proto", flattenFirewallVipProto(o["proto"], d, "proto", sv)); err != nil {
		if !fortiAPIPatch(o["proto"]) {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("traffic_group", flattenFirewallVipTrafficGroup(o["traffic-group"], d, "traffic_group", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("extif", flattenFirewallVipExtif(o["extif"], d, "extif", sv)); err != nil {
		if !fortiAPIPatch(o["extif"]) {
			return fmt.Errorf("Error reading extif: %v", err)
		}
	}

	if err = d.Set("mappedip_min", flattenFirewallVipMappedipMin(o["mappedip_min"], d, "mappedip_min", sv)); err != nil {
		if !fortiAPIPatch(o["mappedip_min"]) {
			return fmt.Errorf("Error reading mappedip_min: %v", err)
		}
	}

	if err = d.Set("mappedport_max", flattenFirewallVipMappedportMax(o["mappedport_max"], d, "mappedport_max", sv)); err != nil {
		if !fortiAPIPatch(o["mappedport_max"]) {
			return fmt.Errorf("Error reading mappedport_max: %v", err)
		}
	}

	if err = d.Set("extport", flattenFirewallVipExtport(o["extport"], d, "extport", sv)); err != nil {
		if !fortiAPIPatch(o["extport"]) {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("portforward", flattenFirewallVipPortforward(o["portforward"], d, "portforward", sv)); err != nil {
		if !fortiAPIPatch(o["portforward"]) {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("extip", flattenFirewallVipExtip(o["extip"], d, "extip", sv)); err != nil {
		if !fortiAPIPatch(o["extip"]) {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("mkey", flattenFirewallVipMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandFirewallVipStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedipMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedportMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipNoNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipTrafficGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtif(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedipMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedportMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPortforward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallVipStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("mappedip_max"); ok {
		t, err := expandFirewallVipMappedipMax(d, v, "mappedip_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip_max"] = t
		}
	}

	if v, ok := d.GetOk("mappedport_min"); ok {
		t, err := expandFirewallVipMappedportMin(d, v, "mappedport_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport_min"] = t
		}
	}

	if v, ok := d.GetOk("no_nat"); ok {
		t, err := expandFirewallVipNoNat(d, v, "no_nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no_nat"] = t
		}
	}

	if v, ok := d.GetOk("proto"); ok {
		t, err := expandFirewallVipProto(d, v, "proto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proto"] = t
		}
	}

	if v, ok := d.GetOk("traffic_group"); ok {
		t, err := expandFirewallVipTrafficGroup(d, v, "traffic_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-group"] = t
		}
	}

	if v, ok := d.GetOk("extif"); ok {
		t, err := expandFirewallVipExtif(d, v, "extif", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extif"] = t
		}
	}

	if v, ok := d.GetOk("mappedip_min"); ok {
		t, err := expandFirewallVipMappedipMin(d, v, "mappedip_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip_min"] = t
		}
	}

	if v, ok := d.GetOk("mappedport_max"); ok {
		t, err := expandFirewallVipMappedportMax(d, v, "mappedport_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport_max"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok {
		t, err := expandFirewallVipExtport(d, v, "extport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok {
		t, err := expandFirewallVipPortforward(d, v, "portforward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok {
		t, err := expandFirewallVipExtip(d, v, "extip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandFirewallVipMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
