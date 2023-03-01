// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure SystemInterfaceChildHaNodeIpList.

package fortiadc

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemInterfaceChildHaNodeIpListList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceChildHaNodeIpListListRead,

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
			"pkey": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSystemInterfaceChildHaNodeIpListListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.GenericGroupRead("/api/system_interface_child_ha_node_ip_list", filter, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterfaceChildHaNodeIpList: %v", err)
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
			if _, ok := i["id"]; ok {
				tmps = append(tmps, i["id"].(string))
			}
		}
	}
	d.Set("mkey_list", tmps)

	d.SetId("DataSourceSystemInterfaceChildHaNodeIpListList" + filter)

	return nil
}
