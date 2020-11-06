package xmrto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client is a xmr.to API client
type Client interface {
	// The order parameter endpoint supplies information about whether new orders can be created.
	// In this case, this endpoint provides the current price, order limits, etc. for newly created orders.
	GetOrderParameters() (*ResponseGetOrderParameters, error)
	// The order creation endpoint allows to create a new order at the current price.
	// The user has to supply a bitcoin destination address and amount to create the order.
	CreateOrder(*RequestCreateOrder) (*ResponseCreateOrder, error)
	// The order status endpoint allows users to query the status of an order, thereby obtaining payment details and order processing progress.
	GetOrderStatus(*RequestGetOrderStatus) (*ResponseGetOrderStatus, error)
	// The order status endpoint allows users to query the recent price of an order by displaying the convertion price of XMR to BTC
	GetOrderPrice(*RequestGetOrderPrice) (*ResponseGetOrderPrice, error)
}

type client struct {
	httpcl *http.Client
	config *Config
}

// New returns a new xmr.to API client
func New(config *Config) Client {
	cfg := &Config{
		APIBaseAddress:         APIBaseAddress,
		APIVersion:             APIVersion,
		APIConversionDirection: APIConversionDirection,
		Testnet:                config.Testnet,
	}
	if len(config.APIBaseAddress) == 0 {
		cfg.APIBaseAddress = APIBaseAddress
	}
	if len(config.APIVersion) == 0 {
		cfg.APIVersion = APIVersion
	}
	if len(config.APIConversionDirection) == 0 {
		cfg.APIConversionDirection = APIConversionDirection
	}

	if cfg.Testnet {
		cfg.APIBaseAddress = APITestnetBaseAddress
	}
	cl := &client{
		config: cfg,
	}

	return cl
}

// Helper function
func (c *client) do(method, path string, in, out interface{}) error {
	var payload []byte
	var err error

	payload = nil
	if in != nil {
		payload, err = json.Marshal(in)
		if err != nil {
			return err
		}
	}
	endpoint, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.config.APIBaseAddress, c.config.APIVersion, c.config.APIConversionDirection))
	if err != nil {
		return nil
	}
	endpoint.Path = fmt.Sprintf("%s/%s/", endpoint.Path, path)

	req, err := http.NewRequest(method, fmt.Sprintf("%s", endpoint), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	httpcl := http.DefaultClient
	resp, err := httpcl.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		xmrtoerror := &APIError{}
		err = json.Unmarshal(body, xmrtoerror)
		if err != nil {
			return err
		}
		return xmrtoerror
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return json.Unmarshal(body, out)
}

// Methods
func (c *client) GetOrderParameters() (resp *ResponseGetOrderParameters, err error) {
	err = c.do("GET", "order_parameter_query", nil, &resp)
	if err != nil {
		return nil, err
	}
	return
}

func (c *client) CreateOrder(req *RequestCreateOrder) (resp *ResponseCreateOrder, err error) {
	err = c.do("POST", "order_create", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

func (c *client) GetOrderStatus(req *RequestGetOrderStatus) (resp *ResponseGetOrderStatus, err error) {
	err = c.do("POST", "order_status_query", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

func (c *client) GetOrderPrice(req *RequestGetOrderPrice) (resp *ResponseGetOrderPrice, err error) {
	err = c.do("POST", "order_check_price", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}
