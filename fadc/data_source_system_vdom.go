// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system vdom.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemVdom() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemVdomRead,
		Schema: map[string]*schema.Schema{
			"concurrentsession": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"l7rps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rs": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"l7cps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"l4cps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslthroughput": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslcps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hc": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ep": &schema.Schema{
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
func dataSourceSystemVdomRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemVdom: type error")
	}

	o, err := c.ReadSystemVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemVdom: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemVdom(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemVdom from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemVdomConcurrentsession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomVs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomL7Rps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomRs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomL7Cps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomL4Cps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomLu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSslthroughput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSslcps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomHc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomUg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomEp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemVdom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("concurrentsession", dataSourceFlattenSystemVdomConcurrentsession(o["concurrentsession"], d, "concurrentsession")); err != nil {
		if !fortiAPIPatch(o["concurrentsession"]) {
			return fmt.Errorf("Error reading concurrentsession: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemVdomStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vs", dataSourceFlattenSystemVdomVs(o["vs"], d, "vs")); err != nil {
		if !fortiAPIPatch(o["vs"]) {
			return fmt.Errorf("Error reading vs: %v", err)
		}
	}

	if err = d.Set("l7rps", dataSourceFlattenSystemVdomL7Rps(o["l7rps"], d, "l7rps")); err != nil {
		if !fortiAPIPatch(o["l7rps"]) {
			return fmt.Errorf("Error reading l7rps: %v", err)
		}
	}

	if err = d.Set("rs", dataSourceFlattenSystemVdomRs(o["rs"], d, "rs")); err != nil {
		if !fortiAPIPatch(o["rs"]) {
			return fmt.Errorf("Error reading rs: %v", err)
		}
	}

	if err = d.Set("inbound", dataSourceFlattenSystemVdomInbound(o["inbound"], d, "inbound")); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("sp", dataSourceFlattenSystemVdomSp(o["sp"], d, "sp")); err != nil {
		if !fortiAPIPatch(o["sp"]) {
			return fmt.Errorf("Error reading sp: %v", err)
		}
	}

	if err = d.Set("l7cps", dataSourceFlattenSystemVdomL7Cps(o["l7cps"], d, "l7cps")); err != nil {
		if !fortiAPIPatch(o["l7cps"]) {
			return fmt.Errorf("Error reading l7cps: %v", err)
		}
	}

	if err = d.Set("l4cps", dataSourceFlattenSystemVdomL4Cps(o["l4cps"], d, "l4cps")); err != nil {
		if !fortiAPIPatch(o["l4cps"]) {
			return fmt.Errorf("Error reading l4cps: %v", err)
		}
	}

	if err = d.Set("lu", dataSourceFlattenSystemVdomLu(o["lu"], d, "lu")); err != nil {
		if !fortiAPIPatch(o["lu"]) {
			return fmt.Errorf("Error reading lu: %v", err)
		}
	}

	if err = d.Set("sslthroughput", dataSourceFlattenSystemVdomSslthroughput(o["sslthroughput"], d, "sslthroughput")); err != nil {
		if !fortiAPIPatch(o["sslthroughput"]) {
			return fmt.Errorf("Error reading sslthroughput: %v", err)
		}
	}

	if err = d.Set("sslcps", dataSourceFlattenSystemVdomSslcps(o["sslcps"], d, "sslcps")); err != nil {
		if !fortiAPIPatch(o["sslcps"]) {
			return fmt.Errorf("Error reading sslcps: %v", err)
		}
	}

	if err = d.Set("hc", dataSourceFlattenSystemVdomHc(o["hc"], d, "hc")); err != nil {
		if !fortiAPIPatch(o["hc"]) {
			return fmt.Errorf("Error reading hc: %v", err)
		}
	}

	if err = d.Set("ug", dataSourceFlattenSystemVdomUg(o["ug"], d, "ug")); err != nil {
		if !fortiAPIPatch(o["ug"]) {
			return fmt.Errorf("Error reading ug: %v", err)
		}
	}

	if err = d.Set("outbound", dataSourceFlattenSystemVdomOutbound(o["outbound"], d, "outbound")); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("ep", dataSourceFlattenSystemVdomEp(o["ep"], d, "ep")); err != nil {
		if !fortiAPIPatch(o["ep"]) {
			return fmt.Errorf("Error reading ep: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemVdomMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
