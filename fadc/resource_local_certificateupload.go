// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocalCertUpload() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLocalCertUploadRead,
		Update: resourceLocalCertUploadUpdate,
		Create: resourceLocalCertUploadCreate,
		Delete: resourceLocalCertUploadDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
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
			/*"subject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ca_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"validfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"validto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},*/
		},
	}
}
func resourceLocalCertUploadCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing LocalCertUpload: type error")
	}

	obj, err := getObjectLocalCertUpload(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LocalCertUpload resource while getting object: %v", err)
	}

	_, err = c.CreateLocalCertUpload(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating LocalCertUpload resource: %v", err)
	}

	d.SetId(mkey)

	//return resourceLocalCertUploadRead(d, m)
	return nil
}
func resourceLocalCertUploadUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceLocalCertUploadRead(d, m)
}
func resourceLocalCertUploadDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLocalCertUpload(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LocalCertUpload resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceLocalCertUploadRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLocalCertUpload(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LocalCertUpload resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLocalCertUpload(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LocalCertUpload resource from API: %v", err)
	}
	return nil
}
func flattenLocalCertUploadSubject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadValidfrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadValidto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadIssuer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadiStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadiCAtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadFingerprint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLocalCertUploadDest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenLocalCertUploadMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLocalCertUpload(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	/*if err = d.Set("subject", flattenLocalCertUploadSubject(o["subject"], d, "subject", sv)); err != nil {
		if !fortiAPIPatch(o["subject"]) {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("validfrom", flattenLocalCertUploadValidfrom(o["validfrom"], d, "validfrom", sv)); err != nil {
		if !fortiAPIPatch(o["validfrom"]) {
			return fmt.Errorf("Error reading validfrom: %v", err)
		}
	}

	if err = d.Set("validto", flattenLocalCertUploadValidto(o["validto"], d, "validto", sv)); err != nil {
		if !fortiAPIPatch(o["validto"]) {
			return fmt.Errorf("Error reading validto: %v", err)
		}
	}

	if err = d.Set("type", flattenLocalCertUploadType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ca_type", flattenLocalCertUploadiCAtype(o["ca_type"], d, "ca_type", sv)); err != nil {
		if !fortiAPIPatch(o["ca_type"]) {
			return fmt.Errorf("Error reading ca_type: %v", err)
		}
	}

	if err = d.Set("issuer", flattenLocalCertUploadIssuer(o["issuer"], d, "issuer", sv)); err != nil {
		if !fortiAPIPatch(o["issuer"]) {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("fingerprint", flattenLocalCertUploadFingerprint(o["fingerprint"], d, "fingerprint", sv)); err != nil {
		if !fortiAPIPatch(o["fingerprint"]) {
			return fmt.Errorf("Error reading fingerprint: %v", err)
		}
	}*/

	if err = d.Set("mkey", flattenLocalCertUploadMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandLocalCertUploadType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLocalCertUploadUpload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLocalCertUploadPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLocalCertUploadCert(d *schema.ResourceData, v interface{}, pre string, sv string, writer *multipart.Writer) (interface{}, error) {
	return v, nil
}

func expandLocalCertUploadKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLocalCertUploadMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLocalCertUploadVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLocalCertUpload(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	body := &bytes.Buffer{}
	obj := make(map[string]interface{})
	writer := multipart.NewWriter(body)

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLocalCertUploadType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "type", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	if v, ok := d.GetOk("upload"); ok {
		t, err := expandLocalCertUploadUpload(d, v, "upload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "upload", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	if v, ok := d.GetOk("passwd"); ok {
		t, err := expandLocalCertUploadPasswd(d, v, "passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "passwd", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandLocalCertUploadCert(d, v, "cert", sv, writer)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "cert", t.(string), true)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandLocalCertUploadKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "key", t.(string), true)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandLocalCertUploadMkey(d, v, "mkey", sv)
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
		t, err := expandLocalCertUploadVdom(d, v, "vdom", sv)
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
