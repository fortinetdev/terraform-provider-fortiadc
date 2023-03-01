// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system sdn connector.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSdnConnectorRead,
		Update: resourceSystemSdnConnectorUpdate,
		Create: resourceSystemSdnConnectorCreate,
		Delete: resourceSystemSdnConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_region_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sap_ms_http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aws_secretkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sap_icm_http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secret_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"aws_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"aws_accesskey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sap_dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_ha_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sap_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sap_sidadm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oci_compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSystemSdnConnectorCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSdnConnector: type error")
	}

	obj, err := getObjectSystemSdnConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnector resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdnConnector(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnector resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSdnConnectorRead(d, m)
}
func resourceSystemSdnConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemSdnConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnector resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdnConnector(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnector resource: %v", err)
	}

	return resourceSystemSdnConnectorRead(d, m)
}
func resourceSystemSdnConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemSdnConnector(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemSdnConnector(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnector(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnector resource from API: %v", err)
	}
	return nil
}
func flattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSapMsHttpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorAwsSecretkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSapIcmHttpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciTenantId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorAwsRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciUserId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorAwsAccesskey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSapDnsSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciHaStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSapPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSapSidadm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciCompartmentId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemSdnConnectorStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("oci_region_type", flattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type", sv)); err != nil {
		if !fortiAPIPatch(o["oci-region-type"]) {
			return fmt.Errorf("Error reading oci-region-type: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSdnConnectorPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("oci_cert", flattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert", sv)); err != nil {
		if !fortiAPIPatch(o["oci-cert"]) {
			return fmt.Errorf("Error reading oci-cert: %v", err)
		}
	}

	if err = d.Set("sap_ms_http_port", flattenSystemSdnConnectorSapMsHttpPort(o["sap-ms-http-port"], d, "sap_ms_http_port", sv)); err != nil {
		if !fortiAPIPatch(o["sap-ms-http-port"]) {
			return fmt.Errorf("Error reading sap-ms-http-port: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemSdnConnectorMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("aws_secretkey", flattenSystemSdnConnectorAwsSecretkey(o["aws_secretkey"], d, "aws_secretkey", sv)); err != nil {
		if !fortiAPIPatch(o["aws_secretkey"]) {
			return fmt.Errorf("Error reading aws_secretkey: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemSdnConnectorUpdateInterval(o["update_interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update_interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSdnConnectorType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("sap_icm_http_port", flattenSystemSdnConnectorSapIcmHttpPort(o["sap-icm-http-port"], d, "sap_icm_http_port", sv)); err != nil {
		if !fortiAPIPatch(o["sap-icm-http-port"]) {
			return fmt.Errorf("Error reading sap-icm-http-port: %v", err)
		}
	}

	if err = d.Set("secret_token", flattenSystemSdnConnectorSecretToken(o["secret_token"], d, "secret_token", sv)); err != nil {
		if !fortiAPIPatch(o["secret_token"]) {
			return fmt.Errorf("Error reading secret_token: %v", err)
		}
	}

	if err = d.Set("oci_tenant_id", flattenSystemSdnConnectorOciTenantId(o["oci-tenant-id"], d, "oci_tenant_id", sv)); err != nil {
		if !fortiAPIPatch(o["oci-tenant-id"]) {
			return fmt.Errorf("Error reading oci-tenant-id: %v", err)
		}
	}

	if err = d.Set("oci_region", flattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region", sv)); err != nil {
		if !fortiAPIPatch(o["oci-region"]) {
			return fmt.Errorf("Error reading oci-region: %v", err)
		}
	}

	if err = d.Set("aws_region", flattenSystemSdnConnectorAwsRegion(o["aws_region"], d, "aws_region", sv)); err != nil {
		if !fortiAPIPatch(o["aws_region"]) {
			return fmt.Errorf("Error reading aws_region: %v", err)
		}
	}

	if err = d.Set("oci_user_id", flattenSystemSdnConnectorOciUserId(o["oci-user-id"], d, "oci_user_id", sv)); err != nil {
		if !fortiAPIPatch(o["oci-user-id"]) {
			return fmt.Errorf("Error reading oci-user-id: %v", err)
		}
	}

	if err = d.Set("aws_accesskey", flattenSystemSdnConnectorAwsAccesskey(o["aws_accesskey"], d, "aws_accesskey", sv)); err != nil {
		if !fortiAPIPatch(o["aws_accesskey"]) {
			return fmt.Errorf("Error reading aws_accesskey: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", flattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam", sv)); err != nil {
		if !fortiAPIPatch(o["use-metadata-iam"]) {
			return fmt.Errorf("Error reading use-metadata-iam: %v", err)
		}
	}

	if err = d.Set("sap_dns_suffix", flattenSystemSdnConnectorSapDnsSuffix(o["sap-dns-suffix"], d, "sap_dns_suffix", sv)); err != nil {
		if !fortiAPIPatch(o["sap-dns-suffix"]) {
			return fmt.Errorf("Error reading sap-dns-suffix: %v", err)
		}
	}

	if err = d.Set("oci_ha_status", flattenSystemSdnConnectorOciHaStatus(o["oci-ha-status"], d, "oci_ha_status", sv)); err != nil {
		if !fortiAPIPatch(o["oci-ha-status"]) {
			return fmt.Errorf("Error reading oci-ha-status: %v", err)
		}
	}

	if err = d.Set("sap_password", flattenSystemSdnConnectorSapPassword(o["sap-password"], d, "sap_password", sv)); err != nil {
		if !fortiAPIPatch(o["sap-password"]) {
			return fmt.Errorf("Error reading sap-password: %v", err)
		}
	}

	if err = d.Set("sap_sidadm", flattenSystemSdnConnectorSapSidadm(o["sap-sidadm"], d, "sap_sidadm", sv)); err != nil {
		if !fortiAPIPatch(o["sap-sidadm"]) {
			return fmt.Errorf("Error reading sap-sidadm: %v", err)
		}
	}

	if err = d.Set("oci_compartment_id", flattenSystemSdnConnectorOciCompartmentId(o["oci-compartment-id"], d, "oci_compartment_id", sv)); err != nil {
		if !fortiAPIPatch(o["oci-compartment-id"]) {
			return fmt.Errorf("Error reading oci-compartment-id: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSdnConnectorServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	return nil
}

func expandSystemSdnConnectorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSapMsHttpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAwsSecretkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSapIcmHttpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSecretToken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciTenantId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAwsRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciUserId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAwsAccesskey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUseMetadataIam(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSapDnsSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciHaStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSapSidadm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciCompartmentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnector(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSdnConnectorStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("oci_region_type"); ok {
		t, err := expandSystemSdnConnectorOciRegionType(d, v, "oci_region_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-type"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemSdnConnectorPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("oci_cert"); ok {
		t, err := expandSystemSdnConnectorOciCert(d, v, "oci_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-cert"] = t
		}
	}

	if v, ok := d.GetOk("sap_ms_http_port"); ok {
		t, err := expandSystemSdnConnectorSapMsHttpPort(d, v, "sap_ms_http_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sap-ms-http-port"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSdnConnectorMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("aws_secretkey"); ok {
		t, err := expandSystemSdnConnectorAwsSecretkey(d, v, "aws_secretkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws_secretkey"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {
		t, err := expandSystemSdnConnectorUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update_interval"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemSdnConnectorType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("sap_icm_http_port"); ok {
		t, err := expandSystemSdnConnectorSapIcmHttpPort(d, v, "sap_icm_http_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sap-icm-http-port"] = t
		}
	}

	if v, ok := d.GetOk("secret_token"); ok {
		t, err := expandSystemSdnConnectorSecretToken(d, v, "secret_token", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret_token"] = t
		}
	}

	if v, ok := d.GetOk("oci_tenant_id"); ok {
		t, err := expandSystemSdnConnectorOciTenantId(d, v, "oci_tenant_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-tenant-id"] = t
		}
	}

	if v, ok := d.GetOk("oci_region"); ok {
		t, err := expandSystemSdnConnectorOciRegion(d, v, "oci_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region"] = t
		}
	}

	if v, ok := d.GetOk("aws_region"); ok {
		t, err := expandSystemSdnConnectorAwsRegion(d, v, "aws_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws_region"] = t
		}
	}

	if v, ok := d.GetOk("oci_user_id"); ok {
		t, err := expandSystemSdnConnectorOciUserId(d, v, "oci_user_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-user-id"] = t
		}
	}

	if v, ok := d.GetOk("aws_accesskey"); ok {
		t, err := expandSystemSdnConnectorAwsAccesskey(d, v, "aws_accesskey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws_accesskey"] = t
		}
	}

	if v, ok := d.GetOk("use_metadata_iam"); ok {
		t, err := expandSystemSdnConnectorUseMetadataIam(d, v, "use_metadata_iam", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-metadata-iam"] = t
		}
	}

	if v, ok := d.GetOk("sap_dns_suffix"); ok {
		t, err := expandSystemSdnConnectorSapDnsSuffix(d, v, "sap_dns_suffix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sap-dns-suffix"] = t
		}
	}

	if v, ok := d.GetOk("oci_ha_status"); ok {
		t, err := expandSystemSdnConnectorOciHaStatus(d, v, "oci_ha_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-ha-status"] = t
		}
	}

	if v, ok := d.GetOk("sap_password"); ok {
		t, err := expandSystemSdnConnectorSapPassword(d, v, "sap_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sap-password"] = t
		}
	}

	if v, ok := d.GetOk("sap_sidadm"); ok {
		t, err := expandSystemSdnConnectorSapSidadm(d, v, "sap_sidadm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sap-sidadm"] = t
		}
	}

	if v, ok := d.GetOk("oci_compartment_id"); ok {
		t, err := expandSystemSdnConnectorOciCompartmentId(d, v, "oci_compartment_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-compartment-id"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemSdnConnectorServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	return &obj, nil
}
