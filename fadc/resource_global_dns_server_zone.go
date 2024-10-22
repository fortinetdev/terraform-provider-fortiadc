// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server zone

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerZone() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerZoneRead,
		Update: resourceGlobalDnsServerZoneUpdate,
		Create: resourceGlobalDnsServerZoneCreate,
		Delete: resourceGlobalDnsServerZoneDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dnssec_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dnssec_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dnssec_keysize": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dsset_info_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"negative_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"responsible_mail": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"primary_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"primary_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"primary_server_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward_host_enable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forwarders": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"notify_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"also_notify_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_transfer": &schema.Schema{
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

func resourceGlobalDnsServerZoneCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing GlobalDnsServerZone: type error")
	}

	obj, err := getObjectGlobalDnsServerZone(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerZone resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_zone"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating GlobalDnsServerZone resource: %v", err)
	}

	d.SetId(mkey)

	return resourceGlobalDnsServerZoneRead(d, m)
}

func resourceGlobalDnsServerZoneUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalDnsServerZone(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerZone resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_zone_update?mkey=" + escapeURLString(mkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerZone resource: %v", err)
	}

	return resourceGlobalDnsServerZoneRead(d, m)
}

func resourceGlobalDnsServerZoneDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_zone?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerZone resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGlobalDnsServerZoneRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_zone?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerZone resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerZone(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerZone resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerZone(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenGlobalDnsServerZone(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("type", flattenGlobalDnsServerZone(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenGlobalDnsServerZone(o["domain_name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain_name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("dns_policy", flattenGlobalDnsServerZone(o["dns_policy"], d, "dns_policy", sv)); err != nil {
		if !fortiAPIPatch(o["dns_policy"]) {
			return fmt.Errorf("Error reading dns_policy: %v", err)
		}
	}

	if err = d.Set("dnssec_status", flattenGlobalDnsServerZone(o["dnssec_status"], d, "dnssec_status", sv)); err != nil {
		if !fortiAPIPatch(o["dnssec_status"]) {
			return fmt.Errorf("Error reading dnssec_status: %v", err)
		}
	}

	if err = d.Set("dnssec_algorithm", flattenGlobalDnsServerZone(o["dnssec_algorithm"], d, "dnssec_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["dnssec_algorithm"]) {
			return fmt.Errorf("Error reading dnssec_algorithm: %v", err)
		}
	}

	if err = d.Set("dnssec_keysize", flattenGlobalDnsServerZone(o["dnssec_keysize"], d, "dnssec_keysize", sv)); err != nil {
		if !fortiAPIPatch(o["dnssec_keysize"]) {
			return fmt.Errorf("Error reading dnssec_keysize: %v", err)
		}
	}

	if err = d.Set("dsset_info_list", flattenGlobalDnsServerZone(o["dsset_info_list"], d, "dsset_info_list", sv)); err != nil {
		if !fortiAPIPatch(o["dsset_info_list"]) {
			return fmt.Errorf("Error reading dsset_info_list: %v", err)
		}
	}

	if err = d.Set("ttl", flattenGlobalDnsServerZone(o["ttl"], d, "ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ttl"]) {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("serial", flattenGlobalDnsServerZone(o["serial"], d, "serial", sv)); err != nil {
		if !fortiAPIPatch(o["serial"]) {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("negative_ttl", flattenGlobalDnsServerZone(o["negative_ttl"], d, "negative_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["negative_ttl"]) {
			return fmt.Errorf("Error reading negative_ttl: %v", err)
		}
	}

	if err = d.Set("responsible_mail", flattenGlobalDnsServerZone(o["responible_mail"], d, "responsible_mail", sv)); err != nil {
		if !fortiAPIPatch(o["responible_mail"]) {
			return fmt.Errorf("Error reading responsible_mail: %v", err)
		}
	}

	if err = d.Set("primary_server_name", flattenGlobalDnsServerZone(o["primary_server_name"], d, "primary_server_name", sv)); err != nil {
		if !fortiAPIPatch(o["primary_server_name"]) {
			return fmt.Errorf("Error reading primary_server_name: %v", err)
		}
	}

	if err = d.Set("primary_server_ip", flattenGlobalDnsServerZone(o["primary_server_ip"], d, "primary_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["primary_server_ip"]) {
			return fmt.Errorf("Error reading primary_server_ip: %v", err)
		}
	}

	if err = d.Set("primary_server_ip6", flattenGlobalDnsServerZone(o["primary_server_ip6"], d, "primary_server_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["primary_server_ip6"]) {
			return fmt.Errorf("Error reading primary_server_ip6: %v", err)
		}
	}

	if err = d.Set("forward_host_enable", flattenGlobalDnsServerZone(o["forward_host_enable"], d, "forward_host_enable", sv)); err != nil {
		if !fortiAPIPatch(o["forward_host_enable"]) {
			return fmt.Errorf("Error reading forward_host_enable: %v", err)
		}
	}

	if err = d.Set("forward", flattenGlobalDnsServerZone(o["forward"], d, "forward", sv)); err != nil {
		if !fortiAPIPatch(o["forward"]) {
			return fmt.Errorf("Error reading forward: %v", err)
		}
	}

	if err = d.Set("forwarders", flattenGlobalDnsServerZone(o["forwarders"], d, "forwarders", sv)); err != nil {
		if !fortiAPIPatch(o["forwarders"]) {
			return fmt.Errorf("Error reading forwarders: %v", err)
		}
	}

	if err = d.Set("notify_status", flattenGlobalDnsServerZone(o["notify_status"], d, "notify_status", sv)); err != nil {
		if !fortiAPIPatch(o["notify_status"]) {
			return fmt.Errorf("Error reading notify_status: %v", err)
		}
	}

	if err = d.Set("also_notify_list", flattenGlobalDnsServerZone(o["also_notify_list"], d, "also_notify_list", sv)); err != nil {
		if !fortiAPIPatch(o["also_notify_list"]) {
			return fmt.Errorf("Error reading also_notify_list: %v", err)
		}
	}

	if err = d.Set("allow_transfer", flattenGlobalDnsServerZone(o["allow_transfer"], d, "allow_transfer", sv)); err != nil {
		if !fortiAPIPatch(o["allow_transfer"]) {
			return fmt.Errorf("Error reading allow_transfer: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerZone(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain_name"] = t
		}
	}

	if v, ok := d.GetOk("dns_policy"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "dns_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns_policy"] = t
		}
	}

	if v, ok := d.GetOk("dnssec_status"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "dnssec_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssec_status"] = t
		}
	}

	if v, ok := d.GetOk("dnssec_algorithm"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "dnssec_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssec_algorithm"] = t
		}
	}

	if v, ok := d.GetOk("dnssec_keysize"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "dnssec_keysize", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssec_keysize"] = t
		}
	}

	if v, ok := d.GetOk("dsset_info_list"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "dsset_info_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsset_info_list"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "serial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("negative_ttl"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "negative_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negative_ttl"] = t
		}
	}

	if v, ok := d.GetOk("responsible_mail"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "responsible_mail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["responible_mail"] = t
		}
	}

	if v, ok := d.GetOk("primary_server_name"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "primary_server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_server_name"] = t
		}
	}

	if v, ok := d.GetOk("primary_server_ip"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "primary_server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_server_ip"] = t
		}
	}

	if v, ok := d.GetOk("forward_host_enable"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "forward_host_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward_host_enable"] = t
		}
	}

	if v, ok := d.GetOk("forward"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward"] = t
		}
	}

	if v, ok := d.GetOk("forwarders"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "forwarders", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarders"] = t
		}
	}

	if v, ok := d.GetOk("notify_status"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "notify_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify_status"] = t
		}
	}

	if v, ok := d.GetOk("also_notify_list"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "also_notify_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["also_notify_list"] = t
		}
	}

	if v, ok := d.GetOk("allow_transfer"); ok {
		t, err := expandGlobalDnsServerZone(d, v, "allow_transfer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow_transfer"] = t
		}
	}

	return &obj, nil
}
