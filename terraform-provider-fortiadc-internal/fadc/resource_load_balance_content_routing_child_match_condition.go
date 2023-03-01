// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance content routing child match condition.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceContentRoutingChildMatchCondition() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceContentRoutingChildMatchConditionRead,
		Update: resourceLoadBalanceContentRoutingChildMatchConditionUpdate,
		Create: resourceLoadBalanceContentRoutingChildMatchConditionCreate,
		Delete: resourceLoadBalanceContentRoutingChildMatchConditionDelete,

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
func resourceLoadBalanceContentRoutingChildMatchConditionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	obj, err := getObjectLoadBalanceContentRoutingChildMatchCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRoutingChildMatchCondition resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceContentRoutingChildMatchCondition(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceContentRoutingChildMatchCondition resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceContentRoutingChildMatchConditionRead(d, m)
}
func resourceLoadBalanceContentRoutingChildMatchConditionUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceContentRoutingChildMatchCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRoutingChildMatchCondition resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceContentRoutingChildMatchCondition(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceContentRoutingChildMatchCondition resource: %v", err)
	}

	return resourceLoadBalanceContentRoutingChildMatchConditionRead(d, m)
}
func resourceLoadBalanceContentRoutingChildMatchConditionDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceContentRoutingChildMatchCondition(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceContentRoutingChildMatchCondition resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceContentRoutingChildMatchConditionRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceContentRoutingChildMatchCondition: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceContentRoutingChildMatchCondition(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRoutingChildMatchCondition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceContentRoutingChildMatchCondition(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceContentRoutingChildMatchCondition resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceContentRoutingChildMatchConditionIgnorecase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingChildMatchConditionReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingChildMatchConditionObject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingChildMatchConditionContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingChildMatchConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceContentRoutingChildMatchConditionMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceContentRoutingChildMatchCondition(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ignorecase", flattenLoadBalanceContentRoutingChildMatchConditionIgnorecase(o["ignorecase"], d, "ignorecase", sv)); err != nil {
		if !fortiAPIPatch(o["ignorecase"]) {
			return fmt.Errorf("Error reading ignorecase: %v", err)
		}
	}

	if err = d.Set("reverse", flattenLoadBalanceContentRoutingChildMatchConditionReverse(o["reverse"], d, "reverse", sv)); err != nil {
		if !fortiAPIPatch(o["reverse"]) {
			return fmt.Errorf("Error reading reverse: %v", err)
		}
	}

	if err = d.Set("object", flattenLoadBalanceContentRoutingChildMatchConditionObject(o["object"], d, "object", sv)); err != nil {
		if !fortiAPIPatch(o["object"]) {
			return fmt.Errorf("Error reading object: %v", err)
		}
	}

	if err = d.Set("content", flattenLoadBalanceContentRoutingChildMatchConditionContent(o["content"], d, "content", sv)); err != nil {
		if !fortiAPIPatch(o["content"]) {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalanceContentRoutingChildMatchConditionType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceContentRoutingChildMatchConditionMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceContentRoutingChildMatchConditionIgnorecase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingChildMatchConditionReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingChildMatchConditionObject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingChildMatchConditionContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingChildMatchConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceContentRoutingChildMatchConditionMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceContentRoutingChildMatchCondition(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ignorecase"); ok {
		t, err := expandLoadBalanceContentRoutingChildMatchConditionIgnorecase(d, v, "ignorecase", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignorecase"] = t
		}
	}

	if v, ok := d.GetOk("reverse"); ok {
		t, err := expandLoadBalanceContentRoutingChildMatchConditionReverse(d, v, "reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reverse"] = t
		}
	}

	if v, ok := d.GetOk("object"); ok {
		t, err := expandLoadBalanceContentRoutingChildMatchConditionObject(d, v, "object", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object"] = t
		}
	}

	if v, ok := d.GetOk("content"); ok {
		t, err := expandLoadBalanceContentRoutingChildMatchConditionContent(d, v, "content", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalanceContentRoutingChildMatchConditionType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceContentRoutingChildMatchConditionMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
