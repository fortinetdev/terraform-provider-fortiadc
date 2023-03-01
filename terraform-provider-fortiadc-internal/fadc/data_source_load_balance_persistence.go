// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance persistence.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalancePersistence() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancePersistenceRead,
		Schema: map[string]*schema.Schema{
			"match_across_servers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"keyvalue_relation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_secure": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_httponly": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sess_kw_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"avpno": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_maskbits": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_custom_attr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso8583_bitmap_relation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv4_maskbits": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keyword": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_custom_attr_val": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"override_connection_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cookie_samesite": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"radius_attribute_relation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceLoadBalancePersistenceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalancePersistence(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePersistence: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalancePersistence(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalancePersistence from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalancePersistenceMatchAcrossServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceKeyvalueRelation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceCookieHttponly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceSessKwType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceAvpno(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceIpv6Maskbits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceCookieCustomAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceIso8583BitmapRelation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceIpv4Maskbits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceKeyword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceCookieCustomAttrVal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceOverrideConnectionLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceCookieSamesite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalancePersistenceRadiusAttributeRelation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalancePersistence(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("match_across_servers", dataSourceFlattenLoadBalancePersistenceMatchAcrossServers(o["match-across-servers"], d, "match_across_servers")); err != nil {
		if !fortiAPIPatch(o["match-across-servers"]) {
			return fmt.Errorf("Error reading match-across-servers: %v", err)
		}
	}

	if err = d.Set("keyvalue_relation", dataSourceFlattenLoadBalancePersistenceKeyvalueRelation(o["keyvalue-relation"], d, "keyvalue_relation")); err != nil {
		if !fortiAPIPatch(o["keyvalue-relation"]) {
			return fmt.Errorf("Error reading keyvalue-relation: %v", err)
		}
	}

	if err = d.Set("cookie_secure", dataSourceFlattenLoadBalancePersistenceCookieSecure(o["cookie_secure"], d, "cookie_secure")); err != nil {
		if !fortiAPIPatch(o["cookie_secure"]) {
			return fmt.Errorf("Error reading cookie_secure: %v", err)
		}
	}

	if err = d.Set("cookie_httponly", dataSourceFlattenLoadBalancePersistenceCookieHttponly(o["cookie_httponly"], d, "cookie_httponly")); err != nil {
		if !fortiAPIPatch(o["cookie_httponly"]) {
			return fmt.Errorf("Error reading cookie_httponly: %v", err)
		}
	}

	if err = d.Set("sess_kw_type", dataSourceFlattenLoadBalancePersistenceSessKwType(o["sess_kw_type"], d, "sess_kw_type")); err != nil {
		if !fortiAPIPatch(o["sess_kw_type"]) {
			return fmt.Errorf("Error reading sess_kw_type: %v", err)
		}
	}

	if err = d.Set("avpno", dataSourceFlattenLoadBalancePersistenceAvpno(o["avpno"], d, "avpno")); err != nil {
		if !fortiAPIPatch(o["avpno"]) {
			return fmt.Errorf("Error reading avpno: %v", err)
		}
	}

	if err = d.Set("ipv6_maskbits", dataSourceFlattenLoadBalancePersistenceIpv6Maskbits(o["ipv6-maskbits"], d, "ipv6_maskbits")); err != nil {
		if !fortiAPIPatch(o["ipv6-maskbits"]) {
			return fmt.Errorf("Error reading ipv6-maskbits: %v", err)
		}
	}

	if err = d.Set("cookie_custom_attr", dataSourceFlattenLoadBalancePersistenceCookieCustomAttr(o["cookie_custom_attr"], d, "cookie_custom_attr")); err != nil {
		if !fortiAPIPatch(o["cookie_custom_attr"]) {
			return fmt.Errorf("Error reading cookie_custom_attr: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalancePersistenceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("iso8583_bitmap_relation", dataSourceFlattenLoadBalancePersistenceIso8583BitmapRelation(o["iso8583_bitmap-relation"], d, "iso8583_bitmap_relation")); err != nil {
		if !fortiAPIPatch(o["iso8583_bitmap-relation"]) {
			return fmt.Errorf("Error reading iso8583_bitmap-relation: %v", err)
		}
	}

	if err = d.Set("ipv4_maskbits", dataSourceFlattenLoadBalancePersistenceIpv4Maskbits(o["ipv4-maskbits"], d, "ipv4_maskbits")); err != nil {
		if !fortiAPIPatch(o["ipv4-maskbits"]) {
			return fmt.Errorf("Error reading ipv4-maskbits: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalancePersistenceMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("keyword", dataSourceFlattenLoadBalancePersistenceKeyword(o["keyword"], d, "keyword")); err != nil {
		if !fortiAPIPatch(o["keyword"]) {
			return fmt.Errorf("Error reading keyword: %v", err)
		}
	}

	if err = d.Set("cookie_domain", dataSourceFlattenLoadBalancePersistenceCookieDomain(o["cookie_domain"], d, "cookie_domain")); err != nil {
		if !fortiAPIPatch(o["cookie_domain"]) {
			return fmt.Errorf("Error reading cookie_domain: %v", err)
		}
	}

	if err = d.Set("cookie_custom_attr_val", dataSourceFlattenLoadBalancePersistenceCookieCustomAttrVal(o["cookie_custom_attr_val"], d, "cookie_custom_attr_val")); err != nil {
		if !fortiAPIPatch(o["cookie_custom_attr_val"]) {
			return fmt.Errorf("Error reading cookie_custom_attr_val: %v", err)
		}
	}

	if err = d.Set("override_connection_limit", dataSourceFlattenLoadBalancePersistenceOverrideConnectionLimit(o["override-connection-limit"], d, "override_connection_limit")); err != nil {
		if !fortiAPIPatch(o["override-connection-limit"]) {
			return fmt.Errorf("Error reading override-connection-limit: %v", err)
		}
	}

	if err = d.Set("cookie_samesite", dataSourceFlattenLoadBalancePersistenceCookieSamesite(o["cookie_samesite"], d, "cookie_samesite")); err != nil {
		if !fortiAPIPatch(o["cookie_samesite"]) {
			return fmt.Errorf("Error reading cookie_samesite: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenLoadBalancePersistenceTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("radius_attribute_relation", dataSourceFlattenLoadBalancePersistenceRadiusAttributeRelation(o["radius-attribute-relation"], d, "radius_attribute_relation")); err != nil {
		if !fortiAPIPatch(o["radius-attribute-relation"]) {
			return fmt.Errorf("Error reading radius-attribute-relation: %v", err)
		}
	}

	return nil
}
