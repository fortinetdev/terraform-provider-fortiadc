// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf sig profile category id list group

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafSigProfileCategoryIdListGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafSigProfileCategoryIdListGroupRead,
		Update: resourceSecurityWafSigProfileCategoryIdListGroupUpdate,
		Create: resourceSecurityWafSigProfileCategoryIdListGroupUpdate,
		Delete: resourceSecurityWafSigProfileCategoryIdListGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"values": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mkey": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSecurityWafSigProfileCategoryIdListGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := ""
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
		return fmt.Errorf("Error describing SecurityWafSigProfileCategoryIdListGroup: type error")
	}

	obj, err := getObjectSecurityWafSigProfileCategoryIdListGroup(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafWebAttackSignatureChildSignature resource while getting object: %v", err)
	}

	path := "/api/security_waf_sig_profile_category_id_list_group?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(pkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafSigProfileCategoryIdListGroup resource: %v", err)
	}

	d.SetId("SecurityWafSigProfileCategoryIdListGroup")
	return resourceSecurityWafSigProfileCategoryIdListGroupRead(d, m)
}

func resourceSecurityWafSigProfileCategoryIdListGroupDelete(d *schema.ResourceData, m interface{}) error {

	d.SetId("")

	return nil
}

func resourceSecurityWafSigProfileCategoryIdListGroupRead(d *schema.ResourceData, m interface{}) error {

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
		return fmt.Errorf("Error describing SecurityWafSigProfileCategoryIdListGroup: type error")
	}

	path := "/api/security_waf_web_attack_signature_child_category?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	o, err := c.StandardRead("", vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafSigProfileCategoryIdListGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafSigProfileCategoryIdListGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafSigProfileCategoryIdListGroup resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafSigProfileCategoryIdListGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafSigProfileCategoryIdListGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenSecurityWafSigProfileCategoryIdListGroup(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("values", flattenSecurityWafSigProfileCategoryIdListGroup(o["values"], d, "values", sv)); err != nil {
		if !fortiAPIPatch(o["values"]) {
			return fmt.Errorf("Error reading values: %v", err)
		}
	}

	return nil
}

func expandSecurityWafSigProfileCategoryIdListGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafSigProfileCategoryIdListGroup(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		if setArgNil {
			obj["type"] = nil
		} else {
			t, err := expandSecurityWafSigProfileCategoryIdListGroup(d, v, "type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["type"] = t
			}
		}
	}

	if v, ok := d.GetOk("values"); ok {
		if setArgNil {
			obj["values"] = nil
		} else {
			t, err := expandSecurityWafSigProfileCategoryIdListGroup(d, v, "vaules", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["values"] = t
			}
		}
	}

	return &obj, nil
}
