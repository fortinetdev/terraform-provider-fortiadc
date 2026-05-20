// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure security waf adpative learning child url list

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityWafAdaptiveLearningChildUrlList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityWafAdaptiveLearningChildUrlListRead,
		Update: resourceSecurityWafAdaptiveLearningChildUrlListUpdate,
		Create: resourceSecurityWafAdaptiveLearningChildUrlListCreate,
		Delete: resourceSecurityWafAdaptiveLearningChildUrlListDelete,

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
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url": &schema.Schema{
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

func resourceSecurityWafAdaptiveLearningChildUrlListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	obj, err := getObjectSecurityWafAdaptiveLearningChildUrlList(d, c.Fv, false)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdaptiveLearningChildUrlList resource while getting object: %v", err)
	}

	path := "/api/security_waf_adaptive_learning_child_url_list?pkey=" + escapeURLString(pkey)
	_, err = c.StandardCreate(obj, vdom, path)
	if err != nil {
		return fmt.Errorf("Error creating SecurityWafAdaptiveLearningChildUrlList resource: %v", err)
	}

	id := "SecurityWafAdaptiveLearningChildUrlList"

	d.SetId(id)

	return resourceSecurityWafAdaptiveLearningChildUrlListRead(d, m)
}

func resourceSecurityWafAdaptiveLearningChildUrlListUpdate(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

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
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	obj, err := getObjectSecurityWafAdaptiveLearningChildUrlList(d, c.Fv, true)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAdaptiveLearning resource while getting object: %v", err)
	}

	path := "/api/security_waf_adaptive_learning_child_url_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	_, err = c.StandardUpdate(obj, mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error updating SecurityWafAdaptiveLearning resource: %v", err)
	}

	return resourceSecurityWafAdaptiveLearningChildUrlListRead(d, m)
}

func resourceSecurityWafAdaptiveLearningChildUrlListDelete(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

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
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	mkey := ""

	t = d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	path := "/api/security_waf_adaptive_learning_child_url_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	err := c.StandardDelete(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error clearing SecurityWafAdaptiveLearningChildUrlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSecurityWafAdaptiveLearningChildUrlListRead(d *schema.ResourceData, m interface{}) error {
	//mkey := d.Id()

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
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	pkey := ""

	t = d.Get("pkey")
	if v, ok := t.(string); ok {
		pkey = v
	} else if v, ok := t.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SecurityWafAdaptiveLearningChildUrlList: type error")
	}

	path := "/api/security_waf_adaptive_learning_child_url_list?pkey=" + escapeURLString(pkey)
	path += "&mkey="
	path += escapeURLString(mkey)
	o, err := c.StandardRead(mkey, vdom, path)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdaptiveLearningChildUrlList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityWafAdaptiveLearningChildUrlList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityWafAdaptiveLearningChildUrlList resource from API: %v", err)
	}
	return nil
}

func flattenSecurityWafAdaptiveLearningChildUrlList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityWafAdaptiveLearningChildUrlList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSecurityWafAdaptiveLearningChildUrlList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("host_status", flattenSecurityWafAdaptiveLearningChildUrlList(o["host_status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host_status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}

	if err = d.Set("host", flattenSecurityWafAdaptiveLearningChildUrlList(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("url", flattenSecurityWafAdaptiveLearningChildUrlList(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func expandSecurityWafAdaptiveLearningChildUrlList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityWafAdaptiveLearningChildUrlList(d *schema.ResourceData, sv string, action bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if action == true {
		if v, ok := d.GetOk("mkey"); ok {
			t, err := expandSecurityWafAdaptiveLearningChildUrlList(d, v, "mkey", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mkey"] = t
			}
		}
	}

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandSecurityWafAdaptiveLearningChildUrlList(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_status"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandSecurityWafAdaptiveLearningChildUrlList(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {
		t, err := expandSecurityWafAdaptiveLearningChildUrlList(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
