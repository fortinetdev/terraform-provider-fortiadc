package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"errors"
	"github.com/terraform-providers/terraform-provider-fortiadc/adc-sdk/auth"
	"github.com/terraform-providers/terraform-provider-fortiadc/adc-sdk/config"
	"github.com/terraform-providers/terraform-provider-fortiadc/adc-sdk/request"
	"strconv"
)

// MultValue describes the nested structure in the results
type MultValue struct {
	Name string `json:"name"`
}

// MultValues describes the nested structure in the results
type MultValues []MultValue

// FortiSDKClient describes the global FortiOS plugin client instance
type FortiSDKClient struct {
	Config  config.Config
	Retries int
	Fv      string
}

// ExtractString extracts strings from result and put them into a string array,
// and return the string array
func ExtractString(members []MultValue) []string {
	vs := make([]string, 0, len(members))
	for _, v := range members {
		c := v.Name
		vs = append(vs, c)
	}
	return vs
}

// EscapeURLString escapes the string so it can be safely placed inside a URL query
func EscapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) (*FortiSDKClient, error) {
	c := &FortiSDKClient{}

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	vsave := client.Timeout
	client.Timeout = time.Second * 30
	v, err := c.GetDeviceVersion()
	if err != nil {
		//return nil, fmt.Errorf("%s", err)
		c.Fv = "Unknown (fortiadc non-global user can't get version)"
	} else {
		c.Fv = v
	}
	client.Timeout = vsave

	return c, nil
}

// NewRequest creates the request to FortiOS for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
}

// GetDeviceVersion gets the version of FortiOS
// It returns version as string
func (c *FortiSDKClient) GetDeviceVersion() (version string, err error) {
	HTTPMethod := "GET"

	path := "/api/platform/version"
	if c.Config.Auth.Vdom != "" {
		path = fmt.Sprintf("/api/platform/version?vdom=%s", c.Config.Auth.Vdom)
	}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send2(2, true)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request, %s", err)
		return "", err
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body, %s", err)
		return "", err
	}

	var result map[string]interface{}
	var payload map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["payload"] == nil {
			err = fmt.Errorf("no payload in response")
			return "", err
		} else {
			switch v := result["payload"].(type) {
			case map[string]interface{}:
				payload = result["payload"].(map[string]interface{})
			default:
				err = fmt.Errorf("no version in response1 %v", v)
				return "", err
			}
		}

		if payload["version"] == nil {
			err = fmt.Errorf("no version in response")
			return "", err
		}

		return payload["version"].(string), err
	}

	return "", err
}

func fortiAPIHttpStatus404Checking(result map[string]interface{}) (b404 bool) {
	b404 = false

	if result != nil {
		if result["http_status"] != nil && result["http_status"] == 404.0 {
			b404 = true
			return
		}
	}

	return
}

func fortiAPIErrorFormat(result map[string]interface{}, body string) (err error) {
	if result != nil {
		if result["payload"] != nil {
			err = nil
			if _, ok := result["payload"].(float64); ok {
				ret_code := result["payload"].(float64)
				if ret_code != 0 {
					str_code := strconv.FormatInt(int64(ret_code), 10)
					err = errors.New(fadc_err[str_code])
				}
			}
			return
		}
		err = fmt.Errorf("\n%v", body)
		return
	}

	// Authorization Required, etc. | Attention: scalable here
	err = fmt.Errorf("\n%v", body)
	return
}

//Build input data by sdk
