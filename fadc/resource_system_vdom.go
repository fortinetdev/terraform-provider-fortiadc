// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system vdom.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemVdom() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemVdomRead,
		Update: resourceSystemVdomUpdate,
		Create: resourceSystemVdomCreate,
		Delete: resourceSystemVdomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"concurrentsession": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vs": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"l7rps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rs": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"l7cps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"l4cps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sslthroughput": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sslcps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hc": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ep": &schema.Schema{
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
		},
	}
}
func resourceSystemVdomCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVdom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdom(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemVdomRead(d, m)
}
func resourceSystemVdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemVdom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdom(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource: %v", err)
	}

	return resourceSystemVdomRead(d, m)
}
func resourceSystemVdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdom resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemVdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemVdom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource from API: %v", err)
	}
	return nil
}
func flattenSystemVdomConcurrentsession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomVs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomL7Rps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomRs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomInbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomL7Cps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomL4Cps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomLu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomSslthroughput(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomSslcps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomHc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomUg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomOutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomEp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("concurrentsession", flattenSystemVdomConcurrentsession(o["concurrentsession"], d, "concurrentsession", sv)); err != nil {
		if !fortiAPIPatch(o["concurrentsession"]) {
			return fmt.Errorf("Error reading concurrentsession: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemVdomStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vs", flattenSystemVdomVs(o["vs"], d, "vs", sv)); err != nil {
		if !fortiAPIPatch(o["vs"]) {
			return fmt.Errorf("Error reading vs: %v", err)
		}
	}

	if err = d.Set("l7rps", flattenSystemVdomL7Rps(o["l7rps"], d, "l7rps", sv)); err != nil {
		if !fortiAPIPatch(o["l7rps"]) {
			return fmt.Errorf("Error reading l7rps: %v", err)
		}
	}

	if err = d.Set("rs", flattenSystemVdomRs(o["rs"], d, "rs", sv)); err != nil {
		if !fortiAPIPatch(o["rs"]) {
			return fmt.Errorf("Error reading rs: %v", err)
		}
	}

	if err = d.Set("inbound", flattenSystemVdomInbound(o["inbound"], d, "inbound", sv)); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("sp", flattenSystemVdomSp(o["sp"], d, "sp", sv)); err != nil {
		if !fortiAPIPatch(o["sp"]) {
			return fmt.Errorf("Error reading sp: %v", err)
		}
	}

	if err = d.Set("l7cps", flattenSystemVdomL7Cps(o["l7cps"], d, "l7cps", sv)); err != nil {
		if !fortiAPIPatch(o["l7cps"]) {
			return fmt.Errorf("Error reading l7cps: %v", err)
		}
	}

	if err = d.Set("l4cps", flattenSystemVdomL4Cps(o["l4cps"], d, "l4cps", sv)); err != nil {
		if !fortiAPIPatch(o["l4cps"]) {
			return fmt.Errorf("Error reading l4cps: %v", err)
		}
	}

	if err = d.Set("lu", flattenSystemVdomLu(o["lu"], d, "lu", sv)); err != nil {
		if !fortiAPIPatch(o["lu"]) {
			return fmt.Errorf("Error reading lu: %v", err)
		}
	}

	if err = d.Set("sslthroughput", flattenSystemVdomSslthroughput(o["sslthroughput"], d, "sslthroughput", sv)); err != nil {
		if !fortiAPIPatch(o["sslthroughput"]) {
			return fmt.Errorf("Error reading sslthroughput: %v", err)
		}
	}

	if err = d.Set("sslcps", flattenSystemVdomSslcps(o["sslcps"], d, "sslcps", sv)); err != nil {
		if !fortiAPIPatch(o["sslcps"]) {
			return fmt.Errorf("Error reading sslcps: %v", err)
		}
	}

	if err = d.Set("hc", flattenSystemVdomHc(o["hc"], d, "hc", sv)); err != nil {
		if !fortiAPIPatch(o["hc"]) {
			return fmt.Errorf("Error reading hc: %v", err)
		}
	}

	if err = d.Set("ug", flattenSystemVdomUg(o["ug"], d, "ug", sv)); err != nil {
		if !fortiAPIPatch(o["ug"]) {
			return fmt.Errorf("Error reading ug: %v", err)
		}
	}

	if err = d.Set("outbound", flattenSystemVdomOutbound(o["outbound"], d, "outbound", sv)); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("ep", flattenSystemVdomEp(o["ep"], d, "ep", sv)); err != nil {
		if !fortiAPIPatch(o["ep"]) {
			return fmt.Errorf("Error reading ep: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemVdomMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemVdomConcurrentsession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomVs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomL7Rps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomRs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomInbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomL7Cps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomL4Cps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomLu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSslthroughput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSslcps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomHc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomUg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomOutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomEp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("concurrentsession"); ok {
		t, err := expandSystemVdomConcurrentsession(d, v, "concurrentsession", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrentsession"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemVdomStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vs"); ok {
		t, err := expandSystemVdomVs(d, v, "vs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vs"] = t
		}
	}

	if v, ok := d.GetOk("l7rps"); ok {
		t, err := expandSystemVdomL7Rps(d, v, "l7rps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l7rps"] = t
		}
	}

	if v, ok := d.GetOk("rs"); ok {
		t, err := expandSystemVdomRs(d, v, "rs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rs"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok {
		t, err := expandSystemVdomInbound(d, v, "inbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("sp"); ok {
		t, err := expandSystemVdomSp(d, v, "sp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp"] = t
		}
	}

	if v, ok := d.GetOk("l7cps"); ok {
		t, err := expandSystemVdomL7Cps(d, v, "l7cps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l7cps"] = t
		}
	}

	if v, ok := d.GetOk("l4cps"); ok {
		t, err := expandSystemVdomL4Cps(d, v, "l4cps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4cps"] = t
		}
	}

	if v, ok := d.GetOk("lu"); ok {
		t, err := expandSystemVdomLu(d, v, "lu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lu"] = t
		}
	}

	if v, ok := d.GetOk("sslthroughput"); ok {
		t, err := expandSystemVdomSslthroughput(d, v, "sslthroughput", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslthroughput"] = t
		}
	}

	if v, ok := d.GetOk("sslcps"); ok {
		t, err := expandSystemVdomSslcps(d, v, "sslcps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslcps"] = t
		}
	}

	if v, ok := d.GetOk("hc"); ok {
		t, err := expandSystemVdomHc(d, v, "hc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hc"] = t
		}
	}

	if v, ok := d.GetOk("ug"); ok {
		t, err := expandSystemVdomUg(d, v, "ug", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ug"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok {
		t, err := expandSystemVdomOutbound(d, v, "outbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("ep"); ok {
		t, err := expandSystemVdomEp(d, v, "ep", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemVdomMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
