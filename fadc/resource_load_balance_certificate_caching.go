// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance certificate caching.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceCertificateCaching() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceCertificateCachingRead,
		Update: resourceLoadBalanceCertificateCachingUpdate,
		Create: resourceLoadBalanceCertificateCachingCreate,
		Delete: resourceLoadBalanceCertificateCachingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_entries": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_certificate_cache_size": &schema.Schema{
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
func resourceLoadBalanceCertificateCachingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCertificateCaching: type error")
	}

	obj, err := getObjectLoadBalanceCertificateCaching(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCertificateCaching resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceCertificateCaching(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCertificateCaching resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalanceCertificateCachingRead(d, m)
}
func resourceLoadBalanceCertificateCachingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalanceCertificateCaching(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCertificateCaching resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceCertificateCaching(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCertificateCaching resource: %v", err)
	}

	return resourceLoadBalanceCertificateCachingRead(d, m)
}
func resourceLoadBalanceCertificateCachingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalanceCertificateCaching(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceCertificateCaching resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceCertificateCachingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalanceCertificateCaching(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCertificateCaching resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceCertificateCaching(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCertificateCaching resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceCertificateCachingMaxEntries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCertificateCachingMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCertificateCachingMaxCertificateCacheSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceCertificateCaching(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_entries", flattenLoadBalanceCertificateCachingMaxEntries(o["max_entries"], d, "max_entries", sv)); err != nil {
		if !fortiAPIPatch(o["max_entries"]) {
			return fmt.Errorf("Error reading max_entries: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceCertificateCachingMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("max_certificate_cache_size", flattenLoadBalanceCertificateCachingMaxCertificateCacheSize(o["max_certificate_cache_size"], d, "max_certificate_cache_size", sv)); err != nil {
		if !fortiAPIPatch(o["max_certificate_cache_size"]) {
			return fmt.Errorf("Error reading max_certificate_cache_size: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceCertificateCachingMaxEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCertificateCachingMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCertificateCachingMaxCertificateCacheSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceCertificateCaching(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_entries"); ok {
		t, err := expandLoadBalanceCertificateCachingMaxEntries(d, v, "max_entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_entries"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceCertificateCachingMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("max_certificate_cache_size"); ok {
		t, err := expandLoadBalanceCertificateCachingMaxCertificateCacheSize(d, v, "max_certificate_cache_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max_certificate_cache_size"] = t
		}
	}

	return &obj, nil
}
