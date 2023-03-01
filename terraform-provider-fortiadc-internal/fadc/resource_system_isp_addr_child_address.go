// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system isp addr child address.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIspAddrChildAddress() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemIspAddrChildAddressRead,
		Update: resourceSystemIspAddrChildAddressUpdate,
		Create: resourceSystemIspAddrChildAddressCreate,
		Delete: resourceSystemIspAddrChildAddressDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"province": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_netmask": &schema.Schema{
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
func resourceSystemIspAddrChildAddressCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	obj, err := getObjectSystemIspAddrChildAddress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemIspAddrChildAddress resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemIspAddrChildAddress(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemIspAddrChildAddress resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemIspAddrChildAddressRead(d, m)
}
func resourceSystemIspAddrChildAddressUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemIspAddrChildAddress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIspAddrChildAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIspAddrChildAddress(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemIspAddrChildAddress resource: %v", err)
	}

	return resourceSystemIspAddrChildAddressRead(d, m)
}
func resourceSystemIspAddrChildAddressDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemIspAddrChildAddress(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIspAddrChildAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemIspAddrChildAddressRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemIspAddrChildAddress: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemIspAddrChildAddress(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemIspAddrChildAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIspAddrChildAddress(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIspAddrChildAddress resource from API: %v", err)
	}
	return nil
}
func flattenSystemIspAddrChildAddressProvince(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIspAddrChildAddressIpNetmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemIspAddrChildAddressMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIspAddrChildAddress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("province", flattenSystemIspAddrChildAddressProvince(o["province"], d, "province", sv)); err != nil {
		if !fortiAPIPatch(o["province"]) {
			return fmt.Errorf("Error reading province: %v", err)
		}
	}

	if err = d.Set("ip_netmask", flattenSystemIspAddrChildAddressIpNetmask(o["ip-netmask"], d, "ip_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ip-netmask"]) {
			return fmt.Errorf("Error reading ip-netmask: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemIspAddrChildAddressMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemIspAddrChildAddressProvince(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIspAddrChildAddressIpNetmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIspAddrChildAddressMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIspAddrChildAddress(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("province"); ok {
		t, err := expandSystemIspAddrChildAddressProvince(d, v, "province", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["province"] = t
		}
	}

	if v, ok := d.GetOk("ip_netmask"); ok {
		t, err := expandSystemIspAddrChildAddressIpNetmask(d, v, "ip_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-netmask"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemIspAddrChildAddressMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
