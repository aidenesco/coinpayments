package coinpayments

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//IPN holds the data and type from an IPN
type IPN struct {
	ipnInformation
	depositInformation            depositInformation
	withdrawalInformation         withdrawalInformation
	buyerInformation              buyerInformation
	shippingInformation           shippingInformation
	simpleButtonFields            simpleButtonFields
	advancedButtonFields          advancedButtonFields
	shoppingCartButtonFields      shoppingCartButtonFields
	donationButtonFields          donationButtonFields
	apiGeneratedTransactionFields apiGeneratedTransactionFields
}

type ipnInformation struct {
	IPNVersion string
	IPNType    string
	IPNMode    string
	IPNId      string
	Merchant   string
}

type simpleIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	simpleButtonFields
}

//ToSimpleIPN returns the data from the "simple" ipn type
func (i *IPN) ToSimpleIPN() (*simpleIPN, error) {
	if i.IPNType != "simple" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'simple'")
	}
	return &simpleIPN{
		ipnInformation:      i.ipnInformation,
		buyerInformation:    i.buyerInformation,
		shippingInformation: i.shippingInformation,
		simpleButtonFields:  i.simpleButtonFields,
	}, nil
}

type buttonIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	advancedButtonFields
}

//ToButtonIPN returns the data from the "button" ipn type
func (i *IPN) ToButtonIPN() (*buttonIPN, error) {
	if i.IPNType != "button" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'button'")
	}
	return &buttonIPN{
		ipnInformation:       i.ipnInformation,
		buyerInformation:     i.buyerInformation,
		shippingInformation:  i.shippingInformation,
		advancedButtonFields: i.advancedButtonFields,
	}, nil
}

type cartIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	shoppingCartButtonFields
}

//ToCartIPN returns the data from the "cart" ipn type
func (i *IPN) ToCartIPN() (*cartIPN, error) {
	if i.IPNType != "cart" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'cart'")
	}
	return &cartIPN{
		ipnInformation:           i.ipnInformation,
		buyerInformation:         i.buyerInformation,
		shippingInformation:      i.shippingInformation,
		shoppingCartButtonFields: i.shoppingCartButtonFields,
	}, nil
}

type donationIPN struct {
	ipnInformation
	buyerInformation
	shippingInformation
	donationButtonFields
}

//ToDonationIPN returns the data from the "donation" ipn type
func (i *IPN) ToDonationIPN() (*donationIPN, error) {
	if i.IPNType != "donation" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'donation'")
	}
	return &donationIPN{
		ipnInformation:       i.ipnInformation,
		buyerInformation:     i.buyerInformation,
		shippingInformation:  i.shippingInformation,
		donationButtonFields: i.donationButtonFields,
	}, nil
}

type depositIPN struct {
	ipnInformation
	depositInformation
}

//ToDepositIPN returns the data from the "deposit" ipn type
func (i *IPN) ToDepositIPN() (*depositIPN, error) {
	if i.IPNType != "deposit" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'deposit'")
	}
	return &depositIPN{
		ipnInformation:     i.ipnInformation,
		depositInformation: i.depositInformation,
	}, nil
}

type withdrawalIPN struct {
	ipnInformation
	withdrawalInformation
}

//ToWithdrawalIPN returns the data from the "withdrawal" ipn type
func (i *IPN) ToWithdrawalIPN() (*withdrawalIPN, error) {
	if i.IPNType != "withdrawal" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'withdrawal'")
	}
	return &withdrawalIPN{
		ipnInformation:        i.ipnInformation,
		withdrawalInformation: i.withdrawalInformation,
	}, nil
}

type apiIPN struct {
	ipnInformation
	apiGeneratedTransactionFields
}

//ToApiIPN returns the data from the "api" ipn type
func (i *IPN) ToApiIPN() (*apiIPN, error) {
	if i.IPNType != "api" {
		return nil, fmt.Errorf("coinpayments: IPN type not 'api'")
	}
	return &apiIPN{
		ipnInformation:                i.ipnInformation,
		apiGeneratedTransactionFields: i.apiGeneratedTransactionFields,
	}, nil
}

