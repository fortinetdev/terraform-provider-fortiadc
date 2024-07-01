package fortiadc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fortinetdev/terraform-provider-fortiadc/adc-sdk/auth"
	//forticlient "github.com/fortinetdev/forti-sdk-go/fortiadc/sdkcore"
	forticlient "github.com/fortinetdev/terraform-provider-fortiadc/adc-sdk/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Hostname        string
	Token           string
	Insecure        *bool
	CABundle        string
	CABundleContent string
	Vdom            string
	HTTPProxy       string
	HTTPSProxy      string
	NOProxy         string

	PeerAuth   string
	CaCert     string
	ClientCert string
	ClientKey  string
}

// FortiClient contains the basic FortiOS SDK connection information to FortiOS
// It can be used to as a client of FortiOS for the plugin
// Now FortiClient contains two kinds of clients:
// Client is for FortiGate
// Client Fottimanager is for FortiManager
type FortiClient struct {
	//to sdk client
	Client *forticlient.FortiSDKClient
}

func set_log() {
	f, err := os.OpenFile("/tmp/adc", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	//set_log()
	bFOSExist := bFortiOSHostnameExist(c)

	if bFOSExist {
		err := createFortiOSClient(&fClient, c)
		if err != nil {
			return nil, fmt.Errorf("Error create fortios client: %v", err)
		}
	}

	if !bFOSExist {
		return nil, fmt.Errorf("FortiADC at least one of their hostnames should be set")
	}

	return &fClient, nil
}

func bFortiOSHostnameExist(c *Config) bool {
	if c.Hostname == "" {
		if os.Getenv("FORTIADC_ACCESS_HOSTNAME") == "" {
			return false
		}
	}

	return true
}

func createFortiOSClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Hostname, c.Token, c.CABundle, c.CABundleContent, c.PeerAuth, c.CaCert, c.ClientCert, c.ClientKey, c.Vdom, c.HTTPProxy, c.HTTPSProxy, c.NOProxy)

	if auth.Hostname == "" {
		_, err := auth.GetEnvHostname()
		if err != nil {
			return fmt.Errorf("Error reading Hostname")
		}
	}

	if auth.Token == "" {
		_, err := auth.GetEnvToken()
		if err != nil {
			return fmt.Errorf("Error reading Token")
		}
	}

	if auth.CABundle == "" {
		auth.GetEnvCABundle()
	}

	if auth.PeerAuth == "" {
		_, err := auth.GetEnvPeerAuth()
		if err != nil {
			return fmt.Errorf("Error reading PeerAuth")
		}
	}
	if auth.CaCert == "" {
		_, err := auth.GetEnvCaCert()
		if err != nil {
			return fmt.Errorf("Error reading CaCert")
		}
	}
	if auth.ClientCert == "" {
		_, err := auth.GetEnvClientCert()
		if err != nil {
			return fmt.Errorf("Error reading ClientCert")
		}
	}
	if auth.ClientKey == "" {
		_, err := auth.GetEnvClientKey()
		if err != nil {
			return fmt.Errorf("Error reading ClientKey")
		}
	}
	if auth.HTTPProxy == "" {
		_, err := auth.GetEnvHTTPProxy()
		if err != nil {
			return fmt.Errorf("Error reading HTTP proxy")
		}
	}
	if auth.HTTPSProxy == "" {
		_, err := auth.GetEnvHTTPSProxy()
		if err != nil {
			return fmt.Errorf("Error reading HTTPS proxy")
		}
	}
	if auth.NOProxy == "" {
		_, err := auth.GetEnvNOProxy()
		if err != nil {
			return fmt.Errorf("Error reading NO proxy")
		}
	}

	pool := x509.NewCertPool()

	if auth.CABundle != "" {
		if auth.CABundleContent != "" {
			return fmt.Errorf("\"cabundlefile\" and \"cabundlecontent\" could not exist at the same time! Please only configure one of them. If you are not configure \"cabundlefile\", please check the environment variable \"FORTIADC_CA_CABUNDLE\".")
		}

		f, err := os.Open(auth.CABundle)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %v", err)
		}
		defer f.Close()

		caBundle, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %v", err)
		}

		if !pool.AppendCertsFromPEM([]byte(caBundle)) {
			return fmt.Errorf("Error reading CA Bundle")
		}
		config.RootCAs = pool
	} else if auth.CABundleContent != "" {
		if !pool.AppendCertsFromPEM([]byte(auth.CABundleContent)) {
			return fmt.Errorf("Error reading CA Bundle")
		}
		config.RootCAs = pool
	}

	if auth.PeerAuth == "enable" {
		if auth.CaCert != "" {
			caCertFile := auth.CaCert
			caCert, err := ioutil.ReadFile(caCertFile)
			if err != nil {
				return fmt.Errorf("client ioutil.ReadFile couldn't load cacert file: %v", err)
			}

			pool.AppendCertsFromPEM(caCert)
		}

		if auth.ClientCert == "" {
			return fmt.Errorf("User Cert file doesn't exist!")
		}

		if auth.ClientKey == "" {
			return fmt.Errorf("User Key file doesn't exist!")
		}

		clientCert, err := tls.LoadX509KeyPair(auth.ClientCert, auth.ClientKey)
		if err != nil {
			return fmt.Errorf("Client ioutil.ReadFile couldn't load clientCert/clientKey file: %v", err)
		}

		config.Certificates = []tls.Certificate{clientCert}
	}

	if c.Insecure == nil {
		// defaut value for Insecure is false
		b, _ := auth.GetEnvInsecure()
		config.InsecureSkipVerify = b
	} else {
		config.InsecureSkipVerify = *c.Insecure
	}

	if config.InsecureSkipVerify == false && auth.CABundle == "" && auth.CABundleContent == "" {
		return fmt.Errorf("Error getting CA Bundle, CA Bundle should be set when insecure is false")
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	disable_proxy := false

	if auth.NOProxy != "" {
		var noproxy *url.URL
		noproxy, err2 := url.Parse(auth.NOProxy)
		if err2 != nil {
			fmt.Errorf("%w", err2)
		}
		noproxy_slice := strings.Split(noproxy.String(), ",")
		for index, line := range noproxy_slice {
			if line == "*" || line == c.Hostname {
				log.Printf("[LOG] no_proxy match: %w %s ", index, line)
				disable_proxy = true
			}
		}
	}

	if (auth.HTTPProxy != "" || auth.HTTPSProxy != "") && disable_proxy == false {
		tr = &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: config,
		}
	}

	if auth.HTTPProxy != "" && disable_proxy == false {
		var httpProxy *url.URL
		httpProxy, err := url.Parse(auth.HTTPProxy)
		if err != nil {
			return fmt.Errorf("Error parsing HTTP Proxy: %w", err)
		}
		tr.Proxy = http.ProxyURL(httpProxy)
	}

	if auth.HTTPSProxy != "" && disable_proxy == false {
		var httpsProxy *url.URL
		httpsProxy, err := url.Parse(auth.HTTPSProxy)
		if err != nil {
			return fmt.Errorf("Error parsing HTTPS Proxy: %w", err)
		}
		tr.Proxy = http.ProxyURL(httpsProxy)
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 250,
	}

	fc, err := forticlient.NewClient(auth, client)

	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}

	fClient.Client = fc

	return nil
}
