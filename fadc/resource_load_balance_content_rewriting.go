// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content rewriting.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceContentRewriting() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceContentRewritingRead,
		Update: resourceLoadBalanceContentRewritingUpdate,
		Create: resourceLoadBalanceContentRewritingCreate,
		Delete: resourceLoadBalanceContentRewritingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"referer_content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"referer_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"header_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"location": &schema.Schema{
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
func resourceLoadBalanceContentRewritingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRewriting: type error")
	}

	obj, err := getObjectLoadBalanceContentRewriting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRewriting resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceContentRewriting(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRewriting resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceContentRewritingRead(d, m)
}
func resourceLoadBalanceContentRewritingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceContentRewriting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRewriting resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceContentRewriting(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRewriting resource: %v", err)
	}

	return resourceLoadBalanceContentRewritingRead(d, m)
}
func resourceLoadBalanceContentRewritingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceContentRewriting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceContentRewriting resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceContentRewritingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceContentRewriting(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRewriting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceContentRewriting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRewriting resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceContentRewritingRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingRefererContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingRefererStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingHostStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingUrlStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingActionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingHeaderValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingUrlContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingHostContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceContentRewriting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("redirect", flattenLoadBalanceContentRewritingRedirect(o["redirect"], d, "redirect", sv)); err != nil {
		if !fortiAPIPatch(o["redirect"]) {
			return fmt.Errorf("Error reading redirect: %v", err)
		}
	}

	if err = d.Set("header_name", flattenLoadBalanceContentRewritingHeaderName(o["header_name"], d, "header_name", sv)); err != nil {
		if !fortiAPIPatch(o["header_name"]) {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("referer_content", flattenLoadBalanceContentRewritingRefererContent(o["referer_content"], d, "referer_content", sv)); err != nil {
		if !fortiAPIPatch(o["referer_content"]) {
			return fmt.Errorf("Error reading referer_content: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceContentRewritingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("referer_status", flattenLoadBalanceContentRewritingRefererStatus(o["referer_status"], d, "referer_status", sv)); err != nil {
		if !fortiAPIPatch(o["referer_status"]) {
			return fmt.Errorf("Error reading referer_status: %v", err)
		}
	}

	if err = d.Set("comments", flattenLoadBalanceContentRewritingComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("host_status", flattenLoadBalanceContentRewritingHostStatus(o["host_status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host_status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}

	if err = d.Set("url_status", flattenLoadBalanceContentRewritingUrlStatus(o["url_status"], d, "url_status", sv)); err != nil {
		if !fortiAPIPatch(o["url_status"]) {
			return fmt.Errorf("Error reading url_status: %v", err)
		}
	}

	if err = d.Set("action_type", flattenLoadBalanceContentRewritingActionType(o["action_type"], d, "action_type", sv)); err != nil {
		if !fortiAPIPatch(o["action_type"]) {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if err = d.Set("header_value", flattenLoadBalanceContentRewritingHeaderValue(o["header_value"], d, "header_value", sv)); err != nil {
		if !fortiAPIPatch(o["header_value"]) {
			return fmt.Errorf("Error reading header_value: %v", err)
		}
	}

	if err = d.Set("action", flattenLoadBalanceContentRewritingAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("url_content", flattenLoadBalanceContentRewritingUrlContent(o["url_content"], d, "url_content", sv)); err != nil {
		if !fortiAPIPatch(o["url_content"]) {
			return fmt.Errorf("Error reading url_content: %v", err)
		}
	}

	if err = d.Set("host_content", flattenLoadBalanceContentRewritingHostContent(o["host_content"], d, "host_content", sv)); err != nil {
		if !fortiAPIPatch(o["host_content"]) {
			return fmt.Errorf("Error reading host_content: %v", err)
		}
	}

	if err = d.Set("location", flattenLoadBalanceContentRewritingLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceContentRewritingRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingRefererContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingRefererStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingHostStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingUrlStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingActionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingHeaderValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingUrlContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingHostContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceContentRewriting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("redirect"); ok {
		t, err := expandLoadBalanceContentRewritingRedirect(d, v, "redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok {
		t, err := expandLoadBalanceContentRewritingHeaderName(d, v, "header_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header_name"] = t
		}
	}

	if v, ok := d.GetOk("referer_content"); ok {
		t, err := expandLoadBalanceContentRewritingRefererContent(d, v, "referer_content", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referer_content"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceContentRewritingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("referer_status"); ok {
		t, err := expandLoadBalanceContentRewritingRefererStatus(d, v, "referer_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referer_status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandLoadBalanceContentRewritingComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandLoadBalanceContentRewritingHostStatus(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_status"] = t
		}
	}

	if v, ok := d.GetOk("url_status"); ok {
		t, err := expandLoadBalanceContentRewritingUrlStatus(d, v, "url_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url_status"] = t
		}
	}

	if v, ok := d.GetOk("action_type"); ok {
		t, err := expandLoadBalanceContentRewritingActionType(d, v, "action_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action_type"] = t
		}
	}

	if v, ok := d.GetOk("header_value"); ok {
		t, err := expandLoadBalanceContentRewritingHeaderValue(d, v, "header_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header_value"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandLoadBalanceContentRewritingAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("url_content"); ok {
		t, err := expandLoadBalanceContentRewritingUrlContent(d, v, "url_content", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url_content"] = t
		}
	}

	if v, ok := d.GetOk("host_content"); ok {
		t, err := expandLoadBalanceContentRewritingHostContent(d, v, "host_content", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_content"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {
		t, err := expandLoadBalanceContentRewritingLocation(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	return &obj, nil
}
