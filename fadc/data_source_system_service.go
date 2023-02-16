// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system service.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemServiceRead,
		Schema: map[string]*schema.Schema{
			"destination_port_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_port_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"specify_source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceSystemServiceRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemService: type error")
	}

	o, err := c.ReadSystemService(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemService: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemService(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemService from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemServiceDestinationPortMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceSourcePortMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceProtocolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceDestinationPortMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceSpecifySourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceSourcePortMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemServiceMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("destination_port_max", dataSourceFlattenSystemServiceDestinationPortMax(o["destination-port-max"], d, "destination_port_max")); err != nil {
		if !fortiAPIPatch(o["destination-port-max"]) {
			return fmt.Errorf("Error reading destination-port-max: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenSystemServiceProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("source_port_min", dataSourceFlattenSystemServiceSourcePortMin(o["source-port-min"], d, "source_port_min")); err != nil {
		if !fortiAPIPatch(o["source-port-min"]) {
			return fmt.Errorf("Error reading source-port-min: %v", err)
		}
	}

	if err = d.Set("protocol_type", dataSourceFlattenSystemServiceProtocolType(o["protocol-type"], d, "protocol_type")); err != nil {
		if !fortiAPIPatch(o["protocol-type"]) {
			return fmt.Errorf("Error reading protocol-type: %v", err)
		}
	}

	if err = d.Set("destination_port_min", dataSourceFlattenSystemServiceDestinationPortMin(o["destination-port-min"], d, "destination_port_min")); err != nil {
		if !fortiAPIPatch(o["destination-port-min"]) {
			return fmt.Errorf("Error reading destination-port-min: %v", err)
		}
	}

	if err = d.Set("specify_source_port", dataSourceFlattenSystemServiceSpecifySourcePort(o["specify-source-port"], d, "specify_source_port")); err != nil {
		if !fortiAPIPatch(o["specify-source-port"]) {
			return fmt.Errorf("Error reading specify-source-port: %v", err)
		}
	}

	if err = d.Set("source_port_max", dataSourceFlattenSystemServiceSourcePortMax(o["source-port-max"], d, "source_port_max")); err != nil {
		if !fortiAPIPatch(o["source-port-max"]) {
			return fmt.Errorf("Error reading source-port-max: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemServiceMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
