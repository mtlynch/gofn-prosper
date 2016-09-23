package thin

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPlaceBidSuccessfulResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/orders/",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testContentType(t, r, "application/json")
			fmt.Fprint(w, `{
	"order_id":"067e6162-3b6f-4ae2-a171-2470b63dff00",
				"bid_requests": [
					{
						"listing_id": 215032,
						"bid_status": "PENDING",
						"bid_amount": 32
					}
				],
				"effective_yield": 0.0842,
				"estimated_loss": 0.0324,
				"estimated_return": 0.0518,
				"source": "API",
				"order_status": "IN_PROGRESS",
				"order_date": "2015-09-17 19:54:58 +0000"
}`)
		},
	)

	client := Client{
		baseUrl:      server.URL,
		tokenManager: mockTokenManager{},
	}
	got, err := client.PlaceBid([]BidRequest{
		{215032, 32},
	})
	if err != nil {
		t.Fatalf("client.PlaceBid failed: %v", err)
	}

	want := OrderResponse{
		OrderId: "067e6162-3b6f-4ae2-a171-2470b63dff00",
		BidStatus: []BidStatus{
			{
				BidRequest: BidRequest{
					ListingId: 215032,
					BidAmount: 32.0,
				},
				Status: "PENDING",
			},
		},
		OrderStatus: "IN_PROGRESS",
		OrderDate:   "2015-09-17 19:54:58 +0000",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("client.PlaceBid returned %+v, want %+v", got, want)
	}
}

func TestPlaceBidServerError(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/orders/",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testContentType(t, r, "application/json")
			http.Error(w, "", 500)
			fmt.Fprint(w, `{
	"code":"SYS0001",
	"message":"Application Error"
}`)
		},
	)

	client := Client{
		baseUrl:      server.URL,
		tokenManager: mockTokenManager{},
	}
	errWant := errors.New(`request failed: 500 Internal Server Error - { "code":"SYS0001", "message":"Application Error" }`)
	_, errGot := client.PlaceBid([]BidRequest{
		{215032, 32},
	})
	if !reflect.DeepEqual(errGot, errWant) {
		t.Fatalf("got:\n%v\n, want:\n%v", errGot, errWant)
	}
}

func TestOrderStatusSuccessfulResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/orders/90cf709d-81d6-416a-89f2-ba6ab8146ef2",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{
	"order_id":"90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				"bid_requests": [
					{
						"listing_id": 2211270,
						"bid_amount": 100,
						"bid_status": "INVESTED",
						"bid_result": "BID_SUCCEEDED",
						"bid_amount_placed": 100
					}
				],
				"effective_yield": 0.0842,
				"estimated_loss": 0.0324,
				"estimated_return": 0.0518,
				"order_amount": 100,
				"order_amount_placed": 100,
				"order_amount_invested": 100,
				"source": "API",
				"order_status": "COMPLETED",
				"order_date": "2015-09-17 19:54:58 +0000"
}`)
		},
	)

	client := Client{
		baseUrl:      server.URL,
		tokenManager: mockTokenManager{},
	}
	got, err := client.OrderStatus("90cf709d-81d6-416a-89f2-ba6ab8146ef2")
	if err != nil {
		t.Fatalf("client.OrderStatus failed: %v", err)
	}

	want := OrderResponse{
		OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
		BidStatus: []BidStatus{
			{
				BidRequest: BidRequest{
					ListingId: 2211270,
					BidAmount: 100.0,
				},
				Status:          "INVESTED",
				BidResult:       "BID_SUCCEEDED",
				BidAmountPlaced: 100.0,
			},
		},
		OrderStatus: "COMPLETED",
		OrderDate:   "2015-09-17 19:54:58 +0000",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("client.OrderStatus returned %#v, want %#v", got, want)
	}
}
