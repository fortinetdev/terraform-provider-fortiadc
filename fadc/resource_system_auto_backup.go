// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system auto backup.

package fortiadc

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAutoBackup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAutoBackupRead,
		Update: resourceSystemAutoBackupUpdate,
		Create: resourceSystemAutoBackupUpdate,
		Delete: resourceSystemAutoBackupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"overwrite_config": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"config_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scheduled_backup_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scheduled_backup_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"storage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scheduled_backup_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scheduled_backup_day": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
func resourceSystemAutoBackupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAutoBackup(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoBackup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoBackup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoBackup resource: %v", err)
	}

	d.SetId("SystemAutoBackup")
	return resourceSystemAutoBackupRead(d, m)
}
func resourceSystemAutoBackupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAutoBackup(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoBackup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoBackup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAutoBackup resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemAutoBackupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAutoBackup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoBackup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoBackup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoBackup resource from API: %v", err)
	}
	return nil
}
func flattenSystemAutoBackupUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupOverwriteConfig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupConfigPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupScheduledBackupStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupScheduledBackupFrequency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupStorage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupScheduledBackupTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupScheduledBackupDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupFolder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoBackupPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutoBackup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("username", flattenSystemAutoBackupUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("overwrite_config", flattenSystemAutoBackupOverwriteConfig(o["overwrite-config"], d, "overwrite_config", sv)); err != nil {
		if !fortiAPIPatch(o["overwrite-config"]) {
			return fmt.Errorf("Error reading overwrite-config: %v", err)
		}
	}

	if err = d.Set("config_password", flattenSystemAutoBackupConfigPassword(o["config-password"], d, "config_password", sv)); err != nil {
		if !fortiAPIPatch(o["config-password"]) {
			return fmt.Errorf("Error reading config-password: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_status", flattenSystemAutoBackupScheduledBackupStatus(o["scheduled-backup-status"], d, "scheduled_backup_status", sv)); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-status"]) {
			return fmt.Errorf("Error reading scheduled-backup-status: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_frequency", flattenSystemAutoBackupScheduledBackupFrequency(o["scheduled-backup-frequency"], d, "scheduled_backup_frequency", sv)); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-frequency"]) {
			return fmt.Errorf("Error reading scheduled-backup-frequency: %v", err)
		}
	}

	if err = d.Set("storage", flattenSystemAutoBackupStorage(o["storage"], d, "storage", sv)); err != nil {
		if !fortiAPIPatch(o["storage"]) {
			return fmt.Errorf("Error reading storage: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_time", flattenSystemAutoBackupScheduledBackupTime(o["scheduled-backup-time"], d, "scheduled_backup_time", sv)); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-time"]) {
			return fmt.Errorf("Error reading scheduled-backup-time: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemAutoBackupAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("scheduled_backup_day", flattenSystemAutoBackupScheduledBackupDay(o["scheduled-backup-day"], d, "scheduled_backup_day", sv)); err != nil {
		if !fortiAPIPatch(o["scheduled-backup-day"]) {
			return fmt.Errorf("Error reading scheduled-backup-day: %v", err)
		}
	}

	if err = d.Set("folder", flattenSystemAutoBackupFolder(o["folder"], d, "folder", sv)); err != nil {
		if !fortiAPIPatch(o["folder"]) {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemAutoBackupPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutoBackupPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func expandSystemAutoBackupUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupOverwriteConfig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupConfigPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupScheduledBackupStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupScheduledBackupFrequency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupStorage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupScheduledBackupTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupScheduledBackupDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupFolder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoBackupPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoBackup(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("username"); ok {
		if setArgNil {
			obj["username"] = nil
		} else {
			t, err := expandSystemAutoBackupUsername(d, v, "username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username"] = t
			}
		}
	}

	if v, ok := d.GetOk("overwrite_config"); ok {
		if setArgNil {
			obj["overwrite_config"] = nil
		} else {
			t, err := expandSystemAutoBackupOverwriteConfig(d, v, "overwrite_config", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["overwrite-config"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_password"); ok {
		if setArgNil {
			obj["config_password"] = nil
		} else {
			t, err := expandSystemAutoBackupConfigPassword(d, v, "config_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-password"] = t
			}
		}
	}

	if v, ok := d.GetOk("scheduled_backup_status"); ok {
		if setArgNil {
			obj["scheduled_backup_status"] = nil
		} else {
			t, err := expandSystemAutoBackupScheduledBackupStatus(d, v, "scheduled_backup_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scheduled-backup-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("scheduled_backup_frequency"); ok {
		if setArgNil {
			obj["scheduled_backup_frequency"] = nil
		} else {
			t, err := expandSystemAutoBackupScheduledBackupFrequency(d, v, "scheduled_backup_frequency", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scheduled-backup-frequency"] = t
			}
		}
	}

	if v, ok := d.GetOk("storage"); ok {
		if setArgNil {
			obj["storage"] = nil
		} else {
			t, err := expandSystemAutoBackupStorage(d, v, "storage", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["storage"] = t
			}
		}
	}

	if v, ok := d.GetOk("scheduled_backup_time"); ok {
		if setArgNil {
			obj["scheduled_backup_time"] = nil
		} else {
			t, err := expandSystemAutoBackupScheduledBackupTime(d, v, "scheduled_backup_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scheduled-backup-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("address"); ok {
		if setArgNil {
			obj["address"] = nil
		} else {
			t, err := expandSystemAutoBackupAddress(d, v, "address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address"] = t
			}
		}
	}

	if v, ok := d.GetOk("scheduled_backup_day"); ok {
		if setArgNil {
			obj["scheduled_backup_day"] = nil
		} else {
			t, err := expandSystemAutoBackupScheduledBackupDay(d, v, "scheduled_backup_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scheduled-backup-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("folder"); ok {
		if setArgNil {
			obj["folder"] = nil
		} else {
			t, err := expandSystemAutoBackupFolder(d, v, "folder", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["folder"] = t
			}
		}
	}

	if v, ok := d.GetOk("password"); ok {
		if setArgNil {
			obj["password"] = nil
		} else {
			t, err := expandSystemAutoBackupPassword(d, v, "password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["password"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {
			t, err := expandSystemAutoBackupPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	return &obj, nil
}
