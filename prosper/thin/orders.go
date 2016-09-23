package thin

import (
	"bytes"
	"encoding/json"
)

type (
	BidRequest struct {
		ListingId int64   `json:"listing_id"`
		BidAmount float64 `json:"bid_amount"`
	}

	orderParams struct {
		BidRequests []BidRequest `json:"bid_requests"`
	}

	BidStatus struct {
		BidRequest
		Status          string  `json:"bid_status"`
		BidResult       string  `json:"bid_result"`
		BidAmountPlaced float64 `json:"bid_amount_placed"`
	}

	OrderResponse struct {
		//TODO: Add support for effective_yield, estimated_loss, estimated_return,
		// source
		OrderId     string      `json:"order_id"`
		BidStatus   []BidStatus `json:"bid_requests"`
		OrderStatus string      `json:"order_status"`
		OrderDate   string      `json:"order_date"`
	}
)

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

func (c Client) OrderStatus(orderId string) (response OrderResponse, err error) {
	err = c.DoRequest("GET", c.baseUrl+"/orders/"+orderId, nil, &response)
	if err != nil {
		return OrderResponse{}, err
	}
	return response, nil
}
