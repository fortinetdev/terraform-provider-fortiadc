// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security ztna profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityZtnaProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityZtnaProfileRead,
		Update: resourceSecurityZtnaProfileUpdate,
		Create: resourceSecurityZtnaProfileCreate,
		Delete: resourceSecurityZtnaProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"log_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
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
		},
	}
}
func resourceSecurityZtnaProfileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityZtnaProfile: type error")
	}

	obj, err := getObjectSecurityZtnaProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityZtnaProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSecurityZtnaProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SecurityZtnaProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityZtnaProfileRead(d, m)
}
func resourceSecurityZtnaProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityZtnaProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityZtnaProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityZtnaProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SecurityZtnaProfile resource: %v", err)
	}

	return resourceSecurityZtnaProfileRead(d, m)
}
func resourceSecurityZtnaProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSecurityZtnaProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SecurityZtnaProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSecurityZtnaProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSecurityZtnaProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SecurityZtnaProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityZtnaProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityZtnaProfile resource from API: %v", err)
	}
	return nil
}
func flattenSecurityZtnaProfileLogStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityZtnaProfileComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityZtnaProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityZtnaProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("log_status", flattenSecurityZtnaProfileLogStatus(o["log_status"], d, "log_status", sv)); err != nil {
		if !fortiAPIPatch(o["log_status"]) {
			return fmt.Errorf("Error reading log_status: %v", err)
		}
	}

	if err = d.Set("comments", flattenSecurityZtnaProfileComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSecurityZtnaProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityZtnaProfileLogStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityZtnaProfileComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityZtnaProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityZtnaProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_status"); ok {
		t, err := expandSecurityZtnaProfileLogStatus(d, v, "log_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSecurityZtnaProfileComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityZtnaProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
