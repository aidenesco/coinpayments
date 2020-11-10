package coinpayments

import "net/url"

//getBasicInfoResponse is the api response of a "get_basic_info" call
type getBasicInfoResponse struct {
	Username   string `json:"username"`
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
	PublicName string `json:"public_name"`
}

//GetBasicInfo calls the "get_basic_info" command
func (c *Client) GetBasicInfo(optionals ...OptionalValue) (*getBasicInfoResponse, error) {
	values := &url.Values{}
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getBasicInfoResponse `json:"result"`
	}
	if err := c.call("get_basic_info", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//ratesResponse is the api response of a "rates" call
type ratesResponse map[string]struct {
	IsFiat       int      `json:"is_fiat"`
	RateBTC      string   `json:"rate_btc"`
	LastUpdate   string   `json:"last_update"`
	TxFee        string   `json:"tx_fee"`
	Status       string   `json:"status"`
	Name         string   `json:"name"`
	Confirms     string   `json:"confirms"`
	Capabilities []string `json:"capabilities"`
	Accepted     int      `json:"accepted"`
}

//Rates calls the "rates" command
func (c *Client) Rates(optionals ...OptionalValue) (*ratesResponse, error) {
	values := &url.Values{}
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *ratesResponse `json:"result"`
	}
	if err := c.call("rates", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
