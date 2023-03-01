// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system interface child ha node ip list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemInterfaceChildHaNodeIpList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemInterfaceChildHaNodeIpListRead,
		Update: resourceSystemInterfaceChildHaNodeIpListUpdate,
		Create: resourceSystemInterfaceChildHaNodeIpListCreate,
		Delete: resourceSystemInterfaceChildHaNodeIpListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
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
func resourceSystemInterfaceChildHaNodeIpListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	obj, err := getObjectSystemInterfaceChildHaNodeIpList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceChildHaNodeIpList resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemInterfaceChildHaNodeIpList(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceChildHaNodeIpList resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemInterfaceChildHaNodeIpListRead(d, m)
}
func resourceSystemInterfaceChildHaNodeIpListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemInterfaceChildHaNodeIpList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceChildHaNodeIpList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceChildHaNodeIpList(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceChildHaNodeIpList resource: %v", err)
	}

	return resourceSystemInterfaceChildHaNodeIpListRead(d, m)
}
func resourceSystemInterfaceChildHaNodeIpListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemInterfaceChildHaNodeIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceChildHaNodeIpList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemInterfaceChildHaNodeIpListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemInterfaceChildHaNodeIpList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceChildHaNodeIpList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceChildHaNodeIpList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceChildHaNodeIpList resource from API: %v", err)
	}
	return nil
}
func flattenSystemInterfaceChildHaNodeIpListNode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildHaNodeIpListIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceChildHaNodeIpListAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemInterfaceChildHaNodeIpListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemInterfaceChildHaNodeIpList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("node", flattenSystemInterfaceChildHaNodeIpListNode(o["node"], d, "node", sv)); err != nil {
		if !fortiAPIPatch(o["node"]) {
			return fmt.Errorf("Error reading node: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceChildHaNodeIpListIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystemInterfaceChildHaNodeIpListAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("fadc_id", flattenSystemInterfaceChildHaNodeIpListId(o["id"], d, "fadc_id", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}

	return nil
}

func expandSystemInterfaceChildHaNodeIpListNode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildHaNodeIpListIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildHaNodeIpListAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceChildHaNodeIpListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceChildHaNodeIpList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("node"); ok {
		t, err := expandSystemInterfaceChildHaNodeIpListNode(d, v, "node", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemInterfaceChildHaNodeIpListIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandSystemInterfaceChildHaNodeIpListAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("fadc_id"); ok {
		t, err := expandSystemInterfaceChildHaNodeIpListId(d, v, "fadc_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