//ParseIPN takes a http request and parses the IPN information from it
func (c *Client) ParseIPN(r *http.Request) (*IPN, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("coinpayments: error reading request body - %v", err)
	}

	if c.ipnSecret != "" {
		hmac := r.Header.Get("HMAC")

		genHMAC, err := c.makeIPNHMAC(string(data))
		if err != nil {
			return nil, fmt.Errorf("coinpayments: error generating ipn HMAC - %v", err)
		}

		if hmac != genHMAC {
			return nil, fmt.Errorf("coinpayments: could not validate server HMAC")
		}
	}

	values, err := url.ParseQuery(string(data))
	if err != nil {
		return nil, err
	}

	ipn := &IPN{
		ipnInformation: ipnInformation{
			IPNVersion: values.Get("ipn_version"),
			IPNType:    values.Get("ipn_type"),
			IPNMode:    values.Get("ipn_mode"),
			IPNId:      values.Get("ipn_id"),
			Merchant:   values.Get("merchant"),
		},
	}

	switch ipn.IPNType {
	case "simple":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.simpleButtonFields = simpleButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			Net:              values.Get("net"),
			ItemAmount:       values.Get("item_amount"),
			ItemName:         values.Get("item_name"),
			ItemDescription:  values.Get("item_desc"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Option1Name:      values.Get("on1"),
			Option1Value:     values.Get("ov1"),
			Option2Name:      values.Get("on2"),
			Option2Value:     values.Get("ov2"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "button":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.advancedButtonFields = advancedButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			Net:              values.Get("net"),
			ItemAmount:       values.Get("item_amount"),
			ItemName:         values.Get("item_name"),
			Quantity:         values.Get("quantity"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Option1Name:      values.Get("on1"),
			Option1Value:     values.Get("ov1"),
			Option2Name:      values.Get("on2"),
			Option2Value:     values.Get("ov2"),
			Extra:            values.Get("extra"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "cart":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.shoppingCartButtonFields = shoppingCartButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			ItemName:         values.Get("item_name_#"),
			ItemAmount:       values.Get("item_amount_#"),
			ItemQuantity:     values.Get("item_quantity_#"),
			ItemNumber:       values.Get("item_number_#"),
			Option1Name:      values.Get("item_on1_#"),
			Option1Value:     values.Get("item_ov1_#"),
			Option2Name:      values.Get("item_on2_#"),
			Option2Value:     values.Get("item_ov2_#"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Extra:            values.Get("extra"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "donation":
		ipn.buyerInformation = buyerInformation{
			FirstName: values.Get("first_name"),
			LastName:  values.Get("last_name"),
			Company:   values.Get("company"),
			Email:     values.Get("email"),
		}

		ipn.shippingInformation = shippingInformation{
			Address1:    values.Get("address1"),
			Address2:    values.Get("address2"),
			City:        values.Get("city"),
			State:       values.Get("state"),
			ZipCode:     values.Get("zip"),
			Country:     values.Get("country"),
			CountryName: values.Get("country_name"),
			Phone:       values.Get("phone"),
		}

		ipn.donationButtonFields = donationButtonFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Subtotal:         values.Get("subtotal"),
			Shipping:         values.Get("shipping"),
			Tax:              values.Get("tax"),
			Fee:              values.Get("fee"),
			Net:              values.Get("net"),
			ItemName:         values.Get("item_name"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			Option1Name:      values.Get("on1"),
			Option1Value:     values.Get("ov1"),
			Option2Name:      values.Get("on2"),
			Option2Value:     values.Get("ov2"),
			Extra:            values.Get("extra"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	case "deposit":
		ipn.depositInformation = depositInformation{
			TransactionID: values.Get("txn_id"),
			Address:       values.Get("address"),
			DestTag:       values.Get("dest_tag"),
			Status:        values.Get("status"),
			StatusText:    values.Get("status_text"),
			Currency:      values.Get("currency"),
			Confirms:      values.Get("confirms"),
			Amount:        values.Get("amount"),
			Amounti:       values.Get("amounti"),
			Fee:           values.Get("fee"),
			Feei:          values.Get("feei"),
			FiatCoin:      values.Get("fiat_coin"),
			FiatAmount:    values.Get("fiat_amount"),
			FiatAmounti:   values.Get("fiat_amounti"),
			FiatFee:       values.Get("fiat_fee"),
			FiatFeei:      values.Get("fiat_feei"),
		}
	case "withdrawal":
		ipn.withdrawalInformation = withdrawalInformation{
			ID:            values.Get("id"),
			Status:        values.Get("status"),
			StatusText:    values.Get("status_text"),
			Address:       values.Get("address"),
			TransactionID: values.Get("txn_id"),
			Currency:      values.Get("currency"),
			Amount:        values.Get("amount"),
			Amounti:       values.Get("amounti"),
		}
	case "api":
		ipn.apiGeneratedTransactionFields = apiGeneratedTransactionFields{
			Status:           values.Get("status"),
			StatusText:       values.Get("status_text"),
			TransactionID:    values.Get("txn_id"),
			Currency1:        values.Get("currency1"),
			Currency2:        values.Get("currency2"),
			Amount1:          values.Get("amount1"),
			Amount2:          values.Get("amount2"),
			Fee:              values.Get("fee"),
			BuyerName:        values.Get("buyer_name"),
			Email:            values.Get("email"),
			ItemName:         values.Get("item_name"),
			ItemNumber:       values.Get("item_number"),
			Invoice:          values.Get("invoice"),
			Custom:           values.Get("custom"),
			SendTransaction:  values.Get("send_tx"),
			ReceivedAmount:   values.Get("received_amount"),
			ReceivedConfirms: values.Get("received_confirms"),
		}
	}

	return ipn, nil
}

type depositInformation struct {
	TransactionID string
	Address       string
	DestTag       string
	Status        string
	StatusText    string
	Currency      string
	Confirms      string
	Amount        string
	Amounti       string
	Fee           string
	Feei          string
	FiatCoin      string
	FiatAmount    string
	FiatAmounti   string
	FiatFee       string
	FiatFeei      string
}

type withdrawalInformation struct {
	ID            string
	Status        string
	StatusText    string
	Address       string
	TransactionID string
	Currency      string
	Amount        string
	Amounti       string
}

type buyerInformation struct {
	FirstName string
	LastName  string
	Company   string
	Email     string
}

type shippingInformation struct {
	Address1    string
	Address2    string
	City        string
	State       string
	ZipCode     string
	Country     string
	CountryName string
	Phone       string
}

type simpleButtonFields struct {
	Status           string
	StatusText       string
	TransactionID    string
	Currency1        string
	Currency2        string
	Amount1          string
	Amount2          string
	Subtotal         string
	Shipping         string
	Tax              string
	Fee              string
	Net              string
	ItemAmount       string
	ItemName         string
	ItemDescription  string
	ItemNumber       string
	Invoice          string
	Custom           string
	Option1Name      string
	Option1Value     string
	Option2Name      string
	Option2Value     string
	SendTransaction  string
	ReceivedAmount   string
	ReceivedConfirms string
}

type advancedButtonFields struct {
	Status           string
	StatusText       string
	TransactionID    string
	Currency1        string
	Currency2        string
	Amount1          string
	Amount2          string
	Subtotal         string
	Shipping         string
	Tax              string
	Fee              string
	Net              string
	ItemAmount       string
	ItemName         string
	Quantity         string
	ItemNumber       string
	Invoice          string
	Custom           string
	Option1Name      string
	Option1Value     string
	Option2Name      string
	Option2Value     string
	Extra            string
	SendTransaction  string
	ReceivedAmount   string
	ReceivedConfirms string
}

type shoppingCartButtonFields struct {
	Status           string
	StatusText       string
	TransactionID    string
	Currency1        string
	Currency2        string
	Amount1          string
	Amount2          string
	Subtotal         string
	Shipping         string
	Tax              string
	Fee              string
	ItemName         string
	ItemAmount       string
	ItemQuantity     string
	ItemNumber       string
	Option1Name      string
	Option1Value     string
	Option2Name      string
	Option2Value     string
	Invoice          string
	Custom           string
	Extra            string
	SendTransaction  string
	ReceivedAmount   string
	ReceivedConfirms string
}

type donationButtonFields struct {
	Status           string
	StatusText       string
	TransactionID    string
	Currency1        string
	Currency2        string
	Amount1          string
	Amount2          string
	Subtotal         string
	Shipping         string
	Tax              string
	Fee              string
	Net              string
	ItemName         string
	ItemNumber       string
	Invoice          string
	Custom           string
	Option1Name      string
	Option1Value     string
	Option2Name      string
	Option2Value     string
	Extra            string
	SendTransaction  string
	ReceivedAmount   string
	ReceivedConfirms string
}

type apiGeneratedTransactionFields struct {
	Status           string
	StatusText       string
	TransactionID    string
	Currency1        string
	Currency2        string
	Amount1          string
	Amount2          string
	Fee              string
	BuyerName        string
	Email            string
	ItemName         string
	ItemNumber       string
	Invoice          string
	Custom           string
	SendTransaction  string
	ReceivedAmount   string
	ReceivedConfirms string
}
