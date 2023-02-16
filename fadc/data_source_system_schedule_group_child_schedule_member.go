// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system schedule group child schedule member.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemScheduleGroupChildScheduleMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemScheduleGroupChildScheduleMemberRead,
		Schema: map[string]*schema.Schema{
			"startdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enddate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"day_of_month": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"day_of_week": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"starttime_of_startdate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"endtime_of_enddate": &schema.Schema{
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
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceSystemScheduleGroupChildScheduleMemberRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemScheduleGroupChildScheduleMember(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemScheduleGroupChildScheduleMember(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemScheduleGroupChildScheduleMember from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemScheduleGroupChildScheduleMemberStartdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberEnddate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberDayOfMonth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberDayOfWeek(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberStarttimeOfStartdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberEndtimeOfEnddate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemScheduleGroupChildScheduleMemberMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemScheduleGroupChildScheduleMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("startdate", dataSourceFlattenSystemScheduleGroupChildScheduleMemberStartdate(o["startdate"], d, "startdate")); err != nil {
		if !fortiAPIPatch(o["startdate"]) {
			return fmt.Errorf("Error reading startdate: %v", err)
		}
	}

	if err = d.Set("enddate", dataSourceFlattenSystemScheduleGroupChildScheduleMemberEnddate(o["enddate"], d, "enddate")); err != nil {
		if !fortiAPIPatch(o["enddate"]) {
			return fmt.Errorf("Error reading enddate: %v", err)
		}
	}

	if err = d.Set("day_of_month", dataSourceFlattenSystemScheduleGroupChildScheduleMemberDayOfMonth(o["day-of-month"], d, "day_of_month")); err != nil {
		if !fortiAPIPatch(o["day-of-month"]) {
			return fmt.Errorf("Error reading day-of-month: %v", err)
		}
	}

	if err = d.Set("day_of_week", dataSourceFlattenSystemScheduleGroupChildScheduleMemberDayOfWeek(o["day-of-week"], d, "day_of_week")); err != nil {
		if !fortiAPIPatch(o["day-of-week"]) {
			return fmt.Errorf("Error reading day-of-week: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemScheduleGroupChildScheduleMemberType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("starttime_of_startdate", dataSourceFlattenSystemScheduleGroupChildScheduleMemberStarttimeOfStartdate(o["starttime-of-startdate"], d, "starttime_of_startdate")); err != nil {
		if !fortiAPIPatch(o["starttime-of-startdate"]) {
			return fmt.Errorf("Error reading starttime-of-startdate: %v", err)
		}
	}

	if err = d.Set("endtime_of_enddate", dataSourceFlattenSystemScheduleGroupChildScheduleMemberEndtimeOfEnddate(o["endtime-of-enddate"], d, "endtime_of_enddate")); err != nil {
		if !fortiAPIPatch(o["endtime-of-enddate"]) {
			return fmt.Errorf("Error reading endtime-of-enddate: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemScheduleGroupChildScheduleMemberMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
