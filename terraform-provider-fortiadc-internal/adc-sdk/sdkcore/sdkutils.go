package forticlient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func createUpload(c *FortiSDKClient, method string, path string, params *map[string]interface{}, output map[string]interface{}, vdomparam string) (err error) {
	tmp := *params
	if _, ok := tmp["data"]; !ok {
		err = fmt.Errorf("no data.")
		return
	}
	if _, ok := tmp["content_type"]; !ok {
		err = fmt.Errorf("no content_type.")
		return
	}
	data := string(tmp["data"].([]uint8))
	bytes := bytes.NewBuffer([]byte(data))

	req := c.NewRequest(method, path, nil, bytes)
	err = req.Send4(vdomparam, tmp["content_type"].(string))
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output["vdom"] = result["vdom"]
		}

		if result["mkey"] != nil {
			output["mkey"] = result["mkey"]
		}

		if _, ok := result["payload"].(float64); ok {
			ret_code := result["payload"].(float64)
			if ret_code != 0 {
				str_code := strconv.FormatInt(int64(ret_code), 10)
				err = errors.New(fadc_err[str_code])
			}
		}
	}

	return
}

func createUpdate(c *FortiSDKClient, method string, path string, params *map[string]interface{}, output map[string]interface{}, vdomparam string) (err error) {
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(method, path, nil, bytes)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output["vdom"] = result["vdom"]
		}

		if result["mkey"] != nil {
			output["mkey"] = result["mkey"]
		}

		if _, ok := result["payload"].(float64); ok {
			ret_code := result["payload"].(float64)
			if ret_code != 0 {
				str_code := strconv.FormatInt(int64(ret_code), 10)
				err = errors.New(fadc_err[str_code])
			}
		}
	}

	return
}

func delete(c *FortiSDKClient, method string, path string, vdomparam string) (err error) {

	req := c.NewRequest(method, path, nil, nil)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func read(c *FortiSDKClient, method string, path string, bcomplex bool, vdomparam string) (mapTmp map[string]interface{}, err error) {
	req := c.NewRequest(method, path, nil, nil)
	if vdomparam == "global" {
		vdomparam = ""
	}
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	log.Printf("path %s", path)
	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		switch result["payload"].(type) {
		case map[string]interface{}:
			mapTmp = result["payload"].(map[string]interface{})
			for k, v := range mapTmp {
				switch v.(type) {
				case string:
					mapTmp[k] = strings.TrimSuffix(v.(string), " ")
				}
			}
		default:
			mapTmp = (result["payload"].([]interface{}))[0].(map[string]interface{})
		}
	}

	return
}
