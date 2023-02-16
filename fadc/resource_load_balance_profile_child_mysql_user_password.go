// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance profile child mysql user password.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalanceProfileChildMysqlUserPassword() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalanceProfileChildMysqlUserPasswordRead,
		Update: resourceLoadBalanceProfileChildMysqlUserPasswordUpdate,
		Create: resourceLoadBalanceProfileChildMysqlUserPasswordCreate,
		Delete: resourceLoadBalanceProfileChildMysqlUserPasswordDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
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
func resourceLoadBalanceProfileChildMysqlUserPasswordCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlUserPassword: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlUserPassword: type error")
	}

	obj, err := getObjectLoadBalanceProfileChildMysqlUserPassword(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildMysqlUserPassword resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateLoadBalanceProfileChildMysqlUserPassword(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceProfileChildMysqlUserPassword resource: %v", err)
	}

	d.SetId(new_id)

	return resourceLoadBalanceProfileChildMysqlUserPasswordRead(d, m)
}
func resourceLoadBalanceProfileChildMysqlUserPasswordUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlUserPassword: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectLoadBalanceProfileChildMysqlUserPassword(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildMysqlUserPassword resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceProfileChildMysqlUserPassword(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceProfileChildMysqlUserPassword resource: %v", err)
	}

	return resourceLoadBalanceProfileChildMysqlUserPasswordRead(d, m)
}
func resourceLoadBalanceProfileChildMysqlUserPasswordDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlUserPassword: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteLoadBalanceProfileChildMysqlUserPassword(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceProfileChildMysqlUserPassword resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalanceProfileChildMysqlUserPasswordRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceProfileChildMysqlUserPassword: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadLoadBalanceProfileChildMysqlUserPassword(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildMysqlUserPassword resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceProfileChildMysqlUserPassword(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceProfileChildMysqlUserPassword resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalanceProfileChildMysqlUserPasswordUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlUserPasswordPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalanceProfileChildMysqlUserPasswordMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalanceProfileChildMysqlUserPassword(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("username", flattenLoadBalanceProfileChildMysqlUserPasswordUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("password", flattenLoadBalanceProfileChildMysqlUserPasswordPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalanceProfileChildMysqlUserPasswordMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLoadBalanceProfileChildMysqlUserPasswordUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlUserPasswordPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceProfileChildMysqlUserPasswordMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceProfileChildMysqlUserPassword(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("username"); ok {
		t, err := expandLoadBalanceProfileChildMysqlUserPasswordUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandLoadBalanceProfileChildMysqlUserPasswordPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalanceProfileChildMysqlUserPasswordMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
