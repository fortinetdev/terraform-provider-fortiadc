// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system sdn connector.

package fortiadc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSdnConnectorRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_region_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sap_ms_http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aws_secretkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sap_icm_http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_accesskey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sap_dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_ha_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sap_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sap_sidadm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnConnector(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error describing SystemSdnConnector: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}
	err = dataSourceRefreshObjectSystemSdnConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSdnConnector from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}
func dataSourceFlattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSapMsHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorMkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAwsSecretkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSapIcmHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAwsRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAwsAccesskey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSapDnsSuffix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciHaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSapSidadm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemSdnConnectorStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("oci_region_type", dataSourceFlattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type")); err != nil {
		if !fortiAPIPatch(o["oci-region-type"]) {
			return fmt.Errorf("Error reading oci-region-type: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemSdnConnectorPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("oci_cert", dataSourceFlattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert")); err != nil {
		if !fortiAPIPatch(o["oci-cert"]) {
			return fmt.Errorf("Error reading oci-cert: %v", err)
		}
	}

	if err = d.Set("sap_ms_http_port", dataSourceFlattenSystemSdnConnectorSapMsHttpPort(o["sap-ms-http-port"], d, "sap_ms_http_port")); err != nil {
		if !fortiAPIPatch(o["sap-ms-http-port"]) {
			return fmt.Errorf("Error reading sap-ms-http-port: %v", err)
		}
	}

	if err = d.Set("mkey", dataSourceFlattenSystemSdnConnectorMkey(o["mkey"], d, "mkey")); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("aws_secretkey", dataSourceFlattenSystemSdnConnectorAwsSecretkey(o["aws_secretkey"], d, "aws_secretkey")); err != nil {
		if !fortiAPIPatch(o["aws_secretkey"]) {
			return fmt.Errorf("Error reading aws_secretkey: %v", err)
		}
	}

	if err = d.Set("update_interval", dataSourceFlattenSystemSdnConnectorUpdateInterval(o["update_interval"], d, "update_interval")); err != nil {
		if !fortiAPIPatch(o["update_interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemSdnConnectorType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("sap_icm_http_port", dataSourceFlattenSystemSdnConnectorSapIcmHttpPort(o["sap-icm-http-port"], d, "sap_icm_http_port")); err != nil {
		if !fortiAPIPatch(o["sap-icm-http-port"]) {
			return fmt.Errorf("Error reading sap-icm-http-port: %v", err)
		}
	}

	if err = d.Set("secret_token", dataSourceFlattenSystemSdnConnectorSecretToken(o["secret_token"], d, "secret_token")); err != nil {
		if !fortiAPIPatch(o["secret_token"]) {
			return fmt.Errorf("Error reading secret_token: %v", err)
		}
	}

	if err = d.Set("oci_tenant_id", dataSourceFlattenSystemSdnConnectorOciTenantId(o["oci-tenant-id"], d, "oci_tenant_id")); err != nil {
		if !fortiAPIPatch(o["oci-tenant-id"]) {
			return fmt.Errorf("Error reading oci-tenant-id: %v", err)
		}
	}

	if err = d.Set("oci_region", dataSourceFlattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region")); err != nil {
		if !fortiAPIPatch(o["oci-region"]) {
			return fmt.Errorf("Error reading oci-region: %v", err)
		}
	}

	if err = d.Set("aws_region", dataSourceFlattenSystemSdnConnectorAwsRegion(o["aws_region"], d, "aws_region")); err != nil {
		if !fortiAPIPatch(o["aws_region"]) {
			return fmt.Errorf("Error reading aws_region: %v", err)
		}
	}

	if err = d.Set("oci_user_id", dataSourceFlattenSystemSdnConnectorOciUserId(o["oci-user-id"], d, "oci_user_id")); err != nil {
		if !fortiAPIPatch(o["oci-user-id"]) {
			return fmt.Errorf("Error reading oci-user-id: %v", err)
		}
	}

	if err = d.Set("aws_accesskey", dataSourceFlattenSystemSdnConnectorAwsAccesskey(o["aws_accesskey"], d, "aws_accesskey")); err != nil {
		if !fortiAPIPatch(o["aws_accesskey"]) {
			return fmt.Errorf("Error reading aws_accesskey: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", dataSourceFlattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam")); err != nil {
		if !fortiAPIPatch(o["use-metadata-iam"]) {
			return fmt.Errorf("Error reading use-metadata-iam: %v", err)
		}
	}

	if err = d.Set("sap_dns_suffix", dataSourceFlattenSystemSdnConnectorSapDnsSuffix(o["sap-dns-suffix"], d, "sap_dns_suffix")); err != nil {
		if !fortiAPIPatch(o["sap-dns-suffix"]) {
			return fmt.Errorf("Error reading sap-dns-suffix: %v", err)
		}
	}

	if err = d.Set("oci_ha_status", dataSourceFlattenSystemSdnConnectorOciHaStatus(o["oci-ha-status"], d, "oci_ha_status")); err != nil {
		if !fortiAPIPatch(o["oci-ha-status"]) {
			return fmt.Errorf("Error reading oci-ha-status: %v", err)
		}
	}

	if err = d.Set("sap_password", dataSourceFlattenSystemSdnConnectorSapPassword(o["sap-password"], d, "sap_password")); err != nil {
		if !fortiAPIPatch(o["sap-password"]) {
			return fmt.Errorf("Error reading sap-password: %v", err)
		}
	}

	if err = d.Set("sap_sidadm", dataSourceFlattenSystemSdnConnectorSapSidadm(o["sap-sidadm"], d, "sap_sidadm")); err != nil {
		if !fortiAPIPatch(o["sap-sidadm"]) {
			return fmt.Errorf("Error reading sap-sidadm: %v", err)
		}
	}

	if err = d.Set("oci_compartment_id", dataSourceFlattenSystemSdnConnectorOciCompartmentId(o["oci-compartment-id"], d, "oci_compartment_id")); err != nil {
		if !fortiAPIPatch(o["oci-compartment-id"]) {
			return fmt.Errorf("Error reading oci-compartment-id: %v", err)
		}
	}

	if err = d.Set("server", dataSourceFlattenSystemSdnConnectorServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	return nil
}
