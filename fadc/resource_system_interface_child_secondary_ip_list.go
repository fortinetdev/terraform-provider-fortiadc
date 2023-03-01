// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface child secondary ip list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemInterfaceChildSecondaryIpList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemInterfaceChildSecondaryIpListRead,
		Update: resourceSystemInterfaceChildSecondaryIpListUpdate,
		Create: resourceSystemInterfaceChildSecondaryIpListCreate,
		Delete: resourceSystemInterfaceChildSecondaryIpListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"traffic_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"floating_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"floating": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fadc_id": &schema.Schema{
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
func resourceSystemInterfaceChildSecondaryIpListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("fadc_id")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	obj, err := getObjectSystemInterfaceChildSecondaryIpList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceChildSecondaryIpList resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemInterfaceChildSecondaryIpList(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceChildSecondaryIpList resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemInterfaceChildSecondaryIpListRead(d, m)
}
func resourceSystemInterfaceChildSecondaryIpListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemInterfaceChildSecondaryIpList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceChildSecondaryIpList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceChildSecondaryIpList(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceChildSecondaryIpList resource: %v", err)
	}

	return resourceSystemInterfaceChildSecondaryIpListRead(d, m)
}
func resourceSystemInterfaceChildSecondaryIpListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemInterfaceChildSecondaryIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceChildSecondaryIpList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemInterfaceChildSecondaryIpListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildSecondaryIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemInterfaceChildSecondaryIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceChildSecondaryIpList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceChildSecondaryIpList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceChildSecondaryIpList resource from API: %v", err)
	}
	return nil
}
func flattenSystemInterfaceChildSecondaryIpListTrafficGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildSecondaryIpListIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceChildSecondaryIpListFloatingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildSecondaryIpListAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildSecondaryIpListFloating(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildSecondaryIpListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemInterfaceChildSecondaryIpList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("traffic_group", flattenSystemInterfaceChildSecondaryIpListTrafficGroup(o["traffic-group"], d, "traffic_group", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-group"]) {
			return fmt.Errorf("Error reading traffic-group: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceChildSecondaryIpListIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("floating_ip", flattenSystemInterfaceChildSecondaryIpListFloatingIp(o["floating-ip"], d, "floating_ip", sv)); err != nil {
		if !fortiAPIPatch(o["floating-ip"]) {
			return fmt.Errorf("Error reading floating-ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystemInterfaceChildSecondaryIpListAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("floating", flattenSystemInterfaceChildSecondaryIpListFloating(o["floating"], d, "floating", sv)); err != nil {
		if !fortiAPIPatch(o["floating"]) {
			return fmt.Errorf("Error reading floating: %v", err)
		}
	}

	if err = d.Set("fadc_id", flattenSystemInterfaceChildSecondaryIpListId(o["id"], d, "fadc_id", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}

	return nil
}

func expandSystemInterfaceChildSecondaryIpListTrafficGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildSecondaryIpListIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildSecondaryIpListFloatingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildSecondaryIpListAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildSecondaryIpListFloating(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildSecondaryIpListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceChildSecondaryIpList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("traffic_group"); ok {
		t, err := expandSystemInterfaceChildSecondaryIpListTrafficGroup(d, v, "traffic_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-group"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemInterfaceChildSecondaryIpListIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("floating_ip"); ok {
		t, err := expandSystemInterfaceChildSecondaryIpListFloatingIp(d, v, "floating_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["floating-ip"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandSystemInterfaceChildSecondaryIpListAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("floating"); ok {
		t, err := expandSystemInterfaceChildSecondaryIpListFloating(d, v, "floating", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["floating"] = t
		}
	}

	if v, ok := d.GetOk("fadc_id"); ok {
		t, err := expandSystemInterfaceChildSecondaryIpListId(d, v, "fadc_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
