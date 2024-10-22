// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf web attack signature child signature.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafWebAttackSignatureChildSignature() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafWebAttackSignatureChildSignatureRead,
		Update: resourceSecurityWafWebAttackSignatureChildSignatureUpdate,
		Create: resourceSecurityWafWebAttackSignatureChildSignatureUpdate,
		Delete: resourceSecurityWafWebAttackSignatureChildSignatureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_name": &schema.Schema{
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
func resourceSecurityWafWebAttackSignatureChildSignatureUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafWebAttackSignatureChildSignature: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafWebAttackSignatureChildSignature: type error")
	}

	obj, err := getObjectSecurityWafWebAttackSignatureChildSignature(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafWebAttackSignatureChildSignature resource while getting object: %v", err)
	}

	path := "/api/security_waf_web_attack_signature_child_signature?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafWebAttackSignatureChildSignature resource: %v", err)
	}

	d.SetId(mkey)
	return resourceSecurityWafWebAttackSignatureChildSignatureRead(d, m)
}
func resourceSecurityWafWebAttackSignatureChildSignatureDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafWebAttackSignatureChildSignature: type error")
	}

	obj, err := getObjectSecurityWafWebAttackSignatureChildSignature(d, true, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafWebAttackSignatureChildSignature resource while getting object: %v", err)
	}

	path := "/api/security_waf_web_attack_signature_child_signature?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafWebAttackSignatureChildSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSecurityWafWebAttackSignatureChildSignatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	t := d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafWebAttackSignatureChildSignature: type error")
	}

	path := "/api/security_waf_web_attack_signature_child_signature?pkey=" + escapeURLString(pkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafWebAttackSignatureChildSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafWebAttackSignatureChildSignature(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafWebAttackSignatureChildSignature resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafWebAttackSignatureChildSignature(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafWebAttackSignatureChildSignature(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafWebAttackSignatureChildSignature(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("status", flattenSecurityWafWebAttackSignatureChildSignature(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("exception_name", flattenSecurityWafWebAttackSignatureChildSignature(o["exception_name"], d, "exception_name", sv)); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}

	return nil
}

func expandSecurityWafWebAttackSignatureChildSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafWebAttackSignatureChildSignature(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		if setArgNil {
			obj["mkey"] = nil
		} else {
			t, err := expandSecurityWafWebAttackSignatureChildSignature(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSecurityWafWebAttackSignatureChildSignature(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("exception_name"); ok {
		if setArgNil {
			obj["exception_name"] = nil
		} else {
			t, err := expandSecurityWafWebAttackSignatureChildSignature(d, v, "exception_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["exception_name"] = t
			}
		}
	}

	return &obj, nil
}
