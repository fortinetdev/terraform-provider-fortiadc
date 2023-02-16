// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system global.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemGlobalRead,
		Schema: map[string]*schema.Schema{
			"vdom_admin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"telnet_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"https_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_sync_enable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sys_global_language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_sync_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_device_latitude": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_router": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_intermediate_ca_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_system": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_default_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_device_longtitude": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "SystemGlobal"

	o, err := c.ReadSystemGlobal(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemGlobal: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemGlobal from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemGlobalVdomAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTelnetPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHttpsServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalConfigSyncEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVdomMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSysGlobalLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalConfigSyncPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDeviceLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiRouter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDefaultIntermediateCaGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUseDefaultHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDeviceLongtitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vdom_admin", dataSourceFlattenSystemGlobalVdomAdmin(o["vdom-admin"], d, "vdom_admin")); err != nil {
		if !fortiAPIPatch(o["vdom-admin"]) {
			return fmt.Errorf("Error reading vdom-admin: %v", err)
		}
	}

	if err = d.Set("telnet_port", dataSourceFlattenSystemGlobalTelnetPort(o["telnet-port"], d, "telnet_port")); err != nil {
		if !fortiAPIPatch(o["telnet-port"]) {
			return fmt.Errorf("Error reading telnet-port: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", dataSourceFlattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner")); err != nil {
		if !fortiAPIPatch(o["pre-login-banner"]) {
			return fmt.Errorf("Error reading pre-login-banner: %v", err)
		}
	}

	if err = d.Set("https_server_cert", dataSourceFlattenSystemGlobalHttpsServerCert(o["https-server-cert"], d, "https_server_cert")); err != nil {
		if !fortiAPIPatch(o["https-server-cert"]) {
			return fmt.Errorf("Error reading https-server-cert: %v", err)
		}
	}

	if err = d.Set("hostname", dataSourceFlattenSystemGlobalHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("config_sync_enable", dataSourceFlattenSystemGlobalConfigSyncEnable(o["config-sync-enable"], d, "config_sync_enable")); err != nil {
		if !fortiAPIPatch(o["config-sync-enable"]) {
			return fmt.Errorf("Error reading config-sync-enable: %v", err)
		}
	}

	if err = d.Set("vdom_mode", dataSourceFlattenSystemGlobalVdomMode(o["vdom-mode"], d, "vdom_mode")); err != nil {
		if !fortiAPIPatch(o["vdom-mode"]) {
			return fmt.Errorf("Error reading vdom-mode: %v", err)
		}
	}

	if err = d.Set("sys_global_language", dataSourceFlattenSystemGlobalSysGlobalLanguage(o["sys-global-language"], d, "sys_global_language")); err != nil {
		if !fortiAPIPatch(o["sys-global-language"]) {
			return fmt.Errorf("Error reading sys-global-language: %v", err)
		}
	}

	if err = d.Set("gui_log", dataSourceFlattenSystemGlobalGuiLog(o["gui-log"], d, "gui_log")); err != nil {
		if !fortiAPIPatch(o["gui-log"]) {
			return fmt.Errorf("Error reading gui-log: %v", err)
		}
	}

	if err = d.Set("ssh_port", dataSourceFlattenSystemGlobalSshPort(o["ssh-port"], d, "ssh_port")); err != nil {
		if !fortiAPIPatch(o["ssh-port"]) {
			return fmt.Errorf("Error reading ssh-port: %v", err)
		}
	}

	if err = d.Set("config_sync_port", dataSourceFlattenSystemGlobalConfigSyncPort(o["config-sync-port"], d, "config_sync_port")); err != nil {
		if !fortiAPIPatch(o["config-sync-port"]) {
			return fmt.Errorf("Error reading config-sync-port: %v", err)
		}
	}

	if err = d.Set("gui_device_latitude", dataSourceFlattenSystemGlobalGuiDeviceLatitude(o["gui-device-latitude"], d, "gui_device_latitude")); err != nil {
		if !fortiAPIPatch(o["gui-device-latitude"]) {
			return fmt.Errorf("Error reading gui-device-latitude: %v", err)
		}
	}

	if err = d.Set("gui_router", dataSourceFlattenSystemGlobalGuiRouter(o["gui-router"], d, "gui_router")); err != nil {
		if !fortiAPIPatch(o["gui-router"]) {
			return fmt.Errorf("Error reading gui-router: %v", err)
		}
	}

	if err = d.Set("default_intermediate_ca_group", dataSourceFlattenSystemGlobalDefaultIntermediateCaGroup(o["default-intermediate-ca-group"], d, "default_intermediate_ca_group")); err != nil {
		if !fortiAPIPatch(o["default-intermediate-ca-group"]) {
			return fmt.Errorf("Error reading default-intermediate-ca-group: %v", err)
		}
	}

	if err = d.Set("admin_idle_timeout", dataSourceFlattenSystemGlobalAdminIdleTimeout(o["admin-idle-timeout"], d, "admin_idle_timeout")); err != nil {
		if !fortiAPIPatch(o["admin-idle-timeout"]) {
			return fmt.Errorf("Error reading admin-idle-timeout: %v", err)
		}
	}

	if err = d.Set("https_port", dataSourceFlattenSystemGlobalHttpsPort(o["https-port"], d, "https_port")); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https-port: %v", err)
		}
	}

	if err = d.Set("https_redirect", dataSourceFlattenSystemGlobalHttpsRedirect(o["https-redirect"], d, "https_redirect")); err != nil {
		if !fortiAPIPatch(o["https-redirect"]) {
			return fmt.Errorf("Error reading https-redirect: %v", err)
		}
	}

	if err = d.Set("gui_system", dataSourceFlattenSystemGlobalGuiSystem(o["gui-system"], d, "gui_system")); err != nil {
		if !fortiAPIPatch(o["gui-system"]) {
			return fmt.Errorf("Error reading gui-system: %v", err)
		}
	}

	if err = d.Set("http_port", dataSourceFlattenSystemGlobalHttpPort(o["http-port"], d, "http_port")); err != nil {
		if !fortiAPIPatch(o["http-port"]) {
			return fmt.Errorf("Error reading http-port: %v", err)
		}
	}

	if err = d.Set("use_default_hostname", dataSourceFlattenSystemGlobalUseDefaultHostname(o["use-default-hostname"], d, "use_default_hostname")); err != nil {
		if !fortiAPIPatch(o["use-default-hostname"]) {
			return fmt.Errorf("Error reading use-default-hostname: %v", err)
		}
	}

	if err = d.Set("gui_device_longtitude", dataSourceFlattenSystemGlobalGuiDeviceLongtitude(o["gui-device-longtitude"], d, "gui_device_longtitude")); err != nil {
		if !fortiAPIPatch(o["gui-device-longtitude"]) {
			return fmt.Errorf("Error reading gui-device-longtitude: %v", err)
		}
	}

	return nil
}
