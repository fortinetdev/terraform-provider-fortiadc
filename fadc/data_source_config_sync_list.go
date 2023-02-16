// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  config sync list.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceConfigSyncList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceConfigSyncListRead,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceConfigSyncListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("server-ip")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing ConfigSyncList: type error")
	}

	o, err := c.ReadConfigSyncList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing ConfigSyncList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectConfigSyncList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing ConfigSyncList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenConfigSyncListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenConfigSyncListMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectConfigSyncList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", dataSourceFlattenConfigSyncListComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenConfigSyncListStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("filename", dataSourceFlattenConfigSyncListFilename(o["filename"], d, "filename")); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("server_ip", dataSourceFlattenConfigSyncListServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if !fortiAPIPatch(o["server-ip"]) {
			return fmt.Errorf("Error reading server-ip: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenConfigSyncListPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenConfigSyncListType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenConfigSyncListPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenConfigSyncListMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
