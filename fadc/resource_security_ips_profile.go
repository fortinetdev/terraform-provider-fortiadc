// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security ips profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityIpsProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityIpsProfileRead,
		Update: resourceSecurityIpsProfileUpdate,
		Create: resourceSecurityIpsProfileCreate,
		Delete: resourceSecurityIpsProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ips_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
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
func resourceSecurityIpsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("ips_profile_name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityIpsProfile: type error")
	}

	obj, err := getObjectSecurityIpsProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityIpsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSecurityIpsProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SecurityIpsProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityIpsProfileRead(d, m)
}
func resourceSecurityIpsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityIpsProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityIpsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityIpsProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SecurityIpsProfile resource: %v", err)
	}

	return resourceSecurityIpsProfileRead(d, m)
}
func resourceSecurityIpsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSecurityIpsProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SecurityIpsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSecurityIpsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSecurityIpsProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SecurityIpsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityIpsProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityIpsProfile resource from API: %v", err)
	}
	return nil
}
func flattenSecurityIpsProfileIpsProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityIpsProfileComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityIpsProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ips_profile_name", flattenSecurityIpsProfileIpsProfileName(o["ips_profile_name"], d, "ips_profile_name", sv)); err != nil {
		if !fortiAPIPatch(o["ips_profile_name"]) {
			return fmt.Errorf("Error reading ips_profile_name: %v", err)
		}
	}

	if err = d.Set("comments", flattenSecurityIpsProfileComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func expandSecurityIpsProfileIpsProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityIpsProfileComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityIpsProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ips_profile_name"); ok {
		t, err := expandSecurityIpsProfileIpsProfileName(d, v, "ips_profile_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips_profile_name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSecurityIpsProfileComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
