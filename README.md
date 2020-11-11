# CoinPayments [![PkgGoDev](https://pkg.go.dev/badge/github.com/aidenesco/coinpayments)](https://pkg.go.dev/github.com/aidenesco/coinpayments) [![Go Report Card](https://goreportcard.com/badge/github.com/aidenesco/coinpayments)](https://goreportcard.com/report/github.com/aidenesco/coinpayments)

This package is a Coinpayments.net client library. See official documentation [here](https://www.coinpayments.net/apidoc-intro)

## Installation

```sh
go get -u github.com/aidenesco/coinpayments
```

## Usage

```go
import "github.com/aidenesco/coinpayments"

client := coinpayments.NewClient("public key", "private key")

balances, _ := client.Balances()

fmt.Println(balances)
```

