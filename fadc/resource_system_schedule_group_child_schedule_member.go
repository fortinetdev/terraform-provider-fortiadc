// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system schedule group child schedule member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemScheduleGroupChildScheduleMember() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemScheduleGroupChildScheduleMemberRead,
		Update: resourceSystemScheduleGroupChildScheduleMemberUpdate,
		Create: resourceSystemScheduleGroupChildScheduleMemberCreate,
		Delete: resourceSystemScheduleGroupChildScheduleMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"startdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"enddate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"day_of_month": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"day_of_week": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"starttime_of_startdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"endtime_of_enddate": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceSystemScheduleGroupChildScheduleMemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember: type error")
	}

	obj, err := getObjectSystemScheduleGroupChildScheduleMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleGroupChildScheduleMember resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateSystemScheduleGroupChildScheduleMember(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleGroupChildScheduleMember resource: %v", err)
	}

	d.SetId(new_id)

	return resourceSystemScheduleGroupChildScheduleMemberRead(d, m)
}
func resourceSystemScheduleGroupChildScheduleMemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectSystemScheduleGroupChildScheduleMember(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleGroupChildScheduleMember resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemScheduleGroupChildScheduleMember(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleGroupChildScheduleMember resource: %v", err)
	}

	return resourceSystemScheduleGroupChildScheduleMemberRead(d, m)
}
func resourceSystemScheduleGroupChildScheduleMemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteSystemScheduleGroupChildScheduleMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemScheduleGroupChildScheduleMember resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemScheduleGroupChildScheduleMemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadSystemScheduleGroupChildScheduleMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleGroupChildScheduleMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemScheduleGroupChildScheduleMember(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleGroupChildScheduleMember resource from API: %v", err)
	}
	return nil
}
func flattenSystemScheduleGroupChildScheduleMemberStartdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberEnddate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberDayOfMonth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberDayOfWeek(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberStarttimeOfStartdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberEndtimeOfEnddate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleGroupChildScheduleMemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemScheduleGroupChildScheduleMember(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("startdate", flattenSystemScheduleGroupChildScheduleMemberStartdate(o["startdate"], d, "startdate", sv)); err != nil {
		if !fortiAPIPatch(o["startdate"]) {
			return fmt.Errorf("Error reading startdate: %v", err)
		}
	}

	if err = d.Set("enddate", flattenSystemScheduleGroupChildScheduleMemberEnddate(o["enddate"], d, "enddate", sv)); err != nil {
		if !fortiAPIPatch(o["enddate"]) {
			return fmt.Errorf("Error reading enddate: %v", err)
		}
	}

	if err = d.Set("day_of_month", flattenSystemScheduleGroupChildScheduleMemberDayOfMonth(o["day-of-month"], d, "day_of_month", sv)); err != nil {
		if !fortiAPIPatch(o["day-of-month"]) {
			return fmt.Errorf("Error reading day-of-month: %v", err)
		}
	}

	if err = d.Set("day_of_week", flattenSystemScheduleGroupChildScheduleMemberDayOfWeek(o["day-of-week"], d, "day_of_week", sv)); err != nil {
		if !fortiAPIPatch(o["day-of-week"]) {
			return fmt.Errorf("Error reading day-of-week: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemScheduleGroupChildScheduleMemberType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("starttime_of_startdate", flattenSystemScheduleGroupChildScheduleMemberStarttimeOfStartdate(o["starttime-of-startdate"], d, "starttime_of_startdate", sv)); err != nil {
		if !fortiAPIPatch(o["starttime-of-startdate"]) {
			return fmt.Errorf("Error reading starttime-of-startdate: %v", err)
		}
	}

	if err = d.Set("endtime_of_enddate", flattenSystemScheduleGroupChildScheduleMemberEndtimeOfEnddate(o["endtime-of-enddate"], d, "endtime_of_enddate", sv)); err != nil {
		if !fortiAPIPatch(o["endtime-of-enddate"]) {
			return fmt.Errorf("Error reading endtime-of-enddate: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemScheduleGroupChildScheduleMemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemScheduleGroupChildScheduleMemberStartdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberEnddate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberDayOfMonth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberDayOfWeek(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberStarttimeOfStartdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberEndtimeOfEnddate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleGroupChildScheduleMemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemScheduleGroupChildScheduleMember(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("startdate"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberStartdate(d, v, "startdate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startdate"] = t
		}
	}

	if v, ok := d.GetOk("enddate"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberEnddate(d, v, "enddate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enddate"] = t
		}
	}

	if v, ok := d.GetOk("day_of_month"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberDayOfMonth(d, v, "day_of_month", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day-of-month"] = t
		}
	}

	if v, ok := d.GetOk("day_of_week"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberDayOfWeek(d, v, "day_of_week", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day-of-week"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("starttime_of_startdate"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberStarttimeOfStartdate(d, v, "starttime_of_startdate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["starttime-of-startdate"] = t
		}
	}

	if v, ok := d.GetOk("endtime_of_enddate"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberEndtimeOfEnddate(d, v, "endtime_of_enddate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endtime-of-enddate"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemScheduleGroupChildScheduleMemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
