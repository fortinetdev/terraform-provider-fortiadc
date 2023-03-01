// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content rewriting child match condition.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceContentRewritingChildMatchCondition() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceContentRewritingChildMatchConditionRead,
		Update: resourceLoadBalanceContentRewritingChildMatchConditionUpdate,
		Create: resourceLoadBalanceContentRewritingChildMatchConditionCreate,
		Delete: resourceLoadBalanceContentRewritingChildMatchConditionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ignorecase": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reverse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceLoadBalanceContentRewritingChildMatchConditionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRewritingChildMatchCondition: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceContentRewritingChildMatchCondition: type error")
	}

	obj, err := getObjectLoadBalanceContentRewritingChildMatchCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRewritingChildMatchCondition resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceContentRewritingChildMatchCondition(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRewritingChildMatchCondition resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceContentRewritingChildMatchConditionRead(d, m)
}
func resourceLoadBalanceContentRewritingChildMatchConditionUpdate(d *schema.ResourceData, m interface{}) error {
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

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceContentRewritingChildMatchCondition: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceContentRewritingChildMatchCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRewritingChildMatchCondition resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceContentRewritingChildMatchCondition(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRewritingChildMatchCondition resource: %v", err)
	}

	return resourceLoadBalanceContentRewritingChildMatchConditionRead(d, m)
}
func resourceLoadBalanceContentRewritingChildMatchConditionDelete(d *schema.ResourceData, m interface{}) error {
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

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceContentRewritingChildMatchCondition: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceContentRewritingChildMatchCondition(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceContentRewritingChildMatchCondition resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceContentRewritingChildMatchConditionRead(d *schema.ResourceData, m interface{}) error {
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

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceContentRewritingChildMatchCondition: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceContentRewritingChildMatchCondition(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRewritingChildMatchCondition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceContentRewritingChildMatchCondition(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRewritingChildMatchCondition resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceContentRewritingChildMatchConditionIgnorecase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingChildMatchConditionReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingChildMatchConditionObject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingChildMatchConditionContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingChildMatchConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRewritingChildMatchConditionMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceContentRewritingChildMatchCondition(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ignorecase", flattenLoadBalanceContentRewritingChildMatchConditionIgnorecase(o["ignorecase"], d, "ignorecase", sv)); err != nil {
		if !fortiAPIPatch(o["ignorecase"]) {
			return fmt.Errorf("Error reading ignorecase: %v", err)
		}
	}

	if err = d.Set("reverse", flattenLoadBalanceContentRewritingChildMatchConditionReverse(o["reverse"], d, "reverse", sv)); err != nil {
		if !fortiAPIPatch(o["reverse"]) {
			return fmt.Errorf("Error reading reverse: %v", err)
		}
	}

	if err = d.Set("object", flattenLoadBalanceContentRewritingChildMatchConditionObject(o["object"], d, "object", sv)); err != nil {
		if !fortiAPIPatch(o["object"]) {
			return fmt.Errorf("Error reading object: %v", err)
		}
	}

	if err = d.Set("content", flattenLoadBalanceContentRewritingChildMatchConditionContent(o["content"], d, "content", sv)); err != nil {
		if !fortiAPIPatch(o["content"]) {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceContentRewritingChildMatchConditionType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceContentRewritingChildMatchConditionMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceContentRewritingChildMatchConditionIgnorecase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingChildMatchConditionReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingChildMatchConditionObject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingChildMatchConditionContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingChildMatchConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRewritingChildMatchConditionMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceContentRewritingChildMatchCondition(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ignorecase"); ok {
		t, err := expandLoadBalanceContentRewritingChildMatchConditionIgnorecase(d, v, "ignorecase", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignorecase"] = t
		}
	}

	if v, ok := d.GetOk("reverse"); ok {
		t, err := expandLoadBalanceContentRewritingChildMatchConditionReverse(d, v, "reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reverse"] = t
		}
	}

	if v, ok := d.GetOk("object"); ok {
		t, err := expandLoadBalanceContentRewritingChildMatchConditionObject(d, v, "object", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object"] = t
		}
	}

	if v, ok := d.GetOk("content"); ok {
		t, err := expandLoadBalanceContentRewritingChildMatchConditionContent(d, v, "content", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceContentRewritingChildMatchConditionType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceContentRewritingChildMatchConditionMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
