// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system certificate intermediate ca.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateIntermediateCa() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateIntermediateCaRead,
		Update: resourceSystemCertificateIntermediateCaUpdate,
		Create: resourceSystemCertificateIntermediateCaCreate,
		Delete: resourceSystemCertificateIntermediateCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scep_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ca_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSystemCertificateIntermediateCaCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCa: type error")
	}

	obj, err := getObjectSystemCertificateIntermediateCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCa resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateIntermediateCa(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCa resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateIntermediateCaRead(d, m)
}
func resourceSystemCertificateIntermediateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateIntermediateCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateIntermediateCa(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCa resource: %v", err)
	}

	return resourceSystemCertificateIntermediateCaRead(d, m)
}
func resourceSystemCertificateIntermediateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateIntermediateCa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateIntermediateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemCertificateIntermediateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateIntermediateCa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateIntermediateCa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCa resource from API: %v", err)
	}
	return nil
}
func flattenSystemCertificateIntermediateCaMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateIntermediateCa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateIntermediateCaMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateIntermediateCaScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}
func expandSystemCertificateIntermediateCaCaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}
func expandSystemCertificateIntermediateCaHostHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateIntermediateCaMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateIntermediateCa(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateIntermediateCaMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}
	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandSystemCertificateIntermediateCaScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep_url"] = t
		}
	}
	if v, ok := d.GetOk("ca_id"); ok {
		t, err := expandSystemCertificateIntermediateCaCaId(d, v, "ca_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca_id"] = t
		}
	}
	if v, ok := d.GetOk("host_header"); ok {
		t, err := expandSystemCertificateIntermediateCaHostHeader(d, v, "host_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host_header"] = t
		}
	}

	return &obj, nil
}
