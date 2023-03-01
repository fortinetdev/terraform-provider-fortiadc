// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface child trust ip list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemInterfaceChildTrustIpList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemInterfaceChildTrustIpListRead,
		Update: resourceSystemInterfaceChildTrustIpListUpdate,
		Create: resourceSystemInterfaceChildTrustIpListCreate,
		Delete: resourceSystemInterfaceChildTrustIpListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"ip_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip6_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
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
func resourceSystemInterfaceChildTrustIpListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	obj, err := getObjectSystemInterfaceChildTrustIpList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceChildTrustIpList resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemInterfaceChildTrustIpList(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceChildTrustIpList resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemInterfaceChildTrustIpListRead(d, m)
}
func resourceSystemInterfaceChildTrustIpListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemInterfaceChildTrustIpList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceChildTrustIpList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceChildTrustIpList(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceChildTrustIpList resource: %v", err)
	}

	return resourceSystemInterfaceChildTrustIpListRead(d, m)
}
func resourceSystemInterfaceChildTrustIpListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemInterfaceChildTrustIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceChildTrustIpList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemInterfaceChildTrustIpListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildTrustIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemInterfaceChildTrustIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceChildTrustIpList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceChildTrustIpList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceChildTrustIpList resource from API: %v", err)
	}
	return nil
}
func flattenSystemInterfaceChildTrustIpListIp6Max(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildTrustIpListIp6Min(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildTrustIpListIpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildTrustIpListIpNetmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceChildTrustIpListIp6Netmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceChildTrustIpListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildTrustIpListIpMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildTrustIpListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemInterfaceChildTrustIpList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip6_max", flattenSystemInterfaceChildTrustIpListIp6Max(o["ip6-max"], d, "ip6_max", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-max"]) {
			return fmt.Errorf("Error reading ip6-max: %v", err)
		}
	}

	if err = d.Set("ip6_min", flattenSystemInterfaceChildTrustIpListIp6Min(o["ip6-min"], d, "ip6_min", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-min"]) {
			return fmt.Errorf("Error reading ip6-min: %v", err)
		}
	}

	if err = d.Set("ip_max", flattenSystemInterfaceChildTrustIpListIpMax(o["ip-max"], d, "ip_max", sv)); err != nil {
		if !fortiAPIPatch(o["ip-max"]) {
			return fmt.Errorf("Error reading ip-max: %v", err)
		}
	}

	if err = d.Set("ip_netmask", flattenSystemInterfaceChildTrustIpListIpNetmask(o["ip-netmask"], d, "ip_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ip-netmask"]) {
			return fmt.Errorf("Error reading ip-netmask: %v", err)
		}
	}

	if err = d.Set("ip6_netmask", flattenSystemInterfaceChildTrustIpListIp6Netmask(o["ip6-netmask"], d, "ip6_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-netmask"]) {
			return fmt.Errorf("Error reading ip6-netmask: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemInterfaceChildTrustIpListType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ip_min", flattenSystemInterfaceChildTrustIpListIpMin(o["ip-min"], d, "ip_min", sv)); err != nil {
		if !fortiAPIPatch(o["ip-min"]) {
			return fmt.Errorf("Error reading ip-min: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemInterfaceChildTrustIpListName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func expandSystemInterfaceChildTrustIpListIp6Max(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListIp6Min(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListIpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListIpNetmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListIp6Netmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListIpMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildTrustIpListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceChildTrustIpList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip6_max"); ok {
		t, err := expandSystemInterfaceChildTrustIpListIp6Max(d, v, "ip6_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-max"] = t
		}
	}

	if v, ok := d.GetOk("ip6_min"); ok {
		t, err := expandSystemInterfaceChildTrustIpListIp6Min(d, v, "ip6_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-min"] = t
		}
	}

	if v, ok := d.GetOk("ip_max"); ok {
		t, err := expandSystemInterfaceChildTrustIpListIpMax(d, v, "ip_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-max"] = t
		}
	}

	if v, ok := d.GetOk("ip_netmask"); ok {
		t, err := expandSystemInterfaceChildTrustIpListIpNetmask(d, v, "ip_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-netmask"] = t
		}
	}

	if v, ok := d.GetOk("ip6_netmask"); ok {
		t, err := expandSystemInterfaceChildTrustIpListIp6Netmask(d, v, "ip6_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-netmask"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemInterfaceChildTrustIpListType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ip_min"); ok {
		t, err := expandSystemInterfaceChildTrustIpListIpMin(d, v, "ip_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-min"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemInterfaceChildTrustIpListName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
