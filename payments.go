package coinpayments

import "net/url"

//createTransactionResponse is the api response of a "create_transaction" call
type createTransactionResponse struct {
	Amount         string `json:"amount"`
	Address        string `json:"address"`
	DestTag        string `json:"dest_tag"`
	TxnId          string `json:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed"`
	Timeout        int    `json:"timeout"`
	CheckoutURL    string `json:"checkout_url"`
	StatusURL      string `json:"status_url"`
	QRCodeURL      string `json:"qrcode_url"`
}

//CreateTransaction calls the "create_transaction" command
func (c *Client) CreateTransaction(amount, currency1, currency2, buyerEmail string, optionals ...OptionalValue) (*createTransactionResponse, error) {
	values := &url.Values{}
	values.Set("amount", amount)
	values.Set("currency1", currency1)
	values.Set("currency2", currency2)
	values.Set("buyer_email", buyerEmail)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *createTransactionResponse `json:"result"`
	}

	if err := c.call("create_transaction", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getCallbackAddressResponse is the api response of a "get_callback_address" call
type getCallbackAddressResponse struct {
	Address string `json:"address"`
	PubKey  string `json:"pubkey"`
	DestTag string `json:"dest_tag"`
}

//GetCallbackAddress calls the "get_callback_address" command
func (c *Client) GetCallbackAddress(currency string, optionals ...OptionalValue) (*getCallbackAddressResponse, error) {
	values := &url.Values{}
	values.Set("currency", currency)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getCallbackAddressResponse `json:"result"`
	}

	if err := c.call("get_callback_address", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getTxInfoResponse is the api response of a "get_tx_info" call
type getTxInfoResponse struct {
	TimeCreated      int    `json:"time_created"`
	TimeExpires      int    `json:"time_expires"`
	Status           int    `json:"status"`
	StatusText       string `json:"status_text"`
	Type             string `json:"type"`
	Coin             string `json:"coin"`
	Amount           int    `json:"amount"`
	Amountf          string `json:"amountf"`
	Received         int    `json:"received"`
	Receivedf        string `json:"receivedf"`
	ReceivedConfirms int    `json:"recv_confirms"`
	PaymentAddress   string `json:"payment_address"`
	Checkout         struct {
		Currency   string        `json:"currency"`
		Amount     int           `json:"amount"`
		Test       int           `json:"test"`
		ItemNumber string        `json:"item_number"`
		ItemName   string        `json:"item_name"`
		Details    []interface{} `json:"details"`
		Invoice    string        `json:"invoice"`
		Custom     string        `json:"custom"`
		IPNURL     string        `json:"ipn_url"`
		Amountf    int           `json:"amountf"`
	} `json:"checkout,omitempty"`
	Shipping []interface{} `json:"shipping,omitempty"`
}

//GetTxInfo calls the "get_tx_info" command
func (c *Client) GetTxInfo(txid string, optionals ...OptionalValue) (*getTxInfoResponse, error) {
	values := &url.Values{}
	values.Set("txid", txid)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getTxInfoResponse `json:"result"`
	}

	if err := c.call("get_tx_info", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getTxInfoMultiResponse is the api response of a "get_tx_info_multi" call
type getTxInfoMultiResponse map[string]struct {
	Error            string `json:"error"`
	TimeCreated      int    `json:"time_created"`
	TimeExpires      int    `json:"time_expires"`
	Status           int    `json:"status"`
	StatusText       string `json:"status_text"`
	Type             string `json:"type"`
	Coin             string `json:"coin"`
	Amount           int    `json:"amount"`
	Amountf          string `json:"amountf"`
	Received         int    `json:"received"`
	Recievedf        string `json:"recievedf"`
	RecievedConfirms int    `json:"recv_confirms"`
	PaymentAddress   string `json:"payment_address"`
}

//GetTxInfoMulti calls the "get_tx_info_multi" command
func (c *Client) GetTxInfoMulti(txid string, optionals ...OptionalValue) (*getTxInfoMultiResponse, error) {
	values := &url.Values{}
	values.Set("txid", txid)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getTxInfoMultiResponse `json:"result"`
	}

	if err := c.call("get_tx_info_multi", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getTxIdsResponse is the api response of a "get_tx_ids" call
type getTxIdsResponse []string

//GetTxIds calls the "get_tx_ids" command
func (c *Client) GetTxIds(optionals ...OptionalValue) (*getTxIdsResponse, error) {
	values := &url.Values{}
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getTxIdsResponse `json:"result"`
	}
	if err := c.call("get_tx_ids", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
