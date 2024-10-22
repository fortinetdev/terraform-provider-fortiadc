// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf web attack signature child category.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafWebAttackSignatureChildCategory() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafWebAttackSignatureChildCategoryRead,
		Update: resourceSecurityWafWebAttackSignatureChildCategoryUpdate,
		Create: resourceSecurityWafWebAttackSignatureChildCategoryUpdate,
		Delete: resourceSecurityWafWebAttackSignatureChildCategoryDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkey": &schema.Schema{
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
						"name": &schema.Schema{
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

func resourceSecurityWafWebAttackSignatureChildCategoryUpdate(d *schema.ResourceData, m interface{}) error {
	d.SetId("SecurityWafWebAttackSignatureChildCategory")
	return resourceSecurityWafWebAttackSignatureChildCategoryRead(d, m)
}

func resourceSecurityWafWebAttackSignatureChildCategoryDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func resourceSecurityWafWebAttackSignatureChildCategoryRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafWebAttackSignatureChildCategory: type error")
	}

	path := "/api/security_waf_web_attack_signature_child_category?pkey=" + escapeURLString(pkey)
	o, err := c.StandardRead("", vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafWebAttackSignatureChildCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafWebAttackSignatureChildCategory(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafSigProfileCategoryIdListGroup resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafWebAttackSignatureChildCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafWebAttackSignatureChildCategory(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("values", flattenSecurityWafWebAttackSignatureChildCategory(o["values"], d, "values", sv)); err != nil {
		if !fortiAPIPatch(o["values"]) {
			return fmt.Errorf("Error reading values: %v", err)
		}
	}

	return nil
}
