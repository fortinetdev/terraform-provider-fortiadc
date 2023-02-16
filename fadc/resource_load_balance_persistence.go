// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance persistence.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoadBalancePersistence() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLoadBalancePersistenceRead,
		Update: resourceLoadBalancePersistenceUpdate,
		Create: resourceLoadBalancePersistenceCreate,
		Delete: resourceLoadBalancePersistenceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"match_across_servers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"keyvalue_relation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_secure": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_httponly": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sess_kw_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"avpno": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipv6_maskbits": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_custom_attr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"iso8583_bitmap_relation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipv4_maskbits": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keyword": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_custom_attr_val": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"override_connection_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_samesite": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"radius_attribute_relation": &schema.Schema{
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
func resourceLoadBalancePersistenceCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalancePersistence: type error")
	}

	obj, err := getObjectLoadBalancePersistence(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePersistence resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalancePersistence(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalancePersistence resource: %v", err)
	}

	d.SetId(mkey)

	return resourceLoadBalancePersistenceRead(d, m)
}
func resourceLoadBalancePersistenceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectLoadBalancePersistence(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePersistence resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalancePersistence(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalancePersistence resource: %v", err)
	}

	return resourceLoadBalancePersistenceRead(d, m)
}
func resourceLoadBalancePersistenceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteLoadBalancePersistence(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalancePersistence resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLoadBalancePersistenceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadLoadBalancePersistence(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePersistence resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalancePersistence(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalancePersistence resource from API: %v", err)
	}
	return nil
}
func flattenLoadBalancePersistenceMatchAcrossServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceKeyvalueRelation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceCookieHttponly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceSessKwType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceAvpno(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceIpv6Maskbits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceCookieCustomAttr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceIso8583BitmapRelation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceIpv4Maskbits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceKeyword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceCookieCustomAttrVal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceOverrideConnectionLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceCookieSamesite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLoadBalancePersistenceRadiusAttributeRelation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLoadBalancePersistence(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("match_across_servers", flattenLoadBalancePersistenceMatchAcrossServers(o["match-across-servers"], d, "match_across_servers", sv)); err != nil {
		if !fortiAPIPatch(o["match-across-servers"]) {
			return fmt.Errorf("Error reading match-across-servers: %v", err)
		}
	}

	if err = d.Set("keyvalue_relation", flattenLoadBalancePersistenceKeyvalueRelation(o["keyvalue-relation"], d, "keyvalue_relation", sv)); err != nil {
		if !fortiAPIPatch(o["keyvalue-relation"]) {
			return fmt.Errorf("Error reading keyvalue-relation: %v", err)
		}
	}

	if err = d.Set("cookie_secure", flattenLoadBalancePersistenceCookieSecure(o["cookie_secure"], d, "cookie_secure", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_secure"]) {
			return fmt.Errorf("Error reading cookie_secure: %v", err)
		}
	}

	if err = d.Set("cookie_httponly", flattenLoadBalancePersistenceCookieHttponly(o["cookie_httponly"], d, "cookie_httponly", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_httponly"]) {
			return fmt.Errorf("Error reading cookie_httponly: %v", err)
		}
	}

	if err = d.Set("sess_kw_type", flattenLoadBalancePersistenceSessKwType(o["sess_kw_type"], d, "sess_kw_type", sv)); err != nil {
		if !fortiAPIPatch(o["sess_kw_type"]) {
			return fmt.Errorf("Error reading sess_kw_type: %v", err)
		}
	}

	if err = d.Set("avpno", flattenLoadBalancePersistenceAvpno(o["avpno"], d, "avpno", sv)); err != nil {
		if !fortiAPIPatch(o["avpno"]) {
			return fmt.Errorf("Error reading avpno: %v", err)
		}
	}

	if err = d.Set("ipv6_maskbits", flattenLoadBalancePersistenceIpv6Maskbits(o["ipv6-maskbits"], d, "ipv6_maskbits", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-maskbits"]) {
			return fmt.Errorf("Error reading ipv6-maskbits: %v", err)
		}
	}

	if err = d.Set("cookie_custom_attr", flattenLoadBalancePersistenceCookieCustomAttr(o["cookie_custom_attr"], d, "cookie_custom_attr", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_custom_attr"]) {
			return fmt.Errorf("Error reading cookie_custom_attr: %v", err)
		}
	}

	if err = d.Set("type", flattenLoadBalancePersistenceType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("iso8583_bitmap_relation", flattenLoadBalancePersistenceIso8583BitmapRelation(o["iso8583_bitmap-relation"], d, "iso8583_bitmap_relation", sv)); err != nil {
		if !fortiAPIPatch(o["iso8583_bitmap-relation"]) {
			return fmt.Errorf("Error reading iso8583_bitmap-relation: %v", err)
		}
	}

	if err = d.Set("ipv4_maskbits", flattenLoadBalancePersistenceIpv4Maskbits(o["ipv4-maskbits"], d, "ipv4_maskbits", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-maskbits"]) {
			return fmt.Errorf("Error reading ipv4-maskbits: %v", err)
		}
	}

	if err = d.Set("mkey", flattenLoadBalancePersistenceMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("keyword", flattenLoadBalancePersistenceKeyword(o["keyword"], d, "keyword", sv)); err != nil {
		if !fortiAPIPatch(o["keyword"]) {
			return fmt.Errorf("Error reading keyword: %v", err)
		}
	}

	if err = d.Set("cookie_domain", flattenLoadBalancePersistenceCookieDomain(o["cookie_domain"], d, "cookie_domain", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_domain"]) {
			return fmt.Errorf("Error reading cookie_domain: %v", err)
		}
	}

	if err = d.Set("cookie_custom_attr_val", flattenLoadBalancePersistenceCookieCustomAttrVal(o["cookie_custom_attr_val"], d, "cookie_custom_attr_val", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_custom_attr_val"]) {
			return fmt.Errorf("Error reading cookie_custom_attr_val: %v", err)
		}
	}

	if err = d.Set("override_connection_limit", flattenLoadBalancePersistenceOverrideConnectionLimit(o["override-connection-limit"], d, "override_connection_limit", sv)); err != nil {
		if !fortiAPIPatch(o["override-connection-limit"]) {
			return fmt.Errorf("Error reading override-connection-limit: %v", err)
		}
	}

	if err = d.Set("cookie_samesite", flattenLoadBalancePersistenceCookieSamesite(o["cookie_samesite"], d, "cookie_samesite", sv)); err != nil {
		if !fortiAPIPatch(o["cookie_samesite"]) {
			return fmt.Errorf("Error reading cookie_samesite: %v", err)
		}
	}

	if err = d.Set("timeout", flattenLoadBalancePersistenceTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("radius_attribute_relation", flattenLoadBalancePersistenceRadiusAttributeRelation(o["radius-attribute-relation"], d, "radius_attribute_relation", sv)); err != nil {
		if !fortiAPIPatch(o["radius-attribute-relation"]) {
			return fmt.Errorf("Error reading radius-attribute-relation: %v", err)
		}
	}

	return nil
}

func expandLoadBalancePersistenceMatchAcrossServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceKeyvalueRelation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceCookieHttponly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceSessKwType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceAvpno(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceIpv6Maskbits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceCookieCustomAttr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceIso8583BitmapRelation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceIpv4Maskbits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceKeyword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceCookieCustomAttrVal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceOverrideConnectionLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceCookieSamesite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLoadBalancePersistenceRadiusAttributeRelation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalancePersistence(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("match_across_servers"); ok {
		t, err := expandLoadBalancePersistenceMatchAcrossServers(d, v, "match_across_servers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-across-servers"] = t
		}
	}

	if v, ok := d.GetOk("keyvalue_relation"); ok {
		t, err := expandLoadBalancePersistenceKeyvalueRelation(d, v, "keyvalue_relation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyvalue-relation"] = t
		}
	}

	if v, ok := d.GetOk("cookie_secure"); ok {
		t, err := expandLoadBalancePersistenceCookieSecure(d, v, "cookie_secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_secure"] = t
		}
	}

	if v, ok := d.GetOk("cookie_httponly"); ok {
		t, err := expandLoadBalancePersistenceCookieHttponly(d, v, "cookie_httponly", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_httponly"] = t
		}
	}

	if v, ok := d.GetOk("sess_kw_type"); ok {
		t, err := expandLoadBalancePersistenceSessKwType(d, v, "sess_kw_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sess_kw_type"] = t
		}
	}

	if v, ok := d.GetOk("avpno"); ok {
		t, err := expandLoadBalancePersistenceAvpno(d, v, "avpno", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["avpno"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_maskbits"); ok {
		t, err := expandLoadBalancePersistenceIpv6Maskbits(d, v, "ipv6_maskbits", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-maskbits"] = t
		}
	}

	if v, ok := d.GetOk("cookie_custom_attr"); ok {
		t, err := expandLoadBalancePersistenceCookieCustomAttr(d, v, "cookie_custom_attr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_custom_attr"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLoadBalancePersistenceType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("iso8583_bitmap_relation"); ok {
		t, err := expandLoadBalancePersistenceIso8583BitmapRelation(d, v, "iso8583_bitmap_relation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iso8583_bitmap-relation"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_maskbits"); ok {
		t, err := expandLoadBalancePersistenceIpv4Maskbits(d, v, "ipv4_maskbits", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-maskbits"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLoadBalancePersistenceMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("keyword"); ok {
		t, err := expandLoadBalancePersistenceKeyword(d, v, "keyword", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyword"] = t
		}
	}

	if v, ok := d.GetOk("cookie_domain"); ok {
		t, err := expandLoadBalancePersistenceCookieDomain(d, v, "cookie_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_domain"] = t
		}
	}

	if v, ok := d.GetOk("cookie_custom_attr_val"); ok {
		t, err := expandLoadBalancePersistenceCookieCustomAttrVal(d, v, "cookie_custom_attr_val", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_custom_attr_val"] = t
		}
	}

	if v, ok := d.GetOk("override_connection_limit"); ok {
		t, err := expandLoadBalancePersistenceOverrideConnectionLimit(d, v, "override_connection_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-connection-limit"] = t
		}
	}

	if v, ok := d.GetOk("cookie_samesite"); ok {
		t, err := expandLoadBalancePersistenceCookieSamesite(d, v, "cookie_samesite", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie_samesite"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandLoadBalancePersistenceTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("radius_attribute_relation"); ok {
		t, err := expandLoadBalancePersistenceRadiusAttributeRelation(d, v, "radius_attribute_relation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-attribute-relation"] = t
		}
	}

	return &obj, nil
}
