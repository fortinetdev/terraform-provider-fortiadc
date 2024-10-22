// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure global dns server general.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalDnsServerGeneral() *schema.Resource {
	return &schema.Resource{
		Read:   resourceGlobalDnsServerGeneralRead,
		Update: resourceGlobalDnsServerGeneralUpdate,
		Create: resourceGlobalDnsServerGeneralUpdate,
		Delete: resourceGlobalDnsServerGeneralDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"gds_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"recurision_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dnssec_validate_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"minimal_responses": &schema.Schema{
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
			"listen_on_all_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dohs_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dohs_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dohs_interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"doh_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"doh_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"doh_interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dot_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dot_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dot_interface_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_system_dns_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"forwarders": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traffic_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"response_ratelimit": &schema.Schema{
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
func resourceGlobalDnsServerGeneralUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalDnsServerGeneral(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerGeneral resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_general?mkey=-1"
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerGeneral resource: %v", err)
	}

	d.SetId("GlobalDnsServerGeneral")
	return resourceFirewallPolicyRead(d, m)
}
func resourceGlobalDnsServerGeneralDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectGlobalDnsServerGeneral(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GlobalDnsServerGeneral resource while getting object: %v", err)
	}

	path := "/api/global_dns_server_general?mkey=-1"
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing GlobalDnsServerGeneral resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceGlobalDnsServerGeneralRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/global_dns_server_general"
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerGeneral resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGlobalDnsServerGeneral(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GlobalDnsServerGeneral resource from API: %v", err)
	}
	return nil
}

func flattenGlobalDnsServerGeneral(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGlobalDnsServerGeneral(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("gds_status", flattenGlobalDnsServerGeneral(o["gds_status"], d, "gds_status", sv)); err != nil {
		if !fortiAPIPatch(o["gds_status"]) {
			return fmt.Errorf("Error reading gds_status: %v", err)
		}
	}

	if err = d.Set("recurision_status", flattenGlobalDnsServerGeneral(o["recurision_status"], d, "recurision_status", sv)); err != nil {
		if !fortiAPIPatch(o["recurision_status"]) {
			return fmt.Errorf("Error reading recurision_status: %v", err)
		}
	}

	if err = d.Set("dnssec_validate_status", flattenGlobalDnsServerGeneral(o["dnssec_validate_status"], d, "dnssec_validate_status", sv)); err != nil {
		if !fortiAPIPatch(o["dnssec_validate_status"]) {
			return fmt.Errorf("Error reading dnssec_validate_status: %v", err)
		}
	}

	if err = d.Set("minimal_responses", flattenGlobalDnsServerGeneral(o["minimal_responses"], d, "minimal_responses", sv)); err != nil {
		if !fortiAPIPatch(o["minimal_responses"]) {
			return fmt.Errorf("Error reading minimal_responses: %v", err)
		}
	}

	if err = d.Set("ipv4_accessed_status", flattenGlobalDnsServerGeneral(o["ipv4-accessed-status"], d, "ipv4_accessed_status", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-accessed-status"]) {
			return fmt.Errorf("Error reading ipv4_accessed_status: %v", err)
		}
	}

	if err = d.Set("ipv6_accessed_status", flattenGlobalDnsServerGeneral(o["ipv6_accessed_status"], d, "ipv6_accessed_status", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6_accessed_status"]) {
			return fmt.Errorf("Error reading ipv6_accessed_status: %v", err)
		}
	}

	if err = d.Set("listen_on_all_interface", flattenGlobalDnsServerGeneral(o["listen-on_all_interface"], d, "listen_on_all_interface", sv)); err != nil {
		if !fortiAPIPatch(o["listen-on_all_interface"]) {
			return fmt.Errorf("Error reading listen_on_all_interface: %v", err)
		}
	}

	if err = d.Set("interface_list", flattenGlobalDnsServerGeneral(o["interface_list"], d, "interface_list", sv)); err != nil {
		if !fortiAPIPatch(o["interface_list"]) {
			return fmt.Errorf("Error reading interface_list: %v", err)
		}
	}

	if err = d.Set("dohs_status", flattenGlobalDnsServerGeneral(o["dohs_status"], d, "dohs_status", sv)); err != nil {
		if !fortiAPIPatch(o["dohs_status"]) {
			return fmt.Errorf("Error reading dohs_status: %v", err)
		}
	}

	if err = d.Set("dohs_port", flattenGlobalDnsServerGeneral(o["dohs_port"], d, "dohs_port", sv)); err != nil {
		if !fortiAPIPatch(o["dohs_port"]) {
			return fmt.Errorf("Error reading dohs_port: %v", err)
		}
	}

	if err = d.Set("dohs_interface_list", flattenGlobalDnsServerGeneral(o["dohs_interface_list"], d, "dohs_interface_list", sv)); err != nil {
		if !fortiAPIPatch(o["dohs_interface_list"]) {
			return fmt.Errorf("Error reading dohs_interface_list: %v", err)
		}
	}

	if err = d.Set("doh_status", flattenGlobalDnsServerGeneral(o["doh_status"], d, "doh_status", sv)); err != nil {
		if !fortiAPIPatch(o["doh_status"]) {
			return fmt.Errorf("Error reading doh_status: %v", err)
		}
	}

	if err = d.Set("doh_port", flattenGlobalDnsServerGeneral(o["doh_port"], d, "doh_port", sv)); err != nil {
		if !fortiAPIPatch(o["doh_port"]) {
			return fmt.Errorf("Error reading doh_port: %v", err)
		}
	}

	if err = d.Set("doh_interface_list", flattenGlobalDnsServerGeneral(o["doh_interface_list"], d, "doh_interface_list", sv)); err != nil {
		if !fortiAPIPatch(o["doh_interface_list"]) {
			return fmt.Errorf("Error reading doh_interface_list: %v", err)
		}
	}

	if err = d.Set("dot_status", flattenGlobalDnsServerGeneral(o["dot_status"], d, "dot_status", sv)); err != nil {
		if !fortiAPIPatch(o["dot_status"]) {
			return fmt.Errorf("Error reading dot_status: %v", err)
		}
	}

	if err = d.Set("dot_port", flattenGlobalDnsServerGeneral(o["dot_port"], d, "dot_port", sv)); err != nil {
		if !fortiAPIPatch(o["dot_port"]) {
			return fmt.Errorf("Error reading dot_port: %v", err)
		}
	}

	if err = d.Set("dot_interface_list", flattenGlobalDnsServerGeneral(o["dot_interface_list"], d, "dot_interface_list", sv)); err != nil {
		if !fortiAPIPatch(o["dot_interface_list"]) {
			return fmt.Errorf("Error reading dot_interface_list: %v", err)
		}
	}

	if err = d.Set("certificate", flattenGlobalDnsServerGeneral(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("forward", flattenGlobalDnsServerGeneral(o["forward"], d, "forward", sv)); err != nil {
		if !fortiAPIPatch(o["forward"]) {
			return fmt.Errorf("Error reading forward: %v", err)
		}
	}

	if err = d.Set("use_system_dns_server", flattenGlobalDnsServerGeneral(o["use-system_dns_server"], d, "use_system_dns_server", sv)); err != nil {
		if !fortiAPIPatch(o["use-system_dns_server"]) {
			return fmt.Errorf("Error reading use_system_dns_server: %v", err)
		}
	}

	if err = d.Set("forwarders", flattenGlobalDnsServerGeneral(o["forwarders"], d, "forwarders", sv)); err != nil {
		if !fortiAPIPatch(o["forwarders"]) {
			return fmt.Errorf("Error reading forwarders: %v", err)
		}
	}

	if err = d.Set("traffic_log", flattenGlobalDnsServerGeneral(o["traffic-log"], d, "traffic_log", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-log"]) {
			return fmt.Errorf("Error reading traffic_log: %v", err)
		}
	}

	if err = d.Set("response_ratelimit", flattenGlobalDnsServerGeneral(o["response_ratelimit"], d, "response_ratelimit", sv)); err != nil {
		if !fortiAPIPatch(o["response_ratelimit"]) {
			return fmt.Errorf("Error reading response_ratelimit: %v", err)
		}
	}

	return nil
}

func expandGlobalDnsServerGeneral(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGlobalDnsServerGeneral(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("gds_status"); ok {
		if setArgNil {
			obj["gds_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "gds_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gds_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("recurision_status"); ok {
		if setArgNil {
			obj["recurision_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "recurision_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["recurision_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("dnssec_validate_status"); ok {
		if setArgNil {
			obj["dnssec_validate_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dnssec_validate_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dnssec_validate_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("minimal_responses"); ok {
		if setArgNil {
			obj["minimal_responses"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "minimal_responses", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["minimal_responses"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv4_accessed_status"); ok {
		if setArgNil {
			obj["ipv4-accessed-status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "ipv4_accessed_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv4-accessed-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_accessed_status"); ok {
		if setArgNil {
			obj["ipv6_accessed_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "ipv6_accessed_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6_accessed_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("listen_on_all_interface"); ok {
		if setArgNil {
			obj["listen-on_all_interface"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "listen_on_all_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["listen-on_all_interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_list"); ok {
		if setArgNil {
			obj["interface_list"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "interface_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface_list"] = t
			}
		}
	}

	if v, ok := d.GetOk("dohs_status"); ok {
		if setArgNil {
			obj["dohs_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dohs_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dohs_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("dohs_port"); ok {
		if setArgNil {
			obj["dohs_port"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dohs_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dohs_port"] = t
			}
		}
	}

	if v, ok := d.GetOk("dohs_interface_list"); ok {
		if setArgNil {
			obj["dohs_interface_list"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dohs_interface_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dohs_interface_list"] = t
			}
		}
	}

	if v, ok := d.GetOk("doh_status"); ok {
		if setArgNil {
			obj["doh_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "doh_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["doh_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("doh_port"); ok {
		if setArgNil {
			obj["doh_port"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "doh_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["doh_port"] = t
			}
		}
	}

	if v, ok := d.GetOk("doh_interface_list"); ok {
		if setArgNil {
			obj["doh_interface_list"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "doh_interface_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["doh_interface_list"] = t
			}
		}
	}

	if v, ok := d.GetOk("dot_status"); ok {
		if setArgNil {
			obj["dot_status"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dot_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dot_status"] = t
			}
		}
	}

	if v, ok := d.GetOk("dot_port"); ok {
		if setArgNil {
			obj["dot_port"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dot_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dot_port"] = t
			}
		}
	}

	if v, ok := d.GetOk("dot_interface_list"); ok {
		if setArgNil {
			obj["dot_interface_list"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "dot_interface_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dot_interface_list"] = t
			}
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		if setArgNil {
			obj["certificate"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("forward"); ok {
		if setArgNil {
			obj["forward"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "forward", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward"] = t
			}
		}
	}

	if v, ok := d.GetOk("use_system_dns_server"); ok {
		if setArgNil {
			obj["use-system_dns_server"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "use_system_dns_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["use-system_dns_server"] = t
			}
		}
	}

	if v, ok := d.GetOk("forwarders"); ok {
		if setArgNil {
			obj["forwarders"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "forwarders", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forwarders"] = t
			}
		}
	}

	if v, ok := d.GetOk("traffic_log"); ok {
		if setArgNil {
			obj["traffic-log"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "traffic_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["traffic-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("response_ratelimit"); ok {
		if setArgNil {
			obj["response_ratelimit"] = nil
		} else {
			t, err := expandGlobalDnsServerGeneral(d, v, "response_ratelimit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["response_ratelimit"] = t
			}
		}
	}

	return &obj, nil
}
