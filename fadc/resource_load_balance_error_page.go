// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance error page.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceErrorPage() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceErrorPageRead,
		Update: resourceLoadBalanceErrorPageUpdate,
		Create: resourceLoadBalanceErrorPageCreate,
		Delete: resourceLoadBalanceErrorPageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vpath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file": &schema.Schema{
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
func resourceLoadBalanceErrorPageCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceErrorPage: type error")
	}

	obj, err := getObjectLoadBalanceErrorPage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceErrorPage resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceErrorPage(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceErrorPage resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceErrorPageRead(d, m)
}
func resourceLoadBalanceErrorPageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceErrorPage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceErrorPage resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceErrorPage(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceErrorPage resource: %v", err)
	}

	return resourceLoadBalanceErrorPageRead(d, m)
}
func resourceLoadBalanceErrorPageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceErrorPage(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceErrorPage resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceErrorPageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceErrorPage(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceErrorPage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceErrorPage(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceErrorPage resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceErrorPageMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceErrorPageVpath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceErrorPageFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceErrorPage(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenLoadBalanceErrorPageMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("vpath", flattenLoadBalanceErrorPageVpath(o["vpath"], d, "vpath", sv)); err != nil {
		if !fortiAPIPatch(o["vpath"]) {
			return fmt.Errorf("Error reading vpath: %v", err)
		}
	}

	if err = d.Set("file", flattenLoadBalanceErrorPageFile(o["file"], d, "file", sv)); err != nil {
		if !fortiAPIPatch(o["file"]) {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceErrorPageMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceErrorPageVpath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceErrorPageFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceErrorPage(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceErrorPageMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("vpath"); ok {
		t, err := expandLoadBalanceErrorPageVpath(d, v, "vpath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpath"] = t
		}
	}

	if v, ok := d.GetOk("file"); ok {
		t, err := expandLoadBalanceErrorPageFile(d, v, "file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file"] = t
		}
	}

	return &obj, nil
}
