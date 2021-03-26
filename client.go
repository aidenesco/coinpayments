package coinpayments

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//ClientOption is an option used to modify a client
type ClientOption func(client *Client)

//OptionalValue is an option used to add values to an api request
type OptionalValue func(values *url.Values)

//Client allows programmatic access to the coinpayments api
type Client struct {
	client     *http.Client
	privateKey string
	publicKey  string
	ipnSecret  string
}

//NewClient returns a new Client with the applied options
func NewClient(publicKey, privateKey string, options ...ClientOption) *Client {
	client := &Client{
		privateKey: privateKey,
		publicKey:  publicKey,
		client:     http.DefaultClient,
	}

	for _, o := range options {
		o(client)
	}
	return client
}

//WithHTTPClient is an option that makes the Client use the provided http client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.client = httpClient
	}
}

//WithIPNSecret is an option that makes the Client use the provided secret
func WithIPNSecret(secret string) ClientOption {
	return func(client *Client) {
		client.ipnSecret = secret
	}
}

//WithOptionalValue is an option that adds values to an api request
func WithOptionalValue(key, value string) OptionalValue {
	return func(values *url.Values) {
		values.Set(key, value)
	}
}

func (c *Client) call(cmd string, values *url.Values, response interface{}) error {

	values.Add("key", c.publicKey)
	values.Add("version", apiVersion)
	values.Add("cmd", cmd)
	values.Add("format", apiFormat)

	sData := values.Encode()

	dataHMAC, err := c.makeHMAC(sData)
	if err != nil {
		return fmt.Errorf("coinpayments: error making HMAC - %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(sData))
	if err != nil {
		return fmt.Errorf("coinpayments: error making api request - %v", err)
	}

	req.Header.Add("HMAC", dataHMAC)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(sData)))
	req.Header.Add("User-Agent", "github.com/aidenesco/coinpayments")
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("coinpayments: error doing api request - %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("coinpayments: api call returned unexpected status: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("coinpayments: error reading api response body - %v", err)
	}

	var errResp errResponse

	if err := json.Unmarshal(body, &errResp); err != nil {
		return fmt.Errorf("coinpayments: error unmarshaling api error response - %v", err)
	}

	if errResp.Error != apiSuccess {
		return fmt.Errorf("coinpayments: api error - %v", errResp.Error)
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("coinpayments: error unmarshaling response json - %v", err)
	}

	return nil
}

func (c *Client) makeHMAC(data string) (string, error) {
	hash := hmac.New(sha512.New, []byte(c.privateKey))
	if _, err := hash.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (c *Client) makeIPNHMAC(data string) (string, error) {
	hash := hmac.New(sha512.New, []byte(c.ipnSecret))
	if _, err := hash.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func addOptionals(opts []OptionalValue, values *url.Values) {
	for _, v := range opts {
		v(values)
	}
}

type errResponse struct {
	Error string `json:"error"`
}
