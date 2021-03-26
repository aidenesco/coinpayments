package coinpayments

import "net/url"

//getPBNInfoResponse is the api response of a "get_pbn_info" call
type getPBNInfoResponse struct {
	PBNTag       string `json:"pbntag"`
	Merchant     string `json:"merchant"`
	ProfileName  string `json:"profile_name"`
	ProfileURL   string `json:"profile_url"`
	ProfileEmail string `json:"profile_email"`
	ProfileImage string `json:"profile_image"`
	MemberSince  int    `json:"member_since"`
	Feedback     struct {
		Positive int    `json:"pos"`
		Negative int    `json:"neg"`
		Neutral  string `json:"neut"`
		Total    int    `json:"total"`
		Percent  string `json:"percent"`
	} `json:"feedback"`
}

//GetPBNInfo calls the "get_pbn_info" command
func (c *Client) GetPBNInfo(pbntag string, optionals ...OptionalValue) (*getPBNInfoResponse, error) {
	values := &url.Values{}
	values.Set("pbntag", pbntag)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getPBNInfoResponse `json:"result"`
	}
	if err := c.call("get_pbn_info", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//getPBNListResponse is the api response of a "get_pbn_list" call
type getPBNListResponse []struct {
	TagID       string `json:"tagid"`
	PBGTag      string `json:"pbgtag"`
	TimeExpires int    `json:"time_expires"`
}

//GetPBNList calls the "get_pbn_list" command
func (c *Client) GetPBNList(optionals ...OptionalValue) (*getPBNListResponse, error) {
	values := &url.Values{}
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *getPBNListResponse `json:"result"`
	}
	if err := c.call("get_pbn_list", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//buyPBNTagsResponse is the api response of a "buy_pbn_tags" call
type buyPBNTagsResponse []interface{}

//BuyPBNTags calls the "buy_pbn_tags" command
func (c *Client) BuyPBNTags(coin string, optionals ...OptionalValue) (*buyPBNTagsResponse, error) {
	values := &url.Values{}
	values.Set("coin", coin)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *buyPBNTagsResponse `json:"result"`
	}
	if err := c.call("buy_pbn_tags", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//claimPBNTagResponse is the api response of a "claim_pbn_tag" call
type claimPBNTagResponse []interface{}

//ClaimPBNTag calls the "claim_pbn_tag" command
func (c *Client) ClaimPBNTag(tagid, name string, optionals ...OptionalValue) (*claimPBNTagResponse, error) {
	values := &url.Values{}
	values.Set("tagid", tagid)
	values.Set("name", name)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *claimPBNTagResponse `json:"result"`
	}
	if err := c.call("claim_pbn_tag", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//updatePBNTagResponse is the api response of a "update_pbn_tag" call
type updatePBNTagResponse []interface{}

//UpdatePBNTag calls the "update_pbn_tag" command
func (c *Client) UpdatePBNTag(tagid string, optionals ...OptionalValue) (*updatePBNTagResponse, error) {
	values := &url.Values{}
	values.Set("tagid", tagid)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *updatePBNTagResponse `json:"result"`
	}
	if err := c.call("update_pbn_tag", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//renewPBNTagResponse is the api response of a "renew_pbn_tag" call
type renewPBNTagResponse []interface{}

//RenewPBNTag calls the "renew_pbn_tag" command
func (c *Client) RenewPBNTag(tagid, coin string, optionals ...OptionalValue) (*renewPBNTagResponse, error) {
	values := &url.Values{}
	values.Set("tagid", tagid)
	values.Set("coin", coin)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *renewPBNTagResponse `json:"result"`
	}
	if err := c.call("renew_pbn_tag", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//deletePBNTagResponse is the api response of a "delete_pbn_tag" call
type deletePBNTagResponse []interface{}

//DeletePBNTag calls the "delete_pbn_tag" command
func (c *Client) DeletePBNTag(tagid string, optionals ...OptionalValue) (*deletePBNTagResponse, error) {
	values := &url.Values{}
	values.Set("tagid", tagid)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *deletePBNTagResponse `json:"result"`
	}
	if err := c.call("delete_pbn_tag", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

//claimPBNCouponResponse is the api response of a "claim_pbn_coupon" call
type claimPBNCouponResponse struct {
	TagID string `json:"tagid"`
}

//ClaimPBNCoupon calls the "claim_pbn_coupon" command
func (c *Client) ClaimPBNCoupon(coupon string, optionals ...OptionalValue) (*claimPBNCouponResponse, error) {
	values := &url.Values{}
	values.Set("coupon", coupon)
	addOptionals(optionals, values)

	var resp struct {
		errResponse
		Result *claimPBNCouponResponse `json:"result"`
	}
	if err := c.call("claim_pbn_coupon", values, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
