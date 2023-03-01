// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUploadErrorPage() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUploadErrorPageRead,
		Update: resourceUploadErrorPageUpdate,
		Create: resourceUploadErrorPageCreate,
		Delete: resourceUploadErrorPageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"errorpagefile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
func resourceUploadErrorPageCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UploadErrorPage: type error")
	}

	obj, err := getObjectUploadErrorPage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UploadErrorPage resource while getting object: %v", err)
	}

	_, err = c.CreateUploadErrorPage(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating UploadErrorPage resource: %v", err)
	}

	d.SetId(mkey)

	//return resourceUploadErrorPageRead(d, m)
	return nil
}
func resourceUploadErrorPageUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUploadErrorPage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClientSslProfile resource while getting object: %v", err)
	}

	_, err = c.CreateUploadErrorPage(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceClientSslProfile resource: %v", err)
	}

	return resourceUploadErrorPageRead(d, m)
}
func resourceUploadErrorPageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUploadErrorPage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UploadErrorPage resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceUploadErrorPageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUploadErrorPage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UploadErrorPage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUploadErrorPage(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UploadErrorPage resource from API: %v", err)
	}
	return nil
}
func flattenUploadErrorPageSubject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageValidfrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageValidto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageIssuer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageiStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageFingerprint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUploadErrorPageDest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenUploadErrorPageMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUploadErrorPage(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenUploadErrorPageMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUploadErrorPageVpath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadErrorPageUpdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadErrorPageErrorPageFile(d *schema.ResourceData, v interface{}, pre string, sv string, writer *multipart.Writer) (interface{}, error) {
	return v, nil
}

func expandUploadErrorPageMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUploadErrorPageVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUploadErrorPage(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	body := &bytes.Buffer{}
	obj := make(map[string]interface{})
	writer := multipart.NewWriter(body)

	if v, ok := d.GetOk("vpath"); ok {
		t, err := expandUploadErrorPageVpath(d, v, "vpath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "vpath", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	if v, ok := d.GetOk("update"); ok {
		t, err := expandUploadErrorPageUpdate(d, v, "update", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "update", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("errorpagefile"); ok {
		t, err := expandUploadErrorPageErrorPageFile(d, v, "errorpagefile", sv, writer)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "errorPageFile", t.(string), true)
			if err != nil {
				writer.Close()
				return &obj, err
			}
			filename := filepath.Base(t.(string))
			set_multipart_parameter(writer, "local_pc", filename, false)
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUploadErrorPageMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "mkey", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandUploadErrorPageVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "vdom", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	writer.Close()
	obj["content_type"] = writer.FormDataContentType()
	obj["data"] = body.Bytes()

	return &obj, nil
}
