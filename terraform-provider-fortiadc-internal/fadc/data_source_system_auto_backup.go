// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system auto backup.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemAutoBackup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutoBackupRead,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overwrite_config": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_backup_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_backup_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_backup_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_backup_day": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSystemAutoBackupRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := "SystemAutoBackup"

	o, err := c.ReadSystemAutoBackup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoBackup: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemAutoBackup(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoBackup from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemAutoBackupUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupOverwriteConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupConfigPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupScheduledBackupStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupScheduledBackupFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupStorage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupScheduledBackupTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupScheduledBackupDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoBackupPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutoBackup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("username", dataSourceFlattenSystemAutoBackupUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("overwrite_config", dataSourceFlattenSystemAutoBackupOverwriteConfig(o["overwrite-config"], d, "overwrite_config")); err != nil {
		if !fortiAPIPatch(o["overwrite-config"]) {
			return fmt.Errorf("Error reading overwrite-config: %v", err)
		}
	}

	if err = d.Set("config_password", dataSourceFlattenSystemAutoBackupConfigPassword(o["config-password"], d, "config_password")); err != nil {
		if !fortiAPIPatch(o["config-password"]) {
			return fmt.Errorf("Error reading config-password: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_status", dataSourceFlattenSystemAutoBackupScheduledBackupStatus(o["scheduled-backup-status"], d, "scheduled_backup_status")); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-status"]) {
			return fmt.Errorf("Error reading scheduled-backup-status: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_frequency", dataSourceFlattenSystemAutoBackupScheduledBackupFrequency(o["scheduled-backup-frequency"], d, "scheduled_backup_frequency")); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-frequency"]) {
			return fmt.Errorf("Error reading scheduled-backup-frequency: %v", err)
		}
	}

	if err = d.Set("storage", dataSourceFlattenSystemAutoBackupStorage(o["storage"], d, "storage")); err != nil {
		if !fortiAPIPatch(o["storage"]) {
			return fmt.Errorf("Error reading storage: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_time", dataSourceFlattenSystemAutoBackupScheduledBackupTime(o["scheduled-backup-time"], d, "scheduled_backup_time")); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-time"]) {
			return fmt.Errorf("Error reading scheduled-backup-time: %v", err)
		}
	}

	if err = d.Set("address", dataSourceFlattenSystemAutoBackupAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_day", dataSourceFlattenSystemAutoBackupScheduledBackupDay(o["scheduled-backup-day"], d, "scheduled_backup_day")); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-day"]) {
			return fmt.Errorf("Error reading scheduled-backup-day: %v", err)
		}
	}

	if err = d.Set("folder", dataSourceFlattenSystemAutoBackupFolder(o["folder"], d, "folder")); err != nil {
		if !fortiAPIPatch(o["folder"]) {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenSystemAutoBackupPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemAutoBackupPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}
