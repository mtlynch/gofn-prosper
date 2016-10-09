# Go Forth 'n Prosper (gofn-prosper)

[![Build Status](https://travis-ci.org/mtlynch/gofn-prosper.svg?branch=master)](https://travis-ci.org/mtlynch/gofn-prosper)
[![Coverage Status](https://coveralls.io/repos/github/mtlynch/gofn-prosper/badge.svg?branch=master)](https://coveralls.io/github/mtlynch/gofn-prosper?branch=master)
[![GoDoc](https://godoc.org/github.com/mtlynch/gofn-prosper?status.svg)](https://godoc.org/github.com/mtlynch/gofn-prosper)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtlynch/gofn-prosper)](https://goreportcard.com/report/github.com/mtlynch/gofn-prosper)

## Overview

**Go Forth 'n Prosper** is a set of API Bindings for the Prosper Peer to Peer Lending API.

Currently, about 80% of the functionality is implemented.

## Example Usage

For an example of a project that uses gofn-prosper, see:

 * [ProsperBot](https://github.com/mtlynch/prosperbot) - An automated note buyer and account tracker.

## API Documentation

The sections below give a high-level overview of some gofn-propser features. For full details, refer to the [gofn-prosper GoDoc](https://godoc.org/github.com/mtlynch/gofn-prosper).

### Creating a Client

All Prosper actions take place through a `Client` interface. To create a `Client` instance, do the following: 

```go
import (
  "github.com/mtlynch/gofn-prosper/prosper"
  "github.com/mtlynch/gofn-prosper/prosper/auth"
)

client := prosper.NewClient(auth.ClientCredentials{
  ClientID:     "your client ID",
  ClientSecret: "your client secret",
  Username:     "your Prosper username",
  Password:     "your Prosper password",
  })
```

`ClientID` and `ClientSecret` are the values Prosper assigns to you when you generate OAuth credentials on the [OAuth Settings](https://www.prosper.com/oauth#/settings) page.

*Note*: If it seems strange to you that you need to enter your OAuth credentials **and** your username/password, it is strange, but this is a requirement of the Prosper API.

### Account Information

The [`Account`](https://godoc.org/github.com/mtlynch/gofn-prosper/prosper#Client.Account) API allows clients to retrieve information about their account, such as available balance and account value:

```go
account, err := client.Account(prosper.AccountParams{})
if err != nil {
  fmt.Printf("Failed to retrieve account information: %v", err)
  return
}
fmt.Printf("Your account has $%.2f in cash and a total value of $%.2f\n",
  account.AvailableCashBalance, account.TotalAccountValue)
```

```text
Your account has $250.00 in cash and a total value of $1893.91
```

### Searching Available Notes

The `Search` API allows clients to search the available listings, filtering by listing criteria.

This snippet shows how to search for loans with prosper ratings of A or B, where the borrower earns between $50,000 and $75,000 per year:

```go
searchResp, err := client.Search(prosper.SearchParams{
  Offset: 0,
  Limit:  5,
  Filter: prosper.SearchFilter{
    Rating:      []prosper.Rating{prosper.RatingA, prosper.RatingB},
    IncomeRange: []prosper.IncomeRange{prosper.Between50kAnd75k},
  },
})
if err != nil {
  fmt.Printf("Failed to search available note listings: %v\n", err)
  return
}
fmt.Printf("Found %d matching notes, showing first %d\n",
  searchResp.TotalCount, searchResp.ResultCount)
for i, listing := range searchResp.Results {
  fmt.Printf("%2d: ID: %v  Loan Amount: $%5.0f  Yield: %.2f%%\n",
    i+1, listing.ListingNumber, listing.ListingAmount,
    listing.EffectiveYield*100.0)
```

```text
Found 27 matching notes, showing first 5
 1: ID: 5492410  Loan Amount: $15000  Yield: 9.01%
 2: ID: 5302541  Loan Amount: $14700  Yield: 9.70%
 3: ID: 5515002  Loan Amount: $15500  Yield: 7.99%
 4: ID: 5298533  Loan Amount: $20000  Yield: 6.85%
 5: ID: 5511744  Loan Amount: $ 6000  Yield: 8.58%
```

### Buying a Note

The `PlaceBid` API allows clients to make a bid on a Prosper listing.

```go
orderResp, err := client.PlaceBid(prosper.BidRequest{
  ListingID: 5492410,
  BidAmount: 25.0,
})
if err != nil {
  fmt.Printf("Failed to place bid: %v\n", err)
  return
}
fmt.Printf("Successfully placed order. Order ID: %v\n", orderResp.OrderID)
```

```text
Successfully placed order. Order ID: a9d4b52b-34cb-4112-bf57-671a18efecdd
```

### Checking Order Status

```go
orderID := prosper.OrderID("a9d4b52b-34cb-4112-bf57-671a18efecdd")
orderResp, err := client.OrderStatus(orderID)
if err != nil {
  fmt.Printf("Failed to check order status: %v\n", err)
  return
}
if orderResp.OrderStatus == prosper.OrderCompleted {
  fmt.Printf("Order %v is complete\n", orderID)
} else {
  fmt.Printf("Order %v is in progress\n", orderID)
}
```

```text
Order a9d4b52b-34cb-4112-bf57-671a18efecdd is complete
```

## Relationship to Prosper, Inc.
**Go Forth 'n Prosper** is not affiliated with Prosper, Inc. It is released as an independently maintained library for interacting with Prosper's public API.

## License

Go Forth 'n Propser is released under the [Apache 2 License](LICENSE).
