// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  user adfs publish.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceUserAdfsPublish() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserAdfsPublishRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_server_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"preauth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"relying_party": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"proxy": &schema.Schema{
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
func dataSourceUserAdfsPublishRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserAdfsPublish: type error")
	}

	o, err := c.ReadUserAdfsPublish(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing UserAdfsPublish: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectUserAdfsPublish(d, o)
	if err != nil {
		return fmt.Errorf("Error describing UserAdfsPublish from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenUserAdfsPublishStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserAdfsPublishBackendServerUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserAdfsPublishExternalUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserAdfsPublishPreauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserAdfsPublishRelyingParty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserAdfsPublishMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserAdfsPublishProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectUserAdfsPublish(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenUserAdfsPublishStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("backend_server_url", dataSourceFlattenUserAdfsPublishBackendServerUrl(o["backend-server-url"], d, "backend_server_url")); err != nil {
		if !fortiAPIPatch(o["backend-server-url"]) {
			return fmt.Errorf("Error reading backend-server-url: %v", err)
		}
	}

	if err = d.Set("external_url", dataSourceFlattenUserAdfsPublishExternalUrl(o["external-url"], d, "external_url")); err != nil {
		if !fortiAPIPatch(o["external-url"]) {
			return fmt.Errorf("Error reading external-url: %v", err)
		}
	}

	if err = d.Set("preauth", dataSourceFlattenUserAdfsPublishPreauth(o["preauth"], d, "preauth")); err != nil {
		if !fortiAPIPatch(o["preauth"]) {
			return fmt.Errorf("Error reading preauth: %v", err)
		}
	}

	if err = d.Set("relying_party", dataSourceFlattenUserAdfsPublishRelyingParty(o["relying-party"], d, "relying_party")); err != nil {
		if !fortiAPIPatch(o["relying-party"]) {
			return fmt.Errorf("Error reading relying-party: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenUserAdfsPublishMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("proxy", dataSourceFlattenUserAdfsPublishProxy(o["proxy"], d, "proxy")); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	return nil
}
