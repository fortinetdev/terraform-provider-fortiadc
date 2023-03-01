// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security antivirus profile.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSecurityAntivirusProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityAntivirusProfileRead,
		Schema: map[string]*schema.Schema{
			"uncomp_nest_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"emulator": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"analytics_db": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_virus_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"analytics_max_upload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uncomp_size_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oversize": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ftgd_analytics": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"options": &schema.Schema{
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
func dataSourceSecurityAntivirusProfileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SecurityAntivirusProfile: type error")
	}

	o, err := c.ReadSecurityAntivirusProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SecurityAntivirusProfile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSecurityAntivirusProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SecurityAntivirusProfile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSecurityAntivirusProfileUncompNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileAvBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileAnalyticsDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileAvVirusLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileAnalyticsMaxUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileUncompSizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileOversize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileFtgdAnalytics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSecurityAntivirusProfileMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSecurityAntivirusProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("uncomp_nest_limit", dataSourceFlattenSecurityAntivirusProfileUncompNestLimit(o["uncomp-nest-limit"], d, "uncomp_nest_limit")); err != nil {
		if !fortiAPIPatch(o["uncomp-nest-limit"]) {
			return fmt.Errorf("Error reading uncomp-nest-limit: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", dataSourceFlattenSecurityAntivirusProfileStreamingContentBypass(o["streaming-content-bypass"], d, "streaming_content_bypass")); err != nil {
		if !fortiAPIPatch(o["streaming-content-bypass"]) {
			return fmt.Errorf("Error reading streaming-content-bypass: %v", err)
		}
	}

	if err = d.Set("emulator", dataSourceFlattenSecurityAntivirusProfileEmulator(o["emulator"], d, "emulator")); err != nil {
		if !fortiAPIPatch(o["emulator"]) {
			return fmt.Errorf("Error reading emulator: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", dataSourceFlattenSecurityAntivirusProfileScanBzip2(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if !fortiAPIPatch(o["scan-bzip2"]) {
			return fmt.Errorf("Error reading scan-bzip2: %v", err)
		}
	}

	if err = d.Set("av_block_log", dataSourceFlattenSecurityAntivirusProfileAvBlockLog(o["av-block-log"], d, "av_block_log")); err != nil {
		if !fortiAPIPatch(o["av-block-log"]) {
			return fmt.Errorf("Error reading av-block-log: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSecurityAntivirusProfileComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("analytics_db", dataSourceFlattenSecurityAntivirusProfileAnalyticsDb(o["analytics-db"], d, "analytics_db")); err != nil {
		if !fortiAPIPatch(o["analytics-db"]) {
			return fmt.Errorf("Error reading analytics-db: %v", err)
		}
	}

	if err = d.Set("oversize_limit", dataSourceFlattenSecurityAntivirusProfileOversizeLimit(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if !fortiAPIPatch(o["oversize-limit"]) {
			return fmt.Errorf("Error reading oversize-limit: %v", err)
		}
	}

	if err = d.Set("av_virus_log", dataSourceFlattenSecurityAntivirusProfileAvVirusLog(o["av-virus-log"], d, "av_virus_log")); err != nil {
		if !fortiAPIPatch(o["av-virus-log"]) {
			return fmt.Errorf("Error reading av-virus-log: %v", err)
		}
	}

	if err = d.Set("analytics_max_upload", dataSourceFlattenSecurityAntivirusProfileAnalyticsMaxUpload(o["analytics-max-upload"], d, "analytics_max_upload")); err != nil {
		if !fortiAPIPatch(o["analytics-max-upload"]) {
			return fmt.Errorf("Error reading analytics-max-upload: %v", err)
		}
	}

	if err = d.Set("uncomp_size_limit", dataSourceFlattenSecurityAntivirusProfileUncompSizeLimit(o["uncomp-size-limit"], d, "uncomp_size_limit")); err != nil {
		if !fortiAPIPatch(o["uncomp-size-limit"]) {
			return fmt.Errorf("Error reading uncomp-size-limit: %v", err)
		}
	}

	if err = d.Set("oversize", dataSourceFlattenSecurityAntivirusProfileOversize(o["oversize"], d, "oversize")); err != nil {
		if !fortiAPIPatch(o["oversize"]) {
			return fmt.Errorf("Error reading oversize: %v", err)
		}
	}

	if err = d.Set("ftgd_analytics", dataSourceFlattenSecurityAntivirusProfileFtgdAnalytics(o["ftgd-analytics"], d, "ftgd_analytics")); err != nil {
		if !fortiAPIPatch(o["ftgd-analytics"]) {
			return fmt.Errorf("Error reading ftgd-analytics: %v", err)
		}
	}

	if err = d.Set("options", dataSourceFlattenSecurityAntivirusProfileOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSecurityAntivirusProfileMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}
