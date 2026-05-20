// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf adpative learning

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAdaptiveLearning() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafAdaptiveLearningRead,
		Update: resourceSecurityWafAdaptiveLearningUpdate,
		Create: resourceSecurityWafAdaptiveLearningCreate,
		Delete: resourceSecurityWafAdaptiveLearningDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sampling_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fp_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fp_exp_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"least_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
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

func resourceSecurityWafAdaptiveLearningCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearning: type error")
	}

	obj, err := getObjectSecurityWafAdaptiveLearning(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdaptiveLearning resource while getting object: %v", err)
	}

	path := "/api/security_waf_adaptive_learning"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdaptiveLearning resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafAdaptiveLearningRead(d, m)
}

func resourceSecurityWafAdaptiveLearningUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafAdaptiveLearning(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAdaptiveLearning resource while getting object: %v", err)
	}

	path := "/api/security_waf_adaptive_learning?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAdaptiveLearning resource: %v", err)
	}

	return resourceSecurityWafAdaptiveLearningRead(d, m)
}

func resourceSecurityWafAdaptiveLearningDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_adaptive_learning?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAdaptiveLearning resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafAdaptiveLearningRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_adaptive_learning?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdaptiveLearning resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAdaptiveLearning(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdaptiveLearning resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAdaptiveLearning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAdaptiveLearning(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAdaptiveLearning(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("status", flattenSecurityWafAdaptiveLearning(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sampling_rate", flattenSecurityWafAdaptiveLearning(o["sampling_rate"], d, "sampling_rate", sv)); err != nil {
		if !fortiAPIPatch(o["sampling_rate"]) {
			return fmt.Errorf("Error reading sampling_rate: %v", err)
		}
	}

	if err = d.Set("fp_threshold", flattenSecurityWafAdaptiveLearning(o["fp_threshold"], d, "fp_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["fp_threshold"]) {
			return fmt.Errorf("Error reading fp_threshold: %v", err)
		}
	}

	if err = d.Set("fp_exp_id", flattenSecurityWafAdaptiveLearning(o["fp_exp_id"], d, "fp_exp_id", sv)); err != nil {
		if !fortiAPIPatch(o["fp_exp_id"]) {
			return fmt.Errorf("Error reading fp_exp_id: %v", err)
		}
	}

	if err = d.Set("least_time", flattenSecurityWafAdaptiveLearning(o["least_time"], d, "least_time", sv)); err != nil {
		if !fortiAPIPatch(o["least_time"]) {
			return fmt.Errorf("Error reading least_time: %v", err)
		}
	}

	if err = d.Set("action", flattenSecurityWafAdaptiveLearning(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAdaptiveLearning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAdaptiveLearning(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sampling_rate"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "sampling_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sampling_rate"] = t
		}
	}

	if v, ok := d.GetOk("fp_threshold"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "fp_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp_threshold"] = t
		}
	}

	if v, ok := d.GetOk("fp_exp_id"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "fp_exp_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp_exp_id"] = t
		}
	}

	if v, ok := d.GetOk("least_time"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "least_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["least_time"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandSecurityWafAdaptiveLearning(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	return &obj, nil
}
