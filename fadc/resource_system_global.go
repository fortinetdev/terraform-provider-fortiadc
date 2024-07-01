// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system global.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemGlobalRead,
		Update: resourceSystemGlobalUpdate,
		Create: resourceSystemGlobalUpdate,
		Delete: resourceSystemGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdom_admin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"telnet_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"https_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"config_sync_enable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sys_global_language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"config_sync_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_device_latitude": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_router": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_intermediate_ca_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"admin_idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_system": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_default_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_device_longtitude": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_second": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobal(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource: %v", err)
	}

	d.SetId("SystemGlobal")
	return resourceSystemGlobalRead(d, m)
}
func resourceSystemGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemGlobal(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobal(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemGlobal(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource from API: %v", err)
	}
	return nil
}
func flattenSystemGlobalVdomAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTelnetPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHttpsServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalConfigSyncEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVdomMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSysGlobalLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalConfigSyncPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDeviceLatitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiRouter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDefaultIntermediateCaGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHttpsRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiSystem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHttpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUseDefaultHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDeviceLongtitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPrimaryDNS(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSecondaryDNS(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vdom_admin", flattenSystemGlobalVdomAdmin(o["vdom-admin"], d, "vdom_admin", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-admin"]) {
			return fmt.Errorf("Error reading vdom-admin: %v", err)
		}
	}

	if err = d.Set("telnet_port", flattenSystemGlobalTelnetPort(o["telnet-port"], d, "telnet_port", sv)); err != nil {
		if !fortiAPIPatch(o["telnet-port"]) {
			return fmt.Errorf("Error reading telnet-port: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", flattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner", sv)); err != nil {
		if !fortiAPIPatch(o["pre-login-banner"]) {
			return fmt.Errorf("Error reading pre-login-banner: %v", err)
		}
	}

	if err = d.Set("https_server_cert", flattenSystemGlobalHttpsServerCert(o["https-server-cert"], d, "https_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["https-server-cert"]) {
			return fmt.Errorf("Error reading https-server-cert: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemGlobalHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("config_sync_enable", flattenSystemGlobalConfigSyncEnable(o["config-sync-enable"], d, "config_sync_enable", sv)); err != nil {
		if !fortiAPIPatch(o["config-sync-enable"]) {
			return fmt.Errorf("Error reading config-sync-enable: %v", err)
		}
	}

	if err = d.Set("vdom_mode", flattenSystemGlobalVdomMode(o["vdom-mode"], d, "vdom_mode", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-mode"]) {
			return fmt.Errorf("Error reading vdom-mode: %v", err)
		}
	}

	if err = d.Set("sys_global_language", flattenSystemGlobalSysGlobalLanguage(o["sys-global-language"], d, "sys_global_language", sv)); err != nil {
		if !fortiAPIPatch(o["sys-global-language"]) {
			return fmt.Errorf("Error reading sys-global-language: %v", err)
		}
	}

	if err = d.Set("gui_log", flattenSystemGlobalGuiLog(o["gui-log"], d, "gui_log", sv)); err != nil {
		if !fortiAPIPatch(o["gui-log"]) {
			return fmt.Errorf("Error reading gui-log: %v", err)
		}
	}

	if err = d.Set("ssh_port", flattenSystemGlobalSshPort(o["ssh-port"], d, "ssh_port", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-port"]) {
			return fmt.Errorf("Error reading ssh-port: %v", err)
		}
	}

	if err = d.Set("config_sync_port", flattenSystemGlobalConfigSyncPort(o["config-sync-port"], d, "config_sync_port", sv)); err != nil {
		if !fortiAPIPatch(o["config-sync-port"]) {
			return fmt.Errorf("Error reading config-sync-port: %v", err)
		}
	}

	if err = d.Set("gui_device_latitude", flattenSystemGlobalGuiDeviceLatitude(o["gui-device-latitude"], d, "gui_device_latitude", sv)); err != nil {
		if !fortiAPIPatch(o["gui-device-latitude"]) {
			return fmt.Errorf("Error reading gui-device-latitude: %v", err)
		}
	}

	if err = d.Set("gui_router", flattenSystemGlobalGuiRouter(o["gui-router"], d, "gui_router", sv)); err != nil {
		if !fortiAPIPatch(o["gui-router"]) {
			return fmt.Errorf("Error reading gui-router: %v", err)
		}
	}

	if err = d.Set("default_intermediate_ca_group", flattenSystemGlobalDefaultIntermediateCaGroup(o["default-intermediate-ca-group"], d, "default_intermediate_ca_group", sv)); err != nil {
		if !fortiAPIPatch(o["default-intermediate-ca-group"]) {
			return fmt.Errorf("Error reading default-intermediate-ca-group: %v", err)
		}
	}

	if err = d.Set("admin_idle_timeout", flattenSystemGlobalAdminIdleTimeout(o["admin-idle-timeout"], d, "admin_idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["admin-idle-timeout"]) {
			return fmt.Errorf("Error reading admin-idle-timeout: %v", err)
		}
	}

	if err = d.Set("https_port", flattenSystemGlobalHttpsPort(o["https-port"], d, "https_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https-port: %v", err)
		}
	}

	if err = d.Set("https_redirect", flattenSystemGlobalHttpsRedirect(o["https-redirect"], d, "https_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["https-redirect"]) {
			return fmt.Errorf("Error reading https-redirect: %v", err)
		}
	}

	if err = d.Set("gui_system", flattenSystemGlobalGuiSystem(o["gui-system"], d, "gui_system", sv)); err != nil {
		if !fortiAPIPatch(o["gui-system"]) {
			return fmt.Errorf("Error reading gui-system: %v", err)
		}
	}

	if err = d.Set("http_port", flattenSystemGlobalHttpPort(o["http-port"], d, "http_port", sv)); err != nil {
		if !fortiAPIPatch(o["http-port"]) {
			return fmt.Errorf("Error reading http-port: %v", err)
		}
	}

	if err = d.Set("use_default_hostname", flattenSystemGlobalUseDefaultHostname(o["use-default-hostname"], d, "use_default_hostname", sv)); err != nil {
		if !fortiAPIPatch(o["use-default-hostname"]) {
			return fmt.Errorf("Error reading use-default-hostname: %v", err)
		}
	}

	if err = d.Set("gui_device_longtitude", flattenSystemGlobalGuiDeviceLongtitude(o["gui-device-longtitude"], d, "gui_device_longtitude", sv)); err != nil {
		if !fortiAPIPatch(o["gui-device-longtitude"]) {
			return fmt.Errorf("Error reading gui-device-longtitude: %v", err)
		}
	}

	if err = d.Set("ip_primary", flattenSystemGlobalPrimaryDNS(o["ip_primary"], d, "ip_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip_primary"]) {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	if err = d.Set("ip_second", flattenSystemGlobalSecondaryDNS(o["ip_second"], d, "ip_second", sv)); err != nil {
		if !fortiAPIPatch(o["ip_second"]) {
			return fmt.Errorf("Error reading ip_second: %v", err)
		}
	}

	return nil
}

func expandSystemGlobalVdomAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTelnetPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBanner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpsServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalConfigSyncEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVdomMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSysGlobalLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalConfigSyncPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDeviceLatitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiRouter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDefaultIntermediateCaGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpsRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiSystem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUseDefaultHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDeviceLongtitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPrimaryDNS(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSecondaryDNS(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vdom_admin"); ok {
		if setArgNil {
			obj["vdom_admin"] = nil
		} else {
			t, err := expandSystemGlobalVdomAdmin(d, v, "vdom_admin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom-admin"] = t
			}
		}
	}

	if v, ok := d.GetOk("telnet_port"); ok {
		if setArgNil {
			obj["telnet_port"] = nil
		} else {
			t, err := expandSystemGlobalTelnetPort(d, v, "telnet_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["telnet-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("pre_login_banner"); ok {
		if setArgNil {
			obj["pre_login_banner"] = nil
		} else {
			t, err := expandSystemGlobalPreLoginBanner(d, v, "pre_login_banner", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pre-login-banner"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_server_cert"); ok {
		if setArgNil {
			obj["https_server_cert"] = nil
		} else {
			t, err := expandSystemGlobalHttpsServerCert(d, v, "https_server_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-server-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		if setArgNil {
			obj["hostname"] = nil
		} else {
			t, err := expandSystemGlobalHostname(d, v, "hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_sync_enable"); ok {
		if setArgNil {
			obj["config_sync_enable"] = nil
		} else {
			t, err := expandSystemGlobalConfigSyncEnable(d, v, "config_sync_enable", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-sync-enable"] = t
			}
		}
	}

	if v, ok := d.GetOk("vdom_mode"); ok {
		if setArgNil {
			obj["vdom_mode"] = nil
		} else {
			t, err := expandSystemGlobalVdomMode(d, v, "vdom_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("sys_global_language"); ok {
		if setArgNil {
			obj["sys_global_language"] = nil
		} else {
			t, err := expandSystemGlobalSysGlobalLanguage(d, v, "sys_global_language", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sys-global-language"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui_log"); ok {
		if setArgNil {
			obj["gui_log"] = nil
		} else {
			t, err := expandSystemGlobalGuiLog(d, v, "gui_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_port"); ok {
		if setArgNil {
			obj["ssh_port"] = nil
		} else {
			t, err := expandSystemGlobalSshPort(d, v, "ssh_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_sync_port"); ok {
		if setArgNil {
			obj["config_sync_port"] = nil
		} else {
			t, err := expandSystemGlobalConfigSyncPort(d, v, "config_sync_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-sync-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui_device_latitude"); ok {
		if setArgNil {
			obj["gui_device_latitude"] = nil
		} else {
			t, err := expandSystemGlobalGuiDeviceLatitude(d, v, "gui_device_latitude", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui-device-latitude"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui_router"); ok {
		if setArgNil {
			obj["gui_router"] = nil
		} else {
			t, err := expandSystemGlobalGuiRouter(d, v, "gui_router", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui-router"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_intermediate_ca_group"); ok {
		if setArgNil {
			obj["default_intermediate_ca_group"] = nil
		} else {
			t, err := expandSystemGlobalDefaultIntermediateCaGroup(d, v, "default_intermediate_ca_group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-intermediate-ca-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_idle_timeout"); ok {
		if setArgNil {
			obj["admin_idle_timeout"] = nil
		} else {
			t, err := expandSystemGlobalAdminIdleTimeout(d, v, "admin_idle_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-idle-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_port"); ok {
		if setArgNil {
			obj["https_port"] = nil
		} else {
			t, err := expandSystemGlobalHttpsPort(d, v, "https_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_redirect"); ok {
		if setArgNil {
			obj["https_redirect"] = nil
		} else {
			t, err := expandSystemGlobalHttpsRedirect(d, v, "https_redirect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-redirect"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui_system"); ok {
		if setArgNil {
			obj["gui_system"] = nil
		} else {
			t, err := expandSystemGlobalGuiSystem(d, v, "gui_system", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui-system"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_port"); ok {
		if setArgNil {
			obj["http_port"] = nil
		} else {
			t, err := expandSystemGlobalHttpPort(d, v, "http_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("use_default_hostname"); ok {
		if setArgNil {
			obj["use_default_hostname"] = nil
		} else {
			t, err := expandSystemGlobalUseDefaultHostname(d, v, "use_default_hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["use-default-hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui_device_longtitude"); ok {
		if setArgNil {
			obj["gui_device_longtitude"] = nil
		} else {
			t, err := expandSystemGlobalGuiDeviceLongtitude(d, v, "gui_device_longtitude", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui-device-longtitude"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_primary"); ok {
		if setArgNil {
			obj["ip_primary"] = nil
		} else {
			t, err := expandSystemGlobalPrimaryDNS(d, v, "ip_primay", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip_primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_second"); ok {
		if setArgNil {
			obj["ip_second"] = nil
		} else {
			t, err := expandSystemGlobalSecondaryDNS(d, v, "ip_second", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip_second"] = t
			}
		}
	}

	return &obj, nil
}
