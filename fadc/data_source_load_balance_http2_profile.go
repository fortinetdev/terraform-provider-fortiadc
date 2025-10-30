// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance http2 profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceHttp2Profile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceHttp2ProfileRead,
		Schema: map[string]*schema.Schema{
			"max_concurrent_stream": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_frame_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_constraint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_receive_window": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header_table_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_header_list_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backend_http2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_max_receive_window": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_concurrent_stream": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_proto_mode_https": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_proto_mode_http": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backend_multiplex_mode": &schema.Schema{
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
func dataSourceLoadBalanceHttp2ProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceHttp2Profile: type error")
	}

	o, err := c.ReadLoadBalanceHttp2Profile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceHttp2Profile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceHttp2Profile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceHttp2Profile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceHttp2ProfileMaxConcurrentStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileMaxFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfilePriorityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileSslConstraint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileMaxReceiveWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileHeaderTableSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileUpgradeMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileMaxHeaderListSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileBackendHttp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileBackendMaxReceiveWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileBackendConcurrentStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileBackendProtoModeHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileBackendProtoModeHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceHttp2ProfileBackendMultiplexMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceHttp2Profile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_concurrent_stream", dataSourceFlattenLoadBalanceHttp2ProfileMaxConcurrentStream(o["max_concurrent_stream"], d, "max_concurrent_stream")); err != nil {
		if !fortiAPIPatch(o["max_concurrent_stream"]) {
			return fmt.Errorf("Error reading max_concurrent_stream: %v", err)
		}
	}

	if err = d.Set("max_frame_size", dataSourceFlattenLoadBalanceHttp2ProfileMaxFrameSize(o["max_frame_size"], d, "max_frame_size")); err != nil {
		if !fortiAPIPatch(o["max_frame_size"]) {
			return fmt.Errorf("Error reading max_frame_size: %v", err)
		}
	}

	if err = d.Set("priority_mode", dataSourceFlattenLoadBalanceHttp2ProfilePriorityMode(o["priority_mode"], d, "priority_mode")); err != nil {
		if !fortiAPIPatch(o["priority_mode"]) {
			return fmt.Errorf("Error reading priority_mode: %v", err)
		}
	}

	if err = d.Set("ssl_constraint", dataSourceFlattenLoadBalanceHttp2ProfileSslConstraint(o["ssl_constraint"], d, "ssl_constraint")); err != nil {
		if !fortiAPIPatch(o["ssl_constraint"]) {
			return fmt.Errorf("Error reading ssl_constraint: %v", err)
		}
	}

	if err = d.Set("max_receive_window", dataSourceFlattenLoadBalanceHttp2ProfileMaxReceiveWindow(o["max_receive_window"], d, "max_receive_window")); err != nil {
		if !fortiAPIPatch(o["max_receive_window"]) {
			return fmt.Errorf("Error reading max_receive_window: %v", err)
		}
	}

	if err = d.Set("header_table_size", dataSourceFlattenLoadBalanceHttp2ProfileHeaderTableSize(o["header_table_size"], d, "header_table_size")); err != nil {
		if !fortiAPIPatch(o["header_table_size"]) {
			return fmt.Errorf("Error reading header_table_size: %v", err)
		}
	}

	if err = d.Set("upgrade_mode", dataSourceFlattenLoadBalanceHttp2ProfileUpgradeMode(o["upgrade_mode"], d, "upgrade_mode")); err != nil {
		if !fortiAPIPatch(o["upgrade_mode"]) {
			return fmt.Errorf("Error reading upgrade_mode: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceHttp2ProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("max_header_list_size", dataSourceFlattenLoadBalanceHttp2ProfileMaxHeaderListSize(o["max_header_list_size"], d, "max_header_list_size")); err != nil {
		if !fortiAPIPatch(o["max_header_list_size"]) {
			return fmt.Errorf("Error reading max_header_list_size: %v", err)
		}
	}

	if err = d.Set("backend_http2", dataSourceFlattenLoadBalanceHttp2ProfileBackendHttp2(o["backend_http2"], d, "backend_http2")); err != nil {
		if !fortiAPIPatch(o["backend_http2"]) {
			return fmt.Errorf("Error reading backend_http2: %v", err)
		}
	}

	if err = d.Set("backend_max_receive_window", dataSourceFlattenLoadBalanceHttp2ProfileBackendMaxReceiveWindow(o["backend_max_receive_window"], d, "backend_max_receive_window")); err != nil {
		if !fortiAPIPatch(o["backend_max_receive_window"]) {
			return fmt.Errorf("Error reading backend_max_receive_window: %v", err)
		}
	}

	if err = d.Set("backend_concurrent_stream", dataSourceFlattenLoadBalanceHttp2ProfileBackendConcurrentStream(o["backend_concurrent_stream"], d, "backend_concurrent_stream")); err != nil {
		if !fortiAPIPatch(o["backend_concurrent_stream"]) {
			return fmt.Errorf("Error reading backend_concurrent_stream: %v", err)
		}
	}

	if err = d.Set("backend_proto_mode_https", dataSourceFlattenLoadBalanceHttp2ProfileBackendProtoModeHttps(o["backend_proto_mode_https"], d, "backend_proto_mode_https")); err != nil {
		if !fortiAPIPatch(o["backend_proto_mode_https"]) {
			return fmt.Errorf("Error reading backend_proto_mode_https: %v", err)
		}
	}

	if err = d.Set("backend_proto_mode_http", dataSourceFlattenLoadBalanceHttp2ProfileBackendProtoModeHttp(o["backend_proto_mode_http"], d, "backend_proto_mode_http")); err != nil {
		if !fortiAPIPatch(o["backend_proto_mode_http"]) {
			return fmt.Errorf("Error reading backend_proto_mode_http: %v", err)
		}
	}

	if err = d.Set("backend_multiplex_mode", dataSourceFlattenLoadBalanceHttp2ProfileBackendMultiplexMode(o["backend_multiplex_mode"], d, "backend_multiplex_mode")); err != nil {
		if !fortiAPIPatch(o["backend_multiplex_mode"]) {
			return fmt.Errorf("Error reading backend_multiplex_mode: %v", err)
		}
	}

	return nil
}
