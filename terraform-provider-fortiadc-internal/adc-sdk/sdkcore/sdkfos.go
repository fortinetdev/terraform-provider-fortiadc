// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Description: SDK for FortiOS Provider

package forticlient

type creatUpdateOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

func (c *FortiSDKClient) ReadSystemScriptingUpload(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_scripting"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteSystemScriptingUpload(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_scripting"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateSystemScriptingUpload(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/system_scripting/upload"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) ReadUploadCaptchaCust(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_captcha_profile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteUploadCaptchaCust(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_captcha_profile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateUploadCaptchaCust(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/upload/captcha_cust"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) ReadUploadErrorPage(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/load_balance_error_page"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteUploadErrorPage(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/load_balance_error_page"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateUploadErrorPage(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/upload/error_page"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) ReadLocalCertUpload(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_local"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteLocalCertUpload(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_local"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateLocalCertUpload(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/upload/certificate_local"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) ReadCertificateIntmedCaUpload(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_intermediate_ca"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteCertificateIntmedCaUpload(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_intermediate_ca"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateCertificateIntmedCaUpload(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/upload/certificate_intmed_ca"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) ReadCaCertUpload(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_ca"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteCaCertUpload(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_ca"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateCaCertUpload(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/upload/certificate_ca"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) ReadCertificateCrlUpload(mkey string, vdomparam string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_certificate_crl"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdomparam)
	return
}

func (c *FortiSDKClient) DeleteCertificateCrlUpload(mkey string, vdomparam string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/system_certificate_crl"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdomparam)
	return
}

func (c *FortiSDKClient) CreateCertificateCrlUpload(params *map[string]interface{}, vdomparam string) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/upload/certificate_crl"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdomparam)
	return
}

func (c *FortiSDKClient) UpdateSystemHaMgmtUpdate(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/system_ha"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) CreateVdom(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/vdom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteVdom(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/vdom"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateVdom(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/vdom"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadVdom(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/system_vdom"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}
