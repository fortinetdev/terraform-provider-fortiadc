// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global load balance setting

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalLoadBalanceSetting() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalLoadBalanceSettingRead,
		Update: resourceGlobalLoadBalanceSettingUpdate,
		Create: resourceGlobalLoadBalanceSettingUpdate,
		Delete: resourceGlobalLoadBalanceSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ca_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trust_ca_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trust_inter_ca_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipv4_accessed_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipv6_accessed_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"listen_on_interface_all": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"listen_on_interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"retry_num": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"retry_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mask_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mask_length6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"aging_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persist_mask_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persist_mask_length6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persist_aging_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"user_defined_cert": &schema.Schema{
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

func resourceGlobalLoadBalanceSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalLoadBalanceSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceSetting resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_setting?mkey=-1"
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceSetting resource: %v", err)
	}

	d.SetId("GlobalLoadBalanceSetting")

	return resourceGlobalLoadBalanceSettingRead(d, m)
}

func resourceGlobalLoadBalanceSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalLoadBalanceSetting(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalLoadBalanceSetting resource while getting object: %v", err)
	}

	path := "/api/global_load_balance_setting?mkey=-1"
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalLoadBalanceSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalLoadBalanceSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_load_balance_setting"
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalLoadBalanceSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalLoadBalanceSetting resource from API: %v", err)
	}
	return nil
}

func flattenGlobalLoadBalanceSettingAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingCAVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingTrustCAGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingTrustInterCAGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingIpv4AccessedStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingIpv6AccessedStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingListenOnInterfaceAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingListenOnInterfaceList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingRetryNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingRetryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingMaskLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingMaskLength6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingAgingPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingPersistMaskLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingPersistMaskLength6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingPersistAgingPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenGlobalLoadBalanceSettingUserDefinedCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalLoadBalanceSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auth_type", flattenGlobalLoadBalanceSettingAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth-type: %v", err)
		}
	}

	if err = d.Set("password", flattenGlobalLoadBalanceSettingPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("ca_verify", flattenGlobalLoadBalanceSettingCAVerify(o["ca-verify"], d, "ca_verify", sv)); err != nil {
		if !fortiAPIPatch(o["ca-verify"]) {
			return fmt.Errorf("Error reading ca-verify: %v", err)
		}
	}

	if err = d.Set("trust_ca_group", flattenGlobalLoadBalanceSettingTrustCAGroup(o["trust-ca-group"], d, "trust_ca_group", sv)); err != nil {
		if !fortiAPIPatch(o["trust-ca-group"]) {
			return fmt.Errorf("Error reading trust-ca-group: %v", err)
		}
	}

	if err = d.Set("trust_inter_ca_group", flattenGlobalLoadBalanceSettingTrustInterCAGroup(o["trust-inter-ca-group"], d, "trust_inter_ca_group", sv)); err != nil {
		if !fortiAPIPatch(o["trust-inter-ca-group"]) {
			return fmt.Errorf("Error reading trust-inter-ca-group: %v", err)
		}
	}

	if err = d.Set("ipv4_accessed_status", flattenGlobalLoadBalanceSettingIpv4AccessedStatus(o["ipv4-accessed-status"], d, "ipv4_accessed_status", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-accessed-status"]) {
			return fmt.Errorf("Error reading ipv4-accessed-status: %v", err)
		}
	}

	if err = d.Set("ipv6_accessed_status", flattenGlobalLoadBalanceSettingIpv6AccessedStatus(o["ipv6-accessed-status"], d, "ipv6_accessed_status", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-accessed-status"]) {
			return fmt.Errorf("Error reading ipv6-accessed-status: %v", err)
		}
	}

	if err = d.Set("listen_on_interface_all", flattenGlobalLoadBalanceSettingListenOnInterfaceAll(o["listen-on-interface-all"], d, "listen_on_interface_all", sv)); err != nil {
		if !fortiAPIPatch(o["listen-on-interface-all"]) {
			return fmt.Errorf("Error reading listen-on-interface-all: %v", err)
		}
	}

	if err = d.Set("listen_on_interface_list", flattenGlobalLoadBalanceSettingListenOnInterfaceList(o["listen-on-interface-list"], d, "listen_on_interface_list", sv)); err != nil {
		if !fortiAPIPatch(o["listen-on-interface-list"]) {
			return fmt.Errorf("Error reading listen-on-interface-list: %v", err)
		}
	}

	if err = d.Set("port", flattenGlobalLoadBalanceSettingPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("proto", flattenGlobalLoadBalanceSettingProto(o["proto"], d, "proto", sv)); err != nil {
		if !fortiAPIPatch(o["proto"]) {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("retry_num", flattenGlobalLoadBalanceSettingRetryNum(o["retry_num"], d, "retry_num", sv)); err != nil {
		if !fortiAPIPatch(o["retry_num"]) {
			return fmt.Errorf("Error reading retry_num: %v", err)
		}
	}

	if err = d.Set("retry_interval", flattenGlobalLoadBalanceSettingRetryInterval(o["retry_interval"], d, "retry_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retry_interval"]) {
			return fmt.Errorf("Error reading retry_interval: %v", err)
		}
	}

	if err = d.Set("mask_length", flattenGlobalLoadBalanceSettingMaskLength(o["mask_length"], d, "mask_length", sv)); err != nil {
		if !fortiAPIPatch(o["mask_length"]) {
			return fmt.Errorf("Error reading mask_length: %v", err)
		}
	}

	if err = d.Set("mask_length6", flattenGlobalLoadBalanceSettingMaskLength6(o["mask-length6"], d, "mask_length6", sv)); err != nil {
		if !fortiAPIPatch(o["mask-length6"]) {
			return fmt.Errorf("Error reading mask-length6: %v", err)
		}
	}

	if err = d.Set("aging_period", flattenGlobalLoadBalanceSettingAgingPeriod(o["aging_period"], d, "aging_period", sv)); err != nil {
		if !fortiAPIPatch(o["aging_period"]) {
			return fmt.Errorf("Error reading aging_period: %v", err)
		}
	}

	if err = d.Set("persist_mask_length", flattenGlobalLoadBalanceSettingPersistMaskLength(o["persist_mask_length"], d, "persist_mask_length", sv)); err != nil {
		if !fortiAPIPatch(o["persist_mask_length"]) {
			return fmt.Errorf("Error reading persist_mask_length: %v", err)
		}
	}

	if err = d.Set("persist_mask_length6", flattenGlobalLoadBalanceSettingPersistMaskLength6(o["persist_mask-length6"], d, "persist_mask_length6", sv)); err != nil {
		if !fortiAPIPatch(o["persist_mask-length6"]) {
			return fmt.Errorf("Error reading persist_mask-length6: %v", err)
		}
	}

	if err = d.Set("persist_aging_period", flattenGlobalLoadBalanceSettingPersistAgingPeriod(o["persist_aging_period"], d, "persist_aging_period", sv)); err != nil {
		if !fortiAPIPatch(o["persist_aging_period"]) {
			return fmt.Errorf("Error reading persist_aging_period: %v", err)
		}
	}

	if err = d.Set("cert", flattenGlobalLoadBalanceSettingCert(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("user_defined_cert", flattenGlobalLoadBalanceSettingUserDefinedCert(o["user-defined-cert"], d, "user_defined_cert", sv)); err != nil {
		if !fortiAPIPatch(o["user-defined-cert"]) {
			return fmt.Errorf("Error reading user-defined-cert: %v", err)
		}
	}

	return nil
}

func expandGlobalLoadBalanceSettingAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingCAVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingTrustCAGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingTrustInterCAGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingIpv4AccessedStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingIpv6AccessedStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingListenOnInterfaceAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingListenOnInterfaceList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingRetryNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingRetryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingMaskLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingMaskLength6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingAgingPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingPersistMaskLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingPersistMaskLength6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingPersistAgingPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandGlobalLoadBalanceSettingUserDefinedCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalLoadBalanceSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_type"); ok {
		if setArgNil {
			obj["auth-type"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingAuthType(d, v, "auth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("password"); ok {
		if setArgNil {
			obj["password"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingPassword(d, v, "password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["password"] = t
			}
		}
	}

	if v, ok := d.GetOk("ca_verify"); ok {
		if setArgNil {
			obj["ca-verify"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingCAVerify(d, v, "ca_verify", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ca-verify"] = t
			}
		}
	}

	if v, ok := d.GetOk("trust_ca_group"); ok {
		if setArgNil {
			obj["trust-ca-group"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingTrustCAGroup(d, v, "trust_ca_group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trust-ca-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("trust_inter_ca_group"); ok {
		if setArgNil {
			obj["trust-inter-ca-group"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingTrustInterCAGroup(d, v, "trust_inter_ca-group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trust-inter-ca-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv4_accessed_status"); ok {
		if setArgNil {
			obj["ipv4-accessed-status"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingIpv4AccessedStatus(d, v, "ipv4_accessed_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv4-accessed-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_accessed_status"); ok {
		if setArgNil {
			obj["ipv6-accessed-status"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingIpv6AccessedStatus(d, v, "ipv6_accessed_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-accessed-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("listen_on_interface_all"); ok {
		if setArgNil {
			obj["listen-on-interface-all"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingListenOnInterfaceAll(d, v, "listen_on_interface_all", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["listen-on-interface-all"] = t
			}
		}
	}

	if v, ok := d.GetOk("listen_on_interface_list"); ok {
		if setArgNil {
			obj["listen-on-interface-list"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingListenOnInterfaceList(d, v, "listen_on_interface_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["listen-on-interface-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	if v, ok := d.GetOk("proto"); ok {
		if setArgNil {
			obj["proto"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingProto(d, v, "proto", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proto"] = t
			}
		}
	}

	if v, ok := d.GetOk("retry_num"); ok {
		if setArgNil {
			obj["retry_num"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingRetryNum(d, v, "retry_num", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["retry_num"] = t
			}
		}
	}

	if v, ok := d.GetOk("retry_interval"); ok {
		if setArgNil {
			obj["retry_interval"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingRetryInterval(d, v, "retry_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["retry_interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("mask_length"); ok {
		if setArgNil {
			obj["mask_length"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingMaskLength(d, v, "mask_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mask_length"] = t
			}
		}
	}

	if v, ok := d.GetOk("mask_length6"); ok {
		if setArgNil {
			obj["mask-length6"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingMaskLength6(d, v, "mask_length6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mask-length6"] = t
			}
		}
	}

	if v, ok := d.GetOk("aging_period"); ok {
		if setArgNil {
			obj["aging_period"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingAgingPeriod(d, v, "aging_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["aging_period"] = t
			}
		}
	}

	if v, ok := d.GetOk("persist_mask_length"); ok {
		if setArgNil {
			obj["persist_mask_length"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingPersistMaskLength(d, v, "persist_mask_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["persist_mask_length"] = t
			}
		}
	}

	if v, ok := d.GetOk("persist_mask_length6"); ok {
		if setArgNil {
			obj["persist_mask-length6"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingPersistMaskLength6(d, v, "persist_mask_length6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["persist_mask-length6"] = t
			}
		}
	}

	if v, ok := d.GetOk("persist_aging_period"); ok {
		if setArgNil {
			obj["persist_aging_period"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingPersistAgingPeriod(d, v, "persist_aging_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["persist_aging_period"] = t
			}
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		if setArgNil {
			obj["cert"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingCert(d, v, "cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("user_defined_cert"); ok {
		if setArgNil {
			obj["user-defined-cert"] = nil
		} else {
			t, err := expandGlobalLoadBalanceSettingUserDefinedCert(d, v, "user_defined_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["user-defined-cert"] = t
			}
		}
	}

	return &obj, nil
}
