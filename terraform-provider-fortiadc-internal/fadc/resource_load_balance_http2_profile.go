// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance http2 profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceHttp2Profile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceHttp2ProfileRead,
		Update: resourceLoadBalanceHttp2ProfileUpdate,
		Create: resourceLoadBalanceHttp2ProfileCreate,
		Delete: resourceLoadBalanceHttp2ProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_concurrent_stream": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_frame_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"priority_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_constraint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_receive_window": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"header_table_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"upgrade_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_header_list_size": &schema.Schema{
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
func resourceLoadBalanceHttp2ProfileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceHttp2Profile: type error")
	}

	obj, err := getObjectLoadBalanceHttp2Profile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceHttp2Profile resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceHttp2Profile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceHttp2Profile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceHttp2ProfileRead(d, m)
}
func resourceLoadBalanceHttp2ProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceHttp2Profile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceHttp2Profile resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceHttp2Profile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceHttp2Profile resource: %v", err)
	}

	return resourceLoadBalanceHttp2ProfileRead(d, m)
}
func resourceLoadBalanceHttp2ProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceHttp2Profile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceHttp2Profile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceHttp2ProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceHttp2Profile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceHttp2Profile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceHttp2Profile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceHttp2Profile resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceHttp2ProfileMaxConcurrentStream(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileMaxFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfilePriorityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileSslConstraint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileMaxReceiveWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileHeaderTableSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileUpgradeMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceHttp2ProfileMaxHeaderListSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceHttp2Profile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_concurrent_stream", flattenLoadBalanceHttp2ProfileMaxConcurrentStream(o["max_concurrent_stream"], d, "max_concurrent_stream", sv)); err != nil {
		if !fortiAPIPatch(o["max_concurrent_stream"]) {
			return fmt.Errorf("Error reading max_concurrent_stream: %v", err)
		}
	}

	if err = d.Set("max_frame_size", flattenLoadBalanceHttp2ProfileMaxFrameSize(o["max_frame_size"], d, "max_frame_size", sv)); err != nil {
		if !fortiAPIPatch(o["max_frame_size"]) {
			return fmt.Errorf("Error reading max_frame_size: %v", err)
		}
	}

	if err = d.Set("priority_mode", flattenLoadBalanceHttp2ProfilePriorityMode(o["priority_mode"], d, "priority_mode", sv)); err != nil {
		if !fortiAPIPatch(o["priority_mode"]) {
			return fmt.Errorf("Error reading priority_mode: %v", err)
		}
	}

	if err = d.Set("ssl_constraint", flattenLoadBalanceHttp2ProfileSslConstraint(o["ssl_constraint"], d, "ssl_constraint", sv)); err != nil {
		if !fortiAPIPatch(o["ssl_constraint"]) {
			return fmt.Errorf("Error reading ssl_constraint: %v", err)
		}
	}

	if err = d.Set("max_receive_window", flattenLoadBalanceHttp2ProfileMaxReceiveWindow(o["max_receive_window"], d, "max_receive_window", sv)); err != nil {
		if !fortiAPIPatch(o["max_receive_window"]) {
			return fmt.Errorf("Error reading max_receive_window: %v", err)
		}
	}

	if err = d.Set("header_table_size", flattenLoadBalanceHttp2ProfileHeaderTableSize(o["header_table_size"], d, "header_table_size", sv)); err != nil {
		if !fortiAPIPatch(o["header_table_size"]) {
			return fmt.Errorf("Error reading header_table_size: %v", err)
		}
	}

	if err = d.Set("upgrade_mode", flattenLoadBalanceHttp2ProfileUpgradeMode(o["upgrade_mode"], d, "upgrade_mode", sv)); err != nil {
		if !fortiAPIPatch(o["upgrade_mode"]) {
			return fmt.Errorf("Error reading upgrade_mode: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceHttp2ProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("max_header_list_size", flattenLoadBalanceHttp2ProfileMaxHeaderListSize(o["max_header_list_size"], d, "max_header_list_size", sv)); err != nil {
		if !fortiAPIPatch(o["max_header_list_size"]) {
			return fmt.Errorf("Error reading max_header_list_size: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceHttp2ProfileMaxConcurrentStream(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileMaxFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfilePriorityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileSslConstraint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileMaxReceiveWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileHeaderTableSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileUpgradeMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceHttp2ProfileMaxHeaderListSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceHttp2Profile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_concurrent_stream"); ok {
		t, err := expandLoadBalanceHttp2ProfileMaxConcurrentStream(d, v, "max_concurrent_stream", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_concurrent_stream"] = t
		}
	}

	if v, ok := d.GetOk("max_frame_size"); ok {
		t, err := expandLoadBalanceHttp2ProfileMaxFrameSize(d, v, "max_frame_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_frame_size"] = t
		}
	}

	if v, ok := d.GetOk("priority_mode"); ok {
		t, err := expandLoadBalanceHttp2ProfilePriorityMode(d, v, "priority_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority_mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_constraint"); ok {
		t, err := expandLoadBalanceHttp2ProfileSslConstraint(d, v, "ssl_constraint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl_constraint"] = t
		}
	}

	if v, ok := d.GetOk("max_receive_window"); ok {
		t, err := expandLoadBalanceHttp2ProfileMaxReceiveWindow(d, v, "max_receive_window", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_receive_window"] = t
		}
	}

	if v, ok := d.GetOk("header_table_size"); ok {
		t, err := expandLoadBalanceHttp2ProfileHeaderTableSize(d, v, "header_table_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header_table_size"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_mode"); ok {
		t, err := expandLoadBalanceHttp2ProfileUpgradeMode(d, v, "upgrade_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade_mode"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceHttp2ProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("max_header_list_size"); ok {
		t, err := expandLoadBalanceHttp2ProfileMaxHeaderListSize(d, v, "max_header_list_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_header_list_size"] = t
		}
	}

	return &obj, nil
}
