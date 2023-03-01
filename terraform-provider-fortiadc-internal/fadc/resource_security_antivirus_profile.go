// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  security antivirus profile.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityAntivirusProfile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSecurityAntivirusProfileRead,
		Update: resourceSecurityAntivirusProfileUpdate,
		Create: resourceSecurityAntivirusProfileCreate,
		Delete: resourceSecurityAntivirusProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"uncomp_nest_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"emulator": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"av_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"analytics_db": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"av_virus_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"analytics_max_upload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"uncomp_size_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oversize": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ftgd_analytics": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"options": &schema.Schema{
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
func resourceSecurityAntivirusProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSecurityAntivirusProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SecurityAntivirusProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSecurityAntivirusProfile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SecurityAntivirusProfile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSecurityAntivirusProfileRead(d, m)
}
func resourceSecurityAntivirusProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSecurityAntivirusProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SecurityAntivirusProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityAntivirusProfile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SecurityAntivirusProfile resource: %v", err)
	}

	return resourceSecurityAntivirusProfileRead(d, m)
}
func resourceSecurityAntivirusProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSecurityAntivirusProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SecurityAntivirusProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSecurityAntivirusProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSecurityAntivirusProfile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SecurityAntivirusProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSecurityAntivirusProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SecurityAntivirusProfile resource from API: %v", err)
	}
	return nil
}
func flattenSecurityAntivirusProfileUncompNestLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileEmulator(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileScanBzip2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileAvBlockLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileAnalyticsDb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileOversizeLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileAvVirusLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileAnalyticsMaxUpload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileUncompSizeLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileOversize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileFtgdAnalytics(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSecurityAntivirusProfileMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSecurityAntivirusProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("uncomp_nest_limit", flattenSecurityAntivirusProfileUncompNestLimit(o["uncomp-nest-limit"], d, "uncomp_nest_limit", sv)); err != nil {
		if !fortiAPIPatch(o["uncomp-nest-limit"]) {
			return fmt.Errorf("Error reading uncomp-nest-limit: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", flattenSecurityAntivirusProfileStreamingContentBypass(o["streaming-content-bypass"], d, "streaming_content_bypass", sv)); err != nil {
		if !fortiAPIPatch(o["streaming-content-bypass"]) {
			return fmt.Errorf("Error reading streaming-content-bypass: %v", err)
		}
	}

	if err = d.Set("emulator", flattenSecurityAntivirusProfileEmulator(o["emulator"], d, "emulator", sv)); err != nil {
		if !fortiAPIPatch(o["emulator"]) {
			return fmt.Errorf("Error reading emulator: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenSecurityAntivirusProfileScanBzip2(o["scan-bzip2"], d, "scan_bzip2", sv)); err != nil {
		if !fortiAPIPatch(o["scan-bzip2"]) {
			return fmt.Errorf("Error reading scan-bzip2: %v", err)
		}
	}

	if err = d.Set("av_block_log", flattenSecurityAntivirusProfileAvBlockLog(o["av-block-log"], d, "av_block_log", sv)); err != nil {
		if !fortiAPIPatch(o["av-block-log"]) {
			return fmt.Errorf("Error reading av-block-log: %v", err)
		}
	}

	if err = d.Set("comments", flattenSecurityAntivirusProfileComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("analytics_db", flattenSecurityAntivirusProfileAnalyticsDb(o["analytics-db"], d, "analytics_db", sv)); err != nil {
		if !fortiAPIPatch(o["analytics-db"]) {
			return fmt.Errorf("Error reading analytics-db: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenSecurityAntivirusProfileOversizeLimit(o["oversize-limit"], d, "oversize_limit", sv)); err != nil {
		if !fortiAPIPatch(o["oversize-limit"]) {
			return fmt.Errorf("Error reading oversize-limit: %v", err)
		}
	}

	if err = d.Set("av_virus_log", flattenSecurityAntivirusProfileAvVirusLog(o["av-virus-log"], d, "av_virus_log", sv)); err != nil {
		if !fortiAPIPatch(o["av-virus-log"]) {
			return fmt.Errorf("Error reading av-virus-log: %v", err)
		}
	}

	if err = d.Set("analytics_max_upload", flattenSecurityAntivirusProfileAnalyticsMaxUpload(o["analytics-max-upload"], d, "analytics_max_upload", sv)); err != nil {
		if !fortiAPIPatch(o["analytics-max-upload"]) {
			return fmt.Errorf("Error reading analytics-max-upload: %v", err)
		}
	}

	if err = d.Set("uncomp_size_limit", flattenSecurityAntivirusProfileUncompSizeLimit(o["uncomp-size-limit"], d, "uncomp_size_limit", sv)); err != nil {
		if !fortiAPIPatch(o["uncomp-size-limit"]) {
			return fmt.Errorf("Error reading uncomp-size-limit: %v", err)
		}
	}

	if err = d.Set("oversize", flattenSecurityAntivirusProfileOversize(o["oversize"], d, "oversize", sv)); err != nil {
		if !fortiAPIPatch(o["oversize"]) {
			return fmt.Errorf("Error reading oversize: %v", err)
		}
	}

	if err = d.Set("ftgd_analytics", flattenSecurityAntivirusProfileFtgdAnalytics(o["ftgd-analytics"], d, "ftgd_analytics", sv)); err != nil {
		if !fortiAPIPatch(o["ftgd-analytics"]) {
			return fmt.Errorf("Error reading ftgd-analytics: %v", err)
		}
	}

	if err = d.Set("options", flattenSecurityAntivirusProfileOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSecurityAntivirusProfileMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSecurityAntivirusProfileUncompNestLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileEmulator(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileScanBzip2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileAvBlockLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileAnalyticsDb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileOversizeLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileAvVirusLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileAnalyticsMaxUpload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileUncompSizeLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileOversize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileFtgdAnalytics(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSecurityAntivirusProfileMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityAntivirusProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("uncomp_nest_limit"); ok {
		t, err := expandSecurityAntivirusProfileUncompNestLimit(d, v, "uncomp_nest_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncomp-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("streaming_content_bypass"); ok {
		t, err := expandSecurityAntivirusProfileStreamingContentBypass(d, v, "streaming_content_bypass", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-content-bypass"] = t
		}
	}

	if v, ok := d.GetOk("emulator"); ok {
		t, err := expandSecurityAntivirusProfileEmulator(d, v, "emulator", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emulator"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok {
		t, err := expandSecurityAntivirusProfileScanBzip2(d, v, "scan_bzip2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("av_block_log"); ok {
		t, err := expandSecurityAntivirusProfileAvBlockLog(d, v, "av_block_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-block-log"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSecurityAntivirusProfileComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("analytics_db"); ok {
		t, err := expandSecurityAntivirusProfileAnalyticsDb(d, v, "analytics_db", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-db"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok {
		t, err := expandSecurityAntivirusProfileOversizeLimit(d, v, "oversize_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("av_virus_log"); ok {
		t, err := expandSecurityAntivirusProfileAvVirusLog(d, v, "av_virus_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-virus-log"] = t
		}
	}

	if v, ok := d.GetOk("analytics_max_upload"); ok {
		t, err := expandSecurityAntivirusProfileAnalyticsMaxUpload(d, v, "analytics_max_upload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-max-upload"] = t
		}
	}

	if v, ok := d.GetOk("uncomp_size_limit"); ok {
		t, err := expandSecurityAntivirusProfileUncompSizeLimit(d, v, "uncomp_size_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncomp-size-limit"] = t
		}
	}

	if v, ok := d.GetOk("oversize"); ok {
		t, err := expandSecurityAntivirusProfileOversize(d, v, "oversize", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_analytics"); ok {
		t, err := expandSecurityAntivirusProfileFtgdAnalytics(d, v, "ftgd_analytics", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-analytics"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandSecurityAntivirusProfileOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSecurityAntivirusProfileMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
