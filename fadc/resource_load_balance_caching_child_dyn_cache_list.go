// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance caching child dyn cache list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceCachingChildDynCacheList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceCachingChildDynCacheListRead,
		Update: resourceLoadBalanceCachingChildDynCacheListUpdate,
		Create: resourceLoadBalanceCachingChildDynCacheListCreate,
		Delete: resourceLoadBalanceCachingChildDynCacheListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"invalid_uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"uri": &schema.Schema{
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
func resourceLoadBalanceCachingChildDynCacheListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	obj, err := getObjectLoadBalanceCachingChildDynCacheList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCachingChildDynCacheList resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceCachingChildDynCacheList(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceCachingChildDynCacheList resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceCachingChildDynCacheListRead(d, m)
}
func resourceLoadBalanceCachingChildDynCacheListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceCachingChildDynCacheList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCachingChildDynCacheList resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceCachingChildDynCacheList(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceCachingChildDynCacheList resource: %v", err)
	}

	return resourceLoadBalanceCachingChildDynCacheListRead(d, m)
}
func resourceLoadBalanceCachingChildDynCacheListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceCachingChildDynCacheList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceCachingChildDynCacheList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceCachingChildDynCacheListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceCachingChildDynCacheList: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceCachingChildDynCacheList(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCachingChildDynCacheList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceCachingChildDynCacheList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceCachingChildDynCacheList resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceCachingChildDynCacheListInvalidUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingChildDynCacheListAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingChildDynCacheListUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceCachingChildDynCacheListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceCachingChildDynCacheList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("invalid_uri", flattenLoadBalanceCachingChildDynCacheListInvalidUri(o["invalid_uri"], d, "invalid_uri", sv)); err != nil {
		if !fortiAPIPatch(o["invalid_uri"]) {
			return fmt.Errorf("Error reading invalid_uri: %v", err)
		}
	}

	if err = d.Set("age", flattenLoadBalanceCachingChildDynCacheListAge(o["age"], d, "age", sv)); err != nil {
		if !fortiAPIPatch(o["age"]) {
			return fmt.Errorf("Error reading age: %v", err)
		}
	}

	if err = d.Set("uri", flattenLoadBalanceCachingChildDynCacheListUri(o["uri"], d, "uri", sv)); err != nil {
		if !fortiAPIPatch(o["uri"]) {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceCachingChildDynCacheListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceCachingChildDynCacheListInvalidUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingChildDynCacheListAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingChildDynCacheListUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceCachingChildDynCacheListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceCachingChildDynCacheList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("invalid_uri"); ok {
		t, err := expandLoadBalanceCachingChildDynCacheListInvalidUri(d, v, "invalid_uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid_uri"] = t
		}
	}

	if v, ok := d.GetOk("age"); ok {
		t, err := expandLoadBalanceCachingChildDynCacheListAge(d, v, "age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["age"] = t
		}
	}

	if v, ok := d.GetOk("uri"); ok {
		t, err := expandLoadBalanceCachingChildDynCacheListUri(d, v, "uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceCachingChildDynCacheListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
