package coinpayments

import "net/url"

//balancesResponse is the api response of a "balances" call
type balancesResponse map[string]struct {
	Balance  int    `json:"balance"`
	Balancef string `json:"balancef"`
	Status   string `json:"status"`
}

//Balances calls the "balances" command
func (c *Client) Balances(optionals ...OptionalValue) (*balancesResponse, error) {
	values := &url.Values{}
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *balancesResponse `json:"result"`
	}
	if err := c.call("balances", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getDepositAddressResponse is the api response of a "get_deposit_address" call
type getDepositAddressResponse struct {
	Address string `json:"address"`
	PubKey  string `json:"pubkey"`
	DestTag int    `json:"dest_tag"`
}

//GetDepositAddress calls the "get_deposit_address" command
func (c *Client) GetDepositAddress(currency string, optionals ...OptionalValue) (*getDepositAddressResponse, error) {
	values := &url.Values{}
	values.Set("currency", currency)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getDepositAddressResponse `json:"result"`
	}
	if err := c.call("get_deposit_address", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//createTransferResponse is the api response of a "create_transfer" call
type createTransferResponse struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

//CreateTransfer calls the "create_transfer" command
func (c *Client) CreateTransfer(amount, currency string, optionals ...OptionalValue) (*createTransferResponse, error) {
	values := &url.Values{}
	values.Set("amount", amount)
	values.Set("currency", currency)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *createTransferResponse `json:"result"`
	}

	if err := c.call("create_transfer", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//createWithdrawalResponse is the api response of a "create_withdrawal" call
type createWithdrawalResponse struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Amount string `json:"amount"`
}

//CreateWithdrawal calls the "create_withdrawal" command
func (c *Client) CreateWithdrawal(amount, currency string, optionals ...OptionalValue) (*createWithdrawalResponse, error) {
	values := &url.Values{}
	values.Set("amount", amount)
	values.Set("currency", currency)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *createWithdrawalResponse `json:"result"`
	}
	if err := c.call("create_withdrawal", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//cancelWithdrawalResponse is the api response of a "cancel_withdrawal" call
type cancelWithdrawalResponse struct{}

//CancelWithdrawal calls the "create_withdrawal" command
func (c *Client) CancelWithdrawal(id string, optionals ...OptionalValue) (*cancelWithdrawalResponse, error) {
	values := &url.Values{}
	values.Set("id", id)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *cancelWithdrawalResponse `json:"result"`
	}
	if err := c.call("cancel_withdrawal", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//convertResponse is the api response of a "convert" call
type convertResponse struct {
	ID string `json:"id"`
}

//Convert calls the "convert" command
func (c *Client) Convert(amount, from, to string, optionals ...OptionalValue) (*convertResponse, error) {
	values := &url.Values{}
	values.Set("amount", amount)
	values.Set("from", from)
	values.Set("to", to)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *convertResponse `json:"result"`
	}
	if err := c.call("convert", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//convertLimitsResponse is the api response of a "convert_limits" call
type convertLimitsResponse struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

//ConvertLimits calls the "convert_limits" command
func (c *Client) ConvertLimits(from, to string, optionals ...OptionalValue) (*convertLimitsResponse, error) {
	values := &url.Values{}
	values.Set("from", from)
	values.Set("to", to)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *convertLimitsResponse `json:"result"`
	}
	if err := c.call("convert_limits", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getWithdrawalHistoryResponse is the api response of a "get_withdrawal_history" call
type getWithdrawalHistoryResponse []struct {
	ID          string `json:"id"`
	TimeCreated int    `json:"time_created"`
	Status      int    `json:"status"`
	StatusText  string `json:"status_text"`
	Coin        string `json:"coin"`
	Amount      int    `json:"amount"`
	Amountf     string `json:"amountf"`
	Note        string `json:"note"`
	SendAddress string `json:"send_address"`
	SendDestTag string `json:"send_dest_tag"`
	SendTXID    string `json:"send_txid"`
}

//GetWithdrawalHistory calls the "get_withdrawal_history" command
func (c *Client) GetWithdrawalHistory(optionals ...OptionalValue) (*getWithdrawalHistoryResponse, error) {
	values := &url.Values{}
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getWithdrawalHistoryResponse `json:"result"`
	}
	if err := c.call("get_withdrawal_history", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getWithdrawalInfoResponse is the api response of a "get_withdrawal_info" call
type getWithdrawalInfoResponse struct {
	TimeCreated int    `json:"time_created"`
	Status      int    `json:"status"`
	StatusText  string `json:"status_text"`
	Coin        string `json:"coin"`
	Amount      int    `json:"amount"`
	Amountf     string `json:"amountf"`
	Note        string `json:"note"`
	SendAddress string `json:"send_address"`
	SendTXID    string `json:"send_txid"`
}

//GetWithdrawalInfo calls the "get_withdrawal_info" command
func (c *Client) GetWithdrawalInfo(id string, optionals ...OptionalValue) (*getWithdrawalInfoResponse, error) {
	values := &url.Values{}
	values.Set("id", id)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getWithdrawalInfoResponse `json:"result"`
	}
	if err := c.call("get_withdrawal_info", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getConversionInfoResponse is the api response of a "get_conversion_info" call
type getConversionInfoResponse struct {
	TimeCreated string `json:"time_created"`
	Status      int    `json:"status"`
	StatusText  string `json:"status_text"`
	Coin1       string `json:"coin1"`
	Coin2       string `json:"coin2"`
	AmountSent  int    `json:"amount_sent"`
	AmountSentf string `json:"amount_sentf"`
	Received    int    `json:"received"`
	Receivedf   string `json:"receivedf"`
}

//GetConversionInfo calls the "get_conversion_info" command
func (c *Client) GetConversionInfo(id string, optionals ...OptionalValue) (*getConversionInfoResponse, error) {
	values := &url.Values{}
	values.Set("id", id)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getConversionInfoResponse `json:"result"`
	}

	if err := c.call("get_conversion_info", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
