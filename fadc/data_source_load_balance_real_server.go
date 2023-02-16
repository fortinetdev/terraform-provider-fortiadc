// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  load balance real server.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceLoadBalanceRealServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalanceRealServerRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_addr_private": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
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
func dataSourceLoadBalanceRealServerRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LoadBalanceRealServer: type error")
	}

	o, err := c.ReadLoadBalanceRealServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceRealServer: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectLoadBalanceRealServer(d, o)
	if err != nil {
		return fmt.Errorf("Error describing LoadBalanceRealServer from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenLoadBalanceRealServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSdnConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerSdnAddrPrivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerInstance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenLoadBalanceRealServerMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectLoadBalanceRealServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenLoadBalanceRealServerStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server_type", dataSourceFlattenLoadBalanceRealServerServerType(o["server_type"], d, "server_type")); err != nil {
		if !fortiAPIPatch(o["server_type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("sdn_connector", dataSourceFlattenLoadBalanceRealServerSdnConnector(o["sdn_connector"], d, "sdn_connector")); err != nil {
		if !fortiAPIPatch(o["sdn_connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	if err = d.Set("sdn_addr_private", dataSourceFlattenLoadBalanceRealServerSdnAddrPrivate(o["sdn_addr_private"], d, "sdn_addr_private")); err != nil {
		if !fortiAPIPatch(o["sdn_addr_private"]) {
			return fmt.Errorf("Error reading sdn_addr_private: %v", err)
		}
	}

	if err = d.Set("fqdn", dataSourceFlattenLoadBalanceRealServerFqdn(o["FQDN"], d, "fqdn")); err != nil {
		if !fortiAPIPatch(o["FQDN"]) {
			return fmt.Errorf("Error reading FQDN: %v", err)
		}
	}

	if err = d.Set("address6", dataSourceFlattenLoadBalanceRealServerAddress6(o["address6"], d, "address6")); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("instance", dataSourceFlattenLoadBalanceRealServerInstance(o["instance"], d, "instance")); err != nil {
		if !fortiAPIPatch(o["instance"]) {
			return fmt.Errorf("Error reading instance: %v", err)
		}
	}

	if err = d.Set("address", dataSourceFlattenLoadBalanceRealServerAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenLoadBalanceRealServerType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenLoadBalanceRealServerMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
