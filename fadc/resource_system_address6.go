// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system address6.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAddress6() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAddress6Read,
		Update: resourceSystemAddress6Update,
		Create: resourceSystemAddress6Create,
		Delete: resourceSystemAddress6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip6_network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_min": &schema.Schema{
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
		},
	}
}
func resourceSystemAddress6Create(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAddress6: type error")
	}

	obj, err := getObjectSystemAddress6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAddress6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAddress6(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAddress6 resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAddress6Read(d, m)
}
func resourceSystemAddress6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAddress6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAddress6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAddress6(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAddress6 resource: %v", err)
	}

	return resourceSystemAddress6Read(d, m)
}
func resourceSystemAddress6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemAddress6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemAddress6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAddress6(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAddress6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAddress6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAddress6 resource from API: %v", err)
	}
	return nil
}
func flattenSystemAddress6Ip6Network(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAddress6Ip6Max(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAddress6Ip6Min(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAddress6Type(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAddress6Mkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAddress6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip6_network", flattenSystemAddress6Ip6Network(o["ip6-network"], d, "ip6_network", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-network"]) {
			return fmt.Errorf("Error reading ip6-network: %v", err)
		}
	}

	if err = d.Set("ip6_max", flattenSystemAddress6Ip6Max(o["ip6-max"], d, "ip6_max", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-max"]) {
			return fmt.Errorf("Error reading ip6-max: %v", err)
		}
	}

	if err = d.Set("ip6_min", flattenSystemAddress6Ip6Min(o["ip6-min"], d, "ip6_min", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-min"]) {
			return fmt.Errorf("Error reading ip6-min: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemAddress6Type(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemAddress6Mkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemAddress6Ip6Network(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAddress6Ip6Max(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAddress6Ip6Min(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAddress6Type(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAddress6Mkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAddress6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip6_network"); ok {
		t, err := expandSystemAddress6Ip6Network(d, v, "ip6_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-network"] = t
		}
	}

	if v, ok := d.GetOk("ip6_max"); ok {
		t, err := expandSystemAddress6Ip6Max(d, v, "ip6_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-max"] = t
		}
	}

	if v, ok := d.GetOk("ip6_min"); ok {
		t, err := expandSystemAddress6Ip6Min(d, v, "ip6_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-min"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemAddress6Type(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemAddress6Mkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
