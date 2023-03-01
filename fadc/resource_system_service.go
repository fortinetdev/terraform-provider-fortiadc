// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system service.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemService() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemServiceRead,
		Update: resourceSystemServiceUpdate,
		Create: resourceSystemServiceCreate,
		Delete: resourceSystemServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"destination_port_max": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_port_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"protocol_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination_port_min": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"specify_source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_port_max": &schema.Schema{
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
func resourceSystemServiceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemService resource while getting object: %v", err)
	}

	_, err = c.CreateSystemService(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemService resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemServiceRead(d, m)
}
func resourceSystemServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemService resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemService(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemService resource: %v", err)
	}

	return resourceSystemServiceRead(d, m)
}
func resourceSystemServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemService(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemService resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemService(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemService(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemService resource from API: %v", err)
	}
	return nil
}
func flattenSystemServiceDestinationPortMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceSourcePortMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceProtocolType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceDestinationPortMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceSpecifySourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceSourcePortMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemServiceMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemService(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("destination_port_max", flattenSystemServiceDestinationPortMax(o["destination-port-max"], d, "destination_port_max", sv)); err != nil {
		if !fortiAPIPatch(o["destination-port-max"]) {
			return fmt.Errorf("Error reading destination-port-max: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemServiceProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("source_port_min", flattenSystemServiceSourcePortMin(o["source-port-min"], d, "source_port_min", sv)); err != nil {
		if !fortiAPIPatch(o["source-port-min"]) {
			return fmt.Errorf("Error reading source-port-min: %v", err)
		}
	}

	if err = d.Set("protocol_type", flattenSystemServiceProtocolType(o["protocol-type"], d, "protocol_type", sv)); err != nil {
		if !fortiAPIPatch(o["protocol-type"]) {
			return fmt.Errorf("Error reading protocol-type: %v", err)
		}
	}

	if err = d.Set("destination_port_min", flattenSystemServiceDestinationPortMin(o["destination-port-min"], d, "destination_port_min", sv)); err != nil {
		if !fortiAPIPatch(o["destination-port-min"]) {
			return fmt.Errorf("Error reading destination-port-min: %v", err)
		}
	}

	if err = d.Set("specify_source_port", flattenSystemServiceSpecifySourcePort(o["specify-source-port"], d, "specify_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["specify-source-port"]) {
			return fmt.Errorf("Error reading specify-source-port: %v", err)
		}
	}

	if err = d.Set("source_port_max", flattenSystemServiceSourcePortMax(o["source-port-max"], d, "source_port_max", sv)); err != nil {
		if !fortiAPIPatch(o["source-port-max"]) {
			return fmt.Errorf("Error reading source-port-max: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemServiceMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemServiceDestinationPortMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceSourcePortMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceProtocolType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceDestinationPortMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceSpecifySourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceSourcePortMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemServiceMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemService(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("destination_port_max"); ok {
		t, err := expandSystemServiceDestinationPortMax(d, v, "destination_port_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-port-max"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemServiceProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("source_port_min"); ok {
		t, err := expandSystemServiceSourcePortMin(d, v, "source_port_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port-min"] = t
		}
	}

	if v, ok := d.GetOk("protocol_type"); ok {
		t, err := expandSystemServiceProtocolType(d, v, "protocol_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-type"] = t
		}
	}

	if v, ok := d.GetOk("destination_port_min"); ok {
		t, err := expandSystemServiceDestinationPortMin(d, v, "destination_port_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-port-min"] = t
		}
	}

	if v, ok := d.GetOk("specify_source_port"); ok {
		t, err := expandSystemServiceSpecifySourcePort(d, v, "specify_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["specify-source-port"] = t
		}
	}

	if v, ok := d.GetOk("source_port_max"); ok {
		t, err := expandSystemServiceSourcePortMax(d, v, "source_port_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port-max"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemServiceMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
