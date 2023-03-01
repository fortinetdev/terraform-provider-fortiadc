// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance pagespeed child resource control.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePagespeedChildResourceControl() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePagespeedChildResourceControlRead,
		Update: resourceLoadBalancePagespeedChildResourceControlUpdate,
		Create: resourceLoadBalancePagespeedChildResourceControlCreate,
		Delete: resourceLoadBalancePagespeedChildResourceControlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"rewrite_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fetch_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"origin_domain_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
func resourceLoadBalancePagespeedChildResourceControlCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	obj, err := getObjectLoadBalancePagespeedChildResourceControl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePagespeedChildResourceControl resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalancePagespeedChildResourceControl(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePagespeedChildResourceControl resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalancePagespeedChildResourceControlRead(d, m)
}
func resourceLoadBalancePagespeedChildResourceControlUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalancePagespeedChildResourceControl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePagespeedChildResourceControl resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePagespeedChildResourceControl(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePagespeedChildResourceControl resource: %v", err)
	}

	return resourceLoadBalancePagespeedChildResourceControlRead(d, m)
}
func resourceLoadBalancePagespeedChildResourceControlDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalancePagespeedChildResourceControl(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePagespeedChildResourceControl resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePagespeedChildResourceControlRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePagespeedChildResourceControl: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalancePagespeedChildResourceControl(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePagespeedChildResourceControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePagespeedChildResourceControl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePagespeedChildResourceControl resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePagespeedChildResourceControlRewriteDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedChildResourceControlFetchDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedChildResourceControlMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePagespeedChildResourceControlOriginDomainPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePagespeedChildResourceControl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("rewrite_domain", flattenLoadBalancePagespeedChildResourceControlRewriteDomain(o["rewrite_domain"], d, "rewrite_domain", sv)); err != nil {
		if !fortiAPIPatch(o["rewrite_domain"]) {
			return fmt.Errorf("Error reading rewrite_domain: %v", err)
		}
	}

	if err = d.Set("fetch_domain", flattenLoadBalancePagespeedChildResourceControlFetchDomain(o["fetch_domain"], d, "fetch_domain", sv)); err != nil {
		if !fortiAPIPatch(o["fetch_domain"]) {
			return fmt.Errorf("Error reading fetch_domain: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePagespeedChildResourceControlMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("origin_domain_pattern", flattenLoadBalancePagespeedChildResourceControlOriginDomainPattern(o["origin_domain_pattern"], d, "origin_domain_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["origin_domain_pattern"]) {
			return fmt.Errorf("Error reading origin_domain_pattern: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePagespeedChildResourceControlRewriteDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedChildResourceControlFetchDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedChildResourceControlMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePagespeedChildResourceControlOriginDomainPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePagespeedChildResourceControl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rewrite_domain"); ok {
		t, err := expandLoadBalancePagespeedChildResourceControlRewriteDomain(d, v, "rewrite_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rewrite_domain"] = t
		}
	}

	if v, ok := d.GetOk("fetch_domain"); ok {
		t, err := expandLoadBalancePagespeedChildResourceControlFetchDomain(d, v, "fetch_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fetch_domain"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePagespeedChildResourceControlMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("origin_domain_pattern"); ok {
		t, err := expandLoadBalancePagespeedChildResourceControlOriginDomainPattern(d, v, "origin_domain_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin_domain_pattern"] = t
		}
	}

	return &obj, nil
}
