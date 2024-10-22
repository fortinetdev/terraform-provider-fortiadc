// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf threshold based detection

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafThresholdBasedDetection() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafThresholdBasedDetectionRead,
		Update: resourceSecurityWafThresholdBasedDetectionUpdate,
		Create: resourceSecurityWafThresholdBasedDetectionCreate,
		Delete: resourceSecurityWafThresholdBasedDetectionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crawler_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"response_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crawler_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crawler_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crawler_occurrence_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crawler_occurrence_within": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_scraping_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_occurrence_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content_occurrence_within": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attack_detection_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attack_modules": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attack_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attack_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attack_occurrence_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attack_occurrence_within": &schema.Schema{
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

func resourceSecurityWafThresholdBasedDetectionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafThresholdBasedDetection: type error")
	}

	obj, err := getObjectSecurityWafThresholdBasedDetection(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafThresholdBasedDetection resource while getting object: %v", err)
	}

	path := "/api/security_waf_threshold_based_detection"
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafThresholdBasedDetection resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityWafThresholdBasedDetectionRead(d, m)
}

func resourceSecurityWafThresholdBasedDetectionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityWafThresholdBasedDetection(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafThresholdBasedDetection resource while getting object: %v", err)
	}

	path := "/api/security_waf_threshold_based_detection?mkey=" + escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafThresholdBasedDetection resource: %v", err)
	}

	return resourceSecurityWafThresholdBasedDetectionRead(d, m)
}

func resourceSecurityWafThresholdBasedDetectionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_threshold_based_detection?mkey=" + escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafThresholdBasedDetection resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafThresholdBasedDetectionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	path := "/api/security_waf_threshold_based_detection?mkey=" + escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafThresholdBasedDetection resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafThresholdBasedDetection(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafThresholdBasedDetection resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafThresholdBasedDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafThresholdBasedDetection(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafThresholdBasedDetection(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("comments", flattenSecurityWafThresholdBasedDetection(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("crawler_status", flattenSecurityWafThresholdBasedDetection(o["crawler-status"], d, "crawler_status", sv)); err != nil {
		if !fortiAPIPatch(o["crawler-status"]) {
			return fmt.Errorf("Error reading crawler_status: %v", err)
		}
	}

	if err = d.Set("response_code", flattenSecurityWafThresholdBasedDetection(o["response-code"], d, "response_code", sv)); err != nil {
		if !fortiAPIPatch(o["response-code"]) {
			return fmt.Errorf("Error reading response_code: %v", err)
		}
	}

	if err = d.Set("crawler_action", flattenSecurityWafThresholdBasedDetection(o["crawler-action"], d, "crawler_action", sv)); err != nil {
		if !fortiAPIPatch(o["crawler-action"]) {
			return fmt.Errorf("Error reading crawler_action: %v", err)
		}
	}

	if err = d.Set("crawler_severity", flattenSecurityWafThresholdBasedDetection(o["crawler-severity"], d, "crawler_severity", sv)); err != nil {
		if !fortiAPIPatch(o["crawler-severity"]) {
			return fmt.Errorf("Error reading crawler_severity: %v", err)
		}
	}

	if err = d.Set("crawler_occurrence_limit", flattenSecurityWafThresholdBasedDetection(o["crawler-occurrence-limit"], d, "crawler_occurrence_limit", sv)); err != nil {
		if !fortiAPIPatch(o["crawler-occurrence-limit"]) {
			return fmt.Errorf("Error reading crawler_occurrence_limit: %v", err)
		}
	}

	if err = d.Set("crawler_occurrence_within", flattenSecurityWafThresholdBasedDetection(o["crawler-occurrence-within"], d, "crawler_occurrence_within", sv)); err != nil {
		if !fortiAPIPatch(o["crawler-occurrence-within"]) {
			return fmt.Errorf("Error reading crawler_occurrence_within: %v", err)
		}
	}

	if err = d.Set("content_scraping_status", flattenSecurityWafThresholdBasedDetection(o["content-scraping-status"], d, "content_scraping_status", sv)); err != nil {
		if !fortiAPIPatch(o["content-scraping-status"]) {
			return fmt.Errorf("Error reading content_scraping_status: %v", err)
		}
	}

	if err = d.Set("content_type", flattenSecurityWafThresholdBasedDetection(o["content-type"], d, "content_type", sv)); err != nil {
		if !fortiAPIPatch(o["content-type"]) {
			return fmt.Errorf("Error reading content_type: %v", err)
		}
	}

	if err = d.Set("content_action", flattenSecurityWafThresholdBasedDetection(o["content-action"], d, "content_action", sv)); err != nil {
		if !fortiAPIPatch(o["content-action"]) {
			return fmt.Errorf("Error reading content_action: %v", err)
		}
	}

	if err = d.Set("content_severity", flattenSecurityWafThresholdBasedDetection(o["content-severity"], d, "content_severity", sv)); err != nil {
		if !fortiAPIPatch(o["content-severity"]) {
			return fmt.Errorf("Error reading content_severity: %v", err)
		}
	}

	if err = d.Set("content_occurrence_limit", flattenSecurityWafThresholdBasedDetection(o["content-occurrence-limit"], d, "content_occurrence_limit", sv)); err != nil {
		if !fortiAPIPatch(o["content-occurrence-limit"]) {
			return fmt.Errorf("Error reading content_occurrence_limit: %v", err)
		}
	}

	if err = d.Set("content_occurrence_within", flattenSecurityWafThresholdBasedDetection(o["content-occurrence-within"], d, "content_occurrence_within", sv)); err != nil {
		if !fortiAPIPatch(o["content-occurrence-within"]) {
			return fmt.Errorf("Error reading content_occurrence_within: %v", err)
		}
	}

	if err = d.Set("attack_detection_status", flattenSecurityWafThresholdBasedDetection(o["attack-detection-status"], d, "attack_detection_status", sv)); err != nil {
		if !fortiAPIPatch(o["attack-detection-status"]) {
			return fmt.Errorf("Error reading attack_detection_status: %v", err)
		}
	}

	if err = d.Set("attack_modules", flattenSecurityWafThresholdBasedDetection(o["attack-modules"], d, "attack_modules", sv)); err != nil {
		if !fortiAPIPatch(o["attack-modules"]) {
			return fmt.Errorf("Error reading attack_modules: %v", err)
		}
	}

	if err = d.Set("attack_action", flattenSecurityWafThresholdBasedDetection(o["attack-action"], d, "attack_action", sv)); err != nil {
		if !fortiAPIPatch(o["attack-action"]) {
			return fmt.Errorf("Error reading attack_action: %v", err)
		}
	}

	if err = d.Set("attack_severity", flattenSecurityWafThresholdBasedDetection(o["attack-severity"], d, "attack_severity", sv)); err != nil {
		if !fortiAPIPatch(o["attack-severity"]) {
			return fmt.Errorf("Error reading attack_severity: %v", err)
		}
	}

	if err = d.Set("attack_occurrence_limit", flattenSecurityWafThresholdBasedDetection(o["attack-occurrence-within"], d, "attack_occurrence_limit", sv)); err != nil {
		if !fortiAPIPatch(o["attack-occurrence-within"]) {
			return fmt.Errorf("Error reading attack_occurrence_limit: %v", err)
		}
	}

	if err = d.Set("attack_occurrence_within", flattenSecurityWafThresholdBasedDetection(o["attack-occurrence-within"], d, "attack_occurrence_within", sv)); err != nil {
		if !fortiAPIPatch(o["attack-occurrence-within"]) {
			return fmt.Errorf("Error reading attack_occurrence_within: %v", err)
		}
	}

	return nil
}

func expandSecurityWafThresholdBasedDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafThresholdBasedDetection(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("crawler_status"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "crawler_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawler-status"] = t
		}
	}

	if v, ok := d.GetOk("response_code"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "response_code", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-code"] = t
		}
	}

	if v, ok := d.GetOk("crawler_action"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "crawler_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawler-action"] = t
		}
	}

	if v, ok := d.GetOk("crawler_severity"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "crawler_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawler-severity"] = t
		}
	}

	if v, ok := d.GetOk("crawler_occurrence_limit"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "crawler_occurrence_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawler-occurrence-limit"] = t
		}
	}

	if v, ok := d.GetOk("crawler_occurrence_within"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "crawler_occurrence_within", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawler-occurrence-within"] = t
		}
	}

	if v, ok := d.GetOk("content_scraping_status"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "content_scraping_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-scraping-status"] = t
		}
	}

	if v, ok := d.GetOk("content_type"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "content_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-type"] = t
		}
	}

	if v, ok := d.GetOk("content_action"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "content_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-action"] = t
		}
	}

	if v, ok := d.GetOk("content_severity"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "content_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-severity"] = t
		}
	}

	if v, ok := d.GetOk("content_occurrence_limit"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "content_occurrence_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-occurrence-limit"] = t
		}
	}

	if v, ok := d.GetOk("content_occurrence_within"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "content_occurrence_within", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-occurrence-within"] = t
		}
	}

	if v, ok := d.GetOk("attack_detection_status"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "attack_detection_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attack-detection-status"] = t
		}
	}

	if v, ok := d.GetOk("attack_modules"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "attack_modules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attack-modules"] = t
		}
	}

	if v, ok := d.GetOk("attack_action"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "attack_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attack-action"] = t
		}
	}

	if v, ok := d.GetOk("attack_severity"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "attack_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attack-severity"] = t
		}
	}

	if v, ok := d.GetOk("attack_occurrence_limit"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "attack_occurrence_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attack-occurrence-limit"] = t
		}
	}

	if v, ok := d.GetOk("attack_occurrence_within"); ok {
		t, err := expandSecurityWafThresholdBasedDetection(d, v, "attack_occurrence_within", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attack-occurrence-within"] = t
		}
	}

	return &obj, nil
}
