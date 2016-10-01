package thin

import (
	"bytes"
	"encoding/json"
)

type (
	// BidRequest represents the JSON object passed to PlaceBid to place an order
	// for a listing.
	BidRequest struct {
		ListingId int64   `json:"listing_id"`
		BidAmount float64 `json:"bid_amount"`
	}

	orderParams struct {
		BidRequests []BidRequest `json:"bid_requests"`
	}

	// BidStatus represents part of the JSON response from the OrderStatus method.
	// It is the order status for a single bid request.
	BidStatus struct {
		BidRequest
		Status          string  `json:"bid_status"`
		BidResult       string  `json:"bid_result"`
		BidAmountPlaced float64 `json:"bid_amount_placed"`
	}

	// OrderResponse represents the full JSON response from the Prosper order
	// APIs.
	OrderResponse struct {
		//TODO: Add support for effective_yield, estimated_loss, estimated_return,
		// source
		OrderId     string      `json:"order_id"`
		BidStatus   []BidStatus `json:"bid_requests"`
		OrderStatus string      `json:"order_status"`
		OrderDate   string      `json:"order_date"`
	}
)

// PlaceBid places a bid for the given listing at the given bid amount. Wraps
// the Prosper POST /orders/ API described at:
// https://developers.prosper.com/docs/investor/orders-api/
func (c Client) PlaceBid(br []BidRequest) (response OrderResponse, err error) {
	reqBody, err := json.Marshal(orderParams{BidRequests: br})
	if err != nil {
		return OrderResponse{}, err
	}
	err = c.DoRequest("POST", c.baseUrl+"/orders/", bytes.NewReader(reqBody), &response)
	if err != nil {
		return OrderResponse{}, err
	}
	return response, nil
}

// OrderStatus retrieves the status of the given Propser Order ID. Wraps the
// Prosper /orders/{order_id}/listings API described at:
// https://developers.prosper.com/docs/investor/orders-api/
func (c Client) OrderStatus(orderId string) (response OrderResponse, err error) {
	err = c.DoRequest("GET", c.baseUrl+"/orders/"+orderId, nil, &response)
	if err != nil {
		return OrderResponse{}, err
	}
	return response, nil
}
