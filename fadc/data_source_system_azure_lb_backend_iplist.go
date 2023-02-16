// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure SystemAzureLbBackendIp.

package fortiadc

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemAzureLbBackendIpList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAzureLbBackendIpListRead,

		Schema: map[string]*schema.Schema{
			"mkey_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func dataSourceSystemAzureLbBackendIpListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	filter := d.Get("filter").(string)
	if filter != "" {
		filter = escapeFilter(filter)
	}
	filters := parse_filter(filter)

	o, err := c.GenericGroupRead("/api/system_azure_lb_backend_ip", filter, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemAzureLbBackendIp: %v", err)
	}

	var tmps []string
	if o != nil {
		if len(o) == 0 || o[0] == nil {
			return nil
		}

		for _, r := range o {
			i := r.(map[string]interface{})

			if !run_filter(filters, i) {
				continue
			}
			if _, ok := i["mkey"]; ok {
				tmps = append(tmps, i["mkey"].(string))
			}
		}
	}
	d.Set("mkey_list", tmps)

	d.SetId("DataSourceSystemAzureLbBackendIpList" + filter)

	return nil
}
